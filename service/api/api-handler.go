package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.PUT("/accounts/:userId/edit", rt.wrap(rt.updateAccount))
	rt.router.GET("/users/:username/profile", rt.wrap(rt.getProfile))
	rt.router.GET("/users/:username/photos", rt.wrap(rt.getPhotos))
	rt.router.POST("/users/:username/photos", rt.wrap(rt.createPhoto))
	rt.router.GET("/images/:photoId", rt.wrap(rt.getImage))
	rt.router.GET("/users/:username/stream", rt.wrap(rt.getStream))
	rt.router.DELETE("/users/:username/photos/:photoId", rt.wrap(rt.deletePhoto))
	rt.router.GET("/users/:username/photos/:photoId/comments", rt.wrap(rt.getComments))
	rt.router.POST("/users/:username/photos/:photoId/comments", rt.wrap(rt.createComment))
	rt.router.DELETE("/users/:username/photos/:photoId/comments/:commentId", rt.wrap(rt.deleteComment))
	rt.router.GET("/users/:username/photos/:photoId/likes", rt.wrap(rt.getLikes))
	rt.router.PUT("/users/:username/photos/:photoId/likes/:userId", rt.wrap(rt.updateLike))
	rt.router.DELETE("/users/:username/photos/:photoId/likes/:userId", rt.wrap(rt.deleteLike))
	rt.router.DELETE("/users/:username/followers/{followername}", rt.wrap(rt.deleteFollower))
	rt.router.GET("/users/:username/followers/", rt.wrap(rt.getFollowers))
	rt.router.GET("/users/:username/followings/", rt.wrap(rt.getFollowings))
	rt.router.PUT("/users/:username/followings/{followingname}", rt.wrap(rt.updateFollowings))
	rt.router.DELETE("/users/:username/followings/{followingname}", rt.wrap(rt.deleteFollowing))
	rt.router.GET("/users/:username/bans/", rt.wrap(rt.getBans))
	rt.router.PUT("/users/:username/bans/{bannedname}", rt.wrap(rt.updateBan))
	rt.router.DELETE("/users/:username/bans/{bannedname}", rt.wrap(rt.deleteBan))
	rt.router.POST("/session", rt.wrap(rt.doLogin))
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
