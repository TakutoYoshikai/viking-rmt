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

  if !requests.Transfer("person2", "password2", request.Id) {
    t.Error("Transferが失敗した")
  }
  t.Log("Requests終了")
}
