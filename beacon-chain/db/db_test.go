package db

import "github.com/Dorol-Chain/Prysm-drl/v5/beacon-chain/db/kv"

var _ Database = (*kv.Store)(nil)
