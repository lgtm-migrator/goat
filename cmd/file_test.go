package cmd

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jtrugman/goat/model"
)

func TestReadYaml(t *testing.T) {
	args := []string{"../example/testRead.yaml"}
	got := readYaml(args)

	want := model.Kid{}
	want.Job.Command.Port = "wlo1"
	want.Job.Command.Operation = "add"
	want.Job.Command.Bitrate.BitrateValue = 3.8
	want.Job.Command.Bitrate.BitrateUnit = "mbit"
	want.Job.Command.Latency = 5.3
	want.Job.Command.PktLoss = 0.5
	want.Job.Command.Jitter = 30.7
	want.Job.Timer.TimeValue = 30.1
	want.Job.Timer.TimeUnit = "seconds"
	want.Job.Link = "downlink"

	if cmp.Equal(got, want) == false {
		t.Errorf("got %+v\n, wanted %+v\n", got, want)
	}
}