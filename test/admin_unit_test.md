# API Unit Test Results and Behaviours For Admin Routes

MongoDB User-Collection content before starting admin Unit Testing:

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

> I'm logged in as swat15 which is a admin account.

## **1. Admin Create New User**

Method:

```
POST
```

Path:

```bash
http://127.0.0.1:10222/api/v1/admin/user
```

Request:

```bash
curl -X 'POST' \
  'http://127.0.0.1:10222/api/v1/admin/user' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzkwMzYwMzQsImlkIjoiNjQxMjJjNzFkOWQ2OGExMTk5Zjc4Mjk1Iiwicm9sZSI6ImFkbWluIiwidXNlcm5hbWUiOiJzd2F0MTUifQ.arnTnbGsa2ZLduSwaKPzTT8VFad_NdVdJoDtmFf8rEs' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "Sam Harnoor",
  "organisation": "google",
  "password": "1234",
  "username": "Sam10"
}'
```

Response:

It does not return any type token.
In order get the token created user needs to login.

```json
Status Code

200 OK

Response Body

{
  "_id": "641401b0c2f42a404fde43ea",
  "name": "Sam Harnoor",
  "username": "Sam10",
  "organisation": "google",
  "role": "user"
}
```

Posssible Responses Status Code:

| Code &nbsp; &nbsp; | Description &nbsp; &nbsp; | Response Body DataType |
| ------------------ | ------------------------- | ---------------------- |
| 200                | Ok                        | Object                 |
| 400                | Bad Request               | String                 |
| 401                | Unauthorized              | String                 |
| 500                | Internal server error     | String                 |

<br>
<hr>
<br>

## **2. Admin Delete A User**

Method:

```
DELETE
```

Path:

```bash
http://127.0.0.1:10222/api/v1/admin/user/<userid>
```

Request:

```bash
curl -X 'DELETE' \
  'http://127.0.0.1:10222/api/v1/admin/user/64122c8bd9d68a1199f78296' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzkwMzYwMzQsImlkIjoiNjQxMjJjNzFkOWQ2OGExMTk5Zjc4Mjk1Iiwicm9sZSI6ImFkbWluIiwidXNlcm5hbWUiOiJzd2F0MTUifQ.arnTnbGsa2ZLduSwaKPzTT8VFad_NdVdJoDtmFf8rEs'
```

Response:

```json
Status Code

200 OK

Response Body

{
  "_id": "64122c8bd9d68a1199f78296",
  "name": "Varsha Yadav",
  "username": "varsha11",
  "password": "+WmjGVFGNN3vW8pkegb/dVmmxu3UPgXbxiU100G3kXw=",
  "organisation": "google",
  "role": "user"
}
```

Posssible Responses Status Code:

| Code &nbsp; &nbsp; | Description &nbsp; &nbsp; | Response Body DataType |
| ------------------ | ------------------------- | ---------------------- |
| 200                | Ok                        | Object                 |
| 400                | Bad Request               | String                 |
| 401                | Unauthorized              | String                 |
| 500                | Internal server error     | String                 |

<br>
<hr>
<br>

## **3. Admin GET One User**

Method:

```
GET
```

Path:

```bash
http://127.0.0.1:10222/api/v1/admin/user/<userId>
```

Request:

```bash
curl -X 'GET' \
  'http://127.0.0.1:10222/api/v1/admin/user/641401b0c2f42a404fde43ea' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzkwMzYwMzQsImlkIjoiNjQxMjJjNzFkOWQ2OGExMTk5Zjc4Mjk1Iiwicm9sZSI6ImFkbWluIiwidXNlcm5hbWUiOiJzd2F0MTUifQ.arnTnbGsa2ZLduSwaKPzTT8VFad_NdVdJoDtmFf8rEs'
```

Response:

