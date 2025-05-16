## User Login server

### Requirements
- Create a small REST service in Nodejs with Database – Postgresql
- Create a table with the names users : Column names are Name, Password Email, Organization
- You need to create Four endpoints:
1. Health check endpoint without authentication that returns code 200 OK when service is loaded
2. User Signup, Signin, and Dashboard endpoint
> Fields for Signup – Name, Email, Password, Organization<br>
> Fields for Signin - Email and password
- Please create a web page that uses this REST service to show the user signup and sign-in page. After successfully logging in, the user should be redirected to a new page

### Implementation Details
- this app is created using express framework and postgres database
- database details -> user: 'postgres', host: 'localhost', database: 'postgres', password: 'postgres', port: 5432,
- to run application
> npm install<br>
> npm run dev

### API documentation

1. POST 'http://localhost:5000/api/user/register' \
-- data '{
    "email":"muskanmangla70@gmail.com",
    "name":"muskan",
    "password":"XXXXXXXX",
    "organization":"manan"
}'


2. POST 'http://localhost:5000/api/user/login' \
-- data '{
    "email":"muskanmangla70@gmail.com",
    "password":"XXXXXXXX"
}'


3. GET 'http://localhost:5000/status'

4. GET 'http://localhost:5000/dashboard' \
--header 'Content-Type: application/json' \
--header 'Authorization: ••••••' \
--data '{
   "id":2
}'


