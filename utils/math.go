package utils

import (
	"fmt"
	"log"
)

type DataConfig struct {
	name   string
	depend []string
}

var cfg = []DataConfig{
	{
		name:   "wrk1",
		depend: []string{"srv1", "srv2", "db1", "db2"},
	},
	{
		name:   "wrk2",
		depend: []string{"srv1", "db2"},
	},
	{
		name: "db1",
	},
	{
		name:   "db2",
		depend: []string{"db1"},
	},
	{
		name:   "srv1",
		depend: []string{"srv2"},
	},
	{
		name:   "srv2",
		depend: []string{"db2"},
	},
	{
		name:   "db3",
		depend: []string{"wrk1"},
	},
}

var started []DataConfig

func filter(target []DataConfig, name string) bool {
	for _, t := range target {
		if t.name == name {
			return true
		}
	}
	return false
}

func start(toStart DataConfig) {
	started = append(started, toStart)
}

func buildDependTree() []DataConfig {
	for {
		for _, v := range cfg {
			// Check is current config started
			if filter(started, v.name) {
				continue
			}

			if len(v.depend) == 0 {
				// Start config
				start(v)
			} else {
				calcStarted := 0
				// Fetch all depends
				for _, vdepend := range v.depend {
					if filter(started, vdepend) {
						calcStarted++
					}
					// Check is dapendency exist
					if !filter(cfg, vdepend) {
						log.Fatal("Dependency not found: ", vdepend)
					}
				}
				// If all depends started
				if len(v.depend) == calcStarted {
					// Start config
					start(v)
				}
			}
		}
		if len(started) == len(cfg) {
			break
		}
	}
	return started
}

func TestWrkDependencies() {
	fmt.Printf("Cfg records: %d \n", len(cfg))
	dependToStart := buildDependTree()
	fmt.Printf("Started: %v\n", dependToStart)
}
