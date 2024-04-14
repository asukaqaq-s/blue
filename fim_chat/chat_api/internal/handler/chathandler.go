package handler

import (
	"context"
	"encoding/json"
	"fim_server/common/models/ctype"
	"fim_server/common/response"
	"fim_server/common/service/redis_service"
	"fim_server/fim_chat/chat_api/internal/svc"
	"fim_server/fim_chat/chat_api/internal/types"
	"fim_server/fim_chat/chat_models"
	"fim_server/fim_file/file_rpc/files"
	"fim_server/fim_user/user_models"
	"fim_server/fim_user/user_rpc/types/user_rpc"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type UserWsInfo struct {
	UserInfo    user_models.UserModel      // 用户信息
	WsClientMap map[string]*websocket.Conn // 这个用户管理的所有ws客户端
	CurrentConn *websocket.Conn            // 当前的连接对象
}

var UserOnlineWsMap = map[uint]*UserWsInfo{}
var VideoCallMap = map[string]time.Time{} // 用户对话的起始时间

func chatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatRequest
		if err := httpx.ParseHeaders(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		var upGrader = websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// 鉴权 true表示放行，false表示拦截
				return true
			},
		}

		conn, err := upGrader.Upgrade(w, r, nil)
		if err != nil {
			logx.Error(err)
			response.Response(r, w, nil, err)
			return
		}

		addr := conn.RemoteAddr().String()
		defer func() {
			conn.Close()

			userWsInfo, ok := UserOnlineWsMap[req.UserID]
			if ok {
				// 删除的退出的那个ws信息
				delete(userWsInfo.WsClientMap, addr)
			}
			if userWsInfo != nil && len(userWsInfo.WsClientMap) == 0 {
				// 全退完了
				delete(UserOnlineWsMap, req.UserID)
				svcCtx.Redis.HDel("online", fmt.Sprintf("%d", req.UserID))
			}
		}()
		// 调用户服务，获取当前用户信息
		res, err := svcCtx.UserRpc.UserInfo(context.Background(), &user_rpc.UserInfoRequest{
			UserId: uint32(req.UserID),
		})
		if err != nil {
			logx.Error(err)
			response.Response(r, w, nil, err)
			return
		}

		var userInfo user_models.UserModel
		err = json.Unmarshal(res.Data, &userInfo)
		if err != nil {
			logx.Error(err)
			response.Response(r, w, nil, err)
			return
		}
		userWsInfo, ok := UserOnlineWsMap[req.UserID]
		if !ok {
			userWsInfo = &UserWsInfo{
				UserInfo: userInfo,
				WsClientMap: map[string]*websocket.Conn{
					addr: conn,
				},
				CurrentConn: conn, // 当前的连接对象
			}
			// 代表这个用户第一次来
			UserOnlineWsMap[req.UserID] = userWsInfo
		}
		_, ok1 := userWsInfo.WsClientMap[addr]
		if !ok1 {
			// 代表这个用户二开及以上
			UserOnlineWsMap[req.UserID].WsClientMap[addr] = conn
			// 把当前连接对象更换
			UserOnlineWsMap[req.UserID].CurrentConn = conn
		}

		// 把在线的用户存进redis
		svcCtx.Redis.HSet("online", fmt.Sprintf("%d", req.UserID), req.UserID)
		// 遍历在线的用户， 和当前这个人是好友的，就给他发送好友在线

		// 先把所有在线的用户id取出来，以及待确认的用户id，然后传到用户rpc服务中
		// [1,2,3]  3
		// 在rpc服务中，去判断哪些用户是好友关系

		//if userInfo.UserConfModel.FriendOnline {
		// 如果用户开启了好友上线提醒
		// 查一下自己的好友是不是上线了
		friendRes, err := svcCtx.UserRpc.FriendList(context.Background(), &user_rpc.FriendListRequest{
			User: uint32(req.UserID),
		})
		// 3 [3,4,5]
		if err != nil {
			logx.Error(err)
			response.Response(r, w, nil, err)
			return
		}
		logx.Infof("用户上线：%s 用户id: %d", userInfo.Nickname, req.UserID)

		for _, info := range friendRes.FriendList {
			friend, ok := UserOnlineWsMap[uint(info.UserId)]
			if ok {
				text := fmt.Sprintf("好友 %s 上线了", UserOnlineWsMap[req.UserID].UserInfo.Nickname)
				// 判断用户是否开了好友上线提醒
				if friend.UserInfo.UserConfModel.FriendOnline {
					// 好友上线了
					//friend.Conn.WriteMessage(websocket.TextMessage, []byte(text))
					sendWsMapMsg(friend.WsClientMap, []byte(text))
				}
			}
		}
		// 查一下自己的好友列表，返回用户id列表，看看在不在这个UserWsMap中，在的话，就给自己发个好友上线的消息
		//}
		for {
			// 消息类型，消息，错误
			_, p, err1 := conn.ReadMessage()
			if err1 != nil {
				// 用户断开聊天
				fmt.Println(err1)
				break
			}
			// 目前这里做不到实时更新
			// 要做到实时更新，把用户的这些配置放到缓存里面去
			// 用户聊天之前就向缓存里面去拿用户的相关配置信息 拿不到的情况下，去调用户rpc方法，然后缓存到缓存里面
			// 在后台，把用户的配置更新之后，让这条缓存失效即可
			if userInfo.UserConfModel.CurtailChat {
				SendTipErrMsg(conn, "当前用户被限制聊天")
				continue
			}

			var request ChatRequest
			err2 := json.Unmarshal(p, &request)
			if err2 != nil {
				// 用户乱发消息
				logx.Error(err2)
				SendTipErrMsg(conn, "参数解析失败")
				continue
			}
			if request.RevUserID != req.UserID {
				// 判断你聊天的这个人是不是你的好友
				isFriendRes, err := svcCtx.UserRpc.IsFriend(context.Background(), &user_rpc.IsFriendRequest{
					User1: uint32(req.UserID),
					User2: uint32(request.RevUserID),
				})
				if err != nil {
					// 用户乱发消息
					logx.Error(err2)
					SendTipErrMsg(conn, "用户服务错误")
					continue
				}

				if !isFriendRes.IsFriend {
					SendTipErrMsg(conn, "你们还不是好友呢")
					continue
				}
			}
			// 判断type  1 - 12
			if !(request.Msg.Type >= 1 && request.Msg.Type <= 12) {
				SendTipErrMsg(conn, "消息类型错误")
				continue
			}
			// 判断是否是文件类型
			switch request.Msg.Type {
			case ctype.TextMsgType:

				if request.Msg.TextMsg == nil {
					SendTipErrMsg(conn, "请输入消息内容")
					continue
				}

				if request.Msg.TextMsg.Content == "" {
					SendTipErrMsg(conn, "请输入消息内容")
					continue
				}

			case ctype.FileMsgType:
				if request.Msg.FileMsg == nil {
					SendTipErrMsg(conn, "请上传文件")
					return
				}

				// 如果是文件类型，那么就要去请求文件rpc服务
				nameList := strings.Split(request.Msg.FileMsg.Src, "/")
				if len(nameList) == 0 {
					SendTipErrMsg(conn, "请上传文件")
					continue
				}
				fileID := nameList[len(nameList)-1]
				fileResponse, err3 := svcCtx.FileRpc.FileInfo(context.Background(), &files.FileInfoRequest{
					FileId: fileID,
				})
				if err3 != nil {
					logx.Error(err3)
					SendTipErrMsg(conn, err3.Error())
					continue
				}
				request.Msg.FileMsg.Title = fileResponse.FileName
				request.Msg.FileMsg.Size = fileResponse.FileSize
				request.Msg.FileMsg.Type = fileResponse.FileType
			case ctype.WithdrawMsgType:
				// 撤回消息的消息id是必填的
				if request.Msg.WithdrawMsg == nil {
					SendTipErrMsg(conn, "撤回消息id必填")
					continue
				}
				if request.Msg.WithdrawMsg.MsgID == 0 {
					SendTipErrMsg(conn, "撤回消息id必填")
					continue
				}

				// 自己只能撤回自己的
				// 找这个消息是谁发的
				var msgModel chat_models.ChatModel
				err = svcCtx.DB.Take(&msgModel, request.Msg.WithdrawMsg.MsgID).Error
				if err != nil {
					SendTipErrMsg(conn, "消息不存在")
					continue
				}

				// 已经是撤回消息的，不能再撤回了
				if msgModel.MsgType == ctype.WithdrawMsgType {
					SendTipErrMsg(conn, "撤回消息不能再撤回了")
					continue
				}

				// 判断是不是自己发的
				if msgModel.SendUserID != req.UserID {
					SendTipErrMsg(conn, "只能撤回自己的消息")
					continue
				}

				// 判断消息的时间，小于2分钟的才能撤回
				now := time.Now()
				subTime := now.Sub(msgModel.CreatedAt)
				if subTime >= time.Minute*2 {
					SendTipErrMsg(conn, "只能撤回两分钟以内的消息哦~")
					continue
				}
				// 撤回逻辑
				// 收到撤回请求之后，服务端这边把原消息类型修改为撤回消息类型，并且记录原消息
				// 然后通知前端的收发双方，重新拉取聊天记录

				var content = "撤回了一条消息"
				if userInfo.UserConfModel.RecallMessage != nil {
					content = *userInfo.UserConfModel.RecallMessage
				}
				content = "你" + content
				// 前端可以判断，这个消息如果不是isMe，就可以把你替换成对方的昵称

				originMsg := msgModel.Msg
				originMsg.WithdrawMsg = nil // 这里可能会出现循环引用，所以拷贝了这个值，并且把撤回消息置空了

				svcCtx.DB.Model(&msgModel).Updates(chat_models.ChatModel{
					MsgPreview: "[撤回消息] - " + content,
					MsgType:    ctype.WithdrawMsgType,
					Msg: ctype.Msg{
						Type: ctype.WithdrawMsgType,
						WithdrawMsg: &ctype.WithdrawMsg{
							Content:   content,
							MsgID:     request.Msg.WithdrawMsg.MsgID,
							OriginMsg: &originMsg,
						},
					},
				})
			case ctype.ReplyMsgType:
				// 回复消息
				// 先校验
				if request.Msg.ReplyMsg == nil || request.Msg.ReplyMsg.MsgID == 0 {
					SendTipErrMsg(conn, "回复消息id必填")
					continue
				}

				// 找这个原消息
				var msgModel chat_models.ChatModel
				err = svcCtx.DB.Take(&msgModel, request.Msg.ReplyMsg.MsgID).Error
				if err != nil {
					SendTipErrMsg(conn, "消息不存在")
					continue
				}

				// 不能回复撤回消息
				if msgModel.MsgType == ctype.WithdrawMsgType {
					SendTipErrMsg(conn, "该消息已撤回")
					continue
				}

				// 回复的这个消息，必须是你自己或者当前和你聊天这个人发出来的

				// 原消息必须是 当前你要和对方聊的  原消息就会有一个 发送人id和接收人id，  我们聊天也会有一个发送人id和接收人id
				// 因为回复消息可以回复自己的，也可以回复别人的
				// 如果回复只能回复别人的？那么条件怎么写?
				if !((msgModel.SendUserID == req.UserID && msgModel.RevUserID == request.RevUserID) ||
					(msgModel.SendUserID == request.RevUserID && msgModel.RevUserID == req.UserID)) {
					SendTipErrMsg(conn, "只能回复自己或者对方的消息")
					continue
				}

				userBaseInfo, err5 := redis_service.GetUserBaseInfo(svcCtx.Redis, svcCtx.UserRpc, msgModel.SendUserID)
				if err5 != nil {
					logx.Error(err5)
					SendTipErrMsg(conn, err5.Error())
					continue
				}

				request.Msg.ReplyMsg.Msg = &msgModel.Msg
				request.Msg.ReplyMsg.UserID = msgModel.SendUserID
				request.Msg.ReplyMsg.UserNickName = userBaseInfo.NickName
				request.Msg.ReplyMsg.OriginMsgDate = msgModel.CreatedAt
			case ctype.QuoteMsgType:
				// 回复消息
				// 先校验
				if request.Msg.QuoteMsg == nil || request.Msg.QuoteMsg.MsgID == 0 {
					SendTipErrMsg(conn, "引用消息id必填")
					continue
				}

				// 找这个原消息
				var msgModel chat_models.ChatModel
				err = svcCtx.DB.Take(&msgModel, request.Msg.QuoteMsg.MsgID).Error
				if err != nil {
					SendTipErrMsg(conn, "消息不存在")
					continue
				}

				// 不能回复撤回消息
				if msgModel.MsgType == ctype.WithdrawMsgType {
					SendTipErrMsg(conn, "该消息已撤回")
					continue
				}

				// 回复的这个消息，必须是你自己或者当前和你聊天这个人发出来的

				// 原消息必须是 当前你要和对方聊的  原消息就会有一个 发送人id和接收人id，  我们聊天也会有一个发送人id和接收人id
				// 因为回复消息可以回复自己的，也可以回复别人的
				// 如果回复只能回复别人的？那么条件怎么写?
				if !((msgModel.SendUserID == req.UserID && msgModel.RevUserID == request.RevUserID) ||
					(msgModel.SendUserID == request.RevUserID && msgModel.RevUserID == req.UserID)) {
					SendTipErrMsg(conn, "只能回复自己或者对方的消息")
					continue
				}

				userBaseInfo, err5 := redis_service.GetUserBaseInfo(svcCtx.Redis, svcCtx.UserRpc, msgModel.SendUserID)
				if err5 != nil {
					logx.Error(err5)
					SendTipErrMsg(conn, err5.Error())
					continue
				}

				request.Msg.QuoteMsg.Msg = &msgModel.Msg
				request.Msg.QuoteMsg.UserID = msgModel.SendUserID
				request.Msg.QuoteMsg.UserNickName = userBaseInfo.NickName
				request.Msg.QuoteMsg.OriginMsgDate = msgModel.CreatedAt
			case ctype.VideoCallMsgType:
				data := request.Msg.VideoCallMsg
				// 先判断对方是否在线
				_, ok2 := UserOnlineWsMap[request.RevUserID]
				if !ok2 {
					SendTipErrMsg(conn, "对方不在线")
					return
				}

				key := fmt.Sprintf("%d_%d", userInfo.ID, request.RevUserID)

				switch data.Flag {
				case 0:
					// 给自己的页面展示一个等待对方接听的一个弹框
					conn.WriteJSON(ChatResponse{
						Msg: ctype.Msg{
							VideoCallMsg: &ctype.VideoCallMsg{
								Flag: 1,
							},
						},
					})
					// 给对方的页面展示一个等待接听的一个弹框
					sendRevUserMsg(request.RevUserID, ctype.Msg{
						VideoCallMsg: &ctype.VideoCallMsg{
							Flag: 2,
						},
					})
				case 1: // 自己挂断
					sendRevUserMsg(request.RevUserID, ctype.Msg{
						VideoCallMsg: &ctype.VideoCallMsg{
							Flag: 3,
							Msg:  "发起者已挂断",
						},
					})
				case 2: // 对方挂断
					// 对方点击挂断，那么它的目标就是revUserID，也就是上面的conn
					sendRevUserMsg(request.RevUserID, ctype.Msg{
						VideoCallMsg: &ctype.VideoCallMsg{
							Flag: 4,
							Msg:  "用户拒绝了你的视频通话",
						},
					})
				case 3: // 对方接受
					// 让发送者准备去发offer
					sendRevUserMsg(request.RevUserID, ctype.Msg{
						VideoCallMsg: &ctype.VideoCallMsg{
							Flag: 5, // 让发送者准备去发offer
							Type: "create_offer",
						},
					})
				case 4: //我方正常挂断
					// 算你们的通话时长
					// 从发offer开始，算一个开始时间，到这里算一个结束时间，就是视频通话的时间
					startTime, ok3 := VideoCallMap[key]
					if !ok3 {
						fmt.Println("没有获取到起始时间")
						continue
					}
					subTime := time.Now().Sub(startTime)
					fmt.Printf("用户正常挂断， 视频通话时长为 %s\n", subTime)
				case 5: // 对方挂断
					key = fmt.Sprintf("%d_%d", request.RevUserID, userInfo.ID)
					startTime, ok3 := VideoCallMap[key]
					if !ok3 {
						fmt.Println("没有获取到起始时间")
						continue
					}
					subTime := time.Now().Sub(startTime)
					fmt.Printf("对方正常挂断， 视频通话时长为 %s\n", subTime)
				}

				switch data.Type {
				case "offer": // offer
					sendRevUserMsg(request.RevUserID, ctype.Msg{
						VideoCallMsg: &ctype.VideoCallMsg{
							Type: "offer",
							Data: data.Data,
						},
					})
					VideoCallMap[key] = time.Now()
				case "answer": // 应答
					conn.WriteJSON(ChatResponse{
						Msg: ctype.Msg{
							VideoCallMsg: &ctype.VideoCallMsg{
								Type: "answer",
								Data: data.Data,
							},
						},
					})
				case "offer_ice":
					sendRevUserMsg(request.RevUserID, ctype.Msg{
						VideoCallMsg: &ctype.VideoCallMsg{
							Type: "offer_ice",
							Data: data.Data,
						},
					})
				case "answer_ice":
					conn.WriteJSON(ChatResponse{
						Msg: ctype.Msg{
							VideoCallMsg: &ctype.VideoCallMsg{
								Type: "answer_ice",
								Data: data.Data,
							},
						},
					})
				}
				// 自己这方可以挂断

				// 对方也可以挂断

				// 如果对方开了多个浏览器，只用找其中的一个，找第一个
				continue
			}

			// 先入库
			msgID := InsertMsgByChat(svcCtx.DB, request.RevUserID, req.UserID, request.Msg)
			// 看看目标用户在不在线  给发送双方都要发消息
			SendMsgByUser(svcCtx, request.RevUserID, req.UserID, request.Msg, msgID)
		}
	}
}

