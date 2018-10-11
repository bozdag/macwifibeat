package beater

import (
	"os/exec"
	"regexp"
	"strconv"
)

type WifiStrength struct {
	ReceivedSignalStrengthIndication int
	Noise                            int
	TransmissionRate                 int
	Ssid                             string
}

func CollectWifiStrength() (*WifiStrength, error)  {

	macWifiCommand := "/System/Library/PrivateFrameworks/" +
		"Apple80211.framework/Versions/Current/Resources/airport"
	macWifiCommandArgument := "-I"
	out, err := exec.Command(macWifiCommand, macWifiCommandArgument).Output()
	if err != nil {
		return nil, err
	}
	strength := &WifiStrength{
		ReceivedSignalStrengthIndication: parseRssi(string(out)),
		Noise:                            parseNoise(string(out)),
		TransmissionRate:                 parseTxRate(string(out)),
		Ssid:                             parseSsid(string(out)),
	}

	return strength, nil
}

func parseRssi(output string) int {
	rssi := 0
	match := regexp.MustCompile("agrCtlRSSI:[[:space:]]+(-?[[:digit:]]+)").FindStringSubmatch(output)
	if match != nil {
		if i, err := strconv.Atoi(match[1]); err == nil {
			rssi = i
		}
	}
	return rssi
}

func parseNoise(output string) int {
	noise := 0
	match := regexp.MustCompile("agrCtlNoise:[[:space:]]+(-?[[:digit:]]+)").FindStringSubmatch(output)
	if match != nil {
		if i, err := strconv.Atoi(match[1]); err == nil {
			noise = i
		}
	}
	return noise
}

func parseTxRate(output string) int {
	txRate := 0
	match := regexp.MustCompile("lastTxRate:[[:space:]]+(-?[[:digit:]]+)").FindStringSubmatch(output)
	if match != nil {
		if i, err := strconv.Atoi(match[1]); err == nil {
			txRate = i
		}
	}
	return txRate
}

func parseSsid(output string) string {
	ssid := ""
	match := regexp.MustCompile("SSID:[[:space:]]+([[:graph:]]+)").FindStringSubmatch(output)
	if match != nil {
		ssid = match[1]
	}
	return ssid
}
