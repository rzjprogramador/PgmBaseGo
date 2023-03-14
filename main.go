package main

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/rzjprogramador/PgmBaseGo/utils"
	// "github.com/rzjprogramador/PgmBaseGo/tecnica"
)

func main() {
	test := utils.UseLoadEnv("TEST")
	nome := utils.UseLoadEnv("NOME")
	idade, _ := strconv.Atoi(utils.UseLoadEnv("IDADE"))
	tipoIdade := reflect.TypeOf(idade)
	
	fmt.Println(test, nome, idade)
	fmt.Println(tipoIdade)
}
