package Admin

import (
	"context"
	"fim_server/fim_group/group_models"
	"gorm.io/gorm"

	"fim_server/fim_group/group_api/internal/svc"
	"fim_server/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupListRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupListRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupListRemoveLogic {
	return &GroupListRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GroupListRemove 删除群
func (l *GroupListRemoveLogic) GroupListRemove(req *types.GroupListRemoveRequest) (resp *types.GroupListRemoveResponse, err error) {
	var groupList []group_models.GroupModel
	l.svcCtx.DB.Preload("MemberList").Preload("GroupMsgList").Find(&groupList, "id in ?", req.IdList)
	for _, model := range groupList {
		logx.Infof("删除群聊id %d, 群名称 %s", model.ID, model.Title)
		err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
			if len(model.GroupMsgList) > 0 {
				err = tx.Delete(&model.GroupMsgList).Error
				if err != nil {
					return err
				}
			}
			logx.Infof("关联删除群消息总数 %d", len(model.GroupMsgList))
			if len(model.MemberList) > 0 {
				err = tx.Delete(&model.MemberList).Error
				if err != nil {
					return err
				}
			}
			logx.Infof("关联删除群用户总数 %d", len(model.MemberList))
			var topModelList []group_models.GroupUserTopModel
			err = tx.Find(&topModelList, "group_id = ?", model.ID).Delete(&topModelList).Error
			if err != nil {
				return err
			}
			logx.Infof("关联删除用户置顶总数 %d", len(topModelList))
			var verifyList []group_models.GroupVerifyModel
			err = tx.Find(&verifyList, "group_id = ?", model.ID).Delete(&verifyList).Error
			if err != nil {
				return err
			}
			logx.Infof("关联删除群验证消息总数 %d", len(verifyList))
			var userDeleteMsgList []group_models.GroupUserMsgDeleteModel
			err = tx.Find(&userDeleteMsgList, "group_id = ?", model.ID).Delete(&userDeleteMsgList).Error
			if err != nil {
				return err
			}
			logx.Infof("关联删除用户删除聊天记录总数 %d", len(userDeleteMsgList))
			err = tx.Delete(&model).Error
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			logx.Error(err)
			continue
		}
		logx.Infof("删除成功")
	}

	return
}
