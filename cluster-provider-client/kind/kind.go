package kindprovider

import (
	"context"
	"strings"
	"sync"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/klog/v2"
	kind "sigs.k8s.io/kind/pkg/cluster"
	"sigs.k8s.io/logical-cluster"

	clusterprovider "github.com/kcp-dev/edge-mc/cluster-provider-client/cluster"
	edgeclient "github.com/kcp-dev/edge-mc/pkg/client/clientset/versioned"
)

// KindClusterProvider is a cluster provider that works with a local Kind instance.
type KindClusterProvider struct {
	kindProvider *kind.Provider
	clientset    edgeclient.Interface
	providerName string
}

// New creates a new KindClusterProvider
func New(clientset edgeclient.Interface, providerName string) KindClusterProvider {
	kindProvider := kind.NewProvider()
	return KindClusterProvider{
		kindProvider: kindProvider,
		providerName: providerName,
		clientset:    clientset,
	}
}

func (k KindClusterProvider) Create(ctx context.Context,
	name logical.Name,
	opts clusterprovider.Options) (clusterprovider.LogicalClusterInfo, error) {
	var resCluster clusterprovider.LogicalClusterInfo

	err := k.kindProvider.Create(string(name), kind.CreateWithKubeconfigPath(opts.KubeconfigPath))
	if err != nil {
		if strings.HasPrefix(err.Error(), "node(s) already exist for a cluster with the name") {
			// TODO: check whether it's the same cluster and return success if true
		} else {
			return resCluster, err
		}
	}

	cfg, err := k.kindProvider.KubeConfig(string(name), false)
	if err != nil {
		return resCluster, err
	}
	resCluster = *clusterprovider.New(cfg, opts)
	return resCluster, err
}

func (k KindClusterProvider) Delete(ctx context.Context,
	name logical.Name,
	opts clusterprovider.Options) error {

	return k.kindProvider.Delete(string(name), opts.KubeconfigPath)
}

func (k KindClusterProvider) List() ([]logical.Name, error) {
	list, err := k.kindProvider.List()
	if err != nil {
		return nil, err
	}
	// TODO: what's the right way to cast []string into []logical.Name ??
	logicalNameList := make([]logical.Name, 0, len(list))
	for _, cluster := range list {
		logicalNameList = append(logicalNameList, logical.Name(cluster))
	}
	return logicalNameList, err
}

func (k KindClusterProvider) Watch() (clusterprovider.Watcher, error) {
	return &KindWatcher{
		ch:        make(chan clusterprovider.WatchEvent),
		provider:  &k,
		clientset: k.clientset}, nil
}

type KindWatcher struct {
	init      sync.Once
	wg        sync.WaitGroup
	ch        chan clusterprovider.WatchEvent
	cancel    context.CancelFunc
	provider  *KindClusterProvider
	clientset edgeclient.Interface
}

func (k *KindWatcher) Stop() {
	if k.cancel != nil {
		k.cancel()
	}
	k.wg.Wait()
	close(k.ch)
}

func (k *KindWatcher) ResultChan() <-chan clusterprovider.WatchEvent {
	k.init.Do(func() {
		ctx, cancel := context.WithCancel(context.Background())
		logger := klog.FromContext(ctx)
		logger.Info("test1")
		k.cancel = cancel
		setClusters := sets.NewString()

		listProviders, err := k.clientset.EdgeV1alpha1().ClusterProviderDescs().Get(ctx, newClusterConfig.Spec.ProviderName, v1.GetOptions{})
		if err != nil {
			logger.Error(err, "failed to get the provider resource")
			return err
		}

		k.wg.Add(1)
		go func() {
			defer k.wg.Done()
			for {
				select {
				// TODO replace the 2 with a param at the cluster-provider-client level
				case <-time.After(2 * time.Second):
					provider := kind.NewProvider()
					list, err := provider.List()
					if err != nil {
						// TODO add logging
						logger.Error(err, "Getting provider list.")
						continue
					}
					newSetClusters := sets.NewString(list...)
					// Check for new clusters.
					for _, cl := range newSetClusters.Difference(setClusters).UnsortedList() {
						logger.Info("Detected a new cluster")
						k.ch <- clusterprovider.WatchEvent{
							Type: watch.Added,
							Name: logical.Name(cl),
						}
					}
					// Check for deleted clusters.
					for _, cl := range setClusters.Difference(newSetClusters).UnsortedList() {
						logger.Info("Detected cluster was deleted.")
						k.ch <- clusterprovider.WatchEvent{
							Type: watch.Deleted,
							Name: logical.Name(cl),
						}
					}
					setClusters = newSetClusters
				case <-ctx.Done():
					return
				}
			}
		}()
	})
	return k.ch
}
