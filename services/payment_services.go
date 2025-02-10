package services

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"ticketing-system/models"
	"ticketing-system/repositories"
	"time"

	"github.com/allegro/bigcache/v3"

	"github.com/joho/godotenv"
)

var cache, _ = bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))

type PaymentServiceInterface interface {
	GetAllPayments() ([]models.Payment, error)
	CreatePayment(payment *models.Payment) (*models.Payment, error)
	GetSinglePayment(id uint) ([]models.Payment, error)
	UpdatePayment(id uint, payment models.Payment) (*models.Payment, error)
	DeletePayment(id uint) (*models.Payment, error)
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

func (s *PaymentService) MpesaGetAccessToken() (string, error) {
	// get mpesa token
	// store it in BigCache
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return "", err
	}

	url := "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"

	// Encode username and password to Base64
	username := os.Getenv("username")
	password := os.Getenv("password")
	credentials := base64.StdEncoding.EncodeToString([]byte((username) + ":" + password))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "GoClient")
	req.Header.Add("Authorization", "Basic "+credentials)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	accessToken, ok := response["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("access_token not found in response")
	}
	StoreAccessTokenInCache(accessToken)

	return accessToken, nil

}

func StoreAccessTokenInCache(access_token string) error {
	// store access toke to cache
	if err := cache.Set("access-token", []byte(access_token)); err != nil {
		return err

	}
	return nil
}

func MpesaOnlinePayment() {
	// get access token from cache
	access_token, _ := cache.Get("access-token")
	url := "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"

	businessShortCode := "174379"
	timestamp := getTimestamp()
	passkey := "YOUR_PASSKEY_HERE"
	password := generatePassword(businessShortCode, passkey, timestamp)

	transactionType := "CustomerPayBillOnline"
	amount := "1"
	partyA := "254708374149"
	partyB := "174379"
	phoneNumber := "254708374149"
	callbackURL := "https://linnric.com"
	accountReference := "Test"
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
		return
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "GoClient")
	req.Header.Set("Authorization", "Bearer "+string(access_token))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response Status:", res.Status)
	fmt.Println("Response Body:", string(body))
}

func generatePassword(shortcode, passkey, timestamp string) string {
	data := shortcode + passkey + timestamp
	hash := sha256.Sum256([]byte(data))
	return base64.StdEncoding.EncodeToString(hash[:])
}

func getTimestamp() string {
	return time.Now().Format("20060102150405")
}
