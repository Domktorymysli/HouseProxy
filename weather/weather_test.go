package weather

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	client := NewClient("aaaa")

	response, _ := client.GetForecast("warsaw")

	fmt.Println(response)

}
