## User Login server

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


3. GET 'http://localhost:5000/status'**

4. GET 'http://localhost:5000/dashboard' \
--header 'Content-Type: application/json' \
--header 'Authorization: ••••••' \
--data '{
   "id":2
}'


