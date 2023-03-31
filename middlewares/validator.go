package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Validate(inputName string) gin.HandlerFunc {
	return func(c *gin.Context) {

		var validationResult bool
		// var validationMessage string

		switch inputName {
		case "apply":
			// validationResult = validateApply(c)
			if validationResult == true {
				c.Next()
			}
		default:
			c.Next()
		}
		c.Abort()
	}

}

// func validateApply(c *gin.Context) bool {
// 	// fmt.Println("We are inside the validator for apply")
// 	ByteBody, _ := ioutil.ReadAll(c.Request.Body)
// 	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(ByteBody))
// 	var json dto.Apply
// 	if err := c.ShouldBindJSON(&json); err != nil {

// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return false
// 	}

// 	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(ByteBody))

// 	return true
// }
