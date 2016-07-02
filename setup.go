// Copyright 2015 Robert S. Gerus. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type whConfig struct {
	Listen  string `yaml:"listen"`
	DB      string `yaml:"db"`
	LogFile string `yaml:"log_file"`
}

func setup() whConfig {
	var err error
	var data []byte
	var config whConfig

	if len(os.Args) != 2 {
		log.Fatalln("Usage:", os.Args[0], "<configuration file>")
	}

	data, err = ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalln("Error reading configuration file:", err)
	}

	if err = yaml.Unmarshal(data, &config); err != nil {
		log.Fatalln("Error parsing configuration file:", err)
	}

	if config.LogFile != "" {
		logfile, err := os.OpenFile(config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalln("Error opening logfile:", err)
		}

		log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	}

	return config
}
