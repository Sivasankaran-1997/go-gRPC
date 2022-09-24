package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"grpc_curd/users/server/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	"strings"
	"go.mongodb.org/mongo-driver/mongo/options"
)



type User struct {
	UserID string `bson:"userid"`
	UserName string `bson:"username"`
	UserEmail string `bson:"useremail"`
	UserPassword string `bson:"userpassword"`
	UserPhoneno string `bson:"userphoneno"`
	UserAddress string `bson:"useraddress"`

}

func (user *User) CreateVaildate() error {
	if strings.TrimSpace(user.UserID) == "" {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("User ID is Required"),
		)
	}

	if strings.TrimSpace(user.UserName) == "" {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("User Name is Required"),
		)
	}

	if strings.TrimSpace(user.UserEmail) == "" {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("User Email is Required"),
		)
	}

	if strings.TrimSpace(user.UserPassword) == "" {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("User Password is Required"),
		)
	}

	if strings.TrimSpace(user.UserPhoneno) == "" {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("User Phoneno is Required"),
		)
	}

	if strings.TrimSpace(user.UserAddress) == "" {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("User Address is Required"),
		)
	}
	return nil
}

func (user *User) UpdateVaildate() error {
	
	if strings.TrimSpace(user.UserName) == "" {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("User Name is Required"),
		)
	}

	if strings.TrimSpace(user.UserEmail) == "" {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("User Email is Required"),
		)
	}

	if strings.TrimSpace(user.UserPassword) == "" {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("User Password is Required"),
		)
	}


	if strings.TrimSpace(user.UserPhoneno) == "" {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("User Phoneno is Required"),
		)
	}

	if strings.TrimSpace(user.UserAddress) == "" {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("User Address is Required"),
		)
	}
	return nil
}



var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func (user *User) Create() (*User,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

	emailCount, _ := userCollection.CountDocuments(ctx, bson.M{"useremail": user.UserEmail})
	emailPhoneno, _ := userCollection.CountDocuments(ctx, bson.M{"userphoneno": user.UserPhoneno})
	defer cancel()

	if emailCount > 0  {
		return nil, status.Errorf(
			codes.AlreadyExists,
			fmt.Sprintf("Email ID is Exists"),
		)
	}
	if emailPhoneno > 0 {
		fmt.Println("Phone count",emailPhoneno)
		return nil, status.Errorf(
			codes.AlreadyExists,
			fmt.Sprintf(" Phone No is Exists"),
		)
	}

	_, err := userCollection.InsertOne(ctx, user)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error on create user"),
		)
	}

	return user, nil
}

func (user *User) GetUser() error{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

	filter := bson.M{"useremail": user.UserEmail,"userpassword":user.UserPassword}

	err := userCollection.FindOne(ctx, filter).Decode(&user)

	defer cancel()

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("User Email is Not Exists"),
		)
	}
	return nil
}

func (user *User) Delete() error {


	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

	filter := bson.M{"useremail": user.UserEmail,"userpassword":user.UserPassword}
	result, err := userCollection.DeleteOne(ctx, filter)
	defer cancel()

	if result.DeletedCount == 0 {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("No Record Found"),
		)
	}

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Email Not Found"),
		)
	}

	return  nil

}

func (user *User) Update() (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

	emailCount, _ := userCollection.CountDocuments(ctx, bson.M{"useremail": user.UserEmail,"userpassword" : user.UserPassword})

	if emailCount == 0  {
		return nil, status.Errorf(
			codes.AlreadyExists,
			fmt.Sprintf("Email ID is Not Exists"),
		)
	}

	filter := bson.M{"useremail": user.UserEmail,"userpassword" : user.UserPassword}

	emailPhoneno, _ := userCollection.CountDocuments(ctx, bson.M{"userphoneno": user.UserPhoneno})

	if emailPhoneno > 0 {
		return nil, status.Errorf(
			codes.AlreadyExists,
			fmt.Sprintf(" Phone No is Exists"),
		)
	}

	updateValue := bson.M{"$set": bson.M{"username": user.UserName,"userphoneno" :user.UserPhoneno,"useraddress" : user.UserAddress}}

	opts := options.Update().SetUpsert(true)

	result, err := userCollection.UpdateOne(ctx, filter, updateValue, opts)

	defer cancel()

	if result.ModifiedCount == 0 {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Data Not Modified"),
		)
	}


	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Data Not Updated"),
		)
	}

	return user, nil

}
