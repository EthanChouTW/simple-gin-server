package main

import (
	"fmt"
	"net/http"

	"github.com/EthanChouTW/studygroup/src/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

type LoginCommand struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateCommand struct {
	Topic         string `json:"topic"`
	Class         string `json:"class"`
	ExecutiveTime int    `json:"executiveTime"`
}

// Payload type is for passing JSON format
type Payload struct {
	Username  string `json:"username,omitempty"`
	IconEmoji string `json:"icon_emoji,omitempty"`
	Channel   string `json:"channel,omitempty"`
	Text      string `json:"text,omitempty"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api")
	v1 := api.Group("/v1")
	v1.Use(middleware.DummyMiddleware())
	{
		v1.POST("/login", func(c *gin.Context) {
			sendSlack("enenenenenenen")
			var loginCmd LoginCommand
			c.BindJSON(&loginCmd)
			fmt.Println(loginCmd)
			c.JSON(http.StatusOK, gin.H{"user": loginCmd})
		})

		v1.POST("/consume", func(c *gin.Context) {
			// 看哪一個topics來決定要消耗哪一個

			// topics-> bulkupload
			// if bulkupload {
			// 	consume bulkupload channel
			// 	publish to immediately channel
			// }

			// topics -> immediately
			// if immediately
			// consume bulkupload channel -> check class
			// if class == bulkUpload && timestamp 到期 ->
			// -> post cooresponding api to alleypin server
		})

		v1.POST("/publish", func(c *gin.Context) {
			// 直接publish 到指定topics
			// 建立空白設定檔。
			var createCommand CreateCommand
			c.BindJSON(&createCommand)
			c.JSON(http.StatusOK, gin.H{"status": "success", "data": createCommand})
		})

	}
	router.Run()
}

func redirectPolicyFunc(req gorequest.Request, via []gorequest.Request) error {
	return fmt.Errorf("Incorrect token (redirection)")
}

func send(url string, proxy string, content interface{}) []error {
	fmt.Println(content)
	request := gorequest.New().Proxy(proxy)
	resp, _, err := request.
		Post(url).
		RedirectPolicy(redirectPolicyFunc).
		Send(content).
		End()

	if err != nil {
		fmt.Println("Error: ", err)
	}
	if resp.StatusCode >= 400 {
		return []error{fmt.Errorf("Error sending msg. Status: %v", resp.Status)}
	}
	return nil
}

func sendSlack(message string) error {
	webhookURL := ""
	payload := Payload{
		Text:      message,
		Username:  "enen",
		IconEmoji: ":smile",
		Channel:   "#testing",
	}
	fmt.Printf(" Platform, message content")
	err := send(webhookURL, "", payload)
	if err != nil {
		return err[0]
	}
	return nil
}
