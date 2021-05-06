package models

import (
	"github.com/jinzhu/gorm"
)

type Course struct {
	gorm.Model
	Name            string `json:"name"`
	Details         string `json:"details"`
	NoOfSemesters   int8 `json:"noOfSemesters"`
	AdmissionsOpen  bool `json:"admissionsOpen"`
	PreviousBatches []Batch `json:"previousBatches"`
}
