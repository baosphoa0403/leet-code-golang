package main

import (
	"fmt"

	"leet-code-golang.com/array"
)

func main() {
	// res := array.TwoSumV2([]int{3, 2, 3, 7}, 9)
	// (3:0) (4:1) (3:0)
	// fmt.Println("res", res)
	// array.ContainDuplicate()
	// array.ValidAnagram()
	// a := array.IsAnagramV2()
	// fmt.Println("a", a)

	// res := array.GroupAnagrams([]string{
	// 	"eat", "tea", "tan", "ate", "nat", "bat",
	// })
	// fmt.Println("res", res)
	// res := array.TopKFrequentV2([]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2)
	// fmt.Println("res", res)
	// array.GroupAnagram()
	// array.TopKFrequenceElement()
	// abc := make(map[int]string, 0)
	// abc[1] = "hello"
	// for k, v := range abc {
	// 	fmt.Println(k, v)

	// }

	// a := []int{1, 1, 1}
	// fmt.Println("a: ", a[1:])
	// k := 2
	// res := array.SubarraySumV2(a, k)
	// fmt.Println("res: ", res)
	nums := []int{-2, 0, 3, -5, 2, -1}

	abc := array.Constructor(nums)
	res := abc.SumRange(0, 2)
	fmt.Println("res: ", res)

	// array.IsSubsequence()
	// array.LongestConsecutive()
	// array.FirstMissingPositive()
	// i := 0
	// for i < 5 { // This acts as a 'while' loop
	// 	fmt.Println(i)
	// 	i++
	// }
	// fmt.Printf("%q\n", chars)

	// twopointer.IsPalindrome()
	// twopointer.ReverseString()
	// twopointer.ReverseWords()
	// twopointer.TwoSum()

	// twopointerslowfast.RemoveDuplicates()
	// twopointerslowfast.RemoveElement()
	// twopointerslowfast.SortedSquares()
	// twopointerslowfast.ThreeSum()
	// twopointerslowfast.FourSum()

	// slidewindow.FindMaxAverage()
	// slidewindow.MaxProfit()
	// slidewindow.CharacterReplacement()

	// slidewindow.MinWindow()
	// linkedlist.ReverseList()
	// linkedlist.MiddleNode()
	// linkedlist.RemoveNthFromEnd()
	//linkedlist.MergeTwoLists()

	//linkedlist.DetectCycle()
	//linkedlist.IsPalindrome()
	// linkedlist.ReorderList()
	// linkedlist.ReverseBetween()
	// linkedlist.SortList()
	// linkedlist.InsertionSortList()
	// linkedlist.RemoveDuplicateFromSortList()
	// linkedlist.PartitionList()
	// linkedlist.SwapPairs()
	// linkedlist.AddTwoNumbers()
	// linkedlist.RotateRight()
	// stack.ValidParentheses()
	// stack.DailyTemperatures()

	// tree.MaxDepth()
	// tree.InvertTree()
	// tree.IsSameTree()
	// tree.IsSubtree()
	// tree.DiameterOfBinaryTree()
	// tree.IsBalanced()
	// tree.PreorderTraversal()
	// tree.InorderTraversal()
	// tree.PostorderTraversal()
	// tree.LevelOrder()

	// binarysearch
	// binarysearch.BinarySearch()
	// binarysearch.SearchInRotatedStoredArray()

}
