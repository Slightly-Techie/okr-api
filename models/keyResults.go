package models

import "gorm.io/gorm"

type KeyResult struct {
	gorm.Model

	KeyResultId string `json:"key_resultId"`
	ObjectiveId string `json:"objectiveId"`
	KeyResult   string `json:"key_result"`
	Status      string `json:"status"`
}
