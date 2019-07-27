package model

func migration(){
	//根据Model自动迁移数据库表格
	DB.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8").AutoMigrate(&User{})
}
