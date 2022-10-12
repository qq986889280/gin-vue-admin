import service from '@/utils/request'

// @Tags FaCharge
// @Summary 创建FaCharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FaCharge true "创建FaCharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /faCharge/createFaCharge [post]
export const createFaCharge = (data) => {
  return service({
    url: '/faCharge/createFaCharge',
    method: 'post',
    data
  })
}

// @Tags FaCharge
// @Summary 删除FaCharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FaCharge true "删除FaCharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /faCharge/deleteFaCharge [delete]
export const deleteFaCharge = (data) => {
  return service({
    url: '/faCharge/deleteFaCharge',
    method: 'delete',
    data
  })
}

// @Tags FaCharge
// @Summary 删除FaCharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FaCharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /faCharge/deleteFaCharge [delete]
export const deleteFaChargeByIds = (data) => {
  return service({
    url: '/faCharge/deleteFaChargeByIds',
    method: 'delete',
    data
  })
}

// @Tags FaCharge
// @Summary 更新FaCharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FaCharge true "更新FaCharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /faCharge/updateFaCharge [put]
export const updateFaCharge = (data) => {
  return service({
    url: '/faCharge/updateFaCharge',
    method: 'put',
    data
  })
}

// @Tags FaCharge
// @Summary 用id查询FaCharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FaCharge true "用id查询FaCharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /faCharge/findFaCharge [get]
export const findFaCharge = (params) => {
  return service({
    url: '/faCharge/findFaCharge',
    method: 'get',
    params
  })
}

// @Tags FaCharge
// @Summary 分页获取FaCharge列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取FaCharge列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /faCharge/getFaChargeList [get]
export const getFaChargeList = (params) => {
  return service({
    url: '/faCharge/getFaChargeList',
    method: 'get',
    params
  })
}
