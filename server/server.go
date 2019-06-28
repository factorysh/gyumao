package server

import (
	"net/http"
	"time"

	"io/ioutil"

	"github.com/influxdata/influxdb/models"
	log "github.com/sirupsen/logrus"
	"github.com/factorysh/gyumao/rule"
)

type Server struct {
	points chan models.Points
	rules  *rule.Rules
}

func New(rules *rule.Rules) *Server {
	s := &Server{
		points: make(chan models.Points, 1024),
		rules:  rules,
	}
	return s
}

func (s *Server) Start() error {
	for {
		points := <-s.points
		for _, point := range points {
			if err := s.rules.Filter(point); err != nil {
				return err
			}
		}
	}
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
	s.points <- points
	w.WriteHeader(http.StatusOK)
}
