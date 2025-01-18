package interfaces

import (
	"context"
	"e-wallet-notification/cmd/proto/notification"
	"e-wallet-notification/internal/models"
)

type IEmailExternal interface {
	SendEmail() error
}

type IEmailRepository interface {
	GetTemplate(ctx context.Context, templateName string) (models.NotificationTemplate, error)
	InsertNotificationHistory(ctx context.Context, notification *models.NotificationHistory) error
}

type IEmailService interface {
	SendEmail(ctx context.Context, req models.InternalNotificationRequest) error
}

type IEmailAPI interface {
	SendNotification(ctx context.Context, req *notification.SendNotificationRequest) (*notification.SendNotificationResponse, error)
}
