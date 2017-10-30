package auth

import (
	"github.com/xuyuntech/wechatshop/pkg/manager"
	"net/http"
	"github.com/xuyuntech/wechatshop"
	"golang.org/x/net/context"
)

type AuthRequired struct {
	manager manager.Manager
}

func NewAuthRequired(m manager.Manager) *AuthRequired {
	return &AuthRequired{
		manager: m,
	}
}

func (a *AuthRequired) HandlerFuncWithNext(w http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	token := r.Header.Get("X-Access-Token")
	data, err := a.manager.VerifyAuthToken(token)
	if err != nil {
		xuyuntech.Resp(w, 1,nil,err.Error())
		return
	}
	if next != nil {
		next(w ,r.WithContext(context.WithValue(r.Context(), "user", data)))
	}
}