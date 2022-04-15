
package utils

import (
  "os"
  "sync"
)

var (
  configuration Configuration
  onceConfig sync.Once
)

type Configuration struct {
  Dsn string
  DBName string
  TgBotApiKey string
}

func Config() *Configuration {
  onceConfig.Do(
    func(){
      configuration = Configuration{
        Dsn: os.Getenv("MONGO_DSN"),
        DBName: os.Getenv("MONGO_DB"),
        TgBotApiKey: os.Getenv("TG_TOKEN"),
      }
    })
//sosu bibu
  return &configuration
}