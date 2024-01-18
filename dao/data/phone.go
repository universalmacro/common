package data

type PhoneNumber struct {
	CountryCode string `json:"countryCode" gorm:"type:CHAR(6);index:,unique,composite:phone_number"`
	Number      string `json:"number" gorm:"type:CHAR(11);index:,unique,composite:phone_number;column:phone_number"`
}
