package repository

import (
	"context"
	"e-wallet-notification/internal/models"
	"gorm.io/gorm"
)

type EmailRepo struct {
	DB *gorm.DB
}

func (r *EmailRepo) GetTemplate(ctx context.Context, templateName string) (models.NotificationTemplate, error) {
	var (
		resp models.NotificationTemplate
	)
	err := r.DB.Where("template_name = ?", templateName).Last(&resp).Error
	return resp, err
}

func (r *EmailRepo) InsertNotificationHistory(ctx context.Context, notification *models.NotificationHistory) error {
	return r.DB.Create(notification).Error
}
