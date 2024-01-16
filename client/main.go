package main

import (
//	"fmt"
	"log"
	"net"
	"time"
	"github.com/spf13/viper"
)

// Connection function
func Connection(config Config) (connection net.Conn, err error){
  var reserror error
  for i := 0; i <= 5 ; i++ {
    // Initiating a connection
    conn, err := net.Dial("tcp", config.Server)
    if err != nil {
      reserror = err
      if i < 5 {
      log.Println("Cannot obtain connection, retrying ...")
      time.Sleep(time.Second * 5) 
      }
    } else {
      return conn, nil
    }
  }
  return nil, reserror
}


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

  //conn, err := net.Dial("tcp", config.Server)
  conn, err := Connection(config)
  if err != nil {
    log.Fatal("Could not reach server ", err)
  }else {
    log.Println("Connted to Server !!!")
  }

  //Send data to the server
  data := []byte("Hello Server!")

  for {
  log.Println("Writing Data ...")
  _, err = conn.Write(data)
  if err != nil {
    log.Println("Error: ", err)
    panic("Connection closed !!!")
  }
  time.Sleep(time.Second * 10)

  defer conn.Close()
  }
}
