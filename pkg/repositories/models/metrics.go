package models

import (
	"time"
)



// By convention, gorm foreign key references are of the form {ModelName}ID
type SpansModel struct {
	BaseModel

	taskId string `gorm:"primary_key"`
	
	StartTime    *time.Time
	
	EndTime *time.Time

	OperationId string
	
	
}

