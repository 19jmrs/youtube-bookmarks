package handlers

import (
	"context"
	"net/http"

	"github.com/19jmrs/youtube-bookmarks/db"
	"github.com/19jmrs/youtube-bookmarks/types"
	"github.com/19jmrs/youtube-bookmarks/views"
)

type UserHandler struct {
	userStore db.UserStore
}


func NewUserHandler(userStore db.UserStore) *UserHandler{
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	err := h.userStore.CreateUser(email, password)

}