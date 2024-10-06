package business_logic

import (
	"encoding/json"

	"kalbenutritionals.com/pman/app/data_access"
	model_request "kalbenutritionals.com/pman/app/helper/model/request"
	model_response "kalbenutritionals.com/pman/app/helper/model/response"
)

type AuthBL struct {
	AuthDAL *data_access.AuthDAL
}

func NewAuthBL(authDAL *data_access.AuthDAL) *AuthBL {
	return &AuthBL{AuthDAL: authDAL}
}

func (bl *AuthBL) GetTokenAccess() (string, error) {
	var tokenResponse model_request.TokenResponse
	response, err := bl.AuthDAL.GetRefreshToken()

	errJson := json.Unmarshal(response, &tokenResponse)
	if errJson != nil {
		err = errJson
	}
	return tokenResponse.AccessToken, err
}

func (bl *AuthBL) Login(body []byte, headers map[string]string) (*model_response.SigninResponse, error) {
	var signinResponse model_response.SigninResponse
	response, err := bl.AuthDAL.Login(body, headers)

	errJson := json.Unmarshal(response, &signinResponse)
	if errJson != nil {
		err = errJson
	}

	return &signinResponse, err
}
