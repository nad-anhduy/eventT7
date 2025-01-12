package transport

import (
	"eventT7/modules/Event/biz"
	"eventT7/modules/Event/model"
	"eventT7/modules/Event/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func HandlerCheckTop(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		var listTop3Vote []model.ResponseTopVoteStruct

		id := uuid.New().String()
		log.Printf(`Soul T7 - [%s] - Process check vote`, id)

		store := storage.NewSQLStore(db)
		bizz := biz.NewTopVotePerSession(store)
		log.Printf(`Soul T7 - [%s] - Query data`, id)
		sessions, err := bizz.NewListSession()
		if err != nil {
			log.Printf(`Soul T7 - [%s] - Query fail with error: %v`, id, err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "List session fail", "status": false, "data": nil})
			return
		}
		for _, session := range sessions {
			resultPerSession, err := bizz.TopVoteNewSession(session)
			if err == nil {
				listTop3Vote = append(listTop3Vote, model.ResponseTopVoteStruct{
					Session: session,
					Items:   resultPerSession,
				})
			}
		}
		log.Printf(`Soul T7 - [%s] - Query data success`, id)
		ctx.JSON(http.StatusOK, gin.H{"error": "", "message": "success", "status": true, "data": listTop3Vote})
		return
	}
}
