package cot

import "strings"

var messages = map[string]string{
	"b-a-g":         "Alert: GeoFenceBreach",
	"b-a-o-can":     "Alert: Cancel",
	"b-a-o-opn":     "Alert: TroopsInContact",
	"b-a-o-pan":     "Alert: RingTheBell",
	"b-a-o-tbl":     "Alert: 911",
	"b-d":           "Detection",
	"b-d-a":         "Detection/Acoustic",
	"b-d-a-c":       "Detection/Acoustic/Cyclostationary",
	"b-d-a-i":       "Detection/Acoustic/Impulsive",
	"b-d-a-v":       "Detection/Acoustic/Voice",
	"b-d-c":         "Detection/CBRNE",
	"b-d-c-b":       "Detection/CBRNE/BioChem",
	"b-d-c-b-b":     "Detection/CBRNE/BioChem/Biological",
	"b-d-c-b-c":     "Detection/CBRNE/BioChem/Chemical",
	"b-d-c-e":       "Detection/CBRNE/Explosive",
	"b-d-c-e-d":     "Detection/CBRNE/Explosive/Device",
	"b-d-c-n":       "Detection/CBRNE/NuclearRadiological",
	"b-d-c-n-n":     "Detection/CBRNE/NuclearRadiological/Nuclear",
	"b-d-c-n-n-b":   "Detection/CBRNE/NuclearRadiological/Nuclear/Bomb",
	"b-d-c-n-n-sm":  "Detection/CBRNE/NuclearRadiological/Nuclear/Special Nuclear Material",
	"b-d-c-n-r":     "Detection/CBRNE/NuclearRadiological/Radiation",
	"b-d-c-n-r-dd":  "Detection/CBRNE/NuclearRadiological/Radiological/Dispersal Device (Dirty Bomb)",
	"b-d-i":         "Detection/Impact",
	"b-d-i-m":       "Detection/Impact/Mortar",
	"b-d-l":         "Detection/Launch",
	"b-d-l-b":       "Detection/Launch/Bullet",
	"b-d-l-m":       "Detection/Launch/Mortar",
	"b-d-m":         "Detection/Motion",
	"b-d-s":         "Detection/Seismic",
	"b-f-t-r":       "Photo",
	"b-i":           "Imagery",
	"b-i-v":         "Video feed",
	"b-i-x-i":       "Photo",
	"b-l":           "Alarm",
	"b-m":           "Mapping",
	"b-m-c":         "Chemlight",
	"b-m-c-b":       "Chemlight Blue",
	"b-m-c-g":       "Chemlight Green",
	"b-m-c-r":       "Chemlight Red",
	"b-m-c-y":       "Chemlight Yellow",
	"b-m-d":         "Map drawing",
	"b-m-p":         "Designated point",
	"b-m-p-a":       "Aimpoint",
	"b-m-p-c":       "Control point",
	"b-m-p-c-cp":    "Contact point",
	"b-m-p-c-ip":    "Initial point",
	"b-m-p-c-z":     "Black triangle",
	"b-m-p-i":       "Initial point",
	"b-m-p-s":       "Sensor point",
	"b-m-p-s-m":     "Point",
	"b-m-p-s-p-i":   "Point of interest",
	"b-m-p-s-p-loc": "Sensor",
	"b-m-p-s-p-op":  "Observation point",
	"b-m-p-t":       "Target point",
	"b-m-p-w":       "Checkpoint",
	"b-m-p-w-GOTO":  "Waypoint flag",
	"b-m-r":         "Route",
	"b-r":           "Report",
	"b-r-f-h-c":     "Casevac",
	"b-t-f":         "Chat message",
	"b-t-f-d":       "Chat delivery receipt",
	"b-t-f-p":       "Chat pending receipt",
	"b-t-f-r":       "Chat read receipt",
	"b-t-f-s":       "Chat delivery failure",
	"b-w":           "Weather",
	"c":             "Capability",
	"r":             "Area restrictions",
	"r-c":           "Contaminated",
	"r-c-x-b":       "Biological attack",
	"r-c-x-c":       "Chemical attack",
	"r-f":           "Flight restrictions",
	"r-o":           "Occupied",
	"r-u":           "Unsafe",
	"t":             "Tasking",
	"t-a":           "Air request",
	"t-b":           "Add Subscription",
	"t-b-c":         "Delete Subscription",
	"t-b-q":         "durable messaging control message",
	"t-k":           "Strike request",
	"t-p":           "Pairing request",
	"t-x":           "Extended request",
	"t-x-c-i-d":     "Unset incognito mode",
	"t-x-c-i-e":     "Set incognito mode",
	"t-x-c-m":       "Metrics message",
	"t-x-c-t":       "Ping",
	"t-x-c-t-r":     "Pong",
	"t-x-d-d":       "Delete by link",
	"t-x-m-c":       "Mission Change Notification",
	"t-x-m-c-l":     "Mission Log Added Notification",
	"t-x-m-n":       "Mission Creation Notification",
	"u-d":           "Drawing",
	"u-d-c-c":       "Drawing Shapes – Circle",
	"u-d-c-e":       "Drawing Shapes – Ellipse",
	"u-d-f":         "Drawing Shapes - Multiline",
	"u-d-f-m":       "Drawing Shapes – Freehand",
	"u-d-r":         "Drawing Shapes – Rectangle",
	"u-r-b-c-c":     "Range & Bearing - Circle",
	"u-rb-a":        "Range & Bearing – Line",
	"y":             "Task Reply",
	"y-a":           "Ack",
	"y-a-r":         "Rcvd",
	"y-a-w":         "Wilco",
	"y-c":           "Tasking Complete",
	"y-c-f":         "Fail",
	"y-c-s":         "Complete: Success",
	"y-s":           "Status",
}

func GetMsgType(t string) string {
	if strings.HasPrefix(t, "a-") {
		return "Atom"
	}

	if tt, ok := messages[t]; ok {
		return tt
	}

	return "Unknown"
}
