package gormimpl

import (
	"context"
	"fmt"

	

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

func (r *MetricRepo) Create() error {
	fmt.Println("Hello World, I am a in meteics_repo.go create ")

	return nil
}



// Returns an instance of TaskExecutionRepoInterface
func NewMetricsRepoRepo(
	db *gorm.DB, errorTransformer flyteAdminDbErrors.ErrorTransformer, scope promutils.Scope) interfaces.TaskExecutionRepoInterface {
	metrics := newMetrics(scope)
	return &MetricsRepo{
		db:               db,
		errorTransformer: errorTransformer,
		metrics:          metrics,
	}
}
