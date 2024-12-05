package handler

import (
	"gin/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetRoomById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id param is not correct")
		return
	}

	room, err := h.services.GetRoomById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, room)
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var input models.Room

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateRoom(input)
	if err != "" {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetAllRooms(c *gin.Context) {

}

func (h *Handler) DeleteRoom(c *gin.Context) {

}
