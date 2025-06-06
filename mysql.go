package arpc

import (
	"fmt"
	"log"
	"os"

	"github.com/Houserqu/arpc/gorm_ext"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Mysql *gorm.DB
var Mysql1 *gorm.DB // 等同于 Mysql
var Mysql2 *gorm.DB
var Mysql3 *gorm.DB

func InitMysql() {
	// 注册自定义序列化器
	schema.RegisterSerializer("timestamp", gorm_ext.TimestampSerializer{})

	if viper.GetString("mysql.host") != "" && !viper.GetBool("mysql.disable") {
		Mysql = NewMysql(MysqlConfig{
			Host:     viper.GetString("mysql.host"),
			Port:     viper.GetString("mysql.port"),
			User:     viper.GetString("mysql.user"),
			Password: viper.GetString("mysql.password"),
			Database: viper.GetString("mysql.database"),
		})
		Mysql1 = Mysql
		log.Println("mysql connect success")
	}

	if viper.GetString("mysql2.host") != "" && !viper.GetBool("mysql2.disable") {
		Mysql2 = NewMysql(MysqlConfig{
			Host:     viper.GetString("mysql2.host"),
			Port:     viper.GetString("mysql2.port"),
			User:     viper.GetString("mysql2.user"),
			Password: viper.GetString("mysql2.password"),
			Database: viper.GetString("mysql2.database"),
		})
		log.Println("mysql2 connect success")
	}

	if viper.GetString("mysql3.host") != "" && !viper.GetBool("mysql3.disable") {
		Mysql2 = NewMysql(MysqlConfig{
			Host:     viper.GetString("mysql3.host"),
			Port:     viper.GetString("mysql3.port"),
			User:     viper.GetString("mysql3.user"),
			Password: viper.GetString("mysql3.password"),
			Database: viper.GetString("mysql3.database"),
		})
		log.Println("mysql3 connect success")
	}
}

type MysqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func NewMysql(config MysqlConfig) *gorm.DB {
	var dbLogger logger.Interface

	if viper.GetBool("dev") {
		// 开发环境使用自定义日志配置
		dbLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				LogLevel: logger.Info, // 日志级别
				Colorful: false,       // 禁用彩色打印
			},
		)
	} else {
		// 生产环境适应默认配置（只输出慢 sql 和错误日志）
		dbLogger = logger.Default
	}

	dsn := fmt.Sprint(config.User, ":", config.Password, "@tcp(", config.Host, ":", config.Port, ")/", config.Database, "?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{
		Logger: dbLogger,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
