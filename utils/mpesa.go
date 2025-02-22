package utils

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/joho/godotenv"
)

var cache, _ = bigcache.New(context.Background(), bigcache.DefaultConfig(50*time.Minute))

func MpesaGetAccessToken() (string, error) {
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
	fmt.Println("The ACCESS TOKEN =============>", accessToken)

	return accessToken, nil

}

func StoreAccessTokenInCache(access_token string) error {
	// store access toke to cache
	// cache.Delete("access-token")
	if err := cache.Set("access-token", []byte(access_token)); err != nil {
		return err

	}

	return nil
}
