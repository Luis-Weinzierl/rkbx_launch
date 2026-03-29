package helpers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne/v2/data/binding"
)

type BoundRkbxConfig struct {
	// ---------- General ----------
	App_licenseKey binding.String
	App_autoUpdate binding.Bool
	App_debug      binding.Bool

	// ---------- BeatKeeper ----------
	Keeper_rekordboxVersion   binding.String
	Keeper_updateRate         binding.Int
	Keeper_slowUpdateEveryNth binding.Int
	Keeper_delayCompensation  binding.Int
	Keeper_keepWarm           binding.Bool
	Keeper_decks              binding.Int

	// ---------- Ableton Link ----------
	Link_enabled                  binding.Bool
	Link_cumulativeErrorTolerance binding.Float

	// ---------- OSC ----------
	Osc_enabled            binding.Bool
	Osc_source             IPAddress
	Osc_destination        IPAddress
	Osc_sendEveryNth       binding.Int
	Osc_phraseOutputFormat binding.String

	Osc_msg_beatMaster      binding.Bool
	Osc_msg_beatMaster_div1 binding.Bool
	Osc_msg_beatMaster_div2 binding.Bool
	Osc_msg_beatMaster_div4 binding.Bool
	Osc_msg_timeMaster      binding.Bool
	Osc_msg_phraseMaster    binding.Bool

	Osc_msg_beat      binding.Bool
	Osc_msg_beat_div1 binding.Bool
	Osc_msg_beat_div2 binding.Bool
	Osc_msg_beat_div4 binding.Bool
	Osc_msg_time      binding.Bool
	Osc_msg_phrase    binding.Bool

	// ---------- File ----------
	File_enabled  binding.Bool
	File_fileName binding.String

	// ---------- Setlist Logging ----------
	Setlist_enabled   binding.Bool
	Setlist_seperator binding.String
	Setlist_filename  binding.String

	// ---------- sACN ----------
	Sacn_enabled      binding.Bool
	Sacn_source       IPAddress
	Sacn_targets      []IPAddress
	Sacn_priority     binding.Int
	Sacn_universe     binding.Int
	Sacn_startChannel binding.Int
	Sacn_mode         binding.String
	Sacn_sourceName   binding.String
}

type configEntry struct {
	key   string
	value string
}

