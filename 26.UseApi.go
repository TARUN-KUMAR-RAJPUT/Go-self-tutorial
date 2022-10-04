package main

import (
	"encoding/json"
	"fmt"
	// "net/http"
	// "strconv"
	// "strings"

	// "io/ioutil"
	// "net/http"
	"os"
	// "strconv"
	// "strings"
	"time"
	// "reflect"
	// "github.com/google/uuid"
	// "strconv"
	// "github.com/fatih/structs"

	// "bytes"
	// "encoding/json"
	// "fmt"
	"io/ioutil"
	// "mime/multipart"
	"net/http"

	mailer "../api/mail"
	db "../components/db"
	setup "../setup"
)

// type StsData struct {
//     EMPID   string  `json:"empid"`
//     TOKEN   string  `json:"token"`
//     AK      string  `json:"AK"`
//     GLID    string  `json:"glid"`
//     MODID   string  `json:"modid"`
//     COMMENT string  `json:"comment"`
// }

type ResponseSTS struct {
    Status  string  `json:"status"`
    Code    string  `json:"code"`
    Data    string  `json:"data"`
    Message string  `json:"message"`
}

var logs = make(map[string]interface{})


func Send_mail(mailsubject string, mailfrom string, mailto string, mailBody string, mailcc string, IsBodyHTML string, MailCategoryName string) bool {
	
	fmt.Println("Sending Mail...")

	var mailData mailer.MailData
	mailData.TO = mailto
	mailData.MAILFROM = mailfrom
	mailData.SUBJECT = mailsubject
	mailData.BODY = mailBody

	mailData.MAILFROMNAME = "Bulk Stop Mandate"
	mailData.MODID = "WEBERP"
	mailData.HTMLFORMAT = IsBodyHTML
	mailData.CATEGORY = MailCategoryName

	mailResponse := mailer.SendMail(mailData)

	if mailResponse.Status != "Success" {
		return false
	}
    
	fmt.Println("Mail Sent !")

	return true
}


func Rabbit_log(logs map[string]interface{}) {

	fmt.Println("Creating Log...")
	
	today_date := time.Now().Format("2006-01-02")
	file_name := "bulk_stop_mandate_log_" + today_date + ".txt"
	// file_name = "/home/indiamart/public_html/dev-weberp-auto-dialer/marketing/logs/" + file_name
	file_name = "/Projects/Indiamart/Go Scheduler/mylogs/" + file_name
	
	
	file, errf := os.OpenFile(file_name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if errf != nil {
		fmt.Println("error in file permission", errf)
	} else {
		log, err := json.Marshal(logs)
		if err == nil {
			defer file.Close()
			file.WriteString(string(log) + "\n" + "\n")
		} else {
			fmt.Println(err)
		}
	}

	fmt.Println("Log Created !")
}

func UpdateSTSHist() ResponseSTS {

    var stsResponse ResponseSTS
	
	url := "https://dev-merp.intermesh.net/index.php/Userlist/UpdateStsHistory?"
	method := "POST"

	empid := "63509"  //
	token := "erp@16072018"  //done
	AK := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI5NDM2OCIsImV4cCI6MTY2MDcxMjk2MSwiaWF0IjoxNjYwNjI2NTYxLCJpc3MiOiJFTVBMT1lFRSJ9.q-pKN5g9fSG3X5E76Sex1HvXsdev8Q8flH6h4SLeuuI"
	glid := "2003749" //done
	modid := "weberp" //done
	comment := "TestingbyTarun"  // done

	url += "empid=" + empid + "&"
	url += "token=" + token + "&"
	url += "AK=" + AK + "&"
	url += "glid=" + glid + "&"
	url += "modid=" + modid + "&"
	url += "comment=" + comment

	fmt.Println(url)

	client := &http.Client {}

	req, err := http.NewRequest(method, url, nil)
  
	if err != nil {
	  fmt.Println(err)
	}

	req.Header.Add("Content-Type", "application/json")
  
	res, err := client.Do(req)
	
	if err != nil {
	  fmt.Println(err)
	}

	defer res.Body.Close()
  
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
	  fmt.Println(err)
	}

	json.Unmarshal(body, &stsResponse)
    
	return stsResponse
}

func main() {

	setup.StartSetup()
	dbMain, err := db.GetDatabaseConnection("main")

	if err != nil {
		logs["main_read_connection_error"] = err
		fmt.Println("Could not start the main database because of following err " + err.Error() + " ...!!!")

		// Send_mail("Error in scheduler BulkStopMandate while fetching connectionstring", "weberperrors@intermesh.net", "weberperrors@intermesh.net", " Issue occurred on :  "+time.Now().Format("02-01-2006_01:01:01")+"\n"+"Error : "+err.Error(),
		// 	"", "false", "Weberp-schedulers")

		Send_mail("Error in scheduler BulkStopMandate while fetching connectionstring","jatin.minocha@indiamart.com","tarun.rajput@indiamart.com"," Issue occurred on :  " + time.Now().Format("02-01-2006_01:01:01") + "\n" + "Error : " + err.Error(),
		"", "false","Weberp-schedulers")	

		Rabbit_log(logs)
		os.Exit(0)
	} else {
		logs["main_read_connection"] = "successful connection"
		Rabbit_log(logs)
	}
	err = dbMain.Ping()
	if err != nil {
		logs["main_read_connection_error"] = err
		fmt.Println("Could not start the main database because of following err " + err.Error() + " ...!!!")

		// Send_mail("Error in scheduler BulkStopMandate while fetching connectionstring", "weberperrors@intermesh.net", "weberperrors@intermesh.net", " Issue occurred on :  "+time.Now().Format("02-01-2006_01:01:01")+"\n"+"Error : "+err.Error(),
		// 	"", "false", "Weberp-schedulers")

		Send_mail("Error in scheduler BulkStopMandate while fetching connectionstring","jatin.minocha@indiamart.com","tarun.rajput@indiamart.com"," Issue occurred on :  " + time.Now().Format("02-01-2006_01:01:01") + "\n" + "Error : " + err.Error(),
		"", "false","Weberp-schedulers")	

		Rabbit_log(logs)
		os.Exit(0)
	} else {
		logs["main_read_connection"] = "successful connection"
		Rabbit_log(logs)
	}

	fmt.Println("Connection to the main-database established ...!!!")
	fmt.Println("setup completed ")
    
	resp := UpdateSTSHist()
	fmt.Println(resp)
}




