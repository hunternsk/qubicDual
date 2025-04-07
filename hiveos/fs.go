package hiveos

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var fsUrl = "/farms/%s/fs"

func (h *HiveOS) GetFses() (*fsType, error) {
	req, err := http.NewRequest("GET", baseUrl+fmt.Sprintf(fsUrl, h.farmID), nil)
	if err != nil {
		fmt.Println("req", err)
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+h.accessToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("resp", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read", err)
		return nil, err
	}

	var fses fsType
	err = json.Unmarshal(body, &fses)
	if err != nil {
		fmt.Println("unmarsh", err)
		return nil, err
	}

	return &fses, nil
}

type fsType struct {
	Data []struct {
		ID           int  `json:"id"`
		FarmID       int  `json:"farm_id"`
		IsFavorite   bool `json:"is_favorite"`
		WorkersCount int  `json:"workers_count"`
		AppliedAt    int  `json:"applied_at"`
		Items        []struct {
			Coin        string   `json:"coin"`
			Pool        string   `json:"pool"`
			PoolGeo     []string `json:"pool_geo"`
			PoolSsl     bool     `json:"pool_ssl"`
			PoolUrls    []string `json:"pool_urls"`
			WalID       int      `json:"wal_id"`
			DpoolSsl    bool     `json:"dpool_ssl"`
			Miner       string   `json:"miner"`
			MinerConfig struct {
				URL      string `json:"url"`
				Algo     string `json:"algo"`
				Pass     string `json:"pass"`
				Template string `json:"template"`
			} `json:"miner_config"`
		} `json:"items"`
		Name string `json:"name,omitempty"`
	} `json:"data"`
}
