# Ledger App with Go[echo-gorm] & Postgres
There are two roles as user and admin.

### user:
- can create an account
- can check balance
- can deposit money into self account
- can withdraw money from self account
- can view past transactions
- can transfer money to another user

### admin:
- can give admin role to users
- can delete users
- can update users' information
- can see the user's balance at spesific time
- can get all users
- can see the balances of all users

### details:
- The user cannot send money does not have. Balance cannot be negative.
- The user can only have one account and balance with the same email.
- If the user information is updated, a unique email should be used. Each email can have only one account.
- If the users' role is not an admin, most of the services will not used.
- If the user account is deleted, it will be deleted in the wallet with balance.
- When creating an account, the password must be a minimum of 8 characters and contain both letters and numbers at the same time.
- Except for creating a new account, all other services are authorized via the header.
- If a transaction is interrupted, a rollback is performed.
- All transactions are recorded. Any user's balance value from a time in the past can be retrieved this way.
- There are three different tables in total. [Transactions, Users, Wallets]
- If database has not any "admin" role, created one whit this informations
```
- Username: "admin",
- Email:    "admin@admin.com",
- Password: "admin1234",
- Role:     "admin"
```
## POST Methods
### http://localhost:8080/app/services/users/create
```
curl --location 'http://localhost:8080/app/services/users/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "test@test.com",
    "password": "123456789test",
    "username": "Test Test"
}'
```

### http://localhost:8080/app/services/wallets/update
```
curl --location 'http://localhost:8080/app/services/wallets/update' \
--header 'email: test@test.com' \
--header 'password: 123456789test' \
--header 'Content-Type: application/json' \
--data '{
    "amount": 300
}'
```

### http://localhost:8080/app/services/transactions/create
```
curl --location 'http://localhost:8080/app/services/transactions/create' \
--header 'email: test@test.com' \
--header 'password: 123456789test' \
--header 'Content-Type: application/json' \
--data-raw '{
    "recipient_email": "testv2@test.com",
    "amount": 100
}'
```

### http://localhost:8080/app/services/users/update-user
```
curl --location 'http://localhost:8080/app/services/users/update-user' \
--header 'email: test@test.com' \
--header 'password: 123456789test' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id": "2",
    "user": {
        "username": "testv3 test",
        "email": "testv3@test.com",
        "password": "123456789test"
    }
}'
```

### http://localhost:8080/app/services/users/delete-user
```
curl --location 'http://localhost:8080/app/services/users/delete-user' \
--header 'email: test@test.com' \
--header 'password: 123456789test' \
--header 'Content-Type: application/json' \
--data '{
    "user_id": "2"
}'
```

### http://localhost:8080/app/services/users/update-user-role
```
curl --location 'http://localhost:8080/app/services/users/update-user-role' \
--header 'email: test@test.com' \
--header 'password: 123456789test' \
--header 'Content-Type: application/json' \
--data '{
    "user_id": "2",
    "role": "admin"
}'
```

### http://localhost:8080/app/services/wallets/update-wallet-balance-by-user-id/
```
curl --location 'http://localhost:8080/app/services/wallets/update-wallet-balance-by-user-id' \
--header 'email: test@test.com' \
--header 'password: 123456789test' \
--header 'Content-Type: application/json' \
--data '{
    "user_id": "2",
    "amount": -250
}'
```

### http://localhost:8080/app/services/transactions/get-wallet-balance-at-time
```
curl --location 'http://localhost:8080/app/services/transactions/get-wallet-balance-at-time' \
--header 'email: test@test.com' \
--header 'password: 123456789test' \
--header 'Content-Type: application/json' \
--data '{
    "user_id": "2",
    "time": "2023-06-25T03:00:00Z"
}'
```

## GET Methods
### http://localhost:8080/app/services/wallets/get
```
curl --location 'http://localhost:8080/app/services/wallets/get' \
--header 'email: test@test.com' \
--header 'password: 123456789test'
```

### http://localhost:8080/app/services/transactions/get
```
curl --location 'http://localhost:8080/app/services/transactions/get' \
--header 'email: test@test.com' \
--header 'password: 123456789test'
```

### http://localhost:8080/app/services/users/get-all-users
```
curl --location 'http://localhost:8080/app/services/users/get-all-users' \
--header 'email: test@test.com' \
--header 'password: 123456789test'
```

### http://localhost:8080/app/services/users/get-all-users-with-wallet
```
curl --location 'http://localhost:8080/app/services/users/get-all-users-with-wallet' \
--header 'email: test@test.com' \
--header 'password: 123456789test'
```

### http://localhost:8080/app/services/wallets/get-all-wallets
```
curl --location 'http://localhost:8080/app/services/wallets/get-all-wallets' \
--header 'email: test@test.com' \
--header 'password: 123456789test'
```

### http://localhost:8080/app/services/users/get-user-by-id/{id}
```
curl --location 'http://localhost:8080/app/services/users/get-user-by-id/5' \
--header 'email: test@test.com' \
--header 'password: 123456789test'
```

### http://localhost:8080/app/services/wallets/get-wallet-by-id/{id}
```
curl --location 'http://localhost:8080/app/services/wallets/get-wallet-by-id/3' \
--header 'email: test@test.com' \
--header 'password: 123456789test'
```
