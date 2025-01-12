package transport

import (
	"eventT7/common"
	"eventT7/modules/Event/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type denyStruct struct {
	GroupID []string `json:"group_id"`
}

func HandlerDenyGroupId() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		id := uuid.New().String()
		var newRequest denyStruct
		if err := ctx.BindJSON(&newRequest); err != nil {
			log.Printf(`SoulT7 - [%s] - Parse JSON fail with err: %s`, id, err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Undefined variable", "status": false, "data": nil})
			return
		}

		if err := common.Validates(newRequest.GroupID); err != nil {
			log.Printf(`SoulT7 - [%s] - Validated GroupID fail with err: %s`, id, err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Group ID " + err.Error(), "message": "Group ID is NULL", "status": false, "data": nil})
			return
		}

		for _, v := range newRequest.GroupID {
			model.ListGroupDeny = append(model.ListGroupDeny, v)
		}
		log.Printf(`Soul T7 - [%s] - Add group id in blacklist success`, id)
		ctx.JSON(http.StatusOK, gin.H{"error": "", "message": "success", "status": true, "data": nil})
		return

	}
}
