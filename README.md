# Magic Go API

Magic Go API is secured by the Magic Admin SDK for Go.

This Go server is where all of the Magic Go API requests are handled. Once the user has generated a [Decentralised ID Token (DIDT)](https://docs.magic.link/decentralized-id) from the [client side](https://github.com/shahbaz17/frontend-go-api), they can pass it into their Request Header as a Bearer token to hit protected endpoints.

# API Endpoints

- Unprotected Home Endpoint: http://localhost:9000/
- Protected Endpoint: http://localhost:8080/protected

# Dependencies

- [gorilla/mux](https://github.com/gorilla/mux): Lets us build a powerful HTTP router.
- [magic-admin-go/client](https://docs.magic.link/admin-sdk/go/get-started#creating-an-sdk-client-instance): Lets us instantiate the Magic SDK for Go.
- [magic-admin-go/token](https://docs.magic.link/admin-sdk/go/get-started#creating-a-token-instance): Lets us create a Token instance.

# Quickstart

## Magic Setup

1. Sign up for an account on [Magic](https://magic.link/).
2. Create an app.
3. Copy your app's Test Secret Key (you'll need it soon).

## Server Setup

1. `git clone https://github.com/shahbaz17/magic-go-api`
2. `cd magic-go-api`
3. `mv .env.example .env`
4. Paste the Test Secret Key you just copied as the value for `MAGIC_TEST_SECRET_KEY` and `PORT` in .env:
   ```
   MAGIC_TEST_SECRET_KEY=sk_test_XXXXXXXXXX
   PORT=
   ```
5. Run all .go files with `go run .`

## Test with Postman

1. Generate a DID token on the Client side.

   Click [here](https://github.com/shahbaz17/frontend-go-api) to spin up your own local client and generate the DID token there.

2. Pass the DID token as a Bearer token into the Postman Collection’s HTTP Authorization request header and click **save**.
3. Send your requests to the Magic Go API! 🎉
