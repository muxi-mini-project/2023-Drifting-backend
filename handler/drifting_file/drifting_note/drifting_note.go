package drifting_note

import (
	"Drifting/controller"
	"Drifting/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateDriftingNote(c *gin.Context) {
	StudentID := c.Param("student_id")
	id, err := strconv.Atoi(StudentID)
	var NewDriftngNote model.DriftingNote
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	err = c.BindJSON(&NewDriftngNote)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	code, message := controller.CreateDriftingNote(id, NewDriftngNote)
	c.JSON(code, gin.H{"message": message})
}
