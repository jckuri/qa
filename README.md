# Questions and Answers

## Installation

Open a new terminal and execute the following commands:

```
git clone https://github.com/jckuri/qa
cd qa
./clean-and-start-mongodb-in-docker.sh
```

Open another new terminal and execute the following command: 

```
./docker-clean-build-and-run.sh
```

Open yet another new terminal and execute the following commands: 

```
./test_all.sh 
./curl_read_all_questions.sh 
./test_delete_all_questions.sh
./curl_read_all_questions.sh 
./start_godoc.sh
```

## Directory Structure

```
$ tree --dirsfirst
.
├── cmd
│   ├── client_tests
│   │   └── client_test.go
│   └── server
│       └── main.go
├── pkg
│   ├── qa_client
│   │   ├── client_remote_api.go
│   │   ├── json_requests.go
│   │   └── unmarshal_qas.go
│   ├── qa_db
│   │   └── mongodb.go
│   └── qa_gokit
│       ├── endpoint.go
│       ├── server.go
│       ├── service.go
│       └── transport.go
├── clean-and-start-mongodb-in-docker.sh
├── curl_read_all_questions.sh
├── docker-clean-build-and-run.sh
├── Dockerfile
├── go.mod
├── go.sum
├── README.md
├── start_godoc.sh
├── test_all.sh
└── test_delete_all_questions.sh

7 directories, 20 files
```
