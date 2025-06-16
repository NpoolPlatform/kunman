package rpc

import (
	"io"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CheckCode(err error) bool {
	if err == io.EOF ||
		status.Code(err) == codes.Unavailable ||
		status.Code(err) == codes.Canceled ||
		status.Code(err) == codes.Unimplemented {
		return true
	}
	return false
}
