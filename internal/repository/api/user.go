package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"example.com/tracker/internal/entity"
)

func (ac *ApiClient) GetUsersInfo(user *entity.User) (*entity.User, error) {
	resp, err := http.Get(fmt.Sprintf("%s/info?passportSerie=%d&passportNumber=%d", ac.ApiClientConfig.APIURL, user.PassportSerie, user.PassportNumber))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusBadRequest {
			return nil, errors.New("incorrect request")
		} else {
			return nil, errors.New("no response")
		}
	}

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}