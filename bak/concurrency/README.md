# concurrency-pattern

## Run

### DO1

```shell
go run ./app/main.go ./app/do1.go 2>/dev/null
```

### DO2

```shell
go run ./app/main.go ./app/do2.go 2>/dev/null
```

### DO3

```shell
MAX_QUEUES=1000000 go run ./app/main.go ./app/do3.go 2>/dev/null
```

### DO4

```shell
MAX_WORKERS=100 MAX_QUEUES=1000000 go run ./app/main.go ./app/do4.go 2>/dev/null
```

## Apache Bench

```shell
./ab.sh
```
