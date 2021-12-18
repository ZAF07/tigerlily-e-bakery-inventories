package env

import (
	"fmt"
	"os"
)

func SetEnv() {
	serverENV := os.Args[1]
	dbHost := os.Args[2]
	dbUser := os.Args[3]
	dbPassword := os.Args[4]
	dbName := os.Args[5]
	dbSSL := os.Args[6]
	dbPort := os.Args[7]
	
	// port := strconv.Itoa(dbPort)
	os.Setenv("serverenv", serverENV)
	os.Setenv("dbHost", dbHost)
	os.Setenv("dbUser", dbUser)
	os.Setenv("dbPassword", dbPassword)
	os.Setenv("dbName", dbName)
	os.Setenv("dbPort", dbPort)
	os.Setenv("dbSSL", dbSSL)
	fmt.Println("DONE SETTING ENV")
}

func GetDBEnv() (dbString string) {

	env := os.Getenv("serverenv")
	host := os.Getenv("dbHost")
	user := os.Getenv("dbUser")
	password := os.Getenv("dbPassword")
	dbname := os.Getenv("dbName")
	port := os.Getenv("dbPort")
	sslMode := os.Getenv("dbSSL")

	switch env {
	case "PROD":
		fmt.Println("DEVELOPMENT")

		dbString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslMode)
		return
	default:
		dbString = fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s", host, user, dbname, port, sslMode)
	}

	return
}