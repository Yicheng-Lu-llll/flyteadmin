package models

import (
	"time"
)

type TimeItSpans struct {
	BaseModel
	SpansTaskId string `gorm:"index"`

	StartTime    *time.Time
	
	EndTime *time.Time

	OperationId string

}

// By convention, gorm foreign key references are of the form {ModelName}ID
type Spans struct {
	BaseModel
	TaskId string `gorm:"primary_key"`
	Spans []TimeItSpans `gorm:"foreignKey:SpansTaskId;references:TaskId"`
}

func (s TimeItSpans) TableName() string {
	return "timeItSpans"
	}

func (s Spans) TableName() string {
	return "spans"
	}

	
