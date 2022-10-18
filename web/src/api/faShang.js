import service from '@/utils/request'

// @Tags FaShang
// @Summary 创建FaShang
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FaShang true "创建FaShang"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /faShang/createFaShang [post]
export const createFaShang = (data) => {
  return service({
    url: '/faShang/createFaShang',
    method: 'post',
    data
  })
}

// @Tags FaShang
// @Summary 删除FaShang
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FaShang true "删除FaShang"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /faShang/deleteFaShang [delete]
export const deleteFaShang = (data) => {
  return service({
    url: '/faShang/deleteFaShang',
    method: 'delete',
    data
  })
}

// @Tags FaShang
// @Summary 删除FaShang
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FaShang"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /faShang/deleteFaShang [delete]
export const deleteFaShangByIds = (data) => {
  return service({
    url: '/faShang/deleteFaShangByIds',
    method: 'delete',
    data
  })
}

// @Tags FaShang
// @Summary 更新FaShang
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FaShang true "更新FaShang"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /faShang/updateFaShang [put]
export const updateFaShang = (data) => {
  return service({
    url: '/faShang/updateFaShang',
    method: 'put',
    data
  })
}

// @Tags FaShang
// @Summary 用id查询FaShang
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FaShang true "用id查询FaShang"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /faShang/findFaShang [get]
export const findFaShang = (params) => {
  return service({
    url: '/faShang/findFaShang',
    method: 'get',
    params
  })
}

// @Tags FaShang
// @Summary 分页获取FaShang列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取FaShang列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /faShang/getFaShangList [get]
export const getFaShangList = (params) => {
  return service({
    url: '/faShang/getFaShangList',
    method: 'get',
    params
  })
}
