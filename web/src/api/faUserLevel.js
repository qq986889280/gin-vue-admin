import service from '@/utils/request'

// @Tags FaUserLevel
// @Summary 创建FaUserLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FaUserLevel true "创建FaUserLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /faUserLevel/createFaUserLevel [post]
export const createFaUserLevel = (data) => {
  return service({
    url: '/faUserLevel/createFaUserLevel',
    method: 'post',
    data
  })
}

// @Tags FaUserLevel
// @Summary 删除FaUserLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FaUserLevel true "删除FaUserLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /faUserLevel/deleteFaUserLevel [delete]
export const deleteFaUserLevel = (data) => {
  return service({
    url: '/faUserLevel/deleteFaUserLevel',
    method: 'delete',
    data
  })
}

// @Tags FaUserLevel
// @Summary 删除FaUserLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FaUserLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /faUserLevel/deleteFaUserLevel [delete]
export const deleteFaUserLevelByIds = (data) => {
  return service({
    url: '/faUserLevel/deleteFaUserLevelByIds',
    method: 'delete',
    data
  })
}

// @Tags FaUserLevel
// @Summary 更新FaUserLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FaUserLevel true "更新FaUserLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /faUserLevel/updateFaUserLevel [put]
export const updateFaUserLevel = (data) => {
  return service({
    url: '/faUserLevel/updateFaUserLevel',
    method: 'put',
    data
  })
}

// @Tags FaUserLevel
// @Summary 用id查询FaUserLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FaUserLevel true "用id查询FaUserLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /faUserLevel/findFaUserLevel [get]
export const findFaUserLevel = (params) => {
  return service({
    url: '/faUserLevel/findFaUserLevel',
    method: 'get',
    params
  })
}

// @Tags FaUserLevel
// @Summary 分页获取FaUserLevel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取FaUserLevel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /faUserLevel/getFaUserLevelList [get]
export const getFaUserLevelList = (params) => {
  return service({
    url: '/faUserLevel/getFaUserLevelList',
    method: 'get',
    params
  })
}
