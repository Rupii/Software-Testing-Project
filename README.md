# Software-Testing-Project
---------------------------------------
Simple Software testing for cards project Go!


# Go! ![Golang](https://intami.pl//wp-content/uploads/sites/3/2016/09/golang-icon.png)
-------------------------------------
- Go the one of the very few languages which is not written in C.
- Go is buit to leverage the processing power of all the intel cores.
- go routines and channels helps achieving the processing speed
- It is satically typted pure function programming language.

## Run GO!
> there are two types of go programs
> - Exectable (containing main package)
> - Reproducble (containing package other then main)
> - Save under the name filneName.go
  
`>>> run go main.go`
> - Run muliple go programs
`>>> run go main.go secondFile.go`

## Project 
> The project is a simple cards one the underlining principle is testing the code
- Functions of the project
```
newDeck()
print()
deal() // returns the hand and remaining cards
shuffle()
saveToFile() // stores a file in HardDrive
deckFromFile() //Retrives a file from HardDrive
```

## Testing

- Testing in go is not exactly the testing in softwares like Seleneum 
- the test file show end with the _test.go in our case deck_test.go
- the funcitions I will be testing are
```
TestNewDeck()
TestToAndFrom() // check if the file is not carshing while the exchange between the HardDrive 
```

*Run Test*

`>>> go test`
