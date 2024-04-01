#todo app

## Description

This one of my first complete projects. It is a single-paged todo app. I decided to develop this application to learn more about software development and show my coding skills.

## Usage

The application can be used to write down todos. Firstly you must create an account, confirm your email address and then you can use the application

## Configuration

The applications requires a config.json file in the application folder. The config.json should have following structure. Additionally there has to be a mongodb running on port 27017 on the local machine.

```json
{
  USERNAME: <username for the mongodb>,
  PASSWORD: <password for the database>,
  JWT_KEY: <key used to encrypt the JWTs>
}
```

## Architecture

The application runs in a single-paged svelte forntend. The backend is written in go. The backend uses APIs to communicate with the frontend.

### API documentation

#### GET request

Provides all html, css and js files.

##### /

Sends the frontend application

#### POST requests

##### /login

Used to login and get authentication token.
Blueprint for request body

```json
{
  "Password": <your password>,
  "Email": <your email>,
}
```

If the login succeeds you will get a response with the following blueprint. Each token is valid for 24 hours.

```json
{
  "Auth": <if authentication was successfull>,
  "Token": <token used to authenticate>
}
```

##### /register

Creates an account with the given email and password.
Each request body must follow the given template

```json
{
  "Email": <email>,
  "Password": <password for email>,
  "ConfirmedPassword": <retyped password>,
  "FirstName": <your first name>,
  "LastName": <your last name>,
  "NotificationsGranted": <bool> <if yes your browser will get notifications (mainly used for frontend)>
}
```

After the register request, you must confirm your email by clicking on the link by the email you received.

##### /addtodo

Adds a todo to the users todo list.
Request body

```json
{
  "Token": <auth token>,
  "New": {
    "NotificationsEnabled": <bool>,
    "Title": <string>,
    "Description": <string>,
    "AlertTime": <time in unix time>,
    "Date": <date when finished in unix time>,
    "Status": <in progess or done>

} <new todo>
}
```

Response body

```json
{
  "Success": <if operation was successfull>,
  "Note": <string>
}
```

#### PUT requests

##### /update todo

Updates a todo
Request body

```json
{
  "Token": <string>
  "Old": {
    "NotificationsEnabled": <bool>,
    "Title": <string>,
    "Description": <string>,
    "AlertTime": <time in unix time>,
    "Date": <date when finished in unix time>,
    "Status": <in progess or done>

} <old todo>,
"New": {
  "NotificationsEnabled": <bool>,
  "Title": <string>,
  "Description": <string>,
  "AlertTime": <time in unix time>,
  "Date": <date when finished in unix time>,
  "Status": <in progess or done>

}<new todo>
}
```

Response body

```json
{
  "Success": <if operation was successfull>,
  "Note": <string>
}
```
