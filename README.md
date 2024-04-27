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

```
 package api
 service CitySearch
 call SearchByName
 call SearchByCoords
```

#### calls in evans inputs args intercatively