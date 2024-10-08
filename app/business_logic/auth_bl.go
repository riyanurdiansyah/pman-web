package business_logic

import (
	"encoding/json"
	"fmt"
	"log"

	business_logic "kalbenutritionals.com/pman/app/business_logic/interface"
	data_access "kalbenutritionals.com/pman/app/data_access/interface"
	model_request "kalbenutritionals.com/pman/app/helper/model/request"
	model_response "kalbenutritionals.com/pman/app/helper/model/response"
)

type AuthBL struct {
	AuthDAL data_access.IAuthDAL
}

// GetMenus implements business_logic.IAuthBL.

func NewAuthBL(authDAL data_access.IAuthDAL) business_logic.IAuthBL {
	return &AuthBL{AuthDAL: authDAL}

}
func (bl *AuthBL) GetMenus(body []byte, headers map[string]string) ([]model_response.MenuDataResponse, error) {
	var menuResponse model_response.MenuResponse
	var menus []model_response.MenuDataResponse
	response, err := bl.AuthDAL.GetMenus(body, headers)
	errJson := json.Unmarshal(response, &menuResponse)
	if errJson != nil {
		err = errJson
	}

	log.Println("CEK RES API : " + menuResponse.ObjData.ObjData)

	errConvert := json.Unmarshal([]byte(menuResponse.ObjData.ObjData), &menus)

	if errConvert != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	for _, item := range menus {
		fmt.Printf("Menu ID: %d, Menu Name: %s\n", item.IntMenuID, item.TxtMenuName)
	}

	return menus, err
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
