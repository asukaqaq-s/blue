package main

import (
	"fim_server/common/list_query"
	"fim_server/common/models"
	"fim_server/core"
	"fim_server/fim_chat/chat_models"
	"fmt"
)

func main() {
	db := core.InitGorm("root:root@tcp(127.0.0.1:3306)/fim_server_db?charset=utf8mb4&parseTime=True&loc=Local")

	var userId = 1
	type Data struct {
		SU         uint   `gorm:"column:sU"`
		RU         uint   `gorm:"column:rU"`
		MaxDate    string `gorm:"column:maxDate"`
		MaxPreview string `gorm:"column:maxPreview"`
		IsTop      bool   `gorm:"column:isTop"`
	}
	var list []Data

	db.Table("(?) as u", db.Model(&chat_models.ChatModel{}).
		Select("least(send_user_id, rev_user_id)    as sU",
			"greatest(send_user_id, rev_user_id) as rU",
			" max(created_at)   as maxDate",
			"max(msg_preview) as maxPreview").Where("send_user_id = ? or rev_user_id = ?", userId, userId).
		Group("least(send_user_id, rev_user_id)").
		Group("greatest(send_user_id, rev_user_id)")).
		Order("maxDate desc").Limit(1).Offset(0).Scan(&list)
	fmt.Println(list)

	column := fmt.Sprintf("if((select 1 from top_user_models where user_id = %d and (top_user_id = sU or top_user_id = rU)), 1, 0)  as isTop", userId)

	chatList, count, _ := list_query.ListQuery(db, Data{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  1,
			Limit: 10,
			Sort:  "isTop desc, maxDate desc",
		},
		Table: func() (string, any) {
			return "(?) as u", db.Model(&chat_models.ChatModel{}).
				Select("least(send_user_id, rev_user_id) as sU",
					"greatest(send_user_id, rev_user_id) as rU",
					"max(created_at) as maxDate",
					"(select msg_preview from chat_models  where (send_user_id = sU and rev_user_id = rU) or (send_user_id = rU and rev_user_id = sU) order by created_at desc  limit 1) as maxPreview",
					column).
				Where("send_user_id = ? or rev_user_id = ?", userId, userId).
				Group("least(send_user_id, rev_user_id)").
				Group("greatest(send_user_id, rev_user_id)")
		},
	})

	fmt.Println(chatList, count)
	for _, data := range chatList {
		fmt.Println(data.IsTop, data.MaxPreview, data.MaxDate)
	}

}
