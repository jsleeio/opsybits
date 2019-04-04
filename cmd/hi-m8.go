// Copyright © 2018 John Slee <john@sleefamily.org>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
)

var (
	him8Histogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_seconds",
		Help:    "HTTP response time",
		Buckets: []float64{1, 2, 5, 10},
	}, []string{"code"})
	him8Delay   time.Duration
	him8Message string
	him8Path    string
	him8Listen  string
	him8Cmd     = &cobra.Command{
		Use:   "hi-m8",
		Short: "Host a single message on an HTTP endpoint",
		Long: `hi-m8 simply listens on an HTTP endpoint and returns a static message.
Optionally, an artificial delay can be added prior to the message being returned.`,
		Run: func(cmd *cobra.Command, args []string) {
			prometheus.Register(him8Histogram)
			http.Handle("/metrics", promhttp.Handler())
			http.HandleFunc("/healthz", healthz)
			http.HandleFunc(him8Path, him8)
			http.ListenAndServe(him8Listen, nil)
		},
	}
)

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK")
}

func him8(w http.ResponseWriter, r *http.Request) {
	time.Sleep(him8Delay)
	startTime := time.Now()
	defer func() {
		r.Body.Close()
		him8Histogram.WithLabelValues(fmt.Sprint(http.StatusOK)).Observe(time.Since(startTime).Seconds())
	}()
	w.Write([]byte(him8Message))
}

func init() {
	rootCmd.AddCommand(him8Cmd)
	him8Cmd.Flags().DurationVar(&him8Delay, "delay", 0*time.Second, "sleep this duration before responding")
	him8Cmd.Flags().StringVar(&him8Message, "message", "hi m8", "specify a message to be returned to clients")
	him8Cmd.Flags().StringVar(&him8Listen, "listen", ":3000", "[address]:port to bind to")
	him8Cmd.Flags().StringVar(&him8Path, "path", "/", "specify a path to respond to")
}
