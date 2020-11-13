package main

import (
  "encoding/json"
  "fmt"
  "bytes"
  "net/http"
  "io/ioutil"
  "github.com/gin-gonic/gin"
)

type answerOut struct {
	Question_id string `json:"id"`
	Status string `json:"status"`
}

func PostAnswer(c *gin.Context)  {
	token := c.Query("token")
 	var url string = "https://api.mercadolibre.com/answers?access_token=" + token
	resp, err :=http.Post(url, "application/json", c.Request.Body)
 	if err !=nil{
 		c.JSON(resp.StatusCode,err.Error())
		return
	}
	data,err := ioutil.ReadAll(resp.Body)
	if err !=nil{
		c.JSON(500,err.Error())
		return
	}
	var res answerOut
	err = json.Unmarshal(data, &res)
	if err != nil{
		c.JSON(500,err.Error())
	}
	c.JSON(200,res)
}
type Respuesta struct{
	Access_token string
	Token_type string
	Expires_in int
	Scope string
	User_id int
	Refresh_token string

}

func GetToken(c *gin.Context){
	code := c.Query("code")
	client := http.Client{}
 	requestBody, _ := json.Marshal(map[string]string{
		"grant_type": "authorization_code",
		"client_id" : "5291933962243912",
		"client_secret": "tnc3qX88LDPrWWXSMN3cL7OYd4L0y8Ta",
		"code" : code,
		"redirect_uri": "http://localhost:8080/auth",
	})
	request, _ := http.NewRequest("POST", "https://api.mercadolibre.com/oauth/token",bytes.NewBuffer(requestBody))
	request.Header.Set("accept", "application/json")
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	resp, _ := client.Do( request )
	defer resp.Body.Close()
	var res Respuesta
	data, _ := ioutil.ReadAll(resp.Body)
	er := json.Unmarshal(data, &res)
	if er != nil {
		println("There has been an error with the Unmarshal of the json")
		fmt.Println(er)
	}
	c.JSON(200, res)
}

type ItemIds struct {
	Results[] string `results`
}

type Questions[] struct {
	Date_created string `json:"date_created"`
	Item_id       string `json:"item_id"`
	Status       string `json:"status"`
	Text         string `json:"text"`
	Id           int64  `json:"id"`
	Answer       string `json:"answer"`
}

type Question struct {
	Questn Questions `json:"questions""`
}


type Items struct{
	Body struct {
		Id    string
		Title string
		Price float32
		Pictures[] map[string]string
		Available_quantity int
		Sold_quantity int
	}

}

type Item struct{
	Id    string
	Title string
	Price float32
	Quantity int
	SoldQuantity int
	Picture string
	Question Questions
}
type Sales struct {
	Results []struct {
		Payments []struct {
			Reason            string      `json:"reason"`
			StatusCode        interface{} `json:"status_code"`
			TotalPaidAmount   float64     `json:"total_paid_amount"`
			OperationType     string      `json:"operation_type"`
			TransactionAmount float64     `json:"transaction_amount"`
			DateApproved      string      `json:"date_approved"`
			Collector         struct {
				ID int `json:"id"`
			} `json:"collector"`
			CouponID             interface{} `json:"coupon_id"`
			Installments         int         `json:"installments"`
			AuthorizationCode    string      `json:"authorization_code"`
			TaxesAmount          int         `json:"taxes_amount"`
			ID                   int64       `json:"id"`
			DateLastModified     string      `json:"date_last_modified"`
			CouponAmount         int         `json:"coupon_amount"`
			AvailableActions     []string    `json:"available_actions"`
			ShippingCost         float64         `json:"shipping_cost"`
			InstallmentAmount    float64     `json:"installment_amount"`
			DateCreated          string      `json:"date_created"`
			ActivationURI        interface{} `json:"activation_uri"`
			OverpaidAmount       int         `json:"overpaid_amount"`
			CardID               int         `json:"card_id"`
			StatusDetail         string      `json:"status_detail"`
			IssuerID             string      `json:"issuer_id"`
			PaymentMethodID      string      `json:"payment_method_id"`
			PaymentType          string      `json:"payment_type"`
			DeferredPeriod       interface{} `json:"deferred_period"`
			AtmTransferReference struct {
				TransactionID string      `json:"transaction_id"`
				CompanyID     interface{} `json:"company_id"`
			} `json:"atm_transfer_reference"`
			SiteID             string      `json:"site_id"`
			PayerID            int         `json:"payer_id"`
			MarketplaceFee     float64     `json:"marketplace_fee"`
			OrderID            int         `json:"order_id"`
			CurrencyID         string      `json:"currency_id"`
			Status             string      `json:"status"`
			TransactionOrderID interface{} `json:"transaction_order_id"`
		} `json:"payments"`
	}
}

