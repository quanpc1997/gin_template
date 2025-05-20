# Các thư viện đã dùng

## GORM
```sh
go get -u "gorm.io/driver/postgres"
go get -u "gorm.io/gorm"
```


## Migrate
Sử dụng ```golang-migrate```:
```sh
go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Tạo Migrate:
```sh
migrate create -ext sql -dir db/migrations -seq ten_file_migrate
```

Chạy migrate:
```sh
migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/ecommerce?sslmode=disable" up
```

Rollback migrate:
```sh
migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/ecommerce?sslmode=disable" down
```

Rollback to specific version:
```sh
migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/ecommerce?sslmode=disable" down -to 0
```
