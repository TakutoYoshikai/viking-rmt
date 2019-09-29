package model
var NewestItemId int = 0

type Item struct {
  Id int
  ItemId int
  OwnerBankUsername string
  Rarity int
  Name string
  IsBought bool
  TransferRequest *TransferRequest
  BuyerGameUsername string
}

type Items []*Item

func NewItem(itemId int, ownerBankUserName string, name string, rarity int) *Item {
  NewestItemId += 1
  return &Item {
    Id: NewestItemId,
    ItemId: itemId,
    OwnerBankUsername: ownerBankUserName,
    Rarity: rarity,
    Name: name,
    IsBought: false,
  }
}

var items Items = Items{}

func AddItem(itemId int, ownerBankUsername string, name string, rarity int) *Item {
  result := NewItem(itemId, ownerBankUsername, name, rarity)
  items = append(items, result)
  return result
}

func GetItem(id int) *Item {
  for _, item := range items {
    if item.Id == id {
      return item
    }
  }
  return nil
}

