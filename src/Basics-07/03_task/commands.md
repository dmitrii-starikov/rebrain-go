```text
go@linux:~/GolandProjects/rebrain-go/src/Basics-07/01_task$ go test -v ./internal/pkg/util -run TestContainsDuplicate
=== RUN   TestContainsDuplicate
=== PAUSE TestContainsDuplicate
=== CONT  TestContainsDuplicate
=== RUN   TestContainsDuplicate/empty_slice
=== PAUSE TestContainsDuplicate/empty_slice
=== RUN   TestContainsDuplicate/single_element
=== PAUSE TestContainsDuplicate/single_element
=== RUN   TestContainsDuplicate/no_duplicates
=== PAUSE TestContainsDuplicate/no_duplicates
=== RUN   TestContainsDuplicate/has_duplicates_at_beginning
=== PAUSE TestContainsDuplicate/has_duplicates_at_beginning
=== RUN   TestContainsDuplicate/has_duplicates_at_end
=== PAUSE TestContainsDuplicate/has_duplicates_at_end
=== RUN   TestContainsDuplicate/has_duplicates_in_middle
=== PAUSE TestContainsDuplicate/has_duplicates_in_middle
=== RUN   TestContainsDuplicate/multiple_duplicates
=== PAUSE TestContainsDuplicate/multiple_duplicates
=== RUN   TestContainsDuplicate/all_same_elements
=== PAUSE TestContainsDuplicate/all_same_elements
=== RUN   TestContainsDuplicate/negative_numbers_without_duplicates
=== PAUSE TestContainsDuplicate/negative_numbers_without_duplicates
=== RUN   TestContainsDuplicate/negative_numbers_with_duplicates
=== PAUSE TestContainsDuplicate/negative_numbers_with_duplicates
=== RUN   TestContainsDuplicate/zero_with_duplicates
=== PAUSE TestContainsDuplicate/zero_with_duplicates
=== RUN   TestContainsDuplicate/zero_without_duplicates
=== PAUSE TestContainsDuplicate/zero_without_duplicates
=== CONT  TestContainsDuplicate/empty_slice
=== CONT  TestContainsDuplicate/multiple_duplicates
=== CONT  TestContainsDuplicate/negative_numbers_with_duplicates
=== CONT  TestContainsDuplicate/single_element
=== CONT  TestContainsDuplicate/zero_without_duplicates
=== CONT  TestContainsDuplicate/zero_with_duplicates
=== CONT  TestContainsDuplicate/has_duplicates_in_middle
=== CONT  TestContainsDuplicate/has_duplicates_at_end
=== CONT  TestContainsDuplicate/all_same_elements
=== CONT  TestContainsDuplicate/negative_numbers_without_duplicates
=== CONT  TestContainsDuplicate/no_duplicates
=== CONT  TestContainsDuplicate/has_duplicates_at_beginning
--- PASS: TestContainsDuplicate (0.00s)
    --- PASS: TestContainsDuplicate/empty_slice (0.00s)
    --- PASS: TestContainsDuplicate/negative_numbers_with_duplicates (0.00s)
    --- PASS: TestContainsDuplicate/zero_with_duplicates (0.00s)
    --- PASS: TestContainsDuplicate/has_duplicates_in_middle (0.00s)
    --- PASS: TestContainsDuplicate/multiple_duplicates (0.00s)
    --- PASS: TestContainsDuplicate/has_duplicates_at_end (0.00s)
    --- PASS: TestContainsDuplicate/all_same_elements (0.00s)
    --- PASS: TestContainsDuplicate/negative_numbers_without_duplicates (0.00s)
    --- PASS: TestContainsDuplicate/zero_without_duplicates (0.00s)
    --- PASS: TestContainsDuplicate/no_duplicates (0.00s)
    --- PASS: TestContainsDuplicate/single_element (0.00s)
    --- PASS: TestContainsDuplicate/has_duplicates_at_beginning (0.00s)
PASS
ok      module06/internal/pkg/util      (cached)
```

