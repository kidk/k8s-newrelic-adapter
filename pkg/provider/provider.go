package provider

import (
	"sync"

	apimeta "k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"

	"github.com/kubernetes-incubator/custom-metrics-apiserver/pkg/provider"
	"github.com/kidk/k8s-newrelic-adapter/pkg/metriccache"
	"github.com/kidk/k8s-newrelic-adapter/pkg/newrelic"
)

var nsGroupResource = schema.GroupResource{Resource: "namespaces"}

// newRelicProvider is a implementation of provider.MetricsProvider for CloudWatch
type newRelicProvider struct {
	client   dynamic.Interface
	mapper   apimeta.RESTMapper
	nrClient newrelic.Client

	valuesLock  sync.RWMutex
	metricCache *metriccache.MetricCache
}

// NewRelicProvider returns an instance of testingProvider, along with its restful.WebService
// that opens endpoints to post new fake metrics
func NewRelicProvider(client dynamic.Interface, mapper apimeta.RESTMapper, nrClient newrelic.Client, metricCache *metriccache.MetricCache) provider.ExternalMetricsProvider {
	provider := &newRelicProvider{
		client:      client,
		mapper:      mapper,
		nrClient:    nrClient,
		metricCache: metricCache,
	}

	return provider
}
