package data_access

type IAuthDAL interface {
	GetRefreshToken() ([]byte, error)
	Login(body []byte, headers map[string]string) ([]byte, error)
	GetMenus(body []byte, headers map[string]string) ([]byte, error)
}