```text
go@linux:~/GolandProjects/rebrain-go/src/Basics-07/01_task$ go test -v ./internal/pkg/util -run TestIsPalindrome
=== RUN   TestIsPalindrome
=== PAUSE TestIsPalindrome
=== CONT  TestIsPalindrome
=== RUN   TestIsPalindrome/positive_palindrome_single_digit
=== PAUSE TestIsPalindrome/positive_palindrome_single_digit
=== RUN   TestIsPalindrome/positive_palindrome_two_same_digits
=== PAUSE TestIsPalindrome/positive_palindrome_two_same_digits
=== RUN   TestIsPalindrome/positive_palindrome_three_digits
=== PAUSE TestIsPalindrome/positive_palindrome_three_digits
=== RUN   TestIsPalindrome/positive_palindrome_four_digits
=== PAUSE TestIsPalindrome/positive_palindrome_four_digits
=== RUN   TestIsPalindrome/positive_palindrome_five_digits
=== PAUSE TestIsPalindrome/positive_palindrome_five_digits
=== RUN   TestIsPalindrome/positive_palindrome_even_digits
=== PAUSE TestIsPalindrome/positive_palindrome_even_digits
=== RUN   TestIsPalindrome/positive_palindrome_with_zero
=== PAUSE TestIsPalindrome/positive_palindrome_with_zero
=== RUN   TestIsPalindrome/positive_palindrome_1001
=== PAUSE TestIsPalindrome/positive_palindrome_1001
=== RUN   TestIsPalindrome/positive_non-palindrome
=== PAUSE TestIsPalindrome/positive_non-palindrome
=== RUN   TestIsPalindrome/positive_non-palindrome_two_digits
=== PAUSE TestIsPalindrome/positive_non-palindrome_two_digits
=== RUN   TestIsPalindrome/positive_non-palindrome_four_digits
=== PAUSE TestIsPalindrome/positive_non-palindrome_four_digits
=== RUN   TestIsPalindrome/positive_non-palindrome_with_zero
=== PAUSE TestIsPalindrome/positive_non-palindrome_with_zero
=== RUN   TestIsPalindrome/zero_is_palindrome
=== PAUSE TestIsPalindrome/zero_is_palindrome
=== RUN   TestIsPalindrome/negative_number_always_false
=== PAUSE TestIsPalindrome/negative_number_always_false
=== RUN   TestIsPalindrome/negative_single_digit
=== PAUSE TestIsPalindrome/negative_single_digit
=== RUN   TestIsPalindrome/negative_palindrome-like
=== PAUSE TestIsPalindrome/negative_palindrome-like
=== RUN   TestIsPalindrome/number_ending_with_zero_but_not_zero
=== PAUSE TestIsPalindrome/number_ending_with_zero_but_not_zero
=== RUN   TestIsPalindrome/number_ending_with_zero_100
=== PAUSE TestIsPalindrome/number_ending_with_zero_100
=== RUN   TestIsPalindrome/number_ending_with_zero_1010
=== PAUSE TestIsPalindrome/number_ending_with_zero_1010
=== RUN   TestIsPalindrome/large_palindrome
=== PAUSE TestIsPalindrome/large_palindrome
=== RUN   TestIsPalindrome/large_non-palindrome
=== PAUSE TestIsPalindrome/large_non-palindrome
=== RUN   TestIsPalindrome/max_int_palindrome
=== PAUSE TestIsPalindrome/max_int_palindrome
=== CONT  TestIsPalindrome/positive_palindrome_single_digit
=== CONT  TestIsPalindrome/positive_non-palindrome_with_zero
=== CONT  TestIsPalindrome/positive_non-palindrome_four_digits
=== CONT  TestIsPalindrome/positive_palindrome_even_digits
=== CONT  TestIsPalindrome/positive_palindrome_five_digits
=== CONT  TestIsPalindrome/positive_palindrome_four_digits
=== CONT  TestIsPalindrome/positive_palindrome_three_digits
=== CONT  TestIsPalindrome/positive_palindrome_two_same_digits
=== CONT  TestIsPalindrome/positive_non-palindrome_two_digits
=== CONT  TestIsPalindrome/number_ending_with_zero_100
=== CONT  TestIsPalindrome/max_int_palindrome
=== CONT  TestIsPalindrome/number_ending_with_zero_but_not_zero
=== CONT  TestIsPalindrome/large_non-palindrome
=== CONT  TestIsPalindrome/large_palindrome
=== CONT  TestIsPalindrome/negative_palindrome-like
=== CONT  TestIsPalindrome/number_ending_with_zero_1010
=== CONT  TestIsPalindrome/negative_single_digit
=== CONT  TestIsPalindrome/positive_non-palindrome
=== CONT  TestIsPalindrome/positive_palindrome_1001
=== CONT  TestIsPalindrome/negative_number_always_false
=== CONT  TestIsPalindrome/zero_is_palindrome
=== CONT  TestIsPalindrome/positive_palindrome_with_zero
--- PASS: TestIsPalindrome (0.00s)
    --- PASS: TestIsPalindrome/positive_palindrome_single_digit (0.00s)
    --- PASS: TestIsPalindrome/positive_non-palindrome_four_digits (0.00s)
    --- PASS: TestIsPalindrome/positive_palindrome_three_digits (0.00s)
    --- PASS: TestIsPalindrome/positive_palindrome_two_same_digits (0.00s)
    --- PASS: TestIsPalindrome/positive_non-palindrome_two_digits (0.00s)
    --- PASS: TestIsPalindrome/positive_non-palindrome_with_zero (0.00s)
    --- PASS: TestIsPalindrome/positive_palindrome_even_digits (0.00s)
    --- PASS: TestIsPalindrome/number_ending_with_zero_100 (0.00s)
    --- PASS: TestIsPalindrome/max_int_palindrome (0.00s)
    --- PASS: TestIsPalindrome/number_ending_with_zero_but_not_zero (0.00s)
    --- PASS: TestIsPalindrome/large_non-palindrome (0.00s)
    --- PASS: TestIsPalindrome/large_palindrome (0.00s)
    --- PASS: TestIsPalindrome/negative_palindrome-like (0.00s)
    --- PASS: TestIsPalindrome/number_ending_with_zero_1010 (0.00s)
    --- PASS: TestIsPalindrome/negative_single_digit (0.00s)
    --- PASS: TestIsPalindrome/positive_non-palindrome (0.00s)
    --- PASS: TestIsPalindrome/positive_palindrome_five_digits (0.00s)
    --- PASS: TestIsPalindrome/positive_palindrome_four_digits (0.00s)
    --- PASS: TestIsPalindrome/positive_palindrome_1001 (0.00s)
    --- PASS: TestIsPalindrome/negative_number_always_false (0.00s)
    --- PASS: TestIsPalindrome/zero_is_palindrome (0.00s)
    --- PASS: TestIsPalindrome/positive_palindrome_with_zero (0.00s)
PASS
ok      module06/internal/pkg/util      (cached)
```