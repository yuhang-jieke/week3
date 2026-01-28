package server

import (
	"net/http"
	"week3/wei/api-getaway/basic/config"
	"week3/wei/api-getaway/handler/request"

	__ "week3/wei/api-getaway/basic/proto"

	"github.com/gin-gonic/gin"
)

func GoodsAdd(c *gin.Context) {
	var form request.GoodsAdd
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不正确",
		})
		return
	}

	_, err := config.GoodsClient.GoodsAdd(c, &__.GoodsAddReq{
		Name:  form.Name,
		Price: form.Price,
		Stock: int64(form.Stock),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
	return
}
func GoodsShow(c *gin.Context) {
	var form request.GoodsShow
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不正确",
		})
		return
	}

	t, err := config.GoodsClient.GoodsList(c, &__.GoodsListReq{
		Id: int64(form.Id),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "搜索失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "搜索成功",
		"data": t,
	})
	return
}
func Es(c *gin.Context) {
	var form request.Es
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不正确",
		})
		return
	}

	t, err := config.GoodsClient.Es(c, &__.EsReq{
		Keyword: form.KeyWord,
		Page:    int64(form.Page),
		Size:    int64(form.Size),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "搜索失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "搜索成功",
		"data": t,
	})
	return
}