func convertFromConfigMap(configMap map[string]string) BoundRkbxConfig {
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

	applicensekey := configMap["app.licensekey"]
	appauto_update := configMap["app.auto_update"] == "true"
	appdebug := configMap["app.debug"] == "true"
	keeperrekordbox_version := configMap["keeper.rekordbox_version"]
	keeperkeep_warm := configMap["keeper.keep_warm"] == "true"
	linkenabled := configMap["link.enabled"] == "true"
	oscenabled := configMap["osc.enabled"] == "true"
	oscphrase_output_format := configMap["osc.phrase_output_format"]
	oscmsgbeat_master := configMap["osc.msg.beat_master"] == "true"
	oscmsgbeat_masterdiv_1 := configMap["osc.msg.beat_master.div_1"] == "true"
	oscmsgbeat_masterdiv_2 := configMap["osc.msg.beat_master.div_2"] == "true"
	oscmsgbeat_masterdiv_4 := configMap["osc.msg.beat_master.div_4"] == "true"
	oscmsgtime_master := configMap["osc.msg.time_master"] == "true"
	oscmsgphrase_master := configMap["osc.msg.phrase_master"] == "true"
	oscmsgbeat := configMap["osc.msg.beat"] == "true"
	oscmsgbeatdiv_1 := configMap["osc.msg.beat.div_1"] == "true"
	oscmsgbeatdiv_2 := configMap["osc.msg.beat.div_2"] == "true"
	oscmsgbeatdiv_4 := configMap["osc.msg.beat.div_4"] == "true"
	oscmsgtime := configMap["osc.msg.time"] == "true"
	oscmsgphrase := configMap["osc.msg.phrase"] == "true"
	fileenabled := configMap["file.enabled"] == "true"
	filefilename := configMap["file.filename"]
	setlistenabled := configMap["setlist.enabled"] == "true"
	setlistseparator := configMap["setlist.separator"]
	setlistfilename := configMap["setlist.filename"]
	sacnenabled := configMap["sacn.enabled"] == "true"
	sacnmode := configMap["sacn.mode"]
	sacnsource_name := configMap["sacn.source_name"]

	return BoundRkbxConfig{
		App_licenseKey: binding.BindString(&applicensekey),
		App_autoUpdate: binding.BindBool(&appauto_update),
		App_debug:      binding.BindBool(&appdebug),

		Keeper_rekordboxVersion:   binding.BindString(&keeperrekordbox_version),
		Keeper_updateRate:         binding.BindInt(&keeperUpdateRate),
		Keeper_slowUpdateEveryNth: binding.BindInt(&keeperSlowUpdateEveryNth),
		Keeper_delayCompensation:  binding.BindInt(&keeperDelayCompensation),
		Keeper_keepWarm:           binding.BindBool(&keeperkeep_warm),
		Keeper_decks:              binding.BindInt(&keeperDecks),

		Link_enabled:                  binding.BindBool(&linkenabled),
		Link_cumulativeErrorTolerance: binding.BindFloat(&linkCumulativeErrorTolerance),

		Osc_enabled:            binding.BindBool(&oscenabled),
		Osc_source:             oscSource,
		Osc_destination:        oscDestination,
		Osc_sendEveryNth:       binding.BindInt(&oscSendEveryNth),
		Osc_phraseOutputFormat: binding.BindString(&oscphrase_output_format),

		Osc_msg_beatMaster:      binding.BindBool(&oscmsgbeat_master),
		Osc_msg_beatMaster_div1: binding.BindBool(&oscmsgbeat_masterdiv_1),
		Osc_msg_beatMaster_div2: binding.BindBool(&oscmsgbeat_masterdiv_2),
		Osc_msg_beatMaster_div4: binding.BindBool(&oscmsgbeat_masterdiv_4),
		Osc_msg_timeMaster:      binding.BindBool(&oscmsgtime_master),
		Osc_msg_phraseMaster:    binding.BindBool(&oscmsgphrase_master),

		Osc_msg_beat:      binding.BindBool(&oscmsgbeat),
		Osc_msg_beat_div1: binding.BindBool(&oscmsgbeatdiv_1),
		Osc_msg_beat_div2: binding.BindBool(&oscmsgbeatdiv_2),
		Osc_msg_beat_div4: binding.BindBool(&oscmsgbeatdiv_4),
		Osc_msg_time:      binding.BindBool(&oscmsgtime),
		Osc_msg_phrase:    binding.BindBool(&oscmsgphrase),

		File_enabled:  binding.BindBool(&fileenabled),
		File_fileName: binding.BindString(&filefilename),

		Setlist_enabled:   binding.BindBool(&setlistenabled),
		Setlist_seperator: binding.BindString(&setlistseparator),
		Setlist_filename:  binding.BindString(&setlistfilename),

		Sacn_enabled:      binding.BindBool(&sacnenabled),
		Sacn_source:       sacnSource,
		Sacn_targets:      sacnTargets,
		Sacn_priority:     binding.BindInt(&sacnPriority),
		Sacn_universe:     binding.BindInt(&sacnUniverse),
		Sacn_startChannel: binding.BindInt(&sacnStartChannel),
		Sacn_mode:         binding.BindString(&sacnmode),
		Sacn_sourceName:   binding.BindString(&sacnsource_name),
	}
}

