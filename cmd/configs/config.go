package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"sync"
)

var (
	sconf      = SroundConfig{}
	sconfViper *viper.Viper
	once       sync.Once
)

type MySQLConfig struct {
	User     string
	Password string
	Endpoint string
	DBname   string
}

type SroundConfig struct {
	DatabaseType string
	MySQLConfig  MySQLConfig
	LoginType    []string
}

func loadConfig() {
	initViper()

	//database
	sconf.DatabaseType = getString("database.type")
	if sconf.DatabaseType == SG_DB_MYSQL {
		sconf.MySQLConfig.User = getString("database.mysql.user")
		sconf.MySQLConfig.Password = getString("database.mysql.password")
		sconf.MySQLConfig.Endpoint = getString("database.mysql.endpoint")
		sconf.MySQLConfig.DBname = getString("database.mysql.dbname")
	}

	//auth
	sconf.LoginType = getStringSlice("auth.logintype")

}

func initViper() {
	sconfViper = viper.New()
	sconfViper.SetConfigType("yaml")
	sconfViper.SetConfigName(ConfigFileName)
	sconfViper.AddConfigPath(getProjectPath())
}

func GetSroundConfig() SroundConfig {
	once.Do(func() {
		loadConfig()
	})
	return sconf
}

func getProjectPath() string {
	return os.Getenv("GOPATH") + "/src/github.com/EMO_Server"
}

func getString(key string) string {
	if errCheck(key) != nil {
		return ""
	}
	return sconfViper.GetString(key)
}

func getStringSlice(key string) []string {
	if errCheck(key) != nil {
		return nil
	}
	return sconfViper.GetStringSlice(key)
}

func errCheck(key string) error {
	if sconfViper.IsSet(key) != true {
		err := fmt.Errorf("key is not setted: %s", key)
		log.Println(err.Error())
		return err
	}
	return nil
}
