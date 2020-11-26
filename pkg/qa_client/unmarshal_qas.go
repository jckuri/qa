package qa_client

import (
    "encoding/json"
    "fmt"
    "qa/pkg/qa_db"
)

// UnmarshalledQA is a type, which is a map whose keys are strings and whose values are QA structures.
type UnmarshalledQA map[string]qa_db.QA

// ParseQuestionFromJson parses a QA structure from a JSON response.
func ParseQuestionFromJson(json1 string) qa_db.QA {
    bytes := []byte(json1)
    var uqa UnmarshalledQA
    json.Unmarshal(bytes, &uqa)
    return uqa["qa"]
}

// SliceQA is a type, which is a slice of QA structures.
type SliceQA []qa_db.QA

// UnmarshalledQAs is a type, which is a map whose keys are strings and whose values are SliceQA.
type UnmarshalledQAs map[string]SliceQA

// ParseQuestionsFromJson parses a SliceQA from a JSON response.
func ParseQuestionsFromJson(json1 string) SliceQA {
    bytes := []byte(json1)
    var uqas UnmarshalledQAs
    json.Unmarshal(bytes, &uqas)
    return uqas["qas"]
}

// UnmarshalErr parses an error from a JSON response.
func UnmarshalErr(json1 string) error {
    bytes := []byte(json1)
    var error_map map[string]string
    json.Unmarshal(bytes, &error_map)
    if err, ok := error_map["err"]; ok {
        return fmt.Errorf("REMOTE ERROR: %", err)
    } 
    return nil
}
