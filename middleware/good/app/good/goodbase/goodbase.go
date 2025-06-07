package goodbase

import (
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type AppGoodBase interface {
	cruder.CrudBase
	AppID() uuid.UUID
	GoodID() uuid.UUID
	Purchasable() bool
	EnableProductPage() bool
	ProductPage() string
	Online() bool
	Visible() bool
	Name() string
	DisplayIndex() int32
	Banner() string
	CommentCount() uint32
	RecommendCount() uint32
	Score() decimal.Decimal
}

type goodBase struct {
	_ent      *ent.AppGoodBase
	_extraEnt *ent.ExtraInfo
}

func (gb *goodBase) ID() uint32 {
	return gb._ent.ID
}

func (gb *goodBase) EntID() uuid.UUID {
	return gb._ent.EntID
}

func (gb *goodBase) AppID() uuid.UUID {
	return gb._ent.AppID
}

func (gb *goodBase) GoodID() uuid.UUID {
	return gb._ent.GoodID
}

func (gb *goodBase) Purchasable() bool {
	return gb._ent.Purchasable
}

func (gb *goodBase) EnableProductPage() bool {
	return gb._ent.EnableProductPage
}

func (gb *goodBase) ProductPage() string {
	return gb._ent.ProductPage
}

func (gb *goodBase) Online() bool {
	return gb._ent.Online
}

func (gb *goodBase) Visible() bool {
	return gb._ent.Visible
}

func (gb *goodBase) Name() string {
	return gb._ent.Name
}

func (gb *goodBase) DisplayIndex() int32 {
	return gb._ent.DisplayIndex
}

func (gb *goodBase) Banner() string {
	return gb._ent.Banner
}

func (gb *goodBase) CreatedAt() uint32 {
	return gb._ent.CreatedAt
}

func (gb *goodBase) UpdatedAt() uint32 {
	return gb._ent.UpdatedAt
}

func (gb *goodBase) DeletedAt() uint32 {
	return gb._ent.DeletedAt
}

func (gb *goodBase) CommentCount() uint32 {
	return gb._extraEnt.CommentCount
}

func (gb *goodBase) RecommendCount() uint32 {
	return gb._extraEnt.RecommendCount
}

func (gb *goodBase) Score() decimal.Decimal {
	return gb._extraEnt.Score
}
