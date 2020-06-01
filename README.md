# go_cli_flga_example

## introduction

    use flag package to parse input args

    and override flag.Usage to print out custom usage

## how to build

### 1 use make to build binary

### prequest

1 if OS is WINDOW (NOT unix-like), then need GNU-MAKE to install

### command

```script===
make
```
### 2 use go command to build

```golang===
go build -o math -v
```

### cli description

|arguments| description |
| -       | -           |
| -o      | operator default is + |
| -p      | showProcess option default is false|

### usage

```script
./math -o + -p=true 12 3
```