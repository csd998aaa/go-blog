package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Server struct {
	Port string `json:"port"`
	Domain string `json:"domain"`
}

type Mysql struct {
	Database string `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Password string `json:"password"`
}

type Config struct {
	Server `json:"server"`
	Mysql `json:"mysql"`
	DatabaseURL interface{}
}

var config Config

func GetConfig() *Config {
	// 读取json文件
	path,_ := filepath.Abs("conf")
	fi, err := os.Open(path+"/conf.json")
	if err != nil {
		log.Fatal(err)
		return &config
	}
	defer fi.Close()
	fd,err := ioutil.ReadAll(fi)
	json.Unmarshal([]byte(fd),&config)
	config.DatabaseURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Mysql.Port,
		config.Database,
	)
	return &config
}

