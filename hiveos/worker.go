package hiveos

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

func (h *HiveOS) SetWorkerFs(id, fsId int) (bool, error) {
	payload, _ := json.Marshal(map[string]interface{}{
		"fs_id": fmt.Sprintf("%d", fsId),
	})
	req, err := http.NewRequest(http.MethodPatch, baseUrl+fmt.Sprintf("/farms/%s/workers/%d", h.farmID, id), bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("req", err)
		return false, err
	}
	req.Header.Set("Authorization", "Bearer "+h.accessToken)
	req.Header.Set("Content-Type", "application/json")

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
	var hiveOSResponse CommandResponse
	err = json.Unmarshal(body, &hiveOSResponse)
	if err != nil {
		fmt.Println("unmarsh", err)
		return false, err
	}
	if len(hiveOSResponse.Commands) > 0 {
		fmt.Println("Change done")
		return true, nil
	}
	fmt.Println("Change goes wrong")
	spew.Dump(hiveOSResponse)
	return false, nil
}

func (h *HiveOS) SetWorkersData(data map[string]interface{}, workerIds []int) (bool, error) {
	payload, _ := json.Marshal(map[string]interface{}{
		"data":       data,
		"worker_ids": workerIds,
	})
	req, err := http.NewRequest(http.MethodPatch, baseUrl+fmt.Sprintf("/farms/%s/workers", h.farmID), bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("req", err)
		return false, err
	}
	fmt.Println("Change req payload:", fmt.Sprintf("%v", req.Body))
	req.Header.Set("Authorization", "Bearer "+h.accessToken)
	req.Header.Set("Content-Type", "application/json")

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
	var hiveOSResponse CommandResponse
	err = json.Unmarshal(body, &hiveOSResponse)
	if err != nil {
		fmt.Println("unmarsh", err)
		return false, err
	}
	if len(hiveOSResponse.Commands) > 0 {
		fmt.Println("Change done")
		return true, nil
	}
	fmt.Println("Change goes wrong")
	spew.Dump(hiveOSResponse)
	return false, nil
}

func (h *HiveOS) CleanWorkersMessages(workerIds []int, types []string) (bool, error) {
	payload, _ := json.Marshal(map[string]interface{}{
		"worker_ids": workerIds,
		"types": types,
	})
	req, err := http.NewRequest(http.MethodDelete, baseUrl+fmt.Sprintf("/farms/%s/workers/messages", h.farmID), bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("req", err)
		return false, err
	}
	fmt.Println("Clean req payload:", fmt.Sprintf("%v", req.Body))
	req.Header.Set("Authorization", "Bearer "+h.accessToken)
	req.Header.Set("Content-Type", "application/json")

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
	if len(body) > 0 {
		spew.Dump(body)
		return false, nil
	}
	fmt.Println("Clean done")
	return true, nil
}