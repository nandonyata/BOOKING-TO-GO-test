# BOOKING-TO-GO-test

### Setup Alias
```shell
alias air='~/go/bin/air'        
alias migrate='~/go/bin/migrate'
```

### Create Migration
```shell
migrate create -ext sql -dir db/migration table_name
```

### up migration
```shell
migrate -database "postgres://postgres:postgres@localhost:5432/booking-to-go?sslmode=disable" -path db/migration up
```

### down migration
```bash
migrate -database "postgres://postgres:postgres@localhost:5432/booking-to-go?sslmode=disable" -path db/migration down
```