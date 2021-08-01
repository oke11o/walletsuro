package service

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.octolab.org/pointer"

	"github.com/oke11o/walletsuro/internal/config"
	"github.com/oke11o/walletsuro/internal/model"
	"github.com/oke11o/walletsuro/internal/repository"
)

type walletsuroSuite struct {
	suite.Suite
	dbx      *sqlx.DB
	fixtures *testfixtures.Loader
	service  *Service
	cfg      config.Config
}

func Test_RepoSuite(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	suite.Run(t, new(walletsuroSuite))
}

func (suite *walletsuroSuite) SetupSuite() {
	dbCfg, err := cfg()
	require.NoError(suite.T(), err)
	suite.cfg = dbCfg

	suite.T().Logf("DB_DSN: %s", dbCfg)

	dbx, err := sqlx.Open("postgres", dbCfg.PgDSN)
	require.NoError(suite.T(), err)
	suite.dbx = dbx

	// migrations
	driver, err := postgres.WithInstance(dbx.DB, &postgres.Config{})
	require.NoError(suite.T(), err)
	m, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations",
		dbCfg.DbName,
		driver,
	)
	require.NoError(suite.T(), err)

	// TODO: странное дело. Если надо выполнить миграции, то тут ок.
	// Если же миграции уже накатили, тут получаем os.ErrNotExist.
	// Стоит разобраться
	_ = m.Steps(2)
	//require.NoError(suite.T(), err)

	//service
	r, err := repository.New(dbCfg)
	require.NoError(suite.T(), err)
	suite.service = New(r)

	// fixtures
	suite.fixtures, err = testfixtures.New(
		testfixtures.Dialect("postgres"),
		testfixtures.Database(suite.dbx.DB),
		testfixtures.Directory("testdata/fixtures/base"),
	)
	suite.Require().NoError(err)
}

func cfg() (config.Config, error) {
	c, err := config.NewFromEnv()
	if err != nil {
		return c, err
	}
	if c.DbName == "" {
		c.DbName = "walletsuro_test"
	}
	if c.PgDSN == "" {
		c.PgDSN = fmt.Sprintf("host=localhost user=postgres password=postgres dbname=%s sslmode=disable", c.DbName)
	}
	return c, nil
}

func (suite *walletsuroSuite) SetupTest() {
	err := suite.fixtures.Load()
	suite.Require().NoError(err)
}

func (suite *walletsuroSuite) Test_CreateWallet() {
	wallet, err := suite.service.CreateWallet(context.Background(), 1)
	suite.Require().NoError(err)
	expected := model.Wallet{
		UUID:   wallet.UUID,
		UserID: 1,
	}
	suite.Equal(expected, wallet)

	wallet2, err := suite.service.CreateWallet(context.Background(), 2)
	suite.Require().NoError(err)
	expected2 := model.Wallet{
		UUID:   wallet2.UUID,
		UserID: 2,
	}
	suite.Equal(expected2, wallet2)

	// check events
	type event struct {
		ID               int64     `db:"id"`
		TargetWalletUUID string    `db:"target_wallet_uuid"`
		WalletUUID       *string   `db:"from_wallet_uuid"`
		Amount           int64     `db:"amount"`
		Date             time.Time `db:"date"`
		Type             string    `db:"type"`
	}
	var events []event
	sql := "SELECT id, from_wallet_uuid, target_wallet_uuid, amount, type, date FROM events WHERE type=$1"
	err = suite.dbx.Select(&events, sql, "create")
	suite.Require().NoError(err)
	suite.Require().Len(events, 4)
}

func (suite *walletsuroSuite) Test_Deposit() {
	amount := model.NewMoney(200, model.DefaultCurrency)
	UUID, err := uuid.Parse("81da6536-f03e-11eb-9a03-0242ac130003")
	suite.Require().NoError(err)

	wallet, err := suite.service.Deposit(context.Background(), 1, UUID, amount)
	suite.Require().NoError(err)
	expected := model.Wallet{
		UUID:   UUID,
		UserID: 1,
		Amount: model.NewMoney(300, model.DefaultCurrency),
	}
	suite.Equal(expected, wallet)

	wallet2, err := suite.service.Deposit(context.Background(), 1, UUID, amount)
	suite.Require().NoError(err)
	expected2 := model.Wallet{
		UUID:   UUID,
		UserID: 1,
		Amount: model.NewMoney(500, model.DefaultCurrency),
	}
	suite.Equal(expected2, wallet2)

	// check events
	type event struct {
		ID               int64     `db:"id"`
		TargetWalletUUID string    `db:"target_wallet_uuid"`
		WalletUUID       *string   `db:"from_wallet_uuid"`
		Amount           int64     `db:"amount"`
		Date             time.Time `db:"date"`
		Type             string    `db:"type"`
	}
	var events []event
	sql := "SELECT id, from_wallet_uuid, target_wallet_uuid, amount, type, date FROM events WHERE type=$1 ORDER BY id"
	err = suite.dbx.Select(&events, sql, "deposit")
	suite.Require().NoError(err)
	suite.Require().Len(events, 2)
	suite.Require().Equal([]event{
		{
			ID:               events[0].ID,
			TargetWalletUUID: "81da6536-f03e-11eb-9a03-0242ac130003",
			WalletUUID:       nil,
			Amount:           200,
			Date:             events[0].Date,
			Type:             "deposit",
		},
		{
			ID:               events[1].ID,
			TargetWalletUUID: "81da6536-f03e-11eb-9a03-0242ac130003",
			WalletUUID:       nil,
			Amount:           200,
			Date:             events[1].Date,
			Type:             "deposit",
		},
	}, events)
}

