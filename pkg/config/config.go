package config

import "sync"

var (
	dbconfig *DBConfig
	onece	sync.Once
)

type DBConfig struct {
	DBtype		string
	DBuser		string
	DBname		string
	DBpassword	string
	DBendpoint	string
	DBport		int
}

func GetDBConfig() *DBConfig{
	onece.Do(func() {
		loadConfig()
	})

	return dbconfig
}

func loadConfig(){
	dbconfig.DBtype = DBTYPE
	dbconfig.DBuser = DBUSER
	dbconfig.DBname = DBNAME
	dbconfig.DBpassword = DBPASSWORD
	dbconfig.DBendpoint = DBENDPOINT
	dbconfig.DBport = DBPORT
}