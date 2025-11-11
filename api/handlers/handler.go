package handlers

import (
	"workerbee/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Services *services.Services
	Router   *gin.Engine
}
