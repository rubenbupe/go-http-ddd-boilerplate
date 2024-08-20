package di

import (
	"slices"

	"github.com/rubenbupe/go-auth-server/internal/shared/platform/di/app"
	"github.com/rubenbupe/go-auth-server/internal/shared/platform/di/auth"
	"github.com/rubenbupe/go-auth-server/internal/shared/platform/di/shared"
)

var defaultDefs = slices.Concat(app.Defs, auth.Defs, shared.Defs)