func (suite *walletsuroSuite) Test_Transfer() {
	amount := model.NewMoney(200, model.DefaultCurrency)
	fromUUID, err := uuid.Parse("50805aec-eef2-4130-995e-12dde9ef0c1a")
	suite.Require().NoError(err)
	toUUID, err := uuid.Parse("81da6536-f03e-11eb-9a03-0242ac130003")
	suite.Require().NoError(err)

	userID := int64(2)
	wallet, err := suite.service.Transfer(context.Background(), userID, fromUUID, toUUID, amount)
	suite.Require().NoError(err)
	expected := model.Wallet{
		UUID:   fromUUID,
		UserID: userID,
		Amount: model.NewMoney(145, model.DefaultCurrency),
	}
	suite.Equal(expected, wallet)

	_, err = suite.service.Transfer(context.Background(), userID, fromUUID, toUUID, amount)
	suite.Require().Error(err)

	// check events
	type event struct {
		ID               int64     `db:"id"`
		TargetWalletUUID string    `db:"target_wallet_uuid"`
		WalletUUID       *string   `db:"from_wallet_uuid"`
		Amount           int64     `db:"amount"`
		Date             time.Time `db:"date"`
		Type             string    `db:"type"`
	}
	var events []event
	sql := "SELECT id, from_wallet_uuid, target_wallet_uuid, amount, type, date FROM events WHERE type=$1 ORDER BY id"
	err = suite.dbx.Select(&events, sql, "transfer")
	suite.Require().NoError(err)
	suite.Require().Len(events, 1)
	suite.Require().Equal([]event{
		{
			ID:               events[0].ID,
			TargetWalletUUID: "50805aec-eef2-4130-995e-12dde9ef0c1a",
			WalletUUID:       pointer.ToString("81da6536-f03e-11eb-9a03-0242ac130003"),
			Amount:           200,
			Date:             events[0].Date,
			Type:             "transfer",
		},
	}, events)
}

func (suite *walletsuroSuite) Test_Report() {
	suite.loadCustomFixture("report")

	userID := int64(2)
	report, err := suite.service.Report(context.Background(), userID, nil, nil)
	suite.Require().NoError(err)
	expected := []model.ReportData{
		{
			WalletUUID: "50805aec-eef2-4130-995e-12dde9ef0c1a",
			Date:       "2021-07-29T18:16:08Z",
			Type:       "transfer",
			Amount:     "$2.00"},
		{
			WalletUUID: "81da6536-f03e-11eb-9a03-0242ac130003",
			Date:       "2021-07-29T18:37:08Z",
			Type:       "deposit",
			Amount:     "$2.00"},
		{
			WalletUUID: "81da6536-f03e-11eb-9a03-0242ac130003",
			Date:       "2021-07-29T18:37:08Z",
			Type:       "deposit",
			Amount:     "$2.00",
		},
	}
	suite.Len(report, 3)
	suite.Equal(expected, report)

}

func (suite *walletsuroSuite) loadCustomFixture(dir string) {

	fixtures, err := testfixtures.New(
		testfixtures.Dialect("postgres"),
		testfixtures.Database(suite.dbx.DB),
		testfixtures.Directory("testdata/fixtures/"+dir),
	)
	suite.Require().NoError(err)

	err = fixtures.Load()
	suite.Require().NoError(err)
}

// loadFixtures - На будущее. Чтобы можно было генерить фикстуры из go структур. Чтобы для разных тестов использовать один файл фикстур
// Use example
// 	type fixture struct {
//		StatusID int
//	}
//	s.loadFixtures("testdata/fixtures/orders/status_test", []*fixture{
//		{StatusID: 1},
//		{StatusID: 56},
//		{StatusID: -11},
//		{StatusID: 888},
//	})
// Yaml example
//{{range $order := $.Orders}}
//- updated_at: RAW=NOW()
//  type: 'vendor'
//  order_hash: RAW=MD5(random()::text)
//  vendor_id: 1
//  chain_id: 1
//  user_id: 1
//  status_id: 60
//  order_updated_at: RAW={{$order.OrderUpdatedAt}}
//{{end}}
func (suite *walletsuroSuite) loadFixtures(dir string, data interface{}) {
	fxrPGSQL, err := testfixtures.New(
		testfixtures.Dialect("postgres"),
		testfixtures.Database(suite.dbx.DB),
		testfixtures.Template(),
		testfixtures.TemplateData(map[string]interface{}{"Orders": data}),
		testfixtures.Directory(dir),
	)
	suite.Require().NoError(err)

	err = fxrPGSQL.Load()
	suite.Require().NoError(err)
}
