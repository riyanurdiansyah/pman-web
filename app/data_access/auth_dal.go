package data_access

import (
	data_access "kalbenutritionals.com/pman/app/data_access/interface"
	"kalbenutritionals.com/pman/app/helper/api"
	"kalbenutritionals.com/pman/app/helper/constanta"
)

type AuthDAL struct {
}

func NewAuthDAL() data_access.IAuthDAL {
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

func (dal *AuthDAL) GetMenus(body []byte, headers map[string]string) ([]byte, error) {
	response, err := api.PostRequest(constanta.MENU_URL, body, headers)
	return response, err
}
