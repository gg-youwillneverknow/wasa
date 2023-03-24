package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) createPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Read the new content for the fountain from the request body.
	var photo Photo
	var data []byte
	var err error
	var username string
	username = ps.ByName("username")
	data, err = ioutil.ReadAll(r.Body)
	// You have to manually close the body, check docs
	// This is required if you want to use things like
	// Keep-Alive and other HTTP sorcery.
	r.Body.Close()
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Create the fountain in the database. Note that this function will return a new instance of the fountain with the
	// same information, plus the ID.
	dbphoto, err := rt.db.CreatePhoto(username, photo.ToDatabase(data))
	if err != nil {
		// In this case, we have an error on our side. sLog the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't create the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Here we can re-use `fountain` as FromDatabase is overwriting every variabile in the structure.
	photo.FromDatabase(dbphoto)

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photo)
}
