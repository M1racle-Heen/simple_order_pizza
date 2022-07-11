# Pizza_Order

## Pizza_Order service
This project is a simple order pizza. It will provide APIs for the front end to do the following:

1. Create and manage customer accounts composed of personal data.
2. Customers makes Order and Pay for them. The manager gets the order and saves it. The manager sends order details to the kitchen and generates bills.
3. Customers get the final product and pay for it. The manager updates the status of the order.

## DB diagram

![Pizza](https://user-images.githubusercontent.com/70756496/178199802-c3ec1415-d941-46ea-b15e-1fd35265c6b0.png)

DB part was written using Postgres and run through a docker container.

## How to get start
Clone the project to the server directory:

`https://github.com/M1racle-Heen/simple_order_pizza.git`

### You need to install

- [Docker desktop](https://www.docker.com/products/docker-desktop)
- [Golang](https://golang.org/)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

    ```bash

    brew install golang-migrate
    ```
    
- [Sqlc](https://github.com/kyleconroy/sqlc#installation)

    ```bash
    brew install sqlc
    ```

- [Gomock](https://github.com/golang/mock)

    ``` bash
    go install github.com/golang/mock/mockgen@v1.6.0
    ```
## Setup details

- Start postgres container:
    
    ``` bash
    make postgres    
    ```
    
- Start simple_bank database:
    
    ``` bash
    make createdb    
    ```
    
- Run db migration up all versions:
    
    ``` bash
    make migrateup    
    ```
    
- Run db migration up 1 version:
    
    ``` bash
    make migrateup1    
    ```

- Run db migration down all versions:
    
    ``` bash
    make migratedown   
    ```
    
- Run db migration down 1 version:
    
    ``` bash
    make migratedown1    
    ```
## How to run
- Run server:
    
    ``` bash
    make server
    ```
- Run test:

    ``` bash
    make test
    ```
## Hot to generate code
- Generate SQL CRUD with sqlc:
    
    ``` bash
    make sqlc
    ```
- Generate DB mock with gomock:
    
    ``` bash
    make mock
    ```
## How to test

You can download the great tool [Postman](https://www.postman.com/).
To test methods from App
