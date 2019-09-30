package model
var newestItemId int = 0

const (
  ItemStatusSale = iota
  ItemStatusOrdered
  ItemStatusSentItem
  ItemStatusCompleted
)

type Item struct {
  Id int
  GameItemId int
  OwnerBankUsername string
  Rarity int
  Name string
  Status int
  Price uint64
  TransferRequest *TransferRequest
  BuyerGameUsername string
}

type Items map[int]*Item

func (item *Item) Sale() {
  item.Status = ItemStatusSale
}

func (item *Item) Ordered() {
  item.Status = ItemStatusOrdered
}

func (item *Item) SentItem() {
  item.Status = ItemStatusSentItem
}

func (item *Item) Completed() {
  item.Status = ItemStatusCompleted
}

func NewItem(itemId int, ownerBankUserName string, name string, price uint64, rarity int) *Item {
  newestItemId += 1
  return &Item {
    Id: newestItemId,
    GameItemId: itemId,
    OwnerBankUsername: ownerBankUserName,
    Rarity: rarity,
    Name: name,
    Price: price,
  }
}

var items Items = Items{}

func AddItem(gameItemId int, ownerBankUsername string, name string, price uint64, rarity int) *Item {
  item := GetItemByGameItemId(gameItemId)
  if item != nil && item.Status != ItemStatusCompleted {
    return nil
  }
  result := NewItem(gameItemId, ownerBankUsername, name, price, rarity)
  items[result.Id] = result
  return result
}

func GetItem(id int) *Item {
  return items[id]
}

func GetAllItems() []Item {
  var result []Item = []Item{}
  for _, item := range items {
    result = append(result, *item)
  }
  return result
}

func GetItemByGameItemId(gameItemId int) *Item {
  for _, item := range items {
    if item.GameItemId == gameItemId {
      return item
    }
  }
  return nil
}

func ItemCount() int {
  return newestItemId
}
