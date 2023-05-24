package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	totalRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "total_requests",
		Help: "total incoming requests",
	})
	successfulRequests = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "total_successful_requests",
			Help: "request that is successful",
		})
)

func IncRequests(){
	totalRequests.Add(1)
}

func IncSuccessfulRequests(){
	successfulRequests.Add(1)
}



