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

type RkbxLinkConfig struct {
	// ---------- General ----------
	App_licenseKey binding.String

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
	Osc_enabled             binding.Bool
	Osc_source              binding.String
	Osc_destination         binding.String
	Osc_sendEveryNth        binding.Int
	Osc_phraseOutputFormat  binding.String
	Osc_trigger_autorelease binding.Bool

	Osc_msg_masterTime   binding.Bool
	Osc_msg_masterPhrase binding.Bool

	Osc_msg_nTime   binding.Bool
	Osc_msg_nPhrase binding.Bool

	Osc_msg_masterBeatSubdiv         binding.Float
	Osc_msg_masterBeatSubdivEnabled  binding.Bool
	Osc_msg_masterBeatTrigger        binding.Float
	Osc_msg_masterBeatTriggerEnabled binding.Bool

	Osc_msg_nBeatSubdiv         binding.Float
	Osc_msg_nBeatSubdivEnabled  binding.Bool
	Osc_msg_nBeatTrigger        binding.Float
	Osc_msg_nBeatTriggerEnabled binding.Bool

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

	// ---------- Internal ----------
	AvailableRekordboxVersions binding.StringList
	HasUnsavedChanges          binding.Bool
}

func (config *RkbxLinkConfig) IsEvaluation() bool {
	if val, err := config.App_licenseKey.Get(); err == nil {
		return val == "evaluation" || val == ""
	}
	return false
}

func fillFromConfigMap(config *RkbxLinkConfig, configMap map[string]string) {
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
	sacnTargets := strings.Split(sacnTargetsStr, " ")

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
	config.Osc_trigger_autorelease.Set(configMap["osc.trigger_autorelease"] == "true")
	config.Osc_msg_masterTime.Set(configMap["osc.msg.master/time"] == "true")
	config.Osc_msg_masterPhrase.Set(configMap["osc.msg.master/phrase"] == "true")
	config.Osc_msg_nTime.Set(configMap["osc.msg.n/time"] == "true")
	config.Osc_msg_nPhrase.Set(configMap["osc.msg.n/phrase"] == "true")

	OptionalStringToBindings(config.Osc_msg_masterBeatSubdiv, config.Osc_msg_masterBeatSubdivEnabled, configMap["osc.msg.master/beat/subdiv"])
	OptionalStringToBindings(config.Osc_msg_masterBeatTrigger, config.Osc_msg_masterBeatTriggerEnabled, configMap["osc.msg.master/beat/trigger"])
	OptionalStringToBindings(config.Osc_msg_nBeatSubdiv, config.Osc_msg_nBeatSubdivEnabled, configMap["osc.msg.n/beat/subdiv"])
	OptionalStringToBindings(config.Osc_msg_nBeatTrigger, config.Osc_msg_nBeatTriggerEnabled, configMap["osc.msg.n/beat/trigger"])

	config.File_enabled.Set(configMap["file.enabled"] == "true")
	config.File_fileName.Set(configMap["file.filename"])

	config.Setlist_enabled.Set(configMap["setlist.enabled"] == "true")
	config.Setlist_seperator.Set(configMap["setlist.separator"])
	config.Setlist_filename.Set(configMap["setlist.filename"])

	config.Sacn_enabled.Set(configMap["sacn.enabled"] == "true")
	config.Sacn_source.Set(configMap["sacn.source"])
	config.Sacn_targets.Set(sacnTargets)
	config.Sacn_priority.Set(sacnPriority)
	config.Sacn_universe.Set(sacnUniverse)
	config.Sacn_startChannel.Set(sacnStartChannel)
	config.Sacn_mode.Set(configMap["sacn.mode"])
	config.Sacn_sourceName.Set(configMap["sacn.source_name"])

	config.updateAvaliableVersions()
	config.HasUnsavedChanges.Set(false)
}

