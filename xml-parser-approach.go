package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

type Population struct {
	XMLName    xml.Name             `xml:"population"`
	Attributes PopulationAttributes `xml:"attributes,omitempty"`
	Persons    []Person             `xml:"person"`
}

type PopulationAttributes struct {
	XMLName    xml.Name              `xml:"attributes"`
	Attributes []PopulationAttribute `xml:"attribute"`
}

type PopulationAttribute struct {
	XMLName xml.Name `xml:"attribute"`
	Name    string   `xml:"name,attr"`
	Class   string   `xml:"class,attr"`
	Value   string   `xml:",chardata"`
}

type Person struct {
	XMLName    xml.Name             `xml:"person"`
	Id         string               `xml:"id,attr"`
	Attributes PopulationAttributes `xml:"attributes,omitempty"`
	Plans      []Plan               `xml:"plan"`
}

type Plan struct {
	XMLName    xml.Name   `xml:"plan"`
	Score      string     `xml:"score,attr"`
	Selected   string     `xml:"selected,attr"`
	Activities []Activity `xml:"activity"`
	Legs       []Leg      `xml:"leg"`
}

type Activity struct {
	XMLName    xml.Name             `xml:"activity"`
	Type       string               `xml:"type,attr"`
	X          string               `xml:"x,attr"`
	Y          string               `xml:"y,attr"`
	Facility   string               `xml:"facility,attr,omitempty"`
	StartTime  string               `xml:"start_time,attr,omitempty"`
	EndTime    string               `xml:"end_time,attr,omitempty"`
	MaxDur     string               `xml:"max_dur,attr,omitempty"`
	Attributes PopulationAttributes `xml:"attributes,omitempty"`
}

type Leg struct {
	XMLName    xml.Name             `xml:"leg"`
	Mode       string               `xml:"mode,attr"`
	DepTime    string               `xml:"dep_time,attr,omitempty"`
	TravTime   string               `xml:"trav_time,attr,omitempty"`
	Attributes PopulationAttributes `xml:"attributes,omitempty"`
}

// this didn't work because all activities and legs were ordered afterwards
func xml_approach_main() {
	if len(os.Args) != 2 {
		log.Fatal("Please specify path to plans file!")
	}

	xmlFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer xmlFile.Close()

	content, _ := io.ReadAll(xmlFile)
	var population Population
	err = xml.Unmarshal(content, &population)
	if err != nil {
		log.Fatal(err)
	}

	out, err := xml.MarshalIndent(population, "", "    ")
	log.Println(string(out)[:5000])

	outFile, err := os.Create(fmt.Sprintf("cleaned-%s", os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	w := bufio.NewWriter(outFile)
	w.WriteString("<?xml version=\"1.0\" encoding=\"utf-8\"?>\n")
	w.WriteString("<!DOCTYPE population SYSTEM \"http://www.matsim.org/files/dtd/population_v6.dtd\">\n")
	w.Write(out)
}
