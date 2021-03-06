package genesis

import (
	"context"

	"github.com/EpiK-Protocol/go-epik/chain/types"
	bstore "github.com/EpiK-Protocol/go-epik/lib/blockstore"
	"github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/filecoin-project/specs-actors/v2/actors/builtin/retrieval"
	"github.com/filecoin-project/specs-actors/v2/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)

func SetupRetrievalFundActor(bs bstore.Blockstore) (*types.Actor, error) {
	store := adt.WrapStore(context.TODO(), cbor.NewCborStore(bs))

	emptyMap, err := adt.MakeEmptyMap(store).Root()
	if err != nil {
		return nil, err
	}

	multiMap, err := adt.AsMultimap(store, emptyMap)
	if err != nil {
		return nil, err
	}

	emptyMultiMap, err := multiMap.Root()
	if err != nil {
		return nil, err
	}

	vas := retrieval.ConstructState(emptyMap, emptyMultiMap)
	stcid, err := store.Put(store.Context(), vas)
	if err != nil {
		return nil, err
	}

	return &types.Actor{
		Code:    builtin.RetrievalFundActorCodeID,
		Head:    stcid,
		Nonce:   0,
		Balance: types.NewInt(0),
	}, nil
}
