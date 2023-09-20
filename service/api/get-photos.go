package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error
	var photos []Photo
	var username string
	var page uint64
	var limit uint64

	username = ps.ByName("username")
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

	dbphotos, err := rt.db.SelectPhotos(username, page, limit)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't get photos")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, dbphoto := range dbphotos {
		var photo Photo
		photo.FromDatabase(dbphoto)

		photos = append(photos, photo)
	}
	w.WriteHeader(http.StatusOK)
	// Send the list to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photos)

}
