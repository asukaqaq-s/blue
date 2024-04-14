package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fim_server/common/models/ctype"
	"fim_server/fim_chat/chat_rpc/chat"
	"fim_server/fim_user/user_api/internal/svc"
	"fim_server/fim_user/user_api/internal/types"
	"fim_server/fim_user/user_models"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewValidStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidStatusLogic {
	return &ValidStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ValidStatusLogic) ValidStatus(req *types.FriendValidStatusRequest) (resp *types.FriendValidStatusResponse, err error) {

	// 我如果是发起方，那么我只能删除 还得是对方操作了
	// 我如果是接收方，那么我只能 1234
	// 如果要4，那么不能是0

	var friendVerify user_models.FriendVerifyModel
	// 我要操作状态，我自己得是接收方
	err = l.svcCtx.DB.Take(&friendVerify, "id = ?", req.VerifyID).Error
	if err != nil {
		return nil, errors.New("验证记录不存在")
	}

	// 去找我是发起方还是接收方
	if friendVerify.SendUserID == req.UserID {
		// 我就是发起方  发起方只能删除
		switch req.Status {
		case 4:
			// 如果是4，那就是接收方要删除这个验证记录，必须等接收方，处理之后才能删
			if friendVerify.RevStatus == 0 {
				return nil, errors.New("接收方未处理，不可删除")
			}
		default:
			return nil, errors.New("发送方错误的状态")

		}
	} else {
		// 我就是接收方
		switch req.Status {
		case 1, 2, 3:
			// 如果是这些状态，那么revStatus就得是 0
			if friendVerify.RevStatus != 0 {
				return nil, errors.New("不可更改状态")
			}
		case 4:
			// 如果是4，那就是接收方要删除这个验证记录
			if friendVerify.RevStatus == 0 {
				return nil, errors.New("不可删除未处理状态")
			}
		default:
			return nil, errors.New("接收方错误的状态")
		}
	}

	switch req.Status {
	case 1: // 同意
		friendVerify.RevStatus = 1
		// 往好友表里面加
		l.svcCtx.DB.Create(&user_models.FriendModel{
			SendUserID: friendVerify.SendUserID,
			RevUserID:  friendVerify.RevUserID,
		})

		msg := ctype.Msg{
			Type: ctype.TextMsgType,
			TextMsg: &ctype.TextMsg{
				Content: "我们已经是好友了，开始聊天吧！",
			},
		}
		byteData, _ := json.Marshal(msg)

		// 给对方发个消息
		_, err = l.svcCtx.ChatRpc.UserChat(l.ctx, &chat.UserChatRequest{
			SendUserId: uint32(friendVerify.SendUserID),
			RevUserId:  uint32(friendVerify.RevUserID),
			Msg:        byteData,
			SystemMsg:  nil,
		})
		if err != nil {
			logx.Error(err)
		}
	case 2: // 拒绝
		friendVerify.RevStatus = 2
	case 3: // 忽略
		friendVerify.RevStatus = 3
	case 4: // 删除
		if friendVerify.SendUserID == req.UserID {
			// 我是发送方
			friendVerify.SendStatus = 4
		} else {
			friendVerify.RevStatus = 4
		}
		// 一条验证记录，是两个人看的
		//l.svcCtx.DB.Delete(&friendVerify)
		//return nil, nil
	}
	l.svcCtx.DB.Save(&friendVerify)
	return
}
