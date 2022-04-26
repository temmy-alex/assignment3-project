package controllers

import (
	"assignment3/models"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateJsonFile(c *gin.Context) {
	data := models.StatusData{
		models.Status{
			Water: randInt(1, 100),
			Wind:  randInt(1, 100),
		},
	}

	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("json/status.json", file, 0644)
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func GetStatusData(c *gin.Context) {
	data, _ := ioutil.ReadFile("json/status.json")
	var result map[string]interface{}
	json.Unmarshal([]byte(data), &result)

	c.JSON(http.StatusOK, gin.H{"msg": "this worked", "data": result["status"]})
}
