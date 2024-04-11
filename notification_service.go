package main

import (
	"fmt"
	"math/rand"
)

type NotificationService interface {
	CreateNotification(btcPrice float64, marketCap float32, tradeVolume float64, intraDayHighPrice float64) (*Notification, error)
	SendNotification(emails []string)
	ListNotification(status SentStatus) Notifications
}

type SentStatus int

const (
	statusSent        = SentStatus(1)
	statusOutstanding = SentStatus(2)
	statusFailed      = SentStatus(3)
)

type Notification struct {
	ID                int
	Status            SentStatus
	btcPrice          float64
	marketCap         float64
	tradeVolume       float64
	intraDayHighPrice float64
}

type Notifications []*Notification

type notificationService struct {
	repo    NotificationRepo
	emailer Emailer
}

func NewNotificationService() NotificationService {
	return &notificationService{
		repo:    NewInMNotificationRepo(),
		emailer: NewEmailService(),
	}
}

func (ns *notificationService) CreateNotification(btcPrice float64, marketCap float32, tradeVolume float64, intraDayHighPrice float64) (*Notification, error) {
	nm := &NotificationModel{
		ID:                rand.Int() % 100000,
		btcPrice:          btcPrice,
		marketCap:         float64(marketCap),
		tradeVolume:       tradeVolume,
		intraDayHighPrice: intraDayHighPrice,
		Status:            statusOutstanding,
	}
	err := ns.repo.Create(nm)
	return toAPINotification(nm), err
}

func (ns *notificationService) SendNotification(emails []string) {
	notifications := ns.repo.List(statusOutstanding)
	for _, email := range emails {
		for _, ntf := range notifications {
			sub, body, err := FormateNotificationEmail(ntf)
			if err != nil {
				fmt.Println("error: ", err.Error())
				err = ns.repo.Update(ntf.ID, statusFailed)
				if err != nil {
					fmt.Println("error: ", err.Error())
				}
				continue
			}
			err = ns.emailer.SendEmail(email, sub, body)
			if err != nil {
				fmt.Println("error: ", err.Error())
				err = ns.repo.Update(ntf.ID, statusFailed)
				if err != nil {
					fmt.Println("error: ", err.Error())
				}
			} else {
				err = ns.repo.Update(ntf.ID, statusSent)
				if err != nil {
					fmt.Println("error: ", err.Error())
				}
			}
		}
	}
}

func (ns *notificationService) ListNotification(status SentStatus) Notifications {
	notifications := ns.repo.List(status)
	resp := Notifications{}
	for _, notf := range notifications {
		resp = append(resp, toAPINotification(notf))
	}
	return resp
}
