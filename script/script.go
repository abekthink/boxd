// Copyright (c) 2018 ContentBox Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package script

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"strings"

	"github.com/BOXFoundation/boxd/core/types"
	"github.com/BOXFoundation/boxd/crypto"
	"github.com/BOXFoundation/boxd/log"
)

var logger = log.NewLogger("script") // logger

// PayToPubKeyHashScript creates a script to lock a transaction output to the specified address.
func PayToPubKeyHashScript(pubKeyHash []byte) *Script {
	return NewScript().AddOpCode(OPDUP).AddOpCode(OPHASH160).AddOperand(pubKeyHash).AddOpCode(OPEQUALVERIFY).AddOpCode(OPCHECKSIG)
}

// SignatureScript creates a script to unlock a utxo.
func SignatureScript(sig *crypto.Signature, pubKey []byte) *Script {
	return NewScript().AddOperand(sig.Serialize()).AddOperand(pubKey)
}

// StandardCoinbaseSignatureScript returns a standard signature script for coinbase transaction.
func StandardCoinbaseSignatureScript(height uint32) *Script {
	return NewScript().AddOperand(scriptNum(height).Bytes()).AddOperand(scriptNum(0).Bytes())
}

// Script represents scripts
type Script []byte

// NewScript returns an empty script
func NewScript() *Script {
	emptyBytes := make([]byte, 0)
	return (*Script)(&emptyBytes)
}

// NewScriptFromBytes returns a script from byte slice
func NewScriptFromBytes(scriptBytes []byte) *Script {
	script := Script(scriptBytes)
	return &script
}

// AddOpCode adds an opcode to the script
func (s *Script) AddOpCode(opCode OpCode) *Script {
	*s = append(*s, byte(opCode))
	return s
}

// AddOperand adds an operand to the script
func (s *Script) AddOperand(operand []byte) *Script {
	dataLen := len(operand)

	if dataLen < int(OPPUSHDATA1) {
		*s = append(*s, byte(dataLen))
	} else if dataLen <= 0xff {
		*s = append(*s, byte(OPPUSHDATA1), byte(dataLen))
	} else if dataLen <= 0xffff {
		buf := make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, uint16(dataLen))
		*s = append(*s, byte(OPPUSHDATA2))
		*s = append(*s, buf...)
	} else {
		buf := make([]byte, 4)
		binary.LittleEndian.PutUint32(buf, uint32(dataLen))
		*s = append(*s, byte(OPPUSHDATA4))
		*s = append(*s, buf...)
	}

	// Append the actual operand
	*s = append(*s, operand...)
	return s
}

// AddScript appends a script to the script
func (s *Script) AddScript(script *Script) *Script {
	*s = append(*s, (*script)...)
	return s
}

// Validate verifies the script
func Validate(scriptSig, scriptPubKey *Script, tx *types.Transaction, txInIdx int) error {
	// concatenate unlocking & locking scripts
	catScript := NewScript().AddScript(scriptSig).AddOpCode(OPCODESEPARATOR).AddScript(scriptPubKey)
	if err := catScript.evaluate(tx, txInIdx); err != nil {
		return err
	}

	if !scriptPubKey.IsPayToScriptHash() {
		return nil
	}

	// Handle p2sh
	// scriptSig: signature <serialized redeemScript>
	//

	// First operand is signature
	_, sig, newPc, _ := scriptSig.parseNextOp(0)
	newScriptSig := NewScript().AddOperand(sig)

	// Second operand is serialized redeem script
	_, redeemScriptBytes, _, _ := scriptSig.parseNextOp(newPc)
	redeemScript := NewScriptFromBytes(redeemScriptBytes)

	// signature becomes the new scriptSig, redeemScript becomes the new scriptPubKey
	catScript = NewScript().AddScript(newScriptSig).AddOpCode(OPCODESEPARATOR).AddScript(redeemScript)
	return catScript.evaluate(tx, txInIdx)
}

// Evaluate interprets the script and returns error if it fails
// It succeeds if the script runs to completion and the top stack element exists and is true
func (s *Script) evaluate(tx *types.Transaction, txInIdx int) error {
	script := *s
	scriptLen := len(script)
	logger.Debugf("script len %d: %s", scriptLen, s.disasm())

	stack := newStack()
	for pc, scriptPubKeyStart := 0, 0; pc < scriptLen; {
		opCode, operand, newPc, err := s.parseNextOp(pc)
		if err != nil {
			return err
		}
		pc = newPc

		if err := s.execOp(opCode, operand, tx, txInIdx, pc, &scriptPubKeyStart, stack); err != nil {
			return err
		}
	}

	// Succeed if top stack item is true
	return stack.validateTop()
}

