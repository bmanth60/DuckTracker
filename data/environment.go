package data

import "os"

var (
	//gEnvironment global environment config for database
	gEnvironment *Environment
)

//Environment variables for database
type Environment struct {
	DatabaseDSN string
}

//setConfig set configuration environment
func setConfig(env *Environment) *Environment {
	gEnvironment = env
	return gEnvironment
}

//readConfig from environment variables
func readConfig() *Environment {
	return &Environment{
		DatabaseDSN: os.Getenv("DB_DSN"),
	}
}

//getConfig get configuration singleton
func getConfig() *Environment {
	if gEnvironment == nil {
		gEnvironment = readConfig()
	}

	return gEnvironment
}
