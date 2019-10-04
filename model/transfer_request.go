package model

import (
  "net/http"
  "io/ioutil"
  "strconv"
  "encoding/json"
  "viking-rmt/config"
)

type TransferRequest struct {
  Id int
  From string
  To string
  Amount int
  Transfered bool
}

func (transferRequest *TransferRequest) Fetch() bool {
  url := config.BankHost + "/requests/show/" + config.BankUsername + "/" + config.BankPassword + "/" + strconv.Itoa(transferRequest.Id)
  res, err := http.Get(url)
  if err != nil || res.StatusCode != 200 {
    return false
  }
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return false
  }
  jsonBytes := ([]byte)(body)
  request := new(TransferRequest)
  err = json.Unmarshal(jsonBytes, request)
  if err != nil {
    return false
  }
  transferRequest.Id = request.Id
  transferRequest.From = request.From
  transferRequest.To = request.To
  transferRequest.Amount = request.Amount
  transferRequest.Transfered = request.Transfered
  return true
}