type ChatRequest struct {
	RevUserID uint      `json:"revUserID"` // 给谁发
	Msg       ctype.Msg `json:"msg"`
}

// 给一组的ws对象发消息
func sendWsMapMsg(wsMap map[string]*websocket.Conn, byteData []byte) {
	for _, conn := range wsMap {
		conn.WriteMessage(websocket.TextMessage, byteData)
	}
}

// 给目标客户端发消息
func sendRevUserMsg(revUserID uint, msg ctype.Msg) {
	userRes, ok := UserOnlineWsMap[revUserID]
	if !ok {
		return
	}
	for _, conn := range userRes.WsClientMap {
		conn.WriteJSON(ChatResponse{
			SendUser: ctype.UserInfo{},
			RevUser: ctype.UserInfo{
				ID:       userRes.UserInfo.ID,
				NickName: userRes.UserInfo.Nickname,
				Avatar:   userRes.UserInfo.Avatar,
			},
			Msg:       msg,
			CreatedAt: time.Now(),
		})
		break
	}

}

type ChatResponse struct {
	ID        uint           `json:"id"`
	IsMe      bool           `json:"isMe"`
	RevUser   ctype.UserInfo `json:"revUser"`
	SendUser  ctype.UserInfo `json:"sendUser"`
	Msg       ctype.Msg      `json:"msg"`
	CreatedAt time.Time      `json:"created_at"`
}

