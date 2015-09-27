## Chuck-A-Blast!

Chuck-A-Blast is a board game that used to be somewhat popular in my country some time ago.

### How to play

* The board contains 45 holes and 44 pieces. The lucky hole that is not occupied is right in the middle.

* Objective of the game is to get rid of as many pieces as possible.

* There is only one way to remove a piece. When two pieces and a hole is next to each other, with a piece in the middle the piece on the side can be moved over the piece in the middle to the hole in the side. And then the piece in the middle is Blasted and removed from the board.

* You go from `<piece> <piece> <hole>` to `<hole> <hole> <piece>`

* Game goes on until there are no such valid moves.

### How to run
If you have Go,
clone the repo and get [termloop](https://github.com/JoelOtter/termloop) and run `main/main.go`

~~~
git clone https://github.com/RoonyH/go_game_jam/
go get -u github.com/JoelOtter/termloop 
cd go_game_jam
go run main/main.go
~~~

I will add some binaries soon.

### Behind the scenes

This implementation of Chuck-A-Blast is pure Go.

Apart from the standard go packages its only dependancy is the awesome [termloop](https://github.com/JoelOtter/termloop)

### alert!

The code is written in few hours and I am new to go.
The code is pretty bad :( I will refactor it soon.

If you are looking inside I appologize for what ever horrible things you may see.
