// 自动生成模板VideoDetails
package video

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// VideoDetails 结构体
type VideoDetails struct {
	global.GVA_MODEL
	VodId            *int   `json:"vodId" form:"vodId" gorm:"column:vod_id;comment:资源id;size:40;"`
	TypeName         string `json:"typeName" form:"typeName" gorm:"column:type_name;comment:类型;size:255;"`
	TypeId           int    `json:"typeId" form:"typeId" gorm:"column:type_id;comment:类型id;size:40;"`
	TypeId1          *int   `json:"typeId1" form:"typeId1" gorm:"column:type_id_1;comment:类型id1;size:40;"`
	GroupId          *int   `json:"groupId" form:"groupId" gorm:"column:group_id;comment:group_id;size:40;"`
	VodName          string `json:"vodName" form:"vodName" gorm:"column:vod_name;comment:标题;size:255;"`
	VodSub           string `json:"vodSub" form:"vodSub" gorm:"column:vod_sub;comment:简介;type:text"`
	VodEn            string `json:"vodEn" form:"vodEn" gorm:"column:vod_en;comment:拼音;size:255;"`
	VodStatus        *bool  `json:"vodStatus" form:"vodStatus" gorm:"column:vod_status;comment:状态;"`
	VodLetter        string `json:"vodLetter" form:"vodLetter" gorm:"column:vod_letter;comment:字母;size:50;"`
	VodTag           string `json:"vodTag" form:"vodTag" gorm:"column:vod_tag;comment:标签;type:text;"`
	VodClass         string `json:"vodClass" form:"vodClass" gorm:"column:vod_class;comment:类型;size:100;"`
	VodPic           string `json:"vodPic" form:"vodPic" gorm:"column:vod_pic;comment:图片;size:255;"`
	VodPicScreenshot string `json:"vodPicScreenshot" form:"vodPicScreenshot" gorm:"column:vod_pic_screenshot;comment:图片;size:255;"`
	VodActor         string `json:"vodActor" form:"vodActor" gorm:"column:vod_actor;comment:演员;type:text"`
	VodDirector      string `json:"vodDirector" form:"vodDirector" gorm:"column:vod_director;comment:导演;size:255;"`
	VodWriter        string `json:"vodWriter" form:"vodWriter" gorm:"column:vod_writer;comment:编剧;size:255;"`
	VodBehind        string `json:"vodBehind" form:"vodBehind" gorm:"column:vod_behind;comment:后台;size:255;"`
	VodBlurb         string `json:"vodBlurb" form:"vodBlurb" gorm:"column:vod_blurb;comment:简介;type:text"`
	VodRemarks       string `json:"vodRemarks" form:"vodRemarks" gorm:"column:vod_remarks;comment:书签;size:255;"`
	VodPubdate       string `json:"vodPubdate" form:"vodPubdate" gorm:"column:vod_pubdate;comment:发行日期;size:255;"`
	VodTotal         *int   `json:"vodTotal" form:"vodTotal" gorm:"column:vod_total;comment:总计;size:40;"`
	VodSerial        string `json:"vodSerial" form:"vodSerial" gorm:"column:vod_serial;comment:序列;size:255;"`
	VodArea          string `json:"vodArea" form:"vodArea" gorm:"column:vod_area;comment:区域;size:255;"`
	VodLang          string `json:"vodLang" form:"vodLang" gorm:"column:vod_lang;comment:语言;size:255;"`
	VodYear          string `json:"vodYear" form:"vodYear" gorm:"column:vod_year;comment:年代;size:255;"`
	VodHits          *int   `json:"vodHits" form:"vodHits" gorm:"column:vod_hits;comment:;size:40;"`
	VodHitsDay       *int   `json:"vodHitsDay" form:"vodHitsDay" gorm:"column:vod_hits_day;comment:;size:40;"`
	VodHitsWeek      *int   `json:"vodHitsWeek" form:"vodHitsWeek" gorm:"column:vod_hits_week;comment:;size:40;"`
	VodHitsMonth     *int   `json:"vodHitsMonth" form:"vodHitsMonth" gorm:"column:vod_hits_month;comment:;size:40;"`
	VodDuration      string `json:"vodDuration" form:"vodDuration" gorm:"column:vod_duration;comment:时间;size:40;"`
	VodUp            *int   `json:"vodUp" form:"vodUp" gorm:"column:vod_up;comment:;size:40;"`
	VodDown          *int   `json:"vodDown" form:"vodDown" gorm:"column:vod_down;comment:;size:40;"`
	VodScore         string `json:"vodScore" form:"vodScore" gorm:"column:vod_score;comment:;size:255;"`
	VodScoreAll      *int   `json:"vodScoreAll" form:"vodScoreAll" gorm:"column:vod_score_all;comment:;size:40;"`
	VodScoreNum      *int   `json:"vodScoreNum" form:"vodScoreNum" gorm:"column:vod_score_num;comment:;size:40;"`
	VodTime          string `json:"vodTime" form:"vodTime" gorm:"column:vod_time;comment:;size:255;"`
	VodTimeAdd       *int   `json:"vodTimeAdd" form:"vodTimeAdd" gorm:"column:vod_time_add;comment:;size:40;"`
	VodDoubanId      *int   `json:"vodDoubanId" form:"vodDoubanId" gorm:"column:vod_douban_id;comment:;size:40;"`
	VodDoubanScore   string `json:"vodDoubanScore" form:"vodDoubanScore" gorm:"column:vod_douban_score;comment:;size:255;"`
	VodContent       string `json:"vodContent" form:"vodContent" gorm:"column:vod_content;comment:内容;type:text;"`
	VodPlayUrl       string `json:"vodPlayUrl" form:"vodPlayUrl" gorm:"column:vod_play_url;comment:;type:longtext;"`
}

// TableName VideoDetails 表名
func (VideoDetails) TableName() string {
	return "video_details"
}
