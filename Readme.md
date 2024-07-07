# Go User Registration for Online Tournament

## Sign in
```bash
localhost:8080/signin [POST]
{
    "username": "username",
    "password": "password"
}
```
- [x] User can sign up with username and password
- [x] Server side validation

## Sign up
```bash
localhost:8080/signup [POST]
{
    "name": "name",
    "phone": "phone",
    "username": "username",
    "password": "password"
}
```
- [x] User can sign up with Name, Phone Number, Username, Password
- [x] Server side validation

## Register Tournament
```bash
localhost:8080/tournament/register [POST]
{
    "teamName": "teamName",
    "captain": {
        "name": "name",
        "phone": "phone",
        "gender": "male",
    }
    "members": {
          "name": "name",
          "phone": "phone",
          "gender": "female"
    }
}
```
- [x] User can register for a tournament with Team Name, Captain Section including (Name, PhoneNumber, and Gender), and Member Section is same as Captain Section
- [x] Server side validation

## Get User Profile
```bash
localhost:8080/profile [GET]
```
- [x] Get user profile for auto complete the tournament registration form

## Logout
```bash
localhost:8080/logout [POST]
```
- [x] User can logout from the system

## JWT Authentication

## How to use
- Clone the repository
- Run the command `go run main.go`