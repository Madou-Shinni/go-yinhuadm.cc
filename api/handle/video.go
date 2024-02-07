package handle

import (
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/internal/service"
	"github.com/Madou-Shinni/gin-quickstart/pkg/constant"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/response"
	"github.com/gin-gonic/gin"
)

type VideoHandle struct {
	s *service.VideoService
}

func NewVideoHandle() *VideoHandle {
	return &VideoHandle{s: service.NewVideoService()}
}

// Add 创建Video
// @Tags     Video
// @Summary  创建Video
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.Video true "创建Video"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /video [post]
func (cl *VideoHandle) Add(c *gin.Context) {
	var video domain.Video
	if err := c.ShouldBindJSON(&video); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Add(video); err != nil {
		response.Error(c, constant.CODE_ADD_FAILED, constant.CODE_ADD_FAILED.Msg())
		return
	}

	response.Success(c)
}

// Delete 删除Video
// @Tags     Video
// @Summary  删除Video
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.Video true "删除Video"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /video [delete]
func (cl *VideoHandle) Delete(c *gin.Context) {
	var video domain.Video
	if err := c.ShouldBindJSON(&video); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Delete(video); err != nil {
		response.Error(c, constant.CODE_DELETE_FAILED, constant.CODE_DELETE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// DeleteByIds 批量删除Video
// @Tags     Video
// @Summary  批量删除Video
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.Ids true "批量删除Video"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /video/delete-batch [delete]
func (cl *VideoHandle) DeleteByIds(c *gin.Context) {
	var ids request.Ids
	if err := c.ShouldBindJSON(&ids); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.DeleteByIds(ids); err != nil {
		response.Error(c, constant.CODE_DELETE_FAILED, constant.CODE_DELETE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// Update 修改Video
// @Tags     Video
// @Summary  修改Video
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.Video true "修改Video"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /video [put]
func (cl *VideoHandle) Update(c *gin.Context) {
	var video map[string]interface{}
	if err := c.ShouldBindJSON(&video); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Update(video); err != nil {
		response.Error(c, constant.CODE_UPDATE_FAILED, constant.CODE_UPDATE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// Find 查询Video
// @Tags     Video
// @Summary  查询Video
// @accept   application/json
// @Produce  application/json
// @Param    data query     domain.Video true "查询Video"
// @Success  200  {string} string            "{"code":200,"msg":"查询成功","data":{}"}"
// @Router   /video [get]
func (cl *VideoHandle) Find(c *gin.Context) {
	var video domain.Video
	if err := c.ShouldBindQuery(&video); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	res, err := cl.s.Find(video)

	if err != nil {
		response.Error(c, constant.CODE_FIND_FAILED, constant.CODE_FIND_FAILED.Msg())
		return
	}

	response.Success(c, res)
}

// List 查询Video列表
// @Tags     Video
// @Summary  查询Video列表
// @accept   application/json
// @Produce  application/json
// @Param    data query     domain.PageVideoSearch true "查询Video列表"
// @Success  200  {string} string            "{"code":200,"msg":"查询成功","data":{}"}"
// @Router   /video/list [get]
func (cl *VideoHandle) List(c *gin.Context) {
	var video domain.PageVideoSearch
	if err := c.ShouldBindQuery(&video); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	res, err := cl.s.List(video)

	if err != nil {
		response.Error(c, constant.CODE_FIND_FAILED, constant.CODE_FIND_FAILED.Msg())
		return
	}

	response.Success(c, res)
}

// Home 查询Home
// @Tags     Home
// @Summary  查询Video列表
// @accept   application/json
// @Produce  application/json
// @Success  200  {string} string            "{"code":200,"msg":"查询成功","data":{}"}"
// @Router   /home [get]
func (cl *VideoHandle) Home(c *gin.Context) {
	res, err := cl.s.Home()
	if err != nil {
		response.Error(c, constant.CODE_FIND_FAILED, constant.CODE_FIND_FAILED.Msg())
		return
	}

	response.Success(c, res)
}
