package drivers

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"street_stall/biz/config"
	"street_stall/biz/domain/model"
)

var (
	DB *gorm.DB
)

func InitMySQL(config *config.BasicConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.Mysql.UserName, config.Mysql.PassWorld, config.Mysql.Host, config.Mysql.Port, config.Mysql.DB)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("[system][initMysql] open mysql error, dsn=%s", dsn)
	}
	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(1000)

	// 启用Logger，显示详细日志
	db.LogMode(true)

	DB = db
	log.Print("[system][mysql] init mysql driver success!")

	// 数据库建表
	createTable()

	//user := model.User{Username: "lt15991075603", Password: "litao.", UserIdentity:1}
	//DB.Create(&user)

}

func createTable() {
	db := DB
	if !db.HasTable(&model.User{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.User{})
		log.Print("[system][mysql][createTable] create table `user`")
	}
	if !db.HasTable(&model.Visitor{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.Visitor{})
		db.Model(&model.Visitor{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
		log.Print("[system][mysql][createTable] create table `visitor`")
	}
	if !db.HasTable(&model.Merchant{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.Merchant{})
		db.Model(&model.Merchant{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
		log.Print("[system][mysql][createTable] create table `merchant`")
	}
	if !db.HasTable(&model.Place{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.Place{})
		log.Print("[system][mysql][createTable] create table `place`")
	}
	if !db.HasTable(&model.Location{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.Location{})
		db.Model(&model.Location{}).AddForeignKey("place_id", "places(id)", "RESTRICT", "RESTRICT")
		log.Print("[system][mysql][createTable] create table `location`")
	}
	if !db.HasTable(&model.Order{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.Order{})
		db.Model(&model.Order{}).AddForeignKey("location_id", "locations(id)", "RESTRICT", "RESTRICT")
		db.Model(&model.Order{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
		log.Print("[system][mysql][createTable] create table `order`")
	}
	if !db.HasTable(&model.Evaluation{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.Evaluation{})
		db.Model(&model.Evaluation{}).AddForeignKey("merchant_id", "merchants(id)", "RESTRICT", "RESTRICT")
		db.Model(&model.Evaluation{}).AddForeignKey("visitor_id", "visitors(id)", "RESTRICT", "RESTRICT")
		log.Print("[system][mysql][createTable] create table `evaluation`")
	}
	if !db.HasTable(&model.Question{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.Question{})
		db.Model(&model.Question{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
		log.Print("[system][mysql][createTable] create table `question`")
	}
	log.Print("[system][mysql] create table success!")
}
