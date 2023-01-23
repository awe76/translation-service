package models

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/HqOapp/translation-service/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getDBPort(dbname string) int {
	m := config.GetenvMap("SERVICE_MAP")
	key := fmt.Sprintf("%s-db", dbname)
	host, ok := m[key]
	if !ok {
		panic(fmt.Sprintf("Key <%s> is not found in the service map", key))
	}

	p := strings.Split(host, ":")
	if len(p) != 2 {
		panic(fmt.Sprintf("Unexpected value <%s> of the key <%s>, expected format is <host:port>", host, key))
	}

	port, err := strconv.Atoi(p[1])
	if err != nil {
		panic(fmt.Sprintf("Unexpected database port value <%s>, shold be a number", p[1]))
	}
	return port
}

func ConnectDatabase() {
	user := config.Getenv("DB_USER")
	password := config.Getenv("DB_PASSWORD")
	dbname := config.Getenv("DB_NAME")
	port := getDBPort(dbname)

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%d sslmode=disable", user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Translation{})
	db.Exec(createUpdateTranslationsFunc)

	DB = db
}
