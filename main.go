package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type UserAgents []struct {
	AppName    string `json:"appName"`
	Connection struct {
		Downlink      int    `json:"downlink"`
		EffectiveType string `json:"effectiveType"`
		Rtt           int    `json:"rtt"`
	} `json:"connection"`
	Platform       string  `json:"platform"`
	PluginsLength  int     `json:"pluginsLength"`
	Vendor         string  `json:"vendor"`
	UserAgent      string  `json:"userAgent"`
	ViewportHeight int     `json:"viewportHeight"`
	ViewportWidth  int     `json:"viewportWidth"`
	DeviceCategory string  `json:"deviceCategory"`
	ScreenHeight   int     `json:"screenHeight"`
	ScreenWidth    int     `json:"screenWidth"`
	Weight         float64 `json:"weight"`
}

var userAgentsList []string

func loadUserAgents() []string {
	jsonFile, err := os.Open("user-agents.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
	}
	userAgents := new(UserAgents)
	json.Unmarshal(byteValue, userAgents)

	for _, item := range *userAgents {
		if item.AppName == "Netscape" {
			userAgentsList = append(userAgentsList, item.UserAgent)
		}
	}

	return userAgentsList
}

func init() {
	userAgentsList = loadUserAgents()
}

func getRandomUserAgent(c *gin.Context) {
	randIdx := rand.Intn(len(userAgentsList))
	c.String(http.StatusOK, userAgentsList[randIdx])
}

func getFavicon(c *gin.Context) {
	c.JSON(204, nil)
}

func getHealthz(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	router := gin.Default()
	router.GET("/favicon.ico", getFavicon)
	router.GET("/healthz", getHealthz)
	router.GET("/", getRandomUserAgent)
	router.Run(":8080")
}
