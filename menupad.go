package menupad

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"bytes"
	"strings"
)

// Init : init the webserver
func Init() {
	fmt.Println("MenuPad REST API serve at http://localhost:8000")
	http.HandleFunc("/", homeController)
	http.ListenAndServe(":8000", nil)
}

func homeController(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// header should be set first at all	
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	list := []string{ "a", "b", "c", "d" }
	listJSON, err := json.Marshal(list)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write(nil)
		log.Fatal(err)
	}
	// else
	w.WriteHeader(http.StatusOK)
	w.Write(listJSON)
}

// Calculate : sum of x and y
func Calculate(x uint, y uint) uint {
	return x + y
}

// Subtract : subtract y from x
func Subtract(x int, y int) int {
	return x - y
}

// TimeSlotID : from hours, minutes to TimeSlotID
func TimeSlotID(h uint8, m uint8) uint8{
	// h * 6 + minutes / 10
	return h * 6 + m/10
}

// StringFromAssignment : Create a string with expected length from a character
func StringFromAssignment(j int, char rune) string {
	// by default string has empty value
	var s string
	for i:=0; i < j; i++{
		s += string(char)
	}
	return s
}

// StringFromAppendJoin : Create a string with expected length from a character
func StringFromAppendJoin(j int, char rune) string {
	// by default string has empty value
	s := []string{}
	for i:=0; i < j; i++{
		s = append(s, string(char))
	}
	return strings.Join(s, "")
}

// StringFromBuffer : Create a string with expected length from a character
func StringFromBuffer(j int, char rune) string {
	// by default string has empty value
	var buffer bytes.Buffer
	for i:=0; i < j; i++{
		buffer.WriteString(string(char))
	}
	return buffer.String()
}