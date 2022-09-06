package main

import (
	"fmt"
	"strconv"
)

func FindCompound(compound string, it int) (string, int){
	subCompound := string(compound[it])
	it++
	for it < len(compound) {
		if compound[it] >= 'a' && compound[it] <= 'z' {
			subCompound = subCompound + string(compound[it])
		} else {
			break
		}
		it++
	}
	var num string
	for it < len(compound) {
		if compound[it] >= '0' && compound[it] <= '9' {
			num = num + string(compound[it])
		}else {
			break
		}
		it++
	}
	if num != "" {
		num, _ := strconv.Atoi(num)
		return subCompound, num
	}
	return subCompound, 1
}

func main() {
	compound := "CH3(HCl(OH)3)5(Cl)4"

	// Find braces and their multiplying factor
	braces := make(map[int]int) // map[braceNo]multiplyingFactor

	brackets := make([]int, 0)
	ans := make(map[string]int,0)
	var BracketCount = 0
	for i:= 0; i < len(compound) ; i++ {

		if compound[i] == '(' {
			BracketCount++
			brackets = append(brackets, BracketCount)
		}
		if compound[i] == ')' {
			b := brackets[len(brackets)-1]
			brackets = brackets[:len(brackets)-1]
			braces[b],_ = strconv.Atoi(string(compound[i+1]))
		}
	}

	multiplyingFactor := 1
	braceNum := 0
	lastAddedBracketCount := make([]int, 0)
	for i:= 0; i < len(compound) ; i++ {
		if compound[i] == '('{
			braceNum++
			lastAddedBracketCount = append(lastAddedBracketCount, braces[braceNum])
			multiplyingFactor = multiplyingFactor * braces[braceNum]
		}
		if compound[i] == ')' {
			divFactor := lastAddedBracketCount[len(lastAddedBracketCount)-1]
			lastAddedBracketCount = lastAddedBracketCount[:len(lastAddedBracketCount)-1]
			multiplyingFactor = multiplyingFactor/divFactor
		}

		if compound[i] >= 'A' && compound[i] <= 'Z' {
			subCompound, num := FindCompound(compound, i)
			ans[subCompound] =  ans[subCompound] + num*multiplyingFactor
		}
	}

	fmt.Println(ans)
}