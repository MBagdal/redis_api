package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erro ao carregar arquiv .env")
	}

	fmt.Println(os.Getenv("TEST"))
}
