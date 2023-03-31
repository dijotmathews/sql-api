package main

import (
	"github.com/sql-api/initializers"
	"gorm.io/gorm"
)

func init() {
	initializers.LoadEnv()
	initializers.Init()
}

func main() {
	db := initializers.GetDB()
	seedLoanTypes(db)
}

func seedLoanTypes(db *gorm.DB) bool {

	// module.exports = {
	// 	9: '72001',
	// 	10: '72001',
	// 	11: '72003',
	// 	12: '72004',
	// 	13: '72005',
	// 	14: '72005',
	// 	15: '72006',
	// 	16: '72006',
	// 	19: '72001',
	// 	20: '72006',
	//   };

	return false
}
