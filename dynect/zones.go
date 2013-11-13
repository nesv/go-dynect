package dynect

type ZonesResponse struct {
	ResponseBlock
	Data []string `json:"data"`
}

type ZoneResponse struct {
	ResponseBlock
	Data ZoneDataBlock `json:"data"`
}

type ZoneDataBlock struct {
	Serial      int    `json:"serial"`
	SerialStyle string `json:"serial_style"`
	Zone        string `json:"zone"`
	ZoneType    string `json:"zone_type"`
}
