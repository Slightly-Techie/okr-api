package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model

	CompanyId string `json:"companyId"`
	Name      string `json:"name"`
}
