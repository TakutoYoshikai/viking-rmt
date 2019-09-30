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

func (transferRequest *TransferRequest) Fetch() {
  url := config.BankHost + "/requests/show/" + config.BankUsername + "/" + config.BankPassword + "/" + strconv.Itoa(transferRequest.Id)
  res, err := http.Get(url)
  if err != nil || res.StatusCode != 200 {
    return
  }
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return
  }
  jsonBytes := ([]byte)(body)
  request := new(TransferRequest)
  err = json.Unmarshal(jsonBytes, request)
  if err != nil {
    return
  }
  transferRequest.Id = request.Id
  transferRequest.From = request.From
  transferRequest.To = request.To
  transferRequest.Amount = request.Amount
  transferRequest.Transfered = request.Transfered
}
