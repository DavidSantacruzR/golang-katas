package webscraping

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sync"
	"testing"
	"time"
)

func startApi() {
	RunApi()
}

func makeRequest(t *testing.T, wg *sync.WaitGroup) {
	defer wg.Done()
	requestUrl := "http://localhost:8080/vehicleData"
	jsonBody := []byte(`{"plate": "ABC123", "OwnerId": "123456789"}`)
	bodyReader := bytes.NewReader(jsonBody)
	response, err := http.Post(requestUrl, "application/json", bodyReader)
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()
	bodyBytes, bytesErr := io.ReadAll(response.Body)
	if err != nil {
		t.Error(bytesErr)
	}
	t.Logf("Completed request with status %s and body %s\n", response.Status, string(bodyBytes))
	fmt.Println("")
}

func TestPerformanceLowTraffic(t *testing.T) {
	var wg sync.WaitGroup
	requestsNumber := 10
	go startApi()
	time.Sleep(1 * time.Second)
	wg.Add(requestsNumber)
	for i := 0; i < requestsNumber; i++ {
		go makeRequest(t, &wg)
	}
	wg.Wait()
	t.Logf("%d Requests completed successfully", 10)
}

func TestPerformanceMediumTraffic(t *testing.T) {
	var wg sync.WaitGroup
	requestsNumber := 50
	go startApi()
	time.Sleep(1 * time.Second)
	wg.Add(requestsNumber)
	for i := 0; i < requestsNumber; i++ {
		go makeRequest(t, &wg)
	}
	wg.Wait()
	t.Logf("%d Requests completed successfully", 10)
}
