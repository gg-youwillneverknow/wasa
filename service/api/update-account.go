package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) updateAccount(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// The UserID in the path is a string. Let's parse it.
	id, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	var updatedInfo User
	// Read the new username for the UserInfo from the request query.
	username := r.URL.Query().Get("newusername")
	updatedInfo.Username = username
	// The client is not supposed to send us the ID in the body, as the UserID is already specified in the path,
	// and it's immutable. So, here we overwrite the ID in the JSON with the `id` variable (that comes from the URL).
	updatedInfo.ID = id
	var ret database.User
	// Update the UserInfo in the database.
	ret, err = rt.db.UpdateAccount(updatedInfo.ToDatabase())
	if errors.Is(err, database.ErrUserDoesNotExist) {
		// The user (indicated by `id`) does not exist, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		// Note (2): we are adding the error and an additional field (`id`) to the log entry, so that we will receive
		// the identifier of the user that triggered the error.
		ctx.Logger.WithError(err).WithField("id", id).Error("can't update the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	updatedInfo.FromDatabase(ret)

	// Send the list to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(updatedInfo)
}
