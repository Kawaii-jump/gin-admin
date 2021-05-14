package handler

import (
	"time"

	"github.com/Kawaii-jump/gin-admin/logger"
	"github.com/Kawaii-jump/gin-admin/models"
	"github.com/Kawaii-jump/gin-admin/service"
	"github.com/gin-gonic/gin"
)

//HandleRoot handle root
// @Summary 接口验证
// @Description 通过请求验证请求
// @Tags grafana接口
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
		logger.Errorf("bind json error,err:", err)
		ctx.JSON(400, "Bad request")
	}
	return
}

//HandleQuery handle query interface
// @Summary grafana 数据查询接口
// @Description 查询grafana 图表 数据
// @Tags grafana接口
// @Accept json
// @Produce json
// @Param queryRequest body models.QueryRequest true "查询请求"
// @Success 200 {object} models.QueryResponse 返回信息
// @Failure 500 {string} string 错误信息
// @Router /query [post]
func HandleQuery(ctx *gin.Context) {
	var queryRequest models.QueryRequest
	if err := ctx.BindJSON(&queryRequest); err == nil {
		var queryResponse models.QueryResponse
		for _, target := range queryRequest.Targets {
			if target.Type == "timeserie" {
				queryResponse = service.GetQueryDatas(target.Type)
			}
		}
		ctx.JSON(200, queryResponse)
	} else {
		logger.Errorf("bind json error,err:", err)
		ctx.JSON(400, "bad request")
	}
	return
}

//HandleAnnotation handle annotation interface
// @Summary grafana Annotation 接口
// @Description grafana Annotation 接口
// @Tags grafana接口
// @Accept json
// @Produce json
// @Param queryRequest body models.AnnotationRequest true "请求"
// @Success 200 {object} models.AnnotationResponse 返回信息
// @Failure 500 {string} string 错误信息
// @Router /query [post]
func HandleAnnotation(ctx *gin.Context) {
	var annotation models.AnnotationRequest
	if err := ctx.BindJSON(&annotation); err == nil {
		ctx.JSON(200, &models.AnnotationResponse{
			Time:  time.Now().Unix() * 1000,
			Title: "grafana",
		})
	}
}
