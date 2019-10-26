package config

import (
	"github.com/YJinHai/MyIm/internal/pkg/snowflake"
	"strings"

	"github.com/spf13/viper"

	"github.com/YJinHai/MyIm/internal/pkg/mysql"
	"github.com/YJinHai/MyIm/internal/pkg/nats"
)

type Config struct {
	Name string

}

func Init(cfg string) error{

	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	snowflake.SetMachineId(viper.GetInt64("app.machine_id"))

	c.initDB()
	c.initMQ()

	return nil

}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("conf") // 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")     // 设置配置文件格式为YAML
	viper.AutomaticEnv()            // 读取匹配的环境变量
	viper.SetEnvPrefix("APISERVER") // 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	return nil
}

func (c *Config) initDB() {
	db := mysql.NewDatabase()
	dbCfg := mysql.DBInfo{
		Username:  viper.GetString("db.username"),
		Password:	viper.GetString("db.password"),
		Addr:viper.GetString("db.addr"),
		DBName:viper.GetString("db.name"),
	}
	db.Init(&dbCfg)
}

func (c *Config) initMQ() {
	mq := nats.NewNats()
	mqCfg := nats.MQInfo{
		ClusterId:  viper.GetString("nats.cluster_id"),
		ClientId:	viper.GetString("nats.client_id"),
		Url:viper.GetString("nats.url"),
	}
	mq.Init(&mqCfg)
}

