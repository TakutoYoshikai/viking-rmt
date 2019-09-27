package main

import (
  "github.com/gin-gonic/gin"
  "strconv"
  "viking-rmt/model"
  "viking-rmt/requests"
)

func CreateServer() *gin.Engine {
  router := gin.Default()
  router.GET("/items/create/:bank_username/:game_username/:game_item_id", func (ctx *gin.Context) {
    bankUsername := ctx.Param("bank_username")
    gameUsername := ctx.Param("game_username")
    gameItemIdStr := ctx.Param("game_item_id")
    gameItemId, err := strconv.Atoi(gameItemIdStr)
    if err != nil {
      ctx.JSON(400, nil)
      return
    }
    gameItem := requests.GetGameItem(gameUsername, gameItemId)
    if gameItem == nil {
      ctx.JSON(400, nil)
      return
    }
    item := model.AddItem(gameItem.Id, bankUsername, gameItem.Name, gameItem.Rarity)
    ctx.JSON(200, item)
  })
  return router
}
