DBNAME:=example

run:
	go run main.go

migrate/init:
	mysql -u root -h localhost --protocol tcp -e "CREATE DATABASE \`$(DBNAME)\` DEFAULT CHARACTER SET utf8"

migrate/up:
	go run cmd/migrate.go