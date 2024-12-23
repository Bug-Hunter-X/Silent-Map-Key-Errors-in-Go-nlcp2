func main() {
    var m = make(map[string]int)
    m["a"] = 1
    value, ok := m["b"]
    if ok {
        fmt.Println(value) // Access the value if the key exists
    } else {
        fmt.Println("Key not found") // Handle the case where the key doesn't exist
    }
}