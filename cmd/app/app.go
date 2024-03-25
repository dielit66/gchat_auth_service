package main

import (
	"fmt"

	"github.com/xxlifestyle/auth_service/internal/proto/app"
	"github.com/xxlifestyle/auth_service/pkg/config"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Printf("error while parsing config, err: %v", err)
	}
	app.Run(cfg)
}
