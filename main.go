package main

import (
	"database/sql"
	"log"
	"os"

	"example.com/blog/v2/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port in not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL in not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("can't connect to the db", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	r := gin.Default()
	r.POST("/user", apiCfg.handlerCreateUser)
	r.POST("/login", apiCfg.handlerLogin)
	r.GET("/validate", apiCfg.UserAuth, apiCfg.handlerValidate)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")}

}
