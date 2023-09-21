package api

import (
	"net/http"
	"fmt"
	"strconv"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/gofrs/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// httpRouterHandler is the signature for functions that accepts a reqcontext.RequestContext in addition to those
// required by the httprouter package.
type httpRouterHandler func(http.ResponseWriter, *http.Request, httprouter.Params, reqcontext.RequestContext)

// wrap parses the request and adds a reqcontext.RequestContext instance related to the request.
func (rt *_router) wrap(fn httpRouterHandler) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		endpoint := r.URL.Path

		if endpoint != "/session" {
			userId, err := strconv.ParseUint(r.Header.Get("token"), 10, 64)
			fmt.Println(token)
			fmt.Println(r.Header.Get("Authorization"))
			fmt.Println(r.Header.Get("token"))
			fmt.Println("trying")
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			fmt.Println(token)
			fmt.Println(userId)
			fmt.Println("wrapping")
			if (userId!=token){
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		reqUUID, err := uuid.NewV4()
		if err != nil {
			rt.baseLogger.WithError(err).Error("can't generate a request UUID")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var ctx = reqcontext.RequestContext{
			ReqUUID: reqUUID,
		}

		// Create a request-specific logger
		ctx.Logger = rt.baseLogger.WithFields(logrus.Fields{
			"reqid":     ctx.ReqUUID.String(),
			"remote-ip": r.RemoteAddr,
		})

		// Call the next handler in chain (usually, the handler function for the path)
		fn(w, r, ps, ctx)
	}
}
