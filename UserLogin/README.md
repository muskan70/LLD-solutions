## User Login server

- this app is created using express framework and postgres database
- database details -> user: 'postgres', host: 'localhost', database: 'postgres', password: 'postgres', port: 5432,
- to run application
> npm install
> npm run dev

### API documentation

**POST 'http://localhost:5000/api/user/register'
Body : 
{
    "email":"muskanmangla70@gmail.com",
    "name":"muskan",
    "password":"XXXXXXXX",
    "organization":"manan"
}**


**POST 'http://localhost:5000/api/user/login'
Body :
{
    "email":"muskanmangla70@gmail.com",
    "password":"hello76"
}**


**GET 'http://localhost:5000/status'**

**GET 'http://localhost:5000/dashboard' \
--header 'Content-Type: application/json' \
--header 'Authorization: ••••••' \
--data '{
   "id":2
}'**


