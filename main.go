package main

import (
	"fmt"
	"news_cleaner/modules/database"
	"news_cleaner/modules/dupCleaner"
	"news_cleaner/modules/utils"
)

var (
	loopDelay = 10
)

func main() {
	utils.Info(fmt.Sprintf("cleaner started!"))
	go database.MongodbStart()
	dupCleaner.Start(loopDelay)
}
