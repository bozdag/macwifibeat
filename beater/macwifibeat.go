package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/bozdag/macwifibeat/config"
)

// Macwifibeat configuration.
type Macwifibeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of macwifibeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Macwifibeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts macwifibeat.
func (bt *Macwifibeat) Run(b *beat.Beat) error {
	logp.Info("macwifibeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.Period)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		if wifi, err := CollectWifiStrength(); err == nil {
			event := beat.Event{
				Timestamp: time.Now(),
				Fields: common.MapStr{
					"type":    b.Info.Name,
					"rssi": wifi.ReceivedSignalStrengthIndication,
					"noise": wifi.Noise,
					"txRate": wifi.TransmissionRate,
					"ssid": wifi.Ssid,
				},
			}
			bt.client.Publish(event)
			logp.Info("Event sent")
		}
	}
}

// Stop stops macwifibeat.
func (bt *Macwifibeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
