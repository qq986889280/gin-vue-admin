// 自动生成模板FaUser
package user

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
)

// FaUser 结构体
type FaUser struct {
	global.GVA_MODEL
	GroupId      *int       `json:"groupId" form:"groupId" gorm:"column:group_id;comment:组别ID;size:10;"`
	Username     string     `json:"username" form:"username" gorm:"column:username;comment:用户名;size:200;"`
	Nickname     string     `json:"nickname" form:"nickname" gorm:"column:nickname;comment:昵称;size:100;"`
	Password     string     `json:"password" form:"password" gorm:"column:password;comment:密码;size:100;"`
	Password2    string     `json:"password2" form:"password2" gorm:"column:password2;comment:支付密码;size:100;"`
	Password3    string     `json:"password3" form:"password3" gorm:"column:password3;comment:谷歌秘钥;size:100;"`
	Salt         string     `json:"salt" form:"salt" gorm:"column:salt;comment:密码盐;size:30;"`
	Email        string     `json:"email" form:"email" gorm:"column:email;comment:电子邮箱;size:100;"`
	Mobile       string     `json:"mobile" form:"mobile" gorm:"column:mobile;comment:手机号;size:11;"`
	Avatar       string     `json:"avatar" form:"avatar" gorm:"column:avatar;comment:头像;size:255;"`
	Level        int        `json:"level" form:"level" gorm:"column:level;comment:等级;"`
	Gender       *bool      `json:"gender" form:"gender" gorm:"column:gender;comment:性别;"`
	Birthday     *time.Time `json:"birthday" form:"birthday" gorm:"column:birthday;comment:生日;"`
	Bio          string     `json:"bio" form:"bio" gorm:"column:bio;comment:格言;size:100;"`
	Logintime    *int       `json:"logintime" form:"logintime" gorm:"column:logintime;comment:登录时间;size:10;"`
	Loginip      string     `json:"loginip" form:"loginip" gorm:"column:loginip;comment:登录IP;size:50;"`
	Loginfailure *bool      `json:"loginfailure" form:"loginfailure" gorm:"column:loginfailure;comment:失败次数;"`
	Jointime     *int       `json:"jointime" form:"jointime" gorm:"column:jointime;comment:加入时间;size:10;"`
	Viptime      *int       `json:"viptime" form:"viptime" gorm:"column:viptime;comment:会员到期时间;size:10;"`
	Wall1        *float64   `json:"wall1" form:"wall1" gorm:"column:wall1;comment:wall1;size:17;"`
	Wall2        *float64   `json:"wall2" form:"wall2" gorm:"column:wall2;comment:wall2;size:17;"`
	Wall3        *float64   `json:"wall3" form:"wall3" gorm:"column:wall3;comment:wall3;size:17;"`
	Wall4        *float64   `json:"wall4" form:"wall4" gorm:"column:wall4;comment:wall4;size:17;"`
	Wall5        *float64   `json:"wall5" form:"wall5" gorm:"column:wall5;comment:wall5;size:17;"`
	Wall6        *float64   `json:"wall6" form:"wall6" gorm:"column:wall6;comment:wall6;size:17;"`
	Wall7        *float64   `json:"wall7" form:"wall7" gorm:"column:wall7;comment:wall7;size:17;"`
	Wall8        *float64   `json:"wall8" form:"wall8" gorm:"column:wall8;comment:wall8;size:17;"`
	Wall9        *float64   `json:"wall9" form:"wall9" gorm:"column:wall9;comment:wall9;size:17;"`
	Wall10       *float64   `json:"wall10" form:"wall10" gorm:"column:wall10;comment:wall10;size:17;"`
	Status       int        `json:"status" form:"status" gorm:"column:status;comment:状态;"`
	Wall1freeze  *float64   `json:"wall1freeze" form:"wall1freeze" gorm:"column:wall1freeze;comment:wall1;size:17;"`
	Wall2freeze  *float64   `json:"wall2freeze" form:"wall2freeze" gorm:"column:wall2freeze;comment:wall2;size:17;"`
	Wall3freeze  *float64   `json:"wall3freeze" form:"wall3freeze" gorm:"column:wall3freeze;comment:wall3;size:17;"`
	Wall4freeze  *float64   `json:"wall4freeze" form:"wall4freeze" gorm:"column:wall4freeze;comment:wall4;size:17;"`
	Wall5freeze  *float64   `json:"wall5freeze" form:"wall5freeze" gorm:"column:wall5freeze;comment:wall5;size:17;"`
	Wall6freeze  *float64   `json:"wall6freeze" form:"wall6freeze" gorm:"column:wall6freeze;comment:wall6;size:17;"`
	Wall7freeze  *float64   `json:"wall7freeze" form:"wall7freeze" gorm:"column:wall7freeze;comment:wall7;size:17;"`
	Wall8freeze  *float64   `json:"wall8freeze" form:"wall8freeze" gorm:"column:wall8freeze;comment:wall8;size:17;"`
	Wall9freeze  *float64   `json:"wall9freeze" form:"wall9freeze" gorm:"column:wall9freeze;comment:wall9;size:17;"`
	Wall10freeze *float64   `json:"wall10freeze" form:"wall10freeze" gorm:"column:wall10freeze;comment:wall10;size:17;"`
	Tjuser       string     `json:"tjuser" form:"tjuser" gorm:"column:tjuser;comment:推荐人账号;size:32;"`
	Tjid         *int       `json:"tjid" form:"tjid" gorm:"column:tjid;comment:推荐人uid;size:10;"`
	Tpath        string     `json:"tpath" form:"tpath" gorm:"column:tpath;comment:推荐路径;size:255;"`
	Tgno         string     `json:"tgno" form:"tgno" gorm:"column:tgno;comment:邀请码;size:50;"`
	Ztlevel      string     `json:"ztlevel" form:"ztlevel" gorm:"column:ztlevel;comment:直推代数;size:10;"`
	Ztnum        *int       `json:"ztnum" form:"ztnum" gorm:"column:ztnum;comment:直推人数;size:10;"`
	Tdnum        *int       `json:"tdnum" form:"tdnum" gorm:"column:tdnum;comment:团队人数;size:10;"`
	IsAgent      string     `json:"isAgent" form:"isAgent" gorm:"column:is_agent;comment:是否商户;size:1;"`
	IsB          string     `json:"isB" form:"isB" gorm:"column:is_b;comment:是否通过谷歌验证;size:1;"`
	Times        *int       `json:"times" form:"times" gorm:"column:times;comment:次数;size:10;"`
	Todayzt      *int       `json:"todayzt" form:"todayzt" gorm:"column:todayzt;comment:今日直推;size:10;"`
	Tdyj         *int       `json:"tdyj" form:"tdyj" gorm:"column:tdyj;comment:团队业绩;size:10;"`
	Gryj         *int       `json:"gryj" form:"gryj" gorm:"column:gryj;comment:个人业绩;size:10;"`
	Realname     string     `json:"realname" form:"realname" gorm:"column:realname;comment:姓名;size:50;"`
	Idcard       string     `json:"idcard" form:"idcard" gorm:"column:idcard;comment:身份证号;size:255;"`
	Card1        string     `json:"card1" form:"card1" gorm:"column:card1;comment:身份证正面;size:255;"`
	Card2        string     `json:"card2" form:"card2" gorm:"column:card2;comment:身份证反面;size:255;"`
	Card3        string     `json:"card3" form:"card3" gorm:"column:card3;comment:护照;size:255;"`
	IsSm         string     `json:"isSm" form:"isSm" gorm:"column:is_sm;comment:实名状态:0=未实名,1=已实名,2=待审核;size:1;"`
	S1           *float64   `json:"S1" form:"S1" gorm:"column:S1;comment:S1账户;size:19;"`
	S2           *float64   `json:"S2" form:"S2" gorm:"column:S2;comment:S2账户;size:19;"`
	B1total      *float64   `json:"B1total" form:"B1total" gorm:"column:B1total;comment:累计1;size:19;"`
	B2total      *float64   `json:"B2total" form:"B2total" gorm:"column:B2total;comment:累计2;size:19;"`
	B3total      *float64   `json:"B3total" form:"B3total" gorm:"column:B3total;comment:累计3;size:19;"`
	B1yest       *float64   `json:"B1yest" form:"B1yest" gorm:"column:B1yest;comment:昨日1;size:19;"`
	B2yest       *float64   `json:"B2yest" form:"B2yest" gorm:"column:B2yest;comment:昨日2;size:19;"`
	B3yest       *float64   `json:"B3yest" form:"B3yest" gorm:"column:B3yest;comment:昨日3;size:19;"`
	Suan         string     `json:"suan" form:"suan" gorm:"column:suan;comment:算力1;size:255;"`
	Suan2        string     `json:"suan2" form:"suan2" gorm:"column:suan2;comment:算力2;size:255;"`
	IsSim        string     `json:"isSim" form:"isSim" gorm:"column:is_sim;comment:模拟;size:1;"`
	Win          string     `json:"win" form:"win" gorm:"column:win;comment:盈次数;size:255;"`
	Sw1          int        `json:"sw1" form:"sw1" gorm:"column:sw1;comment:VIP;"`
	Sw2          int        `json:"sw2" form:"sw2" gorm:"column:sw2;comment:提现;"`
	Sw3          int        `json:"sw3" form:"sw3" gorm:"column:sw3;comment:闪兑;"`
	Sw4          int        `json:"sw4" form:"sw4" gorm:"column:sw4;comment:OTC;"`
	Sw5          int        `json:"sw5" form:"sw5" gorm:"column:sw5;comment:互转;"`
	Sw6          int        `json:"sw6" form:"sw6" gorm:"column:sw6;comment:奖励;"`
}

// TableName FaUser 表名
func (FaUser) TableName() string {
	return "fa_user"
}

func (u *FaUser) BeforeCreate(tx *gorm.DB) (err error) {

	u.Password = utils.BcryptHash(u.Password)
	u.Password2 = utils.BcryptHash(u.Password2)
	return
}
