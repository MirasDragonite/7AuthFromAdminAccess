## 7AuthFromAdminAccess

* [mkabyken](https://t.me/mirasKabykenov)


### Description:
Test project to mobydev with function login,register user  and creating, updating product.


### Usage
Clone the repository:
```
git clone https://github.com/MirasDragonite/7AuthFromAdminAccess.git
```


####   Run a program:


```
go run ./cmd/main.go
```
To fast access you can use Postman and start working with this URL
```
localhost:8000
```
 

## Documentation

```
POST `localhost:8000/sign-up` to create new user.
{
    "username": "Test",
    "email": "example@email.com",
    "password": "12345678",
    "repeat_password": "12345678"
}

POST `localhost:8000/auth/sign-in` to login into system.
{
    "email": "example@email.com",
    "password": "12345678"
    
}

WARNINING: "Only user with role admin can do the next requests"

POST `localhost:8000/product/create to create new product
{
    "name": "none",
    "category":"none",
    "product_type":"none",
    "year":"none",
    "age_category":"none",
    "chronology":"none",
    "key_words":"none",
    "description":"none",
    "director":"none",
    "producer":"none"
}

POST `localhost:8000/product/update?id='ProductID' to update  product
{
    "name": "new",
    "category":"new",
    "product_type":"new",
    "year":"new",
    "age_category":"new",
    "chronology":"new",
    "key_words":"new",
    "description":"new",
    "director":"new",
    "producer":"new"
}

```
