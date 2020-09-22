package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"wiliwili/util"
)

type DataStruct struct {
	MySql struct {
		Host     string
		Port     string
		User     string
		Password string
		Dbname   string
	}
	LogLever string
	Redis    struct {
		Addr     string
		Password string
		Dbname   string
	}
	SessionSecret string `yaml:"session_secret"`
	GinMode       string `yaml:"gin_mode"`
	UpdatePath    string `yaml:"update_path"`
}

var Data = new(DataStruct)

// Init 初始化配置项
func Init() {
	// 获取配置文件
	if initConf() != nil {
		util.Log().Panic("init error")
		return
	}

	// 设置日志级别
	util.BuildLogger(Data.LogLever)
	//err := errors.New("test log")
	//util.LogD("失败 logD %+v", err)
	util.LogD("Data", Data)
	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	//// 连接数据库
	//model.Database(os.Getenv("MYSQL_DSN"))
	//cache.Redis()
	//
	//// 启动定时任务
	//tasks.CronJob()
}

func initConf() error {
	data, err := ioutil.ReadFile("conf/conf.yaml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, Data)
	if err != nil {
		return err
	}
	return nil
}