// Get the next opcode & operand. Operand only applies to data push opcodes. Also return incremented pc.
func (s *Script) parseNextOp(pc int) (OpCode, Operand, int, error) {
	script := *s
	scriptLen := len(script)
	if pc >= scriptLen {
		return 0, nil, pc, ErrScriptBound
	}

	opCode := OpCode(script[pc])
	pc++

	if opCode > OPPUSHDATA4 {
		return opCode, nil, pc, nil
	}

	var operandSize int
	if opCode < OPPUSHDATA1 {
		// opcode itself encodes operand size
		operandSize = int(opCode)
	} else if opCode == OPPUSHDATA1 {
		if scriptLen-pc < 1 {
			return opCode, nil, pc, ErrNoEnoughDataOPPUSHDATA1
		}
		// 1 byte after opcode encodes operand size
		operandSize = int(script[pc])
		pc++
	} else if opCode == OPPUSHDATA2 {
		if scriptLen-pc < 2 {
			return opCode, nil, pc, ErrNoEnoughDataOPPUSHDATA2
		}
		// 2 bytes after opcode encodes operand size
		operandSize = int(binary.LittleEndian.Uint16(script[pc : pc+2]))
		pc += 2
	} else if opCode == OPPUSHDATA4 {
		if scriptLen-pc < 4 {
			return opCode, nil, pc, ErrNoEnoughDataOPPUSHDATA4
		}
		// 4 bytes after opcode encodes operand size
		operandSize = int(binary.LittleEndian.Uint16(script[pc : pc+4]))
		pc += 4
	}

	if scriptLen-pc < operandSize {
		return opCode, nil, pc, ErrScriptBound
	}
	// Read operand
	operand := Operand(script[pc : pc+operandSize])
	pc += operandSize
	return opCode, operand, pc, nil
}

// Execute an operation
func (s *Script) execOp(opCode OpCode, pushData Operand, tx *types.Transaction,
	txInIdx int, pc int, scriptPubKeyStart *int, stack *Stack) error {

	// Push value
	if opCode <= OPPUSHDATA4 {
		if opCode < OPPUSHDATA1 {
			logger.Debugf("push data len: %d, pc: %d", len(pushData), pc)
		} else {
			logger.Debugf("opcode: %s, push data len: %d, pc: %d", opCodeToName(opCode), len(pushData), pc)
		}
		stack.push(pushData)
		return nil
	} else if opCode <= OP16 && opCode != OPRESERVED {
		sn := scriptNum(opCode) - scriptNum(OP1) + 1
		logger.Debugf("opcode: %s, push data: %d, pc: %d", opCodeToName(opCode), sn, pc)
		stack.push(Operand(sn.Bytes()))
		return nil
	}

	logger.Debugf("opcode: %s, pc: %d", opCodeToName(opCode), pc)
	switch opCode {
	case OPDUP:
		if stack.size() < 1 {
			return ErrInvalidStackOperation
		}
		stack.push(stack.topN(1))

	case OPADD:
		fallthrough
	case OPSUB:
		if stack.size() < 2 {
			return ErrInvalidStackOperation
		}
		op1 := stack.topN(2)
		sn1, err := newScriptNum(op1)
		if err != nil {
			return err
		}
		op2 := stack.topN(1)
		sn2, err := newScriptNum(op2)
		if err != nil {
			return err
		}
		var sn scriptNum
		switch opCode {
		case OPADD:
			sn = sn1 + sn2
		case OPSUB:
			sn = sn1 - sn2
		default:
			return ErrBadOpcode
		}
		stack.pop()
		stack.pop()
		stack.push(sn.Bytes())

	case OPEQUAL:
		fallthrough
	case OPEQUALVERIFY:
		if stack.size() < 2 {
			return ErrInvalidStackOperation
		}
		op1 := stack.topN(2)
		op2 := stack.topN(1)
		// use bytes.Equal() instead of reflect.DeepEqual() for efficiency
		isEqual := bytes.Equal(op1, op2)
		stack.pop()
		stack.pop()
		if isEqual {
			stack.push(operandTrue)
		} else {
			stack.push(operandFalse)
		}
		if opCode == OPEQUALVERIFY {
			if isEqual {
				stack.pop()
			} else {
				return ErrScriptEqualVerify
			}
		}

	case OPHASH160:
		if stack.size() < 1 {
			return ErrInvalidStackOperation
		}
		hash160 := Operand(crypto.Hash160(stack.topN(1)))
		stack.pop()
		stack.push(hash160)

	case OPCODESEPARATOR:
		// scriptPubKey starts after the code separator; pc points to the next byte
		*scriptPubKeyStart = pc

	case OPCHECKSIG:
		fallthrough
	case OPCHECKSIGVERIFY:
		if stack.size() < 2 {
			return ErrInvalidStackOperation
		}
		signature := stack.topN(2)
		publicKey := stack.topN(1)

		// script consists of: scriptSig + OPCODESEPARATOR + scriptPubKey
		scriptPubKey := (*s)[*scriptPubKeyStart:]

		isVerified := verifySig(signature, publicKey, scriptPubKey, tx, txInIdx)

		stack.pop()
		stack.pop()
		if isVerified {
			stack.push(operandTrue)
		} else {
			stack.push(operandFalse)
		}
		if opCode == OPCHECKSIGVERIFY {
			if isVerified {
				stack.pop()
			} else {
				return ErrScriptSignatureVerifyFail
			}
		}

	default:
		return ErrBadOpcode
	}
	return nil
}

