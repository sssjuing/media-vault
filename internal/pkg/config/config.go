package config

import (
	"fmt"
	"log"

	"github.com/samber/lo"
	"github.com/spf13/viper"
	"github.com/sssjuing/media-vault/internal/pkg/utils"
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
		configsDirPath := utils.GetRootPath() + "/configs"
		config.AddConfigPath(configsDirPath)
		if err := config.ReadInConfig(); err != nil {
			log.Fatal("Error on parsing default configuration file. ", err)
		}
	}
}

func GetConfig() *viper.Viper {
	return config
}

func makeDsn(host string) string {
	username := config.GetString("postgres.username")
	password := config.GetString("postgres.password")
	dbname := config.GetString("postgres.dbname")
	port := config.GetString("postgres.port")
	// host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, username, password, dbname, port)
}

func GetPostgresDsn() string {
	host := config.GetString("postgres.host")
	return makeDsn(host)
}

func GetPostgresReplicas() []string {
	hosts := config.GetStringSlice("postgres.replicas")
	return lo.Map(hosts, func(host string, index int) string {
		return makeDsn(host)
	})
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
