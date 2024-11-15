package schematype

import (
	"encoding/json"
)

type Role = string

const (
	MarsAdmin Role = "mars_admin"
)

type UserInfo struct {
	ID      string   `json:"id"`
	Email   string   `json:"email"`
	Name    string   `json:"name"`
	Picture string   `json:"picture"`
	Roles   []string `json:"roles"`

	LogoutUrl string `json:"logout_url"`
}

func (ui *UserInfo) Json() string {
	marshal, _ := json.Marshal(ui)
	return string(marshal)
}

func (ui *UserInfo) GetID() string {
	return ui.ID
}

func (ui *UserInfo) IsAdmin() bool {
	for _, role := range ui.Roles {
		if role == MarsAdmin {
			return true
		}
	}
	return false
}

type UploadType string

const (
	Local UploadType = "local"
	S3    UploadType = "s3"
)
