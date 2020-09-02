// Code generated by monitor-code-gen. DO NOT EDIT.

package varnish

import (
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "telegraf/varnish"

var groupSet = map[string]bool{}

const (
	varnishBackendBusy      = "varnish.backend_busy"
	varnishBackendConn      = "varnish.backend_conn"
	varnishBackendFail      = "varnish.backend_fail"
	varnishBackendRecycle   = "varnish.backend_recycle"
	varnishBackendReq       = "varnish.backend_req"
	varnishBackendReuse     = "varnish.backend_reuse"
	varnishBackendToolate   = "varnish.backend_toolate"
	varnishBackendUnhealthy = "varnish.backend_unhealthy"
	varnishCacheHit         = "varnish.cache_hit"
	varnishCacheHitpass     = "varnish.cache_hitpass"
	varnishCacheMiss        = "varnish.cache_miss"
	varnishClientReq        = "varnish.client_req"
	varnishNBackend         = "varnish.n_backend"
	varnishNLruNuked        = "varnish.n_lru_nuked"
	varnishSessDropped      = "varnish.sess_dropped"
	varnishSessQueued       = "varnish.sess_queued"
	varnishThreadQueueLen   = "varnish.thread_queue_len"
	varnishThreads          = "varnish.threads"
	varnishThreadsCreated   = "varnish.threads_created"
	varnishThreadsFailed    = "varnish.threads_failed"
	varnishThreadsLimited   = "varnish.threads_limited"
)

var metricSet = map[string]monitors.MetricInfo{
	varnishBackendBusy:      {Type: datapoint.Gauge},
	varnishBackendConn:      {Type: datapoint.Gauge},
	varnishBackendFail:      {Type: datapoint.Gauge},
	varnishBackendRecycle:   {Type: datapoint.Gauge},
	varnishBackendReq:       {Type: datapoint.Gauge},
	varnishBackendReuse:     {Type: datapoint.Gauge},
	varnishBackendToolate:   {Type: datapoint.Gauge},
	varnishBackendUnhealthy: {Type: datapoint.Gauge},
	varnishCacheHit:         {Type: datapoint.Counter},
	varnishCacheHitpass:     {Type: datapoint.Counter},
	varnishCacheMiss:        {Type: datapoint.Counter},
	varnishClientReq:        {Type: datapoint.Counter},
	varnishNBackend:         {Type: datapoint.Gauge},
	varnishNLruNuked:        {Type: datapoint.Counter},
	varnishSessDropped:      {Type: datapoint.Gauge},
	varnishSessQueued:       {Type: datapoint.Gauge},
	varnishThreadQueueLen:   {Type: datapoint.Gauge},
	varnishThreads:          {Type: datapoint.Gauge},
	varnishThreadsCreated:   {Type: datapoint.Gauge},
	varnishThreadsFailed:    {Type: datapoint.Gauge},
	varnishThreadsLimited:   {Type: datapoint.Gauge},
}

var defaultMetrics = map[string]bool{
	varnishBackendFail:      true,
	varnishBackendReq:       true,
	varnishBackendUnhealthy: true,
	varnishCacheHit:         true,
	varnishCacheMiss:        true,
	varnishClientReq:        true,
	varnishSessDropped:      true,
	varnishSessQueued:       true,
	varnishThreadQueueLen:   true,
	varnishThreads:          true,
	varnishThreadsCreated:   true,
	varnishThreadsFailed:    true,
	varnishThreadsLimited:   true,
}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:     "telegraf/varnish",
	DefaultMetrics:  defaultMetrics,
	Metrics:         metricSet,
	SendUnknown:     false,
	Groups:          groupSet,
	GroupMetricsMap: groupMetricsMap,
	SendAll:         false,
}