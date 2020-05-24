package controller

import (
	"context"
	"github.com/fun-dev/fun-cloud-api/internal/directory/models"
	"github.com/fun-dev/fun-cloud-api/pkg/cloudauth"
	"github.com/fun-dev/fun-cloud-api/pkg/cloudk8s"
	"github.com/fun-dev/fun-cloud-protobuf/directory/rpc"
)

type DirectoryController struct {}

func (d DirectoryController) List(ctx context.Context, request *rpc.ListDirectoryRequest) (*rpc.ListDirectoryResponse, error) {
	values, err := cloudauth.GetValueFromMetadata(ctx, cloudauth.ContextKeyUserID)
	if err != nil {
		return nil, err
	}
	userId := values[0]
	namespace := userId
	listCmdResult, err := cloudk8s.ExecuteListCmdOnKubectl(request.DirectoryPath,request.ContainerId, namespace)
	if err != nil {
		return nil, err
	}
	parsedLsCmdData, err := models.ParseLsCmdData(listCmdResult, request.DirectoryPath)
	if err != nil {
		return nil, err
	}
	result := &rpc.ListDirectoryResponse{}
	result.ObjectInfo = &rpc.ObjectInfo{
		DataType: _convertDataTypeFromAppToGrpc(parsedLsCmdData.DataType),
		DataPath: parsedLsCmdData.DataPath,
		DataName: parsedLsCmdData.DataName,
	}
	for _, v := range parsedLsCmdData.ChildItems {
		bufObjInfo := &rpc.ObjectInfo{
			DataType:   _convertDataTypeFromAppToGrpc(v.DataType),
			DataPath:   v.DataPath,
			DataName:   v.DataName,
			ChildItems: nil,
		}
		result.ObjectInfo.ChildItems = append(result.ObjectInfo.ChildItems, bufObjInfo)
	}
	return result, nil
}

func NewDirectoryController() rpc.DirectoryServiceServer {
	return &DirectoryController{}
}

func _convertDataTypeFromAppToGrpc(dataType models.DataType) rpc.DataType {
	switch dataType {
	case models.FILE:
		return rpc.DataType_FILE
	case models.DIRECTORY:
		return rpc.DataType_DIRECTORY
	default:
		return 3
	}
}