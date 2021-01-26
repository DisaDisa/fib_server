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
sudo apt install -y protobuf-compiler
```
# Run
```
go mod download
go run main.go
```
# Usage
```
http://localhost:8181/get/X-Y where X and Y indexes of fibonacci sequence
```

# Todo
1. Tests
2. gRPC
3. Use big.Int