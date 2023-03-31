package main

import (
	"github.com/sql-api/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.Init()
}

func main() {

}
