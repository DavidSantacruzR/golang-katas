package main

import (
	"fmt"
	"golang/error_handling"
)

func FetchAPIData() (string, error) {
	return "NA", fmt.Errorf("NO DATA AVAILABLE")
}

func main() {
	data, fetchError := FetchAPIData()
	result := error_handling.UnwrapOrElse(data, fetchError, func(e error) string {
		/*
			Make here a retry logic to fetch data from another source. Or a retry if there was a timeout error.
		*/
		return "THERE IS DATA"
	})
	fmt.Println(result)
}
