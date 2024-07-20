package service

import (
	"sync"
	"todo-list/internal/cache"
	"todo-list/internal/service/jsonlog"
)

type Application struct {
	Config Config
	Logger *jsonlog.Logger
	WG     sync.WaitGroup
	Cache  *cache.Cache
}

const version = "1.0"
