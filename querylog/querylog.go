package querylog

import (
	"net"
	"time"

	"github.com/AdguardTeam/AdGuardHome/dnsfilter"
	"github.com/miekg/dns"
)

// QueryLog - main interface
type QueryLog interface {
	Close()
	Add(question *dns.Msg, answer *dns.Msg, result *dnsfilter.Result, elapsed time.Duration, addr net.Addr, upstream string)
	GetData() []map[string]interface{}
}

// Config - configuration object
type Config struct {
	BaseDir  string
	Interval time.Duration
}

// New - create instance
func New(conf Config) QueryLog {
	return newQueryLog(conf)
}
