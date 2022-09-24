package dto

import (
	"grpc_curd/users/server/repository"
	pb "grpc_curd/users/proto"
)


func NewUserRequest(data *pb.ProtoUser) *repository.User{
	return &repository.User{
		UserID : data.ProtoID,
		UserName : data.ProtoName,
		UserEmail : data.ProtoEmail,
		UserPassword : data.ProtoPassword,
		UserPhoneno : data.ProtoPhoneno,
		UserAddress : data.ProtoAddress,

	}
}


func NewUserResponse(data *repository.User) *pb.ProtoCreateReponse{
	user := &pb.ProtoUser{
		ProtoID : data.UserID,
		ProtoName : data.UserName,
		ProtoEmail : data.UserEmail,
		ProtoPassword : data.UserPassword,
		ProtoPhoneno : data.UserPhoneno,
		ProtoAddress : data.UserAddress,
	}
	return &pb.ProtoCreateReponse{
		Res : user,
	}
}



func NewGetUserResponse(data *repository.User) *pb.ProtoGetResponse{
	user := &pb.ProtoUser{
		ProtoID : data.UserID,
		ProtoName : data.UserName,
		ProtoEmail : data.UserEmail,
		ProtoPhoneno : data.UserPhoneno,
		ProtoAddress : data.UserAddress,
	}
	return &pb.ProtoGetResponse{
		Res : user,
	}
}

func NewUpdateRequest(data *pb.ProtoUser) *repository.User{
	return &repository.User{
		UserName : data.ProtoName,
		UserEmail : data.ProtoEmail,
		UserPassword : data.ProtoPassword,
		UserPhoneno : data.ProtoPhoneno,
		UserAddress : data.ProtoAddress,
	}
}

func NewUpdateResponse(data *repository.User) *pb.ProtoUpdateResponse{
	user := &pb.ProtoUser{
		ProtoID : data.UserID,
		ProtoName : data.UserName,
		ProtoEmail : data.UserEmail,
		ProtoPhoneno : data.UserPhoneno,
		ProtoAddress : data.UserAddress,
	}
	return &pb.ProtoUpdateResponse{
		Updateres : user,
	}
}