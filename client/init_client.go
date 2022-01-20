package client

import (
	"log"
	"yblog/model"

	"github.com/MonkeyIsMe/MyTool/client"
	"gorm.io/gorm"
)

var MysqlCli *gorm.DB

// InitClient 初始化一个MySQL的client
// todo 改成从配置文件读取
func InitClient() error {
	var err error
	proxy := client.MysqlProxy{
		Account:  "root",
		Password: "421992850",
		IP:       "127.0.0.1",
		Port:     "3306",
		DBName:   "myblog",
	}
	MysqlCli, err = client.NewClient(proxy)
	if err != nil {
		log.Fatalf("Init DB Client Error %+v", err)
		return err
	}

	err = model.SetClient(MysqlCli)
	if err != nil {
		log.Fatalf("Set DB Client Error %+v", err)
		return err
	}

	return nil
}
