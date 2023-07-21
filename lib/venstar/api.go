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

const (
	ApiInfo       = ""
	QueryInfo     = "query/info"
	QuerySensors  = "query/sensors"
	QueryRunTimes = "query/runtimes"
)

type InfoBody struct {
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

func Info() {

}
