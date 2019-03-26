package main

import (
	"strings"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//a struct to rep user address
/*type Address struct {
	gorm.Model
	BuildingNumber int `gorm:"not null";json:"building_number"`
	Street string `gorm:"not null";json:"street"`
	Town string `gorm:"not null";json:"town"`
	City string `gorm:"not null";json:"city";`
	Country string `gorm:"not null";json:"country"`
}*/

//a struct to rep user account
type Account struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null;json:"name"`
        Email string `gorm:"unique";json:"email"`
	Phone string `gorm:"unique";json:"phone"`
	Address string `json:"address"`
	//FullAddress Address  //relation to Address model
	//FullAddressID int
	AuthType string `json:"authtype";gorm:"-"` // this is not a model field, it'll be ignored by gorm
	Password string
}


//Validate incoming user details...
func (account *Account) Validate() (map[string] interface{}, bool) {

	if account.Name == "" {
                return Message(false, "Name is required"), false
        }

	if account.Email != "" {
                if !strings.Contains(account.Email, "@") || !strings.Contains(account.Email, ".com") {
                        return Message(false, "invalid Email"), false
                }

		temp := &Account{}
	        //check for errors and duplicate emails
		err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
		if err != nil && err != gorm.ErrRecordNotFound {
		        return Message(false, "Connection error. Please retry"), false
		}
	        if temp.Email != "" {
			return Message(false, "Email address already in use by another user."), false
		}
        }
	if account.Phone != "" {
                temp := &Account{}
                //check for errors and duplicate emails
                err := GetDB().Table("accounts").Where("phone = ?", account.Phone).First(temp).Error
                if err != nil && err != gorm.ErrRecordNotFound {
                        return Message(false, "Connection error. Please retry"), false
                }
		if temp.Phone != "" {
                        return Message(false, "Phone number already in use by another user."), false
                }
        } else if account.Email == "" && account.Phone == "" {
		return Message(false, "Email or Phone number is required"), false
	}

	var authType = strings.ToLower(account.AuthType)

	if authType == "" {
		return Message(false, "Authentication Type is required"), false
	} else if authType != "password" && authType != "fingerprint" && authType != "eye-detection" {
		return Message(false, "Please enter a valid Authentication type."), false
	}

	if authType == "fingerprint" || authType == "eye-detection" {
		return Message(false, "the finger print and eye detection Authentication has not yet implemented"), false
	}

	if account.Password == "" {
		return Message(false, "Password is required"), false
	}

	if len(account.Password) < 6 {
                return Message(false, "Password is weak"), false
        }

	return Message(false, "Requirement passed"), true
}

func (account *Account) Create() (map[string] interface{}) {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)

	if account.ID <= 0 {
		return Message(false, "Failed to create account, connection error.")
	}

	account.Password = "" //will not return the password

	response := Message(true, "Account has been created")
	response["account"] = account
	return response
}

