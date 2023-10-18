package config

import (
	"github.com/spf13/viper"
)

var (
	GoogleCredentialFileName string
	ProductionEnv            string = "production"
	ApplicationEnv           string
	DefaultColor             string = "#FF0000"
	DefaultBorderSize        int    = 2
)

func InitializeConfig() {

	// Set the file name of the configurations file
	viper.SetConfigFile(`.env`)

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.IsSet(`GOOGLE_CREDENTIAL_FILE_NAME`) {
		GoogleCredentialFileName = viper.GetString(`GOOGLE_CREDENTIAL_FILE_NAME`)
	}

	if viper.IsSet(`APPLICATION_ENV`) {
		ApplicationEnv = viper.GetString(`APPLICATION_ENV`)
	}
}
