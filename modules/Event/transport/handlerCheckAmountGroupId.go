package transport

import (
	"eventT7/modules/Event/biz"
	"eventT7/modules/Event/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func HandlerCheckAmountGroupId(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		id := uuid.New().String()
		log.Printf(`Soul T7 - [%s] - Process check vote`, id)

		store := storage.NewSQLStore(db)
		bizz := biz.NewCountGroupIdRecord(store)
		log.Printf(`Soul T7 - [%s] - Query data`, id)
		resp, err := bizz.CountNewGroupIdRecord()
		if err != nil {
			log.Printf(`Soul T7 - [%s] - Query fail with error: %v`, id, err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Count fail ", "status": false, "data": nil})
			return
		}
		log.Printf(`Soul T7 - [%s] - Query data success`, id)
		ctx.JSON(http.StatusOK, gin.H{"error": "", "message": "success", "status": true, "data": resp})
		return
	}
}
