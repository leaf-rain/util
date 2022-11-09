package main

import (
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"net/http"
)

func main() {
	endpoint := "http://127.0.0.1:2379"
	path := "k1"

	viper.RemoteConfig = &Config{}

	v := viper.New()
	v.AddRemoteProvider("etcd", endpoint, path)
	v.SetConfigType("json")
	v.ReadRemoteConfig()
	v.WatchRemoteConfigOnChannel()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, v.GetString("service.password"))
	})

	http.ListenAndServe(":8080", nil)
}
