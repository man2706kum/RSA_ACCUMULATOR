package main

import (
	"fmt"
	"math/big"

	"./Acc"
	verify "./verification"
	witness "./witness"
)

func main() {

	//Generation of a Hidden group order
	key := Acc.Rsa_keygen(12)
	fmt.Println("Public Key:", &key)

	//Example set
	U := []big.Int{*big.NewInt(123), *big.NewInt(124), *big.NewInt(125), *big.NewInt(126)}

	//Generate the accumulator for th above set
	Accumulator := Acc.Generate_Acc(key, U)
	fmt.Println("Acc:", Accumulator)

	//witness of a member
	W1 := witness.Generate_witness(*big.NewInt(125), key, U)

	//Verification
	if verify.Verify(*big.NewInt(125), W1, Accumulator.Acc, key.N) {
		fmt.Printf("%v is a valid member\n", big.NewInt(125))
	} else {
		fmt.Printf("%v is not a member\n", big.NewInt(125))
	}

	//witness of a non-member
	W2 := witness.Generate_witness(*big.NewInt(127), key, U)
	if verify.Verify(*big.NewInt(15), W2, Accumulator.Acc, key.N) {
		fmt.Printf("%v is a valid member\n", big.NewInt(127))
	} else {
		fmt.Printf("%v is not a member\n", big.NewInt(127))
	}

	Accumulator.Add_member(*big.NewInt(127))
	//fmt.Println("Acc:", Accumulator)

	Accumulator.Delete_member(*big.NewInt(126))
	//fmt.Println("Acc", Accumulator)

	//pre computation of witness-------------------------------------
	list1 := make(map[string]big.Int, len(Accumulator.U))
	w := &witness.Witness_list{Acc: Accumulator.Acc, List: list1}

	w.Precompute_witness(Accumulator.G, Accumulator.U, Accumulator)
	fmt.Println("witness", w.List)
	//fmt.Println("Set:", Accumulator.U)
	W3 := witness.Generate_witness(*big.NewInt(123), key, Accumulator.U)
	W4 := witness.Generate_witness(*big.NewInt(124), key, Accumulator.U)
	W5 := witness.Generate_witness(*big.NewInt(125), key, Accumulator.U)
	W6 := witness.Generate_witness(*big.NewInt(127), key, Accumulator.U)
	fmt.Println("witnesses:", W3, W4, W5, W6)
}
