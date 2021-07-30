package sys

import (
	"math"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/wailsapp/wails"
)

type Stats struct {
	log *wails.CustomLogger
}
type CpuUsage struct {
	Average int `json:"avg"`
}
func (s *Stats) WailsInit(runtime *wails.Runtime) error {
	s.log=runtime.Log.New("Stats")
	go func() {
		for {
			runtime.Events.Emit("usage", s.GetCpuUsage())
			time.Sleep(500)
		}
	}()
	return nil
}

func (s *Stats) GetCpuUsage() *CpuUsage {
	per, err := cpu.Percent(time.Second, false)
	if err != nil {
		s.log.Errorf("Cannot get CPU Stats %s", err.Error())
		return nil
	}
	return &CpuUsage{
		Average: int(math.Round(per[0])),
	}
}
