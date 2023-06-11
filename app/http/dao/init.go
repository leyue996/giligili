package dao

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var Db *gorm.DB

func Database(connRead string) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead,
		DefaultStringSize:         256,  //string类型默认长度
		DisableDatetimePrecision:  true, //禁止datetime精度
		DontSupportRenameIndex:    true, //重命名索引,就要把索引先删除在重建
		DontSupportRenameColumn:   true, //用change重名名列
		SkipInitializeWithVersion: false,
	}),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, //命名字段后面不加s

			}})
	if err != nil {
		return
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(20)  //设置连接池
	sqlDB.SetMaxIdleConns(100) //最大连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	Db = db
	//err = sqlDB.Ping()
	//if err != nil {
	//	panic(err)
	//}
	migration()

}

func NewDbClient(ctx context.Context) *gorm.DB {
	db := Db
	return db.WithContext(ctx)
}
