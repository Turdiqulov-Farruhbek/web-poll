package handlers

import (
	"context"
	"fmt"
	"net/http"

	pb "auth-service/genprotos"

	"github.com/gin-gonic/gin"
)

// CreateQuestion godoc
// @Summary Create a new question
// @Description Create a new question with the given details
// @Tags question
// @Accept json
// @Produce json
// @Param question body pb.QuestionCreateReq true "Question creation request"
// @Success 200 {object} string "Question created successfully"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /question [post]
// @Security BearerAuth
func (h *HTTPHandler) CreateQuestion(c *gin.Context) {
	var req pb.QuestionCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Question.Create(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Question created successfully"})
}

// UpdateQuestion godoc
// @Summary Update an existing question
// @Description Update an existing question with the given details
// @Tags question
// @Accept json
// @Produce json
// @Param question body pb.QuestionUpdateReq true "Question update request"
// @Success 200 {object} string "Question updated successfully"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /question [put]
// @Security BearerAuth
func (h *HTTPHandler) UpdateQuestion(c *gin.Context) {
	var req pb.QuestionUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Question.Update(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Question updated successfully"})
}

// DeleteQuestion godoc
// @Summary Delete a question
// @Description Delete a question by its ID
// @Tags question
// @Accept json
// @Produce json
// @Param id path string true "Question ID"
// @Success 200 {object} string "Question deleted successfully"
// @Failure 500 {object} string "Server error"
// @Router /question/{id} [delete]
// @Security BearerAuth
func (h *HTTPHandler) DeleteQuestion(c *gin.Context) {
	id := c.Param("id")
	req := pb.ByID{Id: id}
	_, err := h.Question.Delete(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Question deleted successfully"})
}

// GetQuestionByID godoc
// @Summary Get a question by ID
// @Description Get a question by its ID
// @Tags question
// @Accept json
// @Produce json
// @Param id path string true "Question ID"
// @Success 200 {object} pb.QuestionGetByIDRes
// @Failure 500 {object} string "Server error"
// @Router /question/{id} [get]
// @Security BearerAuth
func (h *HTTPHandler) GetQuestionByID(c *gin.Context) {
	id := c.Param("id")
	req := pb.ByID{Id: id}
	res, err := h.Question.GetByID(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetAllQuestions godoc
// @Summary Get all questions
// @Description Get all questions with pagination
// @Tags question
// @Accept json
// @Produce json
// @Param poll_id path string true "Poll ID"
// @Success 200 {object} pb.QuestionGetAllRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /questions/{poll_id} [get]
// @Security BearerAuth
func (h *HTTPHandler) GetAllQuestions(c *gin.Context) {
	poll_id := c.Param("poll_id")
	fmt.Println("poll_id: ", poll_id)
	req := pb.QuestionGetAllReq{PollId: &poll_id}

	res, err := h.Question.GetAll(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
