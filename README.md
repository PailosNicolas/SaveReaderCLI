# SaveReaderCLI
This is a CLI implementation of [GoPkmSaveReader](https://github.com/PailosNicolas/GoPkmSaveReader) using [Bubbles](https://github.com/charmbracelet/bubbles), I wanted to test the save reader module as an user perspective so I could give myself feedback, find mistakes and improve on them.

## Usage
You will need at least Go **1.22** (It probably works with earlier versions but I haven't tested them).
```bash
go run main.go
```
If you want to compile it:
```bash
go build main.go && ./main.go
```

So far you can do two main things, read a save file and read a pokemon file.

### Reading a save file allows you to:
    -General information
    -Check team information
    -Export team member (it can be read)

### Reading a pokemon file allows you to:
    -General information
    -Check stats (same view as reading save file)
    -Check moves information.