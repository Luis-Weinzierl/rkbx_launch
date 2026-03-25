package helpers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type RkbxConfig struct {
	// ---------- General ----------
	App_licenseKey string
	App_autoUpdate bool
	App_debug      bool

	// ---------- BeatKeeper ----------
	Keeper_rekordboxVersion   string
	Keeper_updateRate         int
	Keeper_slowUpdateEveryNth int
	Keeper_delayCompensation  int
	Keeper_keepWarm           bool
	Keeper_decks              int

	// ---------- Ableton Link ----------
	Link_enabled                  bool
	Link_cumulativeErrorTolerance float64

	// ---------- OSC ----------
	Osc_enabled            bool
	Osc_source             IPAddress
	Osc_destination        IPAddress
	Osc_sendEveryNth       int
	Osc_phraseOutputFormat string

	Osc_msg_beatMaster      bool
	Osc_msg_beatMaster_div1 bool
	Osc_msg_beatMaster_div2 bool
	Osc_msg_beatMaster_div4 bool
	Osc_msg_timeMaster      bool
	Osc_msg_phraseMaster    bool

	Osc_msg_beat      bool
	Osc_msg_beat_div1 bool
	Osc_msg_beat_div2 bool
	Osc_msg_beat_div4 bool
	Osc_msg_time      bool
	Osc_msg_phrase    bool

	// ---------- File ----------
	File_enabled  bool
	File_fileName string

	// ---------- Setlist Logging ----------
	Setlist_enabled   bool
	Setlist_seperator string
	Setlist_filename  string

	// ---------- sACN ----------
	Sacn_enabled      bool
	Sacn_source       IPAddress
	Sacn_targets      []IPAddress
	Sacn_priority     int
	Sacn_universe     int
	Sacn_startChannel int
	Sacn_mode         string
	Sacn_sourceName   string
}

type configEntry struct {
	key   string
	value string
}

