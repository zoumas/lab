# Endpoints for the Users resource

## POST /users

Description: Creates a new user

Request Body:
```
{
    "name": "zoumas",
    "password": "zoumas password"
}
```

NOTE: Document the rules about the username and password here

Response Body: None

Response Status Code: 
* 201 if a user is created. 
* 401 if something went wrong when processing the request or the body doesn't match with the rules set for name and passwords. 
* 500 if for some reason the user can't be stored in the database.

## GET /users

Description: Returns all users

Request Body: None

Example Response Body:
```
[
    {
        "id": "8424384b-adb6-4f69-9d21-8a9468dbdaab",
        "created_at": "2023-11-27T13:55:34.106989+02:00",
        "updated_at": "2023-11-27T13:55:34.106989+02:00",
        "name": "zoumas"
    },
    {
        "id": "b377f914-6e07-4e8b-9944-3d9af22f7ab3",
        "created_at": "2023-11-27T14:10:38.666645+02:00",
        "updated_at": "2023-11-27T14:10:38.666645+02:00",
        "name": "doukas"
    }
]
```

Response Status Code:
* 500 if for some reason the users can't be retrieved from the database.
* 200 if everything went well.

## GET /users/{name}

Description: Returns a single user

Request Body: None

Example Request: `GET /users/zoumas`

Response Body:
```
{
    "id": "8424384b-adb6-4f69-9d21-8a9468dbdaab",
    "created_at": "2023-11-27T13:55:34.106989+02:00",
    "updated_at": "2023-11-27T13:55:34.106989+02:00",
    "name": "zoumas"
}
```

