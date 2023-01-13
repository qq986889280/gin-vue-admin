package model

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
type Request struct {
}

type Response struct {
}

type VideoDetail struct {
	GVA_MODEL
	VodID            int64  `json:"vod_id" gorm:"index;column:vod_id;comment:资源id;size:40;"`
	TypeName         string `json:"type_name" gorm:"column:type_name;comment:类型;size:255;"`
	TypeID           int64  `json:"type_id" gorm:"column:type_id;comment:类型id;size:40;"`
	TypeID1          int64  `json:"type_id_1" gorm:"column:type_id_1;comment:类型id1;size:40;"`
	GroupID          int64  `json:"group_id" gorm:"column:group_id;comment:group_id;size:40;"`
	VodName          string `json:"vod_name" gorm:"column:vod_name;comment:标题;size:255;"`
	VodSub           string `json:"vod_sub" gorm:"column:vod_sub;comment:简介;type:text;"`
	VodEn            string `json:"vod_en" gorm:"column:vod_en;comment:拼音;size:255;"`
	VodStatus        int64  `json:"vod_status" gorm:"column:vod_status;comment:状态;size:3;"`
	VodLetter        string `json:"vod_letter" gorm:"column:vod_letter;comment:字母;size:50;"`
	VodTag           string `json:"vod_tag" gorm:"column:vod_tag;comment:标签;type:text;"`
	VodClass         string `json:"vod_class" gorm:"column:vod_class;comment:类型;size:100;"`
	VodPic           string `json:"vod_pic" gorm:"column:vod_pic;comment:图片;size:255;"`
	VodPicScreenshot string `json:"vod_pic_screenshot" gorm:"column:vod_pic_screenshot;comment:图片;size:255;"`
	VodActor         string `json:"vod_actor" gorm:"column:vod_actor;comment:演员;type:text;"`
	VodDirector      string `json:"vod_director" gorm:"column:vod_director;comment:导演;size:255;"`
	VodWriter        string `json:"vod_writer" gorm:"column:vod_writer;comment:编剧;size:255;"`
	VodBehind        string `json:"vod_behind"  gorm:"column:vod_behind;comment:后台;size:255;"`
	VodBlurb         string `json:"vod_blurb"  gorm:"column:vod_blurb;comment:简介;type:text;"`
	VodRemarks       string `json:"vod_remarks"  gorm:"column:vod_remarks;comment:书签;size:255;"`
	VodPubdate       string `json:"vod_pubdate"  gorm:"column:vod_pubdate;comment:发行日期;size:255;"`
	VodTotal         int64  `json:"vod_total"  gorm:"column:vod_total;comment:总计;size:40;"`
	VodSerial        string `json:"vod_serial" gorm:"column:vod_serial;comment:序列;size:255;"`
	VodArea          string `json:"vod_area" gorm:"column:vod_area;comment:区域;size:255;"`
	VodLang          string `json:"vod_lang" gorm:"column:vod_lang;comment:语言;size:255;"`
	VodYear          string `json:"vod_year" gorm:"column:vod_year;comment:年代;size:255;"`
	VodHits          int64  `json:"vod_hits" gorm:"column:vod_hits;size:40;"`
	VodHitsDay       int64  `json:"vod_hits_day" gorm:"column:vod_hits_day;size:40;"`
	VodHitsWeek      int64  `json:"vod_hits_week" gorm:"column:vod_hits_week;size:40;"`
	VodHitsMonth     int64  `json:"vod_hits_month" gorm:"column:vod_hits_month;size:40;"`
	VodDuration      string `json:"vod_duration" gorm:"column:vod_duration;comment:时间;size:40;"`
	VodUp            int64  `json:"vod_up" gorm:"column:vod_up;size:40;"`
	VodDown          int64  `json:"vod_down" gorm:"column:vod_down;size:40;"`
	VodScore         string `json:"vod_score" gorm:"column:vod_score;size:255;"`
	VodScoreAll      int64  `json:"vod_score_all" gorm:"column:vod_score_all;size:255;"`
	VodScoreNum      int64  `json:"vod_score_num" gorm:"column:vod_score_num;size:255;"`
	VodTime          string `json:"vod_time" gorm:"column:vod_time;size:255;"`
	VodTimeAdd       int64  `json:"vod_time_add" gorm:"column:vod_time_add;size:40;"`
	VodDoubanID      int64  `json:"vod_douban_id" gorm:"column:vod_douban_id;size:40;"`
	VodDoubanScore   string `json:"vod_douban_score" gorm:"column:vod_douban_score;size:255;"`
	VodContent       string `json:"vod_content" gorm:"column:vod_content;type:text;"`
	VodPlayURL       string `json:"vod_play_url" gorm:"column:vod_play_url;type:longtext;"`
}

type TkzyRe struct {
	Code      int            `json:"code"`
	Msg       string         `json:"msg"`
	Page      int            `json:"page"`
	Pagecount int            `json:"pagecount"`
	Limit     string         `json:"limit"`
	Total     int            `json:"total"`
	List      *[]VideoDetail `json:"list"`
}
