# Deployment Rules

## **Setup**

> **Before starting, make sure you have git, docker and docker-compose installed in your system**

Step 1:

Clone this repository and open it in terminal

```bash
$ git clone https://github.com/HousewareHQ/houseware---backend-engineering-octernship-ShikharY10.git
```

Step 2:

cd into the root of this golang project

```bash
$ cd /houseware---backend-engineering-octernship-ShikharY10/backend
```

Step 3:

Build a Docker image of this golang project

```bash
$ docker build -f deployments/Dockerfile -t hw-api-server .
```

Step 4:

Start MongoDB and Redis on Docker using docker-compose

```bash
$ docker-compose -f deployments/db-docker-compose.yaml up -d
```

Step 5:

Run the server docker image inside a docker container

```bash
$ docker run --net=host -d  hw-api-server sh
```

<br>

### [Next Step](https://github.com/HousewareHQ/houseware---backend-engineering-octernship-ShikharY10/tree/main/backend/api): Test APIs
