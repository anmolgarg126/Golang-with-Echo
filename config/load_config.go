package config

func LoadConfig() {
	setProperties()
	mongoConfig()
	configureSentry()

}
