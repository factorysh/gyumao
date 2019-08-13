package crusher

import (
	"io/ioutil"
	"net/http"
	"time"

	_consumer "github.com/factorysh/gyumao/consumer"
	_point "github.com/factorysh/gyumao/point"
	"github.com/factorysh/gyumao/probes"
	"github.com/factorysh/gyumao/rule"
	"github.com/influxdata/influxdb/models"
	log "github.com/sirupsen/logrus"
)

// Crusher reads Points and trigger things
type Crusher struct {
	points   chan models.Points
	rules    rule.Rules
	consumer _consumer.Consumer
	probes   probes.Probes
}

// New Crusher
func New(rules rule.Rules, consumer _consumer.Consumer) *Crusher {
	return &Crusher{
		points:   make(chan models.Points, 1024),
		rules:    rules,
		consumer: consumer,
	}
}

// Start the Crusher
func (p *Crusher) Start() {
	for {
		points := <-p.points
		for _, point := range points {
			if err := p.rules.Filter(point,
				func(r *rule.Rule, point models.Point) error {
					return p.consumer.Consume(_point.New(point, r))
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
