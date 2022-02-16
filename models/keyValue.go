package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Store struct {
	keyValue map[string]string `json:"keyValue"`
	path     string
}

func CreateStore(dbName string, saveInterval time.Duration) *Store {
	store := &Store{keyValue: map[string]string{}}

	store.path, _ = os.Getwd()
	store.path += "/" + dbName
	if _, err := os.Stat(store.path); err == nil {
		fmt.Println("DB file exist loading data from file")
		store.loadDataFromFile(store.path)
	} else {
		os.Create(store.path)
	}
	ticker := time.NewTicker(saveInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				store.saveDataToFile(store.path)
			}
		}
	}()
	return store
}

func (k Store) Get(key string) string {
	return k.keyValue[key]
}

func (k Store) Set(key string, value string) {
	k.keyValue[key] = value
}

func (k Store) GetAll() map[string]string {
	return k.keyValue
}

func (k Store) FlushAllData() {
	k.saveDataToFile(k.path)
}

func (k Store) loadDataFromFile(path string) []byte {
	jsonFile, err := os.Open(path)
	byteValue, _ := ioutil.ReadAll(jsonFile)

	defer jsonFile.Close()
	json.Unmarshal(byteValue, &k.keyValue)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Load Successful")
	return byteValue
}

func (k Store) saveDataToFile(path string) {
	file, _ := json.Marshal(k.keyValue)
	_ = ioutil.WriteFile(path, file, 0644)

}
