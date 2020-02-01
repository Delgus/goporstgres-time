package main

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Record - simple row
type Record struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Timestamp   time.Time `db:"timestamp_t"`
	TimestampTZ time.Time `db:"timestamptz_t"`
	Time        time.Time `db:"time_t"`
	TimeTZ      time.Time `db:"timetz_t"`
	Date        time.Time `db:"date_t"`
	Interval    string    `db:"interval_t"`
}

func main() {
	// open connection with postgres
	dbUTC, err := sqlx.Open("postgres", "user=postgres password=123456 dbname=postgres host=localhost port=5000 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer dbUTC.Close()

	dbMoscow, err := sqlx.Open("postgres", "user=postgres password=123456 dbname=postgres host=localhost port=5001 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer dbMoscow.Close()

	// Moscow office
	loc, err := time.LoadLocation(`Europe/Moscow`)
	if err != nil {
		log.Fatal(err)
	}
	moscowTime := time.Date(2020, 1, 1, 12, 12, 12, 0, loc)
	moscowRecord := newRecord(`MoscowRecord`, moscowTime)

	// UTC office
	utcTime := time.Date(2020, 1, 1, 9, 12, 12, 0, time.UTC)
	utcRecord := newRecord(`UTCRecord`, utcTime)

	// Unicorn office
	unicornTime := time.Date(2020, 1, 1, 10, 12, 12, 0, time.FixedZone(`Unicorn`, 3600))
	unicornRecord := newRecord(`UnicornRecord`, unicornTime)

	// Save records in databases
	if err := insertRecord(dbMoscow, moscowRecord); err != nil {
		log.Fatal(err)
	}
	if err := insertRecord(dbUTC, moscowRecord); err != nil {
		log.Fatal(err)
	}

	if err := insertRecord(dbMoscow, utcRecord); err != nil {
		log.Fatal(err)
	}
	if err := insertRecord(dbUTC, utcRecord); err != nil {
		log.Fatal(err)
	}
	if err := insertRecord(dbMoscow, unicornRecord); err != nil {
		log.Fatal(err)
	}
	if err := insertRecord(dbUTC, unicornRecord); err != nil {
		log.Fatal(err)
	}
}

func newRecord(title string, t time.Time) *Record {
	return &Record{
		Title:       title,
		Timestamp:   t,
		TimestampTZ: t,
		Time:        t,
		TimeTZ:      t,
		Date:        t,
		Interval:    `5h`,
	}
}

func insertRecord(db *sqlx.DB, r *Record) error {
	query := "insert into records values (nextval('records_id_seq'), $1, $2, $3, $4, $5, $6, $7)"
	_, err := db.Exec(query, r.Title, r.Timestamp, r.TimestampTZ, r.Time, r.TimeTZ, r.Date, r.Interval)
	return err
}
