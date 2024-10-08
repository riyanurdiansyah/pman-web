package model_response

type MenuResponse struct {
	BitSuccess    bool                `json:"bitSuccess"`
	ObjData       MenuObjDataResponse `json:"objData"`
	TxtMessage    string              `json:"txtMessage"`
	TxtStackTrace string              `json:"txtStackTrace"`
	TxtGUID       string              `json:"txtGUID"`
}

type MenuObjDataResponse struct {
	ObjData string `json:"objData"`
}

type MenuDataResponse struct {
	IntMenuID      int                `json:"intMenuID"`
	IntProgramID   int                `json:"intProgramID"`
	IntParentID    int                `json:"intParentID"`
	IntModuleID    int                `json:"intModuleID"`
	TxtMenuCode    string             `json:"txtMenuCode"`
	TxtMenuName    string             `json:"txtMenuName"`
	TxtPrefix      string             `json:"txtPrefix"`
	TxtDescription string             `json:"txtDescription"`
	TxtLink        string             `json:"txtLink"`
	IntOrder       int                `json:"intOrder"`
	BitActive      bool               `json:"bitActive"`
	TxtGUID        string             `json:"txtGUID"`
	TxtInsertedBy  string             `json:"txtInsertedBy"`
	TxtUpdatedBy   string             `json:"txtUpdatedBy"`
	MProgram       interface{}        `json:"mProgram"`
	BitDeleted     *bool              `json:"bitDeleted"`
	IntCompanyID   *int               `json:"intCompanyID"`
	TxtIcon        *string            `json:"txtIcon"`
	ItemList       []MenuDataResponse `json:"itemList"`
}
