package global

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"jlb_shop_go/config"
)

const (
	SfaLoginKey     = "staff-"
	SfaLoginExpired = int64(86400)
)

var (
	Db      *gorm.DB
	Redis   *redis.Client
	Config  config.App
	Viper   *viper.Viper
	Ctx     = context.Background()
	StaffID int
)
