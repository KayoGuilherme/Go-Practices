package database

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type Database struct {
    db *gorm.DB
}

func NewDatabase() (*Database, error) {
    dsn := "host=localhost user=admin password=secret dbname=mydb port=5432 sslmode=disable"

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return &Database{db: db}, nil
}

func (database *Database) Close() {
    sqlDB, err := database.db.DB()
    if err != nil {
        return
    }
    sqlDB.Close()
}

func (database *Database) GetDB() *gorm.DB {
    return database.db
}