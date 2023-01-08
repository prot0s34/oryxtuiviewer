// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

// +build ignore

package main

import (
	"image"
	"log"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"strconv"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"fmt"
	"math"
)

var titles []string
var losses []float64

func main() {

	// parse oryxspioenkop.com as datasource with goquery
	webPage := "https://www.oryxspioenkop.com/2022/02/attack-on-europe-documenting-equipment.html"
	resp, err := http.Get(webPage)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find("h3").Each(func(i int, s *goquery.Selection) {
	  info := s.Text()
	  // parse title/losecount and append to title/count arrays
	  if strings.ContainsAny(info, "(") == true {
		  lose := strings.Split(strings.Split(info, " (")[1], ",")[0]
		  lose64, _ := strconv.ParseFloat(lose, 64)
		  losses = append(losses, lose64)
		  title := strings.Split(info, " (")[0]
		  titles = append(titles, title)
	  }
  })

	all := strings.Split(strings.Split(doc.Find("h3:contains(' - ')").Text(), "- ")[1], ",")[0]
  	destroyed := strings.Split(strings.Split(doc.Find("h3:contains(' - ')").Text(), "destroyed: ")[1], ",")[0]
  	damaged := strings.Split(strings.Split(doc.Find("h3:contains(' - ')").Text(), "damaged: ")[1], ",")[0]
  	abandoned := strings.Split(strings.Split(doc.Find("h3:contains(' - ')").Text(), "abandoned: ")[1], ",")[0]
  	captured := strings.Split(doc.Find("h3:contains(' - ')").Text(), "captured: ")[1]
//	all64, _ := strconv.ParseFloat(all, 64)
  	destroyed64, _ := strconv.ParseFloat(destroyed, 64)
  	damaged64, _ := strconv.ParseFloat(damaged, 64)
  	abandoned64, _ := strconv.ParseFloat(abandoned, 64)
  	captured64, _ := strconv.ParseFloat(captured, 64)

//  init tui, widgets, render & events
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	header := widgets.NewParagraph()
	header.Text = "Press <q> to quit, Press <h> or <l> to switch tabs"
	header.SetRect(0, 0, 50, 1)
	header.Border = false
	header.TextStyle.Bg = ui.ColorBlue

	p2 := widgets.NewParagraph()
	p2.Text = "Vizualize Russia's Losses approved by Oryx\nNavigation:\nPress <q> to quit\nPress <h> or <l> to switch tabs\n\nData by https://www.oryxspioenkop.com/"
	p2.Title = "About"
	p2.SetRect(2, 5, 46, 13)
	p2.BorderStyle.Fg = ui.ColorYellow

	p3 := widgets.NewParagraph()
	p3.Title = "Summary"
	p3.Text = "[Armor lose:](bg:red) " + all
	p3.SetRect(2, 31, 46, 34)
	p3.BorderStyle.Fg = ui.ColorYellow

	pc := widgets.NewPieChart()
	pc.Title = "Summary stat's"
	pc.SetRect(47, 5, 110, 34)
	pc.Data = []float64{destroyed64, damaged64, abandoned64, captured64}
	summary := []string{"destroyed", "damaged", "abandoned", "captured"}
	pc.AngleOffset = -.5 * math.Pi
	pc.LabelFormatter = func(i int, v float64) string {
		return fmt.Sprintf("%.00f %q", v, summary[i])
	}
	pc.BorderStyle.Fg = ui.ColorYellow

	c := ui.NewCanvas()
	c.SetRect(2, 13, 50, 40)
	c.SetLine(image.Pt(15, 80), image.Pt(45, 85), ui.ColorWhite)
	c.SetLine(image.Pt(45, 85), image.Pt(40, 55), ui.ColorWhite)
	c.SetLine(image.Pt(40, 55), image.Pt(55, 55), ui.ColorWhite)
	c.SetLine(image.Pt(55, 55), image.Pt(50, 85), ui.ColorWhite)
	c.SetLine(image.Pt(50, 85), image.Pt(80, 80), ui.ColorWhite)
	c.SetLine(image.Pt(81, 80), image.Pt(80, 95), ui.ColorWhite)
	c.SetLine(image.Pt(80, 95), image.Pt(50, 90), ui.ColorWhite)
	c.SetLine(image.Pt(50, 90), image.Pt(55, 120), ui.ColorWhite)
	c.SetLine(image.Pt(55, 120), image.Pt(40, 120), ui.ColorWhite)
	c.SetLine(image.Pt(40, 120), image.Pt(45, 90), ui.ColorWhite)
	c.SetLine(image.Pt(45, 90), image.Pt(15, 95), ui.ColorWhite)
	c.SetLine(image.Pt(15, 95), image.Pt(14, 80), ui.ColorWhite)

	heavy := widgets.NewBarChart()
	heavy.Title = "Armor Losses - Heavy & Infantry Vehicles"
	heavy.Data = losses[0:6]
	heavy.SetRect(3, 40, 180, 3)
	heavy.BarWidth = 29
	heavy.Labels = titles[0:6]
	heavy.BarColors = []ui.Color{ui.ColorYellow, ui.ColorBlue}

	special := widgets.NewBarChart()
	special.Title = "Armor Losses - Special & Command Vehicles"
	special.Data = losses[6:10]
	special.SetRect(3, 40, 180, 3)
	special.BarWidth = 43
	special.Labels = titles[6:10]
	special.BarColors = []ui.Color{ui.ColorYellow, ui.ColorBlue}

	artmlr := widgets.NewBarChart()
	artmlr.Title = "Armor Losses - Artillery & MRLs"
	artmlr.Data = losses[10:13]
	artmlr.SetRect(3, 40, 180, 3)
	artmlr.BarWidth = 60
	artmlr.Labels = titles[10:13]
	artmlr.BarColors = []ui.Color{ui.ColorYellow, ui.ColorBlue}

	antiair := widgets.NewBarChart()
	antiair.Title = "Armor Losses - Anti-Air Systems"
	antiair.Data = losses[13:16]
	antiair.SetRect(3, 40, 180, 3)
	antiair.BarWidth = 60
	antiair.Labels = titles[13:16]
	antiair.BarColors = []ui.Color{ui.ColorYellow, ui.ColorBlue}

	intel := widgets.NewBarChart()
	intel.Title = "Armor Losses - Radars, Jammers & Deception Systems"
	intel.Data = losses[16:18]
	intel.SetRect(3, 40, 180, 3)
	intel.BarWidth = 90
	intel.Labels = titles[16:18]
	intel.BarColors = []ui.Color{ui.ColorYellow, ui.ColorBlue}

	airuav := widgets.NewBarChart()
	airuav.Title = "Armor Losses - Aircraft, Heli & UAV"
	airuav.Data = losses[18:22]
	airuav.SetRect(3, 40, 180, 3)
	airuav.BarWidth = 43
	airuav.Labels = titles[18:22]
	airuav.BarColors = []ui.Color{ui.ColorYellow, ui.ColorBlue}

	ship := widgets.NewBarChart()
	ship.Title = "Armor Losses - Naval Ships"
	ship.Data = losses[22:23]
	ship.SetRect(3, 40, 180, 3)
	ship.BarWidth = 180
	ship.Labels = titles[22:23]
	ship.BarColors = []ui.Color{ui.ColorYellow, ui.ColorBlue}

	light := widgets.NewBarChart()
	light.Title = "Armor Losses - Trucks, Vehicles & Jeeps"
	light.Data = losses[23:24]
	light.SetRect(3, 40, 180, 3)
	light.BarWidth = 180
	light.Labels = titles[23:24]
	light.BarColors = []ui.Color{ui.ColorYellow, ui.ColorBlue}

	tabpane := widgets.NewTabPane("Menu", "HEAVY", "SPECIAL", "ARTILLERY & MLRS", "ANTI-AIR", "INTELLIGENT SYS", "AIR & UAV", "SHIPS", "OTHER VEHICLES")
	tabpane.SetRect(0, 1, 110, 4)
	tabpane.Border = true

	renderTab := func() {
		switch tabpane.ActiveTabIndex {
		case 0:
			ui.Render(p2)
			ui.Render(c)
			ui.Render(pc)
			ui.Render(p3)
		case 1:
			ui.Render(heavy)
		case 2:
			ui.Render(special)
		case 3:
			ui.Render(artmlr)
		case 4:
			ui.Render(antiair)
		case 5:
			ui.Render(intel)
		case 6:
			ui.Render(airuav)
		case 7:
			ui.Render(ship)
		case 8:
			ui.Render(light)
		}
	}

	ui.Render(header, tabpane, p2, c, pc, p3)

	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "h":
			tabpane.FocusLeft()
			ui.Clear()
			ui.Render(header, tabpane)
			renderTab()
		case "l":
			tabpane.FocusRight()
			ui.Clear()
			ui.Render(header, tabpane)
			renderTab()
		}
	}
}
