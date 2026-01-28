package inits

import (
	"fmt"
	"week3/wei/user-server/basic/config"

	"github.com/spf13/viper"
)

func ConfigInit() {
	viper.SetConfigFile("C:\\Users\\ZhuanZ\\Desktop\\week3\\week3\\wei\\user-server\\dev.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config.GlobalConf)
	if err != nil {
		return
	}
	fmt.Printf("配置文件读取成功")
}