// verify if signature is right
// scriptPubKey is the locking script of the utxo tx input tx.Vin[txInIdx] references
func verifySig(sigStr []byte, publicKeyStr []byte, scriptPubKey []byte, tx *types.Transaction, txInIdx int) bool {
	sig, err := crypto.SigFromBytes(sigStr)
	if err != nil {
		logger.Debugf("Deserialize signature failed")
		return false
	}
	publicKey, err := crypto.PublicKeyFromBytes(publicKeyStr)
	if err != nil {
		logger.Debugf("Deserialize public key failed")
		return false
	}

	sigHash, err := CalcTxHashForSig(scriptPubKey, tx, txInIdx)
	if err != nil {
		logger.Debugf("Calculate signature hash failed")
		return false
	}

	return sig.VerifySignature(publicKey, sigHash)
}

// CalcTxHashForSig calculates the hash of a tx input, used for signature
func CalcTxHashForSig(scriptPubKey []byte, tx *types.Transaction, txInIdx int) (*crypto.HashType, error) {
	if txInIdx >= len(tx.Vin) {
		return nil, ErrInputIndexOutOfBound
	}

	// We do not want to change the original tx script sig, so make a copy
	oldScriptSigs := make([][]byte, 0, len(tx.Vin))

	for i, txIn := range tx.Vin {
		oldScriptSigs = append(oldScriptSigs, txIn.ScriptSig)

		if i != txInIdx {
			// Blank out other inputs' signatures
			txIn.ScriptSig = nil
		} else {
			// Replace scriptSig with referenced scriptPubKey
			txIn.ScriptSig = scriptPubKey
		}
	}

	// force to recompute hash instead of getting from cached hash since tx has changed
	sigHash, err := tx.CalcTxHash()

	// recover script sig
	for i, txIn := range tx.Vin {
		txIn.ScriptSig = oldScriptSigs[i]
	}
	return sigHash, err
}

// diaasm disassembles script in human readable format. If the script fails to parse, the returned string will
// contain the disassembled script up to the failure point, appended by the string '[Error: error info]'
func (s *Script) disasm() string {
	var str []string

	for pc := 0; pc < len(*s); {
		opCode, operand, newPc, err := s.parseNextOp(pc)
		if err != nil {
			str = append(str, "[Error: "+err.Error()+"]")
			return strings.Join(str, " ")
		}
		if operand != nil {
			str = append(str, hex.EncodeToString(operand))
		} else {
			str = append(str, opCodeToName(opCode))
		}
		pc = newPc
	}

	return strings.Join(str, " ")
}

// IsPayToScriptHash returns if the script is p2sh
func (s *Script) IsPayToScriptHash() bool {
	// OP_HASH160 <160-bit redeemp script hash> OP_EQUAL
	script := *s
	return len(script) == 23 && OpCode(script[0]) == OPHASH160 && script[1] == 20 && OpCode(script[22]) == OPEQUAL
}

// ExtractAddress returns address within the script
func (s *Script) ExtractAddress() (types.Address, error) {
	//TODO: only applies to p2pkh
	// p2pkh scriptPubKey: OPDUP OPHASH160 <pubKeyHash> OPEQUALVERIFY OPCHECKSIG
	_, _, newPc, _ := s.parseNextOp(0)
	_, _, newPc, _ = s.parseNextOp(newPc)
	// Third operand is pubKeyHash
	_, pubKeyHash, _, _ := s.parseNextOp(newPc)

	return types.NewAddressPubKeyHash(pubKeyHash)
}
