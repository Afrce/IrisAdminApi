package middleware

import (
	"IrisAdminApi/models"
	"encoding/json"
	"github.com/kataras/iris"
	"time"
)

func RecordApiLog(ctx iris.Context) {
	log := models.ApiLog{}
	log.Url = ctx.Path()
	log.UserId, _ = ctx.Values().GetUint("user_id")
	Params := make(map[string]interface{})
	for k, v := range ctx.FormValues() {
		if len(v) > 1 {
			Params[k] = v
		} else {
			Params[k] = v[0]
		}
	}
	params, _ := json.Marshal(Params)
	log.Params = string(params)
	log.Time = time.Now().Local()

	models.RecordLogs(log)
	ctx.Next()
}
