package api

import (
	"errors"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var username string
	username = ps.ByName("username")
	var bannedusername string
	bannedusername = ps.ByName("bannedname")

	err := rt.db.DeleteBan(username, bannedusername)
	if errors.Is(err, database.ErrBanDoesNotExist) {
		ctx.Logger.WithError(err).Error("can't delete the ban")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if errors.Is(err, database.ErrUserDoesNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("can't delete the ban")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
