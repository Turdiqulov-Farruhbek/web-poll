package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	pb "auth-service/genprotos"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// toAlphaString converts a column number to an Excel-style column name (e.g., 1 -> A, 27 -> AA).
func toAlphaString(n int) string {
	var columnName string
	for n > 0 {
		remainder := (n - 1) % 26
		columnName = string('A'+remainder) + columnName
		n = (n - 1) / 26
	}
	return columnName
}

// GetResults godoc
// @Summary Get results in excel file
// @Description Get overall results writing in excel file
// @Tags results
// @Accept json
// @Produce json
// @Success 200 {file} application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
// @Failure 404 {object} string "Poll not found"
// @Failure 500 {object} string "Server error"
// @Router /results [GET]
// @Security BearerAuth
func (h *HTTPHandler) GetUserResultsInExcel(c *gin.Context) {
	results, err := h.Result.GetResultsInExcel(c, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	f := excelize.NewFile()

	pollResults := make(map[int][]*pb.ResultRes)
	for _, result := range results.Results {
		pollResults[int(result.PollNum)] = append(pollResults[int(result.PollNum)], result)
	}

	for pollNum, pollRes := range pollResults {
		sheetName := fmt.Sprintf("Poll %d", pollNum)
		index, err := f.NewSheet(sheetName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create sheet", "details": err.Error()})
			return
		}

		headers := []string{
			"Ismi", "Familiyasi", "Jinsi", "Email", "Telefon raqami",
			"Ish tajribasi", "Darajasi", "So'rovnoma raqami", "Jami natijalar",
		}

		maxQuestions := 0
		for _, result := range pollRes {
			if len(result.Answers) > maxQuestions {
				maxQuestions = len(result.Answers)
			}
		}

		for i := 1; i <= maxQuestions; i++ {
			headers = append(headers, fmt.Sprintf("%d - savol", i))
		}

		for i, header := range headers {
			cell := fmt.Sprintf("%s1", toAlphaString(i+1))
			f.SetCellValue(sheetName, cell, header)
		}

		rowNum := 2
		for _, result := range pollRes {
			row := []interface{}{
				result.User.Name,
				result.User.Surname,
				result.User.Gender,
				result.User.Email,
				result.User.PhoneNumber,
				result.User.WorkingExperience,
				result.User.LevelType,
				fmt.Sprintf("%d - so'rovnoma", result.PollNum),
			}

			totalPoints := 0
			for _, answer := range result.Answers {
				row = append(row, answer.AnswerPoint)
				totalPoints += int(answer.AnswerPoint)
			}

			row = append(row[:8], append([]interface{}{totalPoints}, row[8:]...)...)

			for i, cellValue := range row {
				cell := fmt.Sprintf("%s%d", toAlphaString(i+1), rowNum)
				f.SetCellValue(sheetName, cell, cellValue)
			}
			rowNum++
		}

		f.SetActiveSheet(index)
	}

	var buffer bytes.Buffer
	if err := f.Write(&buffer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write Excel file", "details": err.Error()})
		return
	}

	fileName := "results.xlsx"
	filePath := filepath.Join("", fileName)
	if err := os.WriteFile(filePath, buffer.Bytes(), 0644); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save Excel file", "details": err.Error()})
		return
	}

	// c.JSON(http.StatusOK, results)
	c.Header("Content-Disposition", "attachment; filename=results.xlsx")
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buffer.Bytes())
	os.Remove(filePath)
}
