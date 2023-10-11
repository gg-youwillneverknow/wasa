package api

import (
	"errors"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) updateFollowings(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	username := ps.ByName("username")

	followingusername := ps.ByName("followingname")

	err := rt.db.UpdateFollowings(username, followingusername)

	if errors.Is(err, database.ErrFollowingAlreadyExist) {
		ctx.Logger.WithError(err).Error("can't update the following")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if errors.Is(err, database.ErrUserDoesNotExist) {
		ctx.Logger.WithError(err).Error("can't update the following")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		ctx.Logger.WithError(err).Error("can't update the following")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
