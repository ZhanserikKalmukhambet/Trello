package handler

import (
	"github.com/ZhanserikKalmukhambet/Trello/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var newUser entity.RegisterRequest

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := entity.User{
		Username: newUser.Username,
		Email:    newUser.Email,
		Password: newUser.Password,
	}
	id, err := h.srvs.Register(user)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"user_id": id})
}

func (h *Handler) signIn(ctx *gin.Context) {

}
