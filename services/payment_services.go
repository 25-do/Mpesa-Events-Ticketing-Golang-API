package services

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"ticketing-system/models"
	"ticketing-system/repositories"
	"time"

	"github.com/allegro/bigcache/v3"
)

var cache, _ = bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))

type PaymentServiceInterface interface {
	GetAllPayments() ([]models.Payment, error)
	CreatePayment(payment *models.Payment) (*models.Payment, error)
	GetSinglePayment(id uint) ([]models.Payment, error)
	UpdatePayment(id uint, payment models.Payment) (*models.Payment, error)
	DeletePayment(id uint) (*models.Payment, error)
	MpesaOnlinePayment(amounts string, phonenumber string, orgernizerId uint) error
	// MpesaGetAccessToken() (string, error)
}

type PaymentService struct {
	PaymentRepo repositories.PaymentRepositoryInterface
}

func NewPaymentService(paymentRepo repositories.PaymentRepositoryInterface) PaymentServiceInterface {
	return &PaymentService{PaymentRepo: paymentRepo}
}

func (dc *PaymentService) GetAllPayments() ([]models.Payment, error) {
	return dc.PaymentRepo.GetAllPayments()

}

func (dc *PaymentService) CreatePayment(payment *models.Payment) (*models.Payment, error) {

	return dc.PaymentRepo.CreatePayment(payment)
}

func (s *PaymentService) GetSinglePayment(id uint) ([]models.Payment, error) {
	return s.PaymentRepo.GetSinglePayment(id)
}

func (s *PaymentService) UpdatePayment(id uint, payment models.Payment) (*models.Payment, error) {
	return s.PaymentRepo.UpdatePayment(id, payment)
}

func (s *PaymentService) DeletePayment(id uint) (*models.Payment, error) {
	return s.PaymentRepo.DeletePayment(id)
}

func (s *PaymentService) MpesaOnlinePayment(amounts string, phonenumber string, orgernizerId uint) error {

	// get access token from cache
	fmt.Println("this is the ID ==================>", orgernizerId)
	org, _ := s.PaymentRepo.OrganizerId(uint(orgernizerId))
	access_token, _ := cache.Get("access-token")
	fmt.Println("My access is ++++++++++++", access_token)
	url := "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"

	businessShortCode := "174379"
	timestamp := "20160216165627"
	// passkey := "YOUR_PASSKEY_HERE"
	password := "MTc0Mzc5YmZiMjc5ZjlhYTliZGJjZjE1OGU5N2RkNzFhNDY3Y2QyZTBjODkzMDU5YjEwZjc4ZTZiNzJhZGExZWQyYzkxOTIwMTYwMjE2MTY1NjI3"

	transactionType := "CustomerPayBillOnline"
	amount := amounts
	partyA := phonenumber
	partyB := org.TillPayBillNumber
	phoneNumber := phonenumber
	callbackURL := "https://linnric.com"
	accountReference := org.AccountReference
	transactionDesc := "Test"

	payload := []byte(fmt.Sprintf(`{
		"BusinessShortCode": "%s",
		"Password": "%s",
		"Timestamp": "%s",
		"TransactionType": "%s",
		"Amount": "%s",
		"PartyA": "%s",
		"PartyB": "%s",
		"PhoneNumber": "%s",
		"CallBackURL": "%s",
		"AccountReference": "%s",
		"TransactionDesc": "%s"
	}`, businessShortCode, password, timestamp, transactionType, amount, partyA, partyB, phoneNumber, callbackURL, accountReference, transactionDesc))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "GoClient")
	req.Header.Set("Authorization", "Bearer "+string(access_token))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Request error:", err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return err
	}

	fmt.Println("Response Status:", res.Status)
	fmt.Println("Response Body:", string(body))
	return nil
}

func generatePassword(shortcode, passkey, timestamp string) string {
	data := shortcode + passkey + timestamp
	hash := sha256.Sum256([]byte(data))
	return base64.StdEncoding.EncodeToString(hash[:])
}

func getTimestamp() string {
	return time.Now().Format("20060102150405")
}
