package handlers

// CreatePoll godoc
// @Summary Create a new poll
// @Description Create a new poll with poll number, title, and options
// @Tags polls
// @Accept json
// @Produce json
// @Param poll body pb.PollCreateReq true "Poll creation request"
// @Success 201 {object} pb.Void
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /poll [post]
// @Security BearerAuth

// UpdatePoll godoc
// @Summary Update an existing poll
// @Description Update an existing poll with poll number, title, and options
// @Tags polls
// @Accept json
// @Produce json
// @Param poll body pb.PollUpdateReq true "Poll update request"
// @Success 200 {object} pb.Void
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Poll not found"
// @Failure 500 {object} string "Server error"
// @Router /poll [put]
// @Security BearerAuth

// DeletePoll godoc
// @Summary Delete an existing poll
// @Description Delete a poll by its ID
// @Tags polls
// @Accept json
// @Produce json
// @Param id path string true "Poll ID"
// @Success 200 {object} pb.Void
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Poll not found"
// @Failure 500 {object} string "Server error"
// @Router /poll/{id} [delete]
// @Security BearerAuth

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

// GetAllPolls godoc
// @Summary Get all polls
// @Description Get all polls with pagination
// @Tags polls
// @Accept json
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.PollGetAllRes
// @Failure 500 {object} string "Server error"
// @Router /polls [get]
// @Security BearerAuth
