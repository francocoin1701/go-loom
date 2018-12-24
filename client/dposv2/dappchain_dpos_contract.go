// +build evm

package dposv2

import (
	"github.com/loomnetwork/go-loom"
	dpostypes "github.com/loomnetwork/go-loom/builtin/types/dposv2"
	"github.com/loomnetwork/go-loom/client"
)

// DAppChainDPOSContract is a client-side binding for the builtin coin Go contract.
type DAppChainDPOSContract struct {
	contract *client.Contract
	chainID  string

	Address loom.Address
}

func ConnectToDAppChainDPOSContract(loomClient *client.DAppChainRPCClient) (*DAppChainDPOSContract, error) {
	contractAddr, err := loomClient.Resolve("dposV2")
	if err != nil {
		return nil, err
	}

	return &DAppChainDPOSContract{
		contract: client.NewContract(loomClient, contractAddr.Local),
		chainID:  loomClient.GetChainID(),
		Address:  contractAddr,
	}, nil
}

func (dpos *DAppChainDPOSContract) ListCandidates(identity *client.Identity) ([]*dpostypes.CandidateV2, error) {
	owner := loom.Address{
		ChainID: dpos.chainID,
		Local:   identity.LoomAddr.Local,
	}
	req := &dpostypes.ListCandidateRequestV2{}
	var resp dpostypes.ListCandidateResponseV2
	_, err := dpos.contract.StaticCall("ListCandidates", req, owner, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Candidates, err
}

func (dpos *DAppChainDPOSContract) ListValidators(identity *client.Identity) ([]*dpostypes.ValidatorStatisticV2, error) {
	owner := loom.Address{
		ChainID: dpos.chainID,
		Local:   identity.LoomAddr.Local,
	}
	req := &dpostypes.ListValidatorsRequestV2{}
	var resp dpostypes.ListValidatorsResponseV2
	_, err := dpos.contract.StaticCall("ListValidators", req, owner, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Statistics, err
}

func (dpos *DAppChainDPOSContract) RegisterCandidate(identity *client.Identity, pubKey []byte, candidateFee uint64, candidateName string, candidateDescription string, candidateWebsite string) error {
	req := &dpostypes.RegisterCandidateRequestV2{
		PubKey:      pubKey,
		Fee:         candidateFee,
		Name:        candidateName,
		Description: candidateDescription,
		Website:     candidateWebsite,
	}
	_, err := dpos.contract.Call("ListValidators", req, identity.LoomSigner, nil)
	return err
}
