## Конфигурация

* Логин, пароль, хост от Эластики (`ELASTIC_USER`, `ELASTIC_PASSWORD`, `ELASTICSEARCH_ADDRESS`)
* Порт gRPC сервера (`PORT`)
* Дистанция в километрах поиска по координатам (`ELASTIC_DISTANCE`, по умолчанию 10км)

#### Все это можно посмотреть, изменить в .env

---

## Запуск Сервиса

### Запуск ElasticSearch
```
docker compose up -d
```

## Запуск gRPC server

* Для сборки необходим **GO** >= 1.19
* Приложение по умолчанию запущено на порту 50051

```
go mod tidy
```

```
go build -o main cmd/server/main.go 
```

```
./main
```

---

## Консольный клиент

### Проверка, что все работает

Для проверки можно поставить [evans](https://github.com/ktr0731/evans/releases/tag/v0.10.11).

Команда для подключения cli клиента:

```
evans -r repl --host localhost --port 50051
```

```
 package api
```

```
 service CitySearch
```

```
 call SearchByName
```

```
 call SearchByCoords
```

Последние две команды запрашивают параметры интерактивно:

* **SearchByName**  name: string
* **SearchByCoords** lat, long: float