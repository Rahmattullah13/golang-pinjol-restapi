package config

import (
	"fmt"
	"golang-pinjol/model"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic(errEnv)
	}

	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")
	dbUser := os.Getenv("DBUSER")
	dbPassword := os.Getenv("DBPASSWORD")
	dbName := os.Getenv("DBNAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Master_Customer{}, &model.Master_Document_Customer{}, &model.Master_Jobs_customers{}, &model.Master_Loan{}, &model.Master_Payment_History{}, &model.Transactions_Payment_Loan{})
	if err != nil {
		fmt.Println("Error migrating tables: ", err)
		os.Exit(1)
	}

	return db
}

func CloseDB(db *gorm.DB) {
	dbSql, err := db.DB()
	if err != nil {
		panic(err)
	}
	dbSql.Close()
}
