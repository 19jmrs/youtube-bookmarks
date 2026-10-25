package types

import (
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 12

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email             string `bson:"email" json:"email"`
	EncryptedPassword string `bson:"EncryptedPassword" json:"-"`
}

type CreateUserParams struct{	
	Email             string `json:"email"`
	Password string `json:"password"`
}



func NewUserFromParams(params CreateUserParams) (*User, error){
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)

	if err != nil{
		return nil,err
	}
	return &User{		
		Email: params.Email,
		EncryptedPassword: string(encpw),
	}, nil
}


func isEmailValid(e string) bool{
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	return emailRegex.MatchString(e)
}
