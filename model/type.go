package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	IdUser      primitive.ObjectID 	`bson:"id_user,omitempty" json:"id_user,omitempty" example:"123456789"`
	Username   	string             	`bson:"username,omitempty" json:"username,omitempty" example:"username"`
	Password   	string             	`bson:"password,omitempty" json:"password,omitempty" example:"password"`
	IdRole		primitive.ObjectID	`bson:"id_role,omitempty" json:"id_role,omitempty" example:"123456789"`
	Email      	string             	`bson:"email,omitempty" json:"email,omitempty" example:"user@mail.com"`
	Phone      	string             	`bson:"phone,omitempty" json:"phone,omitempty" example:"08123456789"`
	Address    	string             	`bson:"address,omitempty" json:"address,omitempty" example:"Jl. Jalan"`
}

type Roles struct {
	IdRole				primitive.ObjectID  `bson:"id_role,omitempty" json:"id_role,omitempty" example:"123456789"`
	RoleName			string				`bson:"role_name,omitempty" json:"role_name,omitempty" example:"pembeli"`
}

type JWTClaims struct {
	jwt.StandardClaims
	IdUser		primitive.ObjectID `json:"id_user"`
	IdRole		primitive.ObjectID `json:"id_role"`
}