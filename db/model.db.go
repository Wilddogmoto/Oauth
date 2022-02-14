package db

type Services struct {
	Id          int64  `gorm:"primary_key" json:"id"`
	Name        string `json:"name"`
	ClientId    string `json:"client_id"`
	SecretCode  string `json:"secret_code"`
	RedirectUri string `json:"redirect_uri"`
}

type ServiceLogins struct {
	Id    int64  `gorm:"primary_key" json:"id"`
	State string `json:"state"`
}