//Este struct será el enviado como respuesta
type ItemCarrier struct{
	Id    string
	Title string
	Price float32
	Quantity int
	SoldQuantity int
	Picture string
	Question Questions

}

type SalesCarrier struct{
	Id int64
	Title string
	Date string
	Price float64
	PriceTotal float64
}

type ResponseCarrier struct{
	Items []ItemCarrier
	Sales []SalesCarrier
}

func GetItems(c *gin.Context) {
	token := c.Query("token")
	userid := c.Query("userid")
	var url string = "https://api.mercadolibre.com/users/" + userid + "/items/search?access_token=" + token
	var res ItemIds
	getAndMarshall(url, &res, c)
	ch1 := make(chan Item)
	ch2 := make(chan Question)
	var response[] Item
	if len(res.Results) == 0 {
		c.JSON(400, struct {
			Error string}{
			Error: "No items being sold by this user",
		})
		return
	}
	for i := 0; i < len(res.Results); i++ {
		go itemCollector(token, res.Results[i], ch1, c)
		go questionCollector(token, res.Results[i], ch2, c)
		respContainer := <-ch1
		respContainer.Question = (<-ch2).Questn
		response = append(response, respContainer)
	}
	c.JSON(200, response)
}

func itemCollector(token, itemid string, ch1 chan Item, c *gin.Context){
	var url string = "https://api.mercadolibre.com/items?ids="+ itemid +"&attributes=id,price,available_quantity,title,pictures,sold_quantity&access_token=" + token
	var res[] Items
	getAndMarshall(url, &res, c)
	var resp Item
	resp.Quantity = res[0].Body.Available_quantity
	resp.SoldQuantity = res[0].Body.Sold_quantity
	resp.Id = res[0].Body.Id
	resp.Picture = res[0].Body.Pictures[0]["url"]
	resp.Price = res[0].Body.Price
	resp.Title = res[0].Body.Title
	ch1 <- resp
}

func questionCollector(token, itemid string, ch2 chan Question, c *gin.Context)  {
	var url string = "https://api.mercadolibre.com/questions/search?item="+ itemid +"&access_token=" + token
	var res Question
	getAndMarshall(url, &res, c)
	var unansweredQ Question
	if len(res.Questn) == 0 {
		ch2 <- Question{}
	}
	for i := len(res.Questn)-1; i >= 0 ; i-- {
		if res.Questn[i].Status == "UNANSWERED" {
			unansweredQ.Questn = append(unansweredQ.Questn, res.Questn[i])
		}
	}
	ch2 <- unansweredQ
}
func salesCollector(token, userid string, ch3 chan []SalesCarrier, c *gin.Context)  {
	var url string = "https://api.mercadolibre.com/orders/search?seller=" + userid + "&order.status=paid&access_token=" + token
	var res Sales
	var resp []SalesCarrier
	getAndMarshall(url, &res, c)
	if len(res.Results[0].Payments) == 0 {
		ch3 <- []SalesCarrier{}
	}
	for i := 0; i < len(res.Results[0].Payments); i++ {
		var sales SalesCarrier
		sales.Id = res.Results[0].Payments[i].ID
		sales.Title = res.Results[0].Payments[i].Reason
		sales.Date = res.Results[0].Payments[i].DateApproved
		sales.PriceTotal = res.Results[0].Payments[i].TotalPaidAmount
		sales.Price = res.Results[0].Payments[i].TransactionAmount
		resp = append(resp, sales)
	}
	ch3 <- resp

}
func getAndMarshall(url string, res interface{}, c *gin.Context)  {
	req, erro := http.Get(url)
	if req.StatusCode != 200 || erro != nil{
		c.JSON(req.StatusCode, erro.Error())
		return
	}
	defer req.Body.Close()
	data, erro := ioutil.ReadAll(req.Body)
	if erro != nil {
		c.JSON(req.StatusCode, erro.Error())
		return
	}
	erro = json.Unmarshal(data, &res)
	if erro != nil {
		c.JSON(req.StatusCode, erro.Error())
		return
	}
}
func main (){
  router := gin.Default()
	router.GET("/auth",GetToken)
	router.GET("/items/all",GetItems)
	//router.POST("/items/publish", PostItem)
	router.POST("/items/questions/answer",PostAnswer)
	router.Run()
}
