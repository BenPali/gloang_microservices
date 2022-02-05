package service

import (
	"context"
	"fmt"

	"github.com/k4lii/golang_microservices/db/connection"
	"github.com/k4lii/golang_microservices/db/models"
)

// AdsService describes the service.
type AdsService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	CreateAd(ctx context.Context, jwt string, ID uint, title, desc, picture string, price float64) (rs string, err error)
	UpdateAd(ctx context.Context, jwt string, ID uint, addID uint, title, desc, picture string, price float64) (rs string, err error)
	DeleteAd(ctx context.Context, jwt string, ID uint, addID uint) (rs string, err error)
	GetAd(ctx context.Context, addID uint) (string, interface{}, error)
	SearchAd(ctx context.Context, keyword string) (string, interface{}, error)
	GetUserAdsList(ctx context.Context, posterID uint) (s0 string, i1 interface{}, e2 error)
}

type basicAdsService struct{}

func (b *basicAdsService) CreateAd(ctx context.Context, jwt string, ID uint, title string, desc string, picture string, price float64) (rs string, err error) {
	db := connection.DB

	res := db.Table("ads").Create(&models.Ad{
		Title:       title,
		Description: desc,
		Picture:     picture,
		Price:       price,
		Available:   true,

		PosterID: ID,
	})
	if res.Error != nil {
		return "KO", err
	}

	return "ok", err
}
func (b *basicAdsService) UpdateAd(ctx context.Context, jwt string, ID uint, addID uint, title string, desc string, picture string, price float64) (rs string, err error) {
	db := connection.DB

	if title != "" {
		res := db.Table("ads").Where("id = ?", addID).Update("title", title)
		if res.Error != nil {
			return "Ko", res.Error
		}
	}

	if desc != "" {
		res := db.Table("ads").Where("id = ?", addID).Update("desc", desc)
		if res.Error != nil {
			return "Ko", res.Error
		}
	}

	if picture != "" {
		res := db.Table("ads").Where("id = ?", addID).Update("picture", picture)
		if res.Error != nil {
			return "Ko", res.Error
		}
	}
	if price != 0 {
		res := db.Table("ads").Where("id = ?", addID).Update("price", price)
		if res.Error != nil {
			return "Ko", res.Error
		}
	}

	return "Ok", err
}
func (b *basicAdsService) DeleteAd(ctx context.Context, jwt string, ID uint, addID uint) (rs string, err error) {
	db := connection.DB
	var ad models.Ad
	fmt.Println(addID)
	res := db.Table("ads").Where("id = ?", addID).Update("available", false)
	if res.Error != nil {
		return "KO", res.Error
	}
	res = db.Table("ads").Where("id = ?", addID).Delete(&ad)
	if res.Error != nil {
		return "KO", res.Error
	}
	return "Add deleted", res.Error
}

func (b *basicAdsService) GetAd(ctx context.Context, addID uint) (string, interface{}, error) {
	db := connection.DB
	var ad models.Ad
	res := db.Table("ads").Select("id", "poster_id", "title", "description", "price", "available", "picture").Where("id = ?", addID).Scan(&ad)
	if res.Error != nil {
		return "KO", map[string]string{}, res.Error
	}

	ret := map[string]interface{}{
		"id":          ad.ID,
		"poster_id":   ad.PosterID,
		"title":       ad.Title,
		"description": ad.Description,
		"price":       ad.Price,
		"available":   ad.Available,
		"picture":     ad.Picture,
	}
	return "ok", ret, nil
}

func (b *basicAdsService) SearchAd(ctx context.Context, keyword string) (string, interface{}, error) {
	db := connection.DB
	var ad []models.Ad
	res := db.Table("ads").Select("id", "poster_id", "title", "description", "price", "available", "picture").Where("title LIKE ?", "%"+keyword+"%").Scan(&ad)
	if res.Error != nil {
		return "KO", map[string]string{}, res.Error
	}

	ret := make([]map[string]interface{}, 0, len(ad))

	for i := 0; i < len(ad); i++ {
		tmp := map[string]interface{}{
			"id":          ad[i].ID,
			"title":       ad[i].Title,
			"description": ad[i].Description,
			"price":       ad[i].Price,
			"available":   ad[i].Available,
			"picture":     ad[i].Picture,
			"poster_id":   ad[i].PosterID,
		}

		ret = append(ret, tmp)
	}

	return "ok", ret, nil
}

// NewBasicAdsService returns a naive, stateless implementation of AdsService.
func NewBasicAdsService() AdsService {
	return &basicAdsService{}
}

// New returns a AdsService with all of the expected middleware wired in.
func New(middleware []Middleware) AdsService {
	var svc AdsService = NewBasicAdsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicAdsService) GetUserAdsList(ctx context.Context, posterID uint) (s0 string, i1 interface{}, e2 error) {
	// TODO implement the business logic of GetUserAdsList

	db := connection.DB
	var ad []models.Ad
	res := db.Table("ads").Select("title", "description", "price", "available", "picture", "poster_id").Where("poster_id = ?", posterID).Scan(&ad)
	if res.Error != nil {
		return "KO", map[string]string{}, res.Error
	}

	ret := make([]map[string]interface{}, 0, len(ad))

	for i := 0; i < len(ad); i++ {
		tmp := map[string]interface{}{
			"id":          ad[i].ID,
			"title":       ad[i].Title,
			"description": ad[i].Description,
			"price":       ad[i].Price,
			"available":   ad[i].Available,
			"picture":     ad[i].Picture,
			"poster_id":   ad[i].PosterID,
		}

		ret = append(ret, tmp)
	}

	return "ok", ret, nil
}
