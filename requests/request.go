package requests

import (
  "net/http"
  "strconv"
  "io/ioutil"
  "viking-rmt/model"
  "encoding/json"
)

const bankHost string = "http://localhost:8081"
const gameHost string = "http://localhost:8080"

func Transfer(username string, password string, requestId int) *model.BankAccount {
  url := bankHost + "/requests/transfer/" + username + "/" + password + "/" + strconv.Itoa(requestId)
  res, err := http.Get(url)
  if err != nil || res.StatusCode != 200 {
    return nil
  }
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return nil
  }
  jsonBytes := ([]byte)(body)
  bankAccount := new(model.BankAccount)
  err = json.Unmarshal(jsonBytes, bankAccount)
  if err != nil {
    return nil
  }
  return bankAccount
}

func CreateTransferRequest(username string, password string, to string, amount int) *model.TransferRequest {
  url := bankHost + "/requests/create/" + username + "/" + password + "/" + to + "/" + strconv.Itoa(amount)
  res, err := http.Get(url)
  if err != nil || res.StatusCode != 200 {
    return nil
  }
  defer res.Body.Close()
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

func GetTransferRequest(username string, password string, requestId int) *model.TransferRequest {
  url := bankHost + "/requests/show/" + username + "/" + password + "/" + strconv.Itoa(requestId)
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

func GetGameItem(username string, itemId int) *model.GameItem {
  url := gameHost + "/item/" + username + "/" + strconv.Itoa(itemId)
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
  item := new(model.GameItem)
  err = json.Unmarshal(jsonBytes, item)
  if err != nil {
    return nil
  }
  return item
}

func GetGameItems(username string) []model.GameItem {
  url := gameHost + "/items/" + username
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
  items := new([]model.GameItem)
  err = json.Unmarshal(jsonBytes, items)
  if err != nil {
    return nil
  }
  return *items
}
