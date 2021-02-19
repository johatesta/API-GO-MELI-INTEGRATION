package controller

import (
	"github.com/gin-gonic/gin"
)

func ItemsAll(c *gin.Context) {
	token := c.Query("token")
	userid := c.Query("userid")
	var url string = "https://api.mercadolibre.com/users/" + userid + "/items/search?access_token=" + token
	var res ItemIds
	getAndMarshall(url, &res, c)
	var response []ItemCarrier
	if len(res.Results) == 0 {
		c.JSON(400, struct {
			Error string
		}{
			Error: "No items being sold by this user",
		})
		return
	}
	ch1 := make(chan ItemCarrier)
	ch2 := make(chan Question)
	ch3 := make(chan []SalesCarrier)
	go salesCollector(token, userid, ch3, c)
	for i := 0; i < len(res.Results); i++ {
		go itemCollector(token, res.Results[i], ch1, c)
		go questionCollector(token, res.Results[i], ch2, c)
		respContainer := <-ch1
		respContainer.Question = (<-ch2).Questn
		response = append(response, respContainer)
	}
	var finalResp ResponseCarrier
	finalResp.Items = response
	finalResp.Sales = <-ch3
	c.JSON(200, finalResp)
}

func itemCollector(token, itemid string, ch1 chan ItemCarrier, c *gin.Context) {
	var url string = "https://api.mercadolibre.com/items?ids=" + itemid + "&attributes=id,price,available_quantity,title,pictures,sold_quantity&access_token=" + token
	var res []Items
	getAndMarshall(url, &res, c)
	var resp ItemCarrier
	resp.Quantity = res[0].Body.Available_quantity
	resp.SoldQuantity = res[0].Body.Sold_quantity
	resp.Id = res[0].Body.Id
	resp.Picture = res[0].Body.Pictures[0]["url"]
	resp.Price = res[0].Body.Price
	resp.Title = res[0].Body.Title
	ch1 <- resp
}

func questionCollector(token, itemid string, ch2 chan Question, c *gin.Context) {
	var url string = "https://api.mercadolibre.com/questions/search?item=" + itemid + "&access_token=" + token
	var res Question
	getAndMarshall(url, &res, c)
	var unansweredQ Question
	if len(res.Questn) == 0 {
		ch2 <- Question{}
	}
	for i := len(res.Questn) - 1; i >= 0; i-- {
		if res.Questn[i].Status == "UNANSWERED" {
			unansweredQ.Questn = append(unansweredQ.Questn, res.Questn[i])
		}
	}
	ch2 <- unansweredQ
}

func salesCollector(token, userid string, ch3 chan []SalesCarrier, c *gin.Context) {
	var url string = "https://api.mercadolibre.com/orders/search?seller=" + userid + "&order.status=paid&access_token=" + token
	var res Sales
	var resp []SalesCarrier
	getAndMarshall(url, &res, c)
	if len(res.Results) == 0 {
		ch3 <- []SalesCarrier{}
	}
	for i := 0; i < len(res.Results); i++ {
		for j := 0; j < len(res.Results[i].Payments); j++ {
			var sales SalesCarrier
			sales.Id = res.Results[i].Payments[j].ID
			sales.Title = res.Results[i].Payments[j].Reason
			sales.Date = res.Results[i].Payments[j].DateApproved
			sales.PriceTotal = res.Results[i].Payments[j].TotalPaidAmount
			sales.Price = res.Results[i].Payments[j].TransactionAmount
			resp = append(resp, sales)
		}

	}
	ch3 <- resp

}
