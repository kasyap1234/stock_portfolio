package email

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/resend/resend-go/v2"
)

func GenerateEmailToken() string {
	return uuid.New().String()
}
func SendTokenEmail(token string, to string) error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	apiKey := os.Getenv("RESEND_API_KEY")

	client := resend.NewClient(apiKey)
	backendURL := os.Getenv("BACKEND_URL")
	verificationLink := fmt.Sprintf("%s/verify?token=%s", backendURL, token)
	params := &resend.SendEmailRequest{
		From:    "onboarding@resend.dev",
		To:      []string{to},
		Subject: "Verify your email",
		Html:    fmt.Sprintf("<p>Your verification token is <b>%s</b></p><p>Or click <a href='%s'>here</a> to verify.</p>", token, verificationLink),
		Text:    fmt.Sprintf("Your verification token is %s. Or open %s", token, verificationLink),
	}

	_, err = client.Emails.Send(params)
	if err != nil {
		return err
	}
	return nil
}
