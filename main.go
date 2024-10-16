package main

import (
	"fmt"
	"log"
	"os"
	"qubicDual/farm"
	"qubicDual/hiveos"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const (
	epochUrl = "http://qubic1.hk.apool.io:8001/api/qubic/epoch_challenge"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accessToken := os.Getenv("ACCESSTOKEN")
	if accessToken == "" {
		log.Fatal("Set ACCESSTOKEN")
	}

	farmId := os.Getenv("FARMID")
	if farmId == "" {
		log.Fatal("Set FARMID")
	}

	qubicFsIdStr := os.Getenv("QUBICFSID")
	if qubicFsIdStr == "" {
		log.Fatal("Set QUBICFSID")
	}
	qubicFsId, err := strconv.Atoi(qubicFsIdStr)
	if err != nil {
		log.Fatalln("QUBICFSID", err)
	}

	idleFsIdStr := os.Getenv("IDLEFSID")
	if qubicFsIdStr == "" {
		log.Fatal("Set IDLEFSID")
	}
	idleFsId, err := strconv.Atoi(idleFsIdStr)
	if err != nil {
		log.Fatalln("IDLEFSID", err)
	}

	excludeWorkersStr := os.Getenv("EXCLUDEWORKERS")
	excludeWorkers := strings.Split(excludeWorkersStr, ",")
	includeWorkersStr := os.Getenv("INCLUDEWORKERS")
	includeWorkers := strings.Split(includeWorkersStr, ",")

	done := make(chan struct{})
	workers := farm.NewWorkers()
	hiveos := hiveos.New(farmId, accessToken)

	go func() {
		for {
			workers2, err := hiveos.GetWorkers2()
			if err != nil {
				log.Println(err)
				continue
			}
			for _, w2 := range workers2.Data {
				if !slices.Contains(includeWorkers, w2.Name) {
					if excludeWorkersStr == "*" || slices.Contains(excludeWorkers, w2.Name) {
						continue
					}
				}

				if !w2.Stats.Online {
					continue
				}
				workers.Store(w2.ID, farm.WorkerType{
					Name: w2.Name,
					FsId: w2.FlightSheet.ID,
				})
			}

			time.Sleep(time.Second * 30)
		}
	}()

	go func() {
		for {
			if workers.Len() == 0 {
				continue
			}
			idle, err := getQubicEpochChallengeIdle()
			if err == nil {
				for id, worker := range workers.GetAll() {
					if idle && worker.FsId != idleFsId {
						fmt.Println("Setting Idle FS", idleFsId, "for", worker.Name)
						result, err := hiveos.SetWorkerFs(id, idleFsId)
						if err != nil {
							time.Sleep(time.Second * 5)
							break
						}
						if result {
							workers.SetFs(id, idleFsId)
						}
						time.Sleep(time.Millisecond * 500)
					}
					if !idle && worker.FsId != qubicFsId {
						fmt.Println("Setting Qubic FS", qubicFsId, "for", worker.Name)
						result, err := hiveos.SetWorkerFs(id, qubicFsId)
						if err != nil {
							time.Sleep(time.Second * 5)
							break
						}
						if result {
							workers.SetFs(id, qubicFsId)
						}
						time.Sleep(time.Millisecond * 500)
					}
				}

			}
			time.Sleep(time.Second * 5)
		}
	}()

	<-done
}