package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Store struct {
	keyValue map[string]string
	path     string
	sync.Mutex
}

func CreateStore(dbName string, saveInterval time.Duration) *Store {
	store := &Store{keyValue: map[string]string{}}

	store.path, _ = os.Getwd()
	store.path += "/" + dbName
	if _, err := os.Stat(store.path); err == nil {
		fmt.Println("DB file exist loading data from file")
		store.loadDataFromFile(store.path)
	} else {
		_, err := os.Create(store.path)
		if err != nil {
			fmt.Println("Error Creating File")
			os.Exit(-1)
		}
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
	k.Lock()
	value := k.keyValue[key]
	defer k.Unlock()
	return value
}

func (k Store) Set(key string, value string) {
	k.Lock()
	k.keyValue[key] = value
	defer k.Unlock()
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
