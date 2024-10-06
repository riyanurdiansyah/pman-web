package data_access

import (
	"kalbenutritionals.com/pman/app/helper/api"
	"kalbenutritionals.com/pman/app/helper/constanta"
)

type AuthDAL struct {
}

func NewAuthDAL() *AuthDAL {
	return &AuthDAL{}
}

func (dal *AuthDAL) GetRefreshToken() ([]byte, error) {
	response, err := api.PostRefreshToken(constanta.TOKEN_URL, nil)
	return response, err
}

func (dal *AuthDAL) Login(body []byte, headers map[string]string) ([]byte, error) {
	response, err := api.PostRequest(constanta.LOGIN_URL, body, headers)
	return response, err
}
