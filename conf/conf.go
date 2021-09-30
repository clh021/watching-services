package conf

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/linakesi/lnksutils"
	"github.com/spf13/viper"
)

type Config struct {
	// 配置文件内容的路径解析依据配置文件自身的路径
	ConfigDir     string `yaml:"-"`
	ServicePath   string `yaml:"servicePath"` // 只能驼峰，不能-和_连接
	NginxConfPath string `yaml:"nginxConfPath"`
}

func (c *Config) transAbsPath(p *string) bool {
	if !lnksutils.IsDirExist(*p) {
		if !lnksutils.IsDirExist(filepath.Join(c.ConfigDir, *p)) {
			return false
		} else {
			*p = filepath.Join(c.ConfigDir, *p)
		}
	}
	return true
}

// 转化允许配置文件所在目录的相对路径
func (c *Config) transformIfNeed() error {
	if !c.transAbsPath(&c.ServicePath) {
		return errors.New("conf: 'service_path:" + c.ServicePath + "' not exist. ")
	}
	if !c.transAbsPath(&c.NginxConfPath) {
		return errors.New("error: 'nginx_conf_path:" + c.NginxConfPath + "' not exist. ")
	}
	return nil
}

func GetProgramPath() string {
	ex, err := os.Executable()
	if err == nil {
		return filepath.Dir(ex)
	}

	exReal, err := filepath.EvalSymlinks(ex)
	if err != nil {
		panic(err)
	}
	return filepath.Dir(exReal)
}

func GetConf() *Config {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(GetProgramPath())
	viper.AddConfigPath(filepath.Join(GetProgramPath(), "tests"))
	viper.AddConfigPath(filepath.Join(path.Dir(GetProgramPath()), "tests"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	c := Config{}
	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}
	log.Println(color.BlueString("CONF:"), viper.ConfigFileUsed())
	c.ConfigDir = filepath.Dir(viper.ConfigFileUsed())
	if err = c.transformIfNeed(); err != nil {
		panic(err)
	}
	return &c
}
