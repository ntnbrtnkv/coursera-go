# [:deciduous_tree: **hw1_tree**](https://github.com/ntnbrtnkv/coursera-go/tree/master/hw1_tree) - simple [*tree*](https://en.wikipedia.org/wiki/Tree_(command)) command implementation

This programs recursivly lists dir's content. It can inclide/exclude files from listing, print file size.

```
> go run main.go testdata
├───project
├───static
│       ├───a_lorem
│       │       └───ipsum
│       ├───css
│       ├───html
│       ├───js
│       └───z_lorem
│               └───ipsum
└───zline
        └───lorem
                └───ipsum

> go run main.go testdata -f
├───project
│       ├───file.txt (19b)
│       └───gopher.png (70372b)
├───static
│       ├───a_lorem
│       │       ├───dolor.txt (empty)
│       │       ├───gopher.png (70372b)
│       │       └───ipsum
│       │               └───gopher.png (70372b)
│       ├───css
│       │       └───body.css (28b)
│       ├───empty.txt (empty)
│       ├───html
│       │       └───index.html (57b)
│       ├───js
│       │       └───site.js (10b)
│       └───z_lorem
│               ├───dolor.txt (empty)
│               ├───gopher.png (70372b)
│               └───ipsum
│                       └───gopher.png (70372b)
├───zline
│       ├───empty.txt (empty)
│       └───lorem
│               ├───dolor.txt (empty)
│               ├───gopher.png (70372b)
│               └───ipsum
│                       └───gopher.png (70372b)
└───zzfile.txt (empty)
```

## Run

```
go run main.go <path> [-f]
```

By default this command prints pnly directories. If you need include files add `-f` as second argument.

## Tests

```
go tests -v
```