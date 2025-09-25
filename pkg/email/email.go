package email

import (
	"fmt"
	"net/smtp"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/go-laundry/internal/config"
)

type EmailClient struct {
	cfg *config.Config
}

func NewClient(cfg *config.Config) *EmailClient {
	return &EmailClient{
		cfg: cfg,
	}
}

func (e *EmailClient) Send(to []string, subject, body string) error {
	cfg := e.cfg.Email
	msg := fmt.Appendf(nil, "Subject: %s\n\n%s", subject, body)
	auth := smtp.PlainAuth("", cfg.Sender, cfg.AppPassword, cfg.SmtpHost)

	addr := fmt.Sprintf("%s:%s", cfg.SmtpHost, cfg.SmtpPort)
	err := smtp.SendMail(addr, auth, cfg.Sender, to, msg)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to send email: %v", err)
		return err
	}

	log.Info().Msg("Email sent successfully")
	return nil
}
