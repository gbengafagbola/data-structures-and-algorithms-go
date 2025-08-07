
# Character Compression in Go

This Go program compresses a slice of characters using a simple run-length encoding approach. It replaces sequences of the same character with the character followed by the count of repetitions (if more than one).

## Example

**Input:**

```go
[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
```

**Output (in-place):**

```go
[]byte{'a', '2', 'b', '2', 'c', '3'}
```

**Returned Length:**

```
6
```

## How It Works

The function iterates through the character slice and counts consecutive duplicates. Based on the count:

* If a character occurs once, it is added to the result as-is.
* If a character repeats 2–9 times, the count is converted to a single ASCII digit (`'2'` to `'9'`) and appended.
* If a character repeats 10 or more times, the count is converted to multiple ASCII digits (e.g., `'1'`, `'2'` for 12) using `fmt.Sprintf`.

Finally, the compressed result is written back into the original slice (`chars`), and the function returns the length of the compressed slice.

## Usage

```go
import "fmt"

func main() {
    chars := []byte{'a', 'a', 'a', 'b', 'b', 'c'}
    length := compress(chars)
    fmt.Println(chars[:length]) // Output: [a 3 b 2 c]
}
```

## Notes

* This implementation uses a secondary slice (`result`) for simplicity. It can be refactored for true in-place compression if needed.
* It avoids out-of-range panics by checking `i+1 < len(chars)` during iteration.

## Requirements

* Go 1.13+

## License

This project is open-source and free to use under the MIT License.
