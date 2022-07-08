# go-musthave-shortener-tpl
Шаблон репозитория для практического трека «Go в веб-разработке».

# Начало работы

1. Склонируйте репозиторий в любую подходящую директорию на вашем компьютере.
2. В корне репозитория выполните команду `go mod init <name>` (где `<name>` - адрес вашего репозитория на GitHub без префикса `https://`) для создания модуля.

# Обновление шаблона

Чтобы иметь возможность получать обновления автотестов и других частей шаблона выполните следующую команду:

```
git remote add -m main template https://github.com/yandex-praktikum/go-musthave-shortener-tpl.git
```

Для обновления кода автотестов выполните команду:

```
git fetch template && git checkout template/main .github
```

Затем добавьте полученные изменения в свой репозиторий.

## Инремент 1

1. Создание сервера по адресу "http://localhost:8080"
2. Обслуживающий 2 эндпоинта: POST / и GET /{id}
3. POST / в теле запроса принимает строку URL для сокращения
4. GET {id} принимает в качестве параметра идентификатор сокращенного URL, возвращает код 307 и оригинальный URL в HTTP-заголовке Location
5. Учесть некорректные запросы

Запуск проекта

```
go run cmd/shortener/main.go
```

Проверка эндпоинтов

```http request
POST http://localhost:8080
Content-Type: application/x-www-form-urlencoded

https://yandex.ru
```

```http request
GET http://localhost:8080/e9db20

```