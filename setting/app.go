package setting

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

//定义conf类型
//类型里的属性，全是配置文件里的属性
type Conf struct {
	Service     Service     `yaml:"service"`
	MySQLConfig MySQLConfig `yaml:"mysql"`
}

type MySQLConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

type Service struct {
	Url    string  `yaml:"url"`
	Host   string  `yaml:"host"`
	Topics []Topic `yaml:"subscriptions"`
}

type Topic struct {
	Topic   string `yaml:"topic"`
	Address string `yaml:"address"`
}

//读取Yaml配置文件,
//并转换成conf对象
func (conf *Conf) GetConf() *Conf {
	//应该是 绝对地址
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	return conf
}
