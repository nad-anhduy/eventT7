package transport

import (
	"eventT7/common"
	"eventT7/modules/Event/biz"
	"eventT7/modules/Event/model"
	"eventT7/modules/Event/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
	"time"
)

func HanlderCreateRecord(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := uuid.New().String()

		var newRequest model.RequestStruct

		if err := ctx.BindJSON(&newRequest); err != nil {
			log.Printf(`SoulT7 - [%s] - Parse JSON fail with err: %s`, id, err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Undefined variable", "status": false, "data": nil})
			return
		}
		log.Printf(`SoulT7 - [%s] - Raw Request: %v`, id, newRequest)

		if err := newRequest.ImgUrl.Valid(); err != nil {
			log.Printf(`SoulT7 - [%s] - Validated ImgUrl fail with err: %s`, id, err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Img url " + err.Error(), "message": "Img url is NULL", "status": false, "data": nil})
			return
		}

		if err := newRequest.UserName.Valid(); err != nil {
			log.Printf(`SoulT7 - [%s] - Validated UserName fail with err: %s`, id, err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "User name " + err.Error(), "message": "User name is NULL", "status": false, "data": nil})
			return
		}

		if err := newRequest.GroupID.Valid(); err != nil {
			log.Printf(`SoulT7 - [%s] - Validated GroupID fail with err: %s`, id, err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Group ID " + err.Error(), "message": "Group ID is NULL", "status": false, "data": nil})
			return
		}

		if err := biz.ProxyFilterByTimeRequest(id, newRequest.GroupID.String()); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Too many request", "status": false, "data": nil})
			return
		}

		if err := common.Validates(newRequest.ExtraData); err != nil {
			log.Printf(`SoulT7 - [%s] - Validated ExtraData fail with err: %s`, id, err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Extra Data " + err.Error(), "message": "Extra Data is NULL", "status": false, "data": nil})
			return
		}

		if err := newRequest.Session.Valid(); err != nil {
			log.Printf(`SoulT7 - [%s] - Validated Session fail with err: %s`, id, err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Session " + err.Error(), "message": "Session is NULL", "status": false, "data": nil})
			return
		}

		_, ok := common.CheckContaint(newRequest.UserName.String(), newRequest.ExtraData)
		if ok {
			log.Printf(`SoulT7 - [%s] - UserName containt in ExtraData: %s`, id)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username contain in ExtraData ", "message": "Bạn không thể tự vote cho chính mình !!! ", "status": false, "data": nil})
			return
		}

		_, ok = common.CheckContaint(newRequest.GroupID.String(), model.ListGroupDeny)
		if ok {
			log.Printf(`SoulT7 - [%s] - Group ID containt in deny list: %s`, id)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Group ID is deny ", "message": "Group ID đã bị hạn chế !!! ", "status": false, "data": nil})
			return
		}

		store := storage.NewSQLStore(db)
		bizz := biz.NewCreateRecord(store)

		data := model.InsertStructToDB{
			Created:  time.Now().UnixMilli(),
			Uid:      id,
			Session:  newRequest.Session.String(),
			UserName: newRequest.UserName.String(),
			GroupID:  newRequest.GroupID.String(),
			ImgUrl:   newRequest.ImgUrl.String(),
			ExtraX:   strings.Join(newRequest.ExtraData, ","),
		}

		if err := bizz.CreateNewRecord(ctx, data); err != nil {
			log.Printf(`SoulT7 - [%s] - Create record fail with err: %s`, id, err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Increase vote fail", "status": false, "data": nil})
			return
		}
		log.Printf(`SoulT7 - [%s] - Request call increase app`, id)
		biz.CallAppIncrease(id, newRequest.GroupID.String())

		ctx.JSON(http.StatusOK, gin.H{"error": "", "message": "Increase vote success", "status": true, "data": nil})
		return
	}
}
