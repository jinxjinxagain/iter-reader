# iter-reader
iter-reader transform iterator to a io.reader

## install
To start using strkit, install Go and run `go get`;

```sh
$ go get -u github.com/jinxjinxagain/iter-reader
```

## Example
transform a [mgo.Iter](https://godoc.org/gopkg.in/mgo.v2#Iter) to a io.reader

```go
package main

import (
    "github.com/jinxjinxagain/iter-reader"
    "io/ioutil"
)

func main() {
    // TODO: 
    //    Get an iterator from you mongodb
    var reader = ireader.NewReader(func()(result []byte, err error) {
        var res = *Object{}
        var energetic = iter.Next(&res)
        if res != nil {
            result, _ = json.Marshal(res)
        }
        if !energetic {
            err = io.EOF
        }
        return result, err
    })

    var bytes = ioutil.ReadAll(reader)
    fmt.Println(bytes)
}
```

output
```
Documents from your mongodb
```

## Contact
Be free to mail [@jinxjinxagain](jinxjinxagain1994@gmail.com)

## License
strkit source code is available under the MIT[License](/LICENSE).