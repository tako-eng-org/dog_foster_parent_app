package controller

import "fpdapp/models/db"

type (
	Controller struct {
		DbConn *db.Database
	}
)

//*******************************************************************
// 型変換する
//*******************************************************************
func ToInt(i interface{}) int {
	var r, _ = i.(int)
	return r
}
func ToStr(i interface{}) string {
	var r, _ = i.(string)
	return r
}

func ToInt64(i interface{}) uint64 {
	var r, _ = i.(uint64)
	return r
}
