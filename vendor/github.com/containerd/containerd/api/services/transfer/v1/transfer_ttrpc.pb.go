// Code generated by protoc-gen-go-ttrpc. DO NOT EDIT.
// source: github.com/containerd/containerd/api/services/transfer/v1/transfer.proto
package transfer

import (
	context "context"
	ttrpc "github.com/containerd/ttrpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type TTRPCTransferService interface {
	Transfer(context.Context, *TransferRequest) (*emptypb.Empty, error)
}

func RegisterTTRPCTransferService(srv *ttrpc.Server, svc TTRPCTransferService) {
	srv.RegisterService("containerd.services.transfer.v1.Transfer", &ttrpc.ServiceDesc{
		Methods: map[string]ttrpc.Method{
			"Transfer": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req TransferRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Transfer(ctx, &req)
			},
		},
	})
}

type ttrpctransferClient struct {
	client *ttrpc.Client
}

func NewTTRPCTransferClient(client *ttrpc.Client) TTRPCTransferService {
	return &ttrpctransferClient{
		client: client,
	}
}

func (c *ttrpctransferClient) Transfer(ctx context.Context, req *TransferRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.services.transfer.v1.Transfer", "Transfer", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}