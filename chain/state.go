package chain

import (
	"fmt"
	"maps"
	"slices"
	"sync"
)

type State struct {
	mtx         sync.RWMutex
	authority   Address
	balances    map[Address]uint64
	nonces      map[Address]uint64
	lastBlock   SigBlock
	genesisHash Hash
	txs         map[Hash]SigTx
	Pending     *State
}

func NewState(gen SigGenesis) *State {
	return &State{
		authority:   gen.Authority,
		balances:    maps.Clone(gen.Balances),
		nonces:      make(map[Address]uint64),
		genesisHash: gen.Hash(),
		txs:         make(map[Hash]SigTx),
		Pending: &State{
			authority:   gen.Authority,
			balances:    maps.Clone(gen.Balances),
			nonces:      make(map[Address]uint64),
			genesisHash: gen.Hash(),
			txs:         make(map[Hash]SigTx),
		},
	}
}

func (s *State) Clone() *State {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return &State{
		authority:   s.authority,
		balances:    maps.Clone(s.balances),
		nonces:      maps.Clone(s.nonces),
		lastBlock:   s.lastBlock,
		genesisHash: s.genesisHash,
		txs:         maps.Clone(s.txs),
		Pending: &State{
			txs: maps.Clone(s.Pending.txs),
		},
	}
}

func (s *State) ApplyTx(tx SigTx) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	valid, err := VerifyTx(tx)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("tx: invalid transaction signature\n%v\n", tx)
	}
	if tx.Nonce != s.nonces[tx.From]+1 {
		return fmt.Errorf("tx: invalid transaction nonce\n%v\n", tx)
	}
	if s.balances[tx.From] < tx.Value {
		return fmt.Errorf("tx: insufficient account funds\n%v\n", tx)
	}
	s.balances[tx.From] -= tx.Value
	s.balances[tx.To] += tx.Value
	s.nonces[tx.From]++
	s.txs[tx.Hash()] = tx
	return nil
}

func (s *State) CreateBlock(authority Account) (SigBlock, error) {
	pndTxs := make([]SigTx, 0, len(s.Pending.txs))
	for _, tx := range s.Pending.txs {
		pndTxs = append(pndTxs, tx)
	}
	slices.SortFunc(pndTxs, func(a, b SigTx) int {
		if a.Time.Before(b.Time) {
			return -1
		}
		if b.Time.Before(a.Time) {
			return 1
		}
		return 0
	})
	txs := make([]SigTx, 0, len(pndTxs))
	for _, tx := range pndTxs {
		err := s.ApplyTx(tx)
		if err != nil {
			fmt.Printf("tx error: rejected: %v\n", err)
			continue
		}
		txs = append(txs, tx)
	}
	if len(txs) == 0 {
		return SigBlock{}, fmt.Errorf("empty list of valid pending transactions")
	}
	var parent Hash
	if s.lastBlock.Number == 0 {
		parent = s.genesisHash
	} else {
		parent = s.lastBlock.Hash()
	}
	blk, err := NewBlock(s.lastBlock.Number+1, parent, txs)
	if err != nil {
		return SigBlock{}, err
	}
	return authority.SignBlock(blk)
}

func (s *State) ApplyBlock(blk SigBlock) error {
	valid, err := VerifyBlock(blk, s.authority)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("blk error: invalid block signature\n%v", blk)
	}
	if blk.Number != s.lastBlock.Number+1 {
		return fmt.Errorf("blk error: invalid block number\n%v", blk)
	}
	var parent Hash
	if blk.Number == 1 {
		parent = s.genesisHash
	} else {
		parent = s.lastBlock.Hash()
	}
	if blk.Parent != parent {
		return fmt.Errorf("blk error: invalid parent hash\n%v", blk)
	}
	merkleTree, err := MerkleHash(blk.Txs, TxHash, TxPairHash)
	if err != nil {
		return err
	}
	merkleRoot := merkleTree[0]
	if merkleRoot != blk.MerkleRoot {
		return fmt.Errorf("blk error: invalid Merkle root\n%v", blk)
	}
	for _, tx := range blk.Txs {
		err := s.ApplyTx(tx)
		if err != nil {
			return err
		}
	}
	s.lastBlock = blk
	return nil
}
