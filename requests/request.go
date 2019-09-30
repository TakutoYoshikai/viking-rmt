package requests

import (
  "net/http"
  "strconv"
  "io/ioutil"
  "viking-rmt/model"
  "viking-rmt/config"
  "encoding/json"
)


func Transfer(username string, password string, to string, amount uint64) *model.BankAccount {
  url := config.BankHost + "/transfer/" + username + "/" + password + "/" + strconv.Itoa(int(amount)) + "/" + to
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
func TransferWithRequest(username string, password string, requestId int) *model.BankAccount {
  url := config.BankHost + "/requests/transfer/" + username + "/" + password + "/" + strconv.Itoa(requestId)
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

func RmtTransfer(to string, amount uint64) *model.BankAccount {
  return Transfer(config.GameUsername, config.GamePassword, to, amount)
}
func RmtTransferWithRequest(requestId int) *model.BankAccount {
  return TransferWithRequest(config.GameUsername, config.GamePassword, requestId)
}

func CreateTransferRequest(username string, password string, to string, amount uint64) *model.TransferRequest {
  url := config.BankHost + "/requests/create/" + username + "/" + password + "/" + to + "/" + strconv.Itoa(int(amount))
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
  url := config.BankHost + "/requests/show/" + username + "/" + password + "/" + strconv.Itoa(requestId)
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
  url := config.GameHost + "/item/" + username + "/" + strconv.Itoa(itemId)
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
  url := config.GameHost + "/items/" + username
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
  items := new([]model.GameItem)
  err = json.Unmarshal(jsonBytes, items)
  if err != nil {
    return nil
  }
  return *items
}

func GetMyGameItems() model.GameItems {
  return GetGameItems(config.GameUsername)
}

func GiveItem(username string, password string, to string, gameItemId int) bool {
  url := config.GameHost + "/send/" + username + "/" + password + "/" + strconv.Itoa(gameItemId) + "/" + to
  res, err := http.Get(url)
  if err != nil || res.StatusCode != 200 {
    return false
  }
  return true
}
