package model

import "time"

/**
订单模型
 */
type Order struct {
	OrderId    string
	Count      int
	Time       time.Time
	ShowDate   string
	ShowTime   string
	UserId     int
	CinemaName string
	CityName   string
	FilmName   string
	OrderPrice int
	OfferName  string
	CinemaId   int
	Bid        string
	CityId     int
	OfferId    int
}
