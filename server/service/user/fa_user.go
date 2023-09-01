package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
)

type FaUserService struct {
}

// CreateFaUser 创建FaUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (faUserService *FaUserService) CreateFaUser(faUser user.FaUser) (err error) {
	err = global.GVA_DB.Create(&faUser).Error
	return err
}

// DeleteFaUser 删除FaUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (faUserService *FaUserService) DeleteFaUser(faUser user.FaUser) (err error) {
	err = global.GVA_DB.Delete(&faUser).Error
	return err
}

// DeleteFaUserByIds 批量删除FaUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (faUserService *FaUserService) DeleteFaUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]user.FaUser{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateFaUser 更新FaUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (faUserService *FaUserService) UpdateFaUser(faUser user.FaUser) (err error) {
	err = global.GVA_DB.Save(&faUser).Error
	return err
}

// GetFaUser 根据id获取FaUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (faUserService *FaUserService) GetFaUser(id uint) (faUser user.FaUser, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&faUser).Error
	return
}

// GetFaUserInfoList 分页获取FaUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (faUserService *FaUserService) GetFaUserInfoList(info userReq.FaUserSearch) (list []user.FaUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&user.FaUser{}).Preload("FaUserLevel")
	var faUsers []user.FaUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ID > 0 {
		db = db.Where("id = ?", info.ID)
	}
	if info.Username != "" {
		db = db.Where("username = ?", info.Username)
	}
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&faUsers).Error
	return faUsers, total, err
}

// Charge 后台用户充值
// Author [piexlmax](https://github.com/piexlmax)
func (faUserService *FaUserService) Charge(info user.UserCharge) (err error) {
	var FaCaiwuService FaCaiwuService
	if info.Typec == "1" {
		// 充值
		if err = FaCaiwuService.Caiwu(info.Userid, info.Price, info.Ptype, info.Typec, info.Msg); err != nil {
			return
		}
	} else {
		// 扣款
		if err = FaCaiwuService.Caiwu(info.Userid, -info.Price, info.Ptype, info.Typec, info.Msg); err != nil {
			return
		}
	}
	return
}

type Memmbers []user.FaUser

func (faUserService *FaUserService) getTeamTreeMap(ID int) (treeMap map[int][]user.FaUser, err error) {
	var gg []user.FaUser
	treeMap = make(map[int][]user.FaUser)
	err = global.GVA_DB.Model(user.FaUser{}).Where("find_in_set(?,tpath)", ID).Or("id=?", ID).Find(&gg).Error
	for _, v := range gg {
		treeMap[*v.Tjid] = append(treeMap[*v.Tjid], v)
	}
	return treeMap, err
}
func (faUserService *FaUserService) getBaseChildrenList(team *user.FaUser, treeMap map[int][]user.FaUser) (err error) {
	team.Children = treeMap[int(team.ID)]
	for i := 0; i < len(team.Children); i++ {
		err = faUserService.getBaseChildrenList(&team.Children[i], treeMap)
	}
	return err
}

// Team 后台用户团队
// Author [piexlmax](https://github.com/piexlmax)
func (faUserService *FaUserService) Team(ID int) (teams []user.FaUser, err error) {

	treeMap, err := faUserService.getTeamTreeMap(ID)
	teams = treeMap[ID]
	for i := 0; i < len(teams); i++ {
		err = faUserService.getBaseChildrenList(&teams[i], treeMap)
	}
	return teams, err

	// var gg Memmbers
	// err = global.GVA_DB.Model(user.FaUser{}).Where("find_in_set(?,tpath)", ID).Or("id=?", ID).Find(&gg).Error

	// // 生成完全树
	// resp = Tree.GenerateTree(gg.ConvertToINodeArray(), nil)
	// bytes, _ := json.MarshalIndent(gg, "", "\t")
	// fmt.Println(string(bytes))
	// return resp, err
}
