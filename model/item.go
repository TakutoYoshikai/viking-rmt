package model
var newestItemId int = 0

type Item struct {
  Id int
  GameItemId int
  OwnerBankUsername string
  Rarity int
  Name string
  IsBought bool
  Price uint64
  TransferRequest *TransferRequest
  BuyerGameUsername string
}

type Items map[int]*Item

func NewItem(itemId int, ownerBankUserName string, name string, price uint64, rarity int) *Item {
  newestItemId += 1
  return &Item {
    Id: newestItemId,
    GameItemId: itemId,
    OwnerBankUsername: ownerBankUserName,
    Rarity: rarity,
    Name: name,
    Price: price,
    IsBought: false,
  }
}

var items Items = Items{}

func AddItem(itemId int, ownerBankUsername string, name string, price uint64, rarity int) *Item {
  result := NewItem(itemId, ownerBankUsername, name, price, rarity)
  items[result.Id] = result
  return result
}

func GetItem(id int) *Item {
  return items[id]
}

func ItemCount() int {
  return newestItemId
}
