package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/h-u-m-a-n/todo-app"
)

func (h *Handler) createList(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil{
		return
	}

	var input todo.TodoList
	if err:= c.BindJSON(&input); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, input)
	if err!=nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":id,
	})
}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil{
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err!=nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}


func (h *Handler) getListById(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil{
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetListByID(userId, id)
	if err!=nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context){
	
}

func (h *Handler) deleteList(c *gin.Context){
	
}