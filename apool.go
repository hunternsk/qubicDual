package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getQubicEpochChallengeIdle() (bool, error) {
	req, err := http.NewRequest("GET", epochUrl, nil)
	if err != nil {
		fmt.Println("req", err)
		return false, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("resp", err)
		return false, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read", err)
		return false, err
	}

	type QubicEpochChallenge struct {
		Code   int `json:"code"`
		Result struct {
			Epoch            int    `json:"epoch"`
			PoolThreshold    int    `json:"pool_threshold"`
			NetworkThreshold int    `json:"network_threshold"`
			MiningSeed       string `json:"mining_seed"`
			Timestamp        int    `json:"timestamp"`
		} `json:"result"`
		Msg string `json:"msg"`
	}
	var epochChallenge QubicEpochChallenge
	err = json.Unmarshal(body, &epochChallenge)
	if err != nil {
		fmt.Println("unmarsh", err)
		return false, err
	}
	if epochChallenge.Result.MiningSeed == "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=" {
		return true, nil
	}
	return false, nil
}
