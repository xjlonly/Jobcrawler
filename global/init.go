package global

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var once = new(sync.Once)
func Init()  {
	once.Do(func() {
		//设置读取的配置文件
		viper.SetConfigName("config")
		//添加读取的配置文件路径
		viper.AddConfigPath("/etc/crawler/")
		//windows环境下为%GOPATH，linux环境下为$GOPATH
		viper.AddConfigPath("$GOPATH/src/")
		viper.AddConfigPath(App.RootDir+"/config")
		//设置配置文件类型
		viper.SetConfigType("yaml")
		err := viper.ReadInConfig()
		if err != nil{
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	})
}
