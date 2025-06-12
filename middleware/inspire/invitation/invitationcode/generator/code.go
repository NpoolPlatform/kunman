package codegenerator

import (
	"github.com/AmirSoleimani/VoucherCodeGenerator/vcgen"
	"github.com/NpoolPlatform/kunman/framework/wlog"
)

const InvitationCodeLen = 12

func Generate() (string, error) {
	vc, err := vcgen.NewWithOptions(
		vcgen.SetCount(1),
		vcgen.SetPattern("###-###-####"),
		vcgen.SetCharset("1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"),
	)
	if err != nil {
		return "", wlog.Errorf("fail construct invitation code generator: %v", err)
	}
	codes, err := vc.Run()
	if err != nil {
		return "", wlog.Errorf("fail run invitation code generator: %v", err)
	}
	return codes[0], nil
}
