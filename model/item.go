package model
var NewestItemId int = 0

type Item struct {
  Id int
  ItemId int
  Owner string
  IsBought bool
}

func NewItem(itemId int, owner string) Item {
  NewestItemId += 1
  return Item {
    Id: NewestItemId,
    ItemId: itemId,
    Owner: owner,
    IsBought: false,
  }
}

var Items []Item = []Item{}

func AddItem(itemId int, owner string) *Item {
  result := NewItem(itemId, owner)
  Items = append(Items, result)
  return &result
}

func GetItem(id int) *Item {
  for _, item := range Items {
    if item.Id == id {
      return &item
    }
  }
  return nil
}

func getIndexOfItem(id int) int {
  for index, item := range Items {
    if item.Id == id {
      return index
    }
  }
  return -1
}

func Bought(id int) {
  index := getIndexOfItem(id)
  if index == -1 {
    return
  }
  Items[index].IsBought = true
}

func IsBought(id int) bool {
  index := getIndexOfItem(id)
  if index == -1 {
    return false
  }
  return Items[index].IsBought
}
