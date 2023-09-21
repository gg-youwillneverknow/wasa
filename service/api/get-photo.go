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

func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error
	var photo Photo
	var photoId uint64
	var username string
	username = ps.ByName("username")
	photoId, err = strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbphoto, err := rt.db.SelectPhoto(photoId)
	dbphoto.Owner = username
	if errors.Is(err, database.ErrPhotoDoesNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("can't get photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photo.FromDatabase(dbphoto)

	w.WriteHeader(http.StatusOK)
	// Send the list to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photo)

}
