package handlers

import "github.com/19jmrs/youtube-bookmarks/db"

type UserHandler struct {
	userStore db.UserStore
}


func NewUserHandler(userStore db.UserStore) *UserHandler{
	return &UserHandler{
		userStore: userStore,
	}
}