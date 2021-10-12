package domain

import (
	"github.com/cpustejovsky/catchall/internal/core/domain/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// Core manages the set of API's for domain access.
type Core struct {
	store db.Store
}

// NewCore constructs a core for domain api access.
func NewCore(log *zap.SugaredLogger, collection *mongo.Collection) Core {
	return Core{
		store: db.NewStore(log, collection),
	}
}

// func (c Core) Update()
