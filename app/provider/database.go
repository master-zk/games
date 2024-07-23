package provider

import (
	"fmt"
	"games/app/config"
	"games/app/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"reflect"
	"time"
)

type databaseProvider struct {
}

var DatabaseProvider = &databaseProvider{}

func (p *databaseProvider) Register() {
	global.DB.Jenny = p.registerDb(global.DB.Jenny, global.Config.Databases.Jenny)
	global.DB.Xh = p.registerDb(global.DB.Jenny, global.Config.Databases.Xh)
	fmt.Println("-----------------------------------")
	fmt.Printf("t = %T, v= %v, typeOf = %v", global.DB.Jenny, &global.DB.Jenny, reflect.TypeOf(global.DB.Jenny).Kind())
	fmt.Printf("t = %T, v= %v, typeOf = %v", global.DB.Xh, &global.DB.Xh, reflect.TypeOf(global.DB.Xh).Kind())
	fmt.Println("-----------------------------------")
	//	p.registerDb(global.DB.Xh, global.Config.Databases.Xh)
	//	p.registerDb(global.DBJenny, global.Config.Databases.Jenny)
	//	p.registerDb(global.DBXh, global.Config.Databases.Xh)
}

func (p *databaseProvider) registerDb(db *gorm.DB, dbConfig config.DatabaseConfig) *gorm.DB {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database)

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，Mysql 5.6 之前不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		CreateBatchSize: 20, // 默认创建批大小
		NamingStrategy: schema.NamingStrategy{ // 表、列命名策略
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("failed to connect database: %v \n", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(dbConfig.Pool.MaxIdleConn)
	sqlDB.SetMaxOpenConns(dbConfig.Pool.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Duration(dbConfig.Pool.MaxLifetime))
	return db
}
