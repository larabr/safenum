package main

// #include <stdlib.h>
// #include <stdint.h>
import "C"
import (
	"math/big"
	"runtime"

	"github.com/cronokirby/safenum"
)

// if you change these values, change them in dut_go.c aswell!
const chunksize = 16
const measurements = 1e4

// these are fixed throughout the tests
var (
	safeBase      *safenum.Nat
	safeModulus   *safenum.Modulus
	unsafeBase    *big.Int
	unsafeModulus *big.Int
)

//export init_dut
func init_dut() {
	safeBase = new(safenum.Nat).SetUint64(2)
	unsafeBase = big.NewInt(2)

	p, _ := new(big.Int).SetString(
		"FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD1"+
			"29024E088A67CC74020BBEA63B139B22514A08798E3404DD"+
			"EF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245"+
			"E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7ED"+
			"EE386BFB5A899FA5AE9F24117C4B1FE649286651ECE45B3D"+
			"C2007CB8A163BF0598DA48361C55D39A69163FA8FD24CF5F"+
			"83655D23DCA3AD961C62F356208552BB9ED529077096966D"+
			"670C354E4ABC9804F1746C08CA237327FFFFFFFFFFFFFFFF", 16)

	safeModulus = safenum.ModulusFromBytes(p.Bytes())
	unsafeModulus = p
}

//export do_one_computation
func do_one_computation(dataptr *C.uint8_t) C.uint8_t {

	// data is the random exponent
	// modulo and base are constant
	data := makeslice(dataptr, chunksize)
	// safeExp(data)
	unsafeExp(data)
	return 1
}

func safeExp(expBytes []byte) {
	exp := new(safenum.Nat)
	exp.SetBytes(expBytes)
	var res safenum.Nat
	res.Exp(safeBase, exp, safeModulus)
}

func unsafeExp(expBytes []byte) {
	exp := new(big.Int)
	exp.SetBytes(expBytes)
	new(big.Int).Exp(unsafeBase, exp, unsafeModulus)
}

//export prepare_inputs
func prepare_inputs(inputptr *C.uint8_t, classesptr *C.uint8_t) {

	// create slice abstractions
	allinputs := makeslice(inputptr, chunksize*measurements)
	classes := makeslice(classesptr, measurements)
	inputs := make([][]byte, measurements)
	for i := range inputs {
		inputs[i] = allinputs[i*chunksize : (i+1)*chunksize]
	}

	prepareData(&inputs, &classes)

	// run GC now in an attempt to not have
	// it run during computations too much
	runtime.GC()
}

// prepare byte slices .. some with zeroes, some with random data
func prepareData(data *[][]byte, classes *[]byte) {

	for i := range *data {
		(*classes)[i] = randombit()
		if (*classes)[i] == 1 {
			randombytes(&(*data)[i])
		} else {
			// zero exponent
		}
	}

}

func main() {}
