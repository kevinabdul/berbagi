<p align="center" width="100%">
  <a href="https://github.com/kevinabdul/berbagi"><img width="70%" src="https://github.com/kevinabdul/berbagi/blob/bd62cea397eeed4a28546db6d01dc954f939d783/res/logo.png"></a>
</p>

# Berbagi

Find happiness by sharing with others

[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=blue)](https://pkg.go.dev/gorm.io/gorm?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=blue)](https://github.com/labstack/echo)
[![Codecov](https://img.shields.io/badge/coverage-90.7-blue?)](https://github.com/kevinabdul/berbagi)

# Table of Content

- [Description](#description)
- [Feature Overview](#feature-overview)
- [High Level Architecture](#high-level-architecture)
- [Flowchart](#flowchart)
- [Entity Relations Diagram](#entity-relations-diagram)
- [How to use](#how-to-use)
- [Endpoints](#endpoints)
- [OpenAPI Documentation](#openapi-documentation)
- [Contribute](#contribute)
- [Credits](#credits)

# Description
--

# Feature Overview
<img src="https://github.com/urnikrokhiyah/pictures/blob/e41900307780344cca1ce634154608779f6b9497/Untitled%20Diagram(3).drawio(2).png">

# High Level Architecture
<img src="https://github.com/kevinabdul/berbagi/blob/bd62cea397eeed4a28546db6d01dc954f939d783/res/HLA-Berbagi.jpg">


# Entity Relations Diagram
<img src="https://github.com/kevinabdul/berbagi/blob/bd62cea397eeed4a28546db6d01dc954f939d783/res/erd.drawio.png">

# How to use
- Install Go and MySQL
- Clone this repository in your $PATH:
```
$ git clone https://github.com/kevinabdul/berbagi.git
```
- Run `main.go`
```
$ go run main.go
```

# Endpoints

| Method | Endpoint | Description| Authentication | Authorization
|:-----|:--------|:----------| :----------:| :----------:|
| POST  | /register | Register a new user | No | No
| POST | /login | Login existing user| No | No
|---|---|---|---|---|
| GET    | /proficiencies |Get list of all proficiencies | Yes | Yes
| POST | /proficiencies | Register a new proficiencies| Yes | Yes
| PUT   | /proficiencies | Update existing proficiencies | Yes | Yes
| DELETE| /proficiencies | Delete existing proficiencies | Yes | Yes
|---|---|---|---|---|
| GET   | /products | Get products list (query `categoryId` to sort by category) | No | No
|---|---|---|---|---|
| GET | /product-carts | Get User's Cart. Target cart based on "userId" claims in jwt | Yes | Yes
| PUT | /product-carts | Update User's Cart . Target cart based on "userId" claims in jwt | Yes | Yes
| DELETE | /product-carts | Delete User's Cart . Target cart based on "userId" claims in jwt | Yes | Yes
|---|---|---|---|---|
| GET | /volunteers | Get volunteers list by admin | Yes | Yes
| GET | /volunteers/profile | Get volunteers list by admin | Yes | No
|---|---|---|---|---|
| GET | /services | Get appointed services in cart by volunteer | Yes | No
| POST | /services | Add services to cart by volunteer | Yes | No
| UPDATE | /services | Update services in cart by volunteer | Yes | No
| DELETE | /services | Delete services to cart by volunteer | Yes | No
|---|---|---|---|---|
| POST | /services/verification | Confirm done service by volunteer | Yes | No
| GET | /services/verification/:verificationId | Get service confirmation | Yes | No
| GET | /services/display/:verificationId | Show service confirmation certificate | Yes | No
|---|---|---|---|---|
| GET | /nearby/:resource | Show nearby `recipients`/`requests`; query `type` and `range` to sort | Yes | No
|---|---|---|---|---|
| GET | /checkout | Get checked out User's Cart based on jwt's userId claims. | Yes | No
| POST | /checkout | Check out User's Cart based on jwt's userId claims. | Yes | No
|---|---|---|---|---|
| GET | /completion/:verificationId | Get volunteer's completion details | Yes | Yes
| PUT | /completion/:verificationId | Update volunteer's completion status | Yes | Yes
|---|---|---|---|---|
| GET | /certificates/:completionId | Get volunteer's certificate | Yes | Yes
| GET | /certificates/display/:completionId | Show volunteer's certificate | Yes | Yes
|---|---|---|---|---|
| GET | /donation | Get donor donations list; query `resolved` to sort | Yes | Yes
| POST | /donation | Post donations; query `quick` to checkout immediately | Yes | Yes
| GET | /cart/donation | Get donations list in cart | Yes | Yes
| PUT | /cart/donation | Update non-requested donations in cart | Yes | Yes
| DELETE | /cart/donation | Delete donations in cart | Yes | Yes
| POST | /donation/checkout | Checkout one donation from cart | Yes | Yes
|---|---|---|---|---|
| GET | /request | Get requests list; query `resolved` to sort | Yes | Yes
| GET | /request/:field | Get type-specific requests list; query `resolved` to sort | Yes | Yes
| POST | /request/gift | Make gift request | Yes | Yes
| POST | /request/donation | Make donation request | Yes | Yes
| POST | /request/service | Make service request | Yes | Yes
| DELETE | /request/:request_id | Delete request | Yes | Yes
|---|---|---|---|---|
| GET | /payments | Get all pending payment. Return value depends on "userId" claims in jwt | Yes | No
| POST | /payments | Resolves one pending gift payment | Yes | No
| POST | /payments/donation | Resolves one pending donation payment | Yes | No
|---|---|---|---|---|
| GET | /gifts | Get all gifts requested by children; query `status` to sort | Yes | No
|---|---|---|---|---|

<br>


# OpenAPI Documentation
Go to this page to test the api:
[Berbagi-API docs](https://app.swaggerhub.com/apis/arieshta/berbagi-api/1.0.0)
<br>
<br>

## Contribute

**Use issues for everything**

- Contribute by:
  - Reporting issues
  - Creating pull requests
  - Suggesting new features or enhancements
  - Improve or fix documentation
- PR should have:
  - Test case
  - Documentation
  - Example

## Credits

- [Ilham Aris](https://github.com/arieshta) (Author and maintainer)
- [Kevin Abdul](https://github.com/kevinabdul) (Author and maintainer)
- [Urnik Rokhiyah](https://github.com/urnikrokhiyah) (Author and maintainer)
