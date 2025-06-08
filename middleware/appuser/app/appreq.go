package app

import (
	"fmt"

	appcrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/app"
	ctrlcrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/app/control"
	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"

	"github.com/google/uuid"
)

type AppReq struct {
	*appcrud.Req
	Control *ctrlcrud.Req
}

func Grpc2CrudReq(req *npool.AppReq) (*AppReq, error) {
	_req := &AppReq{
		Req: &appcrud.Req{
			Name:        req.Name,
			Logo:        req.Logo,
			Description: req.Description,
		},
		Control: &ctrlcrud.Req{
			SignupMethods:            req.SignupMethods,
			ExtSigninMethods:         req.ExtSigninMethods,
			RecaptchaMethod:          req.RecaptchaMethod,
			KycEnable:                req.KycEnable,
			SigninVerifyEnable:       req.SigninVerifyEnable,
			InvitationCodeMust:       req.InvitationCodeMust,
			CreateInvitationCodeWhen: req.CreateInvitationCodeWhen,
			MaxTypedCouponsPerOrder:  req.MaxTypedCouponsPerOrder,
			Maintaining:              req.Maintaining,
			CommitButtonTargets:      req.CommitButtonTargets,
		},
	}
	if req.EntID != nil {
		id, err := uuid.Parse(*req.EntID)
		if err != nil {
			return nil, err
		}
		_req.EntID = &id
		_req.Control.AppID = &id
	}
	if req.CreatedBy == nil {
		return nil, fmt.Errorf("invalid createdby")
	}
	createdBy, err := uuid.Parse(*req.CreatedBy)
	if err != nil {
		return nil, err
	}
	_req.CreatedBy = &createdBy
	return _req, nil
}
