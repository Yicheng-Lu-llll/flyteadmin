package models

import (
	"time"
)


// Task execution primary key
type TaskExecutionKey struct {
	TaskKey
	NodeExecutionKey
	// *IMPORTANT* This is a pointer to an int in order to allow setting an empty ("0") value according to gorm convention.
	// Because RetryAttempt is part of the TaskExecution primary key is should *never* be null.
	RetryAttempt *uint32 `gorm:"primary_key"`
}

// By convention, gorm foreign key references are of the form {ModelName}ID
type Spans struct {
	BaseModel

	StartTime    *time.Time
	
	EndTime *time.Time
	Id core.Span_OperationId
	
}
