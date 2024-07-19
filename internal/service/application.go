package service

import (
	"proxy-server/internal/cache"
	"proxy-server/internal/service/jsonlog"
	"sync"
)

type Application struct {
	Config Config
	Logger *jsonlog.Logger
	WG     sync.WaitGroup
	Cache  *cache.Cache
}

const version = "1.0"
