package qa_gokit

import (
	"context"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"qa/pkg/qa_db"
)

// Encode Response:

// EncodeResponse encodes all the responses in a very general way because the parameter response is very abstract.
func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
    return json.NewEncoder(w).Encode(response)
}

// Empty Request:

// EmptyRequest is a structure that represents an empty request.
type EmptyRequest struct {
}

// DecodeEmptyRequest decodes an empty request.
func DecodeEmptyRequest(ctx context.Context, r *http.Request) (interface{}, error) {
    var req EmptyRequest
    return req, nil
}

// EmptyResponse is a structure that represents an empty response.
type EmptyResponse struct {
    Err string `json:"err,omitempty"` // Err is a string with an error message.
}

// ID Request:

// IDRequest is a structure that represents a request with an ID.
type IDRequest struct {
    Id string `json:"id"` // Id is a string with the ID of a QA structure.
}

// DecodeIDRequest decodes a IDRequest.
func DecodeIDRequest(ctx context.Context, r *http.Request) (interface{}, error) {
    var req IDRequest
    req.Id = mux.Vars(r)["id"]
    return req, nil
}

// User Request:

// UserRequest is a structure that represents a request with a user.
type UserRequest struct {
    User string `json:"id"` // User is a string with the user of a QA structure.
}

// DecodeUserRequest decodes a UserRequest.
func DecodeUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
    var req UserRequest
    req.User = mux.Vars(r)["user"]
    return req, nil
}

// QuestionRequest:

// QuestionRequest is a structure that represents a request with a QA structure and an error message.
type QuestionRequest struct {
    QA qa_db.QA `json:"qa"` // QA is a QA structure.
    Err  string `json:"err,omitempty"` // Err is a string with an error message.
}

// DecodeQuestionRequest decodes a QuestionRequest.
func DecodeQuestionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
    var req QuestionRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        return nil, err
    }
    return req, nil
}

// GetQA:

// QuestionResponse is a structure that represents a response with a QA structure and an error message.
type QuestionResponse struct {
    QA qa_db.QA `json:"qa"`
    Err  string `json:"err,omitempty"`
}

// GetAll:

// QuestionsResponse is a structure that represents a response with a slice of QA structures and an error message.
type QuestionsResponse struct {
    QAs []qa_db.QA `json:"qas"`
    Err  string `json:"err,omitempty"`
}
