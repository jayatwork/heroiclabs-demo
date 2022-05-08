package main

import (
    "context"
    "database/sql"
    "github.com/heroiclabs/nakama-common/runtime"
)

var (
	errBadInput           = runtime.NewError("input contained invalid data", 3)
	errInternalError      = runtime.NewError("internal server error", 13)
	errGuildAlreadyExists = runtime.NewError("guild name is in use", 6)
	errFullGuild          = runtime.NewError("guild is full", 8)
	errNotAllowed         = runtime.NewError("operation not allowed", 7)
	errNoGuildFound       = runtime.NewError("guild not found", 5)
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
    logger.Info("Hello World!")
    return nil
}
