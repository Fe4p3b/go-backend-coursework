# URL shortener

Сервис для сокращения ссылок.

# API

Создание сокращенной ссылки

``` / ```

POST

Request body:
```
{
    "original_url": "url_to_shorten"
}
```
Response body:
```
{
    "short_url": "shortened_url"
}
```

Переход 

``` /:url ```

url - идентификатор ссылки

Cтатистика

``` /:url/stats ```

url - идентификатор ссылки