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

### lib/pq

Не поддерживает `time.Duration`. issue [здесь](https://github.com/lib/pq/issues/78)

