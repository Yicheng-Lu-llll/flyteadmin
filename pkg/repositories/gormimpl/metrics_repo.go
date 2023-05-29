package gormimpl

import (
	"context"
	"fmt"

	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	

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

func (r *MetricsRepo) Create() error {
	fmt.Println("Hello World, I am a in meteics_repo.go create ")

	return nil
}

func (r *MetricsRepo) List(ctx context.Context, input * core.TaskExecutionIdentifier) error {
	fmt.Println("Hello World, I am a in meteics_repo.go list ")
	fmt.Println(input)

	return nil
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
