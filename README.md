```
docker compose up -d
```


```
go mod tidy
```
```
go build -o main cmd/server/main.go 
```
```
./main
```

Запущено на порту 50051
Для проверки можно поставить [evans](https://github.com/ktr0731/evans/releases/tag/v0.10.11) - там просто бинарник сразу идет

Команда для подключения cli клиента
```
evans -r repl --host localhost --port 50051
```

![изображение](https://github.com/Concerts-Mate/Elastic-Service/assets/28489754/8597e166-4504-491a-ab67-cb0c88b1b606)
