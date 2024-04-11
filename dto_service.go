package main

func toAPINotification(nm *NotificationModel) *Notification {
	return &Notification{
		btcPrice:          nm.btcPrice,
		marketCap:         nm.marketCap,
		tradeVolume:       nm.tradeVolume,
		intraDayHighPrice: nm.intraDayHighPrice,
		ID:                nm.ID,
		Status:            nm.Status,
	}
}
