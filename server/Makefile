create migration file:
	migrate create -ext sql -dir db/migrations add_schema_tables

run migrations:
	docs:
		https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
	mysq/mariadb:
		migrate -path ./db/migrations -database "mysql://username:pw@tcp(localhost:3306)/database" up 1
	postgres:
		migrate -source file://db/migrations -database "postgres://postgres:pass@localhost:5432/shortly?sslmode=disable" -verbose up

