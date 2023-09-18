# go-go-monkey-repl

This is my implementation of the "monkey" langugue interpreter. I made this by following along the wonderful book [Writing an Interpreter in Go](https://interpreterbook.com) by Thorsten Ball. I made a few extentions of my own after finishing, but the bulk of the credit goes to him!

I may play keep playing around with this, but know that it is a toy language, and you should not expect any kind of support if you choose to use it for anything. Though, it is licensed under MIT so there's nothing stoping you from doing that, other than the knowledge that you'll be on your own.

You can run the repl and play aound with it by running `go run main.go`. You can exit the repl by typing entering `quit` (one of my few extentions) or by just killing the process with Ctrl+C on linux. 

Most of the modules come with tests you can run like so:

`go test ./eval`

Other than that, there's not that much to say, other than thank you the Thorsten! I had an amazing time reading the book and I highly recommend you do so too. 

## Not so FAQ
- Q: Why is it called go-go-monkey-repl
- A: I was originally going to just call it "the Go Monkey Repl" since it is a REPL for the Moneky language, written in go, and I intend to write the same in other languages too as an exercise for myself. Then I thought go-go would be a funny inspector Gadget joke.

## Extentions
- 2023-09-18: ability to gracefully exit the repl with the `quit` command.
- 2023-09-18: `iter` builtin function. It gives you an array containing the individual characters when called on a string, and an array with the keys when called on a hash (dictionary). It just returns arrays since loops arent's implemented yet (though I plan to implement at least one) and you can faux iterate over arrays using the builtin `first` function. 
