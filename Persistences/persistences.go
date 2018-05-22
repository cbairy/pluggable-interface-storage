package Persistences

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"pluggable-interface-storage/Pojo"
	"time"
)

type Persistence interface {
	GetChosenPersistence() string
	SetActivity(Pojo.Activity)
	GetActivities() []Pojo.Activity
}

type Db struct {
	name string
}

func (db *Db) GetChosenPersistence() string {
	return "couchdb chosen for audit trail"
}

func (db *Db) SetActivity(activity Pojo.Activity) {
	db.name = "couchdb"
}

func (db *Db) GetActivities() []Pojo.Activity {
	return nil
}

type Bc struct {
	name string
}

func (bc *Bc) GetChosenPersistence() string {
	return "hyperledger chosen for audit trail"
}

func (bc *Bc) SetActivity(activity Pojo.Activity) {
	bc.name = "hyperledger"
	jsonActivity, e := json.Marshal(activity)
	if e != nil {
		panic(e)
	}
	fmt.Println("json:", jsonActivity)
	start := time.Now()
	resp, err := http.Post("http://localhost:3000/api/Activity", "application/json", bytes.NewBuffer(jsonActivity))
	end := time.Now()
	elapsed := end.Sub(start)
	if err != nil {
		panic(err)
	}
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	fmt.Println("response time for Post: ", elapsed)
}

func (bc *Bc) GetActivities() []Pojo.Activity {
	var activities []Pojo.Activity
	fmt.Println("getting activities")
	start := time.Now()
	resp, err := http.Get("http://localhost:3000/api/Activity")
	end := time.Now()
	elapsed := end.Sub(start)
	if err != nil {
		panic(err)
	}
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	fmt.Println("response time for Get: ", elapsed)
	json.Unmarshal(body, &activities)
	return activities
}

const (
	DB = iota
	BC
)

func CreatePersistence(pluggableType uint64) (Persistence, error) {
	switch pluggableType {
	case DB:
		return new(Db), nil
	case BC:
		return new(Bc), nil
	default:
		return nil, errors.New("Invalid Pluggable Type")
	}
}
