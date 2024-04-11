package main

import "fmt"

func main() {

	notificationService := NewNotificationService()

	ntf, err := notificationService.CreateNotification(11.0, 111.0, 432, 2334)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ntf.ID)

	notificationService.SendNotification([]string{"pankaj@gmail.com"})

	ntfs := notificationService.ListNotification(statusSent)

	for _, nft := range ntfs {
		fmt.Println("sent notification: ", nft.ID)
	}

	ntfs = notificationService.ListNotification(statusFailed)

	for _, nft := range ntfs {
		fmt.Println("failed notification: ", nft.ID)
	}

}
