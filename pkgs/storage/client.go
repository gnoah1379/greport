package storage

import (
	"github.com/minio/minio-go/v7"
	"github.com/spf13/viper"
	"greport/pkgs/apikey"
	"sync"
)

var (
	clients  map[string]*minio.Client
	endpoint string
	secure   bool
	once     sync.Once
	mu       sync.RWMutex
)

func loadConfig() {
	endpoint = viper.GetString("minio.endpoint")
	secure = viper.GetBool("minio.secure")
}

func GetClient(apiKey string) (*minio.Client, error) {
	once.Do(loadConfig)
	mu.RLock()
	client, found := clients[apiKey]
	mu.RUnlock()
	if !found {
		credential, err := apikey.GetCredential(apiKey)
		if err != nil {
			return nil, err
		}
		client, err := minio.New(endpoint, &minio.Options{Creds: credential, Secure: secure})
		if err != nil {
			return nil, err
		}
		mu.Lock()
		clients[apiKey] = client
		mu.Unlock()
		return client, nil
	}
	return client, nil
}
