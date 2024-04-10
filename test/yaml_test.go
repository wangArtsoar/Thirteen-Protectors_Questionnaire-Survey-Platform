package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"testing"
)

type MySQLConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"dbname"`
}
type RedisConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Auth string `yaml:"auth"`
}
type NginxProxyConfig struct {
	Counter   int      `yaml:"counter"`
	NginxList []string `yaml:"nginx_list"`
}

type ServiceYaml struct {
	MySQL      MySQLConfig      `yaml:"mysql"`
	Redis      RedisConfig      `yaml:"redis"`
	NginxProxy NginxProxyConfig `yaml:"nginx_proxy"`
}

func TestYaml(t *testing.T) {
	filename := "/web/config.yaml"
	y := new(ServiceYaml)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("read file err %v\n", err)
		return
	}
	err = yaml.Unmarshal(yamlFile, y)
	if err != nil {
		log.Fatalf("yaml unmarshal: %v\n", err)
		return
	}

	fmt.Printf("MySQL host : %s, port : %d, user : %s, password : %s, db_name : %s \n",
		y.MySQL.Host, y.MySQL.Port, y.MySQL.User, y.MySQL.Password, y.MySQL.DbName)

	fmt.Printf("Redis host : %s, port %d, auth : %s\n",
		y.Redis.Host, y.Redis.Port, y.Redis.Auth)

	fmt.Printf("Vip Counter: %d, Vip List : %v\n", y.NginxProxy.Counter, y.NginxProxy.NginxList)

	n, err := yaml.Marshal(y)
	if err != nil {
		log.Fatalf("marshal err : %v\n", err)
		return
	}
	fmt.Printf("yaml marshal : %v\n", string(n))
}
