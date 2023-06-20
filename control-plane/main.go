package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"senkawa.moe/haa-chan/app/config"
	"senkawa.moe/haa-chan/cmd"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("failed to load .env file: %v\n", err)
		os.Setenv("APP_ENV", "production")
		fmt.Println("defaulting to production environment")
	}

	if err := config.LoadConfiguration(config.ParseEnvironment()); err != nil {
		bailWith(fmt.Sprintf("failed to load yaml configuration: %v", err))
	}

	cmd.Execute()
}

func bailWith(message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}
