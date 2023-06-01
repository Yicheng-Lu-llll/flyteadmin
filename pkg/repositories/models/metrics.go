package models

import (
	"time"
)

type Span struct {
	BaseModel

	TaskId string `gorm:"primaryKey;index"`
	Description string `gorm:"primaryKey;index"`

	WorkflowId string 

	StartTime    *time.Time
	
	EndTime *time.Time

	

}


	
