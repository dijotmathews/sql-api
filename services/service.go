package services

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/sql-api/dto"
	"github.com/sql-api/initializers"
)

func QueryService(c *gin.Context) ([]byte, int64, error) {

	// *******************************************
	// Getting the request body
	var qb dto.QueryRequestBody
	if err := c.ShouldBindJSON(&qb); err != nil {
		return []byte(""), 0, err
	}

	var result []map[string]interface{}

	// *******************************************
	//Initializing the DB
	db := initializers.GetDB()

	ret := db.Raw(qb.Query).Find(&result)
	if ret.Error != nil {

		return []byte(""), 0, ret.Error
	}
	bytes, _ := json.Marshal(result)
	return bytes, ret.RowsAffected, nil
}
