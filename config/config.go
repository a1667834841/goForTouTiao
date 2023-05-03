package config

import (
	"fmt"
	"path"
	"runtime"

	"github.com/spf13/viper"
)

// 读取配置文件config
type ConfigInfo struct {
	Cookie string
}

var config ConfigInfo

func init() {
	// 把配置文件读取到结构体上
	dir := getCurrentAbPathByCaller()
	viper.SetConfigName("config")
	viper.AddConfigPath(dir)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	viper.Unmarshal(&config) //将配置文件绑定到config上
	// fmt.Println("config: ", config, "Cookie: ", config.Cookie)

}

func GetConfig() ConfigInfo {
	return config
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
