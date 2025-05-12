Person Data API

##  Описание

Сервис для обработки и хранения данных о людях. Принимает ФИО через REST API, обогащает данные с помощью внешних открытых API (возраст, пол, национальность), сохраняет в PostgreSQL и предоставляет возможность управлять этими данными.

## Функциональность

- Добавление пользователя по ФИО
- Обогащение:
  - Возраст через [`agify.io`](https://api.agify.io)
  - Пол через [`genderize.io`](https://api.genderize.io)
  - Национальность через [`nationalize.io`](https://api.nationalize.io)
- Хранение обогащённых данных в PostgreSQL
- Получение списка людей с пагинацией
- Обновление информации по ID
- Удаление человека по ID
- Swagger-документация (`localhost:8080/swagger/index.html`)
- Логирование через Logrus
- Конфигурация через `.env`

## Технологии

- Golang
- Gorilla Mux
- PostgreSQL
- Logrus
- Swag (Swagger генератор)

## Структура JSON-запроса

```json
{
  "name": "Dmitriy",
  "surname": "Ushakov",
  "patronymic": "Vasilevich" // необязательно
}
