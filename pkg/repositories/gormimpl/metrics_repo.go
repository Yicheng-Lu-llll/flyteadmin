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
	// "github.com/flyteorg/flyteadmin/pkg/repositories/models"
	"gorm.io/gorm"
)

// Implementation of TaskExecutionInterface.
type MetricsRepo struct {
	db               *gorm.DB
	errorTransformer flyteAdminDbErrors.ErrorTransformer
	metrics          gormMetrics
}

func (r *MetricsRepo) Create(input []*core.Span) error {
	fmt.Println("Hello World, I am a in meteics_repo.go create ")
	fmt.Println(input)
	// fmt.Println("new!!!")
	// startedAt := time.Now()
	// endAt := startedAt.Add(time.Second)
	// input := core.Span{
	// 	StartTime: timestamppb.New(startedAt),
	// 	EndTime:   timestamppb.New(endAt),
	// 	Id: &core.Span_OperationId{
	// 		OperationId: "This is a just a sample in meteics_repo.go list! ",
	// 	},
	// }


	// tx := r.db.Create(&input)
	// if tx.Error != nil {
	// 	return r.errorTransformer.ToFlyteAdminError(tx.Error)
	// }
	return nil


}

func (r *MetricsRepo) List(ctx context.Context, input * core.TaskExecutionIdentifier) ([]*core.Span, error) {
	fmt.Println("Hello World, I am a in meteics_repo.go list !!!! wao!!!")
	fmt.Println(input)

	startedAt := time.Now()
	endAt := startedAt.Add(time.Second)
	item := &core.Span{
		StartTime: timestamppb.New(startedAt),
		EndTime:   timestamppb.New(endAt),
		Id: &core.Span_OperationId{
			OperationId: "This is a just a sample in meteics_repo.go list! ",
		},
	}

	
	return []*core.Span{item}, nil
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
