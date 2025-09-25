# Strings in Go

## String Fundamentals

Strings in Go are immutable. Example:
```go
name := "aman"
name = "n" + name  // Go runtime creates new string "naman", old "aman" gets garbage collected
```

Go stores all strings as a slice of bytes (UTF-8 encoded) in memory.
- `byte` = `uint8` in Go
- Each character can be 1-4 bytes depending on UTF-8 encoding
- ASCII characters = 1 byte, Unicode can be up to 4 bytes

## Performance Implications

- **String concatenation with `+`** is O(n) - creates new string each time
- **For multiple concatenations**, use `strings.Builder` for O(1) amortized appends
- **String comparison** is O(n) in worst case - compares byte by byte
- **Memory**: Each concatenation allocates new memory, old string becomes garbage

## Common Operations & Time Complexity

- **Access character**: `s[i]` - O(1) (returns byte, not character)
- **Length**: `len(s)` - O(1) (counts bytes, not characters)
- **Substring**: `s[i:j]` - O(j-i) (creates new string)
- **Search**: `strings.Contains()` - O(n*m) where n=string length, m=pattern length
- **Split**: `strings.Split()` - O(n)
- **Replace**: `strings.Replace()` - O(n)

## Key Patterns for DSA Problems

### 1. Two-Pointer Technique
- Palindrome checking
- String reversal
- Finding pairs/triplets

### 2. Sliding Window
- Longest substring without repeating characters
- Minimum window substring
- Anagram detection in substrings

### 3. Character Frequency Counting
```go
freq := make(map[rune]int)
for _, char := range s {
    freq[char]++
}
```

### 4. ASCII Manipulation
- Convert to array index: `int(char) - int('a')` for lowercase a-z
- Character to number: `int(char) - int('0')` for digits 0-9
- Case conversion: `int(char) - int('a') + int('A')` for upper/lower

## Important Gotchas

### Bytes vs Runes
```go
s := "hello"
s[0]  // Returns byte (uint8), not character

// For Unicode characters, use range:
for i, char := range s {
    // char is rune (int32), i is byte index
}
```

### Length vs Character Count
```go
s := "hello"     // len(s) = 5 (5 bytes)
s := "हैलो"       // len(s) = 12 (12 bytes, 4 characters)

// For character count:
utf8.RuneCountInString(s)  // Returns actual character count
```

### String to Slice Conversion
```go
[]byte(s)    // Creates copy - expensive O(n) operation
[]rune(s)    // Creates copy - expensive O(n) operation
```

## Efficient String Building
```go
// Inefficient - O(n²) for n concatenations
result := ""
for _, str := range strings {
    result += str  // Creates new string each time
}

// Efficient - O(n) amortized
var builder strings.Builder
builder.Grow(estimatedSize)  // Pre-allocate if size known
for _, str := range strings {
    builder.WriteString(str)  // O(1) amortized
}
result := builder.String()
```


## Memory Layout
```
String: "hello"
┌─────────────┬─────────────┐
│   Pointer   │   Length    │  String header (16 bytes on 64-bit)
│  (8 bytes)  │  (8 bytes)  │
└─────────────┴─────────────┘
       │
       ▼
┌───┬───┬───┬───┬───┐
│'h'│'e'│'l'│'l'│'o'│  Underlying byte array (5 bytes)
└───┴───┴───┴───┴───┘
```

## Interview Tips
1. Always consider if you need bytes or runes
2. For ASCII-only problems, byte operations are faster
3. Remember string immutability when analyzing space complexity
4. Use `strings.Builder` for efficient string construction
5. Two-pointer technique works great for palindromes and reversals