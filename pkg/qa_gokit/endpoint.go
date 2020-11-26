package qa_gokit

import (
    "context"
    "github.com/go-kit/kit/endpoint"
    "qa/pkg/qa_db"
)

// Endpoints is a structure that contains all the endpoints.
type Endpoints struct {
    ReadQuestionEndpoint endpoint.Endpoint // The endpoint for the ReadQuestion function.
    ReadAllQuestionsEndpoint endpoint.Endpoint // The endpoint for the ReadAllQuestions function.
    CreateQuestionEndpoint endpoint.Endpoint // The endpoint for the CreateQuestion function.
    UpdateQuestionEndpoint endpoint.Endpoint // The endpoint for the UpdateQuestion function.
    DeleteQuestionEndpoint endpoint.Endpoint // The endpoint for the DeleteQuestion function.
    DeleteAllQuestionsEndpoint endpoint.Endpoint // The endpoint for the DeleteAllQuestions function.
    ReadQuestionsOfUserEndpoint endpoint.Endpoint // The endpoint for the ReadQuestionsOfUser function.
    ReadAnswersOfUserEndpoint endpoint.Endpoint // The endpoint for the ReadAnswersOfUser function.
}

// MakeReadQuestionEndpoint creates the endpoint for the ReadQuestion function.
func MakeReadQuestionEndpoint(srv Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(IDRequest)
        qa1, err := srv.ReadQuestion(ctx, req.Id)
        if err != nil {
            return QuestionResponse{qa_db.QA {}, err.Error()}, nil
        }
        return QuestionResponse{qa1, ""}, nil
    }
}

// MakeReadAllQuestionsEndpoint creates the endpoint for the ReadAllQuestions function.
func MakeReadAllQuestionsEndpoint(srv Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        _ = request.(EmptyRequest)
        qas, err := srv.ReadAllQuestions(ctx)
        if err != nil {
            return QuestionsResponse{nil, err.Error()}, nil
        }
        return QuestionsResponse{qas, ""}, nil
    }
}

// MakeCreateQuestionEndpoint creates the endpoint for the CreateQuestion function.
func MakeCreateQuestionEndpoint(srv Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(QuestionRequest)
        qa1, err := srv.CreateQuestion(ctx, req.QA)
        if err != nil {
            return QuestionResponse{qa_db.QA {}, err.Error()}, nil
        }
        return QuestionResponse{qa1, ""}, nil
    }
}

// MakeUpdateQuestionEndpoint creates the endpoint for the UpdateQuestion function.
func MakeUpdateQuestionEndpoint(srv Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(QuestionRequest)
        err := srv.UpdateQuestion(ctx, req.QA)
        if err != nil {
            return EmptyResponse {err.Error()}, nil
        } else {
            return EmptyResponse {}, err
        }
    }
}

// MakeDeleteQuestionEndpoint creates the endpoint for the DeleteQuestion function.
func MakeDeleteQuestionEndpoint(srv Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(IDRequest)
        err := srv.DeleteQuestion(ctx, req.Id)
        if err != nil {
            return EmptyResponse {err.Error()}, nil
        } else {
            return EmptyResponse {}, err
        }
    }
}

// MakeDeleteAllQuestionsEndpoint creates the endpoint for the DeleteAllQuestions function.
func MakeDeleteAllQuestionsEndpoint(srv Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        _ = request.(EmptyRequest)
        err := srv.DeleteAllQuestions(ctx)
        if err != nil {
            return EmptyResponse {err.Error()}, nil
        } else {
            return EmptyResponse {}, err
        }
    }
}

// MakeReadQuestionsOfUserEndpoint creates the endpoint for the ReadQuestionsOfUser function.
func MakeReadQuestionsOfUserEndpoint(srv Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(UserRequest)
        qas, err := srv.ReadQuestionsOfUser(ctx, req.User)
        if err != nil {
            return QuestionsResponse{nil, err.Error()}, nil
        }
        return QuestionsResponse{qas, ""}, nil
    }
}

// MakeReadAnswersOfUserEndpoint creates the endpoint for the ReadAnswersOfUser function.
func MakeReadAnswersOfUserEndpoint(srv Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(UserRequest)
        qas, err := srv.ReadAnswersOfUser(ctx, req.User)
        if err != nil {
            return QuestionsResponse{nil, err.Error()}, nil
        }
        return QuestionsResponse{qas, ""}, nil
    }
}
