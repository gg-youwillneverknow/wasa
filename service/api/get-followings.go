package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"errors"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getFollowings(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var page uint64
	var limit uint64
	var username string 
	username = ps.ByName("username")
	var err error
	
	if r.URL.Query().Has("page") {
		page, err = strconv.ParseUint(r.URL.Query().Get("page"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if page < 1 || page > 1000 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}
	if r.URL.Query().Has("limit") {
		limit, err = strconv.ParseUint(r.URL.Query().Get("limit"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if limit < 1 || limit > 100 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		limit = 20
	}

	followings, err := rt.db.SelectFollowings(username, page, limit)
	if errors.Is(err, database.ErrUserDoesNotExist) {
		ctx.Logger.WithError(err).Error("can't get followings")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		ctx.Logger.WithError(err).Error("can't get followings")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	
	// Send the list to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(followings)

}
