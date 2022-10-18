package main

import (
	"math/rand"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

type Status struct {
	Water       int    `json:"water"`
	WaterStatus string `json:"waterstatus"`
	Wind        int    `json:"wind"`
	WindStatus  string `json:"windstatus"`
}

var status = []Status{
	{
		Water:       rand.Intn(100),
		WaterStatus: "",
		Wind:        rand.Intn(100),
		WindStatus:  "",
	},
}

func main() {
	var PORT = ":8080"
	StartServer().Run(PORT)
}

func StartServer() *gin.Engine {
	router := gin.Default()
	router.GET("/", getStatus2)
	return router
}

func getStatus2(c *gin.Context) {
	var statusData Status
	html, err := template.ParseFiles("example.html")

	// if err := c.ShouldBindJSON(&statusData); err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	if statusData.Water > 5 {
		statusData.WaterStatus = "aman"
	} else if statusData.Water >= 6 && statusData.Water <= 8 {
		statusData.WaterStatus = "siaga"
	} else if statusData.Water > 8 {
		statusData.WaterStatus = "bahaya"
	}

	if statusData.Wind > 6 {
		statusData.WaterStatus = "aman"
	} else if statusData.Water >= 7 && statusData.Water <= 15 {
		statusData.WaterStatus = "siaga"
	} else if statusData.Water > 15 {
		statusData.WaterStatus = "bahaya"
	}

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	html.Execute(c.Writer, status)

	c.JSON(http.StatusOK, statusData)

}
