package models

import (
    "fmt"
    "log"
    "BuffettFive/config"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var Db *gorm.DB
var err error

func init() {
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        config.Config.DBHost,
        config.Config.DBPort,
        config.Config.DBUser,
        config.Config.DBPassword,
        config.Config.DBName,
    )

    Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }

    Db.AutoMigrate(&Todo{})
}
