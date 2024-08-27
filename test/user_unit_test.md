# API Unit Test Results and Behaviours For User Routes

MongoDB User-Collection content before starting user Unit Testing:

```json
[
	{
		"_id": { "$oid": "6411d25577a6afcd1cf9d71e" },
		"name": "Shikhar Yadav",
		"username": "ShikharY10",
		"password": "+WmjGVFGNN3vW8pkegb/dVmmxu3UPgXbxiU100G3kXw=",
		"organisation": "google",
		"role": "admin"
	},
	{
		"_id": { "$oid": "6412283c21135fb279075a1e" },
		"name": "Shivani Yadav",
		"username": "shivani22",
		"password": "+WmjGVFGNN3vW8pkegb/dVmmxu3UPgXbxiU100G3kXw=",
		"organisation": "SSC",
		"role": "user"
	},
	{
		"_id": { "$oid": "64122c71d9d68a1199f78295" },
		"name": "Swatantra Yadav",
		"username": "swat15",
		"password": "+WmjGVFGNN3vW8pkegb/dVmmxu3UPgXbxiU100G3kXw=",
		"organisation": "google",
		"role": "admin"
	},
	{
		"_id": { "$oid": "64122c8bd9d68a1199f78296" },
		"name": "Varsha Yadav",
		"username": "varsha11",
		"password": "+WmjGVFGNN3vW8pkegb/dVmmxu3UPgXbxiU100G3kXw=",
		"organisation": "google",
		"role": "user"
	},
	{
		"_id": { "$oid": "6412ea9a6c0c9b071e7e9395" },
		"name": "Ananya Raman",
		"username": "Ananya22",
		"password": "+WmjGVFGNN3vW8pkegb/dVmmxu3UPgXbxiU100G3kXw=",
		"organisation": "SSC",
		"role": "user"
	}
]
```

## **1. Singup**

Method:

```
POST
```

Path:

```bash
http://127.0.0.1:10222/api/v1/signup
```

Request:

```bash
    curl -X 'POST' \
    'http://127.0.0.1:10222/api/v1/signup' \
    -H 'accept: application/json' \
    -H 'Content-Type: application/json' \
    -d '{
        "name": "Adarsh Yadav",
        "organisation": "SSC",
        "password": "adarsh",
        "username": "adarsh04"
    }'
```

Response:

The response of this request contains json response body and a **http-only cookie** which is a refresh token.

```json
Status Code

200 OK

Response Body

{
  "id": "6412efe42ddfc8927e6cc444",
  "name": "Adarsh Yadav",
  "organisation": "SSC",
  "role": "user",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzg5NjYyNjAsImlkIjoiNjQxMmVmZTQyZGRmYzg5MjdlNmNjNDQ0Iiwicm9sZSI6InVzZXIiLCJ1c2VybmFtZSI6ImFkYXJzaDA0In0.0Ijptq0hxyWnEpQKABhHHewrO7E7V-JkGe4ynU_2Rok",
  "username": "adarsh04"
}

Headers

Set-Cookie: "refresh=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzkwNTExOTIsImlkIjoiNjQxMWQyNTU3N2E2YWZjZDFjZjlkNzFlIn0.IcmQLD0Us0nHOXf6qA2eQgLVYLLEJ8Xvl8P5tLzrPQI; Path=/; Max-Age=86400; HttpOnly"

```

Posssible Responses Status Code:

| Code &nbsp; &nbsp; | Description &nbsp; &nbsp; |
| ------------------ | ------------------------- |
| 200                | Ok                        |
| 400                | Bad Request               |
| 500                | Internal server error     |

<br>
<hr>
<br>

### **2. Login**

Method:

```
POST
```

Path:

```bash
http://127.0.0.1:10222/api/v1/login
```

Request:

```bash
curl -X 'POST' \
  'http://127.0.0.1:10222/api/v1/login' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "password": "1234",
  "username": "ShikharY10"
}'
```

Response:

The response of this request contains json response body and a **http-only cookie** which is a refresh token.

