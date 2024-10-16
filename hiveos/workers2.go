package hiveos

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var workers2Url = "/farms/%s/workers2?join_type=and&page=1&per_page=500&platform=1%%2C2&sort=name&sort_order=asc&status=online"

func (h *HiveOS) GetWorkers2() (*workers2type, error) {
	req, err := http.NewRequest("GET", baseUrl+fmt.Sprintf(workers2Url, h.farmID), nil)
	if err != nil {
		fmt.Println("req", err)
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+h.accessToken)
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

	var workers2 workers2type
	err = json.Unmarshal(body, &workers2)
	if err != nil {
		fmt.Println("unmarsh", err)
		return nil, err
	}

	return &workers2, nil
}

type workers2type struct {
	Data []struct {
		ID            int      `json:"id"`
		FarmID        int      `json:"farm_id"`
		Platform      int      `json:"platform"`
		Name          string   `json:"name"`
		Active        bool     `json:"active"`
		TagIds        []int    `json:"tag_ids"`
		Password      string   `json:"password"`
		MirrorURL     string   `json:"mirror_url"`
		IPAddresses   []string `json:"ip_addresses"`
		RemoteAddress struct {
			IP string `json:"ip"`
		} `json:"remote_address"`
		Vpn          bool   `json:"vpn"`
		SystemType   string `json:"system_type"`
		NeedsUpgrade bool   `json:"needs_upgrade"`
		LanConfig    struct {
			Dhcp    bool   `json:"dhcp"`
			Address string `json:"address"`
			Gateway string `json:"gateway"`
			DNS     string `json:"dns"`
		} `json:"lan_config"`
		Migrated        bool `json:"migrated"`
		HasMknetAutofan bool `json:"has_mknet_autofan,omitempty"`
		Versions        struct {
			Hive         string `json:"hive"`
			Kernel       string `json:"kernel"`
			AmdDriver    string `json:"amd_driver"`
			NvidiaDriver string `json:"nvidia_driver"`
		} `json:"versions"`
		Stats struct {
			Online         bool     `json:"online"`
			BootTime       int      `json:"boot_time"`
			StatsTime      int      `json:"stats_time"`
			GpusOnline     int      `json:"gpus_online"`
			GpusOffline    int      `json:"gpus_offline"`
			GpusOverheated int      `json:"gpus_overheated"`
			CpusOnline     int      `json:"cpus_online"`
			MinerStartTime int      `json:"miner_start_time"`
			PowerDraw      int      `json:"power_draw"`
			Invalid        bool     `json:"invalid"`
			LowAsr         bool     `json:"low_asr"`
			Overloaded     bool     `json:"overloaded"`
			Overheated     bool     `json:"overheated"`
			Problems       []string `json:"problems"`
		} `json:"stats,omitempty"`
		HardwareInfo struct {
			Motherboard struct {
				Manufacturer string `json:"manufacturer"`
				Model        string `json:"model"`
				Bios         string `json:"bios"`
			} `json:"motherboard"`
			CPU struct {
				ID    string `json:"id"`
				Model string `json:"model"`
				Cores int    `json:"cores"`
				Aes   bool   `json:"aes"`
			} `json:"cpu"`
			Disk struct {
				Model string `json:"model"`
			} `json:"disk"`
			NetInterfaces []struct {
				Mac   string `json:"mac"`
				Iface string `json:"iface"`
			} `json:"net_interfaces"`
		} `json:"hardware_info"`
		HardwareStats struct {
			Df      string    `json:"df"`
			Cpuavg  []float64 `json:"cpuavg"`
			Cputemp []int     `json:"cputemp"`
			Memory  struct {
				Total int `json:"total"`
				Free  int `json:"free"`
			} `json:"memory"`
			CPUCores int `json:"cpu_cores"`
		} `json:"hardware_stats"`
		Options struct {
			MaintenanceMode   int  `json:"maintenance_mode"`
			ShellinaboxEnable bool `json:"shellinabox_enable"`
			SSHEnable         bool `json:"ssh_enable"`
			SSHPasswordEnable bool `json:"ssh_password_enable"`
			VncEnable         bool `json:"vnc_enable"`
		} `json:"options,omitempty"`
		Autofan struct {
			Enabled bool `json:"enabled"`
			Items   []struct {
				Mode          string `json:"mode"`
				MaxFan        int    `json:"max_fan"`
				MinFan        int    `json:"min_fan"`
				TargetTemp    int    `json:"target_temp"`
				CriticalTemp  int    `json:"critical_temp"`
				TargetMemTemp int    `json:"target_mem_temp"`
			} `json:"items"`
			CriticalTemp   int  `json:"critical_temp"`
			NoAmd          bool `json:"no_amd"`
			RebootOnErrors bool `json:"reboot_on_errors"`
			SmartMode      bool `json:"smart_mode"`
		} `json:"autofan"`
		Commands       []interface{} `json:"commands"`
		MessagesCounts struct {
			Danger  int `json:"danger"`
			Warning int `json:"warning"`
			Success int `json:"success"`
			Info    int `json:"info"`
		} `json:"messages_counts,omitempty"`
		MessagesCountsUnresolved struct {
			Danger  int `json:"danger"`
			Warning int `json:"warning"`
			Success int `json:"success"`
			Info    int `json:"info"`
		} `json:"messages_counts_unresolved,omitempty"`
		UnitsCount  int  `json:"units_count"`
		RedTemp     int  `json:"red_temp"`
		RedFan      int  `json:"red_fan"`
		RedAsr      int  `json:"red_asr"`
		RedLa       int  `json:"red_la"`
		RedCPUTemp  int  `json:"red_cpu_temp"`
		RedMemTemp  int  `json:"red_mem_temp"`
		HasAmd      bool `json:"has_amd"`
		HasNvidia   bool `json:"has_nvidia"`
		FlightSheet struct {
			ID     int    `json:"id"`
			FarmID int    `json:"farm_id"`
			Name   string `json:"name"`
			Items  []struct {
				Coin     string `json:"coin"`
				WalID    int    `json:"wal_id"`
				Miner    string `json:"miner"`
				MinerAlt string `json:"miner_alt"`
			} `json:"items"`
		} `json:"flight_sheet,omitempty"`
		Overclock struct {
			Algo string `json:"algo"`
			/*Amd  *struct {
				CoreVddc string `json:"core_vddc"`
				MemState string `json:"mem_state"`
			} `json:"amd,omitempty"`*/
			Nvidia `json:"nvidia,omitempty"`
		} `json:"overclock,omitempty"`
		MinersSummary struct {
			Hashrates []struct {
				Miner  string  `json:"miner"`
				Ver    string  `json:"ver"`
				Algo   string  `json:"algo"`
				Coin   string  `json:"coin"`
				Hash   float64 `json:"hash"`
				Shares struct {
					Accepted int         `json:"accepted"`
					Rejected int         `json:"rejected"`
					Total    int         `json:"total"`
					Ratio    interface{} `json:"ratio"`
				} `json:"shares"`
			} `json:"hashrates"`
		} `json:"miners_summary,omitempty"`
		MinersStats struct {
			Hashrates []struct {
				Miner      string    `json:"miner"`
				Algo       string    `json:"algo"`
				Coin       string    `json:"coin"`
				Hashes     []float64 `json:"hashes"`
				Temps      []int     `json:"temps"`
				Fans       []int     `json:"fans"`
				BusNumbers []int     `json:"bus_numbers"`
			} `json:"hashrates"`
		} `json:"miners_stats,omitempty"`
		Watchdog struct {
			Enabled         bool   `json:"enabled"`
			RestartTimeout  int    `json:"restart_timeout"`
			RebootTimeout   int    `json:"reboot_timeout"`
			CheckPower      bool   `json:"check_power"`
			CheckConnection bool   `json:"check_connection"`
			PowerAction     string `json:"power_action"`
			CheckGpu        bool   `json:"check_gpu"`
			MaxLa           int    `json:"max_la"`
			MinAsr          int    `json:"min_asr"`
			Type            string `json:"type"`
			Options         struct {
				ByMiner []struct {
					Miner   string  `json:"miner"`
					Minhash float64 `json:"minhash"`
				} `json:"by_miner"`
				ByAlgo []interface{} `json:"by_algo"`
			} `json:"options"`
		} `json:"watchdog,omitempty"`
		GpuSummary struct {
			Gpus []struct {
				Name   string `json:"name"`
				Amount int    `json:"amount"`
			} `json:"gpus"`
			MaxTemp int `json:"max_temp"`
			MaxFan  int `json:"max_fan"`
		} `json:"gpu_summary"`
		GpuStats []struct {
			BusID     string  `json:"bus_id"`
			BusNumber int     `json:"bus_number"`
			BusNum    int     `json:"bus_num"`
			Temp      int     `json:"temp"`
			Fan       int     `json:"fan"`
			Power     int     `json:"power"`
			Hash      float64 `json:"hash"`
		} `json:"gpu_stats,omitempty"`
		GpuInfo []struct {
			BusID     string `json:"bus_id"`
			BusNumber int    `json:"bus_number"`
			Brand     string `json:"brand"`
			Model     string `json:"model"`
			Index     int    `json:"index,omitempty"`
			ShortName string `json:"short_name,omitempty"`
			Details   struct {
				Mem       string `json:"mem"`
				MemGb     int    `json:"mem_gb"`
				MemType   string `json:"mem_type"`
				MemOem    string `json:"mem_oem"`
				Vbios     string `json:"vbios"`
				Subvendor string `json:"subvendor"`
				Oem       string `json:"oem"`
			} `json:"details,omitempty"`
			PowerLimit struct {
				Min string `json:"min"`
				Def string `json:"def"`
				Max string `json:"max"`
			} `json:"power_limit,omitempty"`
		} `json:"gpu_info"`
		PsuEfficiency int `json:"psu_efficiency,omitempty"`
		MknetAutofan  struct {
			Fan           int  `json:"fan"`
			Auto          bool `json:"auto"`
			TargetTemp    int  `json:"target_temp"`
			TargetMemTemp int  `json:"target_mem_temp"`
			MinFan        int  `json:"min_fan"`
			MaxFan        int  `json:"max_fan"`
		} `json:"mknet_autofan,omitempty"`
		MknetAutofanInfo struct {
			Model string `json:"model"`
		} `json:"mknet_autofan_info,omitempty"`
		MknetAutofanStats struct {
			Casefan []int `json:"casefan"`
		} `json:"mknet_autofan_stats,omitempty"`
		Description      string `json:"description,omitempty"`
		PersonalSettings struct {
			IsFavorite bool `json:"is_favorite"`
		} `json:"personal_settings,omitempty"`
	} `json:"data"`
	Tags []struct {
		ID           int    `json:"id"`
		TypeID       int    `json:"type_id"`
		FarmID       int    `json:"farm_id"`
		Name         string `json:"name"`
		Color        int    `json:"color"`
		IsAuto       bool   `json:"is_auto"`
		WorkersCount int    `json:"workers_count"`
	} `json:"tags"`
	Pagination struct {
		Total       int `json:"total"`
		Count       int `json:"count"`
		PerPage     int `json:"per_page"`
		CurrentPage int `json:"current_page"`
		TotalPages  int `json:"total_pages"`
	} `json:"pagination"`
}

type Nvidia struct {
	ForceP0       bool   `json:"force_p0"`
	LogoOff       bool   `json:"logo_off"`
	CoreClock     string `json:"core_clock"`
	MemClock      string `json:"mem_clock"`
	PowerLimit    string `json:"power_limit"`
	ReducePower   bool   `json:"reduce_power"`
	LockMemClock  string `json:"lock_mem_clock"`
	LockCoreClock string `json:"lock_core_clock"`
}
