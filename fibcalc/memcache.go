package fibcalc

import (
	"strconv"
	"sync"

	"github.com/bradfitz/gomemcache/memcache"
)

const (
	memcacheAddres = "127.0.0.1:11211"
)

//Cache to access sigltone memcache
var (
	Cache *cache
	once  sync.Once
)

//CacheGet get singletone memcahce
func CacheGet() *cache {
	once.Do(func() {
		Cache = new(cache)
		Cache.CreateMemcache()
	})
	return Cache
}

//FibNumber object type for memcache
type fibNumber struct {
	index string
	value int
}

//Cache for memcache client
type cache struct {
	client *memcache.Client
}

//CreateMemcache init memcahce by addres
func (cache *cache) CreateMemcache() {
	cache.client = memcache.New(memcacheAddres)
}

//GetValue return value from memcache if exists or error
func (cache *cache) GetValue(x int) (int, error) {
	fetchItem, err := cache.client.Get(strconv.Itoa(x))
	if err != nil {
		//fmt.Println("Cache Miss")
		return 0, err
	}
	val, err := strconv.Atoi(string(fetchItem.Value))
	//fmt.Println("Cache Hit")
	return val, nil
}

//SetValue sets index, value for memcache
func (cache *cache) SetValue(index, value int) error {
	bs := []byte(strconv.Itoa(value))
	setItem := memcache.Item{
		Key:   strconv.Itoa(index),
		Value: bs}

	return cache.client.Set(&setItem)
}
