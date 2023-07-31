package main

import (
	"fmt"
	"github.com/ZhanserikKalmukhambet/Trello/internal/app"
	"github.com/ZhanserikKalmukhambet/Trello/internal/config"
)

func main() {
	// set configuration files
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%#v", cfg))

	err = app.Run(cfg)
	if err != nil {
		panic(err)
	}
}
