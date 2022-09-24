package service

import (
	"grpc_curd/users/server/repository"
	"grpc_curd/users/server/utils"
	"github.com/rs/xid"
	
)

var (
	UsersService UsersServiceInterface = &usersService{}
)

type UsersServiceInterface interface {
	CreateUser(data *repository.User)(*repository.User, error)
	GetUser(useremail string,userpass string) (*repository.User, error)
	GetDelete(useremail string,userpass string) (error)
	UpdateUser(data *repository.User)(*repository.User, error)
}

type usersService struct{}

func (s *usersService) CreateUser(data *repository.User)(*repository.User, error) {

	if err := data.CreateVaildate(); err != nil {
		return nil, err
	}

	uid := xid.New()
	data.UserID = uid.String()
	pass := utils.HashPasswordMD5(data.UserPassword)
	data.UserPassword = pass

	result, err := data.Create()

	if err != nil {
		return nil, err
		
	}
	result.UserPassword = ""
	
	return result, nil
} 

func  (s *usersService) GetUser(useremail string,userpass string) (*repository.User, error) {

	pass := utils.HashPasswordMD5(userpass)
	result := &repository.User{UserEmail: useremail,UserPassword:pass}
	if err := result.GetUser(); err != nil {
		return nil, err
	}
	result.UserPassword = ""
	return result, nil
}

func (s *usersService) GetDelete(useremail string,userpass string) (error) {

	pass := utils.HashPasswordMD5(userpass)
	result := &repository.User{UserEmail: useremail,UserPassword:pass}

	if err := result.Delete(); err != nil {
		return  err
	}
	return nil

}

func (s *usersService) UpdateUser(data *repository.User)(*repository.User, error) {
	
	if err := data.UpdateVaildate(); err != nil {
		return nil, err
	}

	pass := utils.HashPasswordMD5(data.UserPassword)
	data.UserPassword = pass
	result, err := data.Update()

	if err != nil {
		return nil, err
		
	}
	return result, nil

}
