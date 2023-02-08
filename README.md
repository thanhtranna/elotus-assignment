## eLotus Assignment


- Total time: 50 hours (please no more!)

## Data Structures And Algorithms

> Time: 2 hours
> Please commit and push the code of this round after 2 hours of receiving the test, then you can start the Hackathon round

1. [Gray Code](/01.gray-code/README.md)

2. [Sum of Distances in Tree](/02.sum-of-distances-in-tree/README.md)

3. [Maximum Length of Repeated Subarray](/03.maximum-length-of-repeated-subarray/README.md)

## Hackathon

### How to run this project?

First, make sure docker running on local machine.

- Clone this project

```bash
# Move to your workspace
cd your-workspace

# Clone this project into your workspace
git clone https://github.com/thanhtranna/elotus-assignment.git

# Move to the project root directory
cd elotus-assignment
```

#### Run with Docker

- Install Docker and Docker Compose.
- Run `docker-compose -f docker-compose.yaml up -d`.
- Access API using `http://localhost:8080`



### Example API Request and Response

- Check ping server

  - Request

  ```bash
  curl --location --request GET 'localhost:8080/api/ping'
  ```

  - Response

  ```json
  {
    "message": "pong"
  }
  ```

- register new user

  - Request

  ```bash
  curl --location --request POST 'localhost:8080/api/user/register' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "name": "Test Name",
        "email": "test_gmail@gmail.com",
        "username": "username",
        "password": "123456"
    }'
  ```

  - Response

  ```json
  {
    "email": "test_gmail@gmail.com",
    "userId": 1,
    "username": "username",
    "created_at": "2023-02-08T10:53:21.295Z",
  }
  ```

- login

  - Request

  ```bash
  curl --location --request POST 'localhost:8080/api/user/login' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "email": "test_gmail@gmail.com",
        "password": "123456"
    }'
  ```

  - Response

  ```json
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRoYW5oLnRyYW4tdm4iLCJlbWFpbCI6InRoYW5oLnRyYW4tdm5Aa2xpa2Rva3Rlci5jb20iLCJleHAiOjE2NzU4NTQwOTZ9.hs0CYDzvNql-02oWMJajUgc5-PaLwXyOk8QWIo18abw"
  }
  ```

- upload image

  - Request

  ```bash
  curl --location --request POST 'localhost:8080/api/secured/upload' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRoYW5oLnRyYW4tdm4iLCJlbWFpbCI6InRoYW5oLnRyYW4tdm5Aa2xpa2Rva3Rlci5jb20iLCJleHAiOjE2NzU4NTQwOTZ9.hs0CYDzvNql-02oWMJajUgc5-PaLwXyOk8QWIo18abw' \
    --form 'data=@"/home/thanhtran/Downloads/PngItem_6205768.png"'
  ```

  - Response

  ```json
  {
    "message": "Successfully!",
    "filename": "PngItem_6205768-pnANPCqdVx.png"
  }
  ```

- show image

  - Request

  ```bash
  curl --location --request GET 'localhost:8080/public/PngItem_6205768-tAkRNbRPlJ.png'
  ```

  - Response

  ```json
    Detail Image
  ```