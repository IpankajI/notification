Ideally, we should use dependecy injection in main function to initialise service

```
notificationService := NewNotificationService(NewInMNotificationRepo(), NewEmailService())
```

correponding change in NewNotificationService function
```
func NewNotificationService(repo NotificationRepo, emailer Emailer) NotificationService {
	return &notificationService{
		repo:    repo,
		emailer: emailer,
	}
}
```
