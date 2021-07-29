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
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/oke11o/walletsuro/internal/config"
	"github.com/oke11o/walletsuro/internal/model"
	"github.com/oke11o/walletsuro/internal/repository"
)

type userRepoSuite struct {
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
	suite.Run(t, new(userRepoSuite))
}

func (rs *userRepoSuite) SetupSuite() {
	dbCfg, err := cfg()
	require.NoError(rs.T(), err)
	rs.cfg = dbCfg

	rs.T().Logf("DB_DSN: %s", dbCfg)

	dbx, err := sqlx.Open("postgres", dbCfg.PgDSN)
	require.NoError(rs.T(), err)
	rs.dbx = dbx

	// migrations
	driver, err := postgres.WithInstance(dbx.DB, &postgres.Config{})
	require.NoError(rs.T(), err)
	m, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations",
		dbCfg.DbName,
		driver,
	)
	require.NoError(rs.T(), err)

	// TODO: странное дело. Если надо выполнить миграции, то тут ок.
	// Если же миграции уже накатили, тут получаем os.ErrNotExist.
	// Стоит разобраться
	_ = m.Steps(2)
	//require.NoError(rs.T(), err)

	//service
	r, err := repository.New(dbCfg)
	require.NoError(rs.T(), err)
	rs.service = New(r)

	// fixtures
	rs.fixtures, err = testfixtures.New(
		testfixtures.Dialect("postgres"),
		testfixtures.Database(rs.dbx.DB),
		testfixtures.Directory("testdata/fixtures"),
	)
	rs.Require().NoError(err)
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

func (rs *userRepoSuite) SetupTest() {
	err := rs.fixtures.Load()
	rs.Require().NoError(err)
}

func (rs *userRepoSuite) Test_CreateWallet() {
	wallet, err := rs.service.CreateWallet(context.Background(), 1)
	rs.Require().NoError(err)
	expected := model.Wallet{
		UUID:   wallet.UUID,
		UserID: 1,
	}
	rs.Equal(expected, wallet)

	wallet2, err := rs.service.CreateWallet(context.Background(), 2)
	rs.Require().NoError(err)
	expected2 := model.Wallet{
		UUID:   wallet2.UUID,
		UserID: 2,
	}
	rs.Equal(expected2, wallet2)

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
	err = rs.dbx.Select(&events, sql, "create")
	rs.Require().NoError(err)
	rs.Require().Len(events, 3)
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
func (s *userRepoSuite) loadFixtures(dir string, data interface{}) {
	fxrPGSQL, err := testfixtures.New(
		testfixtures.Dialect("postgres"),
		testfixtures.Database(s.dbx.DB),
		testfixtures.Template(),
		testfixtures.TemplateData(map[string]interface{}{"Orders": data}),
		testfixtures.Directory(dir),
	)
	s.Require().NoError(err)

	err = fxrPGSQL.Load()
	s.Require().NoError(err)
}
