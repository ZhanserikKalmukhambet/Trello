package main

import (
	"fmt"
	"github.com/ZhanserikKalmukhambet/Trello/internal/config"
	"github.com/ZhanserikKalmukhambet/Trello/pkg/http_server"
)

func main() {
	// set configuration files
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%#v", cfg))

	// initialize server
	server := http_server.New()
}
