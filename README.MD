
## Описание

Сообщения на отправку извлекаются из nats (subject конфигурируется)

Формат сообщений:
```json
{
    "phone": "+7(999) 999-99-99",
    "code": 1234
}
```

Логи отправки сохраняются в PostgreSQL

## Конфигурация

1. Создать в [корне](./) конфигурационный файл .env

Пример: [.env.example](.env.example)

2. Создать в папке [structure](structure) конфигурационный файл .env

Пример: [structure/.env.example](structure/.env.example)

## Запуск
```
cd structure
docker-compose up 
```