package main

import (
	"fmt"
	"os"
	"pluggable-interface-storage/Persistences"
	"pluggable-interface-storage/Pojo"
	"strconv"
)

func main() {
	args := os.Args[1:]
	pluggableType, err := strconv.Atoi(args[0])
	fmt.Println(pluggableType)
	if err != nil {
		fmt.Println("7th argument must be a numeric string")
	}

	pluggablePersistence, err := Persistences.CreatePersistence(uint64(pluggableType))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pluggablePersistence.GetChosenPersistence())
		activityType := args[1]
		userId := args[2]
		documentId := args[3]
		viewingTime, err := strconv.Atoi(args[4])
		viewedPages := args[5]
		noOfDownloads, err := strconv.Atoi(args[6])
		loginTime, err := strconv.Atoi(args[7])
		profileUpdate, err := strconv.ParseBool(args[8])
		if err != nil {
			fmt.Println(err)
		}
		activityId := activityType + userId + documentId

		pluggablePersistence.SetActivity(Pojo.Activity{
			ActivityId:    activityId,
			ActivityType:  activityType,
			UserId:        userId,
			DocumentId:    documentId,
			ViewingTime:   uint64(viewingTime),
			ViewedPages:   viewedPages,
			NoOfDownloads: uint64(noOfDownloads),
			LoginTime:     uint64(loginTime),
			ProfileUpdate: profileUpdate,
		})

		storedActivities := pluggablePersistence.GetActivities()
		fmt.Printf("Results: %v\n", storedActivities)
	}
}
