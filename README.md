
<div align="center">
  <h1>baraka</h1>
  
[![Go Report Card](https://goreportcard.com/badge/github.com/xis/baraka)](https://goreportcard.com/report/github.com/xis/baraka)
[![codecov](https://codecov.io/gh/xis/baraka/branch/master/graph/badge.svg)](https://codecov.io/gh/xis/baraka)
[![Build Status](https://travis-ci.org/xis/baraka.svg?branch=master)](https://travis-ci.org/xis/baraka) 
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/xis/baraka)
  
a tool for handling file uploads for http servers

makes it easier to make operations with files from the http request.
</div>

# **install**
```bash
go get github.com/xis/baraka
```

# **usage**
```go
func main() {
	// create a parser
	parser := baraka.NewParser(baraka.ParserOptions{
		MaxFileSize:   5 << 20,
		MaxFileCount:  5,
		MaxParseCount: 5,
	})

	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		// parsing
		p, err := parser.Parse(c.Request)
		if err != nil {
			fmt.Println(err)
		}
		// saving
		err = p.Save("image_", "./")
		if err != nil {
			fmt.Println(err)
		}
		// getting the part in the []byte format
		parts := p.Content()
		buf := parts[0].Content
		fmt.Println(len(buf))
	})
	router.Run()
}
```
you can use baraka with the other http server libraries, just pass the http.Request to the parser.Parse function.

# **filter function**
filter function is a custom function which filters the files that comes from the request. for example you can read file bytes and identify the file, return true if you wanna pass the file, return false if you don't. 


## filter example
```go
func main() {
	// create a parserr
	parser := baraka.NewParser(baraka.ParserOptions{
		// passing filter function
		Filter: func(data []byte) bool {
			// get first 512 bytes for checking content type
			buf := data[:512]
			// detect the content type
			fileType := http.DetectContentType(buf)
			// if it is jpeg then pass the file
			if fileType == "image/jpeg" {
				return true
			}
			// if not then don't pass
			return false
		},
	})
```
# getting information
```go
	p, err := parser.Parse(c.Request)
	if err != nil {
		fmt.Println(err)
	}
	// prints filenames
	fmt.Println(p.Filenames())
	// prints total files count
	fmt.Println(p.Length())
	// prints content types of files
	fmt.Println(p.ContentTypes())
```

# getting json data
 ```go
	p, err := parser.Parse(c.Request)
	if err != nil {
		fmt.Println(err)
	}
	jsonStrings, err := p.GetJSON()
	if err != nil {
		fmt.Println(err)
	}
```
# more 
*v1.1.1*
[*Handling file uploads simple and memory friendly in Go with Baraka*](https://dev.to/xis/handling-file-uploads-simple-and-memory-friendly-in-go-with-baraka-2h3)

i will make a blog post for v2

# contributing
 pull requests are welcome. please open an issue first to discuss what you would like to change.

 please make sure to update tests as appropriate.

# license
[MIT](https://choosealicense.com/licenses/mit/)
