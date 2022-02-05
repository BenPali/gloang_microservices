package service

import (
	"context"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"bytes"

	"github.com/k4lii/golang_microservices/db/connection"
	"github.com/k4lii/golang_microservices/db/models"
)
var transactionURL = os.Getenv("TRANSACTIONS_URL")
// MessagesService describes the service.
type MessagesService interface {
	AddMessage(ctx context.Context, jwt string, id uint, transactionId uint, messages string) (rs string, err error)
	GetMessage(ctx context.Context, jwt string, id uint, transactionId uint) (rs string, messages interface{}, err error)
}

type transaction struct {
	Rs      string `json:"rs"`
	Err     error  `json:"err"`
	Offers []struct {
		AdID  uint `json:"ad_id"`
		Price  float64  `json:"price"`
		Status string  `json:"status"`
		PosterID uint `json:"poster_id"`
		BuyerID uint `json:"buyer_id"`
	}
}

type basicMessagesService struct{}

func (b *basicMessagesService) AddMessage(ctx context.Context, jwt string, id uint, transactionId uint, messages string) (rs string, err error) {
	// TODO implement the business logic of AddMessage
	db := connection.DB
	var isInTransaction = false
	var transaction transaction
	body, _ := json.Marshal(map[string]string{
		"jwt": jwt,
	})
	bodyp := bytes.NewBuffer(body)
	resp, err := http.Post("http://"+transactionURL+"/list-own-offers", "application/json", bodyp)
	if err != nil {
		return "Ko", err
	}
	defer resp.Body.Close()
	//Retransaction the response body
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "KO", err
	}
	err = json.Unmarshal(body, &transaction)

	if err != nil {
		return "KO", err
	}
	for i := 0; i < len(transaction.Offers); i++ {
		if id == transaction.Offers[i].BuyerID || id == transaction.Offers[i].PosterID {
			isInTransaction = true
			break
		}
	}

	if  isInTransaction == true {
		newMessages := models.Message{
			SenderID:      id,
			Message:       messages,
			TransactionID: transactionId,
		}
		res := db.Table("messages").Create(&newMessages)

		if res.Error != nil {
			return "KO", res.Error
		}
		return "Ok", nil
	} else {
		return "Ko", fmt.Errorf("You have no access to this transaction")
	}
}
func (b *basicMessagesService) GetMessage(ctx context.Context, jwt string, id uint, transactionId uint) (rs string, messages interface{}, err error) {
	// TODO implement the business logic of GetMessage
	db := connection.DB
	var isInTransaction = false
	var transaction transaction
	body, _ := json.Marshal(map[string]string{
		"jwt": jwt,
	})
	bodyp := bytes.NewBuffer(body)
	resp, err := http.Post("http://"+transactionURL+"/list-own-offers", "application/json", bodyp)
	if err != nil {
		return "KO", nil ,err
	}
	defer resp.Body.Close()
	//Retransaction the response body
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "KO", nil ,err
	}
	err = json.Unmarshal(body, &transaction)

	if err != nil {
		return "KO", nil ,err
	}
	fmt.Printf(string(body));
	for i := 0; i < len(transaction.Offers); i++ {
		if id == transaction.Offers[i].BuyerID || id == transaction.Offers[i].PosterID {
			isInTransaction = true
			break
		}
	}

	if  isInTransaction == true {
		var messages []models.Message
		res := db.Table("messages").Select("id, sender_id", "message", "created_at").Where("transaction_id = ?", transactionId).Scan(&messages)
		if res.Error != nil {
			return "KO", nil, res.Error
		}
		ret := make([]map[string]interface{}, 0, len(messages))
		for i := 0; i < len(messages); i++ {
			tmp := map[string]interface{}{
				"id" : messages[i].ID,
				"sender_id":  messages[i].SenderID,
				"message":   messages[i].Message,
				"created_at": messages[i].CreatedAt,
			}
			ret = append(ret, tmp)
		}
		return "Ok", ret, nil
	} else {
		return "Ko", nil, fmt.Errorf("You have no access to this transaction")
	}
}

// NewBasicMessagesService returns a naive, stateless implementation of MessagesService.
func NewBasicMessagesService() MessagesService {
	return &basicMessagesService{}
}

// New returns a MessagesService with all of the expected middleware wired in.
func New(middleware []Middleware) MessagesService {
	var svc MessagesService = NewBasicMessagesService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
