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
	Osc_source             binding.String
	Osc_destination        binding.String
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
	Sacn_source       binding.String
	Sacn_targets      binding.StringList
	Sacn_priority     binding.Int
	Sacn_universe     binding.Int
	Sacn_startChannel binding.Int
	Sacn_mode         binding.String
	Sacn_sourceName   binding.String
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

	sacnTargetsStr := configMap["sacn.targets"]
	sacnTargetsParts := strings.Split(sacnTargetsStr, " ")
	sacnTargets := binding.BindStringList(&sacnTargetsParts)

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
	config.Osc_source.Set(configMap["osc.source"])
	config.Osc_destination.Set(configMap["osc.destination"])
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
	config.Sacn_source.Set(configMap["sacn.source"])
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
	oscSource, _ := config.Osc_source.Get()
	oscDestination, _ := config.Osc_destination.Get()
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
	sacnSource, _ := config.Sacn_source.Get()
	sacnTargets, _ := config.Sacn_targets.Get()
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
		"osc.source":               oscSource,
		"osc.destination":          oscDestination,
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
		"sacn.source":        sacnSource,
		"sacn.targets":       strings.Join(sacnTargets, " "),
		"sacn.priority":      fmt.Sprintf("%d", sacnPriority),
		"sacn.universe":      fmt.Sprintf("%d", sacnUniverse),
		"sacn.start_channel": fmt.Sprintf("%d", sacnStartChannel),
		"sacn.mode":          sacnMode,
		"sacn.source_name":   sacnSourceName,
	}
}

func LoadConfigFile(filePath string, out *BoundRkbxConfig) {
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

func StoreConfigFile(config *BoundRkbxConfig, filePath string) {
	lines := "# This file is auto-generated. Manual changes will be overwritten.\n\n"

	for key, value := range convertToConfigMap(config) {
		lines += fmt.Sprintf("%s %s\n", key, value)
	}

	os.WriteFile(filePath, []byte(lines), 0064)
}

func NewBoundRkbxConfig() BoundRkbxConfig {
	return BoundRkbxConfig{
		App_licenseKey:                binding.NewString(),
		App_autoUpdate:                binding.NewBool(),
		App_debug:                     binding.NewBool(),
		Keeper_rekordboxVersion:       binding.NewString(),
		Keeper_updateRate:             binding.NewInt(),
		Keeper_slowUpdateEveryNth:     binding.NewInt(),
		Keeper_delayCompensation:      binding.NewInt(),
		Keeper_keepWarm:               binding.NewBool(),
		Keeper_decks:                  binding.NewInt(),
		Link_enabled:                  binding.NewBool(),
		Link_cumulativeErrorTolerance: binding.NewFloat(),
		Osc_enabled:                   binding.NewBool(),
		Osc_source:                    binding.NewString(),
		Osc_destination:               binding.NewString(),
		Osc_sendEveryNth:              binding.NewInt(),
		Osc_phraseOutputFormat:        binding.NewString(),
		Osc_msg_beatMaster:            binding.NewBool(),
		Osc_msg_beatMaster_div1:       binding.NewBool(),
		Osc_msg_beatMaster_div2:       binding.NewBool(),
		Osc_msg_beatMaster_div4:       binding.NewBool(),
		Osc_msg_timeMaster:            binding.NewBool(),
		Osc_msg_phraseMaster:          binding.NewBool(),
		Osc_msg_beat:                  binding.NewBool(),
		Osc_msg_beat_div1:             binding.NewBool(),
		Osc_msg_beat_div2:             binding.NewBool(),
		Osc_msg_beat_div4:             binding.NewBool(),
		Osc_msg_time:                  binding.NewBool(),
		Osc_msg_phrase:                binding.NewBool(),
		File_enabled:                  binding.NewBool(),
		File_fileName:                 binding.NewString(),
		Setlist_enabled:               binding.NewBool(),
		Setlist_seperator:             binding.NewString(),
		Setlist_filename:              binding.NewString(),
		Sacn_enabled:                  binding.NewBool(),
		Sacn_source:                   binding.NewString(),
		Sacn_targets:                  binding.NewStringList(),
		Sacn_priority:                 binding.NewInt(),
		Sacn_universe:                 binding.NewInt(),
		Sacn_startChannel:             binding.NewInt(),
		Sacn_mode:                     binding.NewString(),
		Sacn_sourceName:               binding.NewString(),
	}
}
