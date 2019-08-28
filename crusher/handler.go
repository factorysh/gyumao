package crusher

import (
	"io/ioutil"
	"net/http"
	"time"

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
	consumer _point.Consumer
	probes   probes.Probes
}

// New Crusher
func New(rules rule.Rules, consumer _point.Consumer, probes probes.Probes) *Crusher {
	return &Crusher{
		points:   make(chan models.Points, 1024),
		rules:    rules,
		consumer: consumer,
		probes:   probes,
	}
}

// Start the Crusher
func (p *Crusher) Start() {
	for {
		points := <-p.points
		for _, point := range points {
			log.WithField("point", point).Info("crusher.Crusher#Start")
			if err := p.rules.Filter(point, p.probes,
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
	l := log.WithField("request", r)

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	buff, err := ioutil.ReadAll(r.Body)
	if err != nil {
		l.WithError(err).Error("read Body")
		w.WriteHeader(400)
		return
	}
	precision := r.FormValue("precision")
	points, err := models.ParsePointsWithPrecision(buff, time.Now(), precision)
	if err != nil {
		l.WithError(err).Error("Parse points")
		w.WriteHeader(500)
		return
	}
	l.WithField("points", points).Info("Crusher handler")
	p.points <- points
	w.WriteHeader(http.StatusOK)
}
