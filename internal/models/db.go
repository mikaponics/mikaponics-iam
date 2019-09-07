package models // github.com/mikaponics/mikapod-iam/internal/models

import (
    "fmt"
    "log"

    _ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
)


type MikapodDB struct {
	db *sqlx.DB
}


/**
 *  Function initializes our connection with the `postgres` database for this
 *  web-application and saves the db connection instance in a global variable.
 */
func InitDB(dbHost, dbPort, dbUser, dbPassword, dbName string) (*MikapodDB) {
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+ "password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
    dbInstance, err := sqlx.Connect("postgres", psqlInfo)
    if err != nil {
       log.Fatalln(err)
    }
    err = dbInstance.Ping()
    if err != nil {
        panic(err)
    }
    fmt.Println("Database successfully connected!")
    return &MikapodDB{
        db: dbInstance,
    }
}

func (modelLayer *MikapodDB) Shutdown() {
    modelLayer.db.Close()
}
