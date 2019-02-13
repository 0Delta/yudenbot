package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type rawEventDatas struct {
	EventData []rawEventData `json:"events"`
}

type rawEventData struct {
	ID           int         `json:"id"`
	URL          string      `json:"global_id"`
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	StartDateStr string      `json:"start_date"`
	EndDateStr   string      `json:"end_date"`
	Organizer    []organizer `json:"organizer"`
}

type organizer struct {
	ID        int    `json:"id"`
	Organizer string `json:"organizer"`
	Slug      string `json:"slug"`
}

type EventData struct {
	ID          int
	URL         string
	Title       string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	Organizer   string
}

func GetEventsFromWordpress(url string) (events []EventData, err error) {
	nowt := time.Now()
	sdt := nowt.Format("2006/01/02T15:04")
	edt := nowt.Add(24 * time.Hour).Add(1 * time.Minute).Format("2006/01/02T15:04")
	url = "https://" + url + "/?rest_route=/tribe/events/v1/events" + "&start_date=" + sdt + "&end_date=" + edt
	log.Println("Getting URL : ", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Println("HTTP Get error:", err)
		return
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("HTTP Read error:", err)
		return
	}

	return GetEventDatas(byteArray)
}

func GetEventDatas(jsonBytes []byte) (datas []EventData, err error) {
	rawdatas := new(rawEventDatas)
	err = json.Unmarshal(jsonBytes, rawdatas)
	if err != nil {
		log.Println("JSON Unmarshal error:", err)
		return
	}

	for _, rawdata := range rawdatas.EventData {
		d, err := parseEventData(&rawdata)
		if err != nil {
			log.Println("EventData Parse Error:", err)
			continue
		}
		datas = append(datas, d)
	}
	log.Println(datas)
	return
}

func GetEventData(jsonBytes []byte) (data EventData, err error) {
	rawdata := new(rawEventData)
	err = json.Unmarshal(jsonBytes, rawdata)
	if err != nil {
		log.Println("JSON Unmarshal error:", err)
		return
	}

	return parseEventData(rawdata)
}

func parseEventData(rawdata *rawEventData) (data EventData, err error) {
	data.ID = rawdata.ID
	data.URL = rawdata.URL

	// TODO : trim for post
	data.Title = rawdata.Title
	data.Description = rawdata.Description

	data.StartDate, _ = time.Parse("2006-01-02 15:04:05", rawdata.StartDateStr)
	data.EndDate, _ = time.Parse("2006-01-02 15:04:05", rawdata.EndDateStr)
	if len(rawdata.Organizer) == 0 {
		data.Organizer = ""
	} else {
		data.Organizer = rawdata.Organizer[0].Organizer
	}
	return
}
