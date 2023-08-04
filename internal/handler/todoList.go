package handler

import (
	"github.com/ZhanserikKalmukhambet/Trello/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createList(ctx *gin.Context) {
	userID, err := getAuthUserID(ctx)
	if err != nil {
		return
	}

	var input entity.TodoList
	if err := ctx.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.srvs.CreateTodoList(userID, input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"todo_list_id": id})
}

func (h *Handler) getLists(ctx *gin.Context) {

}

func (h *Handler) getListByID(ctx *gin.Context) {

}

func (h *Handler) updateList(ctx *gin.Context) {

}

func (h *Handler) removeList(ctx *gin.Context) {

}
