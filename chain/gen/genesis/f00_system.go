package genesis

import (
	"context"

	"github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/filecoin-project/specs-actors/v2/actors/builtin/system"
	cbor "github.com/ipfs/go-ipld-cbor"

	"github.com/EpiK-Protocol/go-epik/chain/types"
	bstore "github.com/EpiK-Protocol/go-epik/lib/blockstore"
)

func SetupSystemActor(bs bstore.Blockstore) (*types.Actor, error) {
	var st system.State

	cst := cbor.NewCborStore(bs)

	statecid, err := cst.Put(context.TODO(), &st)
	if err != nil {
		return nil, err
	}

	act := &types.Actor{
		Code: builtin.SystemActorCodeID,
		Head: statecid,
	}

	return act, nil
}
