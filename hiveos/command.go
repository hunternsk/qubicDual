package hiveos

type CommandResponse struct {
	Commands []struct {
		Command string `json:"command"`
		ID      int    `json:"id"`
		Data    struct {
			Config   string `json:"config"`
			NvidiaOc string `json:"nvidia_oc"`
			Tweakers string `json:"tweakers"`
			Wallet   string `json:"wallet"`
		} `json:"data"`
	} `json:"commands"`
}
