# LINE TOWN election

<!-- ABOUT THE PROJECT -->
## About The Project

This project is the Solution Engineer Assignment for the full-stack engineer - see the [spec.md](spec.md) file for details

### Built With

* [Go](https://go.dev/)
* [Fiber](https://gofiber.io/)
* [Viper](https://github.com/spf13/viper)
* [Zap](https://github.com/uber-go/zap)
* [Gorm](https://gorm.io/)
* [PostgreSQL](https://www.postgresql.org/)

<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

1. Make sure that you have Go 1.18 or above installed.
2. PostgreSQL database is required.

### Config file
Edit config.yaml at ./cmd/service, i.e.

```
app:
  name: "election-service"
  stage: "development"
  timezone: "Asia/Bangkok"
  cost: 13

database:
  driver: "postgres"
  host: "127.0.0.1"
  port: 5432
  username: "guest"
  password: "guest"
  database: "postgres"
  sslmode: "disable"

server:
  driver: "fiber"
  host: "127.0.0.1"
  port: 8000
  prefork: false
```

### Installation

Download vendor packages:
```sh
go mod download
```

### Run the dev server

```sh
cd cmd/service/
go run main.go
```

dev server running at: http://localhost:8000/

### Run via docker
Let’s run the server automatically via docker-compose:

```sh
docker-compose -p local -f ./docker-compose.yml up -d
```

The sample data is in PostgrstSQL.
Or run manually. but it need to setup PostgrstSQL first.
Let’s build an image called "lte-backend":

```sh
docker build -t lte-backend:0.1.0 .
```

Now that our image is built, we can start a container with the following command, which will serve our app on port 8000.

```sh
docker run --rm -it -p 8000:8000 --name lte-backend lte-backend:0.1.0 
```

<!-- CONTACT -->
## Contact

* **Tanawat Hongthai** - *tana.h@kkumail.com* - [Lazts](https://lazts.com)

<!-- LICENSE -->
## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
