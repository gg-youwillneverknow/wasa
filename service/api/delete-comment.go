package api

import (
	"errors"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	commentId, err := strconv.ParseUint(ps.ByName("commentId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err2 := rt.db.DeleteComment(commentId)
	if errors.Is(err2, database.ErrCommentDoesNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err2 != nil {
		ctx.Logger.WithError(err2).WithField("commentId", commentId).Error("can't delete the fountain")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
