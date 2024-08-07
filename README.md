# LeetCode

LeetCode solutions in Go.

- **[1. Two Sum](https://leetcode.com/problems/two-sum/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/1_two_sum/two_sum.go)
    - **Difficulty:** Easy
    - **Category:** Array
    - **Notes:** use hash map to instantly check for difference value
    - **Tutorial:** [NeetCode](https://youtu.be/KLlXCFG5TnA)
      , [Kevin Naughton Jr.](https://www.youtube.com/watch?v=Aql6zHkONek), [Nick White
      ](https://www.youtube.com/watch?v=BoHO04xVeU0)
    - **Asked By:** Facebook, Google, Amazon, Bloomberg
- **[2. Add Two Numbers](https://leetcode.com/problems/add-two-numbers/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/2_add_two_numbers/main.go)
    - **Difficulty:** Medium
    - **Category:** Linked List
- **[3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/3_longest_substring_without_repeating_characters/main.go)
    - **Difficulty:** Medium
    - **Category:** String
- **[7. Reverse Integer](https://leetcode.com/problems/reverse-integer/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/7_reverse_integer/reverse_integer.go)
    - **Difficulty:** Easy
- **[9. Palindrome Number](https://leetcode.com/problems/palindrome-number/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/9_palindrome_number/palindrome_number.go)
    - **Difficulty:** Easy
- **[11. Container With Most Water](https://leetcode.com/problems/container-with-most-water/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/11_container_with_most_water/container_with_most_water.go)
    - **Difficulty:** Medium
    - **Category:** Array
    - **Tutorials** [NeetCode - Python](https://www.youtube.com/watch?v=UuiTKBwPgAo)
- **[13. Roman To Integer](https://leetcode.com/problems/roman-to-integer/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/13_roman_to_integer/roman_to_integer.go)
    - **Difficulty:** Easy
    - **Category:** String
- **[14. Longest Common Prefix](https://leetcode.com/problems/longest-common-prefix/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/14_longest_common_prefix/longest_common_prefix.go)
    - **Difficulty:** Easy
    - **Category:** String
- **[15. 3Sum](https://leetcode.com/problems/3sum/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/15_3sum/3_sum.go)
    - **Difficulty:** Medium
    - **Category:** Array, Two Pointers, Sorting
    - **Tutorial:** [NeetCode - Python](https://www.youtube.com/watch?v=jzZsG8n2R9A)
      , [Nick White - Java](https://www.youtube.com/watch?v=qJSPYnS35SE)
    - **Notes:** sort array, then for each elem find other 2 elems, like in two sum 2 problem; sort input, for each
      first element, find next two where -a = b+c, if a=prevA, skip a, if b=prevB skip b to elim duplicates; to find b,c
      use two pointers, left/right on remaining list;
- **[17. Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/17_letter_combinations_of_a_phone_number/letter_combinations_of_a_phone_number.go)
    - **Difficulty:** Medium
    - **Category:** String
- **[20. Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/20_valid_parentheses/valid_parentheses.go)
    - **Difficulty:** Easy
    - **Category:** String
    - **Tutorial:** [NeetCode - Python](https://www.youtube.com/watch?v=WTzjTskDFMg)
      , [Nick White - Java](https://www.youtube.com/watch?v=9kmUaXrjizQ)
    - **Note:** push opening brace on stack, pop if matching close brace, at end if stack empty, return true;
- **[21. Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/21_merge_two_sorted_lists/merge_two_sorted_lists.go)
    - **Difficulty:** Easy
    - **Category:** Linked List
    - **Notes:** insert each node from one list into the other
    - **Tutorial:** [NeetCode - Python](https://www.youtube.com/watch?v=XIdigk956u0)
    - **Asked By:** Microsoft
- **[33. Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/33_search_in_rotated_sorted_array/search_in_rotated_sorted_array.go)
    - **Difficulty:** Medium
    - **Category:** Array, Binary search
    - **Notes:** at most two sorted halfs, mid will be apart of left sorted or right sorted, if target is in range of
      sorted portion then search it, otherwise search other half
    - **Tutorial:** [NeetCode - Python](https://www.youtube.com/watch?v=U8XENwh8Oy8)
      , [Nick White - Java](https://www.youtube.com/watch?v=QdVrY3stDD4)
- **[35. Search Insert Position](https://leetcode.com/problems/search-insert-position/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/35_search_insert_position/search_insert_position.go)
    - **Difficulty:** Easy
    - **Category:** Array, Binary search
    - **Notes:** use binary search, to find mid val, check if its match return mid, else return left pointer.
    - **Tutorial:** [NeetCode - Python](https://www.youtube.com/watch?v=K-RYzDZkzCI)
- **[48. Rotate Image](https://leetcode.com/problems/rotate-image/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/48_rotate_image/rotate_image.go)
    - **Difficulty:** Medium
    - **Category:** String, Matrix
- **[49. Group Anagrams](https://leetcode.com/problems/group-anagrams/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/49_group_anagrams/group_anagrams.go)
    - **Difficulty:** Medium
    - **Category:** String, Sorting, Hash Table
    - **Tutorial:** [NeetCode - Python](https://www.youtube.com/watch?v=vzdNOK2oB2E)
      , [Kevin Naughton Jr. - Java](https://www.youtube.com/watch?v=ptgykfAEax8)
    - **Note:** create hash table - key will be a frequency in str, then compare and add to result; for each of 26
      chars, use count of each char in each word as tuple for key in dict, value is the list of anagrams;
    - **Asked By:** Amazonm Google, Uber, Facebook, Microsoft, Apple
- **[53. Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/53_maximum_subarray/maximum_subarray.go)
    - **Difficulty:** Easy
    - **Category:** Arrays
    - **Notes:** pattern: prev subarray can't be negative, dynamic programming: compute max sum for each prefix, sliding
      window
    - **Tutorial:** [NeetCode](https://www.youtube.com/watch?v=5WZl3MMT0Eg)
    - **Asked By:** Amazon
- **[55. Jump Game](https://leetcode.com/problems/maximum-subarray/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/55_jump_game/jump_game.go)
    - **Difficulty:** Medium
    - **Category:** Dynamic Programming, Arrays
    - **Notes:** visualize the recursive tree, cache solution for O(n) time/mem complexity, iterative is O(1) mem, just
      iterate backwards to see if element can reach goal node, if yes, then set it equal to goal node, continue;
    - **Tutorial:** [NeetCode - Python](https://www.youtube.com/watch?v=Yan0cv2cLy8)
      , [Nick White - Java](https://www.youtube.com/watch?v=Zb4eRjuPHbM)
- **[70. Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/70_climbing_stairs/climbing_stairs.go)
    - **Difficulty:** Easy
    - **Category:** Dynamic Programming, Caching - Memoization
    - **Tutorial:** [NeetCode](https://youtu.be/Y0lT9Fck7qI)
      , [Kevin Naughton Jr.](https://www.youtube.com/watch?v=uHAToNgAPaM)
    - **Notes:** subproblem find (n-1) and (n-2), sum = n;
    - **Asked By:** Google
- **[79. Word Search](https://leetcode.com/problems/word-search/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/79_word_search/79_word_search.go)
    - **Difficulty:** Medium
    - **Category:** Matrix
- **[91. Decode Ways](https://leetcode.com/problems/decode-ways/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/)
    - **Difficulty:** Medium
    - **Category:** String, Dynamic Programming
    - **Tutorial:** [NeetCode](https://www.youtube.com/watch?v=6aEyTjOwlJU)
      , [Kevin Naughton Jr.](https://www.youtube.com/watch?v=cQX3yHS0cLo)
    - **Asked By:** Facebook, Microsoft, Uber, Google, Twitter, Bloomberg
- **[94. Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/94_binary_tree_inorder_traversal/binary_tree_inorder_traversal.go)
    - **Difficulty:** Easy
    - **Category:** Tree
- **[98. Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/98_validate_binary_search_tree/validate_binary_search_tree.go)
    - **Difficulty:** Medium
    - **Category:** Tree
- **[100. Same Tree](https://leetcode.com/problems/same-tree/)**: [solution](100_same_tree/same_tree.go)
    - **Difficulty:** Easy
    - **Category:** Tree
    - **Notes:** dfs: recursively check left and right nodes of trees at the same time; iterative bfs compare each level
      of both trees
    - **Tutorial:** [NeetCode - Python](https://www.youtube.com/watch?v=vRbbcKXCxOw)
- **[101. Symmetric Tree](https://leetcode.com/problems/symmetric-tree/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/101_symmetric_tree/symmetric_tree.go)
    - **Difficulty:** Easy
    - **Category:** Tree
- **[104. Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/104_max_depth_of_binary_tree/max_depth_of_binary_tree.go)
    - **Difficulty:** Easy
    - **Category:** Tree
    - **Notes:** recursive dfs to find max-depth of subtrees
    - **Tutorial:** [NeetCode - Python](https://www.youtube.com/watch?v=hTM3phVI6YQ)
- **[108. Convert Sorted Array to Binary Search Tree](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/108_convert_sorted_array_to_binary_search_tree/convert_sorted_array_to_binary_search_tree.go)
    - **Difficulty:** Easy
    - **Category:** Array, Binary Tree
    - **Notes:**
    - **Tutorial:** [Nick White - Java](https://www.youtube.com/watch?v=12omz-VAyRk)
      , [Kevin Naughton Jr. - Java](https://www.youtube.com/watch?v=PZYTs9y4f4o)
- **[121. Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock)**: [solution](https://github.com/rorua/leetcode_go/blob/master/121_best_time_to_buy_sell_stock/best_time_to_buy_sell_stock.go)
    - **Difficulty:** Easy
    - **Category:** Arrays, Dynamic Programming
    - **Tutorial:** [NeetCode](https://youtu.be/1pkOgXD63yU)
    - **Notes:** find local min and search for local max, sliding window;
- **[125. Valid Palindrome](https://leetcode.com/problems/valid-palindrome/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/125_valid_palindrome/valid_palindrome.go)
    - **Difficulty:** Easy
    - **Category:** String, Two Pointers
    - **Tutorial:** [NeetCode](https://youtu.be/1pkOgXD63yU)
    - **Notes:** find local min and search for local max, sliding window;
- **[136. Single Number](https://leetcode.com/problems/single-number/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/136_single_number/single_number.go)
    - **Difficulty:** Easy
    - **Category:** Array
- **[141. Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/141_linked_list_cycle/linked_list_cycle.go)
    - **Difficulty:** Easy
    - **Category:** Linked List
    - **Tutorial:** [NeetCode - Python](https://www.youtube.com/watch?v=gBTe7lFR3vc)
      , [Nick White - Java](https://www.youtube.com/watch?v=6OrZ4wAy4uE)
    - **Notes:** 1 - dict to remember visited nodes; 2 - two pointers at different speeds, if they meet there is loop
- **[144. Binary Tree Preorder Traversal](https://leetcode.com/problems/binary-tree-preorder-traversal/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/144_binary_tree_preorder_traversal/binary_tree_preorder_traversal.go)
    - **Difficulty:** Easy
    - **Category:** Tree
- **[152. Maximum Product Subarray](https://leetcode.com/problems/maximum-product-subarray/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/152_maximum_product_subarray/maximum_product_subarray.go)
    - **Difficulty**: Medium
    - **Category:** Array, Dynamic Programming
    - **Tutorial:** [NeetCode - Python](https://www.youtube.com/watch?v=lXVy6YWFcRM)
      , [Xavier Elon - Java](https://www.youtube.com/watch?v=QQVCKkImH_s)
    - **Notes:** dp: compute max and max-abs-val for each prefix subarr;
- **[153. Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/153_find_min_in_rotated_sorted_array/find_min_in_rotated_sorted_array.go)
    - **Difficulty**: Medium
    - **Category:** Array, Binary
    - **Tutorial:** [NeetCode - Python](https://www.youtube.com/watch?v=nIVW4P8b1VA)
      , [Nick White - Java](https://www.youtube.com/watch?v=IzHR_U8Ly6c)
    - **Notes:** check if half of array is sorted in order to find pivot, arr is guaranteed to be in at most two sorted
      subarrays
- **[155. Min Stack](https://leetcode.com/problems/min-stack/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/155_min_stack/min_stack.go)
    - **Difficulty**: Easy
    - **Category:** Stack
- **[160. Intersection of Two Linked Lists](https://leetcode.com/problems/intersection-of-two-linked-lists/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/160_intersection_of_two_linked_lists/intersection_of_two_linked_lists.go)
    - **Difficulty:** Easy
    - **Category:** Linked List
    - **Notes:** use hash map of visited nodes, return first intercept node...
    - **Tutorial:** [Kevin Naughton Jr.](https://www.youtube.com/watch?v=CPXIkMWNn5Q)
      , [Nick White](https://www.youtube.com/watch?v=IpBfg9d4dmQ)
    - **Asked By:** Amazon, Microsoft, Bloomberg, Oracle, Yahoo
- **[167. Two Sum II - Input array is sorted](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/167_two_sum_2/two_sum_2.go)
    - **Difficulty:** Easy
    - **Category:** Array
- **[169. Majority Element](https://leetcode.com/problems/majority-element/)**
    - [not solved yet](https://github.com/rorua/leetcode_go/blob/master/169_majority_element/majority_element.go)
        - **Difficulty:** Easy
        - **Category:** Array
- **[189. Rotate Array](https://leetcode.com/problems/rotate-array/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/189_rotate_array/rotate_array.go)
    - **Difficulty:** Medium
    - **Category:** Array
- **[190. Reverse Bits](https://leetcode.com/problems/reverse-bits/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/190_reverse_bits/reverse_bits.go)
    - **Difficulty:** Easy
    - **Category:** Bit Manipulation
    - **Tutorial:** [NeetCode](https://www.youtube.com/watch?v=UcoN6UjAI64)
    - **Notes:** reverse each of 32 bits bit by bit
- **[191. Number of 1 Bits](https://leetcode.com/problems/number-of-1-bits/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/191_number_of_1_bits/number_of_1_bits.go)
    - **Difficulty:** Easy
    - **Category:** Bit Manipulation
    - **Notes:** Hamming weight. mod and div are expensive, to divide use bit shift, instead of mod to get 1's place use
      bitwise & 1
    - **Tutorial:** [NeetCode](https://www.youtube.com/watch?v=5Km3utixwZs)
- **[198. House Robber](https://leetcode.com/problems/house-robber/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/198_house_robber/house_robber.go)
    - **Difficulty:** Medium
    - **Category:** Dynamic Programming
    - **Notes:** for each num, get max of prev subarr, or num + prev subarr not including last element, store results of
      prev, and prev not including last element
    - **Tutorial:** [NeetCode - Python](https://www.youtube.com/watch?v=73r3KWiEvyk)
      , [Nick White - Java](https://www.youtube.com/watch?v=ZwDDLAeeBM0)
- **[200. Number of Islands](https://leetcode.com/problems/number-of-islands/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/200_number_of_islands/number_of_islands.go)
    - **Difficulty:** Medium
    - **Category:** Matrix, Array
- **[206. Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/206_reverse_linked_list/reverse_linked_list.go)
    - **Difficulty:** Easy
    - **Category:** Linked List
    - **Notes:** iterate through maintaining cur and prev; recursively reverse, return new head of list
    - **Tutorial:** [NeetCode - Python](https://www.youtube.com/watch?v=G0_I-ZF0S38)
      , [Nick White - Java](https://www.youtube.com/watch?v=NhapasNIKuQ)
- **[217. Contains Duplicate](https://leetcode.com/problems/contains-duplicate/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/217_contains_duplicate/contains_duplicate.go)
    - **Difficulty:** Easy
    - **Category:** Arrays, Hash Table
    - **Notes:** hashmap to get unique values in array, to check for duplicates easily
    - **Tutorial:** [NeetCode - Python](https://www.youtube.com/watch?v=3OamzN90kPg)
- **[226. Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/226_invert_binary_tree/invert_binary_tree.go)
    - **Difficulty:** Easy
    - **Category:** Tree
    - **Notes:** recursive dfs to invert subtrees; bfs to invert levels, use collections.deque; iterative dfs is easy
      with stack if doing pre-order traversal
    - **Tutorial:** [NeetCode - Python](https://www.youtube.com/watch?v=OnSn2XEQ4MY)
- **[234. Palindrome Linked List](https://leetcode.com/problems/palindrome-linked-list/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/234_palindrome_linked_list/palindrome_linked_list.go)
    - **Difficulty:** Easy
    - **Category:** Linked List
- **[235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/235_LCA_of_a_BST/LCA_of_a_BST.go)
    - **Difficulty:** Easy
    - **Category:** Binary Tree
    - **Notes:** compare p, q values to curr node, base case: one is in left, other in right subtree, then curr is lca;
    - **Tutorial:** [NeetCode - Python](https://www.youtube.com/watch?v=gs2LMfuOR9k)
      , [Nick White - Java](https://www.youtube.com/watch?v=6POfA8fruxI)
- **[238. Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/238_product_of_array_except_self/product_of_array_except_self.go)
    - **Difficulty:** Medium
    - **Category:** Array
    - **Tutorial:** [NeetCode - Python](https://www.youtube.com/watch?v=bNvIQI2wAjk)
      , [Nick White - Java](https://www.youtube.com/watch?v=tSRFtR3pv74)
- **[242. Valid Anagram](https://leetcode.com/problems/valid-anagram/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/242_valid_anagram/valid_anagram.go)
    - **Difficulty:** Easy
    - **Category:** String
    - **Tutorial:** [NeetCode - Python](https://www.youtube.com/watch?v=9UtInBqnCgA)
- **[268. Missing Number](https://leetcode.com/problems/missing-number/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/268_missing_number/missing_number.go)
    - **Difficulty:** Easy
    - **Category:** Arrays, Hash Table, Binary
    - **Notes:** 1st - using hashmap, and find absent; 2nd- using xor operator, to find absent; 3rd - calculate sums of
      array and sub it.
    - **Tutorial:** [NeetCode](https://www.youtube.com/watch?v=WnPLSRLSANE)
- **[283. Move Zeroes](https://leetcode.com/problems/move-zeroes/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/283_move_zeroes/move_zeroes.go)
    - **Difficulty:** Easy
    - **Category:** Array
- **[338. Counting Bits](https://leetcode.com/problems/counting-bits/submissions/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/338_counting_bits/counting_bits.go)
    - **Difficulty:** Easy
    - **Category:** Bit Manipulation, Dynamic Programming
    - **Tutorial:** [NeetCode - Python](https://youtu.be/RyBM56RIWrM)
- **[371. Sum of Two Integers](https://leetcode.com/problems/sum-of-two-integers/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/371_sum_of_two_integers/sum_of_two_integers.go)
    - **Difficulty:** Medium
    - **Category:** Math, Bit Manipulation
- **[383. Ransom Note](https://leetcode.com/problems/ransom-note/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/383_ransom_note/ransom_note.go)
    - **Difficulty:** Easy
    - **Category:** String
- **[387. First Unique Character in a String](https://leetcode.com/problems/first-unique-character-in-a-string/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/387_first_unique_character_in_a_string/first_unique_character_in_a_string.go)
    - **Difficulty:** Easy
    - **Category:** String
- **[404. Sum of Left Leaves](https://leetcode.com/problems/sum-of-left-leaves/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/404_sum_of_left_leaves/sum_of_left_leaves.go)
    - **Difficulty:** Easy
    - **Category:** Tree
- **[543. Diameter of Binary Tree](https://leetcode.com/problems/diameter-of-binary-tree/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/543_diameter_of_binary_tree/diameter_of_binary_tree.go)
    - **Difficulty:** Easy
    - **Category:** Tree
- **[617. Merge Two Binary Trees](https://leetcode.com/problems/merge-two-binary-trees/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/617_merge_two_binary_trees/merge_two_binary_trees.go)
    - **Difficulty:** Easy
    - **Category:** Tree
- **[704. Binary Search](https://leetcode.com/problems/binary-search/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/704_binary_search/binary_search.go)
    - **Difficulty:** Easy
    - **Category:** Tree, Array
- **[724. Find Pivot Index](https://leetcode.com/problems/find-pivot-index/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/724_find_pivot_index/find_pivot_index.go)
    - **Difficulty:** Easy
    - **Category:** Array
- **[859. Buddy Strings](https://leetcode.com/problems/buddy-strings/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/859_buddy_strings/buddy_strings.go)
    - **Difficulty:** Easy
    - **Category:** String
- **[912. Sort an Array](https://leetcode.com/problems/sort-an-array/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/912_sort_an_array/sort_an_array.go)
    - **Difficulty:** Medium
    - **Category:** Array
- **[1089. Duplicate Zeros](https://leetcode.com/problems/duplicate-zeros/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/1089_duplicate_zeros/duplicate_zeros.go)
    - **Difficulty:** Easy
    - **Category:** Array
- **[1108. Defanging an IP Address](https://leetcode.com/problems/defanging-an-ip-address/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/1108_defanging_an_ip_address/defanging_an_ip_address.go)
    - **Difficulty:** Easy
    - **Category:** String
- **[1394. Find Lucky Integer in an Array](https://leetcode.com/problems/find-lucky-integer-in-an-array/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/1394_find_lucky_integer/find_lucky_integer.go)
    - **Difficulty:** Easy
    - **Category:** Array
- **[1480. Running Sum of 1d Array](https://leetcode.com/problems/running-sum-of-1d-array/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/1480_running_sum_of_1d_array/running_sum_of_1d_array.go)
    - **Difficulty:** Easy
    - **Category:** Array
- **[1636. Sort Array by Increasing Frequency](https://leetcode.com/problems/sort-array-by-increasing-frequency/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/1636_sort_array_by_increasing_frequency/sort_array_by_increasing_frequency.go)
    - **Difficulty:** Easy
    - **Category:** Array, Hash Table
- **[1790. Check if One String Swap Can Make Strings Equal](https://leetcode.com/problems/check-if-one-string-swap-can-make-strings-equal/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/1790_check_if_one_string_swap_can_make_strings_equal/check_if_one_string_swap_can_make_strings_equal.go)
    - **Difficulty:** Easy
    - **Category:** String
- **[1925. Count Square Sum Triples](https://leetcode.com/problems/count-square-sum-triples/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/1925_count_square_sum_triples/count_square_sum_triples.go)
    - **Difficulty:** Easy
- **[1929. Concatenation of Array](https://leetcode.com/problems/concatenation-of-array/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/1929_concatenation_of_array/concatenation_of_array.go)
    - **Difficulty:** Easy
    - **Category:** Array
- **[2206. Divide Array Into Equal Pairs](https://leetcode.com/problems/divide-array-into-equal-pairs/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/2206_divide_array_into_equal_pairs/divide_array_into_equal_pairs.go)
    - **Difficulty:** Easy
    - **Category:** Array, Hash Table
- **[2418. Sort the People](https://leetcode.com/problems/sort-the-people/)**: [solution](https://github.com/rorua/leetcode_go/blob/master/2418_sort_the_people/sort_the_people.go)
    - **Difficulty:** Easy
    - **Category:** Array, Hash Table