func fillFromConfigMap(config *BoundRkbxConfig, configMap map[string]string) {
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
	config.App_licenseKey.Set(configMap["app.licensekey"])
	config.App_autoUpdate.Set(configMap["app.auto_update"] == "true")
	config.App_debug.Set(configMap["app.debug"] == "true")

	config.Keeper_rekordboxVersion.Set(configMap["keeper.rekordbox_version"])
	config.Keeper_updateRate.Set(keeperUpdateRate)
	config.Keeper_slowUpdateEveryNth.Set(keeperSlowUpdateEveryNth)
	config.Keeper_delayCompensation.Set(keeperDelayCompensation)
	config.Keeper_keepWarm.Set(configMap["keeper.keep_warm"] == "true")
	config.Keeper_decks.Set(keeperDecks)

	config.Link_enabled.Set(configMap["link.enabled"] == "true")
	config.Link_cumulativeErrorTolerance.Set(linkCumulativeErrorTolerance)

	config.Osc_enabled.Set(configMap["osc.enabled"] == "true")
	config.Osc_source = oscSource
	config.Osc_destination = oscDestination
	config.Osc_sendEveryNth.Set(oscSendEveryNth)
	config.Osc_phraseOutputFormat.Set(configMap["osc.phrase_output_format"])

	config.Osc_msg_beatMaster.Set(configMap["osc.msg.beat_master"] == "true")
	config.Osc_msg_beatMaster_div1.Set(configMap["osc.msg.beat_master.div_1"] == "true")
	config.Osc_msg_beatMaster_div2.Set(configMap["osc.msg.beat_master.div_2"] == "true")
	config.Osc_msg_beatMaster_div4.Set(configMap["osc.msg.beat_master.div_4"] == "true")
	config.Osc_msg_timeMaster.Set(configMap["osc.msg.time_master"] == "true")
	config.Osc_msg_phraseMaster.Set(configMap["osc.msg.phrase_master"] == "true")

	config.Osc_msg_beat.Set(configMap["osc.msg.beat"] == "true")
	config.Osc_msg_beat_div1.Set(configMap["osc.msg.beat.div_1"] == "true")
	config.Osc_msg_beat_div2.Set(configMap["osc.msg.beat.div_2"] == "true")
	config.Osc_msg_beat_div4.Set(configMap["osc.msg.beat.div_4"] == "true")
	config.Osc_msg_time.Set(configMap["osc.msg.time"] == "true")
	config.Osc_msg_phrase.Set(configMap["osc.msg.phrase"] == "true")

	config.File_enabled.Set(configMap["file.enabled"] == "true")
	config.File_fileName.Set(configMap["file.filename"])

	config.Setlist_enabled.Set(configMap["setlist.enabled"] == "true")
	config.Setlist_seperator.Set(configMap["setlist.separator"])
	config.Setlist_filename.Set(configMap["setlist.filename"])

	config.Sacn_enabled.Set(configMap["sacn.enabled"] == "true")
	config.Sacn_source = sacnSource
	config.Sacn_targets = sacnTargets
	config.Sacn_priority.Set(sacnPriority)
	config.Sacn_universe.Set(sacnUniverse)
	config.Sacn_startChannel.Set(sacnStartChannel)
	config.Sacn_mode.Set(configMap["sacn.mode"])
	config.Sacn_sourceName.Set(configMap["sacn.source_name"])
}

