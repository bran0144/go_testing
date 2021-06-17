package services

import "fmt"

type pingService struct{}

const (
	pong = "poing"
)

var (
	PingService = pingService{}
)

func (service pingService) HandlePing() (string, error) {
	fmt.Println("doing somthing")
	return pong, nil
}
