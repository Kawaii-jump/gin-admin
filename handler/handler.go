package handler

import (
	"time"

	"github.com/Kawaii-jump/gin-admin/models"
	"github.com/gin-gonic/gin"
)

//HandleLogin doc
// @Summary 登录
// @Description 通过用户名和密码登录
// @Tags 登录
// @Accept mpfd
// @Produce json
// @Param user body models.User true "用户登录信息"
// @Success 200 {object} models.LoginResponse 返回信息
// @Failure 500 {string} string 错误信息
// @Router /login [post]
func HandleLogin(ctx *gin.Context) {

	var user models.User

	if err := ctx.BindJSON(&user); err == nil {
		if user.UserName != "haha" || user.Passworld != "haha" {
			ctx.JSON(200, &models.LoginResponse{Message: "用户名或密码错误！", Token: ""})
			return
		} else {
			ctx.JSON(200, &models.LoginResponse{Message: "登录成功！", Token: user.UserName})
		}
	} else {
		ctx.JSON(400, "格式错误")
	}
}

//HandleGetUserInfo doc
// @Summary 登录
// @Description 通过用户名和密码登录
// @Tags 登录
// @Accept mpfd
// @Produce json
// @Param user body models.User true "用户登录信息"
// @Success 200 {object} models.LoginResponse 返回信息
// @Failure 500 {string} string 错误信息
// @Router /login [post]
func HandleGetUserInfo(ctx *gin.Context) {
	ctx.JSON(200, &models.UserInfo{
		Name:   "test",
		UserID: "10086",
		Access: []string{"test"},
		Token:  "haha",
		Avator: "https://file.iviewui.com/dist/a0e88e83800f138b94d2414621bd9704.png",
	})
}

//HandleRoot handle root
// @Summary 接口验证
// @Description 通过请求验证请求
// @Tags 验证
// @Accept json
// @Produce json
// @Success 200 {string} string 反馈信息
// @Failure 500 {string} string 错误信息
// @Router / [get]
func HandleRoot(ctx *gin.Context) {
	ctx.JSON(200, "success")
}

//HandleSearch handle search interface
// @Summary grafana 查询接口
// @Description 查询grafana matrics 数据
// @Tags 查询接口
// @Accept json
// @Produce json
// @Param searchRequest body models.SearchRequest true "查询请求"
// @Success 200 {object} models.SearchMapResponse 返回信息
// @Failure 500 {string} string 错误信息
// @Router /search [post]
func HandleSearch(ctx *gin.Context) {
	metricsArrary := []string{"health_nodes", "all_nodes"}
	var searchRequest models.SearchRequest

	if err := ctx.BindJSON(&searchRequest); err == nil {
		searchResponse := make(models.SearchMapResponse, 2)
		for i, metrics := range metricsArrary {
			searchResponse = append(searchResponse, models.SearchData{Text: metrics, Value: i})
		}
		ctx.JSON(200, searchResponse)
	} else {
		ctx.JSON(400, "Bad request")
	}
	return
}

//HandleQuery handle query interface
// @Summary grafana 数据查询接口
// @Description 查询grafana 图表 数据
// @Tags 查询接口
// @Accept json
// @Produce json
// @Param searchRequest body models.SearchRequest true "查询请求"
// @Success 200 {object} models.SearchMapResponse 返回信息
// @Failure 500 {string} string 错误信息
// @Router /search [post]
func HandleQuery(ctx *gin.Context) {
	var queryRequest models.QueryRequest
	if err := ctx.BindJSON(&queryRequest); err == nil {
		var queryResponse models.QueryResponse
		for _, target := range queryRequest.Targets {
			if target.Type == "timeserie" {
				querData := &models.QueryData{
					Target:     target.Target,
					Datapoints: [][]interface{}{{1, time.Now().UnixNano()}},
				}
				queryResponse = append(queryResponse, *querData)
			}
		}
		ctx.JSON(200, queryResponse)
	} else {
		ctx.JSON(400, "bad request")
	}
	return
}
