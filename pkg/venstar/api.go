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

type VenstarReply interface{}

type ErrorReply struct {
	Error  bool   `json:"error"`
	Reason string `json:"string"`
}

type ApiInfoReply struct {
	ErrorReply
	ApiVer int    `json:"api_ver"`
	Type   string `json:"type"`
}

type InfoReply struct {
	ErrorReply
	Name           string  `json:"name"`
	Mode           int     `json:"mode"`
	State          int     `json:"state"`
	Fan            int     `json:"fan"`
	FanState       int     `json:"fanstate"`
	TempUnits      int     `json:"tempunits"`
	Schedule       int     `json:"schedule"`
	SchedulePart   int     `json:"schedulepart"`
	Away           int     `json:"away"`
	SpaceTemp      float32 `json:"spacetemp"`
	HeatTemp       float32 `json:"heattemp"`
	CoolTemp       float32 `json:"cooltemp"`
	CoolTempMin    float32 `json:"cooltempmin"`
	CoolTempMax    float32 `json:"cooltempmax"`
	HeatTempMin    float32 `json:"heattempmin"`
	HeatTempMax    float32 `json:"heattempmax"`
	SetPointDelta  float32 `json:"setpointdelta"`
	Humidity       float32 `json:"hum"`
	AvailableModes int     `json:"availablemodes"`
}

type SensorsReply struct {
	ErrorReply
	Sensors []struct {
		Name string `json:"name"`
		Temp int    `json:"temp"`
	} `json:"sensors"`
}

type RuntimesReply struct {
	ErrorReply
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
	ErrorReply
	Alerts []struct {
		Name   string `json:"name"`
		Active bool   `json:"active"`
	}
}
