package qa_gokit

import (
    "context"
    "qa/pkg/qa_db"
)

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

// QAService is an empty structure that will help to implement the abstract functions of the Service interface.
type QAService struct {}

// NewService returns an instance of the QAService structure.
func NewService() Service {
    return QAService {}
}

// ReadQuestion implements an abstract function of the Service interface by using the QAService structure.
func (QAService) ReadQuestion(ctx context.Context, id string) (qa_db.QA, error) {
    qa1, err := qa_db.GetQA(id)
    if err != nil {
        return qa_db.QA {}, err
    } else {
        return qa1, nil
    }
}

// ReadAllQuestions implements an abstract function of the Service interface by using the QAService structure.
func (QAService) ReadAllQuestions(ctx context.Context) ([]qa_db.QA, error) {
    qas, err := qa_db.GetAllQA()
    if err != nil {
        return nil, err
    } else {
        return qas, nil
    }
}

// CreateQuestion implements an abstract function of the Service interface by using the QAService structure.
func (QAService) CreateQuestion(ctx context.Context, qa1 qa_db.QA) (qa_db.QA, error) {
    qa2, err := qa_db.CreateQA(qa1)
    if err != nil {
        return qa2, err
    } else {
        return qa2, nil
    }
}

// UpdateQuestion implements an abstract function of the Service interface by using the QAService structure.
func (QAService) UpdateQuestion(ctx context.Context, qa1 qa_db.QA) error {
    return qa_db.UpdateQA(qa1)
}

// DeleteQuestion implements an abstract function of the Service interface by using the QAService structure.
func (QAService) DeleteQuestion(ctx context.Context, id string) error {
    return qa_db.DeleteQA(id)
}

// DeleteAllQuestions implements an abstract function of the Service interface by using the QAService structure.
func (QAService) DeleteAllQuestions(ctx context.Context) error {
    return qa_db.DeleteAllQAs()
}

// ReadQuestionsOfUser implements an abstract function of the Service interface by using the QAService structure.
func (QAService) ReadQuestionsOfUser(ctx context.Context, quser string) ([]qa_db.QA, error) {
    qas, err := qa_db.GetQuestionsOfUser(quser)
    if err != nil {
        return nil, err
    } else {
        return qas, nil
    }
}

// ReadAnswersOfUser implements an abstract function of the Service interface by using the QAService structure.
func (QAService) ReadAnswersOfUser(ctx context.Context, quser string) ([]qa_db.QA, error) {
    qas, err := qa_db.GetAnswersOfUser(quser)
    if err != nil {
        return nil, err
    } else {
        return qas, nil
    }
}
