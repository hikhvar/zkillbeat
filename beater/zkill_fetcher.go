package beater

import (
	"gopkg.in/square/go-jose.v1/json"
	"net/http"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/hikhvar/zkillbeat/config"
	"github.com/boltdb/bolt"
)

type fetchResult struct {
	kill ZkillPackage
	err  error
}

func FetchKills(cfg config.Config, kills chan<- ZkillPackage, done chan struct{}) {
	fetChan := make(chan fetchResult)
	queueURL := cfg.ZKillboardFeed + "?queueID=" + cfg.QueueID
	db, err := bolt.Open(cfg.SystemCacheFile, 0644, nil)
	if err != nil {
		logp.Err(err.Error())
		close(done)
	}
	defer db.Close()
	sf, err := GetSystemFetcher(db)
	if err != nil {
		logp.Err(err.Error())
		close(done)
	}
	for {
		go fetch(queueURL, fetChan)
		select {
		case <-done:
			return
		case zp := <-fetChan:
			if zp.err != nil {
				logp.Err("Could not fetch kill %v", zp.err)
			} else if (zp.kill.Payload.KillID == 0 ) {
				logp.Warn("Kill ID Equals 0. There is no kill available.")
			} else {
				if err := sf.Annotate(&zp.kill.Payload.Kill.SolarSystem); err != nil {
					logp.Err("Could not annotate system.")
				}
				kills <- zp.kill
			}
		}
	}
}

func fetch(listen_url string, ret chan<- fetchResult) {
	resp, err := http.Get(listen_url)
	var fetchRet fetchResult
	defer func() { ret <- fetchRet }()
	if err != nil {
		fetchRet.err = err
		return
	}
	fetchRet.err = json.NewDecoder(resp.Body).Decode(&fetchRet.kill)
	resp.Body.Close()

}
