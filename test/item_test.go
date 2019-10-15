package test

import (
	"testing"
	"viking-rmt/model"
)

func TestItem(t *testing.T) {
	item := model.AddItem(1, "ownerbank", "ownergame", "item1", 100, 5)
	if item == nil {
		t.Error("itemが作れなかった")
	}
	if item.GameItemId != 1 {
		t.Error("itemのgameitemidが設定されていない")
	}
	if item.OwnerBankUsername != "ownerbank" {
		t.Error("itemのowner bank usernameが設定されていない")
	}
	if item.OwnerGameUsername != "ownergame" {
		t.Error("itemのowner game usernameが設定されていない")
	}
	if item.Name != "item1" {
		t.Error("itemのnameが設定されていない")
	}
	if item.Price != 100 {
		t.Error("itemのpriceが設定されていない")
	}
	if item.Rarity != 5 {
		t.Error("itemのrarityが設定されていない")
	}
	item.Sale()
	if item.Status != model.ItemStatusSale {
		t.Error("itemのstatusがsaleにならなかった")
	}
	item.Ordered()
	if item.Status != model.ItemStatusOrdered {
		t.Error("itemのstatusがorderedにならなかった")
	}
	item.SentItem()
	if item.Status != model.ItemStatusSentItem {
		t.Error("itemのstatusがsent itemにならなかった")
	}
	item.Completed()
	if item.Status != model.ItemStatusCompleted {
		t.Error("itemのstatusがcompletedにならなかった")
	}
	t.Log("Item終了")
}
