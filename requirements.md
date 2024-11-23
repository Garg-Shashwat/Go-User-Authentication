## Requirements

1. Sign up (creation of user) using email and password
2. Sign in
    - Authentication of user credentials
    - A token is returned as response preferably JWT
3. Authorization of token
    - Mechanism of sending token along with a request from client to service
    - Should check for expiry
    - Error handling (proper error codes in each failure scenario)
4. Revocation of token
    - Mechanism of revoking a token from backend
5. Mechanism to refresh a token
    - Client should be able to renew the token before it expires