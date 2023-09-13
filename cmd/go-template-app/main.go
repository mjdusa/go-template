package main

import (
	"os"

	"github.com/go-template-org/go-template-app/internal/runner"
)

func main() {
	os.Exit(runner.Run())
}
