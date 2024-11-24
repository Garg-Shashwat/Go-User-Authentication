## DB Schema

### Users

|Name       |Type       |
|-----------|-----------|
|id         |int        |
|email      |varchar    |
|password   |varchar    |
|created_at |timestamp  |
|updated_at |timestamp  |


### User_tokens

|Name       |Type       |
|-----------|-----------|
|id         |int        |
|user_id    |int        |
|token      |varchar    |
|expires_at |timestamp  |
|revoked    |boolean    |
|created_at |timestamp  |