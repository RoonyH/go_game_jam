## Chuck-A-Blast!

Chuck-A-Blast is a board game that used to be somewhat popular in my country some time ago.

### How to play

* The board contains 45 holes, arranged in a cross shape, and 44 pieces. The lucky hole that is not occupied is right in the middle.

* Objective of the game is to get rid of as many pieces as possible.

* There is only one way to remove a piece. When two pieces and a hole is next to each other, with a piece in the middle the piece on the side can be moved over the piece in the middle to the hole in the side. And then the piece in the middle is Blasted and removed from the board.

* You go from `<piece> <piece> <hole>` to `<hole> <hole> <piece>`

* Game goes on until there are no such valid moves.

* Use arrow keys select a piece to move. Press space to hold the piece. Possible target holes, if any, are now shown. Press the arrow key on the direction of the target hole.

### How to run
If you have Go,
clone the repo and get [termloop](https://github.com/JoelOtter/termloop) and run `main.go` and `game.go`

~~~
git clone https://github.com/RoonyH/go_game_jam
go get -u github.com/JoelOtter/termloop 
cd go_game_jam
go run main.go game.go
~~~

I will add some binaries soon.

### Behind the scenes

This implementation of Chuck-A-Blast is pure [Go](https://golang.org/).

Apart from the standard Go packages its only dependancy is the awesome [termloop](https://github.com/JoelOtter/termloop)

This is written in a weekend for the codelympics project [Go Game Jam](https://codelympics.io/projects/2)

### alert!

The code is written in few hours and I am new to Go.
The code is pretty bad :( I will refactor it soon.

If you are looking inside I appologize for what ever horrible things you may see.
