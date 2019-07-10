package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
	"jobcrawler/global"
	"jobcrawler/model"
	"jobcrawler/util"
)

var dbEngine *xorm.Engine
func init()  {
	global.Init()
	config := model.DBConfig{}
	config.DbUser = viper.GetString("storage.user");
	config.DbPass = viper.GetString("storage.password");
	config.DbHost = viper.GetString("storage.host");
	config.DbPort = viper.GetString("storage.port");
	config.DbName =	viper.GetString("storage.dbname");
	config.DbCharset  = viper.GetString("storage.charset");
	config.DbDriver = viper.GetString("storage.driver")
	initConn(config)
}

func initConn(config model.DBConfig)  {
	dbConnString := getDbConnString(config)
	var err error
	dbEngine,err = xorm.NewEngine(config.DbDriver, dbConnString)
	if err != nil{
		panic(err)
	}

	if errConn := dbEngine.Ping(); errConn != nil{
		panic(errConn)
	}

}

func getDbConnString(config model.DBConfig) string {
	builder := util.NewStringBuilder()
	builder.Append(config.DbUser).Append(":").Append(config.DbPass)
	builder.Append("@tcp(").Append(config.DbHost).Append(":").Append(config.DbPort).Append(")/")
	builder.Append(config.DbName).Append("?charset=").Append(config.DbCharset)
	if config.IsLocalTime {
		builder.Append("&pauseTime=true&loc=Local")
	}
	return  builder.ToString()
}

