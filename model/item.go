package model
var newestItemId int = 0

type Item struct {
  Id int
  ItemId int
  OwnerBankUsername string
  Rarity int
  Name string
  IsBought bool
  Price uint64
  TransferRequest *TransferRequest
  BuyerGameUsername string
}

type Items []*Item

func NewItem(itemId int, ownerBankUserName string, name string, price uint64, rarity int) *Item {
  newestItemId += 1
  return &Item {
    Id: newestItemId,
    ItemId: itemId,
    OwnerBankUsername: ownerBankUserName,
    Rarity: rarity,
    Name: name,
    IsBought: false,
  }
}

var items Items = Items{}

func AddItem(itemId int, ownerBankUsername string, name string, price uint64, rarity int) *Item {
  result := NewItem(itemId, ownerBankUsername, name, price, rarity)
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

func ItemCount() int {
  return newestItemId
}
