package db

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type dbConfig struct {
	Endpoint string `yaml:"endpoint"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Pw       string `yaml:"pw"`
	Db       string `yaml:"db"`
}

const cfgPath = "./config.yaml"

var (
	cfg *dbConfig
	db  *gorm.DB
)

func init() {
	fileBytes, err := os.ReadFile(cfgPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = yaml.Unmarshal(fileBytes, &cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	dsn := "host=" + cfg.Endpoint + " user=" + cfg.User + " password=" + cfg.Pw + " dbname=" + cfg.Db + " port=" + cfg.Port + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func InstanceDB() *gorm.DB {
	return db
}
