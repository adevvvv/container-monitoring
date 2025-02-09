package main

import (
	"log"
	"time"

	"container-monitoring/pinger/api"
	"container-monitoring/pinger/container"
	"container-monitoring/pinger/rabbitmq"
)

func main() {
	ips, err := container.GetContainerIPs()
	if err != nil {
		log.Fatalf("Error getting container IPs: %v", err)
	}

	rmq, err := rabbitmq.NewRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rmq.Close()

	for {
		for _, ip := range ips {
			pingTime, err := container.PingContainer(ip)
			if err != nil {
				log.Printf("Error pinging container %s: %v", ip, err)
				continue
			}

			status := api.PingStatus{
				IP:          ip,
				PingTime:    pingTime,
				LastSuccess: time.Now(),
			}

			// Отправляем данные в API
			err = api.SendToAPI("http://backend:8080/api/v1/status", status)
			if err != nil {
				log.Printf("Error sending to API: %v", err)
			}

			// Отправляем данные в RabbitMQ
			err = rmq.SendMessage(status)
			if err != nil {
				log.Printf("Error sending message to RabbitMQ: %v", err)
			}
		}
		time.Sleep(3 * time.Second)
	}
}