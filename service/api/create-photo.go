package api

import (
	"encoding/json"
	"errors"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func (rt *_router) createPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Read the new content for the fountain from the request body.
	var photo Photo
	var data []byte

	username := ps.ByName("username")

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, _, err2 := r.FormFile("photos")
	if err2 != nil {
		http.Error(w, "Failed to retrieve uploaded file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	// Read the file data
	data, err2 = io.ReadAll(file)
	if err2 != nil {
		http.Error(w, "Failed to read file data", http.StatusInternalServerError)
		return
	}

	dbuserId, err3 := rt.db.SelectUser(username)
	if errors.Is(err3, database.ErrUserDoesNotExist) {
		ctx.Logger.WithError(err3).Error("can't get the user")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err3 != nil {
		ctx.Logger.WithError(err3).Error("can't get the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dbphoto, err4 := rt.db.CreatePhoto(photo.ToDatabase(data, dbuserId, username))
	if err4 != nil {
		ctx.Logger.WithError(err4).Error("can't create the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photo.FromDatabase(dbphoto)
	w.WriteHeader(201)
	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photo)

}
