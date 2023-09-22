package api

import (
	"encoding/json"
	"errors"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) updateAccount(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var updatedInfo User
	var ret database.User

	err := json.NewDecoder(r.Body).Decode(&updatedInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Update the User in the database.
	ret, err = rt.db.UpdateAccount(updatedInfo.ToDatabase())
	if errors.Is(err, database.ErrUserDoesNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("id", updatedInfo.ID).Error("can't update the username")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	updatedInfo.FromDatabase(ret)
	w.WriteHeader(http.StatusOK)

	// Send the new username to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(updatedInfo)
}
