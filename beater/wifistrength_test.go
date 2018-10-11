package beater

import "testing"

func TestCollectWifiStrength(t *testing.T) {
	// Arrange
	var wifiStrength *WifiStrength
	var err error
	// Act
	wifiStrength, err = CollectWifiStrength()
	// Assert
	if err != nil {
		t.Errorf("Error When Collecting Wifi Strength: %v", err)
	}

	if wifiStrength.ReceivedSignalStrengthIndication > 0 || wifiStrength.ReceivedSignalStrengthIndication < -100 {
		t.Errorf("RSSI should be between [0, -100]")
	}

	if wifiStrength.Noise > 0 ||  wifiStrength.Noise < -100 {
		t.Errorf("Noise should be between [0, -100]")
	}

	if wifiStrength.TransmissionRate < 0 {
		t.Errorf("TX Rate should be positive")
	}

	if len(wifiStrength.Ssid) == 0 {
		t.Errorf("SSID can not be empty")
	}
}
