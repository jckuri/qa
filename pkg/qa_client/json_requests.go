package qa_client

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "bytes"
)

// verbose is a boolean that decides if messages appear in the client or not.
var verbose = false

// JsonRequest is a method to call a JSON request. Its parameter data is very abstract.
func JsonRequest(url string, method string, data interface{}) (string, int) {
    json, err := json.Marshal(data)
    if err != nil {
        panic(err)
    }
    req, err := http.NewRequest(method, url, bytes.NewBuffer(json))
    if err != nil {
        panic(err)
    }
    req.Header.Set("Content-Type", "application/json; charset=utf-8")
    client := &http.Client{}
    resp, err := client.Do(req)
    if verbose {
        fmt.Printf("URL: %v\n", url)
        fmt.Printf("Request: %v\n", req)
        fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    return string(body), resp.StatusCode
}

// JsonMapRequest is a method to call a JSON request. Its parameter values is a map whose keys are strings and whose values are very abstract.
func JsonMapRequest(url string, method string, values map[string]interface{}) (string, int) {
    return JsonRequest(url, method, values)
}