func convertFromConfigMap(configMap map[string]string) RkbxConfig {
	keeperUpdateRate, err1 := strconv.Atoi(configMap["keeper.update_rate"])
	keeperSlowUpdateEveryNth, err2 := strconv.Atoi(configMap["keeper.slow_update_every_nth"])
	keeperDelayCompensation, err3 := strconv.Atoi(configMap["keeper.delay_compensation"])
	keeperDecks, err4 := strconv.Atoi(configMap["keeper.decks"])
	oscSendEveryNth, err5 := strconv.Atoi(configMap["osc.send_every_nth"])
	sacnPriority, err6 := strconv.Atoi(configMap["sacn.priority"])
	sacnUniverse, err7 := strconv.Atoi(configMap["sacn.universe"])
	sacnStartChannel, err8 := strconv.Atoi(configMap["sacn.start_channel"])

	linkCumulativeErrorTolerance, err9 := strconv.ParseFloat(configMap["link.cumulative_error_tolerance"], 32)

	oscSource, err10 := parseIPAddress(configMap["osc.source"])
	oscDestination, err11 := parseIPAddress(configMap["osc.destination"])

	sacnSource, err12 := parseIPAddress(configMap["sacn.source"])
	sacnTargetsStr := configMap["sacn.targets"]
	sacnTargetsParts := strings.Split(sacnTargetsStr, " ")
	sacnTargets := make([]IPAddress, len(sacnTargetsParts))
	for i, targetStr := range sacnTargetsParts {
		target, err := parseIPAddress(targetStr)
		if err != nil {
			panic("Error parsing config values")
		}
		sacnTargets[i] = target
	}

	if err1 != nil ||
		err2 != nil ||
		err3 != nil ||
		err4 != nil ||
		err5 != nil ||
		err6 != nil ||
		err7 != nil ||
		err8 != nil ||
		err9 != nil ||
		err10 != nil ||
		err11 != nil ||
		err12 != nil {
		panic("Error parsing config values")
	}

	return RkbxConfig{
		App_licenseKey: configMap["app.licensekey"],
		App_autoUpdate: configMap["app.auto_update"] == "true",
		App_debug:      configMap["app.debug"] == "true",

		Keeper_rekordboxVersion:   configMap["keeper.rekordbox_version"],
		Keeper_updateRate:         keeperUpdateRate,
		Keeper_slowUpdateEveryNth: keeperSlowUpdateEveryNth,
		Keeper_delayCompensation:  keeperDelayCompensation,
		Keeper_keepWarm:           configMap["keeper.keep_warm"] == "true",
		Keeper_decks:              keeperDecks,

		Link_enabled:                  configMap["link.enabled"] == "true",
		Link_cumulativeErrorTolerance: linkCumulativeErrorTolerance,

		Osc_enabled:            configMap["osc.enabled"] == "true",
		Osc_source:             oscSource,
		Osc_destination:        oscDestination,
		Osc_sendEveryNth:       oscSendEveryNth,
		Osc_phraseOutputFormat: configMap["osc.phrase_output_format"],

		Osc_msg_beatMaster:      configMap["osc.msg.beat_master"] == "true",
		Osc_msg_beatMaster_div1: configMap["osc.msg.beat_master.div_1"] == "true",
		Osc_msg_beatMaster_div2: configMap["osc.msg.beat_master.div_2"] == "true",
		Osc_msg_beatMaster_div4: configMap["osc.msg.beat_master.div_4"] == "true",
		Osc_msg_timeMaster:      configMap["osc.msg.time_master"] == "true",
		Osc_msg_phraseMaster:    configMap["osc.msg.phrase_master"] == "true",

		Osc_msg_beat:      configMap["osc.msg.beat"] == "true",
		Osc_msg_beat_div1: configMap["osc.msg.beat.div_1"] == "true",
		Osc_msg_beat_div2: configMap["osc.msg.beat.div_2"] == "true",
		Osc_msg_beat_div4: configMap["osc.msg.beat.div_4"] == "true",
		Osc_msg_time:      configMap["osc.msg.time"] == "true",
		Osc_msg_phrase:    configMap["osc.msg.phrase"] == "true",

		File_enabled:  configMap["file.enabled"] == "true",
		File_fileName: configMap["file.filename"],

		Setlist_enabled:   configMap["setlist.enabled"] == "true",
		Setlist_seperator: configMap["setlist.separator"],
		Setlist_filename:  configMap["setlist.filename"],

		Sacn_enabled:      configMap["sacn.enabled"] == "true",
		Sacn_source:       sacnSource,
		Sacn_targets:      sacnTargets,
		Sacn_priority:     sacnPriority,
		Sacn_universe:     sacnUniverse,
		Sacn_startChannel: sacnStartChannel,
		Sacn_mode:         configMap["sacn.mode"],
		Sacn_sourceName:   configMap["sacn.source_name"],
	}
}

