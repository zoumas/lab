# Endpoint for user login

Requires Basic Authorization Header

`Authorization: Basic username:password`

Request Body: None

Response Body: None

Response Status Code:
* 200 - if the username exists and the password matches the password in the database
* 401 - everything else
