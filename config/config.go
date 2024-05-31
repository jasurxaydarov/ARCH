package config

type Config struct{
	DbUser		string
	DbPassword	string
	DbHost		string
	DbPort		int
	DbName		string
}


func NEwConfig()Config{

	cfg:=Config{}

	cfg.DbUser="jasur"
	cfg.DbPassword="1001"
	cfg.DbHost="localhost"
	cfg.DbPort=5432
	cfg.DbName="postgres"
	return cfg
}
