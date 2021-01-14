package users

import (
	"fmt"
	"github.com/keumda/utils/date"
	"github.com/keumda/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

/*
func something()  {
	user := User{

	}
	if err:= user.Get(); err != nil {
		fmt.Print(err)
		return
	}
}
*/

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError("user is not found")
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

// save user into the DB
func (user *User) Save() *errors.RestErr {
	curr := usersDB[user.Id]
	if curr != nil {
		if curr.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("user email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}
	user.DateCreated = date.GetNowString()
	usersDB[user.Id] = user
	return nil
}
