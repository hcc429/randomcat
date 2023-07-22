package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	successfulRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "requests",
		ConstLabels: map[string]string{"success": "true"},
	})
	failedRequests = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "requests",
			ConstLabels: map[string]string{"success": "false"},
	})
	// TODO: add metric to record cache miss
)



func IncSuccessfulRequests(){
	successfulRequests.Add(1)
}

func IncFailedRequests(){
	failedRequests.Add(1)
}



