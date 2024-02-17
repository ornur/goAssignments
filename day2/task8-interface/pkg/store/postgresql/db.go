package postgresql

import (
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
)

func Connect(host, port, user, password, dbname string) (*sql.DB, error) {
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }

    fmt.Println("Successfully connected to PostgreSQL!")
    return db, nil
}
