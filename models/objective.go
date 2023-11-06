package models

import "gorm.io/gorm"

type Objective struct {
	gorm.Model

	ObjectiveId string `json:"objectiveId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Creator     string `json:"creator"`
	Assignee    string `json:"assignee"`
	CompanyId   string `json:"companyId"`
}
