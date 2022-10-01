package handlers

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"test/internal/config"
	"test/internal/rabbitmq"
	"test/internal/utils"
)

func GetDeviceByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := utils.GetVars(w, r)
	if vars == nil || vars["id"] == "" {
		fmt.Fprintf(w, "{'error':'ID required'}")
	}

	response, err := http.Get(config.GET_DEVICES)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(responseData))
}

func GetDeviceByTypeHandler(w http.ResponseWriter, r *http.Request) {
	vars := utils.GetVars(w, r)
	if vars == nil || vars["type"] == "" {
		fmt.Fprintf(w, "{'error':'Type required'}")
	}

	_page := r.URL.Query().Get("page")
	_limit := r.URL.Query().Get("limit")

	page := 0
	limit := 50

	p, err1 := strconv.Atoi(_page)
	if err1 == nil {
		page = p
	}

	l, err2 := strconv.Atoi(_limit)
	if err2 == nil {
		limit = l
	}

	response, err := http.Get(fmt.Sprintf(config.GET_DEVICES_BY_TYPE, vars["type"], page, limit))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(responseData))
}

func GetDeviceByStatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := utils.GetVars(w, r)
	if vars == nil || vars["status"] == "" {
		fmt.Fprintf(w, "{'error':'status required'}")
	}

	_page := r.URL.Query().Get("page")
	_limit := r.URL.Query().Get("limit")

	page := 0
	limit := 50

	p, err1 := strconv.Atoi(_page)
	if err1 == nil {
		page = p
	}

	l, err2 := strconv.Atoi(_limit)
	if err2 == nil {
		limit = l
	}

	response, err := http.Get(fmt.Sprintf(config.GET_DEVICES_BY_STATUS, vars["status"], page, limit))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(responseData))
}

func GetDevices(w http.ResponseWriter, r *http.Request) {
	_page := r.URL.Query().Get("page")
	_limit := r.URL.Query().Get("limit")

	page := 0
	limit := 50

	p, err1 := strconv.Atoi(_page)
	if err1 == nil {
		page = p
	}

	l, err2 := strconv.Atoi(_limit)
	if err2 == nil {
		limit = l
	}

	response, err := http.Get(fmt.Sprintf(config.GET_ALL_DEVICES, page, limit))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(responseData))
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()
	buf := make([]byte, 1024)
	var b bytes.Buffer
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {

			b.WriteString(string(buf[:n]))
		}
	}
	mq := rabbitmq.Init()
	fmt.Println(mq.Publish(b.String()))
}
