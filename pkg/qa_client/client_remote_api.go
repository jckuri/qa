// qa_client is the package that contains functions to remotely call the GoKit microservices.
// To do so, it is necessary to send JSON requests and to unmarshal the JSON responses.
package qa_client

import (
    "fmt"
    "net/http"
    "qa/pkg/qa_db"
)

// baseUrl is a string with the base URL.
var baseUrl = "http://127.0.0.1:8080"

// GetUrl concatenates the base URL with the path.
func GetUrl(path string) string {
    return fmt.Sprintf("%v%v", baseUrl, path)
}

// MicroServiceError creates an error with a message.
func MicroServiceError(status int) error {
    return fmt.Errorf("MICROSERVICE ERROR: %v %v", status, http.StatusText(status))
}

// ReadQuestion reads 1 question with the corresponding ID.
func ReadQuestion(id string) (qa_db.QA, error) {
    path := fmt.Sprintf("/read_question/%v", id)
    response, status := JsonRequest(GetUrl(path), http.MethodGet, "")
    if status != 200 {
        return qa_db.QA {}, MicroServiceError(status)
    }
    uerr := UnmarshalErr(response)
    if uerr != nil {
        return qa_db.QA {}, uerr
    }
    qa1 := ParseQuestionFromJson(response)
    return qa1, nil
}

// ReadAllQuestions reads all the questions in the server.
func ReadAllQuestions() ([]qa_db.QA, error) {
    response, status := JsonRequest(GetUrl("/read_all_questions"), http.MethodGet, "")
    if status != 200 {
        return nil, MicroServiceError(status)
    }
    uerr := UnmarshalErr(response)
    if uerr != nil {
        return nil, uerr
    }
    qas := ParseQuestionsFromJson(response)
    return qas, nil
}

// CreateQuestion creates a new question.
func CreateQuestion(qa1 qa_db.QA) (qa_db.QA, error) {
    values := make(map[string]interface{})
    values["qa"] = qa1
    response, status := JsonMapRequest(GetUrl("/create_question"), http.MethodPost, values)
    if status != 200 {
        return qa_db.QA {}, MicroServiceError(status)
    }
    uerr := UnmarshalErr(response)
    if uerr != nil {
        return qa_db.QA {}, uerr
    }
    qa2 := ParseQuestionFromJson(response)
    return qa2, nil
}

// UpdateQuestion updates an existent question. If the question does not exist, an error is thrown.
func UpdateQuestion(qa1 qa_db.QA) error {
    values := make(map[string]interface{})
    values["qa"] = qa1
    response, status := JsonMapRequest(GetUrl("/update_question"), http.MethodPut, values)
    if status != 200 {
        return MicroServiceError(status)
    }
    return UnmarshalErr(response)
}

// DeleteQuestion deletes an existent question. If the question does not exist, an error is thrown.
func DeleteQuestion(id string) error {
    path := fmt.Sprintf("/delete_question/%v", id)
    response, status := JsonRequest(GetUrl(path), http.MethodDelete, "")
    if status != 200 {
        return MicroServiceError(status)
    }
    return UnmarshalErr(response)
}

// DeleteAllQuestions deletes all the questions in the server.
func DeleteAllQuestions() error {
    response, status := JsonRequest(GetUrl("/delete_all_questions"), http.MethodDelete, "")
    if status != 200 {
        return MicroServiceError(status)
    }
    return UnmarshalErr(response)
}

// ReadQuestionsOfUser reads all the questions asked by a user.
func ReadQuestionsOfUser(user string) ([]qa_db.QA, error) {
    path := fmt.Sprintf("/read_questions_of_user/%v", user)
    response, status := JsonRequest(GetUrl(path), http.MethodGet, "")
    if status != 200 {
        return nil, MicroServiceError(status)
    }
    uerr := UnmarshalErr(response)
    if uerr != nil {
        return nil, uerr
    }
    qas := ParseQuestionsFromJson(response)
    return qas, nil
}

// ReadAnswersOfUser reads all the answers of a user.
func ReadAnswersOfUser(user string) ([]qa_db.QA, error) {
    path := fmt.Sprintf("/read_answers_of_user/%v", user)
    response, status := JsonRequest(GetUrl(path), http.MethodGet, "")
    if status != 200 {
        return nil, MicroServiceError(status)
    }
    uerr := UnmarshalErr(response)
    if uerr != nil {
        return nil, uerr
    }
    qas := ParseQuestionsFromJson(response)
    return qas, nil
}
