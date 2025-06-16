package depc

// rewrite btcd to support depinc

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
)

const sigHashMask = 0x1f

// WitnessSignature creates an input witness stack for tx to spend BTC sent
// from a previous output to the owner of privKey using the p2wkh script
// template. The passed transaction must contain all the inputs and outputs as
// dictated by the passed hashType. The signature generated observes the new
// transaction digest algorithm defined within BIP0143.
func WitnessSignature(tx *wire.MsgTx, sigHashes *txscript.TxSigHashes, idx int, amt int64, hashType txscript.SigHashType, privKey *btcec.PrivateKey, compress bool) (wire.TxWitness, error) {
	pk := (*btcec.PublicKey)(&privKey.PublicKey)
	var pkData []byte
	if compress {
		pkData = pk.SerializeCompressed()
	} else {
		pkData = pk.SerializeUncompressed()
	}

	pkHash := btcutil.Hash160(pkData)
	sig, err := RawTxInWitnessSignature(tx, sigHashes, idx, amt,
		pkHash, hashType, privKey)
	if err != nil {
		return nil, err
	}

	// A witness script is actually a stack, so we return an array of byte
	// slices here, rather than a single byte slice.
	return wire.TxWitness{sig, pkData}, nil
}

func RawTxInWitnessSignature(tx *wire.MsgTx, sigHashes *txscript.TxSigHashes, idx int, amt int64, pkHash []byte, hashType txscript.SigHashType, key *btcec.PrivateKey) ([]byte, error) {
	hash, err := calcWitnessSignatureHash(pkHash, sigHashes, hashType, tx,
		idx, amt)
	if err != nil {
		return nil, err
	}
	signature, err := key.Sign(hash)
	if err != nil {
		return nil, fmt.Errorf("cannot sign tx input: %s", err)
	}

	return append(signature.Serialize(), byte(hashType)), nil
}

// reference DePINC/depinc > src/script/interpreter.cpp > SignatureHash function
func calcWitnessSignatureHash(pubkeyHash []byte, sigHashes *txscript.TxSigHashes, hashType txscript.SigHashType, tx *wire.MsgTx, idx int, amt int64) ([]byte, error) {
	// As a sanity check, ensure the passed input index for the transaction is valid.
	if idx > len(tx.TxIn)-1 {
		return nil, fmt.Errorf("idx %d but %d txins", idx, len(tx.TxIn))
	}

	// We'll utilize this buffer throughout to incrementally calculate
	// the signature hash for this transaction.
	var sigHash bytes.Buffer

	// First write out, then encode the transaction's version number.
	var bVersion [4]byte
	binary.LittleEndian.PutUint32(bVersion[:], uint32(tx.Version))
	sigHash.Write(bVersion[:])

	// Next write out the possibly pre-calculated hashes for the sequence
	// numbers of all inputs, and the hashes of the previous outs for all
	// outputs.
	var zeroHash chainhash.Hash

	// If anyone can pay isn't active, then we can use the cached
	// hashPrevOuts, otherwise we just write zeroes for the prev outs.
	if hashType&txscript.SigHashAnyOneCanPay == 0 {
		sigHash.Write(sigHashes.HashPrevOuts[:])
	} else {
		sigHash.Write(zeroHash[:])
	}

	// If the sighash isn't anyone can pay, single, or none, the use the
	// cached hash sequences, otherwise write all zeroes for the
	// hashSequence.
	if hashType&txscript.SigHashAnyOneCanPay == 0 &&
		hashType&sigHashMask != txscript.SigHashSingle &&
		hashType&sigHashMask != txscript.SigHashNone {
		sigHash.Write(sigHashes.HashSequence[:])
	} else {
		sigHash.Write(zeroHash[:])
	}

	txIn := tx.TxIn[idx]

	// Next, write the outpoint being spent.
	sigHash.Write(txIn.PreviousOutPoint.Hash[:])
	var bIndex [4]byte
	binary.LittleEndian.PutUint32(bIndex[:], txIn.PreviousOutPoint.Index)
	sigHash.Write(bIndex[:])

	sigHash.Write([]byte{0x19})
	sigHash.Write([]byte{txscript.OP_DUP})
	sigHash.Write([]byte{txscript.OP_HASH160})
	sigHash.Write([]byte{txscript.OP_DATA_20})
	sigHash.Write(pubkeyHash)
	sigHash.Write([]byte{txscript.OP_EQUALVERIFY})
	sigHash.Write([]byte{txscript.OP_CHECKSIG})

	// Next, add the input amount, and sequence number of the input being
	// signed.
	var bAmount [8]byte
	binary.LittleEndian.PutUint64(bAmount[:], uint64(amt))
	sigHash.Write(bAmount[:])
	var bSequence [4]byte
	binary.LittleEndian.PutUint32(bSequence[:], txIn.Sequence)
	sigHash.Write(bSequence[:])

	// If the current signature mode isn't single, or none, then we can
	// re-use the pre-generated hashoutputs sighash fragment. Otherwise,
	// we'll serialize and add only the target output index to the signature
	// pre-image.
	if hashType&txscript.SigHashSingle != txscript.SigHashSingle &&
		hashType&txscript.SigHashNone != txscript.SigHashNone {
		sigHash.Write(sigHashes.HashOutputs[:])
	} else if hashType&sigHashMask == txscript.SigHashSingle && idx < len(tx.TxOut) {
		var b bytes.Buffer
		err := wire.WriteTxOut(&b, 0, 0, tx.TxOut[idx])
		if err != nil {
			return nil, err
		}
		sigHash.Write(chainhash.DoubleHashB(b.Bytes()))
	} else {
		sigHash.Write(zeroHash[:])
	}

	// Finally, write out the transaction's locktime, and the sig hash
	// type.
	var bLockTime [4]byte
	binary.LittleEndian.PutUint32(bLockTime[:], tx.LockTime)
	sigHash.Write(bLockTime[:])
	var bHashType [4]byte
	binary.LittleEndian.PutUint32(bHashType[:], uint32(hashType))
	sigHash.Write(bHashType[:])

	// write salt to tail
	saltBytes := []byte("btchd")
	saltBytes, err := txscript.NewScriptBuilder().AddData(saltBytes).Script()
	if err != nil {
		return nil, err
	}
	sigHash.Write(saltBytes)

	return chainhash.DoubleHashB(sigHash.Bytes()), nil
}
