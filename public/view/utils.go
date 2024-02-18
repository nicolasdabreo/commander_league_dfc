package view

import (
	"context"
	"fmt"
)

func GetFlash(ctx context.Context, kind string) string {
	if msg := ctx.Value(kind); msg != nil {
		return fmt.Sprintf("%s", msg)
	} else {
		return ""
	}
}
