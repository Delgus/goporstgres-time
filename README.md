### Типы даты/времени POSTGRESQL

Свежее этой документации не нашел (
[https://postgrespro.ru/docs/postgresql/9.4/datatype-datetime](https://postgrespro.ru/docs/postgresql/9.4/datatype-datetime)

| Имя                | Размер | Описание                          | Наименьшее значение | Наибольшее значение |
|--------------------|:------:|-----------------------------------|---------------------|---------------------|
| timestamp          |   8    | дата и время (без часового пояса) | 4713 до н. э.       | 294276 н. э.        |
| timestamptz        |   8    | дата и время (с часовым поясом)   | 4713 до н. э.       | 294276 н. э.        |
| date               |   4    | дата (без времени суток)          | 4713 до н. э.       | 5874897 н. э.       |
| time               |   8    | время суток (без даты)            | 00:00:00            | 24:00:00            |
| time with timezone |   12   | только время суток (с поясом)     | 00:00:00+1459       | 24:00:00-1459       |
| interval           |   16   | временной интервал                | -178000000 лет      | 178000000 лет       |

Поля без часового пояса обрезают сдвиг, это очень осложняет работу. Лучше использовать timestamptz вместо timestamp. Хотя не факт. 

[timestamp vs timestamptz](https://medium.com/building-the-system/how-to-store-dates-and-times-in-postgresql-269bda8d6403)

[статья на хабре о жести в таймзонах](https://habr.com/ru/company/mailru/blog/242645)

[истерика на ютуб)))](https://www.youtube.com/watch?v=-5wpm-gesOY)

### Пример cmd/ex1

Запускаем контейнера с Postgres в разных часовых поясах  

```bash
docker-compose up -d 
```

На `localhost:5000` будет слушать Postgres с таймзоной по умолчанию - UTC.
На `localhost:5001` будет слушать Postgres с таймзоной Europe/Moscow

Можете соединиться с ними через любой удобный для вас менеджер баз данных, например DBeaver.

Чтобы посмотреть какие запросы отправил наш скрипт, можно получить список контейнеров:
```bash
docker ps
``` 

Увидим что то вроде того:
```
CONTAINER ID        IMAGE                COMMAND                  CREATED             STATUS              PORTS                    NAMES
90e12a54c7ae        postgres:10-alpine   "docker-entrypoint.s…"   14 hours ago        Up 2 hours          0.0.0.0:5001->5432/tcp   gopostgrestime_db-moscow_1
1dfb45acdaf0        postgres:10-alpine   "docker-entrypoint.s…"   14 hours ago        Up 2 hours          0.0.0.0:5000->5432/tcp   gopostgrestime_db-utc_1
```

А затем из полученного списка выбрать id интересующего нас контейнера и вывести его логи в отдельном терминале
```bash
 docker logs -f 90e12a54c7ae
```

Сбилдим наш пример
```bash
make build
```

Запустим:
```bash
./ex1
```

В логах увидим что то вроде того:
```
2020-02-02 11:46:26.861 MSK [176] LOG:  execute <unnamed>: insert into records values (nextval('records_id_seq'), $1, $2, $3, $4, $5, $6, $7)
2020-02-02 11:46:26.861 MSK [176] DETAIL:  parameters: $1 = 'MoscowRecord_ex1', $2 = '2020-01-01 12:12:12', $3 = '2020-01-01 12:12:12+03', $4 = '12:12:12', $5 = '12:12:12+03', $6 = '2020-01-01', $7 = '05:00:00'
2020-02-02 11:46:26.893 MSK [176] LOG:  execute <unnamed>: insert into records values (nextval('records_id_seq'), $1, $2, $3, $4, $5, $6, $7)
2020-02-02 11:46:26.893 MSK [176] DETAIL:  parameters: $1 = 'UTCRecord_ex1', $2 = '2020-01-01 09:12:12', $3 = '2020-01-01 12:12:12+03', $4 = '09:12:12', $5 = '09:12:12+00', $6 = '2020-01-01', $7 = '05:00:00'
2020-02-02 11:46:26.910 MSK [176] LOG:  execute <unnamed>: insert into records values (nextval('records_id_seq'), $1, $2, $3, $4, $5, $6, $7)
2020-02-02 11:46:26.910 MSK [176] DETAIL:  parameters: $1 = 'UnicornRecord_ex1', $2 = '2020-01-01 10:12:12', $3 = '2020-01-01 12:12:12+03', $4 = '10:12:12', $5 = '10:12:12+01', $6 = '2020-01-01', $7 = '05:00:00'

```

### lib/pq

Не поддерживает `time.Duration`. issue [здесь](https://github.com/lib/pq/issues/78)

Как преобразовывает даты библиотека `lib/pq` для запроса в Postgres.

| Тип поля           | Значение в Go                     | Postgres (UTC)         | Postgres (Europe/Moscow) |
|--------------------|:----------------------------------|:-----------------------|--------------------------|
| timestamp          | 2020-01-01 12:12:12,Europe/Moscow | 2020-01-01 12:12:12    | 2020-01-01 12:12:12      |
| timestamp          | 2020-01-01 09:12:12,UTC           | 2020-01-01 09:12:12    | 2020-01-01 09:12:12      |
| timestamptz        | 2020-01-01 12:12:12,Europe/Moscow | 2020-01-01 09:12:12+00 | 2020-01-01 12:12:12+03   |
| timestamptz        | 2020-01-01 09:12:12,UTC           | 2020-01-01 09:12:12+00 | 2020-01-01 12:12:12+03   |
| date               | 2020-01-01 12:12:12,Europe/Moscow | 2020-01-01             | 2020-01-01               |
| date               | 2020-01-01 09:12:12,UTC           | 2020-01-01             | 2020-01-01               |
| time               | 2020-01-01 12:12:12,Europe/Moscow | 12:12:12               | 12:12:12                 |
| time               | 2020-01-01 09:12:12,UTC           | 09:12:12               | 09:12:12                 |
| time with timezone | 2020-01-01 12:12:12,Europe/Moscow | 12:12:12+03            | 12:12:12+03              |
| time with timezone | 2020-01-01 09:12:12,UTC           | 09:12:12+00            | 09:12:12+00              |