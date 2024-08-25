package handler

import (
	"campaign/fundwave"
	"campaign/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type fundwaveHandler struct {
	service fundwave.Service
}

func NewfundwaveHandler(service fundwave.Service) *fundwaveHandler {
	return &fundwaveHandler{service}
}

// func (h *fundwaveHandler) Getfundwave(c *gin.Context) {
// 	fundwave, err := h.service.Findfundwave(0)
// 	if err != nil {
// 		response := helper.APIResponse("Error to get fundwave", http.StatusBadRequest, "error", nil)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response := helper.APIResponse("List of fundwave", http.StatusOK, "success", fundwave)
// 	c.JSON(http.StatusOK, response)
// }

func (h *fundwaveHandler) Getfundwave(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	fundwave, err := h.service.Getfundwave(userID)
	if err != nil {
		response := helper.APIResponse("Error to get fundwave", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of fundwave", http.StatusOK, "success", fundwave)
	c.JSON(http.StatusOK, response)
}