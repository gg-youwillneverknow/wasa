package api

import (
	"encoding/json"
	"errors"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var username string
	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbuserId, err2 := rt.db.SelectUser(username)

	// Update the fountain in the database.
	if errors.Is(err2, database.ErrUserDoesNotExist) {

		// The fountain (indicated by `id`) does not exist, reject the action indicating an error on the client side.
		dbuserId, err2 = rt.db.CreateUser(username)
		if err2 != nil {
			// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
			// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
			ctx.Logger.WithError(err2).Error("can't create userId")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		_ = json.NewEncoder(w).Encode(dbuserId)
		return
	} else if err2 == nil {

		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(dbuserId)
		return
	} else {

		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err2).Error("can't select userId")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

}
