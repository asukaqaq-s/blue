package logic

import (
	"context"
	"errors"
	"fim_server/common/models/ctype"
	"fim_server/fim_user/user_api/internal/svc"
	"fim_server/fim_user/user_api/internal/types"
	"fim_server/fim_user/user_models"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendLogic {
	return &AddFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFriendLogic) AddFriend(req *types.AddFriendRequest) (resp *types.AddFriendResponse, err error) {

	// 限制用户加好友
	var conf user_models.UserConfModel
	err = l.svcCtx.DB.Take(&conf, "user_id = ?", req.UserID).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	if conf.CurtailAddUser {
		return nil, errors.New("当前用户限制加好友")
	}

	// 如果你们是好友，就不用加了
	var friend user_models.FriendModel
	if friend.IsFriend(l.svcCtx.DB, req.UserID, req.FriendID) {
		return nil, errors.New("你们已经是好友了")
	}

	var userConf user_models.UserConfModel
	err = l.svcCtx.DB.Take(&userConf, "user_id = ?", req.FriendID).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	resp = new(types.AddFriendResponse)
	var verifyModel = user_models.FriendVerifyModel{
		SendUserID:         req.UserID,
		RevUserID:          req.FriendID,
		AdditionalMessages: req.Verify,
	}

	switch userConf.Verification {
	case 0: // 不允许任何人添加
		return nil, errors.New("该用户不允许任何人添加")
	case 1: // 允许任何人添加
		// 直接成为好友
		// 先往验证表里面加一条记录，然后通过
		verifyModel.RevStatus = 1
		var userFriend = user_models.FriendModel{
			SendUserID: req.UserID,
			RevUserID:  req.FriendID,
		}
		l.svcCtx.DB.Create(&userFriend)
	case 2: // 需要验证问题
	case 3: // 需要回答问题
		if req.VerificationQuestion != nil {
			verifyModel.VerificationQuestion = &ctype.VerificationQuestion{
				Problem1: req.VerificationQuestion.Problem1,
				Problem2: req.VerificationQuestion.Problem2,
				Problem3: req.VerificationQuestion.Problem3,
				Answer1:  req.VerificationQuestion.Answer1,
				Answer2:  req.VerificationQuestion.Answer2,
				Answer3:  req.VerificationQuestion.Answer3,
			}
		}
	case 4: // 需要正确回答问题
		// 判断问题是否回答正确
		if req.VerificationQuestion != nil && userConf.VerificationQuestion != nil {
			// 考虑到一个问题，两个问题，三个问题的情况
			var count int
			if userConf.VerificationQuestion.Answer1 != nil && req.VerificationQuestion.Answer1 != nil {
				if *userConf.VerificationQuestion.Answer1 == *req.VerificationQuestion.Answer1 {
					count += 1
				}
			}
			if userConf.VerificationQuestion.Answer2 != nil && req.VerificationQuestion.Answer2 != nil {
				if *userConf.VerificationQuestion.Answer2 == *req.VerificationQuestion.Answer2 {
					count += 1
				}
			}
			if userConf.VerificationQuestion.Answer3 != nil && req.VerificationQuestion.Answer3 != nil {
				if *userConf.VerificationQuestion.Answer3 == *req.VerificationQuestion.Answer3 {
					count += 1
				}
			}
			if count != userConf.ProblemCount() {
				return nil, errors.New("答案错误")
			}
			// 直接加好友
			verifyModel.RevStatus = 1
			verifyModel.VerificationQuestion = userConf.VerificationQuestion
			// 加好友
			var userFriend = user_models.FriendModel{
				SendUserID: req.UserID,
				RevUserID:  req.FriendID,
			}
			l.svcCtx.DB.Create(&userFriend)
		}
	default:
		return nil, errors.New("不支持的验证参数")
	}
	err = l.svcCtx.DB.Create(&verifyModel).Error
	if err != nil {
		logx.Error(err)
		return nil, errors.New("添加好友失败")
	}

	return
}
