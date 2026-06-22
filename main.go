package main

import (
	"fmt"

	twopointerslowfast "leet-code-golang.com/two-pointer-slow-fast"
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
	// nums := []int{2, 7, 11, 15}

	// abc := array.Constructor(nums)
	// res := abc.SumRange(0, 2)
	// fmt.Println("res: ", res)
	// res := array.MaxSubArray(nums)

	// array.IsSubsequence()
	// array.LongestConsecutive()
	// array.FirstMissingPositive()
	// i := 0
	// for i < 5 { // This acts as a 'while' loop
	// 	fmt.Println(i)
	// 	i++
	// }
	// fmt.Printf("%q\n", chars)

	// res := map[int][]int{}
	// res[9] = []int{2, 7, 11, 15}
	// res[1] = []int{-3, -1, 0, 2, 4, 6}
	// res[13] = []int{1, 2, 3, 4, 4, 9}
	// res[8] = []int{1, 2, 3, 4, 4}

	// for k, v := range res {
	// 	res := twopointer.TwoSumV2(v, k)
	// 	fmt.Println("k: ", k, "value: ", v, "res: ", res)
	// }
	// res := twopointer.TwoSumV2(nums, 9)
	// fmt.Println("res: ", res)

	// except := []string{"A man, a plan, a canal: Panama", "race a car", " "}
	// for _, v := range except {
	// 	res := twopointer.IsPalindromeV2(v)
	// 	fmt.Println("s: ", v, "res: ", res)
	// }

	// abc := map[int][][]int{}
	// except := [][]int{
	// 	[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
	// 	[]int{1, 1, 2},
	// 	[]int{1, 1, 2},
	// 	[]int{1, 1, 2},
	// }
	// for _, v := range except {
	// 	res := twopointer.IsPalindromeV2(v)
	// 	fmt.Println("s: ", v, "res: ", res)
	// }

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

	// nums := []int{0, 1, 2, 2, 3, 3, 4}
	// twopointerslowfast.RemoveElementV2(nums, 2)
	// twopointerslowfast.MoveZeroes()
	// nums := []int{-4, -1, -1, 0, 1, 2}
	// abc := twopointerslowfast.ThreeSumV2(nums)
	// fmt.Println("res: ", abc)

	closest := twopointerslowfast.ThreeSumClosest([]int{-1, 2, 1, -4}, 1)
	fmt.Println("closet: ", closest)

}
