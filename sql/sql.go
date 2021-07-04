package sql

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
	"github.com/xuanxuan000/zzutil/log"
	"github.com/xuanxuan000/zzutil/model"
)

func GetMysqlConnection(ormPara string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", ormPara)
	if err != nil {
		return nil, errors.Wrap(err, ormPara)
	}
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 auto_increment=1")
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t" + defaultTableName
	}
	db.LogMode(false)
	return db, err
}

//连接指定数据库
func ConnectDatase(cfg model.MysqlCfg) (gormDB *gorm.DB, err error) {
	log.Debugf("%+v", cfg)
	mysqlPara := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&tls=skip-verify&allowNativePasswords=true",
		cfg.User,
		cfg.PassWord,
		cfg.IP,
		cfg.Port,
		cfg.DBname,
	)
	log.Debug(mysqlPara)
	gormDB, err = GetMysqlConnection(mysqlPara)
	if err != nil {
		log.Warningf("error(%v) connect mysql db info is (%v)", err, cfg)
		if strings.Contains(err.Error(), "server does not support TLS") {
			mysqlPara = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
				cfg.User,
				cfg.PassWord,
				cfg.IP,
				cfg.Port,
				cfg.DBname,
			)
			gormDB, err = GetMysqlConnection(mysqlPara)
			if err != nil {
				log.Errorf("error(%v) connect mysql db info is (%v)", err, cfg)
				return
			}
		} else {
			return
		}
	}
	log.Debug("connect db ok")
	return
}

//ConnectMysql 不指定数据库的链接
func ConnectMysql(cfg model.MysqlCfg) (gormDB *gorm.DB, err error) {
	// var gormDB *gorm.DB
	log.Debugf("%+v", cfg)
	mysqlPara := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=True&loc=Local&tls=skip-verify&allowNativePasswords=true",
		cfg.User,
		cfg.PassWord,
		cfg.IP,
		cfg.Port,
		// cfg.DBname,
	)
	gormDB, err = GetMysqlConnection(mysqlPara)
	if err != nil {
		log.Warningf("error(%v) connect mysql db info is (%v)", err, cfg)
		if strings.Contains(err.Error(), "server does not support TLS") {
			mysqlPara = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=True&loc=Local",
				cfg.User,
				cfg.PassWord,
				cfg.IP,
				cfg.Port,
				// cfg.DBname,
			)
			gormDB, err = GetMysqlConnection(mysqlPara)
			if err != nil {
				log.Errorf("error(%v) connect mysql db info is (%v)", err, cfg)
				return
			}
		} else {
			return
		}
	}
	log.Debug("connect mysql ok")
	return
}

//CreateDatabase 创建数据库
func CreateDatabase(cfg model.MysqlCfg) error {
	db, err := ConnectMysql(cfg)
	if err != nil {
		return err
	}
	SQL := fmt.Sprintf("CREATE DATABASE  IF NOT EXISTS %v default character set utf8mb4 collate utf8mb4_unicode_ci", cfg.DBname)
	err = db.Exec(SQL).Error
	if err != nil {
		return err
	}
	return err
}
