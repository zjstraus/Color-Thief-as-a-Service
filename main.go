/*
Copyright (C) 2020  Zach Strauss

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, version 3.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ericpauley/go-quantize/quantize"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
)

const COLORCOUNT = 10

type requestBody struct {
	URL string `json:"url"`
}

type JSONColor struct {
	R uint32 `json:"r"`
	G uint32 `json:"g"`
	B uint32 `json:"b"`
}

type responseBody struct {
	Palette []JSONColor `json:"palette"`
}

func processPalette(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming data
	decoder := json.NewDecoder(r.Body)
	var reqData requestBody
	readErr := decoder.Decode(&reqData)
	if readErr != nil {
		w.WriteHeader(400)
		w.Write([]byte(readErr.Error()))
		return
	}
	defer r.Body.Close()

	// Get the URL
	resp, respErr := http.Get(reqData.URL)
	if respErr != nil {
		w.WriteHeader(500)
		w.Write([]byte(respErr.Error()))
		return
	}
	defer resp.Body.Close()

	// Build a palette
	img, _, imgDecodeErr := image.Decode(resp.Body)
	if imgDecodeErr != nil {
		w.WriteHeader(500)
		w.Write([]byte(imgDecodeErr.Error()))
		return
	}

	q := quantize.MedianCutQuantizer{}
	palette := q.Quantize(make([]color.Color, 0, COLORCOUNT), img)

	var respBody responseBody
	var red, green, blue uint32
	for _, icolor := range palette {
		red, green, blue, _ = icolor.RGBA()
		respBody.Palette = append(respBody.Palette, JSONColor{
			R: red >> 8,
			G: green >> 8,
			B: blue >> 8,
		})
	}
	json.NewEncoder(w).Encode(respBody)
}

func main() {
	port := flag.Int64("port", 8080, "Port to start HTTP server on")

	flag.Parse()
	http.HandleFunc("/", processPalette)

	log.Printf("Starting HTTP listener on port %d\n", *port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
