CREATE TABLE records (
    id serial not null,
    title varchar(255) not null,
    timestamp_t timestamp not null,
    timestamptz_t timestamptz not null,
    time_t time not null,
    timetz_t time with time zone not null,
    date_t date not null,
    interval_t interval not null
);