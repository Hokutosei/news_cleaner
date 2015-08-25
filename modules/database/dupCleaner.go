package database

import (
	mongodb "gopkg.in/mgo.v2"
	"news_cleaner/modules/utils"
  "gopkg.in/mgo.v2/bson"
)

type DuppedItems struct {
  ID bson.
}

// GetDups query for time Duration of dupped data
func GetDups(from time.Duration, to time.Duration) []string {
	start := time.Now()

	sc := SessionCopy()
	c := sc.DB(Db).C(NewsMainCollection)
	defer sc.Close()

	gte, lte := utils.TimeRange(from, to)
	fmt.Println("query for this times gte: ", gte, " lte: ", lte)

	// get dups
	// db.news_main.aggregate([{$group: {"_id": "$title", total: {$sum:1},items:{$push:"$_id"}}},{$match:{"total":{$gte:2}}}])
}
