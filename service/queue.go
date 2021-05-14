package service

import (
	"context"
	"sync"

	"github.com/Kawaii-jump/gin-admin/models"
)

//Queue slice queue
type Queue struct {
	sync.RWMutex
	ctx  context.Context
	list []models.QueryData
}

//NewQueue create queue
func NewQueue() *Queue {
	list := make([]models.QueryData, 0)
	return &Queue{list: list}
}

//Push add struct to queue
func (q *Queue) Push(data models.QueryData) {
	q.RLock()
	defer q.RUnlock()
	q.list = append(q.list, data)
}

//Pop delete frist struct and return
func (q *Queue) Pop() models.QueryData {
	if len(q.list) == 0 {
		return models.QueryData{}
	}
	q.RLock()
	defer q.RUnlock()
	res := q.list[0]
	q.list = q.list[1:]
	return res
}
