package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) getImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var image []byte

	photoId, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	image, err = rt.db.SelectImage(photoId)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't get image")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	// Send the list to the user.
	w.Header().Set("Content-Type", "image/jpeg")
	_, err = w.Write(image)
	if err != nil {
		// Handle the error. You can log it, send an error response, etc.
		http.Error(w, "Failed to write image data", http.StatusInternalServerError)
		return
	}

}
