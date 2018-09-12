// Package hosts provides address lookups from hosts file.
package hosts

import (
	"io/ioutil"
	"net"
	"strings"
)

// Hosts represents a file containing hosts_sample
type Hosts struct {
	hl       *hostsLineList
	filePath string
}

func New(path string) (*Hosts, error) {

	if path == "" {
		return nil, nil
	}

	h := &Hosts{filePath: path}
	if err := h.loadHostEntries(); err != nil {
		return nil, err
	}

	return h, nil
}

func (h *Hosts) Find(name string) (ipv4List []net.IP, ipv6List []net.IP) {
	name = strings.TrimSuffix(name, ".")
	return h.hl.FindHosts(name)
}

func (h *Hosts) loadHostEntries() error {
	data, err := ioutil.ReadFile(h.filePath)
	if err != nil {
		return err
	}

	h.hl = newHostsLineList(data)

	return nil
}
