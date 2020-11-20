package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type database struct {
	User string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	Dbname string `mapstructure:"dbname"`
}

var DatabaseConf = database{}

func Setup() {
	//当时写这个配置文件的时候差点没被气死，起初使用的是yaml类型的配置文件。代码和教程代码一致，也没报错，就是读不出来数据，我哭了啊。
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("配置信息改变:", e.Name)
	})
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("数据读取失败%v", err)
	}
	err = viper.UnmarshalKey("mysql", &DatabaseConf)
	if err != nil {
		log.Printf("配置数据绑定失败：%v", err)
	}
}
