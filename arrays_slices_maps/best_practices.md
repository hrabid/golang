In Go, there are three main ways to declare a slice, and the best practice depends on your use case (e.g., whether you know the initial elements, capacity, or want to avoid allocations). Let's go through each method and the best practices:


---

✅ 1. Using a Composite Literal (Recommended when you know the initial elements)

s := []int{1, 2, 3}

Best practice: Use this when you know the initial values.

Pros:

Clear and concise.

Allocates and initializes in one line.


Example use case: Static or hardcoded list of values.



---

✅ 2. Using make() (When you want to control length and capacity)

s := make([]int, 5)         // length 5, capacity 5
s := make([]int, 3, 10)     // length 3, capacity 10

Best practice: Use this when you need to preallocate memory for performance, especially in loops or appending data.

Pros:

Prevents repeated memory allocations during growth.

Good for performance in known-size use cases.


Example use case: You are building a slice via append and want to avoid reallocation.



---

✅ 3. Declaring a nil slice (When you don’t need to initialize yet)

var s []int

Best practice: Use this when you want a zero-value slice, and will initialize later or conditionally.

Pros:

Cleanest zero-value declaration.

len(s) == 0 and s == nil.


Example use case: Placeholder for function result or to return an empty slice.



---

❗️ What to avoid:

🚫 Using new([]int)

p := new([]int)   // p is a pointer to a slice

Rarely needed.

Avoid unless you explicitly want a pointer to a slice (uncommon).



---

🧠 Summary of Best Practices:

Method  When to Use

[]T{...}  You know the initial values
make([]T, len, cap) You want to control allocation/performance
var s []T You want a nil slice to initialize later


