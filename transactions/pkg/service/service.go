package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/k4lii/golang_microservices/db/connection"
	"github.com/k4lii/golang_microservices/db/models"
)

var AdURL = os.Getenv("ADS_URL")
var accURL = os.Getenv("ACCOUNT_URL")

type AcceptTransaction struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// TransactionsService describes the service.
type TransactionsService interface {
	MakeOffer(ctx context.Context, jwt string, id uint, adID uint, price float64) (rs string, err error)
	ManageOffer(ctx context.Context, jwt string, id uint, transactionID uint, accepted bool) (rs string, err error)
	ListOwnOffers(ctx context.Context, jwt string, id uint) (rs string, offers interface{}, err error)
}

type basicTransactionsService struct{}

type acc struct {
	Rs      string `json:"rs"`
	Err     error  `json:"err"`
	Profile struct {
		Balance  float64 `json:"balance"`
		Mail     string  `json:"mail"`
		Username string  `json:"username"`
	}
}

type ad struct {
	S0 string `json:"s0"`
	I1 struct {
		Available   bool    `json:"available"`
		Description string  `json:"description"`
		Id          uint    `json:"id"`
		Picture     string  `json:"picture"`
		PosterID    uint    `json:"poster_id"`
		Price       float64 `json:"price"`
		Title       string  `json:"title"`
	}
	E2 error `json:"e2"`
}

func (b *basicTransactionsService) ManageOffer(ctx context.Context, jwt string, id uint, transactionID uint, accepted bool) (rs string, err error) {
	db := connection.DB
	var transaction models.Transaction
	var AcceptTransaction AcceptTransaction
	fmt.Println(accepted)
	db.Table("transactions").Select("price", "buyer_id", "status", "ad_id").Where("id = ?", transactionID).Scan(&transaction)
	if accepted && transaction.Status == "In Progress" {
		body, _ := json.Marshal(map[string]interface{}{
			"jwt":    jwt,
			"amount": transaction.Price,
		})
		bodyp := bytes.NewBuffer(body)
		resp, err := http.Post("http://"+accURL+"/add-to-balance", "application/json", bodyp)
		if err != nil {
			return "Ko", err
		}
		defer resp.Body.Close()
		//Retransaction the response body
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return "KO", err
		}
		err = json.Unmarshal(body, &AcceptTransaction)

		if err != nil {
			return "KO", err
		}

		body, _ = json.Marshal(map[string]interface{}{
			"jwt":    os.Getenv("ADMIN_PASSWORD"),
			"id":     transaction.BuyerID,
			"amount": transaction.Price,
		})
		bodyp = bytes.NewBuffer(body)
		resp, err = http.Post("http://"+accURL+"/susbstract-to-balance", "application/json", bodyp)
		if err != nil {
			return "Ko", err
		}
		defer resp.Body.Close()
		//Retransaction the response body
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return "KO", err
		}
		err = json.Unmarshal(body, &AcceptTransaction)

		if err != nil {
			return "KO", err
		}
		res := db.Table("transactions").Where("id = ?", transactionID).Update("status", "accepted")
		if res.Error != nil {
			return "Ko", res.Error
		}

		body, _ = json.Marshal(map[string]interface{}{
			"jwt":    jwt,
			"add_id": transaction.AdID,
		})
		bodyp = bytes.NewBuffer(body)
		resp, err = http.Post("http://"+AdURL+"/delete-ad", "application/json", bodyp)
		if err != nil {
			return "Ko", err
		}
		defer resp.Body.Close()
		//Retransaction the response body
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return "KO", err
		}
		err = json.Unmarshal(body, &AcceptTransaction)
		if err != nil {
			return "KO", err
		}

		return "OK", nil
	}
	if transaction.Status != "In Progress" {
		return "Ko", fmt.Errorf("already %s", transaction.Status)
	}
	res := db.Table("transactions").Where("id = ?", transactionID).Update("status", "refused")
	if res.Error != nil {
		return "Ko", res.Error
	}
	return "OK", nil
}

func (b *basicTransactionsService) MakeOffer(ctx context.Context, jwt string, id uint, adID uint, price float64) (rs string, err error) {
	db := connection.DB

	var ad ad
	body, _ := json.Marshal(map[string]uint{
		"add_id": adID,
	})
	bodyp := bytes.NewBuffer(body)
	resp, err := http.Post("http://"+AdURL+"/get-ad", "application/json", bodyp)
	if err != nil {
		return "Ko", err
	}
	defer resp.Body.Close()
	//Read the response body
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "KO", err
	}
	err = json.Unmarshal(body, &ad)

	if err != nil {
		return "KO", err
	}

	var acc acc
	body2, _ := json.Marshal(map[string]string{
		"jwt": jwt,
	})
	bodyp2 := bytes.NewBuffer(body2)
	resp2, err2 := http.Post("http://"+accURL+"/get-self-account", "application/json", bodyp2)
	if err2 != nil {
		return "KO", err2
	}
	defer resp2.Body.Close()
	//Read the resp2onse body2
	body2, err = ioutil.ReadAll(resp2.Body)
	if err != nil {
		return "KO", err
	}
	err = json.Unmarshal(body2, &acc)
	if err != nil {
		return "KO", err
	}

	if acc.Profile.Balance >= price && ad.I1.Available {
		res := db.Table("transactions").Create(&models.Transaction{
			Status:   "In Progress",
			Price:    price,
			BuyerID:  id,
			AdID:     adID,
			PosterID: ad.I1.PosterID,
		})
		if res.Error != nil {
			return "KO", res.Error
		}
	} else if ad.I1.Available {
		return "This product is not available", nil
	} else {
		return "not enough money", nil
	}

	return "OK", nil
}

func (b *basicTransactionsService) ListOwnOffers(ctx context.Context, jwt string, id uint) (rs string, offers interface{}, err error) {
	db := connection.DB
	var transactions []models.Transaction
	res := db.Table("transactions").Select("ad_id", "price", "status", "poster_id", "buyer_id").Where("poster_id = ?", id).Or("buyer_id", id).Scan(&transactions)
	if res.Error != nil {
		return "KO", map[string]string{}, res.Error
	}

	ret := make([]map[string]interface{}, 0, len(transactions))
	for i := 0; i < len(transactions); i++ {
		tmp := map[string]interface{}{
			"ad_id":     transactions[i].AdID,
			"price":     transactions[i].Price,
			"status":    transactions[i].Status,
			"poster_id": transactions[i].PosterID,
			"buyer_id":  transactions[i].BuyerID,
		}
		ret = append(ret, tmp)
	}
	return "OK", ret, nil
}

// NewBasicTransactionsService returns a naive, stateless implementation of TransactionsService.
func NewBasicTransactionsService() TransactionsService {
	return &basicTransactionsService{}
}

// New returns a TransactionsService with all of the expected middleware wired in.
func New(middleware []Middleware) TransactionsService {
	var svc TransactionsService = NewBasicTransactionsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
