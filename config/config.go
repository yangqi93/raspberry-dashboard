package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strings"
)

var Conf Config

type Config struct {
	Name  string
	Value *viper.Viper
}

func init() {
	c := Config{
		Name:  "",
		Value: viper.New(),
	}
	if err := c.initConfig(c.Value); err != nil {
		panic(err.Error())
	}
	Conf = c

	// 监控配置文件变化并热加载程序
	//c.watchConfig()

}

func (c *Config) initConfig(cfg *viper.Viper) error {

	cfg.AddConfigPath("conf")
	// 如果没有指定配置文件，则解析默认的配置文件
	cfg.SetConfigName("config")

	cfg.SetConfigType("yaml")     // 设置配置文件格式为YAML
	cfg.AutomaticEnv()            // 读取匹配的环境变量
	cfg.SetEnvPrefix("APISERVER") // 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	cfg.SetEnvKeyReplacer(replacer)
	if err := cfg.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	return nil
}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		//logger.Infof("Config file changed: %s",e.Name)
	})
}
