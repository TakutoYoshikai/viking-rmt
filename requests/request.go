package requests

import (
  "net/http"
  "strconv"
  "io/ioutil"
  "viking-rmt/model"
  "encoding/json"
)

const host string = "http://localhost:8081"

func Transfer(username string, password string, requestId int) bool {
  url := host + "/requests/transfer/" + username + "/" + password + "/" + strconv.Itoa(requestId)
  res, err := http.Get(url)
  defer res.Body.Close()
  if err != nil || res.StatusCode != 200 {
    return false
  }
  return true
}

func CreateTransferRequest(username string, password string, to string, amount int) *model.TransferRequest {
  url := host + "/requests/create/" + username + "/" + password + "/" + to + "/" + strconv.Itoa(amount)
  res, err := http.Get(url)
  defer res.Body.Close()
  if err != nil || res.StatusCode != 200 {
    return nil
  }
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return nil
  }
  jsonBytes := ([]byte)(body)
  request := new(model.TransferRequest)
  err = json.Unmarshal(jsonBytes, request)
  if err != nil {
    return nil
  }
  return request
}
