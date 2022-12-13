package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewPostgresDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", "postgres://sebas:050816@localhost:5432/dbpos?sslmode=disable")
		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Can't do ping db: %v", err)
		}

		fmt.Println("Connected to postgres")
	})
}

//Pool return a unique instace of db
func Pool() *sql.DB {
	return db
}

func stringtoNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

func timeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}
	if !null.Time.IsZero() {
		null.Valid = true
	}
	return null
}
