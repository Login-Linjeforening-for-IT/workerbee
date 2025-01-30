# Beehive Admin API

https://wiki.login.no/en/tekkom/projects/beehive/docs/backend/admin-api

## How to run
1. Create a key pair using `openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 -keyout key.pem -out cert.pem`
2. Create a `.env` file with values for symmetric keys:
```ts
TOKEN_ACCESS_TOKEN_SYMMETRIC_KEY=<your_symmetric_key> 
TOKEN_REFRESH_TOKEN_SYMMETRIC_KEY=<your_symmetric_refresh_key>
```
3. Run `make run`


## Maintainers
- Gjermund H. Pedersen<br>
  gjermuhp@stud.ntnu.no