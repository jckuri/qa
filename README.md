# Questions and Answers

## Golang Development Program - Level 6 Final Project

You are to design the backend side of a system for the following business idea.

We want to build a site called QuestionsAndAnswers.com that will compete with Quora/Stackoverflow and others with 1 major difference. We will only allow 1 answer per question. If someone thinks they have a better answer, they will have to update the existing answer for that question instead of adding another answer. In essence, each question can only have 0 or 1 answer.

The backend should support the following operations:
- Get one question by its ID
- Get a list of all questions
- Get all the questions created by a given user
- Create a new question
- Update an existing question (the statement and/or the answer)
- Delete an existing question

No user tracking or security needed for this version. Database design is up to you.

We would like to receive code that runs, so remember to focus on the MVP functionality. You can document what’s missing that you wish you had more time for? Please think about the different problems you might encounter if the business idea is successful. This would include considerations such as increased load, increased data, and an upvoting feature.

## Solution

Since this website does not require authentication, the database could be very simple requiring just one table:

```
// QA is a structure to store Questions with their Answers.
type QA struct {
    Id string `bson:"id"` // Id is a string that represents the QA.
    Question string `bson:"question"` // Question is a string that contains the question.
    QUser string `bson:"quser"` // QUser is the user who asked the question.
    Answer string `bson:"answer"` // Answer is a string that contains the answer to the question.
    AUser string `bson:"auser"` // AUser is the user who answered the question.
}
```

Then, some functions were implemented in order to access the MongoDB with CRUD operations: Create, Read, Update, and Delete.

After that, GoKit microservices were implemented in order to expose the CRUD operations to the Internet. First, the server side was implemented, and then, the client side. Marshal and unmarshal were needed in order to convert Golang structures to JSON representations, back and forth.

```
// Service is an interface with many abstract functions.
type Service interface {
    ReadQuestion(ctx context.Context, id string) (qa_db.QA, error)
    ReadAllQuestions(ctx context.Context) ([]qa_db.QA, error)
    CreateQuestion(ctx context.Context, qa1 qa_db.QA) (qa_db.QA, error)
    UpdateQuestion(ctx context.Context, qa1 qa_db.QA) error
    DeleteQuestion(ctx context.Context, id string) error
    DeleteAllQuestions(ctx context.Context) error
    ReadQuestionsOfUser(ctx context.Context, quser string) ([]qa_db.QA, error)
    ReadAnswersOfUser(ctx context.Context, auser string) ([]qa_db.QA, error)
}
```

Finally, the remote API of GoKit microservices was tested in the client side with the script `test_all.sh`. Since GoKit microservices are exposed to the Internet, it is also possible to access them via the `curl` command in this script `curl_read_all_questions.sh `.

## Installation

**Prerequisites:** Golang <https://golang.org/> and Docker <https://www.docker.com/>.

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
