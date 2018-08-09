package metric

import (
	"bufio"
	"fmt"
	"github.com/go-clog/clog"
	"github.com/rcrowley/go-metrics"
	"github.com/wuleying/silver-framework/globals"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Addr          *net.TCPAddr
	Registry      metrics.Registry
	FlushInterval time.Duration
	DurationUnit  time.Duration
	Prefix        string
	Percentiles   []float64
}

func Metric(r metrics.Registry, d time.Duration, prefix string, addr *net.TCPAddr) {
	WithConfig(Config{
		Addr:          addr,
		Registry:      r,
		FlushInterval: d,
		DurationUnit:  time.Nanosecond,
		Prefix:        prefix,
		Percentiles:   []float64{0.5, 0.75, 0.95, 0.99, 0.999},
	})
}

func WithConfig(c Config) {
	for range time.Tick(c.FlushInterval) {
		if err := metric(&c); nil != err {
			clog.Fatal(globals.ClogSkipDisplayInfo, err.Error())
		}
	}
}

func Once(c Config) error {
	return metric(&c)
}

func metric(c *Config) error {
	now := time.Now().Unix()
	du := float64(c.DurationUnit)
	flushSeconds := float64(c.FlushInterval) / float64(time.Second)
	conn, err := net.DialTCP("tcp", nil, c.Addr)
	if nil != err {
		return err
	}
	defer conn.Close()
	w := bufio.NewWriter(conn)
	c.Registry.Each(func(name string, i interface{}) {
		switch metric := i.(type) {
		case metrics.Counter:
			count := metric.Count()
			fmt.Fprintf(w, "%s.%s.count %d %d\n", c.Prefix, name, count, now)
			fmt.Fprintf(w, "%s.%s.count_ps %.2f %d\n", c.Prefix, name, float64(count)/flushSeconds, now)
		case metrics.Gauge:
			fmt.Fprintf(w, "%s.%s.value %d %d\n", c.Prefix, name, metric.Value(), now)
		case metrics.GaugeFloat64:
			fmt.Fprintf(w, "%s.%s.value %f %d\n", c.Prefix, name, metric.Value(), now)
		case metrics.Histogram:
			h := metric.Snapshot()
			ps := h.Percentiles(c.Percentiles)
			fmt.Fprintf(w, "%s.%s.count %d %d\n", c.Prefix, name, h.Count(), now)
			fmt.Fprintf(w, "%s.%s.min %d %d\n", c.Prefix, name, h.Min(), now)
			fmt.Fprintf(w, "%s.%s.max %d %d\n", c.Prefix, name, h.Max(), now)
			fmt.Fprintf(w, "%s.%s.mean %.2f %d\n", c.Prefix, name, h.Mean(), now)
			fmt.Fprintf(w, "%s.%s.std-dev %.2f %d\n", c.Prefix, name, h.StdDev(), now)
			for psIdx, psKey := range c.Percentiles {
				key := strings.Replace(strconv.FormatFloat(psKey*100.0, 'f', -1, 64), ".", "", 1)
				fmt.Fprintf(w, "%s.%s.%s-percentile %.2f %d\n", c.Prefix, name, key, ps[psIdx], now)
			}
		case metrics.Meter:
			m := metric.Snapshot()
			fmt.Fprintf(w, "%s.%s.count %d %d\n", c.Prefix, name, m.Count(), now)
			fmt.Fprintf(w, "%s.%s.one-minute %.2f %d\n", c.Prefix, name, m.Rate1(), now)
			fmt.Fprintf(w, "%s.%s.five-minute %.2f %d\n", c.Prefix, name, m.Rate5(), now)
			fmt.Fprintf(w, "%s.%s.fifteen-minute %.2f %d\n", c.Prefix, name, m.Rate15(), now)
			fmt.Fprintf(w, "%s.%s.mean %.2f %d\n", c.Prefix, name, m.RateMean(), now)
		case metrics.Timer:
			t := metric.Snapshot()
			ps := t.Percentiles(c.Percentiles)
			count := t.Count()
			fmt.Fprintf(w, "%s.%s.count %d %d\n", c.Prefix, name, count, now)
			fmt.Fprintf(w, "%s.%s.count_ps %.2f %d\n", c.Prefix, name, float64(count)/flushSeconds, now)
			fmt.Fprintf(w, "%s.%s.min %d %d\n", c.Prefix, name, t.Min()/int64(du), now)
			fmt.Fprintf(w, "%s.%s.max %d %d\n", c.Prefix, name, t.Max()/int64(du), now)
			fmt.Fprintf(w, "%s.%s.mean %.2f %d\n", c.Prefix, name, t.Mean()/du, now)
			fmt.Fprintf(w, "%s.%s.std-dev %.2f %d\n", c.Prefix, name, t.StdDev()/du, now)
			for psIdx, psKey := range c.Percentiles {
				key := strings.Replace(strconv.FormatFloat(psKey*100.0, 'f', -1, 64), ".", "", 1)
				fmt.Fprintf(w, "%s.%s.%s-percentile %.2f %d\n", c.Prefix, name, key, ps[psIdx]/du, now)
			}
			fmt.Fprintf(w, "%s.%s.one-minute %.2f %d\n", c.Prefix, name, t.Rate1(), now)
			fmt.Fprintf(w, "%s.%s.five-minute %.2f %d\n", c.Prefix, name, t.Rate5(), now)
			fmt.Fprintf(w, "%s.%s.fifteen-minute %.2f %d\n", c.Prefix, name, t.Rate15(), now)
			fmt.Fprintf(w, "%s.%s.mean-rate %.2f %d\n", c.Prefix, name, t.RateMean(), now)
		default:
			log.Printf("unable to record metric of type %T\n", i)
		}
		w.Flush()
	})
	return nil
}
