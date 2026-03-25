package main

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
	app_licenseKey string
	app_autoUpdate bool
	app_debug      bool

	// ---------- BeatKeeper ----------
	keeper_rekordboxVersion   string
	keeper_updateRate         int
	keeper_slowUpdateEveryNth int
	keeper_delayCompensation  int
	keeper_keepWarm           bool
	keeper_decks              int

	// ---------- Ableton Link ----------
	link_enabled                  bool
	link_cumulativeErrorTolerance float64

	// ---------- OSC ----------
	osc_enabled            bool
	osc_source             string
	osc_destination        string
	osc_sendEveryNth       int
	osc_phraseOutputFormat string

	osc_msg_beatMaster      bool
	osc_msg_beatMaster_div1 bool
	osc_msg_beatMaster_div2 bool
	osc_msg_beatMaster_div4 bool
	osc_msg_timeMaster      bool
	osc_msg_phraseMaster    bool

	osc_msg_beat      bool
	osc_msg_beat_div1 bool
	osc_msg_beat_div2 bool
	osc_msg_beat_div4 bool
	osc_msg_time      bool
	osc_msg_phrase    bool

	// ---------- File ----------
	file_enabled  bool
	file_fileName string

	// ---------- Setlist Logging ----------
	setlist_enabled   bool
	setlist_seperator string
	setlist_filename  string

	// ---------- sACN ----------
	sacn_enabled      bool
	sacn_source       string
	sacn_targets      []string
	sacn_priority     int
	sacn_universe     int
	sacn_startChannel int
	sacn_mode         string
	sacn_sourceName   string
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

	if err1 != nil ||
		err2 != nil ||
		err3 != nil ||
		err4 != nil ||
		err5 != nil ||
		err6 != nil ||
		err7 != nil ||
		err8 != nil ||
		err9 != nil {
		panic("Error parsing config values")
	}

	return RkbxConfig{
		app_licenseKey: configMap["app.licensekey"],
		app_autoUpdate: configMap["app.auto_update"] == "true",
		app_debug:      configMap["app.debug"] == "true",

		keeper_rekordboxVersion:   configMap["keeper.rekordbox_version"],
		keeper_updateRate:         keeperUpdateRate,
		keeper_slowUpdateEveryNth: keeperSlowUpdateEveryNth,
		keeper_delayCompensation:  keeperDelayCompensation,
		keeper_keepWarm:           configMap["keeper.keep_warm"] == "true",
		keeper_decks:              keeperDecks,

		link_enabled:                  configMap["link.enabled"] == "true",
		link_cumulativeErrorTolerance: linkCumulativeErrorTolerance,

		osc_enabled:            configMap["osc.enabled"] == "true",
		osc_source:             configMap["osc.source"],
		osc_destination:        configMap["osc.destination"],
		osc_sendEveryNth:       oscSendEveryNth,
		osc_phraseOutputFormat: configMap["osc.phrase_output_format"],

		osc_msg_beatMaster:      configMap["osc.msg.beat_master"] == "true",
		osc_msg_beatMaster_div1: configMap["osc.msg.beat_master.div_1"] == "true",
		osc_msg_beatMaster_div2: configMap["osc.msg.beat_master.div_2"] == "true",
		osc_msg_beatMaster_div4: configMap["osc.msg.beat_master.div_4"] == "true",
		osc_msg_timeMaster:      configMap["osc.msg.time_master"] == "true",
		osc_msg_phraseMaster:    configMap["osc.msg.phrase_master"] == "true",

		osc_msg_beat:      configMap["osc.msg.beat"] == "true",
		osc_msg_beat_div1: configMap["osc.msg.beat.div_1"] == "true",
		osc_msg_beat_div2: configMap["osc.msg.beat.div_2"] == "true",
		osc_msg_beat_div4: configMap["osc.msg.beat.div_4"] == "true",
		osc_msg_time:      configMap["osc.msg.time"] == "true",
		osc_msg_phrase:    configMap["osc.msg.phrase"] == "true",

		file_enabled:  configMap["file.enabled"] == "true",
		file_fileName: configMap["file.filename"],

		setlist_enabled:   configMap["setlist.enabled"] == "true",
		setlist_seperator: configMap["setlist.separator"],
		setlist_filename:  configMap["setlist.filename"],

		sacn_enabled:      configMap["sacn.enabled"] == "true",
		sacn_source:       configMap["sacn.source"],
		sacn_targets:      strings.Split(configMap["sacn.targets"], " "),
		sacn_priority:     sacnPriority,
		sacn_universe:     sacnUniverse,
		sacn_startChannel: sacnStartChannel,
		sacn_mode:         configMap["sacn.mode"],
		sacn_sourceName:   configMap["sacn.source_name"],
	}
}

