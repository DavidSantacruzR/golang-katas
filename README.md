<h1>Coding Katas for Golang</h1>
<h2>Exercises</h2>

<h3>Basic:</h3>

<ul>
<li>Create a function that returns a Fibonacci sequence up to N.</li>
<li>Create a function that checks if a word is a palindrome</li>
<li>Create a function that encrypt/decrypt using Caesar algorithm</li>
</ul>

<h3>Medium:</h3>
<ul>
<li>Build your own redis server with SET, GET commands.</li>
</ul>

Include test cases for each of the exercises.

Write an exportable go function.
```go
package MyPackage

func MyFunction() {
	/*Logic stuff here.*/
}
```

Create a test case file with the exercise name ``mypackage_test.go``

Write tests for it.

```go
package MyPackage

import "testing"

func TestMyFunction(t *testing.T) {
	result := ""
	expected := ""
	if result != expected {
		t.Errorf("YourMessageError: got %v; want %v", result, expected)
    }
}
```
Run all your package tests using: ```go test ./MyPackage -v```
