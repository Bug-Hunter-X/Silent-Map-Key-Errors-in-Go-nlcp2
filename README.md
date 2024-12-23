# Silent Map Key Errors in Go

This repository demonstrates a subtle error in Go related to accessing non-existent keys in maps.  Go maps don't throw an error when accessing a key that doesn't exist; instead, they return the zero value for the map's value type. This can lead to unexpected behavior in your programs if you're not careful.  The example provided shows how this can occur and how to mitigate the issue.

## Example

The `bug.go` file shows a simple example that silently returns 0 when accessing the key "b", which doesn't exist in the map.

## Solution

The `bugSolution.go` file demonstrates a more robust approach.  It uses the comma ok idiom to check if a key exists before accessing its value, preventing unexpected results.