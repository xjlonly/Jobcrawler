package global

import (
	"github.com/spf13/viper"
	"jobcrawler/util"
	"os"
	"path/filepath"
	"sync"
	"time"
)


type app struct {
	Name string
	Version string
	Date time.Time

	//项目启动根目录
	RootDir string

	//启动时间
	LaunchTime time.Time
	UpTime time.Duration

	//锁
	locker sync.Mutex
}

var App = &app{}

func  init()  {
	App.Name = os.Args[0]
	App.Version = "1.0"
	App.LaunchTime = time.Now()

	App.RootDir = "."

	if !viper.InConfig("http.config"){
		App.RootDir = inferRootDir()
	}

	fileInfo, err := os.Stat(os.Args[0])
	if err != nil{
		panic(err)
	}

	App.Date = fileInfo.ModTime()
}

func  inferRootDir() string {
	pwd, err := os.Getwd()
	if err != nil{
		panic(err)
	}

	var infer func(dir string) string
	infer = func(dir string) string {
		if util.IsExist(dir + "/config"){
			return  dir;
		}
		if dir == "/" {
			return ""
		}
		return  infer(filepath.Dir(dir))
	}

	return infer(pwd)
}


func (a *app) SetUpdateTime()  {
	a.locker.Lock()
	defer  a.locker.Unlock()
	a.UpTime = time.Now().Sub(a.LaunchTime)
}