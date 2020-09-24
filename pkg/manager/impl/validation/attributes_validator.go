package validation

import (
	"context"
	"fmt"

	"github.com/lyft/flyteadmin/pkg/errors"
	"github.com/lyft/flyteadmin/pkg/manager/impl/shared"
	"github.com/lyft/flyteadmin/pkg/repositories"
	runtimeInterfaces "github.com/lyft/flyteadmin/pkg/runtime/interfaces"
	"github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin"
	"google.golang.org/grpc/codes"
)

var defaultMatchableResource = admin.MatchableResource(-1)

func validateMatchingAttributes(attributes *admin.MatchingAttributes, identifier string) (admin.MatchableResource, error) {
	if attributes == nil {
		return defaultMatchableResource, shared.GetMissingArgumentError(shared.MatchingAttributes)
	}
	if attributes.GetTaskResourceAttributes() != nil {
		return admin.MatchableResource_TASK_RESOURCE, nil
	} else if attributes.GetClusterResourceAttributes() != nil {
		return admin.MatchableResource_CLUSTER_RESOURCE, nil
	} else if attributes.GetExecutionQueueAttributes() != nil {
		return admin.MatchableResource_EXECUTION_QUEUE, nil
	} else if attributes.GetExecutionClusterLabel() != nil {
		return admin.MatchableResource_EXECUTION_CLUSTER_LABEL, nil
	} else if attributes.GetPluginOverrides() != nil {
		return admin.MatchableResource_PLUGIN_OVERRIDE, nil
	}
	return defaultMatchableResource, errors.NewFlyteAdminErrorf(codes.InvalidArgument,
		"Unrecognized matching attributes type for request %s", identifier)
}

func ValidateProjectDomainAttributesUpdateRequest(ctx context.Context,
	db repositories.RepositoryInterface, config runtimeInterfaces.ApplicationConfiguration,
	request admin.ProjectDomainAttributesUpdateRequest) (
	admin.MatchableResource, error) {
	if request.Attributes == nil {
		return defaultMatchableResource, shared.GetMissingArgumentError(shared.Attributes)
	}
	if err := ValidateProjectAndDomain(ctx, db, config, request.Attributes.Project, request.Attributes.Domain); err != nil {
		return defaultMatchableResource, err
	}

	return validateMatchingAttributes(request.Attributes.MatchingAttributes,
		fmt.Sprintf("%s-%s", request.Attributes.Project, request.Attributes.Domain))
}

func ValidateProjectDomainAttributesGetRequest(ctx context.Context, db repositories.RepositoryInterface,
	config runtimeInterfaces.ApplicationConfiguration, request admin.ProjectDomainAttributesGetRequest) error {
	if err := ValidateProjectAndDomain(ctx, db, config, request.Project, request.Domain); err != nil {
		return err
	}

	return nil
}

func ValidateProjectDomainAttributesDeleteRequest(ctx context.Context, db repositories.RepositoryInterface,
	config runtimeInterfaces.ApplicationConfiguration, request admin.ProjectDomainAttributesDeleteRequest) error {
	if err := ValidateProjectAndDomain(ctx, db, config, request.Project, request.Domain); err != nil {
		return err
	}

	return nil
}

func ValidateWorkflowAttributesUpdateRequest(ctx context.Context, db repositories.RepositoryInterface,
	config runtimeInterfaces.ApplicationConfiguration, request admin.WorkflowAttributesUpdateRequest) (
	admin.MatchableResource, error) {
	if request.Attributes == nil {
		return defaultMatchableResource, shared.GetMissingArgumentError(shared.Attributes)
	}
	if err := ValidateProjectAndDomain(ctx, db, config, request.Attributes.Project, request.Attributes.Domain); err != nil {
		return defaultMatchableResource, err
	}
	if err := ValidateEmptyStringField(request.Attributes.Workflow, shared.Name); err != nil {
		return defaultMatchableResource, err
	}

	return validateMatchingAttributes(request.Attributes.MatchingAttributes,
		fmt.Sprintf("%s-%s-%s", request.Attributes.Project, request.Attributes.Domain, request.Attributes.Workflow))
}

func ValidateWorkflowAttributesGetRequest(ctx context.Context, db repositories.RepositoryInterface,
	config runtimeInterfaces.ApplicationConfiguration, request admin.WorkflowAttributesGetRequest) error {
	if err := ValidateProjectAndDomain(ctx, db, config, request.Project, request.Domain); err != nil {
		return err
	}
	if err := ValidateEmptyStringField(request.Workflow, shared.Name); err != nil {
		return err
	}

	return nil
}

func ValidateWorkflowAttributesDeleteRequest(ctx context.Context, db repositories.RepositoryInterface,
	config runtimeInterfaces.ApplicationConfiguration, request admin.WorkflowAttributesDeleteRequest) error {
	if err := ValidateProjectAndDomain(ctx, db, config, request.Project, request.Domain); err != nil {
		return err
	}
	if err := ValidateEmptyStringField(request.Workflow, shared.Name); err != nil {
		return err
	}

	return nil
}

func ValidateListAllMatchableAttributesRequest(request admin.ListMatchableAttributesRequest) error {
	if _, ok := admin.MatchableResource_name[int32(request.ResourceType)]; !ok {
		return shared.GetInvalidArgumentError(shared.ResourceType)
	}
	return nil
}
