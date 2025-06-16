package fil

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/endpoints"
	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/utils"
	"github.com/filecoin-project/go-jsonrpc"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/v0api"
)

const (
	MinNodeNum       = 1
	MaxRetries       = 3
	retriesSleepTime = 200 * time.Millisecond
	ToleranceNum     = 2
	EndpointSep      = `|`
	EndpointInvalid  = `fil endpoint invalid`
	EndpointUnsync   = `filecoin chain unsync`
	dialTimeout      = time.Second * 3
)

type FClientI interface {
	GetNode(ctx context.Context, endpointmgr *endpoints.Manager) (v0api.FullNode, jsonrpc.ClientCloser, error)
	WithClient(ctx context.Context, fn func(v0api.FullNode) (bool, error)) error
}

type FClients struct{}

func (fClients FClients) GetNode(ctx context.Context, endpointmgr *endpoints.Manager) (v0api.FullNode, jsonrpc.ClientCloser, error) {
	endpoint, err := endpointmgr.Peek()
	if err != nil {
		return nil, nil, err
	}

	strs := strings.Split(endpoint, EndpointSep)
	if len(strs) != 2 {
		return nil, nil, fmt.Errorf("%v,%v", EndpointInvalid, endpoint)
	}

	addr := strs[0]
	authToken := strs[1]
	headers := http.Header{"Authorization": []string{"Bearer " + authToken}}

	ctx, cancel := context.WithTimeout(ctx, dialTimeout)
	defer cancel()

	var _api v0api.FullNodeStruct
	// internal has conn pool
	closer, err := jsonrpc.NewMergeClient(ctx, "ws://"+addr+"/rpc/v0", "Filecoin", lotusapi.GetInternalStructs(&_api), headers)
	if err != nil {
		return nil, nil, err
	}

	syncState, err := synchronized(ctx, &_api)
	if err != nil {
		closer()
		return nil, nil, err
	}
	if !syncState {
		closer()
		return nil, nil, fmt.Errorf(EndpointUnsync)
	}
	return &_api, closer, nil
}

func synchronized(ctx context.Context, api *v0api.FullNodeStruct) (bool, error) {
	state, err := api.SyncState(ctx)
	if err != nil {
		return false, err
	}
	if len(state.ActiveSyncs) == 0 {
		return false, nil
	}
	working := -1
	for i, ss := range state.ActiveSyncs {
		switch ss.Stage {
		default:
			working = i
		case lotusapi.StageSyncComplete:
		case lotusapi.StageIdle:
			// not complete, not actively working
		}
	}
	if working == -1 {
		working = len(state.ActiveSyncs) - 1
	}

	ss := state.ActiveSyncs[working]
	var heightDiff int64

	if ss.Base != nil {
		heightDiff = int64(ss.Base.Height())
	}
	if ss.Target != nil {
		heightDiff = int64(ss.Target.Height()) - heightDiff
	} else {
		heightDiff = 0
	}
	if heightDiff < ToleranceNum && (ss.Stage == lotusapi.StageSyncComplete || ss.Stage == lotusapi.StageIdle) {
		return true, nil
	}
	return false, nil
}

func (fClients *FClients) WithClient(ctx context.Context, fn func(c v0api.FullNode) (bool, error)) error {
	var (
		apiErr, err error
		retry       bool
		client      v0api.FullNode
		closer      jsonrpc.ClientCloser
	)
	endpointmgr, err := endpoints.NewManager()
	if err != nil {
		return err
	}

	for i := 0; i < utils.MinInt(MaxRetries, endpointmgr.Len()); i++ {
		if i > 0 {
			time.Sleep(retriesSleepTime)
		}

		client, closer, err = fClients.GetNode(ctx, endpointmgr)
		if err != nil {
			continue
		}

		retry, apiErr = fn(client)
		closer()

		if !retry {
			return err
		}
	}
	if apiErr != nil {
		return apiErr
	}
	return err
}

func Client() FClientI {
	return &FClients{}
}
