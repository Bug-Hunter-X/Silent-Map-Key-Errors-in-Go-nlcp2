func main() {
    var m = make(map[string]int)
    m["a"] = 1
    fmt.Println(m["b"]) // This will print 0, not an error
}