```json

Status Code
200 OK

Response Body
{
  "_id": "641401b0c2f42a404fde43ea",
  "name": "Sam Harnoor",
  "username": "Sam10",
  "password": "+WmjGVFGNN3vW8pkegb/dVmmxu3UPgXbxiU100G3kXw=",
  "organisation": "google",
  "role": "user"
}

```

Posssible Responses Status Code:

| Code &nbsp; &nbsp; | Description &nbsp; &nbsp; | Response Body DataType |
| ------------------ | ------------------------- | ---------------------- |
| 200                | Ok                        | Object                 |
| 400                | Bad Request               | String                 |
| 401                | Unauthorized              | String                 |
| 500                | Internal server error     | String                 |

<br>
<hr>
<br>

## **4. Admin GET Multiple User**

Method:

```
GET
```

Path:

```bash
http://127.0.0.1:10222/api/v1/admin/users?l=<number-of-search-results>&p=<page-number>
```

Request:

```bash
curl -X 'GET' \
  'http://127.0.0.1:10222/api/v1/admin/users?l=2&p=1' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzkwMzYwMzQsImlkIjoiNjQxMjJjNzFkOWQ2OGExMTk5Zjc4Mjk1Iiwicm9sZSI6ImFkbWluIiwidXNlcm5hbWUiOiJzd2F0MTUifQ.arnTnbGsa2ZLduSwaKPzTT8VFad_NdVdJoDtmFf8rEs'
```

Response:

```json

Status Code
200 OK

Response Body
[
  {
    "_id": "6411d25577a6afcd1cf9d71e",
    "name": "Shikhar Yadav",
    "username": "ShikharY10",
    "password": "+WmjGVFGNN3vW8pkegb/dVmmxu3UPgXbxiU100G3kXw=",
    "organisation": "google",
    "role": "admin"
  },
  {
    "_id": "641401b0c2f42a404fde43ea",
    "name": "Sam Harnoor",
    "username": "Sam10",
    "password": "+WmjGVFGNN3vW8pkegb/dVmmxu3UPgXbxiU100G3kXw=",
    "organisation": "google",
    "role": "user"
  }
]

```

Posssible Responses Status Code:

| Code &nbsp; &nbsp; | Description &nbsp; &nbsp; | Response Body DataType |
| ------------------ | ------------------------- | ---------------------- |
| 200                | Ok                        | Object                 |
| 400                | Bad Request               | String                 |
| 401                | Unauthorized              | String                 |
| 500                | Internal server error     | String                 |

<br>
<hr>
<br>

## **5. Admin Create New Admin**

Method:

```
PUT
```

Path:

```bash
http://127.0.0.1:10222/api/v1/admin/new/<userId>
```

Request:

```bash
curl -X 'PUT' \
  'http://127.0.0.1:10222/api/v1/admin/new/641401b0c2f42a404fde43ea' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzkwMzYwMzQsImlkIjoiNjQxMjJjNzFkOWQ2OGExMTk5Zjc4Mjk1Iiwicm9sZSI6ImFkbWluIiwidXNlcm5hbWUiOiJzd2F0MTUifQ.arnTnbGsa2ZLduSwaKPzTT8VFad_NdVdJoDtmFf8rEs'
```

Response:

```json

Status Code
200 Ok

Response Body
"Successfully updated role"

```

Posssible Responses Status Code:

| Code &nbsp; &nbsp; | Description &nbsp; &nbsp; | Response Body DataType |
| ------------------ | ------------------------- | ---------------------- |
| 200                | Ok                        | String                 |
| 400                | Bad Request               | String                 |
| 401                | Unauthorized              | String                 |
| 500                | Internal server error     | String                 |

<br>
MongoDB User-Collection content after starting admin Unit Testing:

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
	},
	{
		"_id": { "$oid": "641401b0c2f42a404fde43ea" },
		"name": "Sam Harnoor",
		"username": "Sam10",
		"password": "+WmjGVFGNN3vW8pkegb/dVmmxu3UPgXbxiU100G3kXw=",
		"organisation": "google",
		"role": "admin"
	}
]
```
