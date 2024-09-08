package handlers

import (
	pb "auth-service/genprotos"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePoll godoc
// @Summary Create a new poll
// @Description Create a new poll with poll number, title, and options
// @Tags polls
// @Accept json
// @Produce json
// @Param poll body pb.PollCreateReq true "Poll creation request"
// @Success 201 {object} string "Successfully created!"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /poll [post]
// @Security BearerAuth
func (h *HTTPHandler) CreatePoll(c *gin.Context) {
	var req pb.PollCreateReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid request payload: "+err.Error())
		return
	}

	for _, feedback := range req.Feedbacks {
		if feedback.From == nil {
			defaultFrom := int32(0)
			feedback.From = &defaultFrom
		}
		if feedback.To == nil {
			defaultTo := int32(0)
			feedback.To = &defaultTo
		}
		if feedback.Text == nil {
			defaultText := ""
			feedback.Text = &defaultText
		}
	}

	for _, option := range req.Options {
		if option.Variant == nil {
			defaultVariant := ""
			option.Variant = &defaultVariant
		}
		if option.Ball == nil {
			defaultBall := int32(0)
			option.Ball = &defaultBall
		}
	}
	_, err := h.Poll.Create(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Server error: "+err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully created!"})
}

// UpdatePoll godoc
// @Summary Update an existing poll
// @Description Update an existing poll with poll number, title, and options
// @Tags polls
// @Accept json
// @Produce json
// @Param id path string true "Poll ID"
// @Param poll body pb.PollUpdateReqSwag true "Poll update request"
// @Success 200 {object} string "Successfully updated"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Poll not found"
// @Failure 500 {object} string "Server error"
// @Router /poll/{id} [put]
// @Security BearerAuth
func (h *HTTPHandler) UpdatePoll(c *gin.Context) {
	id := c.Param("id")
	req := pb.PollUpdateReq{Id: &id}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request payload": err.Error()})
		return
	}

	_, err := h.Poll.Update(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully updated"})
}

// DeletePoll godoc
// @Summary Delete an existing poll
// @Description Delete a poll by its ID
// @Tags polls
// @Accept json
// @Produce json
// @Param id path string true "Poll ID"
// @Success 200 {object} string "Successfully deleted!"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Poll not found"
// @Failure 500 {object} string "Server error"
// @Router /poll/{id} [delete]
// @Security BearerAuth
func (h *HTTPHandler) DeletePoll(c *gin.Context) {
	id := c.Param("id")
	req := pb.ByID{Id: id}

	_, err := h.Poll.Delete(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted!"})
}

// GetPollByID godoc
// @Summary Get a poll by ID
// @Description Get a poll by its ID
// @Tags polls
// @Accept json
// @Produce json
// @Param id path string true "Poll ID"
// @Success 200 {object} pb.PollGetByIDRes
// @Failure 404 {object} string "Poll not found"
// @Failure 500 {object} string "Server error"
// @Router /poll/{id} [get]
// @Security BearerAuth
func (h *HTTPHandler) GetPollByID(c *gin.Context) {
	id := c.Param("id")
	req := pb.ByID{Id: id}

	res, err := h.Poll.GetByID(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetAllPolls godoc
// @Summary Get all polls
// @Description Get all polls with pagination
// @Tags polls
// @Accept json
// @Produce json
// @Success 200 {object} pb.PollGetAllRes
// @Failure 500 {object} string "Server error"
// @Router /polls [get]
// @Security BearerAuth
func (h *HTTPHandler) GetAllPolls(c *gin.Context) {
	var req pb.PollGetAllReq

	res, err := h.Poll.GetAll(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
