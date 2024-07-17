package service

import (
	"proxy-server/internal/service/jsonlog"
	"sync"
)

type Application struct {
	Config Config
	Logger *jsonlog.Logger
	WG     sync.WaitGroup
}

const version = "1.0"
