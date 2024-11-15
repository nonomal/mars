package services

import (
	"context"

	"github.com/duc-cnzj/mars/api/v5/endpoint"
	"github.com/duc-cnzj/mars/v5/internal/mlog"
	"github.com/duc-cnzj/mars/v5/internal/repo"
)

var _ endpoint.EndpointServer = (*endpointSvc)(nil)

type endpointSvc struct {
	endpoint.UnimplementedEndpointServer

	logger mlog.Logger
	epRepo repo.EndpointRepo
}

func NewEndpointSvc(logger mlog.Logger, epRepo repo.EndpointRepo) endpoint.EndpointServer {
	return &endpointSvc{logger: logger.WithModule("services/endpoint"), epRepo: epRepo}
}

func (e *endpointSvc) InNamespace(ctx context.Context, request *endpoint.InNamespaceRequest) (*endpoint.InNamespaceResponse, error) {
	res, err := e.epRepo.InNamespace(ctx, int(request.NamespaceId))
	if err != nil {
		return nil, err
	}
	return &endpoint.InNamespaceResponse{Items: res}, nil
}

func (e *endpointSvc) InProject(ctx context.Context, request *endpoint.InProjectRequest) (*endpoint.InProjectResponse, error) {
	res, err := e.epRepo.InProject(ctx, int(request.ProjectId))
	if err != nil {
		return nil, err
	}
	return &endpoint.InProjectResponse{Items: res}, nil
}
