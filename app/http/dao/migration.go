package dao

//执行数据迁移

func migration() {
	// 自动迁移模式
	Db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{})
}
