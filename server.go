package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"viking-rmt/config"
	"viking-rmt/model"
	"viking-rmt/requests"
)

func CreateServer() *gin.Engine {
	router := gin.Default()
	router.GET("/item/transfered/:item_id", func(ctx *gin.Context) {
		itemIdStr := ctx.Param("item_id")
		itemId, err := strconv.Atoi(itemIdStr)
		if err != nil {
			ctx.JSON(400, nil)
			return
		}
		item := model.GetItem(itemId)
		if item.TransferRequest == nil {
			ctx.JSON(400, nil)
			return
		}
		if !item.TransferRequest.Fetch() {
			ctx.JSON(500, nil)
			return
		}
		if !item.TransferRequest.Transfered {
			ctx.JSON(400, nil)
			return
		}
		item.Transfered()
		ctx.JSON(200, nil)
	})
	router.GET("/item/sent/:item_id", func(ctx *gin.Context) {
		itemIdStr := ctx.Param("item_id")
		itemId, err := strconv.Atoi(itemIdStr)
		if err != nil {
			ctx.JSON(400, nil)
			return
		}
		item := model.GetItem(itemId)
		if item == nil {
			ctx.JSON(404, nil)
			return
		}
		if item.Status != model.ItemStatusTransfered {
			ctx.JSON(400, nil)
			return
		}
		if item.TransferRequest == nil {
			ctx.JSON(400, nil)
			return
		}
		item.TransferRequest.Fetch()
		if !item.TransferRequest.Transfered {
			ctx.JSON(400, nil)
		}
		myGameItems := requests.GetMyGameItems()
		gameItem := myGameItems.GetGameItem(item.GameItemId)
		if gameItem == nil {
			ctx.JSON(400, nil)
			return
		}
		item.SentItem()
		account := requests.Transfer("rmt", "rmt", item.OwnerBankUsername, uint64(float64(item.TransferRequest.Amount)*0.9))
		if account == nil {
			ctx.JSON(500, nil)
			return
		}
		success := requests.GiveItem("rmt", "rmt", item.BuyerGameUsername, gameItem.Id)
		if !success {
			ctx.JSON(500, nil)
			return
		}
		item.Completed()
		ctx.JSON(200, nil)
	})
	router.GET("/item/buy/:item_id/:bank_username/:game_username", func(ctx *gin.Context) {
		bankUsername := ctx.Param("bank_username")
		gameUsername := ctx.Param("game_username")
		itemIdStr := ctx.Param("item_id")
		itemId, err := strconv.Atoi(itemIdStr)
		if err != nil {
			ctx.JSON(400, nil)
			return
		}
		item := model.GetItem(itemId)
		if item == nil {
			ctx.JSON(404, nil)
			return
		}
		if item.Status != model.ItemStatusSale {
			ctx.JSON(400, nil)
			return
		}
		transferRequest := requests.CreateTransferRequest(config.BankUsername, config.BankPassword, bankUsername, item.Price)
		if transferRequest == nil {
			ctx.JSON(500, nil)
			return
		}
		item.TransferRequest = transferRequest
		item.BuyerGameUsername = gameUsername
		item.Ordered()
		ctx.JSON(200, transferRequest)
	})
	router.GET("/items", func(ctx *gin.Context) {
		ctx.JSON(200, model.GetAllItems())
	})
	router.GET("/item/show/:id", func(ctx *gin.Context) {
		itemIdStr := ctx.Param("id")
		itemId, err := strconv.Atoi(itemIdStr)
		if err != nil {
			ctx.JSON(400, nil)
			return
		}
		item := model.GetItem(itemId)
		if item == nil {
			ctx.JSON(404, nil)
		}
		ctx.JSON(200, item)
	})
	router.GET("/item/create/:bank_username/:game_username/:game_item_id/:price", func(ctx *gin.Context) {
		bankUsername := ctx.Param("bank_username")
		gameUsername := ctx.Param("game_username")
		gameItemIdStr := ctx.Param("game_item_id")
		gameItemId, err := strconv.Atoi(gameItemIdStr)
		if err != nil {
			ctx.JSON(400, nil)
			return
		}
		priceStr := ctx.Param("price")
		price, err := strconv.ParseUint(priceStr, 10, 64)
		if err != nil {
			ctx.JSON(400, nil)
			return
		}
		gameItem := requests.GetGameItem(gameUsername, gameItemId)
		if gameItem == nil {
			ctx.JSON(400, nil)
			return
		}
		item := model.AddItem(gameItem.Id, bankUsername, gameUsername, gameItem.Name, price, gameItem.Rarity)
		if item == nil {
			ctx.JSON(400, nil)
			return
		}
		item.Sale()
		ctx.JSON(200, item)
	})
	return router
}
