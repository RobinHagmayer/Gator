package main

import (
	"fmt"

	"github.com/RobinHagmayer/Gator/internal/config"
)

func main() {
	cfg := config.Read()
	cfg.SetUser("robin")
	cfgNew := config.Read()
	fmt.Printf("%+v\n", cfgNew)
}
