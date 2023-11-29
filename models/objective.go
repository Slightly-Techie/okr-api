package models

import "gorm.io/gorm"

type Objective struct {
	gorm.Model

	ObjectiveId string `json:"objectiveId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      string `json:"userid"`
	Assignee    string `json:"assignee"`
	CompanyId   string `json:"companyId"`
}
