package service

import (
	"fmt"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/erolkaldi/agency/pkg/models"
)

func CreateRegisterEmail(user *models.User) *models.Email {
	body := fmt.Sprintf("<!DOCTYPE html><html><body><h1>Wellcome to Go World</h1><h1>Please confirm your email by clicking <a href=\"http://localhost:7071/confirm/" + strconv.Itoa(user.ID) + "\" target=\"_blank\">this</a></h1></br><h1>Best Regards</h1></body></html>")
	return &models.Email{To: user.Email, Subject: "Wellcome to Go World", Body: body}

}

func SendEmail(mail models.Email, smtpInfo models.Smtp) error {
	auth := smtp.PlainAuth("", smtpInfo.Email, smtpInfo.Password, smtpInfo.Host)
	header := make(map[string]string)
	header["From"] = smtpInfo.Email
	header["To"] = mail.To
	header["Subject"] = mail.Subject
	header["MIME-version"] = "1.0"
	header["Content-Type"] = "text/html; charset=\"UTF-8\""

	var emailBody strings.Builder
	for key, value := range header {
		emailBody.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}
	emailBody.WriteString("\r\n" + mail.Body)
	err := smtp.SendMail(smtpInfo.Host+":"+strconv.Itoa(587), auth, smtpInfo.Email, []string{mail.To}, []byte(emailBody.String()))
	if err != nil {
		fmt.Println("Email Send Failed To:" + mail.To)
		return err
	}
	fmt.Println("Email Send Succeeded To:" + mail.To)
	return nil
}
