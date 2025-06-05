package version

import (
	"github.com/NpoolPlatform/kunman/framework/version"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
)

func Version() (*basetypes.VersionResponse, error) {
	info, err := version.GetVersion()
	if err != nil {
		return nil, err
	}
	return &basetypes.VersionResponse{
		Info: info,
	}, nil
}
