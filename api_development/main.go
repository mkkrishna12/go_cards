package main

import (
	"api_developement/db"
	"api_developement/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()
	server.GET("/getEvents", Printl)
	server.POST("/create-event", createEvent)
	server.GET("/get-event/:eventId", getEvent)
	server.PUT("/update-event/:eventId", updateEvent)
	server.Run(":8080")

}

func updateEvent(gc *gin.Context) {

	_, err := strconv.ParseInt(gc.Param("eventId"), 10, 64)

	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect Value has been sent"})
		return
	}

	var eventDetail models.Event
	err = gc.ShouldBind(&eventDetail)

	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"messagse": "Valid date not found"})
		return
	}

	fmt.Println("updating Data", eventDetail)
	eventDetail, err = eventDetail.UpdateEvent()
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server errror"})
		return
	}
	gc.JSON(http.StatusFound, gin.H{"message": eventDetail})
}

func getEvent(gc *gin.Context) {
	eventId, err := strconv.ParseInt(gc.Param("eventId"), 10, 64)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect Value has been sent"})
		return
	}
	eventDetails, err := models.GetEventById(eventId)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"message": "Data coudnot be fetched", "error": err.Error()})
		return
	}
	gc.JSON(http.StatusFound, gin.H{"message": eventDetails})
}

func Printl(gs *gin.Context) {
	events, _ := models.GetAllEvents()
	gs.JSON(http.StatusOK, gin.H{"message": events})
}

func createEvent(gc *gin.Context) {

	var eventDetail models.Event
	err := gc.ShouldBind(&eventDetail)
	fmt.Println("Inserting Data", eventDetail)

	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"messagse": err})
		return
	}

	err = eventDetail.Save()

	if err == nil {
		gc.JSON(http.StatusOK, gin.H{"messagse": err})
		return
	}

	gc.JSON(http.StatusBadRequest, gin.H{"messagses": err})
}
