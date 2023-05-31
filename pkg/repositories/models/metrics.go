package models

import (
	"time"
)

type Span struct {
	// BaseModel
	// ID        uint `gorm:"primary_key"`
	SpansTaskId string 

	StartTime    *time.Time
	
	EndTime *time.Time

	OperationId string

}

// By convention, gorm foreign key references are of the form {ModelName}ID
type Spans struct {
	// BaseModel
	TaskId string `gorm:"primary_key"`
	// Spans []Span `gorm:"foreignKey:SpansTaskId;references:TaskId"`
}

