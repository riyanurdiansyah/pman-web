package business_logic

import model_response "kalbenutritionals.com/pman/app/helper/model/response"

type IAuthBL interface {
	GetTokenAccess() (string, error)
	Login(body []byte, headers map[string]string) (*model_response.SigninResponse, error)
	GetMenus(body []byte, headers map[string]string) (*model_response.SigninResponse, error)
}
