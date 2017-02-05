package beater

import (
	"github.com/boltdb/bolt"
	"strconv"
	"bytes"
	"net/http"
	"encoding/json"
)

type CrestSystem struct {
	EveAPIItem
	Constellation EveAPIItem `json:"constellation"`
	SecurityStatus float64 `json:"securityStatus"`
}

type CrestConstellation struct {
	Name string `json:"name"`
	Region struct{Source string `json:"href"`}
}

type CrestRegion struct {
	Name string `json:"name"`
}

type systemAnnotationError error

func IsAnnotationError(err error) bool {
	_, ok := err.(systemAnnotationError)
	return ok
}

type SystemFetcher struct {
	bucket *bolt.Bucket
}

func GetSystemFetcher(db *bolt.DB) (SystemFetcher, error){
	tx, err := db.Begin(true)
	if err != nil {
		return SystemFetcher{}, err
	}
	b, err := tx.CreateBucketIfNotExists([]byte("systems"))
	if err != nil {
		return SystemFetcher{}, err
	}
	return SystemFetcher{bucket:b}, nil
}

func (sf *SystemFetcher) Annotate(system *SolarSystem) error {
	key := []byte(strconv.Itoa(system.ID))
	if cached := sf.bucket.Get(key); cached != nil {
		return systemAnnotationError(json.NewDecoder(bytes.NewBuffer(cached)).Decode(system))
	}
	var s CrestSystem
	var c CrestConstellation
	var r CrestRegion
	if err := crestGet(system.Source, &s); err != nil {
		return systemAnnotationError(err)
	}
	if err := crestGet(s.Constellation.Source, &c); err != nil {
		return systemAnnotationError(err)
	}
	if err := crestGet(c.Region.Source, &r); err != nil {
		return systemAnnotationError(err)
	}
	system.Region = r.Name
	system.Constellation = c.Name
	system.SecurityStatus = s.SecurityStatus
	value, err := json.Marshal(system)
	if err != nil {
		return systemAnnotationError(err)
	}
	return systemAnnotationError(sf.bucket.Put(key, value))
}

func crestGet(url string, ret interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	return json.NewDecoder(resp.Body).Decode(ret)
}