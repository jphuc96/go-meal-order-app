# Lunch Order

### A backend API server of lunch order application for serving people in an office.

![](logo.png)

## Getting Started

Below instructions will get you into the project and guide you how to run it in local computer and deploy it to a production system. 

### Prerequisites

- Latest version of Docker and docker-compose
- Go 1.12
- PostgreSQL 10.4 (will run in docker, so no need to install locally)

### Installing
#### Step 1: Setup local database
Firstly, clone the repo to $GOPATH/git.d.foundation/datcom directory

```
cd $GOPATH/git.d.foundation/datcom
git clone https://git.d.foundation/datcom/backend.git
```

We will run a local database for testing purpose by deploying the docker-compose.yml file in 'test_postgres' folder.
The docker compose file includes an adminer container for better managing the database.

```
docker-compose -f test_postgres/docker-compose.yml up -d
```
After deploying local db, use following enviroments variable for afterward Go program by putting them into a config.env file

```
export DB_HOST=localhost
export DB_PORT=5432
export DB_NAME=datcom
export DB_USER=postgres
export DB_PASSWORD=datcom
export DB_SSL=disable
```
#### Step 2: Build and run Go program
Build the project:
```
make build
```
Migrate up the database
```
make migrate-up
```
Run the development program
```
source ./config.env
make run
```
The output shoule be ended as belows, and the server should be listening on localhost:8080
```
[GIN-debug] Listening and serving HTTP on 127.0.0.1:8080
```
## Running the tests
The program is an API server. The testing API requests and responses will depend on several authentication by OAuth2 Google Sign In feature that require the combination of frontend web and mobile apps.

## RESTful API definition
Please refer [here](./docs/api.md) for API definitions document.

## Deployment
The real production database will be deployed on the Heroku system.

The project is configured to use Gitlab Runner for running the build and test.

Google Cloud Builds is used to build the Docker image, then it will be deploy to the Google Container Optimized OS running on the Compute Engine service.

The CI/CD pipeline is configured to automatically run the testing, building and deplying on corresponding commit. See .gitlab-ci.yml for more detail.


## Built and Deployed With

* [Gitlab CI]() 
* [Cloud Build](https://cloud.google.com/cloud-build/) - Google Docker image building server
* [Container Optimized OS](https://cloud.google.com/container-optimized-os/) The OS built for containers
* [Heroku](https://heroku.com) - Deploy application effortlessly

## Contributor
### Members
* [Thong Nguyen](https://git.d.foundation/thongnt)
* [Quang Trinh](https://git.d.foundation/quangtk)
* [Dong Le](https://git.d.foundation/donglb)
* [Truc Huynh](https://git.d.foundation/trucht)
* [Phuc Nguyen](https://git.d.foundation/phucnh)
### Mentors and reviewers
* [Hieu Nguyen](https://git.d.foundation/hieunm)
* [Minh Tran](https://git.d.foundation/thminh)
* [Phan Quang Hieu](https://git.d.foundation/hieupq)
* [Quang Le](https://git.d.foundation/quang)
* [Phat Nguyen](https://git.d.foundation/phatnt1995)

## Versioning
* v1.0.0

## License

No License yet