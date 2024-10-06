package data_access

type IAuthDAL interface {
	GetRefreshToken() ([]byte, error)
	Login(username string, password string) ([]byte, error)
}
