package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/metgag/mt-w10d1/internal/configs"
	"github.com/metgag/mt-w10d1/internal/routers"
)

// var users = []User{}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	dbpool, err := configs.InitDB()
	if err != nil {
		log.Printf("unable to create connection pool: %s\n", err)
	}
	defer dbpool.Close()

	if err := configs.PingDB(dbpool); err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		return
	} else {
		log.Printf("connected to database %s as user %s\n", os.Getenv("DB_NAME"), os.Getenv("DB_USER"))
	}

	r := routers.InitRouter(dbpool)

	r.Run()
}
