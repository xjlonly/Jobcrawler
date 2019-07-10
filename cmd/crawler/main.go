package main

import (
	"fmt"
	"github.com/spf13/viper"
	"jobcrawler/dao"
	"jobcrawler/global"
)

func init()  {
	global.Init()
}

func main() {

	viper.SetDefault("http:port", 9090)
	host := viper.GetString("http.host")
	port :=viper.GetString("http.port")

	address := host + ":" + port
	fmt.Println("Server listen at:" + address)
	//log.Fatal(http.ListenAndServe(address, nil))

	rules,err := dao.DefaultRuleRepository.GetList()
	if(err != nil){
		panic(err)
	}
	fmt.Println(rules)
}
