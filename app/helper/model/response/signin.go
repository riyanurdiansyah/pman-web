package model_response

// Mendefinisikan struct untuk response API
type SigninResponse struct {
	BitSuccess    bool     `json:"bitSuccess"`
	ObjData       UserData `json:"objData"` // Menggunakan struct UserData
	TxtMessage    string   `json:"txtMessage"`
	TxtStackTrace string   `json:"txtStackTrace"`
	TxtGUID       string   `json:"txtGUID"`
}

// Mendefinisikan struct untuk data pengguna
type UserData struct {
	IntUserID     int     `json:"intUserID"`
	TxtUserName   string  `json:"txtUserName"`
	TxtFullName   string  `json:"txtFullName"`
	TxtNick       *string `json:"txtNick"` // Bisa nil, jadi menggunakan pointer
	TxtEmployeeID string  `json:"txtEmployeeID"`
	TxtEmail      string  `json:"txtEmail"`
	BitActive     bool    `json:"bitActive"`
	LtRoles       []Role  `json:"ltRoles"` // List of roles
}

// Mendefinisikan struct untuk role
type Role struct {
	IntRoleId      int          `json:"IntRoleId"`
	TxtRoleName    string       `json:"TxtRoleName"`
	BitSuperuser   bool         `json:"BitSuperuser"`
	TxtGuid        string       `json:"TxtGuid"`
	TxtCreatedBy   *string      `json:"TxtCreatedBy"`
	DtmCreatedDate *string      `json:"DtmCreatedDate"`
	TxtUpdatedBy   string       `json:"TxtUpdatedBy"`
	DtmUpdatedDate *string      `json:"DtmUpdatedDate"`
	MRoleAccesses  []RoleAccess `json:"MRoleAccesses"`
}

// Mendefinisikan struct untuk role access
type RoleAccess struct {
	IntRoleAccessId int     `json:"IntRoleAccessId"`
	IntRoleId       int     `json:"IntRoleId"`
	IntModuleId     int     `json:"IntModuleId"`
	BitEdit         bool    `json:"BitEdit"`
	BitView         bool    `json:"BitView"`
	BitDelete       bool    `json:"BitDelete"`
	BitPrint        *bool   `json:"BitPrint"` // Bisa nil, jadi menggunakan pointer
	TxtGuid         string  `json:"TxtGuid"`
	TxtCreatedBy    *string `json:"TxtCreatedBy"`   // Bisa nil, jadi menggunakan pointer
	DtmCreatedDate  *string `json:"DtmCreatedDate"` // Bisa nil, jadi menggunakan pointer
	TxtUpdatedBy    string  `json:"TxtUpdatedBy"`
	DtmUpdatedDate  *string `json:"DtmUpdatedDate"` // Bisa nil, jadi menggunakan pointer
	IntModule       *int    `json:"IntModule"`      // Bisa nil, jadi menggunakan pointer
	IntRole         *int    `json:"IntRole"`        // Bisa nil, jadi menggunakan pointer
}

// Mendefinisikan struct untuk user role
type UserRole struct {
	// Sesuaikan dengan field yang diperlukan di sini
}
