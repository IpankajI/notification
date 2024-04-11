package main

type NotificationRepo interface {
	Create(*NotificationModel) error
	List(status SentStatus) NotificationModels
	Update(ID int, status SentStatus) error
}

// in memory impl
type inMNotificationRepo struct {
	notifications map[int]*NotificationModel
}

func NewInMNotificationRepo() NotificationRepo {
	return &inMNotificationRepo{
		notifications: make(map[int]*NotificationModel),
	}
}

func (in *inMNotificationRepo) Create(notification *NotificationModel) error {

	in.notifications[notification.ID] = notification

	return nil
}
func (in *inMNotificationRepo) List(status SentStatus) NotificationModels {
	notifications := NotificationModels{}
	for _, notification := range in.notifications {
		if status != notification.Status {
			continue
		}
		notifications = append(notifications, notification)
	}
	return notifications
}

func (in *inMNotificationRepo) Update(ID int, status SentStatus) error {
	in.notifications[ID].Status = status
	return nil
}