func convertToConfigMap(config RkbxConfig) map[string]string {
	return map[string]string{
		"app.licensekey":  config.App_licenseKey,
		"app.auto_update": fmt.Sprintf("%v", config.App_autoUpdate),
		"app.debug":       fmt.Sprintf("%v", config.App_debug),

		"keeper.rekordbox_version":     config.Keeper_rekordboxVersion,
		"keeper.update_rate":           fmt.Sprintf("%d", config.Keeper_updateRate),
		"keeper.slow_update_every_nth": fmt.Sprintf("%d", config.Keeper_slowUpdateEveryNth),
		"keeper.delay_compensation":    fmt.Sprintf("%d", config.Keeper_delayCompensation),
		"keeper.keep_warm":             fmt.Sprintf("%v", config.Keeper_keepWarm),
		"keeper.decks":                 fmt.Sprintf("%d", config.Keeper_decks),

		"link.enabled":                    fmt.Sprintf("%v", config.Link_enabled),
		"link.cumulative_error_tolerance": fmt.Sprintf("%f", config.Link_cumulativeErrorTolerance),

		"osc.enabled":              fmt.Sprintf("%v", config.Osc_enabled),
		"osc.source":               config.Osc_source.String(),
		"osc.destination":          config.Osc_destination.String(),
		"osc.send_every_nth":       fmt.Sprintf("%d", config.Osc_sendEveryNth),
		"osc.phrase_output_format": config.Osc_phraseOutputFormat,

		"osc.msg.beat_master":       fmt.Sprintf("%v", config.Osc_msg_beatMaster),
		"osc.msg.beat_master.div_1": fmt.Sprintf("%v", config.Osc_msg_beatMaster_div1),
		"osc.msg.beat_master.div_2": fmt.Sprintf("%v", config.Osc_msg_beatMaster_div2),
		"osc.msg.beat_master.div_4": fmt.Sprintf("%v", config.Osc_msg_beatMaster_div4),
		"osc.msg.time_master":       fmt.Sprintf("%v", config.Osc_msg_timeMaster),
		"osc.msg.phrase_master":     fmt.Sprintf("%v", config.Osc_msg_phraseMaster),

		"osc.msg.beat":       fmt.Sprintf("%v", config.Osc_msg_beat),
		"osc.msg.beat.div_1": fmt.Sprintf("%v", config.Osc_msg_beat_div1),
		"osc.msg.beat.div_2": fmt.Sprintf("%v", config.Osc_msg_beat_div2),
		"osc.msg.beat.div_4": fmt.Sprintf("%v", config.Osc_msg_beat_div4),
		"osc.msg.time":       fmt.Sprintf("%v", config.Osc_msg_time),
		"osc.msg.phrase":     fmt.Sprintf("%v", config.Osc_msg_phrase),

		"file.enabled":  fmt.Sprintf("%v", config.File_enabled),
		"file.filename": config.File_fileName,

		"setlist.enabled":   fmt.Sprintf("%v", config.Setlist_enabled),
		"setlist.separator": config.Setlist_seperator,
		"setlist.filename":  config.Setlist_filename,

		"sacn.enabled":       fmt.Sprintf("%v", config.Sacn_enabled),
		"sacn.source":        config.Sacn_source.String(),
		"sacn.targets":       IPsToString(config.Sacn_targets),
		"sacn.priority":      fmt.Sprintf("%d", config.Sacn_priority),
		"sacn.universe":      fmt.Sprintf("%d", config.Sacn_universe),
		"sacn.start_channel": fmt.Sprintf("%d", config.Sacn_startChannel),
		"sacn.mode":          config.Sacn_mode,
		"sacn.source_name":   config.Sacn_sourceName,
	}
}

func ParseConfigFile(filePath string) RkbxConfig {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	configMap := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, " ", 2)
		if len(parts) == 2 {
			configMap[parts[0]] = strings.TrimSpace(parts[1])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return convertFromConfigMap(configMap)
}

func StoreConfigFile(config RkbxConfig, filePath string) {
	lines := "# This file is auto-generated. Manual changes will be overwritten.\n\n"

	for key, value := range convertToConfigMap(config) {
		lines += fmt.Sprintf("%s %s\n", key, value)
	}

	os.WriteFile(filePath, []byte(lines), 0064)
}

type IPAddress struct {
	Layer1 uint8
	Layer2 uint8
	Layer3 uint8
	Layer4 uint8
	Port   uint16
}

func (ip IPAddress) String() string {
	return fmt.Sprintf("%d.%d.%d.%d:%d", ip.Layer1, ip.Layer2, ip.Layer3, ip.Layer4, ip.Port)
}

func IPsToString(ips []IPAddress) string {
	strs := make([]string, len(ips))
	for i, ip := range ips {
		strs[i] = ip.String()
	}
	return strings.Join(strs, " ")
}

func parseIPAddress(ipString string) (IPAddress, error) {
	parts := strings.Split(ipString, ":")

	ipParts := strings.Split(parts[0], ".")
	if len(ipParts) != 4 {
		return IPAddress{}, fmt.Errorf("invalid IP address format")
	}

	layer1, err1 := strconv.Atoi(ipParts[0])
	layer2, err2 := strconv.Atoi(ipParts[1])
	layer3, err3 := strconv.Atoi(ipParts[2])
	layer4, err4 := strconv.Atoi(ipParts[3])
	port := 0
	var err5 error
	if len(parts) > 1 {
		port, err5 = strconv.Atoi(parts[1])
	}

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		return IPAddress{}, fmt.Errorf("invalid IP address format")
	}

	return IPAddress{
		Layer1: uint8(layer1),
		Layer2: uint8(layer2),
		Layer3: uint8(layer3),
		Layer4: uint8(layer4),
		Port:   uint16(port),
	}, nil
}