func convertToConfigMap(config RkbxConfig) map[string]string {
	return map[string]string{
		"app.licensekey":  config.app_licenseKey,
		"app.auto_update": fmt.Sprintf("%v", config.app_autoUpdate),
		"app.debug":       fmt.Sprintf("%v", config.app_debug),

		"keeper.rekordbox_version":     config.keeper_rekordboxVersion,
		"keeper.update_rate":           fmt.Sprintf("%d", config.keeper_updateRate),
		"keeper.slow_update_every_nth": fmt.Sprintf("%d", config.keeper_slowUpdateEveryNth),
		"keeper.delay_compensation":    fmt.Sprintf("%d", config.keeper_delayCompensation),
		"keeper.keep_warm":             fmt.Sprintf("%v", config.keeper_keepWarm),
		"keeper.decks":                 fmt.Sprintf("%d", config.keeper_decks),

		"link.enabled":                    fmt.Sprintf("%v", config.link_enabled),
		"link.cumulative_error_tolerance": fmt.Sprintf("%f", config.link_cumulativeErrorTolerance),

		"osc.enabled":              fmt.Sprintf("%v", config.osc_enabled),
		"osc.source":               config.osc_source,
		"osc.destination":          config.osc_destination,
		"osc.send_every_nth":       fmt.Sprintf("%d", config.osc_sendEveryNth),
		"osc.phrase_output_format": config.osc_phraseOutputFormat,

		"osc.msg.beat_master":       fmt.Sprintf("%v", config.osc_msg_beatMaster),
		"osc.msg.beat_master.div_1": fmt.Sprintf("%v", config.osc_msg_beatMaster_div1),
		"osc.msg.beat_master.div_2": fmt.Sprintf("%v", config.osc_msg_beatMaster_div2),
		"osc.msg.beat_master.div_4": fmt.Sprintf("%v", config.osc_msg_beatMaster_div4),
		"osc.msg.time_master":       fmt.Sprintf("%v", config.osc_msg_timeMaster),
		"osc.msg.phrase_master":     fmt.Sprintf("%v", config.osc_msg_phraseMaster),

		"osc.msg.beat":       fmt.Sprintf("%v", config.osc_msg_beat),
		"osc.msg.beat.div_1": fmt.Sprintf("%v", config.osc_msg_beat_div1),
		"osc.msg.beat.div_2": fmt.Sprintf("%v", config.osc_msg_beat_div2),
		"osc.msg.beat.div_4": fmt.Sprintf("%v", config.osc_msg_beat_div4),
		"osc.msg.time":       fmt.Sprintf("%v", config.osc_msg_time),
		"osc.msg.phrase":     fmt.Sprintf("%v", config.osc_msg_phrase),

		"file.enabled":  fmt.Sprintf("%v", config.file_enabled),
		"file.filename": config.file_fileName,

		"setlist.enabled":   fmt.Sprintf("%v", config.setlist_enabled),
		"setlist.separator": config.setlist_seperator,
		"setlist.filename":  config.setlist_filename,

		"sacn.enabled":       fmt.Sprintf("%v", config.sacn_enabled),
		"sacn.source":        config.sacn_source,
		"sacn.targets":       strings.Join(config.sacn_targets, " "),
		"sacn.priority":      fmt.Sprintf("%d", config.sacn_priority),
		"sacn.universe":      fmt.Sprintf("%d", config.sacn_universe),
		"sacn.start_channel": fmt.Sprintf("%d", config.sacn_startChannel),
		"sacn.mode":          config.sacn_mode,
		"sacn.source_name":   config.sacn_sourceName,
	}
}

func parseConfigFile(filePath string) RkbxConfig {
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

func storeConfigFile(config RkbxConfig, filePath string) {
	lines := "# This file is auto-generated. Manual changes will be overwritten.\n\n"

	for key, value := range convertToConfigMap(config) {
		lines += fmt.Sprintf("%s %s\n", key, value)
	}

	os.WriteFile(filePath, []byte(lines), 0064)
}
