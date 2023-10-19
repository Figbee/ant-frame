package initialize

import (
	"ant-frame/bootstrap/config"
	"ant-frame/bootstrap/global"
	"github.com/spf13/viper"
	"log"
	"os"
)

const (
	configType = "yml"
	configPath = "./bootstrap/config"
)

func InitConfig() {
	getenv := os.Getenv("APP_ENV")
	cfgPath := getConfigEnv(getenv)

	v, err := loadConfig(cfgPath, configType)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	parseConfig, err := parseConfig(v)
	global.Conf = *parseConfig
}

/*
 * Load config file
 */
func loadConfig(filename string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filename)
	v.SetConfigType(fileType)
	v.AddConfigPath(configPath)
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("Unable to read config: %v", err)
		return nil, err
	}
	return v, nil
}

/*
 * Get config file name
 */
func getConfigEnv(env string) string {
	if env == "prod" {
		return "config-prod"
	} else if env == "dev" {
		return "config-dev"
	} else {
		return "config-dev"
	}
}

/*
 * Parse config file
 */
func parseConfig(v *viper.Viper) (*config.Config, error) {
	var cfg config.Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to parse config: %v", err)
		return nil, err
	}
	return &cfg, nil
}
