# Merchant API Webhook Examples

This repo contains scripts and a few web server implementations in different languages to help you implement your own web server to receive webhook events from iFood.

## Scripts

### Message signature generation

We have a bash script to help you validate if the signature your program is generating is valid. You can use the json string directly, remembering the [importance of formatting when generating the signature](https://developer.ifood.com.br/pt-BR/docs/guides/order/events/delivery-methods/webhook/signature/#exemplos).

```bash
sign.sh -s dummysecret -m '{
    "code":"KEEPALIVE",
    "fullCode":"KEEPALIVE",
    "id":"a38ba215-f949-4b2c-982a-0582a9d0c10e"
}'
```

This script was validated in ubuntu and macos.

## Web servers

All implementations expose a `POST http://localhost:8080/webook` endpoint expecting the values sent by iFood's webhook.

### Receiving events from iFood in local environment

Since the program is running in your local machine, you need a way to expose your machine to the internet. We recommend using [ngrok](https://betterprogramming.pub/ngrok-make-your-localhost-accessible-to-anyone-333b99e44b07) to expose your local webhook endpoint to the internet.

You can have one free static DNS so you don't have to update the webhook configuration with the new URL every time you use ngrok. You need to create this domain using ngrok UI before using it.

Example running ngrok after installing it and setting up a static domain:

```sh
ngrok http --domain=your-specific-free-domain.ngrok-free.app 8080
```

### Javascript

As with any node program, run `npm install` to get all dependencies before running `SECRET=yoursecret npm start`.

### Golang

There are no dependencies, so just run `SECRET=yoursecret go run .`.
