package handler

import "github.com/NpoolPlatform/kunman/framework/wlog"

func (h *Handler) CheckStartEndAt() error {
	if (h.StartAt != nil && h.EndAt != nil) && *h.StartAt > *h.EndAt {
		return wlog.Errorf("invalid startat and endat")
	}
	return nil
}
