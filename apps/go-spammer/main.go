package main

import (
	"fmt"
	"net/http"
	"sort"
	"time"
)

func main() {
	env := newEnv()

	fmt.Println("Concurrency:", env.Concurrency)
	fmt.Println("URL:", env.Url)

	concurrency := env.Concurrency
	callUrl := env.Url

	for {
		fmt.Println("Starting to spam", callUrl, "with", concurrency, "requests")
		spam(concurrency, callUrl)
		time.Sleep(2 * time.Second)
	}
}

func spam(concurrency int, callUrl string) {
	responseTimes := make(chan time.Duration, concurrency)
	done := make(chan bool)

	runStartTime := time.Now()
	for i := 0; i < concurrency; i++ {
		go func() {
			startTime := time.Now()

			resp, err := http.Get(callUrl)
			if err != nil {
				done <- true
				return
			}
			defer resp.Body.Close()

			duration := time.Since(startTime)
			responseTimes <- duration

			done <- true
		}()
	}

	var runTotalTime time.Duration
	go func() {
		for i := 0; i < concurrency; i++ {
			<-done
		}
		close(responseTimes)
		runTotalTime = time.Since(runStartTime)
	}()

	responseTimeSlice := make([]time.Duration, 0)

	for duration := range responseTimes {
		responseTimeSlice = append(responseTimeSlice, duration)
	}

	sort.Slice(responseTimeSlice, func(i, j int) bool {
		return responseTimeSlice[i] < responseTimeSlice[j]
	})

	averageTime := average(responseTimeSlice)
	quantile10 := calculateQuantile(responseTimeSlice, 0.1)
	quantile90 := calculateQuantile(responseTimeSlice, 0.9)
	medianTime := median(responseTimeSlice)

	fmt.Printf("Run of requests to Requests to %s\n", callUrl)
	fmt.Println("Total Time:", runTotalTime)
	fmt.Println("Average response time:", averageTime)
	fmt.Println("Median response time:", medianTime)
	fmt.Println("0.1 Quantile:", quantile10)
	fmt.Println("0.9 Quantile:", quantile90)
}

func calculateQuantile(responseTimes []time.Duration, quantile float64) time.Duration {
	if len(responseTimes) == 0 {
		return 0
	}

	index := int(float64(len(responseTimes)) * quantile)
	return responseTimes[index]
}

func median(responseTimes []time.Duration) time.Duration {
	if len(responseTimes) == 0 {
		return 0
	}

	if len(responseTimes)%2 == 1 {
		return responseTimes[len(responseTimes)/2]
	}

	return (responseTimes[len(responseTimes)/2-1] + responseTimes[len(responseTimes)/2]) / 2
}

func average(responseTimes []time.Duration) time.Duration {
	if len(responseTimes) == 0 {
		return 0
	}

	total := time.Duration(0)
	for _, duration := range responseTimes {
		total += duration
	}

	return total / time.Duration(len(responseTimes))
}
