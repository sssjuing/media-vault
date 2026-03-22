package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func init() {
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("config")
	// config.AddConfigPath("$MEDIA_VAULT_CONFIG_PATH")
	config.AddConfigPath("/etc/media-vault")

	if err := config.ReadInConfig(); err != nil {
		configsDirPath := GetRootPath() + "/configs"
		config.AddConfigPath(configsDirPath)
		if err := config.ReadInConfig(); err != nil {
			log.Fatal("Error on parsing default configuration file. ", err)
		}
	}
}

func GetConfig() *viper.Viper {
	return config
}

func GetMinioPublicUrl() string {
	endpoint := config.GetString("minio.public_endpoint")
	bucketName := config.GetString("minio.bucket_name")
	return fmt.Sprintf("//%s/%s/", endpoint, bucketName)
}

type User struct {
	Username string `mapstructure:"name"`
	Password string `mapstructure:"password"`
}

func GetUsers() ([]User, error) {
	var users []User
	if err := config.UnmarshalKey("server.users", &users); err != nil {
		return nil, fmt.Errorf("error unmarshaling server.users: %w", err)
	}
	return users, nil
}
