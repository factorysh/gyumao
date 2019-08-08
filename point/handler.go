package point

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/factorysh/gyumao/rule"
	"github.com/influxdata/influxdb/models"
	log "github.com/sirupsen/logrus"
)

// Crusher reads Points and trigger things
type Crusher struct {
	points  chan models.Points
	rules   rule.Rules
	context map[string]interface{}
}

// New Crusher
func New(rules rule.Rules, context map[string]interface{}) *Crusher {
	return &Crusher{
		points:  make(chan models.Points, 1024),
		rules:   rules,
		context: context,
	}
}

func (p *Crusher) Start() {
	for {
		points := <-p.points
		for _, point := range points {
			if err := p.rules.Visit(point, p.context, func(point models.Point) error {
				fmt.Println(point)
				return nil
			}); err != nil {
				log.WithError(err)
				continue
			}
		}
	}
}

// ServeHTTP is a HTTP handler implementation
func (p *Crusher) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	body, err := r.GetBody()
	if err != nil {
		log.Error(err)
		w.WriteHeader(400)
		return
	}
	buff, err := ioutil.ReadAll(body)
	if err != nil {
		log.Error(err)
		w.WriteHeader(400)
		return
	}
	points, err := models.ParsePointsWithPrecision(buff, time.Now(), "")
	if err != nil {
		log.Error(err)
		w.WriteHeader(500)
		return
	}
	p.points <- points
	w.WriteHeader(http.StatusOK)
}
