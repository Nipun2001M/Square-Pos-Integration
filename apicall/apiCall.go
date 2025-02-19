package apicall

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const baseURL = "https://connect.squareupsandbox.com/v2/"

type ClientSq struct {
	AcessToken string
	client    *http.Client
}

func GetClient() *ClientSq{
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &ClientSq{
		AcessToken: os.Getenv("SQUARE_ACCESS_TOKEN"),
		client: &http.Client{Timeout: 10 * time.Second},
	}
}
func (c *ClientSq) ApiCall(method string, endpoint string, data interface{}) ([]byte, error)  {
	url := baseURL + endpoint
	fmt.Println("url->",url)
	var req *http.Request
	var err error
	if method == http.MethodGet {
		req, err = http.NewRequest(method, url, nil)
	} else {
		reqb, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return nil, fmt.Errorf("error marshaling request data")
		}
		req, err = http.NewRequest(method, url, bytes.NewBuffer(reqb))
	}
	if err!=nil{
		fmt.Println("error in creating req",err)
		return nil,fmt.Errorf("error in creating request")

	}
	req.Header.Set("Square-Version", "2025-01-23")
	req.Header.Set("Authorization", "Bearer "+c.AcessToken)
	req.Header.Set("Content-Type", "application/json")
	res,errreq:=c.client.Do(req)
	if errreq!=nil{
		log.Fatal("error in make request")
		return nil,fmt.Errorf("error in make request")
	}

	returnData,_:=io.ReadAll(res.Body)
	return returnData,nil



}