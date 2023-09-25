package main

import (
	_ "file-system/internal/database"
	"file-system/internal/global"
	"file-system/internal/web/router"
	"strconv"
)

func main() {
	r := router.NewRouter()

	r.Run(global.WebAddr + ":" + strconv.Itoa(global.WebPort))
}