// SendTipErrMsg 发送错误提示的消息
func SendTipErrMsg(conn *websocket.Conn, msg string) {
	resp := ChatResponse{
		Msg: ctype.Msg{
			Type: ctype.TipMsgType,
			TipMsg: &ctype.TipMsg{
				Status:  "error",
				Content: msg,
			},
		},
		CreatedAt: time.Now(),
	}
	byteData, _ := json.Marshal(resp)
	conn.WriteMessage(websocket.TextMessage, byteData)

}

// InsertMsgByChat 消息入库
func InsertMsgByChat(db *gorm.DB, revUserId uint, sendUserID uint, msg ctype.Msg) (msgID uint) {
	switch msg.Type {
	case ctype.WithdrawMsgType:
		fmt.Println("撤回消息自己是不入库的")
		return
	}
	chatModel := chat_models.ChatModel{
		SendUserID: sendUserID,
		RevUserID:  revUserId,
		MsgType:    msg.Type,
		Msg:        msg,
	}
	chatModel.MsgPreview = chatModel.MsgPreviewMethod()
	err := db.Create(&chatModel).Error
	if err != nil {
		logx.Error(err)
		sendUser, ok := UserOnlineWsMap[sendUserID]
		if !ok {
			return
		}
		SendTipErrMsg(sendUser.CurrentConn, "消息保存失败")
	}
	return chatModel.ID
}

