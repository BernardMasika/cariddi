/*
==========
Cariddi
==========

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see http://www.gnu.org/licenses/.

	@Repository:  https://github.com/edoardottt/cariddi
	@Author:      edoardottt, https://www.edoardoottavianelli.it
*/

package output

import (
	"log"
	"os"
)

//bannerHTML
func BannerHTML(filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString("<html><body><div style='" + "background-color:#4adeff;color:white" + "'><h1>Cariddi</h1>")
	file.WriteString("<ul>")
	file.WriteString("<li><a href='" + "https://github.com/edoardottt/cariddi'" + ">github.com/edoardottt/cariddi</a></li>")
	file.WriteString("<li>edoardottt, <a href='" + "https://www.edoardoottavianelli.it'" + ">edoardoottavianelli.it</a></li>")
	file.WriteString("<li>Released under <a href='" + "http://www.gnu.org/licenses/gpl-3.0.html'" + ">GPLv3 License</a></li></ul></div>")
	file.Close()
}

//AppendOutputToHtml
func AppendOutputToHTML(output string, status string, filename string, isLink bool) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	if isLink {
		var statusColor string
		if status != "" {
			if string(status[0]) == "2" || string(status[0]) == "3" {
				statusColor = "<p style='color:green;display:inline'>" + status + "</p>"
			} else {
				statusColor = "<p style='color:red;display:inline'>" + status + "</p>"
			}
		} else {
			statusColor = status
		}

		if _, err := file.WriteString("<li><a target='_blank' href='" + output + "'>" + output + "</a> " + statusColor + "</li>"); err != nil {
			log.Fatal(err)
		}
	} else {
		if _, err := file.WriteString("<li>" + output + "</li>"); err != nil {
			log.Fatal(err)
		}
	}
	file.Close()
}

//HeaderHTML
func HeaderHTML(header string, filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if _, err := file.WriteString("<h3>" + header + "</h3><ul>"); err != nil {
		log.Fatal(err)
	}
	file.Close()
}

//FooterHTML
func FooterHTML(filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	if _, err := file.WriteString("</ul>"); err != nil {
		log.Fatal(err)
	}
	file.Close()
}

//BannerFooterHTML
func BannerFooterHTML(filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	file.WriteString("<div style='" + "background-color:#4adeff;color:white" + "'>")
	file.WriteString("<ul><li><a href='" + "https://github.com/edoardottt/cariddi'" + ">Contribute to cariddi</a></li>")
	file.WriteString("<li>Released under <a href='" + "http://www.gnu.org/licenses/gpl-3.0.html'" + ">GPLv3 License</a></li></ul></div>")
	file.Close()
}
