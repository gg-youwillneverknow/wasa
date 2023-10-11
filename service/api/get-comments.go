package api

import (
	"encoding/json"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var page uint64
	var limit uint64
	var comments []Comment

	photoId, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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

	dbcomments, err2 := rt.db.SelectComments(photoId, page, limit)
	if err2 != nil {
		ctx.Logger.WithError(err2).Error("can't get comments")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, dbcomment := range dbcomments {
		var comment Comment
		comment.FromDatabase(dbcomment)
		comments = append(comments, comment)
	}

	w.WriteHeader(http.StatusOK)

	// Send the list to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(comments)

}
