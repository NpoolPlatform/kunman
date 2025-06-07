package goodbase

import (
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type GoodBase interface {
	cruder.CrudBase
	GoodType() types.GoodType
	BenefitType() types.BenefitType
	Name() string
	ServiceStartAt() uint32
	StartMode() types.GoodStartMode
	State() types.GoodState
	TestOnly() bool
	BenefitIntervalHours() uint32
	Purchasable() bool
	Online() bool
}

type goodBase struct {
	_ent *ent.GoodBase
}

func (gb *goodBase) ID() uint32 {
	return gb._ent.ID
}

func (gb *goodBase) EntID() uuid.UUID {
	return gb._ent.EntID
}

func (gb *goodBase) GoodType() types.GoodType {
	return types.GoodType(types.GoodType_value[gb._ent.GoodType])
}

func (gb *goodBase) BenefitType() types.BenefitType {
	return types.BenefitType(types.BenefitType_value[gb._ent.BenefitType])
}

func (gb *goodBase) Name() string {
	return gb._ent.Name
}

func (gb *goodBase) ServiceStartAt() uint32 {
	return gb._ent.ServiceStartAt
}

func (gb *goodBase) StartMode() types.GoodStartMode {
	return types.GoodStartMode(types.GoodStartMode_value[gb._ent.StartMode])
}

func (gb *goodBase) TestOnly() bool {
	return gb._ent.TestOnly
}

func (gb *goodBase) BenefitIntervalHours() uint32 {
	return gb._ent.BenefitIntervalHours
}

func (gb *goodBase) Purchasable() bool {
	return gb._ent.Purchasable
}

func (gb *goodBase) Online() bool {
	return gb._ent.Online
}

func (gb *goodBase) State() types.GoodState {
	return types.GoodState(types.GoodState_value[gb._ent.State])
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
