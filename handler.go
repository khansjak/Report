package pck

import (
	"bitbucket.org/Myself/goProject/src/bitbucket.org/tekion/tbaas/tapi"
	"bitbucket.org/tekion/tbaas/apiContext"
	"encoding/csv"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	//"time"
	mMgr "bitbucket.org/tekion/tbaas/mongoManager"
)

func getReport(w http.ResponseWriter,r *http.Request)  {
	ctx := apiContext.UpgradeCtx(r.Context())
	tomorrowStart :=time.Now()
	start:=tomorrowStart.AddDate(0,0,-1)
	tomorrowEnd :=   time.Now()
	var dealer []Dealer
	selectQ:=bson.M{"_id":1,"dealerName":1}
	err:=mMgr.ReadAll(ctx.Tenant,"DealerMaster",nil,selectQ,&dealer)
	if err!=nil{
		fmt.Errorf("failed to fetch dealer ",err)
		fmt.Println("error")
		return
	}
	var details []AppointmentOutput
	findQ:=bson.M{"appointmentDateTime":bson.M{"$gte": start, "$lte": tomorrowEnd}}
	selectQ=bson.M{"dealerID":1,"serviceAdvisorName":1,"lastName":1,"firstName":1,"appointmentDateTime":1,"lastUpdatedByDisplay":1,"status":1}
	err=mMgr.ReadAll(ctx.Tenant,"Appointment",findQ,selectQ,&details)
	if err!=nil{
		fmt.Errorf("failed to fetch appointment ",err)
		fmt.Println("Error")
		return
	}
	var report []Report
	for _,dm:=range dealer{
		for _,app:=range details{
			if app.ID==dm.Id{
				var test Report
				test.Dealer=dm.Name
				test.CustomerName= app.LastName+" "+app.FirstName
				test.ID=app.ID
				test.AppointmentDateTime=app.AppointmentDateTime
				test.AppointmentStatus=app.AppointmentStatus
				test.LastUpdatedBy=app.LastUpdatedBy
				test.ServiceAdvisor=app.ServiceAdvisor
				report=append(report,test)
			}
		}
	}
	csvdatafile, err := os.Create("/Users/paviyarasu/go/src/bitbucket.org/tekion/javed/data.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvdatafile.Close()

	writer := csv.NewWriter(csvdatafile)

	var headerRow Report
	val := reflect.Indirect(reflect.ValueOf(headerRow))
	fmt.Println(val.Type().Field(0).Name)
	fmt.Println(val.Type().Field(1).Name)
	fmt.Println(val.Type().Field(2).Name)
	fmt.Println(val.Type().Field(3).Name)
	var record []string
	record = append(record, "ID")
	record = append(record, "AppointmentStatus")
	record = append(record, "Dealer")
	record = append(record, "ServiceAdvisor")
	record = append(record, "LastUpdatedBy")
	record = append(record, "AppointmentDateTime.String()")
	record = append(record,"CustomerName")
	writer.Write(record)


	for _, worker := range report {
		var record []string
		record = append(record, worker.ID)
		record = append(record, strconv.Itoa(worker.AppointmentStatus))
		record = append(record, worker.Dealer)
		record = append(record, worker.ServiceAdvisor)
		record = append(record, worker.LastUpdatedBy)
		record = append(record, worker.AppointmentDateTime.String())
		record = append(record, worker.CustomerName)
		writer.Write(record)
	}

	// remember to flush!
	writer.Flush()



	tapi.HTTPResponse(ctx,w,http.StatusOK,"Report created",report)

}

