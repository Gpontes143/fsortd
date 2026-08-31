package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Gpontes143/fsortd/internal/config"
)

func main() {
	file, err := os.ReadFile("./config.example.toml")
	if err != nil {
		log.Fatal(err)
	}
	configuracao, err := config.TomlUnmarshal(file)
	if err != nil {
		log.Fatal(err)
	}
	rotas := config.DestinyTable(configuracao.Rules)
	fmt.Println(rotas)
}
