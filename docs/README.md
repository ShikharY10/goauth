# Documentation

## Setup & Starting API server

[Click here](https://github.com/HousewareHQ/houseware---backend-engineering-octernship-ShikharY10/tree/main/backend/deployments)

<br>

## API

[Click here](https://github.com/HousewareHQ/houseware---backend-engineering-octernship-ShikharY10/tree/main/backend/api)

<br>

# Design Decisions

<h2> Reason why I choose <b><span style="color:blue">GIN</span></b> as web framework</h2>

- Past Experience
- Largest community of any go web framework | Most stars on [Github](https://github.com/gin-gonic/gin)
- Simple design and architecture

<h2> Reason why I have choose <b><span style="color:green">MongoDB</span></b> as primary database</h2>

- Ease of getting started
- Support structured or unstructured data
- Json like format to store documents
- Great documentation support for Golang

<b>I have not used any ORM/ODM because MongoDB provides great document support for golang.</b>

<h2> Reason why I have choose <b><span style="color:red">Redis</span></b> as cache</h2>

- In Memory
- Key/Value based storing
- Great community support

I have also used Redis for caching purposes. I have used cache to store the uniqueness and expiry of refresh and access token. By choosing this style of storing expiry and uniqueness of token we have assured that server have full controll over the lifetime of tokens. Let say user 'A' have send a logout request to server, then what server will do is it will go to cache and delete both key for access and refresh token and this way when user 'A' tries access authorized content using expired/Useless token then server will responed with 401 Status code. Because when any user tries to access authorized content using token, server check the uniqueness by matching the hash part of the token with stored hashed part and this way server assures the uniqueness of token, Redis support key lifetime feature and I have used this feature to store the hash of access token for 1 hour and hash refresh token for 24 hours.
