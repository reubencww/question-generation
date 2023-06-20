package config

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var appEnv = Production

type Environment int

const (
	Production Environment = iota
	Development
)

func (e Environment) String() string {
	switch e {
	case Development:
		return "development"
	case Production:
		return "production"
	default:
		// this should never execute
		return "invalid environment"
	}
}

func IsDevelopment() bool {
	return appEnv == Development
}

func IsProduction() bool {
	return appEnv == Production
}

func GetEnvironment() Environment {
	return appEnv
}

func ParseEnvironment() Environment {
	switch configEnv := os.Getenv("APP_ENV"); configEnv {
	case "development", "dev", "local":
		appEnv = Development
	case "production", "prod":
		appEnv = Production
	default:
		fmt.Printf("Unknown environment: %s (Supported values are dev/development, prod/production)\n"+
			"App environment defaulted to: %s\n",
			color.YellowString(configEnv),
			color.GreenString("production"),
		)
		appEnv = Production
	}

	return appEnv
}
