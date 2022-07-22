package main

import (
  "fmt"
	"net/http"
	"tinkoff-api/help"
	"tinkoff-api/model"
  //"tinkoff-api/send"
	"github.com/gin-gonic/gin"


	//"reflect"
)



type N map[string]interface{}


func convert_and_send(notification map[string]interface{}){
  data := model.Data{}
	for i := range notification{
		switch i{
		case "OrderId":
			data.OrderId = notification[i]
		case "PaymentId":
			data.PaymentId = notification[i]
		case "RebillId":
		 	data.RebillId = notification[i]
		case "CardId":
		  data.CardId	= notification[i]

		}
	}

	fmt.Println(data.OrderId)
	fmt.Println(data.PaymentId)
	fmt.Println(data.RebillId)
	fmt.Println(data.CardId)
  send.Send_notification(data)

}

func signal(c *gin.Context) {
	var notification N

	err := c.BindJSON(&notification)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, help.BuildErrorResponse("Failed to bind json", err.Error(), help.EmptyObj{}))
		return
	}

	//data := help.BuildReponse(true, "OK", notification)
  fmt.Println(notification)
  convert_and_send(notification)

	c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte("OK"))


}

func main() {
	r := gin.Default()

	r.GET("/api", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte("OK"))
	})

	r.POST("/ok", signal)

	r.Run(":8080")


}
