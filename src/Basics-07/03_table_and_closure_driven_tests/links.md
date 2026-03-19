* [Closure driven tests: an alternative style to table driven tests in go](https://medium.com/@cep21/closure-driven-tests-an-alternative-style-to-table-driven-tests-in-go-628a41497e5e)
* [Table driver tests golang](https://github.com/golang/go/wiki/TableDrivenTests)
* [Advanced tests go](https://about.sourcegraph.com/go/advanced-testing-in-go)

## Key Differences Between Table Driven and Closure Driven Tests

| Aspect                      | Table Driven                                       | Closure Driven                                          |
|-----------------------------|----------------------------------------------------|---------------------------------------------------------|
| Structure                   | Slice/array of structs with test data              | Constructor function returns func(t *testing.T)         |
| When to Use                 | Same test logic, only input/output data changes	   | Test logic varies between cases or needs specific setup |
| Readability                 | All test cases in one place, easy to add new ones	 | Each case is explicitly visible in t.Run() calls        |
| Flexibility                 | Limited by the test case struct definition         | Can have different preparation logic per test case      |
| Setup Complexity            | Simple, uniform setup for all cases                | Can have complex, unique setup per case                 |
| Code Duplication Complexity | Minimal - test logic written once                  | Slightly more verbose but very explicit                 |
