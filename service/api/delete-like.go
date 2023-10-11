package api

import (
	"errors"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) deleteLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	photoId, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	likerId, err2 := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err2 = rt.db.DeleteLike(photoId, likerId)
	if errors.Is(err2, database.ErrLikeDoesNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err2 != nil {
		ctx.Logger.WithError(err2).WithField("likerId", likerId).Error("can't delete the like")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
