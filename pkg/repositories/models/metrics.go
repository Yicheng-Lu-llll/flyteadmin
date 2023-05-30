package models

import (
	"time"
)

type Span struct {

	ID           uint       `gorm:"primary_key;unique"`
	
	SpansTaskId string 

	StartTime    *time.Time
	
	EndTime *time.Time

	OperationId string

}

// By convention, gorm foreign key references are of the form {ModelName}ID
type Spans struct {
	BaseModel

	TaskId string `gorm:"primary_key;unique"`

	Spans []Span `gorm:"foreignkey:SpansTaskId"`

}

