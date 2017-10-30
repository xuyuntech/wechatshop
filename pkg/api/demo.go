package api

import "net/http"

func (a *api) ServiceNeedAuth(w http.ResponseWriter, r *http.Request) {}
func (a *api) ServicePublic(w http.ResponseWriter, r *http.Request) {}
