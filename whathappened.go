// Copyright 2015 Robert S. Gerus. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// WhatHappened is a small "log" utility to give non-technical people idea about
// what's happening.
package main

import (
	"fmt"
	"log"
	"net/http"
)

var _ = fmt.Println
var _ = http.ListenAndServe

func main() {
	config := setup()

	log.Println("setup:", config)

	http.HandleFunc("/submit", submit)
	http.HandleFunc("/submit_form", submit_form)
	log.Fatalln(http.ListenAndServe(config.Listen, nil))
}
