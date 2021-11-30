package wxMini

import (
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/xinzf/gokit/logger"
	"jlb_shop_go/global"
	"time"
)

const (
	WxUrl              = "https://api.weixin.qq.com"
	AccessTokenKey     = "wx:token"
	AccessTokenExpired = int64(7000)
)

type Server struct {
	appid, secret string
	Cache         _defaultCache
}

type _defaultCache struct {
}

func NewWxMiniServer(appid, secret string) *Server {
	return &Server{appid: appid, secret: secret}
}

type AccessTokenCache interface {
	Get(key string) (string, error)
	Set(key, token string, duration time.Duration) error
}

func (c *_defaultCache) Get(key string) (string, error) {
	key = fmt.Sprintf("%s:%s", AccessTokenKey, key)
	return global.Redis.Get(global.Ctx, key).Result()
}

func (c *_defaultCache) Set(key, token string, duration time.Duration) error {
	key = fmt.Sprintf("%s:%s", AccessTokenKey, key)
	return global.Redis.Set(global.Ctx, key, token, duration).Err()
}

// GetToken 获取access token
func (s *Server) GetToken() (string, error) {
	token, err := s.Cache.Get(AccessTokenKey)
	if token != "" {
		logger.DefaultLogger.Debug("wx:token", "命中缓存")
		return token, nil
	}
	logger.DefaultLogger.Debug("wx:token", "没有命中缓存")

	rsp := &AccessTokenRes{}
	_url := fmt.Sprintf("%s/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", WxUrl, s.appid, s.secret)

	_rsp, _, errs := gorequest.New().Get(_url).EndStruct(&rsp)
	if _rsp.StatusCode != 200 {
		return "", fmt.Errorf("微信服务器异常,httpStatus: %d", _rsp.StatusCode)
	}

	if len(errs) > 0 {
		return "", errs[0]
	}

	if rsp.ErrCode != 0 {
		return "", errors.New(rsp.ErrMsg)
	}

	err = s.Cache.Set(AccessTokenKey, rsp.AccessToken, time.Duration(AccessTokenExpired)*time.Second)
	if err != nil {
		return "", err
	}

	return rsp.AccessToken, nil
}

// GetCode2Session 获取小程序登录信息
func (s Server) GetCode2Session(code string) (rsp *Code2SessionRes, err error) {
	_url := fmt.Sprintf("%s/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", WxUrl, s.appid, s.secret, code)

	_rsp, _, errs := gorequest.New().Get(_url).EndStruct(&rsp)
	if _rsp.StatusCode != 200 {
		return nil, fmt.Errorf("微信服务器异常,httpStatus: %d", _rsp.StatusCode)
	}

	if len(errs) > 0 {
		return nil, errs[0]
	}

	if rsp.ErrCode != 0 {
		return nil, errors.New(rsp.ErrMsg)
	}

	return rsp, nil
}
