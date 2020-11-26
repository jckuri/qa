// qa_tests is the package of the remote client tests.
package qa_tests

import (
    "testing"
    "reflect"
    "qa/pkg/qa_db"
    "qa/pkg/qa_client"
)

// TestDeleteAllQuestions tests the microservice DeleteAllQuestions.
func TestDeleteAllQuestions(t *testing.T) {
    err := qa_client.DeleteAllQuestions()
    if err != nil {
        t.Error(err)    
    }
    qas, err := qa_client.ReadAllQuestions()
    if err != nil {
        t.Error(err)    
    }
    if len(qas) != 0 {
        t.Errorf("Some questions were not deleted.")
    }
}

// CreateQuestionTest is a helper function that helps to test the microservice CreateQuestion.
func CreateQuestionTest(t *testing.T, qa0 qa_db.QA) {
    qa1, err1 := qa_client.CreateQuestion(qa0)
    if err1 != nil {
        t.Error(err1)
    }
    qa1_test, err1_test := qa_client.ReadQuestion(qa1.Id)
    if err1_test != nil {
        t.Error(err1_test)
    }
    if qa1 != qa1_test {
        t.Errorf("Got: %+v, want: %+v", qa1_test, qa1)
    }
}

// create_qas is a helper variable for the function TestCreateQuestions that stores a slice of QA structures.
var create_qas []qa_db.QA = []qa_db.QA {
    qa_db.QA {"", "Where are we?", "jckuri", "", ""},
    qa_db.QA {"", "What are we doing?", "ccedano", "", ""},
    qa_db.QA {"", "Where's Waldo?", "tpeycere", "", ""},
    qa_db.QA {"", "Who are we?", "tpeycere", "", ""},
}

// TestCreateQuestions tests the microservice CreateQuestion.
func TestCreateQuestions(t *testing.T) {
    for _, qa0 := range create_qas {
        CreateQuestionTest(t, qa0)
    }
    qas2, err := qa_client.ReadAllQuestions()
    if err != nil {
        t.Error(err)    
    }
    if len(create_qas) != len(qas2) {
        t.Errorf("Some questions were not created.")
    }
}

// UpdateQuestionTest is a helper function that helps to test the microservice UpdateQuestion.
func UpdateQuestionTest(t *testing.T, qa0 qa_db.QA) {
    err0 := qa_client.UpdateQuestion(qa0)
    if err0 != nil {
        t.Error(err0)
    }
    qa1, err1 := qa_client.ReadQuestion(qa0.Id)
    if err1 != nil {
        t.Error(err1)
    }
    if qa0 != qa1 {
        t.Errorf("Got: %+v, want: %+v", qa1, qa0)
    }
}

// update_qas is a helper variable for the function TestUpdateQuestions that stores a slice of QA structures.
var update_qas []qa_db.QA = []qa_db.QA {
    qa_db.QA {"1", "Where are we?", "jckuri", "We are in Latin America.", "ccedano"},
    qa_db.QA {"2", "What are we doing?", "ccedano", "We are programming a project.", "jckuri"},
    qa_db.QA {"3", "Where's Waldo?", "tpeycere", "Here.", "jckuri"},
}

// TestUpdateQuestions tests the microservice UpdateQuestion.
func TestUpdateQuestions(t *testing.T) {
    for _, qa0 := range update_qas {
        UpdateQuestionTest(t, qa0)
    }
}

// TestUpdateInexistentQuestion tests the microservice UpdateQuestion if it handles inexistent questions in a correct way by throwing an error.
func TestUpdateInexistentQuestion(t *testing.T) {
    err5 := qa_client.UpdateQuestion(qa_db.QA {"5", "Does question 5 exist?", "jckuri", "No, it doesn't.", "ccedano"},)
    if err5 == nil {
        t.Error("Question 5 doesn't exist. So, updating question 5 should have caused an error.")
    }
}

// TestReadQuestionsOfUser tests the microservice ReadQuestionsOfUser.
func TestReadQuestionsOfUser(t *testing.T) {
    questions_got, err1 := qa_client.ReadQuestionsOfUser("ccedano")
    if err1 != nil {
        t.Error(err1)
    }
    questions_want := []qa_db.QA {update_qas[1]}
    if !reflect.DeepEqual(questions_got, questions_want) {
        t.Errorf("Got: %+v, want: %+v", questions_got, questions_want)
    }
}

// TestReadAnswersOfUser tests the microservice ReadAnswersOfUser.
func TestReadAnswersOfUser(t *testing.T) {
    questions_got, err1 := qa_client.ReadAnswersOfUser("jckuri")
    if err1 != nil {
        t.Error(err1)
    }
    questions_want := []qa_db.QA {update_qas[1], update_qas[2]}
    if !reflect.DeepEqual(questions_got, questions_want) {
        t.Errorf("Got: %+v, want: %+v", questions_got, questions_want)
    }
}

// TestDeleteQuestion tests the microservice DeleteQuestion.
func TestDeleteQuestion(t *testing.T) {
    err1 := qa_client.DeleteQuestion("4")
    if err1 != nil {
        t.Error(err1)
    }
    qas2, err2 := qa_client.ReadAllQuestions()
    if err2 != nil {
        t.Error(err2)    
    }
    if !reflect.DeepEqual(qas2, update_qas) {
        t.Errorf("Got: %+v, want: %+v", qas2, update_qas)
    }
}

// TestDeleteInexistentQuestion tests the microservice DeleteQuestion if it handles inexistent questions in a correct way by throwing an error.
func TestDeleteInexistentQuestion(t *testing.T) {
    err5 := qa_client.DeleteQuestion("5")
    if err5 == nil {
        t.Error("Question 5 doesn't exist. So, deleting question 5 should have caused an error.")
    }
}
