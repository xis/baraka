
<div align="center">
  <h1>baraka</h1>
  
[![Go Report Card](https://goreportcard.com/badge/github.com/xis/baraka)](https://goreportcard.com/report/github.com/xis/baraka)
[![codecov](https://codecov.io/gh/xis/baraka/branch/master/graph/badge.svg)](https://codecov.io/gh/xis/baraka)
[![Build Status](https://travis-ci.org/xis/baraka.svg?branch=master)](https://travis-ci.org/xis/baraka) 
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/xis/baraka)
  
a tool for handling file uploads for http servers

makes it easier to save multipart files from http request and to filter them,
prevents unwanted files from getting into memory, extracts json data with files.
</div>

# **install**
```bash
go get -u github.com/xis/baraka
```

# **using**
```go
func main() {
	// create a storage
	storage, err := baraka.NewStorage("./pics/", baraka.Options{})
	if err != nil {
		fmt.Println(err)
	}
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		// parsing
		p, err := storage.Parse(c.Request)
		// or you can use ParseButMax if you need limited size 
		p, err := storage.ParseButMax(32<<20, 5, c.Request)
		if err != nil {
			fmt.Println(err)
		}
		// saving
		p.Store("file_prefix")
	})
	router.Run()
}
```
you can use with other http server libraries, just pass the http.Request to storage.Parse function.

# **filter function**
filter function is a custom function that filters the files that comes from requests. you can read file bytes and identify the file, return true if you wanna pass the file, return false if you dont. 


## filter example
```go
// create a storage
func main() {
	storage, err := baraka.NewStorage("./pics/", baraka.Options{
		// passing filter function
		Filter: func(file *multipart.Part) bool {
			// create a byte array
			b := make([]byte, 512)
			// get the file bytes to created byte array
			file.Read(b)
			// detect the content type
			fileType := http.DetectContentType(b)
			// if it is jpeg then pass the file
			if fileType == "image/jpeg" {
				return true
			}
			// if not then don't pass
			return false
		},
	})
	...codes below...
```
# getting information
```go
... codes above ...
	p, err := storage.Parse(c.Request)
	if err != nil {
		fmt.Println(err)
	}
	// prints filenames
	fmt.Println(p.Filenames())
	// prints total files count
	fmt.Println(p.Length())
	// prints content types of files
	fmt.Println(p.ContentTypes())
... codes below ...
```

# getting json data
 ```go
... codes above ...
	p, err := storage.Parse(c.Request)
	if err != nil {
		fmt.Println(err)
	}
	b, err := p.JSON()
	if err != nil {
		fmt.Println(err)
	}
	var foo Foo
	err := json.Unmarshal(b, foo)
	if err != nil {
		return err
	}
... codes below ...
```
# more 
[Handling file uploads simple and memory friendly in Go with Baraka](https://pkg.go.dev/github.com/xis/baraka)

# contributing
 pull requests are welcome. please open an issue first to discuss what you would like to change.

 please make sure to update tests as appropriate.

# license
[MIT](https://choosealicense.com/licenses/mit/)