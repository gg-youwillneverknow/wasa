package api

import (
	"errors"
	"net/http"
	"fmt"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) updateFollowings(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var username string
	username = ps.ByName("username")

	var followingusername string
	followingusername = ps.ByName("followingname")
	fmt.Println("here")
	err := rt.db.UpdateFollowings(username, followingusername)
	fmt.Println("following attempted")
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
