package main

import (
	"fmt"
	"github.com/ZhanserikKalmukhambet/Trello/internal/config"
)

func main() {
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%#v", cfg))
}
