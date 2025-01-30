### Encode-Decode Algorithm

**Problem**  
Design an algorithm to encode a list of strings to a single string. The encoded string is then decoded back to the original list of strings.

**Description**  
This project demonstrates a custom encoding and decoding algorithm for efficiently converting a list of strings into a single string and then back to the original list. It ensures that the transformation is lossless and can handle special cases like empty strings or strings containing colons.

---

## Functions
- **Encode**: Transforms a list of strings into a single encoded string.
- **Decode**: Reconstructs the original list of strings from the encoded string.
- Handles edge cases like empty strings or special characters (`:`).
- Efficient implementation using a prefix-based encoding scheme or JSON serialization.

---

## Implementation Details

### Encoding Algorithm

The algorithm prefixes each string with its length followed by a colon (`:`). The format ensures that decoding is unambiguous.

**Example**:
- Input: `["neet", "code", "love", "you"]`
- Encoded String: `"4:neet4:code4:love3:you"`

### Decoding Algorithm

The decoding algorithm:
1. Parses the length of each string (up to the first colon `:`).
2. Extracts the corresponding substring using the parsed length.
3. Repeats the process until the entire encoded string is decoded.

**Example**:
- Encoded String: `"4:neet4:code4:love3:you"`
- Output: `["neet", "code", "love", "you"]`

- **Time Complexity**:  
- - encode: O(n . L)
- - decode:O(n . L)
- - Overall: O(n . L)
- -
- **Space Complexity**:
- - Overall: O(n . L)

[source](https://leetcode.com/problems/encode-and-decode-strings/)
--- 