package handler

import (
	"context"
	"net/http"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	connectorPB "github.com/instill-ai/protogen-go/connector/v1alpha"
)

func (h *handler) CreateSourceConnector(ctx context.Context, req *connectorPB.CreateSourceConnectorRequest) (*connectorPB.CreateSourceConnectorResponse, error) {
	resp, err := h.createConnector(ctx, req)
	if err := grpc.SetHeader(ctx, metadata.Pairs("x-http-code", strconv.Itoa(http.StatusCreated))); err != nil {
		return resp.(*connectorPB.CreateSourceConnectorResponse), err
	}
	return resp.(*connectorPB.CreateSourceConnectorResponse), err
}

func (h *handler) ListSourceConnector(ctx context.Context, req *connectorPB.ListSourceConnectorRequest) (*connectorPB.ListSourceConnectorResponse, error) {
	resp, err := h.listConnector(ctx, req)
	return resp.(*connectorPB.ListSourceConnectorResponse), err
}

func (h *handler) GetSourceConnector(ctx context.Context, req *connectorPB.GetSourceConnectorRequest) (*connectorPB.GetSourceConnectorResponse, error) {
	resp, err := h.getConnector(ctx, req)
	return resp.(*connectorPB.GetSourceConnectorResponse), err
}

func (h *handler) UpdateSourceConnector(ctx context.Context, req *connectorPB.UpdateSourceConnectorRequest) (*connectorPB.UpdateSourceConnectorResponse, error) {
	resp, err := h.updateConnector(ctx, req)
	return resp.(*connectorPB.UpdateSourceConnectorResponse), err
}

func (h *handler) DeleteSourceConnector(ctx context.Context, req *connectorPB.DeleteSourceConnectorRequest) (*connectorPB.DeleteSourceConnectorResponse, error) {
	resp, err := h.deleteConnector(ctx, req)
	if err := grpc.SetHeader(ctx, metadata.Pairs("x-http-code", strconv.Itoa(http.StatusNoContent))); err != nil {
		return &connectorPB.DeleteSourceConnectorResponse{}, err
	}
	return resp.(*connectorPB.DeleteSourceConnectorResponse), err
}

func (h *handler) LookUpSourceConnector(ctx context.Context, req *connectorPB.LookUpSourceConnectorRequest) (*connectorPB.LookUpSourceConnectorResponse, error) {
	resp, err := h.lookUpConnector(ctx, req)
	return resp.(*connectorPB.LookUpSourceConnectorResponse), err
}

func (h *handler) RenameSourceConnector(ctx context.Context, req *connectorPB.RenameSourceConnectorRequest) (*connectorPB.RenameSourceConnectorResponse, error) {
	resp, err := h.renameConnector(ctx, req)
	return resp.(*connectorPB.RenameSourceConnectorResponse), err
}
