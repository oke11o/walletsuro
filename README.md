# Walletsuro

Payment system with custom user wallets

## Использование паттерна Money

Используем pattern money. Поэтому `amount `

### Currency

Для MVP работаем только с USD

## Примеры запросов

В директории `http`

## Переменные окружения

PORT - на каком порту поднимать приложение

## POSTGRES_MULTIPLE_DATABASES

Для локально разработки удобно запускать интеграционные тесты с тестовой базой.
Но руками тестировать на локально базе. Поэтому создаем 2 базы данных.
Тестовая - для интеграционных тестов. Перед тестами всегда чистится.

# Logger

В данном приложении не реализован logging и tracing.
Используется logger из стандартной либы го. Пите в stdout.

# Event

1 операции соответствует 1 запись в таблице. Поэтому при получении Report'a придется делать запрос по кошельку From=wallet OR To=wallet. 

# Tools

Генерим в директорию ./bin все необходимые утилиты

- mockgen
- goswagger
- migrate

```makefile
make tools
```

# Linter

```makefile
make lint
```