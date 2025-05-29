package usecase

import (
	"bytes"
	"data-pusher/constant"
	"data-pusher/entity"
	"data-pusher/repository"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type DataUsecase struct {
	Mysql *repository.MysqlCon
}

func (du *DataUsecase) ProcessData(secret string, rawBody []byte) error {
	// Step 1: Unmarshal incoming JSON body
	var js map[string]interface{}
	if err := json.Unmarshal(rawBody, &js); err != nil {
		return errors.New("invalid JSON data")
	}

	// Step 2: Validate secret and get account
	var account entity.Accounts
	err := du.Mysql.Connection.
		Table(constant.ACCOUNT_TABLE_NAME).
		Where("secret = ? AND status = 'active'", secret).
		First(&account).Error
	if err != nil {
		return errors.New("unauthenticated or inactive account")
	}

	// Step 3: Fetch all destinations for the account using GORM
	var destinations []entity.Destination
	err = du.Mysql.Connection.
		Table("destinations").
		Where("account_id = ?", account.AccountID).
		Find(&destinations).Error
	if err != nil {
		return fmt.Errorf("failed to fetch destinations: %w", err)
	}

	// Step 4: Forward the data to each destination
	for _, dest := range destinations {
		urlStr := dest.URL
		method := strings.ToUpper(dest.Method)
		headersStr := dest.Headers

		// Parse headers JSON string
		headers := map[string]string{}
		if err := json.Unmarshal([]byte(headersStr), &headers); err != nil {
			fmt.Println("invalid headers JSON:", err)
			continue
		}

		var req *http.Request
		client := &http.Client{}

		// Build the HTTP request
		if method == "GET" {
			values := url.Values{}
			for k, v := range js {
				values.Add(k, fmt.Sprintf("%v", v))
			}
			req, err = http.NewRequest("GET", urlStr+"?"+values.Encode(), nil)
		} else {
			req, err = http.NewRequest(method, urlStr, bytes.NewBuffer(rawBody))
			if err == nil {
				req.Header.Set("Content-Type", "application/json")
			}
		}

		if err != nil {
			fmt.Println("failed to create request:", err)
			continue
		}

		// Add custom headers
		for k, v := range headers {
			req.Header.Set(k, v)
		}

		// Send the request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("failed to send request to", urlStr, "error:", err)
			continue
		}
		resp.Body.Close()
	}

	return nil
}
