package services

import (
	"bytes"
	"context"
	"e-wallet-notification/external"
	"e-wallet-notification/internal/interfaces"
	"e-wallet-notification/internal/models"
	"github.com/pkg/errors"
	"html/template"
)

type EmailService struct {
	EmailRepo interfaces.IEmailRepository
}

func (s *EmailService) SendEmail(ctx context.Context, req models.InternalNotificationRequest) error {
	emailTemplate, err := s.EmailRepo.GetTemplate(ctx, req.TemplateName)
	if err != nil {
		return errors.Wrap(err, "failed to get email template")
	}

	parse, err := template.New("emailTemplate").Parse(emailTemplate.Body)
	if err != nil {
		return errors.Wrap(err, "failed to parse email template")
	}

	var tpl bytes.Buffer
	err = parse.Execute(&tpl, req.Placeholder)
	if err != nil {
		return errors.Wrap(err, "failed to execute email template")
	}

	email := external.Email{
		To:      req.Recipient,
		Subject: emailTemplate.Subject,
		Body:    tpl.String(),
	}

	err = email.SendEmail()
	if err != nil {
		notificationHistory := models.NotificationHistory{
			Recipient:    req.Recipient,
			TemplateID:   emailTemplate.ID,
			Status:       "failed",
			ErrorMessage: err.Error(),
		}

		s.EmailRepo.InsertNotificationHistory(ctx, &notificationHistory)
		return errors.Wrap(err, "failed to send email")
	}

	notificationHistory := models.NotificationHistory{
		Recipient:  req.Recipient,
		TemplateID: emailTemplate.ID,
		Status:     "success",
	}

	err = s.EmailRepo.InsertNotificationHistory(ctx, &notificationHistory)
	if err != nil {
		return errors.Wrap(err, "failed to insert notification history")
	}

	return nil
}
