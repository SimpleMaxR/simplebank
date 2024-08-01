package api

import "github.com/gin-gonic/gin"

type createUserRequest struct {
	Username string `json:"username" binding:"required,min=4,max=30"`
	Password string `json:"password" binding:"required,min=8,max=40"`
}

func (server *Server) createUser(ctx *gin.Context) {
	// TODO
}
