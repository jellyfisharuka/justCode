package main
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    }
    if list2 == nil {
        return list1
    }
    if list1.Val <= list2.Val {
        list1.Next = mergeTwoLists(list1.Next, list2)
        return list1
    }
    list2.Next = mergeTwoLists(list1, list2.Next)
    return list2
}
func main() {
	int list1 = [1,2,4], list2 = [1,3,4]
	int listT = mergeTwoLists(list1, list2);
	fmt.Print(listT)
}