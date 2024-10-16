package hiveos

const baseUrl = "https://api2.hiveos.farm/api/v2"

type HiveOS struct {
	farmID      string
	accessToken string
}

func New(farmId, accessToken string) *HiveOS {
	return &HiveOS{
		farmID:      farmId,
		accessToken: accessToken,
	}
}
