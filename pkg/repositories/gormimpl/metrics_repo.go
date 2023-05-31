package gormimpl

import (
	"context"
	"fmt"
	// "time"

	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	// "google.golang.org/protobuf/types/known/timestamppb"
	

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


	// startTime := input[0].StartTime.AsTime()
	// endTime := input[0].EndTime.AsTime()
	// item := models.Span{
	// 	StartTime: &startTime,
	// 	EndTime:   &endTime,
	// 	SpansTaskId: taskId.TaskId.Name,
	// 	OperationId: "Hi, I am in database!!! ",
	// }

	myInput := models.Spans{
		TaskId: taskId.TaskId.Name,
		// Spans: []models.Span{item},
		Spans: []models.Span{},
	}


	r.db.Create(&myInput)
	fmt.Println("!!!pupa!!!")
	// if tx.Error != nil {
	// 	return r.errorTransformer.ToFlyteAdminError(tx.Error)
	// }




	return nil


}

func (r *MetricsRepo) List(ctx context.Context, input * core.TaskExecutionIdentifier) ([]*core.Span, error) {
	fmt.Println("Hello World, I am a in meteics_repo.go list")
	fmt.Println(input)

	// startedAt := time.Now()
	// endAt := startedAt.Add(time.Second)
	// item := &core.Span{
	// 	StartTime: timestamppb.New(startedAt),
	// 	EndTime:   timestamppb.New(endAt),
	// 	Id: &core.Span_OperationId{
	// 		OperationId: "This is a just a sample in meteics_repo.go list! ",
	// 	},
	// }



	// // Retrieve spans with a specific TaskId
	// var spans models.Spans
	// r.db.Where("task_id = ?", input.TaskId.Name).Find(&spans).Limit(1)
	// fmt.Println("!!!spans.TaskId!!!", spans.TaskId)


	// // Iterate over the spans
	// for _, span := range spans.Spans {
	// 	// Access span data
	// 	fmt.Println("SpansTaskId:", span.SpansTaskId)
	// 	fmt.Println("StartTime:", span.StartTime)
	// 	fmt.Println("EndTime:", span.EndTime)
	// 	fmt.Println("OperationId:", span.OperationId)
	// }






	
	// return []*core.Span{item}, nil
	return []*core.Span{}, nil
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
