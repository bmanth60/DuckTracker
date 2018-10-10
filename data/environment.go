package data

import "os"

var (
	gEnvironment *Environment
)

//Environment variables for database
type Environment struct {
	DatabaseDSN string
}

func setConfig(env *Environment) *Environment {
	gEnvironment = env
	return gEnvironment
}

func readConfig() *Environment {
	return &Environment{
		DatabaseDSN: os.Getenv("DB_DSN"),
	}
}

func getConfig() *Environment {
	if gEnvironment == nil {
		gEnvironment = readConfig()
	}

	return gEnvironment
}
