package common

import "github.com/micro/go-micro/v2/config"

type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

//获取mysql配置
func GetMysqlConfig(config config.Config, path ...string) *MysqlConfig {
	mysqlConfig := &MysqlConfig{}
	config.Get(path...).Scan(mysqlConfig)

	return mysqlConfig
}
