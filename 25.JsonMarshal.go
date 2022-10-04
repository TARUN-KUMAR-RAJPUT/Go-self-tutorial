package main
import (
    "fmt"
    "encoding/json"
)


type StsData struct {
    EMPID   string  `json:"empid"`
    TOKEN   string  `json:"token"`
    AK      string  `json:"AK"`
    GLID    string  `json:"glid"`
    MODID   string  `json:"modid"`
    COMMENT string  `json:"comment"`
}

func main() {
    
    var stsData StsData

    stsData.EMPID = "94368"
	stsData.TOKEN = "erp@16072018"
	stsData.AK = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI5NDM2OCIsImV4cCI6MTY2MDI5ODY4MiwiaWF0IjoxNjYwMjEyMjgyLCJpc3MiOiJFTVBMT1lFRSJ9.TV5M0-oMbB5A4YM7G6SutoGirVMkWpvQSAASbNVI-Xs"
	stsData.GLID = "28243970"
	stsData.MODID = "weberp"
	stsData.COMMENT = "Hello we are testing"

    bytes, _ := json.Marshal(stsData)

    fmt.Println(string(bytes))
}