package api

import (
	"errors"
	"net/http"
	"strconv"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	photoId, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = rt.db.DeletePhoto(photoId)
	if errors.Is(err, database.ErrPhotoDoesNotExist) {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("photoId", photoId).Error("can't delete the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
