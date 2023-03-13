package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	// "github.com/rzjprogramador/PgmBaseGo/tecnica"
)



func main() {
	// tecnica.ExecuteTecnica()

	// Acessando variavelDeAmbiente definida no arquivo .env

	// 1- carregar variavel - verificando se tem erro
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
	// 2- usar variavelCarregada - passando a variavel como string
	fmt.Println(os.Getenv("TEST"))
}