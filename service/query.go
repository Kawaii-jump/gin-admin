package service

import (
	"net"
	"time"

	"github.com/Kawaii-jump/gin-admin/logger"
	"github.com/Kawaii-jump/gin-admin/models"
)

var queue *Queue
var services []string

func init() {
	queue = NewQueue()
	services = []string{"localhost:8181", "localhost:8181", "localhost:8081"}
}

//GetQueryDatas get query result
func GetQueryDatas(targatType string) []models.QueryData {
	if targatType == "timeserie" {
		return queue.list
	}
	return []models.QueryData{}
}

//ProduceDatas produce query data
func ProduceDatas(period time.Duration) {
	t := time.NewTicker(period)
	for {
		select {
		case <-t.C:
			tt := time.Now().Unix() * 1000
			if len(queue.list) < 1000 {
				res := GetNodesStatus(tt)
				for _, data := range res {
					queue.Push(data)
				}
			} else {
				queue.Pop()
				queue.Pop()
				res := GetNodesStatus(tt)
				for _, data := range res {
					queue.Push(data)
				}
			}
		case <-queue.ctx.Done():
			return
		}
	}
}

//GetNodesStatus get nodes status
func GetNodesStatus(t int64) (res [2]models.QueryData) {
	health_nodes := 0
	all_nodes := len(services)
	for _, server := range services {
		if judgeLive(server) {
			health_nodes++
		}
	}
	res[0] = models.QueryData{
		Target:     "health_nodes",
		Datapoints: [][]interface{}{{health_nodes, t}},
	}
	res[1] = models.QueryData{
		Target:     "all_nodes",
		Datapoints: [][]interface{}{{all_nodes, t}},
	}
	return
}

//GetNodesStatusToTable get mq nodes to grafana table
func GetNodesStatusToTable(target string) (columns []models.TableColumn, rows [][]interface{}) {
	if target == "node" {
		columns = append(columns, models.TableColumn{
			Text: "Server",
			Type: "string",
		})
		columns = append(columns, models.TableColumn{
			Text: "Status",
			Type: "number",
		})
	}

	for _, server := range services {
		if judgeLive(server) {
			rows = append(rows, []interface{}{server, 1})
		} else {
			rows = append(rows, []interface{}{server, 0})
		}
	}
	return
}

func judgeLive(server string) bool {
	conn, err := net.DialTimeout("tcp", server, 3*time.Second)
	if err != nil {
		logger.Errorf("dial server port error,err:", err)
		return false
	}
	if conn != nil {
		return true
	}
	return false
}
