package notify

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// EmailConfig holds email configuration
type EmailConfig struct {
	SMTPHost       string
	SMTPPort       string
	SenderEmail    string
	SenderPass     string
	RecipientEmail string
}

// LoadEmailConfig loads email configuration from environment variables
func LoadEmailConfig() (*EmailConfig, error) {
	godotenv.Load()

	config := &EmailConfig{
		SMTPHost:       "smtp.gmail.com",
		SMTPPort:       "465", // Try SSL port instead of TLS
		SenderEmail:    os.Getenv("SENDER_EMAIL"),
		SenderPass:     os.Getenv("SENDER_PASSWORD"),
		RecipientEmail: os.Getenv("RECIPIENT_EMAIL"),
	}

	if config.SenderEmail == "" || config.SenderPass == "" || config.RecipientEmail == "" {
		return nil, fmt.Errorf("missing email configuration in .env file")
	}

	return config, nil
}

// SendEmail - Core email sending function with SSL support
func SendEmail(subject, body string) error {
	fmt.Println("🔄 Attempting to send email...")

	config, err := LoadEmailConfig()
	if err != nil {
		fmt.Printf("❌ Config error: %v\n", err)
		return fmt.Errorf("failed to load email config: %v", err)
	}

	fmt.Printf("📧 Email config loaded - From: %s, To: %s\n", config.SenderEmail, config.RecipientEmail)

	// Create TLS configuration
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         config.SMTPHost,
	}

	// Gmail SMTP authentication
	auth := smtp.PlainAuth("", config.SenderEmail, config.SenderPass, config.SMTPHost)

	// Create email message with proper headers
	message := fmt.Sprintf("To: %s\r\n", config.RecipientEmail)
	message += fmt.Sprintf("From: %s\r\n", config.SenderEmail)
	message += fmt.Sprintf("Subject: %s\r\n", subject)
	message += "Content-Type: text/plain; charset=utf-8\r\n"
	message += "\r\n"
	message += body

	fmt.Printf("🌐 Connecting to SMTP server: %s:%s\n", config.SMTPHost, config.SMTPPort)

	// Try different connection methods based on port
	var emailErr error
	if config.SMTPPort == "465" {
		// SSL connection
		emailErr = sendEmailSSL(config, auth, message, tlsConfig)
	} else {
		// TLS connection (port 587)
		emailErr = sendEmailTLS(config, auth, message, tlsConfig)
	}

	if emailErr != nil {
		fmt.Printf("❌ SMTP error: %v\n", emailErr)
		return fmt.Errorf("failed to send email: %v", emailErr)
	}

	fmt.Println("✅ Email sent successfully!")
	return nil
}

// sendEmailTLS sends email using STARTTLS (port 587)
func sendEmailTLS(config *EmailConfig, auth smtp.Auth, message string, tlsConfig *tls.Config) error {
	return smtp.SendMail(
		config.SMTPHost+":"+config.SMTPPort,
		auth,
		config.SenderEmail,
		[]string{config.RecipientEmail},
		[]byte(message),
	)
}

// sendEmailSSL sends email using SSL (port 465)
func sendEmailSSL(config *EmailConfig, auth smtp.Auth, message string, tlsConfig *tls.Config) error {
	conn, err := tls.Dial("tcp", config.SMTPHost+":"+config.SMTPPort, tlsConfig)
	if err != nil {
		return err
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, config.SMTPHost)
	if err != nil {
		return err
	}
	defer client.Quit()

	if err = client.Auth(auth); err != nil {
		return err
	}

	if err = client.Mail(config.SenderEmail); err != nil {
		return err
	}

	if err = client.Rcpt(config.RecipientEmail); err != nil {
		return err
	}

	writer, err := client.Data()
	if err != nil {
		return err
	}

	_, err = writer.Write([]byte(message))
	if err != nil {
		return err
	}

	return writer.Close()
}

// EmailToConsole - Send status email like PrintToConsole
func EmailSummary(message string) error {
	subject := "WeatherWatch - Status Update"
	return SendEmail(subject, message)
}

// SendEmailAlert - Send alert email like PrintAlert
func SendEmailAlert(message string) error {
	subject := "🚨 WeatherWatch Alert!"
	emailBody := fmt.Sprintf("🚨 WEATHER ALERT! 🚨\n\n%s\n\nStay safe!\nWeatherWatch Alert System", message)
	return SendEmail(subject, emailBody)
}

// SendEmailStatus - Send status email like PrintStatus (fixed function name)
func SendEmailStatus(message string) error {
	subject := "✅ WeatherWatch - Status OK"
	emailBody := fmt.Sprintf("✅ %s\n\nWeatherWatch is monitoring your location.", message)
	return SendEmail(subject, emailBody)
}

// SendWeatherAlert - Comprehensive weather alert with formatted details
func SendWeatherAlert(location, timestamp string, alerts []string) error {
	subject := fmt.Sprintf("🚨 WeatherWatch Alert - %s", location)

	body := fmt.Sprintf(`Hello!

Your WeatherWatch service has detected weather conditions that exceed your configured thresholds:

📍 Location: %s
🕓 Time: %s

🚨 Alert Conditions:
%s

Please take appropriate precautions based on these weather conditions.

Stay safe!
WeatherWatch Alert System 🌤️

---
This is an automated message from your WeatherWatch monitoring service.
To stop receiving these alerts, please update your monitoring configuration.
`, location, timestamp, "• "+strings.Join(alerts, "\n• "))

	return SendEmail(subject, body)
}
