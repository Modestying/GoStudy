package main

import (
	"fmt"
	"strings"
)

type Servo struct {
	data int
}

type D struct {
	num  int
	name string
}

func (S *Servo) PanLeft() {
	fmt.Println("PAN_LEFT")
}

func (S *Servo) ZoomInStop() {
	fmt.Println("ZOOM_IN_STOP")
	fmt.Println(S.data)
}

func (S *Servo) WithVal(num int, name string) {
	fmt.Println(num, "  ", name)
}
func (S *Servo) WithVal2(d D) {
	fmt.Println(d)
}
func SnakeToCamel(str string) string {
	//fmt.Println(strings.Replace(strings.Title(strings.ToLower(strings.Replace(str, "_", " ", -1))), " ", "", -1))
	return strings.Replace(strings.Title(strings.ToLower(strings.Replace(str, "_", " ", -1))), " ", "", -1)
}

var ControlCommand_name = map[int32]string{
	0:   "PAN_LEFT",
	1:   "PAN_RIGHT",
	2:   "TILT_UP",
	3:   "TILT_DOWN",
	4:   "LEFT_UP",
	5:   "LEFT_DOWN",
	6:   "RIGHT_UP",
	7:   "RIGHT_DOWN",
	8:   "PAN_TILT_STOP",
	9:   "AUTO",
	10:  "AUTO_STOP",
	11:  "POWER_REBOOT",
	12:  "SELF_CHECK",
	40:  "ZOOM_IN",
	41:  "ZOOM_IN_STOP",
	42:  "ZOOM_OUT",
	43:  "ZOOM_OUT_STOP",
	44:  "AUTO_FOCUS",
	45:  "FOCUS_FAR",
	46:  "FOCUS_FAR_STOP",
	47:  "FOCUS_NEAR",
	48:  "FOCUS_NEAR_STOP",
	49:  "IRIS_ADD",
	50:  "IRIS_ADD_STOP",
	51:  "IRIS_DEC",
	52:  "IRIS_DEC_STOP",
	53:  "WIPER_OPEN",
	54:  "WIPER_STOP",
	55:  "COLOR_BW",
	56:  "COLOR_BW_STOP",
	57:  "DEFOG_ON",
	58:  "DEFOG_OFF",
	59:  "POWER_ON",
	60:  "POWER_OFF",
	61:  "AUXILIARY_ON",
	62:  "AUXILIARY_OFF",
	63:  "HEATER_ON",
	64:  "HEATER_OFF",
	110: "ADJUST",
	111: "DIGITAL_ZOOM",
	112: "DIGITAL_ZOOM_STOP",
	113: "FOCUS_COMPENSATION",
	114: "FOCUS_COMPENSATION_STOP",
	115: "CROSS_HAIR",
	116: "CROSS_HAIR_STOP",
	117: "POLAR_BLACK",
	118: "POLAR_WHITE",
	119: "VIDEO_EFFECT",
	120: "VIDEO_EFFECT_RESET",
	121: "LEN_SELF_CHECK",
	190: "QUERY_TEMPERATURE",
	191: "TrackData",
	192: "TRACK_POS",
	200: "PRESET_SET",
	201: "PRESET_CLEAR",
	203: "PRESET_GOTO",
	205: "PRESET_CUSTOMIZED",
	206: "SET_WATCH",
	220: "CRUISE_SET",
	222: "CRUISE_CLEAR",
	223: "CRUISE_RUN",
	224: "CRUISE_STOP",
	225: "PATTERN_CRUISE_START_BOUNDARY",
	226: "PATTERN_CRUISE_STOP_BOUNDARY",
	227: "PATTERN_CRUISE_RUN",
	250: "TRACK_START",
	251: "TRACK_STOP",
}
