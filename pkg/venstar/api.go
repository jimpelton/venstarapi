package venstar

// const query_endpoints = map[string]string{
// 	"api_info": "",
// 	"info":     "query/info",
// 	"sensors":  "query/sensors",
// 	"runtimes": "query/runtimes",
// 	"alerts":   "query/alerts",
// 	// "control": "control",
// 	// "settings": "settings",
// }

type ApiInfoReply struct {
	ApiVer int
	Type   string
}

type InfoReply struct {
	Name           string `json:"name"`
	Mode           int    `json:"mode"`
	State          int    `json:"state"`
	Fan            int    `json:"fan"`
	FanState       int    `json:"fanstate"`
	TempUnits      int    `json:"tempunits"`
	Schedule       int    `json:"schedule"`
	SchedulePart   int    `json:"schedulepart"`
	Away           int    `json:"away"`
	SpaceTemp      int    `json:"spacetemp"`
	HeatTemp       int    `json:"heattemp"`
	CoolTemp       int    `json:"cooltemp"`
	CoolTempMin    int    `json:"cooltempmin"`
	CoolTempMax    int    `json:"cooltempmax"`
	HeatTempMin    int    `json:"heattempmin"`
	HeatTempMax    int    `json:"heattempmax"`
	SetPointDelta  int    `json:"setpointdelta"`
	Humidity       int    `json:"hum"`
	AvailableModes int    `json:"availablemodes"`
}

type SensorsReply struct {
	Sensors []struct {
		Name string `json:"name"`
		Temp int    `json:"temp"`
	} `json:"sensors"`
}

type RuntimesReply struct {
	Runtimes []struct {
		Ts    int `json:"ts"`
		Heat1 int `json:"heat1"`
		Heat2 int `json:"heat2"`
		Cool1 int `json:"cool1"`
		Cool2 int `json:"cool2"`
		Aux1  int `json:"aux1"`
		Aux2  int `json:"aux2"`
		Fc    int `json:"fc"`
	} `json:"runtimes"`
}

type Alerts struct {
	Alerts []struct {
		Name   string `json:"name"`
		Active bool   `json:"active"`
	}
}
