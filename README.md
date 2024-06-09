
# Go: the CRM Backend.

CRM Backend Project Overview

Backend (i.e., server-side portion) of a CRM application. That is, you are not responsible for creating a front-end UI. Instead, while your server is on and active, users will be able to make their own HTTP requests to your server to perform CRUD operations.


## Environment Setup

To run this project, you will need to add the following environment variables to your .env file

`PostBird`: For interact with Postgres Database.

`Postgres`: 16.3.1.. Download and setup here: https://www.postgresql.org/docs/release/16.3/

Or
Download this window version for similar to my setup:
https://s.net.vn/32M8

Note: you have to set up by yourself, if you already has Postgres, you can skip this step.

## Database Setup
-- Manually create database with name: 'go_crm'

-- Connect to the database by postbird..

-- Create the Customer table, import ***init.sql*** script to create database..


## Run Locally

Clone the project

```bash
  git clone https://github.com/hoanghaing/go-crm-backend
```

Go to the project directory

```bash
  cd go-crm
```

Start the server

```bash
  go run main.go
```


## Authors

- [@hainh](https://www.linkedin.com/in/hainhptit/)


## Features

- Getting a single customer through a **/customers/{id}**
- Getting all customers through a the **/customers**
- Creating a customer through a **/customers**
- Updating a customer through a **/customers/{id}**
- Deleting a customer through a **/customers/{id}**


## Warning
- You maybe have to change: 
dsn := "host=localhost user=postgres ..."

In **main.go** file to your own host, user, and password base on your setup