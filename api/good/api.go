package api

import (
	appdelegatedstaking "github.com/NpoolPlatform/kunman/api/good/app/delegatedstaking"
	appfee "github.com/NpoolPlatform/kunman/api/good/app/fee"
	appgood "github.com/NpoolPlatform/kunman/api/good/app/good"
	"github.com/NpoolPlatform/kunman/api/good/app/good/comment"
	default1 "github.com/NpoolPlatform/kunman/api/good/app/good/default"
	"github.com/NpoolPlatform/kunman/api/good/app/good/description"
	displaycolor "github.com/NpoolPlatform/kunman/api/good/app/good/display/color"
	displayname "github.com/NpoolPlatform/kunman/api/good/app/good/display/name"
	"github.com/NpoolPlatform/kunman/api/good/app/good/label"
	"github.com/NpoolPlatform/kunman/api/good/app/good/like"
	"github.com/NpoolPlatform/kunman/api/good/app/good/poster"
	"github.com/NpoolPlatform/kunman/api/good/app/good/recommend"
	appgoodrequired "github.com/NpoolPlatform/kunman/api/good/app/good/required"
	"github.com/NpoolPlatform/kunman/api/good/app/good/score"
	"github.com/NpoolPlatform/kunman/api/good/app/good/topmost"
	topmostconstraint "github.com/NpoolPlatform/kunman/api/good/app/good/topmost/constraint"
	topmostgood "github.com/NpoolPlatform/kunman/api/good/app/good/topmost/good"
	topmostgoodconstraint "github.com/NpoolPlatform/kunman/api/good/app/good/topmost/good/constraint"
	topmostgoodposter "github.com/NpoolPlatform/kunman/api/good/app/good/topmost/good/poster"
	topmostposter "github.com/NpoolPlatform/kunman/api/good/app/good/topmost/poster"
	apppowerrental "github.com/NpoolPlatform/kunman/api/good/app/powerrental"
	apppowerrentalsimulate "github.com/NpoolPlatform/kunman/api/good/app/powerrental/simulate"
	delegatedstaking "github.com/NpoolPlatform/kunman/api/good/delegatedstaking"
	devicetype "github.com/NpoolPlatform/kunman/api/good/device"
	manufacturer "github.com/NpoolPlatform/kunman/api/good/device/manufacturer"
	deviceposter "github.com/NpoolPlatform/kunman/api/good/device/poster"
	fee "github.com/NpoolPlatform/kunman/api/good/fee"
	"github.com/NpoolPlatform/kunman/api/good/good"
	goodcoin "github.com/NpoolPlatform/kunman/api/good/good/coin"
	"github.com/NpoolPlatform/kunman/api/good/good/coin/reward/history"
	malfunction "github.com/NpoolPlatform/kunman/api/good/good/malfunction"
	"github.com/NpoolPlatform/kunman/api/good/good/required"
	powerrental "github.com/NpoolPlatform/kunman/api/good/powerrental"
	"github.com/NpoolPlatform/kunman/api/good/vender/brand"
	"github.com/NpoolPlatform/kunman/api/good/vender/location"

	v1 "github.com/NpoolPlatform/kunman/message/good/gateway/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	v1.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	v1.RegisterGatewayServer(server, &Server{})
	devicetype.Register(server)
	manufacturer.Register(server)
	deviceposter.Register(server)
	brand.Register(server)
	location.Register(server)
	good.Register(server)
	comment.Register(server)
	description.Register(server)
	displayname.Register(server)
	displaycolor.Register(server)
	like.Register(server)
	label.Register(server)
	poster.Register(server)
	recommend.Register(server)
	required.Register(server)
	appgoodrequired.Register(server)
	history.Register(server)
	score.Register(server)
	appgood.Register(server)
	appfee.Register(server)
	fee.Register(server)
	powerrental.Register(server)
	default1.Register(server)
	topmost.Register(server)
	topmostconstraint.Register(server)
	topmostposter.Register(server)
	topmostgood.Register(server)
	topmostgoodconstraint.Register(server)
	topmostgoodposter.Register(server)
	apppowerrentalsimulate.Register(server)
	apppowerrental.Register(server)
	goodcoin.Register(server)
	malfunction.Register(server)
	delegatedstaking.Register(server)
	appdelegatedstaking.Register(server)
}

//nolint:gocyclo,funlen
func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := devicetype.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := manufacturer.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := deviceposter.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := brand.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := location.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := good.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := comment.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := description.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := displayname.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := displaycolor.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := like.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := label.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := poster.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := recommend.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := required.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := appgoodrequired.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := history.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := score.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := appgood.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := fee.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := powerrental.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := appfee.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := default1.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := topmost.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := topmostconstraint.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := topmostposter.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := topmostgood.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := topmostgoodconstraint.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := topmostgoodposter.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := apppowerrentalsimulate.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := apppowerrental.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := goodcoin.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := malfunction.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := delegatedstaking.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := appdelegatedstaking.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
