package boot

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jlb_shop_go/global"
)

//@author: SliverHorn
//@function: Gorm
//@description: 初始化数据库并产生数据库全局变量
//@return: *gorm.DB

func Gorm() *gorm.DB {
	switch global.Config.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

//@author: SliverHorn
//@function: GormMysql
//@description: 初始化Mysql数据库
//@return: *gorm.DB

func GormMysql() *gorm.DB {
	m := global.Config.Mysql
	if m.Dbname == "" {
		return nil
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		//global.GVA_LOG.Error("MySQL启动异常", zap.Any("err", err))
		//os.Exit(0)
		//return nil
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}

}

//@author: SliverHorn
//@function: gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config

//func gormConfig() *gorm.Config {
//	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
//	switch global.Config.Mysql.LogMode {
//	case "silent", "Silent":
//		config.Logger = logger.Default.LogMode(logger.Silent)
//	case "error", "Error":
//		config.Logger = logger.Default.LogMode(logger.Error)
//	case "warn", "Warn":
//		config.Logger = logger.Default.LogMode(logger.Warn)
//	case "info", "Info":
//		config.Logger = logger.Default.LogMode(logger.Info)
//	default:
//		config.Logger = logger.Default.LogMode(logger.Info)
//	}
//	return config
//}
