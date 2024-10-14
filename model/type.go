package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID         	primitive.ObjectID 	`bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	Username   	string             	`bson:"username,omitempty" json:"username,omitempty" example:"user"`
	Password   	string             	`bson:"password,omitempty" json:"password,omitempty" example:"user"`
	IDrole		primitive.ObjectID	`bson:"idrole,omitempty" json:"idrole,omitempty" example:"123456789"`
	Email      	string             	`bson:"email,omitempty" json:"email,omitempty" example:"user"`
	Phone      	string             	`bson:"phone,omitempty" json:"phone,omitempty" example:"08123456789"`
	Address    	string             	`bson:"address,omitempty" json:"address,omitempty" example:"Jl. Jalan"`
}


type Roles struct {
	IDrole				primitive.ObjectID  `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	Role				string				`bson:"role,omitempty" json:"role,omitempty" example:"admin"`
}

type JWTClaims struct {
	jwt.StandardClaims
	UserID			primitive.ObjectID `json:"user_id"`
	RoleID			primitive.ObjectID `json:"role_id"`
}