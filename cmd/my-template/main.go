// Package main is the starting point of my-template
package main

import (
	"flag"
	"os"
	"strconv"
	"sync"

	//_ "github.com/go-sql-driver/mysql"
	"github.com/fsnotify/fsnotify"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"

	restServer "github.com/yunkon-kim/my-template/pkg/api/rest/server"

	// Black import (_) is for running a package's init() function without using its other contents.
	"github.com/rs/zerolog/log"
	_ "github.com/yunkon-kim/my-template/internal/config"
	_ "github.com/yunkon-kim/my-template/internal/logger"
)

func main() {

	log.Info().Msg("starting my-template server")

	// Set the default port number "8056" for the REST API server to listen on
	port := flag.String("port", "8056", "port number for the restapiserver to listen to")
	flag.Parse()

	// Validate port
	if portInt, err := strconv.Atoi(*port); err != nil || portInt < 1 || portInt > 65535 {
		log.Fatal().Msgf("%s is not a valid port number. Please retry with a valid port number (ex: -port=[1-65535]).", *port)
	}
	log.Debug().Msgf("port number: %s", *port)

	//Setup database (meta_db/dat/my-template.s3db)
	log.Info().Msg("setting SQL Database")
	err := os.MkdirAll("./meta_db/dat/", os.ModePerm)
	if err != nil {
		log.Error().Err(err).Msg("error creating directory")
	}
	log.Debug().Msgf("database file path: %s", "./meta_db/dat/my-template.s3db")

	// Watch config file changes
	go func() {
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			log.Debug().Str("file", e.Name).Msg("config file changed")
			err := viper.ReadInConfig()
			if err != nil { // Handle errors reading the config file
				log.Fatal().Err(err).Msg("fatal error in config file")
			}
			// err = viper.Unmarshal(&common.RuntimeConf)
			if err != nil {
				log.Panic().Err(err).Msg("error unmarshaling runtime configuration")
			}
		})
	}()

	// Launch API servers (REST)
	wg := new(sync.WaitGroup)
	wg.Add(1)

	// Start REST Server
	go func() {
		restServer.RunServer(*port)
		wg.Done()
	}()

	wg.Wait()
}
