package send_to

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type MailRequest struct {
	ToEmail     string `json:"to_email"`
	MessageBody string `json:"message_body"`
	Subject     string `json:"subject"`
	Attachment  string `json:"attachment"`
}

func SendEmailToJavaByApi(otp string, email string, purpose string) error {
	// URL API
	postUrl := "http://localhost:8080/email/send_text"
	// Data json
	mailRequest := MailRequest{
		ToEmail:     email,
		MessageBody: "OPT IS " + otp,
		Subject:     "Verify OTP " + purpose,
		Attachment:  "path/to/email",
	}
	// Convert struct to json
	requestBody, err := json.Marshal(mailRequest)
	if err != nil {
		return err
	}
	// Create request
	req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	// PUT header
	req.Header.Set("Content-Type", "application/json")
	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)
	return nil
}
