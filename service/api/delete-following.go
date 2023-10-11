package api

import (
	"errors"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) deleteFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	username := ps.ByName("username")

	followingusername := ps.ByName("followingname")

	err := rt.db.DeleteFollowing(username, followingusername)
	if errors.Is(err, database.ErrFollowingDoesNotExist) {
		ctx.Logger.WithError(err).Error("can't delete the following")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if errors.Is(err, database.ErrUserDoesNotExist) {
		ctx.Logger.WithError(err).Error("can't delete the following")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("followingusername", followingusername).Error("can't delete the following")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
