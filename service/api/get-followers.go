package api

import (
	"encoding/json"
	"errors"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) getFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var page uint64
	var limit uint64
	var username string
	var followers []Follower
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

	dbfollowers, err := rt.db.SelectFollowers(username, page, limit)
	if errors.Is(err, database.ErrUserDoesNotExist) {
		ctx.Logger.WithError(err).Error("can't get followers")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		ctx.Logger.WithError(err).Error("can't get followers")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, dbfollower := range dbfollowers {
		var follower Follower
		follower.FromDatabase(dbfollower)
		followers = append(followers, follower)
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(followers)
	// Send the list to the user.
	w.Header().Set("Content-Type", "application/json")
}
