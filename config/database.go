package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func Database() *sql.DB{
	user, exist := os.LookupEnv("DB_USER")

	if !exist {
		log.Fatal("DB_USER not set in .env")
	}

	pass, exist := os.LookupEnv("DB_PASS")

	if !exist {
		log.Fatal("DB_PASS not set in .env")
	}

	host, exist := os.LookupEnv("DB_HOST")

	if !exist {
		log.Fatal("DB_HOST not set in .env")
	}

	credentials := fmt.Sprintf("%s:%s@(%s:3306)/?charset=utf8&parseTime=True", user, pass, host)

	database, err := sql.Open("mysql", credentials)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database Connection Successful")
	}

	// _, err = database.Exec(`CREATE DATABASE gotodo`)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	_, err = database.Exec(`USE gotodo`)

	if err != nil {
		fmt.Println(err)
	}

	// _, err = database.Exec(`
	// 	CREATE TABLE todos (
	// 	id INT AUTO_INCREMENT,
	// 	title TEXT NOT NULL,
	// 	completed BOOLEAN DEFAULT FALSE,
	// 	color VARCHAR(10),
	// 	start DATETIME,
	// 	end DATETIME,
	// 	PRIMARY KEY (id)
	// 	);
	// `)

	if err != nil {
		fmt.Println(err)
	}

	return database
}