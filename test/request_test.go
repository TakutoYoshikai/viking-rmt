package test

import (
  "testing"
  "viking-rmt/requests"
)

func TestRequests(t *testing.T) {
  request := requests.CreateTransferRequest("person1", "password1", "person2", 500)
  if request == nil {
    t.Error("create transfer requestが失敗した")
  }
  bankAccount := requests.TransferWithRequest("person2", "password2", request.Id)
  if bankAccount == nil {
    t.Error("TransferWithRequestが失敗した")
  }
  bankAccount = requests.Transfer("person2", "password2", "person1", 100)
  if bankAccount == nil {
    t.Error("Transferが失敗した")
  }
  request = requests.GetTransferRequest("person1", "password1", request.Id)
  if request == nil {
    t.Error("get transfer requestが失敗した")
  }

  items := requests.GetGameItems("player1")
  if items == nil {
    t.Error("GetItemsが失敗した")
  }
  item := requests.GetGameItem("player1", items[0].Id)
  if item == nil {
    t.Error("GetItemが失敗した")
  }
  t.Log("Requests終了")
}
