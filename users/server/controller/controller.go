package controller

import (
	"context"
	pb "grpc_curd/users/proto"
	"grpc_curd/users/server/dto"
	"grpc_curd/users/server/service"

)

type userController struct{}

func NewUserControllerServer() pb.UserServiceServer {
	return userController{}
}





func  (userController) CreateUser(ctx context.Context, req *pb.ProtoCreateRequest) (*pb.ProtoCreateReponse,error){

	result := dto.NewUserRequest(req.GetReq())

	data,err := service.UsersService.CreateUser(result)

	if err != nil {
		return nil,err
	}

	response := dto.NewUserResponse(data)

	return response,nil
}

func (userController) GetUser(ctx context.Context, req *pb.ProtoGetRequest) (*pb.ProtoGetResponse,error){

	email := req.GetProtoEmail()
	pass := req.GetProtoPassword()

	data,err := service.UsersService.GetUser(email,pass)
	
	if err != nil {
		return nil,err
	}

	response := dto.NewGetUserResponse(data)

	return response,nil
}

func (userController) GetDelete(ctx context.Context, req *pb.ProtoDeleteRequest) (*pb.ProtoDeleteResponse,error){

	email := req.GetProtoEmail()
	pass := req.GetProtoPassword()

	err := service.UsersService.GetDelete(email,pass)

	if err != nil {
		return nil,err
	}

	return &pb.ProtoDeleteResponse{
		Protoemail : req.ProtoEmail,
	},nil
}

func  (userController) UpdateUser(ctx context.Context, req *pb.ProtoUpdateRequest) (*pb.ProtoUpdateResponse,error){

	result := dto.NewUpdateRequest(req.GetUpdatereq())

	data,err := service.UsersService.UpdateUser(result)

	if err != nil {
		return nil,err
	}

	response := dto.NewUpdateResponse(data)

	return response,nil
}