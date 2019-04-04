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
	"math/rand"
	"os"
	"time"

	"github.com/drhodes/golorem"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"gopkg.in/kyokomi/emoji.v1"
)

// return a randomly-selected emoji character
func randEmoji() string {
	// FIXME: larger corpus!
	corpus := []string{
		":beer:", ":+1:", ":pizza:",
	}
	return emoji.Sprint(corpus[rand.Intn(len(corpus))])
}

var (
	logspewEmojiContent       bool
	logspewEmojiFieldNames    bool
	logspewFieldCount         int
	logspewFieldContentLength int
	logspewInterval           time.Duration
	// logspewCmd represents the logspew command
	logspewCmd = &cobra.Command{
		Use:   "logspew",
		Short: "Emit an infinite stream of random logs",
		Long: `Generate a stream of random JSON-formatted log data and emit it to stdout.
  Optionally include emoji for testing UTF8 conformance.`,
		Run: func(cmd *cobra.Command, args []string) {
			rand.Seed(time.Now().Unix())
			logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
			for i := 0; true; i++ {
				event := logger.Info()
				for i := 0; i < logspewFieldCount; i++ {
					name := fmt.Sprintf("spew%d", i)
					if logspewEmojiFieldNames {
						name = name + randEmoji()
					}
					value := lorem.Sentence(logspewFieldContentLength/3, logspewFieldContentLength)
					if logspewEmojiContent {
						value = value + randEmoji()
					}
					event = event.Str(name, value)
				}
				event.Int("seq", i)
				event.Msg("logspew")
				time.Sleep(logspewInterval)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(logspewCmd)
	logspewCmd.Flags().BoolVar(&logspewEmojiContent, "emoji-in-content", true, "Include emoji in the output fields")
	logspewCmd.Flags().BoolVar(&logspewEmojiFieldNames, "emoji-in-field-names", false, "Include emoji in the output field names")
	logspewCmd.Flags().IntVar(&logspewFieldCount, "field-count", 5, "How many additional JSON fields to each log entry")
	logspewCmd.Flags().IntVar(&logspewFieldContentLength, "field-content-length", 15, "How long each additional JSON fields should be")
	logspewCmd.Flags().DurationVarP(&logspewInterval, "interval", "i", 5*time.Second, "Interval between log entries")
}
