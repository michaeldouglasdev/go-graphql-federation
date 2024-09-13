create-migration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/md_petshop" -verbose up

migrate-down:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/md_petshop" -verbose down

# DOES NOT runs a file called "migrate", instead of runs the command migrate
.PHONY: migrate