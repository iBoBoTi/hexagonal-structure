package helper

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"os"
)

type ConfigStruct struct {
	Address            string `mapstructure:"address"`
	Port               string `mapstructure:"port"`
	Mode               string `mapstructure:"mode"`
	DbHost             string `mapstructure:"db_host"`
	DbName             string `mapstructure:"db_name"`
	AppName            string `mapstructure:"app_name"`
	LogDir             string `mapstructure:"log_dir"`
	LogFile            string `mapstructure:"log_file"`
	ExternalConfigPath string `mapstructure:"external_config_path"`
}

var (
	address            string
	port               string
	mode               string
	dbHost             string
	dbName             string
	externalConfigPath string
)

func LoadConfig() (string, string, string, string, string, string) {
	flag.StringVar(&address, "address", Config.Address, "local host")
	flag.StringVar(&port, "port", Config.Port, "application ports")
	flag.StringVar(&mode, "mode", Config.Mode, "application mode, either dev or prod")
	flag.StringVar(&dbHost, "dbhost", Config.DbHost, "database host")
	flag.StringVar(&dbName, "dbname", Config.DbName, "database name")
	flag.StringVar(&externalConfigPath, "external_config_path", Config.DbName, "external config path")
	flag.Parse()
	for i, val := range flag.Args() {
		os.Args[i] = val
	}
	return address, port, mode, dbHost, dbName, externalConfigPath
}
func LoadEnv(path string) (config ConfigStruct, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("resource-service")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return ConfigStruct{}, err
	}
	err = viper.Unmarshal(&config)
	return
}
func ReturnConfig() ConfigStruct {
	config, err := LoadEnv(".")
	if err != nil {
		log.Println(err)
	}
	if config.ExternalConfigPath != "" {
		viper.Reset()
		config, err = LoadEnv(config.ExternalConfigPath)
		if err != nil {
			log.Println(err)
		}
	}
	return config
}

var Config = ReturnConfig()
