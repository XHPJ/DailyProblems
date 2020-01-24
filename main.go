package main

import "fmt"

func main() {

	//Problem 7
	// fmt.Println(NumOfPossibleDecodes("111"))
	// 	fmt.Println(NumOfPossibleDecodes("25710"))
	// 	fmt.Println(NumOfPossibleDecodes("2571"))

	//Problem 8
	// 	tree := BinaryTree{value: 0, left: nil, right: nil}
	// 	tree.SetLeft(1)
	// 	tree.SetRight(0)
	// 	tree.right.SetLeft(1)
	// 	tree.right.SetRight(0)
	// 	tree.right.left.SetLeft(1)
	// 	tree.right.left.SetRight(1)
	// 	fmt.Println(CountUivalTrees(&tree))

	//Problem 9
	// numbers := []int{2, 4, 6, 2, 5}
	// numbers2 := []int{5, 1, 1, 5}
	// fmt.Println(LargestSumOfNonAdj(numbers))
	// fmt.Println(LargestSumOfNonAdj(numbers2))

	//Problem 10
	//	FunctionScheduler(SaySomething, 4500)

	//Problem 11
	dictionary := []string{"dog", "deer", "deal"}
	predix := "de"
	fmt.Println(PreSelection(dictionary, predix))
}
