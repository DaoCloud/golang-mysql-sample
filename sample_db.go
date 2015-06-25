package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func MustConnectDB() {
    if err := ConnectDatabase(); err != nil {
        panic(err)
    }
    InitDB()
}

func ConnectDatabase() (err error) {

    username := os.Getenv("MYSQL_USERNAME")
    password := os.Getenv("MYSQL_PASSWORD")
    if username == "" {
        username = "root"
    }

    host := os.Getenv("MYSQL_PORT_3306_TCP_ADDR")
    if host == "" {
        host = "localhost"
    }

    port := os.Getenv("MYSQL_PORT_3306_TCP_PORT")
    if port == "" {
        port = "3306"
    }

    database := "test"

    if len(os.Getenv("MYSQL_INSTANCE_NAME")) > 0 {
        database = os.Getenv("MYSQL_INSTANCE_NAME")
    }

    uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

    if db, err = sql.Open("mysql", uri); err != nil {
        return
    }

    err = db.Ping()
    return
}

func InitDB() {
    defer func() {
        if e := recover(); e != nil {
            log.Println(e)
        }
    }()

    CreateTable()
    Insert(&Person{Name: "Ale", Phone: "+55 53 1234 4321"})
    Insert(&Person{Name: "Cla", Phone: "+66 33 1234 5678"})
}

func CreateTable() {
    stmt, err := db.Prepare(createTable)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    _, err = stmt.Exec()
    if err != nil {
        panic(err)
    }
}

var createTable = `
CREATE TABLE IF NOT EXISTS people (
     user_id      INTEGER PRIMARY KEY AUTO_INCREMENT
    ,username     VARCHAR(32)
    ,phone        VARCHAR(32)
);
`

func Drop() {
    stmt, err := db.Prepare(dropTable)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    _, err = stmt.Exec()
    if err != nil {
        panic(err)
    }

    CreateTable()
}

var dropTable = `
DROP TABLE people;
`
