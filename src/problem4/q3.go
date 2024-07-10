/*Task
Provide 3 unique implementations of the following function in JavaScript.
Input: n - any integer
Assuming this input will always produce a result lesser than Number.MAX_SAFE_INTEGER.
Output: return - summation to n, i.e. sum_to_n(5) === 1 + 2 + 3 + 4 + 5 === 15.
*/

//using loop
//O(n) time complexity
func sum_to_n_b(n int) int {
    sum := 0
    for i := 1; i <= n; i++ {
        sum += i
    }
    return sum
}


//using recursion
//O(n) time complexity
var sum_to_n_b = function(n) {
    if(n===1) {
        return 1;
    }
    return n+sum_to_n_b(n-1);
};

//using formula which is n*(n+1)/2
//O(1) time complexity
func sum_to_n_a(n int) int {
    return n * (n + 1) / 2
}
