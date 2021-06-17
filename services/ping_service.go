package services

import "fmt"

type pingService interface {
	HandlePing() (string, error)
}
type pingServiceImpl struct{}

const (
	pong = "poing"
)

var (
	PingService pingService = pingServiceImpl{}
)

func (service pingServiceImpl) HandlePing() (string, error) {
	fmt.Println("doing somthing")
	return pong, nil
}
