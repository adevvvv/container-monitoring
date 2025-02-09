package api

import (
	"container-monitoring/backend/model"
	"container-monitoring/backend/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	pingService service.PingService
}

func NewHandler(ps service.PingService) *Handler {
	return &Handler{pingService: ps}
}

// ListStatuses godoc
// @Summary Get all ping statuses
// @Description Retrieve all ping statuses from the system
// @Tags statuses
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "List of ping statuses"
// @Failure 500 {object} map[string]interface{} "Database error"
// @Router /api/v1/status [get]
func (h *Handler) ListStatuses(c *gin.Context) {
	statuses, err := h.pingService.GetAll()
	if err != nil {
		sendError(c, http.StatusInternalServerError, "Database error", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": statuses})
}

// CreateStatus godoc
// @Summary Create a new ping status
// @Description Create a new ping status in the system
// @Tags statuses
// @Accept json
// @Produce json
// @Param status body model.PingStatus true "Ping Status"
// @Success 201 {object} model.PingStatus "Created ping status"
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Failure 500 {object} map[string]interface{} "Failed to create status"
// @Router /api/v1/status [post]
func (h *Handler) CreateStatus(c *gin.Context) {
	var status model.PingStatus
	if err := c.ShouldBindJSON(&status); err != nil {
		sendError(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := h.pingService.Create(&status); err != nil {
		sendError(c, http.StatusInternalServerError, "Failed to create status", err)
		return
	}

	c.Header("Location", fmt.Sprintf("/api/v1/status/%d", status.ID))
	c.JSON(http.StatusCreated, status)
}

// GetStatus godoc
// @Summary Get a ping status by ID
// @Description Retrieve a specific ping status by its ID
// @Tags statuses
// @Accept json
// @Produce json
// @Param id path int true "Ping Status ID"
// @Success 200 {object} model.PingStatus "Ping status"
// @Failure 400 {object} map[string]interface{} "Invalid ID format"
// @Failure 500 {object} map[string]interface{} "Database error"
// @Router /api/v1/status/{id} [get]
func (h *Handler) GetStatus(c *gin.Context) {
	idStr := c.Param("id")         // Получаем id как строку
	id, err := strconv.Atoi(idStr) // Преобразуем строку в int
	if err != nil {
		sendError(c, http.StatusBadRequest, "Invalid ID format", err)
		return
	}

	status, err := h.pingService.GetByID(id)
	if err != nil {
		sendError(c, http.StatusInternalServerError, "Database error", err)
		return
	}

	c.JSON(http.StatusOK, status)
}

// UpdateStatus godoc
// @Summary Update a ping status
// @Description Update a ping status by its ID
// @Tags statuses
// @Accept json
// @Produce json
// @Param id path int true "Ping Status ID"
// @Param status body model.PingStatus true "Ping Status"
// @Success 200 {object} model.PingStatus "Updated ping status"
// @Failure 400 {object} map[string]interface{} "Invalid request payload or ID format"
// @Failure 500 {object} map[string]interface{} "Failed to update status"
// @Router /api/v1/status/{id} [put]
func (h *Handler) UpdateStatus(c *gin.Context) {
	idStr := c.Param("id")         // Получаем id как строку
	id, err := strconv.Atoi(idStr) // Преобразуем строку в int
	if err != nil {
		sendError(c, http.StatusBadRequest, "Invalid ID format", err)
		return
	}

	var status model.PingStatus
	if err := c.ShouldBindJSON(&status); err != nil {
		sendError(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	status.ID = id
	if err := h.pingService.Update(&status); err != nil {
		sendError(c, http.StatusInternalServerError, "Failed to update status", err)
		return
	}

	c.JSON(http.StatusOK, status)
}

// DeleteStatus godoc
// @Summary Delete a ping status by ID
// @Description Delete a ping status by its ID
// @Tags statuses
// @Accept json
// @Produce json
// @Param id path int true "Ping Status ID"
// @Success 204 {object} nil "No content"
// @Failure 400 {object} map[string]interface{} "Invalid ID format"
// @Failure 500 {object} map[string]interface{} "Failed to delete status"
// @Router /api/v1/status/{id} [delete]
func (h *Handler) DeleteStatus(c *gin.Context) {
	idStr := c.Param("id")         // Получаем id как строку
	id, err := strconv.Atoi(idStr) // Преобразуем строку в int
	if err != nil {
		sendError(c, http.StatusBadRequest, "Invalid ID format", err)
		return
	}

	if err := h.pingService.Delete(id); err != nil {
		sendError(c, http.StatusInternalServerError, "Failed to delete status", err)
		return
	}

	c.Status(http.StatusNoContent)
}

// sendError is a helper function to send error responses.
func sendError(c *gin.Context, code int, message string, details interface{}) {
	c.JSON(code, gin.H{
		"message": message,
		"code":    code,
		"details": details,
	})
}
