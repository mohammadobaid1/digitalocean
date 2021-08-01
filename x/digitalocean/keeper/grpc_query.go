package keeper

import (
	"github.com/mohammadobaid1/digitalocean/x/digitalocean/types"
)

var _ types.QueryServer = Keeper{}
