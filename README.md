
<div align="center">
	<pre>                         
 ▄▄▄▄    ▄▄▄       ██▀███   ▄▄▄       ██ ▄█▀▄▄▄      
▓█████▄ ▒████▄    ▓██ ▒ ██▒▒████▄     ██▄█▒▒████▄    
▒██▒ ▄██▒██  ▀█▄  ▓██ ░▄█ ▒▒██  ▀█▄  ▓███▄░▒██  ▀█▄  
▒██░█▀  ░██▄▄▄▄██ ▒██▀▀█▄  ░██▄▄▄▄██ ▓██ █▄░██▄▄▄▄██ 
░▓█  ▀█▓ ▓█   ▓██▒░██▓ ▒██▒ ▓█   ▓██▒▒██▒ █▄▓█   ▓██▒
░▒▓███▀▒ ▒▒   ▓▒█░░ ▒▓ ░▒▓░ ▒▒   ▓▒█░▒ ▒▒ ▓▒▒▒   ▓▒█░
▒░▒   ░   ▒   ▒▒ ░  ░▒ ░ ▒░  ▒   ▒▒ ░░ ░▒ ▒░ ▒   ▒▒ ░
 ░    ░   ░   ▒     ░░   ░   ░   ▒   ░ ░░ ░  ░   ▒   
 ░            ░  ░   ░           ░  ░░  ░        ░  ░
      ░                                              
	</pre>
  
[![Go Report Card](https://goreportcard.com/badge/github.com/xis/baraka)](https://goreportcard.com/report/github.com/xis/baraka)
[![codecov](https://codecov.io/gh/xis/baraka/branch/master/graph/badge.svg)](https://codecov.io/gh/xis/baraka)
[![Build Status](https://travis-ci.org/xis/baraka.svg?branch=master)](https://travis-ci.org/xis/baraka) 
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/xis/baraka/v2)
  
a tool for handling file uploads for http servers

makes it easier to make operations with files from the http request.
</div>

## Contents
 - [Install](#Install)
 - [Simple Usage](#Simple-Usage)
 - [Filtering Parts](#Filtering-Parts)
 - [More](#More)
 - [Contribute](#Contribute)
 - [License](#License)

## Install
```bash
go get github.com/xis/baraka/v2
```

## Simple Usage
```go
func main() {
	// create a parser
	parser := baraka.NewParser(baraka.ParserOptions{
		MaxFileSize:   5 << 20,
		MaxFileCount:  5,
		MaxParseCount: 5,
	})

	store := baraka.NewFilesystemStorage("./files")

	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		// parse
		request, err := parser.Parse(c.Request)
		if err != nil {
			fmt.Println(err)
		}

		// get the form
		images, err := request.GetForm("images")
		if err != nil {
			fmt.Println(err)
		}

		// save
		for key, image := range images {
			err = store.Save("images", "image_"+strconv.Itoa(key), image)
			if err != nil {
				fmt.Println(err)
			}
		}
	})
	router.Run()
}
```
You can use baraka with the other http server libraries, just pass the http.Request to the parser.Parse function.

## Filtering Parts
You can filter parts by their properties, like part's content type. Parser can inspect the part's bytes and detect the type of the part with the Inspector.

```go
// create a parser
parser := baraka.NewParser(baraka.ParserOptions{
	MaxFileSize:   5 << 20,
	MaxFileCount:  5,
	MaxParseCount: 5,
})

// give parser an inspector
parser.SetInspector(baraka.NewDefaultInspector(512))
// give parser a filter
parser.SetFilter(baraka.NewExtensionFilter(".jpg"))
```

Now parser will inspect the each part and it will just return the jpeg ones from the Parse function. You can make your own Inspector and Filter.

## Contribute
Pull requests are welcome. please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
