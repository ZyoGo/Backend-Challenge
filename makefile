migrate-up:
	migrate -path db/migrations -database "postgres://user:password123@localhost:5432/db_mooi_studio?sslmode=disable" up

migrate-down:	
	migrate -path db/migrations -database "postgres://user:password123@localhost:5432/db_mooi_studio?sslmode=disable" down
