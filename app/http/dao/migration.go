package dao

import "giligili/app/http/model"

//执行数据迁移

func migration() {
	// 自动迁移模式
	err := Db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.User{},
			&model.Video{},
		)

	if err != nil {
		panic(err)
	}
	return
}
