/**
 * @Author: Cc
 * @Description: 描述
 * @File: configs
 * @Version: 1.0.0
 * @Date: 2022/10/12 17:10
 * @Software : GoLand
 */

package confis

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Configs struct {
	config *Config
	Server [][]string `toml:"server"`
}

var Conf Configs

func InitConfigStart() {

	appPath, err := GetAppAbsPath()
	if err != nil {
		fmt.Println("Get path config error ", err.Error())
		panic(err)
	}

	config, err := NewConfig(path.Join(path.Join(appPath, "conf"), "server.toml"), "toml")
	if err != nil {
		fmt.Println("Get path config error ", err.Error())
		panic(err)
	}
	Conf = Configs{config: config}
}

type Config struct {
	*viper.Viper             //viper结构
	changeHandlerFunc func() //监听到文件修改处理方法
	fileName          string
}

func GetAppAbsPath() (path string, err error) {
	path, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Print("failed get app abs path,err:", err)
		return
	}
	return
}

func NewConfig(fileName string, configType string) (*Config, error) {
	confIn, err := load(fileName, configType)
	if err != nil {
		return nil, err
	}

	c := &Config{
		Viper:    confIn,
		fileName: fileName,
	}

	c.changeMonitor()

	return c, nil
}

func (p *Config) changeMonitor() {
	//配置文件修改监听
	p.OnConfigChange(func(event fsnotify.Event) {
		if event.Op != fsnotify.Write {
			return
		}
		err1 := p.ReadInConfig()
		if err1 != nil {
			log.Printf("config:%v was changed,but read err:%v", event.Name, err1)
			return
		}
		log.Println("config change:", event.Name)
		if p.changeHandlerFunc != nil {
			p.changeHandlerFunc()
		}
	})
	p.WatchConfig()
}

func load(fileName string, configType string) (*viper.Viper, error) {
	confIn := viper.New()
	confIn.SetConfigFile(fileName)
	confIn.SetConfigType(configType)
	err := confIn.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return confIn, nil
}
