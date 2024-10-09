package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/refluxwhw/nas/pkg/conf"
)

var gDB *gorm.DB

// DatabaseConf db config, driver can be mysql/sqlite,
//   - driver: mysql
//     dsn: <user>:<passwd>@tcp(<host>:<port>)/<dbname>?charset=utf8mb4&parseTime=True&loc=Local
//   - driver: sqlite
//     dsn: <file-path>
type DatabaseConf struct {
	Driver string
	Dsn    string
}

func Setup() (err error) {
	cfg := conf.GetConfig()
	dbConf := &DatabaseConf{}
	err = cfg.UnmarshalKey("db", dbConf)
	if err != nil {
		return fmt.Errorf("get database config error: %v", err)
	}

	switch dbConf.Driver {
	case "sqlite":
		gDB, err = gorm.Open(sqlite.Open(dbConf.Dsn), &gorm.Config{})

	case "mysql":
		gDB, err = gorm.Open(mysql.Open(dbConf.Dsn), &gorm.Config{})

	default:
		return fmt.Errorf("unsupported database driver: %s", dbConf.Driver)
	}

	// check
	conn, err := gDB.DB()
	if err != nil {
		return fmt.Errorf("get connection error: %v", err)
	}

	if err = conn.Ping(); err != nil {
		return fmt.Errorf("test ping error: %v", err)
	}

	return nil
}

func GetORM() *gorm.DB {
	return gDB
}
