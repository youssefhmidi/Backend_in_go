# Backend_in_go
this is my first time trying to make a backend in go

# infastrucutre 
this will contain two parts

## database
mainly using sqlite as a database and GORM as an ORM

Tables:
- usr Table: ✅
    - ID
    - UserName
    - Password
    - Email
    - DateCreated
- shop Table: ✅
    - ID
    - OwnerID "aka user"
    - Name
    - Category
    - IsPrivet
    - DateCreated
- product Table: ✅
    - ID
    - ShopID
    - Name
    - Descriptions
    - Price
    - Detail "JSON type"
- order Table: ✅
    - ID
    - ShopID
    - OrdererID
    - DateCreated
- productorder Table: "junction table for the 'Many to Many' relation betwwen product and order ✅
    - ID
    - Product ID
    - Order ID

## API
using gin to make http request available

there will be 5 Requests possible
- "/signup": create user and response with an JWT token ✅
- "/login": check for the user and response with an JWT token ✅
- "/me" : returns the current user informations ✅
- "/shop/": an endpoint where users can: create, delete, list or edit their shops
    - It will contain these Requests:
        - create: " /shop/create "✅
        - getall: " /shop/getall "✅
        - delete: " /shop/delete "✅
        - edit:   " /shop/edit "✅
        - access a specific shop: " /shop/get?name='Shop name' " or " /shop/get?id='Shop id' " ✅
        - add products: "/shop/product" `method POST`✅
        - get products  "/shop/product" `method GET`✅
        - post an order "/shop/order" `method POST` ✅
        - get shop's orders and their status "/shop/order" `method GET`✅
- "/browse/": an endpoint where user can: view public shops and buy products
    - It will contain these Requests:
        - get the most recent ,aka the top 10: " /browse " ✅
        - get by filter: " /browse?filter=... " + a payload with data of the filter ✅
    Note : the filter is just an expression like "category=toys" or "name=somthing" in futur update I will make a complex filter type