func convertToConfigMap(config *BoundRkbxConfig) map[string]string {
	appLicenseKey, _ := config.App_licenseKey.Get()
	appAutoUpdate, _ := config.App_autoUpdate.Get()
	appDebug, _ := config.App_debug.Get()

	keeperRekordboxVersion, _ := config.Keeper_rekordboxVersion.Get()
	keeperUpdateRate, _ := config.Keeper_updateRate.Get()
	keeperSlowUpdateEveryNth, _ := config.Keeper_slowUpdateEveryNth.Get()
	keeperDelayCompensation, _ := config.Keeper_delayCompensation.Get()
	keeperKeepWarm, _ := config.Keeper_keepWarm.Get()
	keeperDecks, _ := config.Keeper_decks.Get()

	linkEnabled, _ := config.Link_enabled.Get()
	linkCumulativeErrorTolerance, _ := config.Link_cumulativeErrorTolerance.Get()

	oscEnabled, _ := config.Osc_enabled.Get()
	oscSendEveryNth, _ := config.Osc_sendEveryNth.Get()
	oscPhraseOutputFormat, _ := config.Osc_phraseOutputFormat.Get()

	oscMsgBeatMaster, _ := config.Osc_msg_beatMaster.Get()
	oscMsgBeatMasterDiv1, _ := config.Osc_msg_beatMaster_div1.Get()
	oscMsgBeatMasterDiv2, _ := config.Osc_msg_beatMaster_div2.Get()
	oscMsgBeatMasterDiv4, _ := config.Osc_msg_beatMaster_div4.Get()
	oscMsgTimeMaster, _ := config.Osc_msg_timeMaster.Get()
	oscMsgPhraseMaster, _ := config.Osc_msg_phraseMaster.Get()

	oscMsgBeat, _ := config.Osc_msg_beat.Get()
	oscMsgBeatDiv1, _ := config.Osc_msg_beat_div1.Get()
	oscMsgBeatDiv2, _ := config.Osc_msg_beat_div2.Get()
	oscMsgBeatDiv4, _ := config.Osc_msg_beat_div4.Get()
	oscMsgTime, _ := config.Osc_msg_time.Get()
	oscMsgPhrase, _ := config.Osc_msg_phrase.Get()

	fileEnabled, _ := config.File_enabled.Get()
	fileFileName, _ := config.File_fileName.Get()

	setlistEnabled, _ := config.Setlist_enabled.Get()
	setlistSeparator, _ := config.Setlist_seperator.Get()
	setlistFilename, _ := config.Setlist_filename.Get()

	sacnEnabled, _ := config.Sacn_enabled.Get()
	sacnPriority, _ := config.Sacn_priority.Get()
	sacnUniverse, _ := config.Sacn_universe.Get()
	sacnStartChannel, _ := config.Sacn_startChannel.Get()
	sacnMode, _ := config.Sacn_mode.Get()
	sacnSourceName, _ := config.Sacn_sourceName.Get()

	return map[string]string{
		"app.licensekey":  appLicenseKey,
		"app.auto_update": fmt.Sprintf("%v", appAutoUpdate),
		"app.debug":       fmt.Sprintf("%v", appDebug),

		"keeper.rekordbox_version":     keeperRekordboxVersion,
		"keeper.update_rate":           fmt.Sprintf("%d", keeperUpdateRate),
		"keeper.slow_update_every_nth": fmt.Sprintf("%d", keeperSlowUpdateEveryNth),
		"keeper.delay_compensation":    fmt.Sprintf("%d", keeperDelayCompensation),
		"keeper.keep_warm":             fmt.Sprintf("%v", keeperKeepWarm),
		"keeper.decks":                 fmt.Sprintf("%d", keeperDecks),

		"link.enabled":                    fmt.Sprintf("%v", linkEnabled),
		"link.cumulative_error_tolerance": fmt.Sprintf("%f", linkCumulativeErrorTolerance),

		"osc.enabled":              fmt.Sprintf("%v", oscEnabled),
		"osc.source":               config.Osc_source.String(),
		"osc.destination":          config.Osc_destination.String(),
		"osc.send_every_nth":       fmt.Sprintf("%d", oscSendEveryNth),
		"osc.phrase_output_format": oscPhraseOutputFormat,

		"osc.msg.beat_master":       fmt.Sprintf("%v", oscMsgBeatMaster),
		"osc.msg.beat_master.div_1": fmt.Sprintf("%v", oscMsgBeatMasterDiv1),
		"osc.msg.beat_master.div_2": fmt.Sprintf("%v", oscMsgBeatMasterDiv2),
		"osc.msg.beat_master.div_4": fmt.Sprintf("%v", oscMsgBeatMasterDiv4),
		"osc.msg.time_master":       fmt.Sprintf("%v", oscMsgTimeMaster),
		"osc.msg.phrase_master":     fmt.Sprintf("%v", oscMsgPhraseMaster),

		"osc.msg.beat":       fmt.Sprintf("%v", oscMsgBeat),
		"osc.msg.beat.div_1": fmt.Sprintf("%v", oscMsgBeatDiv1),
		"osc.msg.beat.div_2": fmt.Sprintf("%v", oscMsgBeatDiv2),
		"osc.msg.beat.div_4": fmt.Sprintf("%v", oscMsgBeatDiv4),
		"osc.msg.time":       fmt.Sprintf("%v", oscMsgTime),
		"osc.msg.phrase":     fmt.Sprintf("%v", oscMsgPhrase),

		"file.enabled":  fmt.Sprintf("%v", fileEnabled),
		"file.filename": fileFileName,

		"setlist.enabled":   fmt.Sprintf("%v", setlistEnabled),
		"setlist.separator": setlistSeparator,
		"setlist.filename":  setlistFilename,

		"sacn.enabled":       fmt.Sprintf("%v", sacnEnabled),
		"sacn.source":        config.Sacn_source.String(),
		"sacn.targets":       IPsToString(config.Sacn_targets),
		"sacn.priority":      fmt.Sprintf("%d", sacnPriority),
		"sacn.universe":      fmt.Sprintf("%d", sacnUniverse),
		"sacn.start_channel": fmt.Sprintf("%d", sacnStartChannel),
		"sacn.mode":          sacnMode,
		"sacn.source_name":   sacnSourceName,
	}
}

func ParseConfigFile(filePath string, out *BoundRkbxConfig) {
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

	fillFromConfigMap(out, configMap)
}

func LoadConfigFile(filePath string) BoundRkbxConfig {
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

func StoreConfigFile(config *BoundRkbxConfig, filePath string) {
	lines := "# This file is auto-generated. Manual changes will be overwritten.\n\n"

	for key, value := range convertToConfigMap(config) {
		lines += fmt.Sprintf("%s %s\n", key, value)
	}

	os.WriteFile(filePath, []byte(lines), 0064)
}

type IPAddress struct {
	Layer1 int
	Layer2 int
	Layer3 int
	Layer4 int
	Port   int
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
		Layer1: layer1,
		Layer2: layer2,
		Layer3: layer3,
		Layer4: layer4,
		Port:   port,
	}, nil
}
