package handlers

import (
	pb "auth-service/genprotos"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Send answers godoc
// @Summary Handles results
// @Description incoming results of user
// @Tags answer
// @Accept json
// @Produce json
// @Param poll body pb.IncomingResult true "Poll creation request"
// @Success 201 {object} string "Successfully posted"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /send-answers [POST]
// @Security BearerAuth
func (h *HTTPHandler) SendAnswers(c *gin.Context) {
	var req pb.IncomingResult
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	poll := pb.CreateResultReq{
		UserId: req.UserId,
		PollId: req.PollId,
	}

	id, err := h.Result.CreateResult(c, &poll)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't create result", "details": err.Error()})
		return
	}

	for _, i := range req.Answers {
		_, err := h.Result.SavePollAnswer(c, &pb.SavePollAnswerReq{ResultId: id.ResultId, QuestionId: i.QuestionId, Answer: i.AnswerPoint})
		if err != nil {
			c.JSON(400, gin.H{"error": "Couldn't save answer", "details": err.Error()})
			return
		}
	}

	feedResult, err := h.Result.GetPollResults(c, &pb.ByIDs{ResultId: id.ResultId, PollId: poll.PollId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get results", "details": err.Error()})
		return
	}
	ressss := *feedResult.Feed
	res := ressss[6:]
	fmt.Println([]byte(res))

	c.JSON(201, gin.H{"message": "Successfully posted", "id": id.ResultId, "feedback": res})
}
