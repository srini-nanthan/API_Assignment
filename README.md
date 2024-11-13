Steps to run

1. `docker-compose up`

## Endpoints

**REGISTER API** -
Creates a new user account using an email and password
```bash
curl --location 'http://localhost:8080/register' \
--header 'Content-Type: application/json' \
--data-raw '{"email":"sample@gmail.com", "password":"Admin@123"}'
```
-------------------------------------------------------------------

**LOGIN API** -
Authenticates a user's credentials and, if valid, returns an access token (JWT) and a refresh token for secure session management

curl --location 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{"email":"sample@gmail.com", "password":"Admin@123"}'

-------------------------------------------------------------------

**PROTECTED API** -
Protected API is a secure endpoint that only allows access if you have a valid token, blocking anyone without proper authorization

curl --location 'http://localhost:8080/protected' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNhbXBsZUBnbWFpbC5jb20iLCJleHAiOjE3MzE0OTcyOTZ9.I-8SxJicrari88lOJlQ_PlBGzU1XO0yz6Y_a0MBtJ2s'

-------------------------------------------------------------------

**REVOKE API** -
Revokes the specified token, making it invalid for future requests

curl --location --request POST 'http://localhost:8080/revoke' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNhbXBsZUBnbWFpbC5jb20iLCJleHAiOjE3MzE0OTcyOTZ9.I-8SxJicrari88lOJlQ_PlBGzU1XO0yz6Y_a0MBtJ2s'
