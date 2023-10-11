package models

import "gorm.io/gorm"

type Test struct {
	gorm.Model

	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Company struct {
	gorm.Model

	CompanyId string `json:"companyId"`
	Name      string `json:"name"`
}

type Objective struct {
	gorm.Model

	ObjectiveId string `json:"objectiveId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Creator     string `json:"creator"`
	Assignee    string `json:"assignee"`
	CompanyId   string `json:"companyId"`
}

type KeyResult struct {
	gorm.Model

	KeyResultId string `json:"key_resultId"`
	ObjectiveId string `json:"objectiveId"`
	KeyResult   string `json:"key_result"`
	Status      string `json:"status"`
}