// SendMsgByUser 发消息 给谁发 谁发的
func SendMsgByUser(svcCtx *svc.ServiceContext, revUserId uint, sendUserID uint, msg ctype.Msg, msgID uint) {

	revUser, ok1 := UserOnlineWsMap[revUserId]
	sendUser, ok2 := UserOnlineWsMap[sendUserID]
	resp := ChatResponse{
		ID:        msgID,
		Msg:       msg,
		CreatedAt: time.Now(),
	}

	if ok1 && ok2 && sendUserID == revUserId {
		// 自己给自己发
		resp.RevUser = ctype.UserInfo{
			ID:       revUserId,
			NickName: revUser.UserInfo.Nickname,
			Avatar:   revUser.UserInfo.Avatar,
		}
		resp.SendUser = ctype.UserInfo{
			ID:       sendUserID,
			NickName: sendUser.UserInfo.Nickname,
			Avatar:   sendUser.UserInfo.Avatar,
		}
		byteData, _ := json.Marshal(resp)
		//revUser.Conn.WriteMessage(websocket.TextMessage, byteData)
		sendWsMapMsg(revUser.WsClientMap, byteData)
		return
	}

	// 在线的情况下，我是可以拿到对方的用户信息的
	// 对方不在线的情况下，我只能通过调用户的rpc方法，去获取用户基本信息

	// 不管怎么样，都要给发送者回传消息的
	// 如果接受者不在线，那么我就要去拿接受者的用户信息

	if !ok1 {
		userBaseInfo, err := redis_service.GetUserBaseInfo(svcCtx.Redis, svcCtx.UserRpc, revUserId)
		if err != nil {
			logx.Error(err)
			return
		}
		resp.RevUser = ctype.UserInfo{
			ID:       revUserId,
			NickName: userBaseInfo.NickName,
			Avatar:   userBaseInfo.Avatar,
		}
	} else {
		resp.RevUser = ctype.UserInfo{
			ID:       revUserId,
			NickName: revUser.UserInfo.Nickname,
			Avatar:   revUser.UserInfo.Avatar,
		}
	}

	// 发送者在线
	resp.SendUser = ctype.UserInfo{
		ID:       sendUserID,
		NickName: sendUser.UserInfo.Nickname,
		Avatar:   sendUser.UserInfo.Avatar,
	}
	resp.IsMe = true
	byteData, _ := json.Marshal(resp)

	//sendUser.Conn.WriteMessage(websocket.TextMessage, byteData)
	sendWsMapMsg(sendUser.WsClientMap, byteData)

	if ok1 {
		// 接收者在线
		resp.IsMe = false
		byteData, _ = json.Marshal(resp)
		//revUser.Conn.WriteMessage(websocket.TextMessage, byteData)
		sendWsMapMsg(revUser.WsClientMap, byteData)
	}
}
