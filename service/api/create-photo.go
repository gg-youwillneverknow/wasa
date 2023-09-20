package api

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"errors"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) createPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Read the new content for the fountain from the request body.
	var photo Photo
	var data []byte
	var err error
	username := ps.ByName("username")
	
	if err:= r.ParseMultipartForm(32 << 20); err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	file, _, err := r.FormFile("photos")
	if err != nil {
		http.Error(w, "Failed to retrieve uploaded file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	// Read the file data
	data, err = ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file data", http.StatusInternalServerError)
		return
	}

	dbuserId, err := rt.db.SelectUser(username)
	if errors.Is(err, database.ErrUserDoesNotExist) {
		ctx.Logger.WithError(err).Error("can't get the user")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err!=nil {
		ctx.Logger.WithError(err).Error("can't get the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dbphoto, err := rt.db.CreatePhoto(photo.ToDatabase(data,dbuserId,username))
	if err != nil {
		ctx.Logger.WithError(err).Error("can't create the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photo.FromDatabase(dbphoto)

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photo)

	w.WriteHeader(201)
}
