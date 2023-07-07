# Backend_in_go
this is my first time trying to make a backend in go

# infastrucutre 
this will contain two parts

## database
mainly using sqlite as a database and GORM as an ORM

Tables:
- usr Table:
    - ID
    - UserName
    - Password
    - Email
    - DateCreated
- shop Table:
    - ID
    - OwnerID "aka user"
    - Name
    - Category
    - IsPrivet
    - DateCreated
- Products Table:
    - ID
    - ShopID
    - Name
    - Descriptions
    - Price
    - Detail "JSON type"
- order Table:
    - ID
    - ShopID
    - OrdererID
    - ProductID
    - DateCreated

## API
using gin to make http request available

there will be 5 Requests possible
- "/signup": create user and response with an JWT token
- "/login": check for the user and response with an JWT token
- "/shop/": an endpoint where users can: create, delete, list or edit their shops
    - It will contain these Requests:
        - create: " /shop/create "
        - getall: " /shop/getall "
        - delete: " /shop/delete "
        - edit:   " /shop/edit "
        - access a specific shop: " /shop/'shopId' "
- "/browse/": an endpoint where user can: view public shops and buy products
    - It will contain these Requests:
        - get the most recent ,aka the top 10: " /browse/recents "
        - get by filter: " /browse/filter " + a payload with data of the filter
