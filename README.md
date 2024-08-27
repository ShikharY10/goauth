<div align="center">
  <img src="/media/goauth.png" alt="GoAuth Logo" width="320" height="300">
  <h1>GoAuth</h1>
  <strong>Authorization+Authentication service in Golang</strong>
  <h6>A backend API service in Golang that handles authorization and authentication for a web app</h6>
</div>
<br>

## File Structure

```
📦backend
 ┣ 📂api
 ┃ ┣ 📜README.md
 ┃ ┣ 📜docs.go
 ┃ ┣ 📜swagger.json
 ┃ ┗ 📜swagger.yaml
 ┣ 📂cmd
 ┃ ┣ 📂configs
 ┃ ┃ ┣ 📜cache.go
 ┃ ┃ ┣ 📜db.go
 ┃ ┃ ┗ 📜env.go
 ┃ ┣ 📂controllers
 ┃ ┃ ┗ 📂controller_v1
 ┃ ┃ ┃ ┣ 📜admin.go
 ┃ ┃ ┃ ┗ 📜user.go
 ┃ ┣ 📂handlers
 ┃ ┃ ┣ 📜cache.go
 ┃ ┃ ┣ 📜database.go
 ┃ ┃ ┗ 📜handler.go
 ┃ ┣ 📂middleware
 ┃ ┃ ┣ 📜jwt.go
 ┃ ┃ ┗ 📜role.go
 ┃ ┣ 📂models
 ┃ ┃ ┗ 📜models.go
 ┃ ┣ 📂routes
 ┃ ┃ ┗ 📂routes_v1
 ┃ ┃ ┃ ┣ 📜admin.go
 ┃ ┃ ┃ ┗ 📜user.go
 ┃ ┣ 📂utils
 ┃ ┃ ┣ 📜examiners.go
 ┃ ┃ ┗ 📜helpers.go
 ┃ ┗ 📜main.go
 ┣ 📂deployments
 ┃ ┣ 📜Dockerfile
 ┃ ┣ 📜README.md
 ┃ ┗ 📜db-docker-compose.yaml
 ┣ 📂docs
 ┃ ┗ 📜README.md
 ┣ 📂test
 ┃ ┣ 📜admin_unit_test.md
 ┃ ┗ 📜user_unit_test.md
 ┣ 📜.editorconfig
 ┣ 📜.env
 ┣ 📜.gitignore
 ┣ 📜README.md
 ┣ 📜go.mod
 ┗ 📜go.sum
```

# Documentation

For Documentation, [Click Here](https://github.com/HousewareHQ/houseware---backend-engineering-octernship-ShikharY10/tree/main/backend/docs)
