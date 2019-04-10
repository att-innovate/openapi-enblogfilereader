package core

import (
	"fmt"
	"log"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
)

const (
	INFLUX_SERVER = "http://192.168.1.201"
	INFLUX_PORT   = "8086"
	INFLUX_DBNAME = ""
	INFLUX_TABLE  = ""

	INFLUX_USERNAME = ""
	INFLUX_PASSWORD = ""
)

func PushENBStats(enb Enb) {
	//log.Printf("perform db write action")

	clnt, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     fmt.Sprintf("%s:%s", INFLUX_SERVER, INFLUX_PORT),
		Username: INFLUX_USERNAME,
		Password: INFLUX_PASSWORD,
	})

	if err != nil {
		log.Printf("Error in instantiating inflixdb client %v.", err)
	}

	defer clnt.Close()

	// Create a new point batch
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  INFLUX_DBNAME,
		Precision: "s",
	})

	// Create a point and add to batch
	tags := map[string]string{"ueid": fmt.Sprintf("%v", enb.UEid)}
	fields := map[string]interface{}{
		"ulbr": enb.ULbr,
		"dlbr": enb.DLbr,
		"snr":  enb.SNR,
	}

	pt, err := client.NewPoint(INFLUX_TABLE, tags, fields, time.Now())
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	bp.AddPoint(pt)
	// Write the batch
	if err := clnt.Write(bp); err != nil {
		log.Fatalln("Error: ", err)
	} else {
		//log.Println("write complete")
	}
}
