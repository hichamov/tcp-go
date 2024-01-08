package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/spf13/viper"
)

// This type the server name and port
type Config struct {
  Server string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
  viper.AddConfigPath(path)
  viper.SetConfigName("app")
  viper.SetConfigType("env")

  viper.AutomaticEnv()
  
  err = viper.ReadInConfig()
  if err != nil {
    return
  }

  err = viper.Unmarshal(&config)
  return
}

func main(){
  // Loading configuration
  config, err := LoadConfig(".")
  if err != nil {
    log.Fatal("Connot load config", err)
  }

  conn, err := net.Dial("tcp", config.Server)
  if err != nil {
    fmt.Println(err)
  }

  //Send data to the server
  data := []byte("Hello Server!")

  for { 
  _, err = conn.Write(data)
  if err != nil {
    fmt.Println("Error: ", err)
    return
  }
  time.Sleep(time.Second * 20)

  defer conn.Close()
  }
}
