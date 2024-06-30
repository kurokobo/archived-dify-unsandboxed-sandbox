package main

import (
	"fmt"
	"os"

	"github.com/kurokobo/dify-unsandboxed-sandbox/internal/router"
)

func main() {
	r := router.Default()
	port, exists := os.LookupEnv("SANDBOX_PORT")
	if !exists {
		port = "8194"
	}
	r.Run(fmt.Sprintf("0.0.0.0:%s", port))
}
