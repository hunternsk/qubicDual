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
	args := os.Args[1:]
	if len(args) > 0 {
		switch arg := args[0]; arg {
		case "-fs":
			accessToken := promptUser("Access Token")
			farmId := promptUser("Farm ID")
			hiveos := hiveos.New(farmId, accessToken)

			fses, err := hiveos.GetFses()
			if err != nil {
				fmt.Println("GetFses", err)
				os.Exit(1)
			}
			fmt.Printf("ID | Coins | Name\n")
			for _, fs := range fses.Data {
				var coins []string
				for _, item := range fs.Items {
					coins = append(coins, item.Coin)
				}
				fmt.Printf("%d | %s | %s\n", fs.ID, strings.Join(coins, ","), fs.Name)
			}
			os.Exit(0)
		}
	}

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

				/* if !w2.Stats.Online {
					continue
				} */
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

			setFsId := qubicFsId
			workerIds := []int{}
			idle, err := getQubicEpochChallengeIdle()
			if err == nil {
				if idle {
					setFsId = idleFsId
				}
				for id, worker := range workers.GetAll() {
					if worker.FsId != setFsId {
						workerIds = append(workerIds, id)
					}

				}
				if len(workerIds) > 0 {
					result, err := hiveos.SetWorkersData(map[string]interface{}{"fs_id": setFsId}, workerIds)
					if err != nil {
						time.Sleep(time.Second * 5)
						break
					}
					if result {
						for _, id := range workerIds {
							workers.SetFs(id, setFsId)
						}
					}
				}

			}
			time.Sleep(time.Second * 5)
		}
	}()

	<-done
}