func convertToConfigMap(config *RkbxLinkConfig) map[string]string {
	appLicenseKey, _ := config.App_licenseKey.Get()

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
	oscTriggerAutoRelease, _ := config.Osc_trigger_autorelease.Get()
	oscMsgMasterTime, _ := config.Osc_msg_masterTime.Get()
	oscMsgMasterPhrase, _ := config.Osc_msg_masterPhrase.Get()
	oscMsgNTime, _ := config.Osc_msg_nTime.Get()
	oscMsgNPhrase, _ := config.Osc_msg_nPhrase.Get()

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
		"app.auto_update": "true",
		"app.debug":       "true",
		"app.yesToAll":    "true",

		"display.enabled":  "true",
		"display.interval": "1.0",

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
		"osc.trigger_autorelease":  fmt.Sprintf("%v", oscTriggerAutoRelease),

		"osc.msg.master/time":   fmt.Sprintf("%v", oscMsgMasterTime),
		"osc.msg.master/phrase": fmt.Sprintf("%v", oscMsgMasterPhrase),

		"osc.msg.n/time":   fmt.Sprintf("%v", oscMsgNTime),
		"osc.msg.n/phrase": fmt.Sprintf("%v", oscMsgNPhrase),

		"osc.msg.master/beat/subdiv":  OptionalToString(config.Osc_msg_masterBeatSubdiv, config.Osc_msg_masterBeatSubdivEnabled),
		"osc.msg.master/beat/trigger": OptionalToString(config.Osc_msg_masterBeatTrigger, config.Osc_msg_masterBeatTriggerEnabled),

		"osc.msg.n/beat/subdiv":  OptionalToString(config.Osc_msg_nBeatSubdiv, config.Osc_msg_nBeatSubdivEnabled),
		"osc.msg.n/beat/trigger": OptionalToString(config.Osc_msg_nBeatTrigger, config.Osc_msg_nBeatTriggerEnabled),

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

func LoadConfigFile(filePath string, out *RkbxLinkConfig) {
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
		} else if len(parts) == 1 {
			configMap[parts[0]] = ""
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fillFromConfigMap(out, configMap)
}

func StoreConfigFile(config *RkbxLinkConfig, filePath string) {
	lines := "# This file is auto-generated. Manual changes will be overwritten.\n\n"

	for key, value := range convertToConfigMap(config) {
		lines += fmt.Sprintf("%s %s\n", key, value)
	}

	os.WriteFile(filePath, []byte(lines), 0064)
	config.HasUnsavedChanges.Set(false)
}

func NewBoundRkbxConfig() RkbxLinkConfig {
	config := RkbxLinkConfig{}

	config.AvailableRekordboxVersions = binding.NewStringList()
	config.HasUnsavedChanges = binding.NewBool()

	listener := binding.NewDataListener(func() {
		config.HasUnsavedChanges.Set(true)
	})

	config.App_licenseKey = binding.NewString() // Doesn't change during normal operation so we dont need to track it in the normal routine
	config.Keeper_rekordboxVersion = newStringBindingWithListener(listener)
	config.Keeper_updateRate = newIntBindingWithListener(listener)
	config.Keeper_slowUpdateEveryNth = newIntBindingWithListener(listener)
	config.Keeper_delayCompensation = newIntBindingWithListener(listener)
	config.Keeper_keepWarm = newBoolBindingWithListener(listener)
	config.Keeper_decks = newIntBindingWithListener(listener)
	config.Link_enabled = newBoolBindingWithListener(listener)
	config.Link_cumulativeErrorTolerance = newFloatBindingWithListener(listener)
	config.Osc_enabled = newBoolBindingWithListener(listener)
	config.Osc_source = newStringBindingWithListener(listener)
	config.Osc_destination = newStringBindingWithListener(listener)
	config.Osc_sendEveryNth = newIntBindingWithListener(listener)
	config.Osc_phraseOutputFormat = newStringBindingWithListener(listener)
	config.Osc_trigger_autorelease = newBoolBindingWithListener(listener)
	config.Osc_msg_masterTime = newBoolBindingWithListener(listener)
	config.Osc_msg_masterPhrase = newBoolBindingWithListener(listener)
	config.Osc_msg_nTime = newBoolBindingWithListener(listener)
	config.Osc_msg_nPhrase = newBoolBindingWithListener(listener)
	config.Osc_msg_masterBeatSubdiv = newFloatBindingWithListener(listener)
	config.Osc_msg_masterBeatSubdivEnabled = newBoolBindingWithListener(listener)
	config.Osc_msg_masterBeatTrigger = newFloatBindingWithListener(listener)
	config.Osc_msg_masterBeatTriggerEnabled = newBoolBindingWithListener(listener)
	config.Osc_msg_nBeatSubdiv = newFloatBindingWithListener(listener)
	config.Osc_msg_nBeatSubdivEnabled = newBoolBindingWithListener(listener)
	config.Osc_msg_nBeatTrigger = newFloatBindingWithListener(listener)
	config.Osc_msg_nBeatTriggerEnabled = newBoolBindingWithListener(listener)
	config.File_enabled = newBoolBindingWithListener(listener)
	config.File_fileName = newStringBindingWithListener(listener)
	config.Setlist_enabled = newBoolBindingWithListener(listener)
	config.Setlist_seperator = newStringBindingWithListener(listener)
	config.Setlist_filename = newStringBindingWithListener(listener)
	config.Sacn_enabled = newBoolBindingWithListener(listener)
	config.Sacn_source = newStringBindingWithListener(listener)
	config.Sacn_targets = newStringListBindingWithListener(listener)
	config.Sacn_priority = newIntBindingWithListener(listener)
	config.Sacn_universe = newIntBindingWithListener(listener)
	config.Sacn_startChannel = newIntBindingWithListener(listener)
	config.Sacn_mode = newStringBindingWithListener(listener)
	config.Sacn_sourceName = newStringBindingWithListener(listener)

	config.App_licenseKey.AddListener(binding.NewDataListener(config.updateAvaliableVersions))

	return config
}

func (config RkbxLinkConfig) updateAvaliableVersions() {
	availVersions := []string{"7.2.10", "7.2.8", "7.2.6", "7.2.4", "7.2.3", "7.2.2", "7.1.4"}

	if config.IsEvaluation() {
		availVersions = []string{"7.2.2"}
	}

	config.AvailableRekordboxVersions.Set(availVersions)
}

func newStringBindingWithListener(listener binding.DataListener) binding.String {
	bind := binding.NewString()
	bind.AddListener(listener)
	return bind
}

func newBoolBindingWithListener(listener binding.DataListener) binding.Bool {
	bind := binding.NewBool()
	bind.AddListener(listener)
	return bind
}

func newIntBindingWithListener(listener binding.DataListener) binding.Int {
	bind := binding.NewInt()
	bind.AddListener(listener)
	return bind
}

func newFloatBindingWithListener(listener binding.DataListener) binding.Float {
	bind := binding.NewFloat()
	bind.AddListener(listener)
	return bind
}

func newStringListBindingWithListener(listener binding.DataListener) binding.StringList {
	bind := binding.NewStringList()
	bind.AddListener(listener)
	return bind
}

func OptionalToString(value binding.Float, enabled binding.Bool) string {
	if val, err := value.Get(); err == nil {
		if enabledVal, err2 := enabled.Get(); err2 == nil && enabledVal {
			return fmt.Sprintf("%v", val)
		}
	}
	return ""
}

func OptionalStringToBindings(valueBind binding.Float, enabledBind binding.Bool, valueStr string) {
	if valueStr == "" {
		valueBind.Set(1.0)
		enabledBind.Set(false)
	} else {
		if val, err := strconv.ParseFloat(valueStr, 32); err == nil {
			valueBind.Set(val)
			enabledBind.Set(true)
		} else {
			valueBind.Set(1.0)
			enabledBind.Set(false)
		}
	}
}
