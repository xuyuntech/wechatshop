package auth_token

import (
	"net/http"
	"time"
	"encoding/json"
	"golang.org/x/net/context"
	"github.com/Sirupsen/logrus"
	"github.com/xuyuntech/wechatshop"
	"github.com/xuyuntech/wechatshop/pkg/modules/httplib"
)

type AuthTokenRequired struct {
	authUrl string
}

func NewAuthTokenRequired(url string) *AuthTokenRequired {
	return &AuthTokenRequired{
		authUrl: url,
	}
}

type resultAuth struct {
	Status int `json:"status"`
	Results interface{} `json:"results"`
	Msg string `json:"msg"`
}

func (a *AuthTokenRequired) HandlerFuncWithNext(w http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	token := r.Header.Get("X-Access-Token")
	if token == "" {
		xuyuntech.Resp(w, xuyuntech.STATUS_UNAUTHORIZED, nil)
		return
	}
	logrus.Debugf("AuthToken url: %s", a.authUrl)
	resp, err := httplib.Get(a.authUrl).Header("X-Access-Token", token).SetTimeout(time.Second*3, time.Second*3).Response()
	if err != nil {
		logrus.Errorf("AuthToken err: %v", err)
		xuyuntech.Resp(w, xuyuntech.STATUS_INTERNAL_SERVER_ERR,nil)
		return
	}
	var result resultAuth
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		xuyuntech.Resp(w, xuyuntech.STATUS_INTERNAL_SERVER_ERR,nil, err.Error())
		return
	}
	if result.Status == xuyuntech.STATUS_OK && next != nil {
		logrus.Debugf("AuthToken result: %+v", result)
		next(w, r.WithContext(context.WithValue(r.Context(), "user", result.Results)))
	} else {
		logrus.Debugf("AuthToken failed: %+v", result)
		xuyuntech.Resp(w, xuyuntech.STATUS_UNAUTHORIZED, nil)
	}
}

