# JSON Restful JSON API

## requirements for this module

visit this article to install [link](https://tecadmin.net/install-go-on-centos/ "go language and its enviornment") on your linux system

## Run the server by:
it will run on localhost:8080/

```sh
# clone the repo from github
git clone https://github.com/D0nQuixote/accountsModule.git

# change directory into the projecct
cd accountsModule/

# run the combiled module
./server

# you can combile the project using this command
go build -o filename
```
or you can run the source files by doing this:
```sh
go run ./*.go
```

## The API endpoint documentation

### first endpoint is /api/v1/auth/signup

this endpoint is responsible for creating user accounts
it only have one PUT method that will take json data like this

{
	"name": "UserName",
	"email": "userEmail@gmail.com",
	"phone": "0111111111",
	"address": "User address",
	"authtype": "password" or "fingerprint" or "eye-detection", 
	"password": "UserPassword"
}

the authtype implemented the password authentcation for now.
i'm planning to implement the finger print and eye detection by using an image for each of them.

there will be another endpoint for confirming the user's Email address
