package manager

import (
	"github.com/go-xorm/xorm"
	mAuth "github.com/xuyuntech/wechatshop/pkg/modules/auth"
	"errors"
	"github.com/xuyuntech/wechatshop/pkg/config"
)

type Manager interface {
	VerifyAuthToken(token string) (map[string]interface{}, error)
}

type defaultManager struct {
	engine *xorm.Engine
	config *config.Config
}

func NewDefaultManager(config *config.Config, engine *xorm.Engine) (Manager, error) {
	return &defaultManager{
		config: config,
		engine: engine,
	}, nil
}


func (m defaultManager) VerifyAuthToken(token string) (map[string]interface{}, error) {
	errFailed := errors.New("token 解析失败")
	success, data, err := mAuth.AuthToken(m.config.Auth.TokenKey, token)
	if err != nil {
		return nil, err
	}
	_username := data["username"]
	if _username == nil {
		return nil, errFailed
	}
	_uid := data["uid"]
	if _uid == nil {
		return nil, errFailed
	}
	if !success {
		return nil, errFailed
	}

	return map[string]interface{}{
		"uid": _uid.(string),
		"username": _username.(string),
	}, nil
}
