package api

import (
	"net/http"
	"strconv"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
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
	w.Write(image)

}
