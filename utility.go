package main

import "encoding/json"

func FormateNotificationEmail(ntf *NotificationModel) (string, string, error) {
	body, err := json.Marshal(ntf)
	return "email subject", string(body), err
}
