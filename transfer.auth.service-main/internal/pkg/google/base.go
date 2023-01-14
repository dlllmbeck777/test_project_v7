package google

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/transferMVP/transfer.webapp/internal/config"
	"github.com/transferMVP/transfer.webapp/internal/models"
	"io/ioutil"
	"net/http"
	"time"
)

func GetGoogleUser(access_token string, id_token string) (*models.GoogleUserResult, bool, error) {

	rootUrl := config.Config.GoogleSettings.HostVerifyToken + access_token
	//fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=%s", access_token)
	fmt.Println(rootUrl)
	req, err := http.NewRequest("GET", rootUrl, nil)
	if err != nil {
		return nil, false, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", id_token))

	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, false, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, false, errors.New("could not retrieve user")
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, false, err
	}

	var GoogleUserRes map[string]interface{}

	if err := json.Unmarshal(resBody, &GoogleUserRes); err != nil {
		return nil, false, err
	}

	userBody := &models.GoogleUserResult{
		Id:             GoogleUserRes["id"].(string),
		Email:          GoogleUserRes["email"].(string),
		Verified_email: GoogleUserRes["verified_email"].(bool),
		Name:           GoogleUserRes["name"].(string),
		Given_name:     GoogleUserRes["given_name"].(string),
		Picture:        GoogleUserRes["picture"].(string),
		Locale:         GoogleUserRes["locale"].(string),
	}

	return userBody, userBody.Verified_email, nil
}