```json
Status Code

200 OK

Response Body

{
  "id": "6411d25577a6afcd1cf9d71e",
  "name": "Shikhar Yadav",
  "organisation": "google",
  "role": "admin",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzg5Njg5NzQsImlkIjoiNjQxMWQyNTU3N2E2YWZjZDFjZjlkNzFlIiwicm9sZSI6ImFkbWluIiwidXNlcm5hbWUiOiJTaGlraGFyWTEwIn0.hBCuBxR8fbeokzUYkbPuNS43RSvyWQCmhQy260Vl_mM",
  "username": "ShikharY10"
}

Header

Set-Cookie: "refresh=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzkwNTExOTIsImlkIjoiNjQxMWQyNTU3N2E2YWZjZDFjZjlkNzFlIn0.IcmQLD0Us0nHOXf6qA2eQgLVYLLEJ8Xvl8P5tLzrPQI; Path=/; Max-Age=86400; HttpOnly"
```

Posssible Responses Status Code:

| Code &nbsp; &nbsp; | Description &nbsp; &nbsp; |
| ------------------ | ------------------------- |
| 200                | Ok                        |
| 400                | Bad Request               |
| 500                | Internal server error     |

<br>
<hr>
<br>

### **3. Refresh Access Token**

Method:

```
PUT
```

Path:

```bash
http://127.0.0.1:10222/api/v1/refresh/<id>
```

Request:

```bash
curl -X 'PUT' \
  'http://127.0.0.1:10222/api/v1/refresh/6411d25577a6afcd1cf9d71e' \
  -H 'accept: application/json' \
  -H 'Cookie: refresh=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzkwNTExOTIsImlkIjoiNjQxMWQyNTU3N2E2YWZjZDFjZjlkNzFlIn0.IcmQLD0Us0nHOXf6qA2eQgLVYLLEJ8Xvl8P5tLzrPQI'
```

Response:

```json
Status Code

200 OK

Response Body

{
  "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzg5NzA0NjksImlkIjoiNjQxMWQyNTU3N2E2YWZjZDFjZjlkNzFlIiwicm9sZSI6ImFkbWluIiwidXNlcm5hbWUiOiJTaGlraGFyWTEwIn0.u_2k0jy5w8EgJc6cxHgL1Sv1YK7nkGIJUYOumACA-Zw"
}
```

Posssible Responses Status Code:

| Code &nbsp; &nbsp; | Description &nbsp; &nbsp; |
| ------------------ | ------------------------- |
| 200                | Ok                        |
| 400                | Bad Request               |
| 500                | Internal server error     |

<br>
<hr>
<br>

### **4. Logout**

Method:

```
DELETE
```

Path:

```bash
http://127.0.0.1:10222/api/v1/logout
```

Security:

```json
Bearer Token

    Name: Authorization
    In: Header
    Example:
        Authorization: "Bearer <token>"
```

Request:

```bash
curl -X 'DELETE' \
  'http://127.0.0.1:10222/api/v1/logout' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzg5NzEzOTIsImlkIjoiNjQxMWQyNTU3N2E2YWZjZDFjZjlkNzFlIiwicm9sZSI6ImFkbWluIiwidXNlcm5hbWUiOiJTaGlraGFyWTEwIn0.EI0X31PPXi5pziDGMyVS0FnNpbT7je4aUSTiITEpOvw'
```

Response:

```json
Status Code
200 OK

Response Body
"Successfully Logout"
```

Posssible Responses Status Code:

| Code &nbsp; &nbsp; | Description &nbsp; &nbsp; |
| ------------------ | ------------------------- |
| 200                | Ok                        |
| 401                | Unauthorized              |

<br>
<hr>
<br>

### **5. Get One User**

Method:

```
GET
```

Path:

```bash
http://127.0.0.1:10222/api/v1/user/<username>
```

Security:

```json
Bearer Token

    Name: Authorization
    In: Header
    Example:
        Authorization: "Bearer <token>"
```

Request:

```bash
curl -X 'GET' \
  'http://127.0.0.1:10222/api/v1/user/swat15' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzg5NzE3NjksImlkIjoiNjQxMWQyNTU3N2E2YWZjZDFjZjlkNzFlIiwicm9sZSI6ImFkbWluIiwidXNlcm5hbWUiOiJTaGlraGFyWTEwIn0.i9YwhcH4IJJ1edClV3XgI4LZAkRPWAN98f7WM4-N_kE'
```

