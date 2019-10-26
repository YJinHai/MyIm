package mysql

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/lexkong/log"
	"net/url"
	"sync"
)


var Info *DBInfo
//var DB *Database

type Database interface {
	Init(dbInfo *DBInfo)
	Close()
}

type DBInfo struct {
	Mu *sync.RWMutex
	Username string  `yaml:"username"`
	Password string `yaml:"password"`
	Addr string `yaml:"addr"`
	DBName string `yaml:"db_name"`
}

type database struct {
	Self   *xorm.Engine
	Docker *xorm.Engine
}

func NewDatabase() Database {
	return &database{}
}

func (db *database) Init(dbInfo *DBInfo) {
	Info = &DBInfo{
		Mu: &sync.RWMutex{},
		Username:dbInfo.Username,
		Password:dbInfo.Password,
		Addr:dbInfo.Addr,
		DBName:dbInfo.DBName,
	}

	db.Self = GetSelfDB()
	db.Docker = GetDockerDB()
}

func (db *database) Close() {
	db.Self.Close()
	db.Docker.Close()
}

func openDB(username, password, addr, dbName string) *xorm.Engine {
	timezone :=  "'+8:00'"   //'Asia/Shanghai'"
	fmt.Printf("username, password, addr, dbName:",username, password, addr, dbName)
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&time_zone=%s&loc=%s",
		username,
		password,
		addr,
		dbName,
		true,
		url.QueryEscape(timezone),
		"Asia%2FShanghai")

	db, err := xorm.NewEngine("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", dbName)
	}

	// set for db connection
	setupDB(db)

	return db
}

func setupDB(db *xorm.Engine) {
	//db.LogMode(GetBool("gormlog"))
	db.DB().SetMaxOpenConns(20000) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxIdleConns(0) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
}

func MysqlMid(username,password,addr,dbName string) gin.HandlerFunc {
	return func(c *gin.Context){
		openDB(username,password,addr,dbName)
	}
}

// used for cli
func InitSelfDB(username,password,addr,dbName string) *xorm.Engine {
	return openDB(username,password,addr,dbName)
}

func GetSelfDB() *xorm.Engine {

	Info.Mu.RLock()
	defer Info.Mu.RUnlock()
	return InitSelfDB(Info.Username,Info.Password,Info.Addr,Info.DBName)
}

func InitDockerDB(username,password,addr,dbName string) *xorm.Engine{
	return openDB(username,password,addr,dbName)
}

func GetDockerDB() *xorm.Engine {
	Info.Mu.RLock()
	defer Info.Mu.RUnlock()
	return InitDockerDB(Info.Username,Info.Password,Info.Addr,Info.DBName)
}



