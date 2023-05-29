package interfaces

import (
	"context"

	// "github.com/flyteorg/flyteadmin/pkg/repositories/models"
	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
)

// Defines the interface for interacting with task execution models.
type MetricsRepoInterface interface {
	// Inserts a task execution model into the database store.
	Create() error
	List(ctx context.Context, input *core.TaskExecutionIdentifier) ([]*core.Span, error)

}
