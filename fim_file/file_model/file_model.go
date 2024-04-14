package file_model

import (
	"fim_server/common/models"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"os"
)

type FileModel struct {
	models.Model
	Uid      uuid.UUID `json:"uid"`      // 文件唯一id /api/file/{uuid}
	UserID   uint      `json:"userID"`   // 用户id
	FileName string    `json:"fileName"` // 文件名称
	Size     int64     `json:"size"`     // 文件大小
	Path     string    `json:"path"`     // 文件的实际路径
	Hash     string    `json:"hash"`     // 文件hash
}

func (file *FileModel) WebPath() string {
	return "/api/file/" + file.Uid.String()
}
func (file *FileModel) BeforeDelete(tx *gorm.DB) (err error) {
	logx.Infof("删除文件的名称 %s", file.FileName)
	if file.Path != "" {
		err1 := os.Remove(file.Path)
		if err1 != nil {
			logx.Error(err1)
		} else {
			logx.Infof("文件源地址删除 %s", file.Path)
		}
	}
	return
}
