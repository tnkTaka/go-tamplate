DBNAME:=example

run:
	go run main.go

deps:
	which dep || go get -u github.com/golang/dep/cmd/dep
	dep ensure

migrate/init:
	mysql -u root -h localhost --protocol tcp -e "CREATE DATABASE \`$(DBNAME)\` DEFAULT CHARACTER SET utf8"

migrate/up:
	go run cmd/migrate.go