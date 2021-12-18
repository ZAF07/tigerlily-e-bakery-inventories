package env

import (
	"os"
)

func SetEnv() {
	
	dbHost := os.Args[1]
	dbUser := os.Args[2]
	dbPassword := os.Args[3]
	dbName := os.Args[4]
	dbSSL := os.Args[5]
	dbPort := os.Args[6]
	
	// port := strconv.Itoa(dbPort)
	os.Setenv("host", dbHost)
	os.Setenv("dbUser", dbUser)
	os.Setenv("dbPassword", dbPassword)
	os.Setenv("dbName", dbName)
	os.Setenv("dbPort", dbPort)
	os.Setenv("dbSSL", dbSSL)
}