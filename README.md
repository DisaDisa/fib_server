# Task

Фибоначчи
Реализовать сервис, возвращающий срез последовательности чисел из ряда Фибоначчи.

Сервис должен отвечать на запросы и возвращать ответ. В ответе должны быть перечислены все числа, последовательности Фибоначчи с порядковыми номерами от x до y.

Требования:
1. Требуется реализовать два протокола: HTTP REST и GRPC
2. Кэширование. Сервис не должен повторно вычислять числа из ряда Фибоначчи. Значения необходимо сохранить в Memcache.
3. Код должен быть покрыт тестами.

# Install
```
sudo apt install memcached  
```
(Optional) To build proto files:
```
sudo apt install protobuf-compiler
```
# Run
```
go mod download
go run main.go (from server folder)
```
# Usage
```
http://localhost:8181/get/X-Y where X and Y indexes of fibonacci sequence
```
To test gRPC (client folder):  
```
go run main.go X Y
```
# Rebuild proto (from grpcservice folder)
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpcservice.proto
```

# Todo
1. Tests
2. Use big.Int
3. Refactor