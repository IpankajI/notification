package main

type NotificationModel struct {
	ID                int
	Status            SentStatus
	btcPrice          float64
	marketCap         float64
	tradeVolume       float64
	intraDayHighPrice float64
}

type NotificationModels []*NotificationModel
