# travelly-api
API for travelly app(app, that helps people to organize their travel plan)

# Authorization

All methods returning two tokens `access_token` and `refresh_token`

JSON with tokens example:
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIzMzM0NTc3ODkwIiwibmFtZSI6IkdvZ2kiLCJpYXQiOjM1MTYyMzkwMjJ9.pZm2pmR7FoyI0hwfSF_OMuE7tD3MVqeN6-D2UuVSYnQ.eyJzdWIiOiIxMzM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9yaWFuIiwiaWF0IjoxNTE2MjM5MDIyfQ.UCSQHuC44ByGLwA7F5gcYea2rruRlbH6_kXuVv7_6Rg"
}
```

### GET 
- `api/auth/` - use it to refresh access token 
  
input: http header "Authorization" with two header parts: Bearer and refresh token

### POST
- `api/auth/email_register` - registration via email method
  
input: JSON with email, password, first name, last name and optionally photo url

Example with photos url:
```json
{
  "email": "qwerty@gmail.com",
  "password": "awesome_password",
  "first_name": "John",
  "last_name": "Dorian",
  "photo_url": "http://www.google.com/url?sa=i&url=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FWeb%2FHTML%2FElement%2Fimg&psig=AOvVaw2Sx5WUbOxNOXkF4Px38IOk&ust=1616444812929000&source=images&cd=vfe&ved=0CAIQjRxqFwoTCKilgqmcwu8CFQAAAAAdAAAAABAI"
}
```

Example without photos url:
```json
{
  "email": "qwerty@gmail.com",
  "password": "awesome_password",
  "first_name": "John",
  "last_name": "Dorian"
}
```

- `api/auth/login` - login existing user method

input: JSON with email and password

Example:
```json
{
  "email": "qwerty@gmail.com",
  "password": "awesome_password"
}
```