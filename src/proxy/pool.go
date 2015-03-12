package proxy

const (
	ORIGIN_POOL_KEY         = "origin_proxies_pool"
	STABLE_POOL_KEY         = "stable_proxies_pool"
	HEIGH_PEFORM_POOL_KEY   = "high_perform_proxies_pool"
	MAX_STABLE_RECORD_COUNT = 100
)

type SetPool struct {
	PoolKey string
	Store   *SetStore
}

func (setPool *SetPool) All() {
	setPool.Store.All(setPool.PoolKey)
}

func AddToStablePool(proxyAddress string, sec int, store *HashArrayStore) {
	times := store.Get(STABLE_POOL_KEY, proxyAddress)
	if len(times) > MAX_STABLE_RECORD_COUNT {
		times = times[:len(times)-MAX_STABLE_RECORD_COUNT]
		append(times, sec)
	}
	store.Set(STABLE_POOL_KEY, proxyAddress, times)
}

func DeleteFromStablePool(proxyAddress string, store *HashArrayStore) {
	store.Delete(STABLE_POOL_KEY, proxyAddress, times)
}

func AddToHeighPerformPool(proxyAddress string, store *SetStore) {
	store.Add(HEIGH_PEFORM_POOL_KEY, proxyAddress)
}

func DeleteFromHeightPerfromPool(proxyAddress string, store *SetStore) {
	store.Delete(HEIGH_PEFORM_POOL_KEY, proxyAddress)
}
