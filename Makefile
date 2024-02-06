create_migration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate:
	migrate -path=sql/migrations -database "postgres://root:rooter@localhost:5432/easy-shop?sslmode=disable" -verbose up  

migrate_down:
	migrate -path=sql/migrations -database "postgres://root:rooter@localhost:5432/easy-shop?sslmode=disable" -verbose down

.PHONY: migrate create_migration migrate_down
