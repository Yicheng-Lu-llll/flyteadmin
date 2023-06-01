package gormimpl

import (
	"context"
	"fmt"
	"time"

	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	"google.golang.org/protobuf/types/known/timestamppb"
	

	"github.com/flyteorg/flytestdlib/promutils"

	flyteAdminDbErrors "github.com/flyteorg/flyteadmin/pkg/repositories/errors"
	"github.com/flyteorg/flyteadmin/pkg/repositories/interfaces"
	"github.com/flyteorg/flyteadmin/pkg/repositories/models"
	"gorm.io/gorm"
)

// Implementation of TaskExecutionInterface.
type MetricsRepo struct {
	db               *gorm.DB
	errorTransformer flyteAdminDbErrors.ErrorTransformer
	metrics          gormMetrics
}

func (r *MetricsRepo) Create(input []*core.Span, taskId *core.TaskExecutionIdentifier) error {
	fmt.Println("Hello World, I am a in meteics_repo.go create ")
	// let's use taskId.TaskId.Name as task identifier now
	fmt.Println("taskId.TaskId.Name", taskId.TaskId.Name)



	startTime := timestamppb.New(time.Now()).AsTime()
	endTime := timestamppb.New(time.Now().Add(time.Second)).AsTime()

	itema := models.Span{
		StartTime: &startTime,
		EndTime:   &endTime,
		TaskId: taskId.TaskId.Name,
		Description: "Hi, I am in database 1!!! ",
	}

	itemb := models.Span{
		StartTime: &startTime,
		EndTime:   &endTime,
		TaskId: taskId.TaskId.Name,
		Description: "Hi, I am in database 2!!! ",
	}

	items := []models.Span{itema, itemb}


	r.db.Create(&items)
	fmt.Println("pupa")


	return nil


}

func (r *MetricsRepo) List(ctx context.Context, input * core.TaskExecutionIdentifier) ([]*core.Span, error) {
	fmt.Println("Hello World, I am a in meteics_repo.go list")
	fmt.Println(input)

	

	// Retrieve spans with a specific TaskId
	var items []models.Span
	r.db.Where("task_id = ?", input.TaskId.Name).Find(&items)
	fmt.Println("!!!length:!!!", len(items))
	// print each span in items
	for _, item := range items {
		fmt.Println(item.Description)
	}
	var coreSpans []*core.Span
	return coreSpans, nil

}

// Returns an instance of TaskExecutionRepoInterface
func NewMetricsRepo(
	db *gorm.DB, errorTransformer flyteAdminDbErrors.ErrorTransformer, scope promutils.Scope) interfaces.MetricsRepoInterface {
	metrics := newMetrics(scope)
	return &MetricsRepo{
		db:               db,
		errorTransformer: errorTransformer,
		metrics:          metrics,
	}
}
