package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

	"github.com/hikhvar/zkillbeat/config"
)

const killTimeFormat string = "2006.01.02 15:04:05"

type Zkillbeat struct {
	done   chan struct{}
	config config.Config
	client publisher.Client
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Zkillbeat{
		done:   make(chan struct{}),
		config: config,
	}
	return bt, nil
}

func (bt *Zkillbeat) Run(b *beat.Beat) error {
	logp.Info("zkillbeat is running! Hit CTRL-C to stop it.")

	bt.client = b.Publisher.Connect()
	//ticker := time.NewTicker(bt.config.Period)
	kills := make(chan ZkillPackage)
	go FetchKills(bt.config, kills, bt.done)
	for {
		select {
		case <-bt.done:
			return nil
		case kill := <-kills:
			killTime, err := time.ParseInLocation(killTimeFormat, kill.Payload.Kill.KillTime, time.UTC)
			if err != nil {
				logp.Err("Could not parse Time &v", err)
			}
			event := common.MapStr{
				"@timestamp":    common.Time(killTime),
				"fetched":       common.Time(time.Now()),
				"type":          b.Name,
				"killid":        kill.Payload.KillID,
				"killmail":      kill.Payload.Kill,
				"zkillmetadata": kill.Payload.Metadata,
			}
			bt.client.PublishEvent(event)
			logp.Info("Event sent")
		}

	}
}

func (bt *Zkillbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
