package server

import (
	"fmt"
	"net/http"
	"time"

	"io/ioutil"

	"github.com/influxdata/influxdb/models"
	diskqueue "github.com/nsqio/go-diskqueue"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	queue diskqueue.Interface
}

func (s *Server) Write(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	body, err := r.GetBody()
	if err != nil {
		log.Error(err)
	}
	buff, err := ioutil.ReadAll(body)
	if err != nil {
		log.Error(err)
	}
	points, err := models.ParsePointsWithPrecision(buff, time.Now(), "")
	if err != nil {
		log.Error(err)
	}
	fmt.Println(points)
	err = s.queue.Put(buff)
	if err != nil {
		log.Error(err)
	}
}
