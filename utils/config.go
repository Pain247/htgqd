package utils

import (
	"os"
	"log"
	"github.com/BurntSushi/toml"
)

type ConfigSql struct {
	MYSQL_DB_HOST                  string
	MYSQL_DB_NAME                  string
	MYSQL_DB_USER                  string
	MYSQL_DB_PASS                  string
}
type ConfigServer struct{
	ADDR                           string
}


//function load all config of ssp server follow ConfigServerSSP struct
func LoadConfigSql(nameFileConfig string) ConfigSql{
	var config ConfigSql
	_, err := os.Stat(nameFileConfig)
	if err != nil {
		log.Fatal("Config file is missing: ", nameFileConfig)
	}

	if _, err := toml.DecodeFile(nameFileConfig, &config); err != nil {
		log.Fatal("Wrong parameters!")
		log.Fatal(err)
	}
	return config
}
func LoadConfigServer(nameFileConfig string) ConfigServer{
	var config ConfigServer
	_, err := os.Stat(nameFileConfig)
	if err != nil {
		log.Fatal("Config file is missing: ", nameFileConfig)
	}

	if _, err := toml.DecodeFile(nameFileConfig, &config); err != nil {
		log.Fatal("Wrong parameters!")
		log.Fatal(err)
	}
	return config
}