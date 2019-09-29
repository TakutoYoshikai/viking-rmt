package model

type GameItem struct {
  Id int
  Rarity int
  Name string
}

type GameItems []GameItem

func (gameItems GameItems) GetGameItem(id int) *GameItem {
  for _, gameItem := range gameItems {
    if gameItem.Id == id {
      return &gameItem
    }
  }
  return nil
}
