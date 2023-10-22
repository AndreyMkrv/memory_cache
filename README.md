# In memory cache

Simple implementation of in-memory-cache with using of standart data structures of go lang. 

## Example:

```sh
    func main() {
	cache := cache.New()

	john := NewUser("John", 29)
	sara := NewUser("Sara", 27)

	cache.Set("id-1", john)
	cache.Set("id-2", sara)

	fmt.Println(cache.Get("id-1"))
	fmt.Println(cache.Get("id-2"))

}
```