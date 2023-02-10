import service from '@/utils/request'

// @Tags FaCaiwu
// @Summary 创建FaCaiwu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FaCaiwu true "创建FaCaiwu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /faCaiwu/createFaCaiwu [post]
export const createFaCaiwu = (data) => {
  return service({
    url: '/faCaiwu/createFaCaiwu',
    method: 'post',
    data
  })
}

// @Tags FaCaiwu
// @Summary 删除FaCaiwu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FaCaiwu true "删除FaCaiwu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /faCaiwu/deleteFaCaiwu [delete]
export const deleteFaCaiwu = (data) => {
  return service({
    url: '/faCaiwu/deleteFaCaiwu',
    method: 'delete',
    data
  })
}

// @Tags FaCaiwu
// @Summary 删除FaCaiwu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FaCaiwu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /faCaiwu/deleteFaCaiwu [delete]
export const deleteFaCaiwuByIds = (data) => {
  return service({
    url: '/faCaiwu/deleteFaCaiwuByIds',
    method: 'delete',
    data
  })
}

// @Tags FaCaiwu
// @Summary 更新FaCaiwu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FaCaiwu true "更新FaCaiwu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /faCaiwu/updateFaCaiwu [put]
export const updateFaCaiwu = (data) => {
  return service({
    url: '/faCaiwu/updateFaCaiwu',
    method: 'put',
    data
  })
}

// @Tags FaCaiwu
// @Summary 用id查询FaCaiwu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FaCaiwu true "用id查询FaCaiwu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /faCaiwu/findFaCaiwu [get]
export const findFaCaiwu = (params) => {
  return service({
    url: '/faCaiwu/findFaCaiwu',
    method: 'get',
    params
  })
}

// @Tags FaCaiwu
// @Summary 分页获取FaCaiwu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取FaCaiwu列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /faCaiwu/getFaCaiwuList [get]
export const getFaCaiwuList = (params) => {
  return service({
    url: '/faCaiwu/getFaCaiwuList',
    method: 'get',
    params
  })
}