Response:

```json
Status Code
200 Ok

Response Body
{
  "id": "64122c71d9d68a1199f78295",
  "name": "Swatantra Yadav",
  "organisation": "google",
  "role": "admin",
  "username": "swat15"
}
```

Posssible Responses Status Code:

| Code &nbsp; &nbsp; | Description &nbsp; &nbsp; |
| ------------------ | ------------------------- |
| 200                | Ok                        |
| 400                | Bad Request               |
| 401                | Unauthorized              |
| 500                | Internal server error     |

<br>
<hr>
<br>

### **6. Get Multiple One User**

Method:

```
GET
```

Path:

```bash
http://127.0.0.1:10222/api/v1/users
```

Security:

```json
Bearer Token

    Name: Authorization
    In: Header
    Example:
        Authorization: "Bearer <token>"
```

Request:

```bash
curl -X 'GET' \
  'http://127.0.0.1:10222/api/v1/users?l=2&p=1' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzg5NzE3NjksImlkIjoiNjQxMWQyNTU3N2E2YWZjZDFjZjlkNzFlIiwicm9sZSI6ImFkbWluIiwidXNlcm5hbWUiOiJTaGlraGFyWTEwIn0.i9YwhcH4IJJ1edClV3XgI4LZAkRPWAN98f7WM4-N_kE'
```

Response:

```json
Status Code
200 OK

Response Body:
[
  {
    "_id": "64122c71d9d68a1199f78295",
    "name": "Swatantra Yadav",
    "username": "swat15",
    "organisation": "google",
    "role": "admin"
  },
  {
    "_id": "64122c8bd9d68a1199f78296",
    "name": "Varsha Yadav",
    "username": "varsha11",
    "organisation": "google",
    "role": "user"
  }
]
```

Posssible Responses Status Code:

| Code &nbsp; &nbsp; | Description &nbsp; &nbsp; |
| ------------------ | ------------------------- |
| 200                | Ok                        |
| 400                | Bad Request               |
| 401                | Unauthorized              |
| 500                | Internal server error     |

MongoDB user Collection Content After Unit Testing:

```json
[
	{
		"_id": { "$oid": "6411d25577a6afcd1cf9d71e" },
		"name": "Shikhar Yadav",
		"username": "ShikharY10",
		"password": "+WmjGVFGNN3vW8pkegb/dVmmxu3UPgXbxiU100G3kXw=",
		"organisation": "google",
		"role": "admin"
	},
	{
		"_id": { "$oid": "6412283c21135fb279075a1e" },
		"name": "Shivani Yadav",
		"username": "shivani22",
		"password": "+WmjGVFGNN3vW8pkegb/dVmmxu3UPgXbxiU100G3kXw=",
		"organisation": "SSC",
		"role": "user"
	},
	{
		"_id": { "$oid": "64122c71d9d68a1199f78295" },
		"name": "Swatantra Yadav",
		"username": "swat15",
		"password": "+WmjGVFGNN3vW8pkegb/dVmmxu3UPgXbxiU100G3kXw=",
		"organisation": "google",
		"role": "admin"
	},
	{
		"_id": { "$oid": "64122c8bd9d68a1199f78296" },
		"name": "Varsha Yadav",
		"username": "varsha11",
		"password": "+WmjGVFGNN3vW8pkegb/dVmmxu3UPgXbxiU100G3kXw=",
		"organisation": "google",
		"role": "user"
	},
	{
		"_id": { "$oid": "6412ea9a6c0c9b071e7e9395" },
		"name": "Ananya Raman",
		"username": "Ananya22",
		"password": "+WmjGVFGNN3vW8pkegb/dVmmxu3UPgXbxiU100G3kXw=",
		"organisation": "SSC",
		"role": "user"
	},
	{
		"_id": { "$oid": "6412efe42ddfc8927e6cc444" },
		"name": "Adarsh Yadav",
		"username": "adarsh04",
		"password": "z2vsanAnF1h6sQGYZRGuMToHUb9ShJK1o70lVbNwC0s=",
		"organisation": "SSC",
		"role": "user"
	}
]
```
