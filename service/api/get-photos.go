package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Parse the query string part. To do that, we need to check whether the latitude, longitude and range exists.
	// If latitude and longitude are specified, we parse them, and we filter results for them. If range is specified,
	// the value will be parsed and used as a filter. If it's not specified, 10 will be used as default (as specified in
	// the OpenAPI file).
	// If one of latitude or longitude is not specified (or both), no filter will be applied.

	var err error
	var posts []Photo
	var username string
	var page uint64
	var limit uint64
	var sort string
	username = ps.ByName("username")
	if r.URL.Query().Has("page") {
		page, err = strconv.ParseUint(r.URL.Query().Get("page"), 10, 64)
		if err != nil {
			// The latitude is not a valid float, or it's out of range, reject the request
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if page < 1 || page > 1000 {
			// The value is out of range, reject the request
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}
	if r.URL.Query().Has("limit") {
		limit, err = strconv.ParseUint(r.URL.Query().Get("limit"), 10, 64)
		if err != nil {
			// The longitude is not a valid float, or it's out of range, reject the request
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if limit < 1 || limit > 100 {
			// The value is out of range, reject the request
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		limit = 20
	}
	if r.URL.Query().Has("sort") {
		sort = r.URL.Query().Get("sort")
		if sort != "chronological" || sort != "reverse chronological" {
			// The longitude is not a valid float, or it's out of range, reject the request
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		sort = "reverse chronological"
	}

	photos, err := rt.db.SelectPhotos(username, page, limit, sort)

	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't get photos")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, p := range photos {
		var post Photo
		var comments []Comment
		var likes []Like
		var page uint64 = 1
		var limit uint64 = 20
		dbcomments, err := rt.db.SelectComments(p.ID, page, limit)
		if err != nil {
			// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
			// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
			ctx.Logger.WithError(err).Error("can't get photos")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		dblikes, err := rt.db.SelectLikes(p.ID, page, limit)
		if err != nil {
			// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
			// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
			ctx.Logger.WithError(err).Error("can't get photos")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		for _, c := range dbcomments {
			var comment Comment
			comment.FromDatabase(c)
			comments = append(comments, comment)
		}

		for _, l := range dblikes {
			var like Like
			like.FromDatabase(l)
			likes = append(likes, like)
		}
		post.FromDatabase(p)
		post.Comments = comments
		post.Likes = likes
		posts = append(posts, post)
	}

	// Send the list to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(posts)

}
