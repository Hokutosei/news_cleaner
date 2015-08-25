package dupCleaner

import (
	"fmt"
	"news_cleaner/modules/utils"
	"time"
)

// Start main start cleaner
func Start(loopDelay int) {
	utils.Info(fmt.Sprintf("start cleaning!"))
	for t := range time.Tick(time.Duration(loopDelay) * time.Second) {
		utils.Info(fmt.Sprintf("%v", t))
	}
}
