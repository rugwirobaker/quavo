package main

import (
	"context"
	"log"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rugwirobaker/quavo/models/cache"
)

//CacheService implements cache.Cacheserver
type CacheService struct {
	store map[string][]byte
	mutex *sync.Mutex
}

//NewCacheService is a constructor for CacheService.
func NewCacheService() *CacheService {
	return &CacheService{
		store: make(map[string][]byte),
		mutex: &sync.Mutex{},
	}
}

//Get gets a value given the key.
func (c *CacheService) Get(ctx context.Context, req *cache.GetRequest) (*cache.GetResponse, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	val, ok := c.store[req.Key]
	if !ok {
		log.Printf("key \"%s\" not found\n", req.Key)
		return nil, status.Errorf(codes.NotFound, "key \"%s\" not found", req.Key)
	}
	log.Printf("retrieving key \"%s\"", req.Key)
	return &cache.GetResponse{Value: val}, nil
}

//Set adds a new value to the cache.
func (c *CacheService) Set(ctx context.Context, req *cache.SetRequest) (*cache.SetResponse, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	_, ok := c.store[req.Key]
	if ok {
		log.Printf("key \"%s\"already exists\n", req.Key)
		return nil, status.Errorf(codes.AlreadyExists, "key \"%s\"already exists", req.Key)
	}
	log.Printf("setting key \"%s\"", req.Key)
	c.store[req.Key] = req.Value
	return &cache.SetResponse{Ok: true}, nil
}

//Unset removes a key value pair from the cache.
func (c *CacheService) Unset(ctx context.Context, req *cache.UnsetRequest) (*cache.UnsetResponse, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	_, ok := c.store[req.Key]
	if !ok {
		log.Printf("key \"%s\"does not exists\n", req.Key)
		return &cache.UnsetResponse{Ok: false}, status.Errorf(codes.NotFound, "key \"%s\"does not exists", req.Key)
	}
	log.Printf("deleting key \"%s\"\n", req.Key)
	delete(c.store, req.Key)
	return &cache.UnsetResponse{Ok: true}, nil
}

//Flush deletes all the data in the cache
func (c *CacheService) Flush(ctx context.Context, req *empty.Empty) (*cache.FlushResponse, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if len(c.store) < 1 {
		return &cache.FlushResponse{Ok: false}, status.Errorf(codes.Aborted, "store is already empty")
	}
	log.Printf("flushing the store\n")
	c.store = make(map[string][]byte)
	return &cache.FlushResponse{Ok: true}, nil
}
