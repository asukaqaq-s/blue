package main

import (
	"fim_server/core"
	"fim_server/fim_chat/chat_models"
	"fim_server/fim_file/file_model"
	"fim_server/fim_group/group_models"
	"fim_server/fim_logs/logs_model"
	"fim_server/fim_settings/settings_model"
	"fim_server/fim_user/user_models"
	"flag"
	"fmt"
)

type Options struct {
	DB bool
}

func main() {

	var opt Options
	flag.BoolVar(&opt.DB, "db", false, "db")
	flag.Parse()

	if opt.DB {
		db := core.InitGorm("root:root@tcp(127.0.0.1:3306)/fim_server_db?charset=utf8mb4&parseTime=True&loc=Local")
		err := db.AutoMigrate(
			&user_models.UserModel{},                // 用户表
			&user_models.FriendModel{},              // 好友表
			&user_models.FriendVerifyModel{},        // 好友验证表
			&user_models.UserConfModel{},            // 用户配置表
			&chat_models.ChatModel{},                // 对话表
			&chat_models.TopUserModel{},             // 置顶用户表
			&chat_models.UserChatDeleteModel{},      // 用户删除聊天记录表
			&group_models.GroupModel{},              // 群组表
			&group_models.GroupMemberModel{},        // 群成员表
			&group_models.GroupMsgModel{},           // 群消息表
			&group_models.GroupVerifyModel{},        // 群验证表
			&group_models.GroupUserMsgDeleteModel{}, // 用户删除聊天记录表
			&group_models.GroupUserTopModel{},       // 用户置顶群聊表
			&file_model.FileModel{},                 // 文件表
			&logs_model.LogModel{},                  // 日志表
			&settings_model.SettingsModel{},         // 系统表
		)
		if err != nil {
			fmt.Println("表结构生成失败", err)
			return
		}
		fmt.Println("表结构生成成功！")

	}

}
