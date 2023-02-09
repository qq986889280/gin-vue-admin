import service from '@/utils/request'

// @Tags FaUser
// @Summary 创建FaUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FaUser true "创建FaUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /faUser/createFaUser [post]
export const createFaUser = (data) => {
  return service({
    url: '/faUser/createFaUser',
    method: 'post',
    data
  })
}

// @Tags FaUser
// @Summary 删除FaUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FaUser true "删除FaUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /faUser/deleteFaUser [delete]
export const deleteFaUser = (data) => {
  return service({
    url: '/faUser/deleteFaUser',
    method: 'delete',
    data
  })
}

// @Tags FaUser
// @Summary 删除FaUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FaUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /faUser/deleteFaUser [delete]
export const deleteFaUserByIds = (data) => {
  return service({
    url: '/faUser/deleteFaUserByIds',
    method: 'delete',
    data
  })
}

// @Tags FaUser
// @Summary 更新FaUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FaUser true "更新FaUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /faUser/updateFaUser [put]
export const updateFaUser = (data) => {
  return service({
    url: '/faUser/updateFaUser',
    method: 'put',
    data
  })
}

// @Tags FaUser
// @Summary 用id查询FaUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FaUser true "用id查询FaUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /faUser/findFaUser [get]
export const findFaUser = (params) => {
  return service({
    url: '/faUser/findFaUser',
    method: 'get',
    params
  })
}

// @Tags FaUser
// @Summary 分页获取FaUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取FaUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /faUser/getFaUserList [get]
export const getFaUserList = (params) => {
  return service({
    url: '/faUser/getFaUserList',
    method: 'get',
    params
  })
}
