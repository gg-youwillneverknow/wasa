package api

import (
	"encoding/json"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) createComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var comment Comment

	photoId, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&comment)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !comment.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbcomment, err2 := rt.db.CreateComment(photoId, comment.ToDatabase())
	if err2 != nil {
		ctx.Logger.WithError(err2).Error("can't create the comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	comment.FromDatabase(dbcomment)
	w.WriteHeader(201)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(comment)
}
