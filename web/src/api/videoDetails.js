import service from '@/utils/request'

// @Tags VideoDetails
// @Summary 创建VideoDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VideoDetails true "创建VideoDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoDetails/createVideoDetails [post]
export const createVideoDetails = (data) => {
  return service({
    url: '/videoDetails/createVideoDetails',
    method: 'post',
    data
  })
}

// @Tags VideoDetails
// @Summary 删除VideoDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VideoDetails true "删除VideoDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoDetails/deleteVideoDetails [delete]
export const deleteVideoDetails = (data) => {
  return service({
    url: '/videoDetails/deleteVideoDetails',
    method: 'delete',
    data
  })
}

// @Tags VideoDetails
// @Summary 删除VideoDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除VideoDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoDetails/deleteVideoDetails [delete]
export const deleteVideoDetailsByIds = (data) => {
  return service({
    url: '/videoDetails/deleteVideoDetailsByIds',
    method: 'delete',
    data
  })
}

// @Tags VideoDetails
// @Summary 更新VideoDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VideoDetails true "更新VideoDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /videoDetails/updateVideoDetails [put]
export const updateVideoDetails = (data) => {
  return service({
    url: '/videoDetails/updateVideoDetails',
    method: 'put',
    data
  })
}

// @Tags VideoDetails
// @Summary 用id查询VideoDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.VideoDetails true "用id查询VideoDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /videoDetails/findVideoDetails [get]
export const findVideoDetails = (params) => {
  return service({
    url: '/videoDetails/findVideoDetails',
    method: 'get',
    params
  })
}

// @Tags VideoDetails
// @Summary 分页获取VideoDetails列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取VideoDetails列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoDetails/getVideoDetailsList [get]
export const getVideoDetailsList = (params) => {
  return service({
    url: '/videoDetails/getVideoDetailsList',
    method: 'get',
    params
  })
}
