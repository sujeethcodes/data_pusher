package entity

type Accounts struct {
	Name      string `json:"name"`
	AccountID string `json:"account_id" gorm:"unique;not null"` // Required & Unique
	Email     string `json:"email" gorm:"unique;not null"`      // Required & Unique
	Website   string `json:"website,omitempty"`
	Token     string `json:"token"`
	Status    string `json:"status"`
}

type DeleteReq struct {
	AccountId string `json:"account_id"`
	Status    string `json:"status"`
}
