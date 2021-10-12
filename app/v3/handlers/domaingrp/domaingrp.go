package domaingrp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ardanlabs/service/foundation/web"
	"github.com/cpustejovsky/catchall/internal/core/domain/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Handlers struct {
	Store db.Store
}

func (h Handlers) UpdateDelivered(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	domain := r.URL.Query().Get(":domain_name")
	filter := bson.M{
		"name": domain,
	}
	update := bson.D{
		{"$inc", bson.D{{"delivered", 1}}},
	}

	upsert := true
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}
	data, err := h.Store.Update(ctx, filter, update, opt)
	if err != nil {
		return fmt.Errorf("unable to increment delivered property of domain %v: %w", domain, err)
	}
	return web.Respond(ctx, w, data, http.StatusNoContent)
}

func (h Handlers) UpdateBounced(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	domain := r.URL.Query().Get(":domain_name")
	filter := bson.M{
		"name": domain,
	}
	update := bson.D{
		{"$inc", bson.D{{"bounced", 1}}},
	}

	upsert := true
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}
	data, err := h.Store.Update(ctx, filter, update, opt)
	if err != nil {
		return fmt.Errorf("unable to increment bounced property of domain %v: %w", domain, err)
	}
	return web.Respond(ctx, w, data, http.StatusNoContent)
}

func (h Handlers) CheckStatus(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	name := r.URL.Query().Get(":domain_name")
	filter := bson.M{
		"name": name,
	}
	domain, err := h.Store.Get(ctx, filter)
	if err != nil {
		return fmt.Errorf("unable to find domain %v: %w", domain, err)
	}
	bounced := domain["bounced"]
	delivered := domain["delivered"]
	if bounced == nil {
		bounced = 0
	}
	if delivered == nil {
		delivered = 0
	}
	bouncedInt, ok := bounced.(int32)
	if !ok {
		return fmt.Errorf("Bounced property of domain %v did not convert to integer", domain, err)

	}
	deliveredInt, ok := delivered.(int32)
	if !ok {
		return fmt.Errorf("Delivered property of domain %v did not convert to integer", domain, err)

	}
	if bouncedInt >= 1 {
		return web.Respond(ctx, w, "not a catch-all", http.StatusNoContent)

	}
	if deliveredInt >= 1000 {
		return web.Respond(ctx, w, "catch-all", http.StatusNoContent)

	} else {
		return web.Respond(ctx, w, "unknown", http.StatusNoContent)
	}
}
