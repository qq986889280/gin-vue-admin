package video

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/video"
	videoReq "github.com/flipped-aurora/gin-vue-admin/server/model/video/request"
)

type VideoDetailsService struct {
}

// CreateVideoDetails 创建VideoDetails记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoDetailsService *VideoDetailsService) CreateVideoDetails(videoDetails video.VideoDetails) (err error) {
	err = global.GVA_DB.Create(&videoDetails).Error
	return err
}

// DeleteVideoDetails 删除VideoDetails记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoDetailsService *VideoDetailsService) DeleteVideoDetails(videoDetails video.VideoDetails) (err error) {
	err = global.GVA_DB.Delete(&videoDetails).Error
	return err
}

// DeleteVideoDetailsByIds 批量删除VideoDetails记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoDetailsService *VideoDetailsService) DeleteVideoDetailsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]video.VideoDetails{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateVideoDetails 更新VideoDetails记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoDetailsService *VideoDetailsService) UpdateVideoDetails(videoDetails video.VideoDetails) (err error) {
	err = global.GVA_DB.Save(&videoDetails).Error
	return err
}

// GetVideoDetails 根据id获取VideoDetails记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoDetailsService *VideoDetailsService) GetVideoDetails(id uint) (videoDetails video.VideoDetails, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&videoDetails).Error
	return
}

// GetVideoDetailsInfoList 分页获取VideoDetails记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoDetailsService *VideoDetailsService) GetVideoDetailsInfoList(info videoReq.VideoDetailsSearch) (list []video.VideoDetails, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&video.VideoDetails{})
	var videoDetailss []video.VideoDetails
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.TypeId > 0 {
		db = db.Where("type_id = ?", info.TypeId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&videoDetailss).Error
	return videoDetailss, total, err
}
