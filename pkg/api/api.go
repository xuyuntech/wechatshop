package api

import (
	"github.com/xuyuntech/wechatshop/pkg/config"
	"github.com/xuyuntech/wechatshop"
	tokenAuth "github.com/xuyuntech/wechatshop/pkg/middleware/auth_token"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/context"
	"github.com/xuyuntech/wechatshop/pkg/models"
	"github.com/xuyuntech/wechatshop/pkg/manager"
)

type api struct {
	config *config.Config
	manager manager.Manager
}

type ApiConfig struct {
	Config *config.Config
	Manager manager.Manager
}

func NewApi(config *ApiConfig) (*api, error) {
	return &api{
		config: config.Config,
		manager: config.Manager,
	}, nil
}

func (a *api) Run() error {
	cs := xuyuntech.Cors
	apiAuthRequired := tokenAuth.NewAuthTokenRequired(a.config.Auth.TokenAuthURL)

	globalMux := http.NewServeMux()

	homeRouter := mux.NewRouter()
	homeRouter.HandleFunc("/v1/home/auth/xxx", a.ServiceNeedAuth).Methods("POST")

	authRouter := negroni.New()
	authRouter.Use(negroni.HandlerFunc(apiAuthRequired.HandlerFuncWithNext))
	authRouter.UseHandler(homeRouter)
	globalMux.Handle("/v1/home/auth/", authRouter)

	pubRouter := mux.NewRouter()
	pubRouter.HandleFunc("/v1/home/pub/xxx", a.ServicePublic).Methods("GET")
	globalMux.Handle("/v1/home/pub/", pubRouter)

	s := &http.Server{
		Addr: a.config.Server.Addr,
		Handler: context.ClearHandler(cs.Handler(globalMux)),
	}

	logrus.Infof("Server run: %s", a.config.Server.Addr)

	return s.ListenAndServe()
}


func (a *api) getUserInfo(r *http.Request) *models.User {
	us := r.Context().Value("user")
	if us == nil {
		return nil
	}
	ctx := us.(map[string]interface{})
	if ctx == nil || ctx["uid"] == nil {
		return nil
	}
	return &models.User{
		Id: ctx["uid"].(string),
		Name: ctx["username"].(string),
	}
}