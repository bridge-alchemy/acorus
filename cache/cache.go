package cache

import (
	"errors"

	lru "github.com/hashicorp/golang-lru"

	"github.com/cornerstone-labs/acorus/service/models"
)

const ListSize = 1200000

type LruCache struct {
	lruL1ToL2List *lru.Cache
	lruL2ToL1List *lru.Cache
}

func NewLruCache() *LruCache {
	lruL1ToL2List, err := lru.New(ListSize)
	if err != nil {
		panic(errors.New("Failed to init lruL1ToL2List, err :" + err.Error()))
	}
	lruL2ToL1List, err := lru.New(ListSize)
	if err != nil {
		panic(errors.New("Failed to init lruL2ToL1List, err :" + err.Error()))
	}
	return &LruCache{
		lruL1ToL2List: lruL1ToL2List,
		lruL2ToL1List: lruL2ToL1List,
	}
}

func (lc *LruCache) GetL1ToL2List(key string) (*models.DepositsResponse, error) {
	result, ok := lc.lruL1ToL2List.Get(key)
	if !ok {
		return nil, errors.New("lru get L1ToL2 list fail")
	}
	return result.(*models.DepositsResponse), nil
}

func (lc *LruCache) AddL1ToL2List(key string, data *models.DepositsResponse) {
	lc.lruL1ToL2List.PeekOrAdd(key, data)
}

func (lc *LruCache) GetL2ToL1List(key string) (*models.WithdrawsResponse, error) {
	result, ok := lc.lruL2ToL1List.Get(key)
	if !ok {
		return nil, errors.New("lru get L2ToL1 list fail")
	}
	return result.(*models.WithdrawsResponse), nil
}

func (lc *LruCache) AddL2ToL1List(key string, data *models.WithdrawsResponse) {
	lc.lruL2ToL1List.PeekOrAdd(key, data)
}
