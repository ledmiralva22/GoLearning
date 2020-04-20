package main

import "fmt"

const GlobalLimit = 100
const MaxCacheSize int = 10 * GlobalLimit

const (
	CacheKeyBook = "book_"
	CacheKeyCD   = "cd_"
)

var cache map[string]string

func cacheGet(key string) string {
	return cache[key]
}

func cacheSet(key, val string) {
	if len(cache)+1 >= MaxCacheSize {
		return
	}
	cache[key] = val
}

func GetBook(isbn string) string {
	return cacheGet(CacheKeyBook + isbn)
}

func SetBook(isbn string, name string) {
	cacheSet(CacheKeyBook+isbn, name)
}

func GetCD(sky string) string {
	return cacheGet(CacheKeyCD + sky)
}

func SetCD(sku string, title string) {
	cacheSet(CacheKeyCD+sku, title)
}

//iota and constant expression
//iota always begin in zero
const (
	first  = iota + 6
	second = iota << 2 // iota multiplied by 2, 2 times
)

//iota resets when is used in another constant block
const (
	third = iota
)

func main() {
	cache = make(map[string]string)
	SetBook("1234-5678", "Get Ready to Go")
	SetCD("1234-5678", "Get Ready to Go Audio Book")

	fmt.Println("Book: ", GetBook("1234-5678"))
	fmt.Println("CD: ", GetCD("1234-5678"))

	fmt.Println("Iota values")
	fmt.Println(first, second, third)
}
