# Diamond Decoder

A simple server that communicates with a command-line tool in order to Decode/Encode text-based art

## Technologies used

- Uiverse.io , I styled the interface with open-source code from here
- Radio buttons by Shoh2008
- Input form by Yaseen549
- Process button by zjssun

## Features

- Decoding/Encoding of text
- Defaults to multiline input.
- Handles errors effectively and in-depth

## Installation

### Prerequisites

Before running the server, make sure you have the following installed and or updated:

- Go (1.18 or later)
- Git (for cloning the repo)

### Steps to use and run

## Clone the repo:

````
git clone https://gitea.koodsisu.fi/emilsundkvist/art.git
cd art
````

## Build the project

````
cd decoder
go build -o diamond-decoder
````

## Run the project: Navigate to the interface directory

````
cd interface
go run .
````


## Usage

### Encoded text

- Encoded patterns have to be inside brackets, if brackets are mismatched it will return an error
- There are two arguments inside those brackets, separated by a SINGLE space. Everything proceding will be considered part of the pattern to be expanded
- First argument has to be a number, otherwise it is invalid
- There has to be a second argument, otherwise it is also invalid
- Square brackets are not printable for simplicity

## Example test cases

````
[8 ⣿]⠟⠋⠁[8 ⠀]⠉⠻⣿
[7 ⣿]⠁[13 ⠀]⢺⣿
[7 ⣿][13 ⠀]⠆⠜⣿
[4 ⣿]⠿⠿⠛[18 ⠀]⠉⠻⣿⣿
⣿⣿⡏⠁[5 ⠀]⣀⣠⣤⣤[5 ⣶]⣦⣤⡄[4 ⠀]⢀⣴⣿
⣿⣿⣷⣄⠀⠀⠀⢠⣾[10 ⣿]⢿⡧⠇⢀⣤⣶
[6 ⣿]⣾⣮⣭⣿⡻⣽⣒⠀⣤⣜⣭⠐⢐⣒⠢⢰
[7 ⣿]⣏[6 ⣿]⡟⣾⣿⠂⢈⢿⣷⣞
[8 ⣿]⣽⣿⣿⣷⣶⣾⡿⠿⣿⠗⠈⢻⣿
[12 ⣿]⡿⠻⠋⠉⠑⠀⠀⢘⢻
[7 ⣿]⡿⠟⢹⣿⣿⡇⢀⣶⣶⠴⠶⠀⠀⢽
[6 ⣿]⡿⠀⠀⢸⣿⣿⠀⠀⠣[5 ⠀]⡟⢿⣿
⣿⣿⣿⡿⠟⠋[4 ⠀]⠹⣿⣧⣀[4 ⠀]⡀⣴⠁⢘⡙
⠉⠉⠁[8 ⠀]⠈⠙⢿⠗⠂⠄⠀⣴⡟⠀⠀⡃
Will output: 
⣿⣿⣿⣿⣿⣿⣿⣿⠟⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠉⠻⣿
⣿⣿⣿⣿⣿⣿⣿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢺⣿
⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠆⠜⣿
⣿⣿⣿⣿⠿⠿⠛⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠻⣿⣿
⣿⣿⡏⠁⠀⠀⠀⠀⠀⣀⣠⣤⣤⣶⣶⣶⣶⣶⣦⣤⡄⠀⠀⠀⠀⢀⣴⣿
⣿⣿⣷⣄⠀⠀⠀⢠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢿⡧⠇⢀⣤⣶
⣿⣿⣿⣿⣿⣿⣾⣮⣭⣿⡻⣽⣒⠀⣤⣜⣭⠐⢐⣒⠢⢰
⣿⣿⣿⣿⣿⣿⣿⣏⣿⣿⣿⣿⣿⣿⡟⣾⣿⠂⢈⢿⣷⣞
⣿⣿⣿⣿⣿⣿⣿⣿⣽⣿⣿⣷⣶⣾⡿⠿⣿⠗⠈⢻⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠻⠋⠉⠑⠀⠀⢘⢻
⣿⣿⣿⣿⣿⣿⣿⡿⠟⢹⣿⣿⡇⢀⣶⣶⠴⠶⠀⠀⢽
⣿⣿⣿⣿⣿⣿⡿⠀⠀⢸⣿⣿⠀⠀⠣⠀⠀⠀⠀⠀⡟⢿⣿
⣿⣿⣿⡿⠟⠋⠀⠀⠀⠀⠹⣿⣧⣀⠀⠀⠀⠀⡀⣴⠁⢘⡙
⠉⠉⠁⠀⠀⠀⠀⠀⠀⠀⠀⠈⠙⢿⠗⠂⠄⠀⣴⡟⠀⠀⡃
````

````
Test input #2:
Input:
___________________________________________________/\\\__
 __/\\\____________________________________________\/\\\__
  _\/\\\____________________________________________\/\\\__
   _\/\\\\\\\\________/\\\\\________/\\\\\___________\/\\\__
    _\/\\\////\\\____/\\\///\\\____/\\\///\\\____/\\\\\\\\\__
     _\/\\\\\\\\/____/\\\__\//\\\__/\\\__\//\\\__/\\\////\\\__
      _\/\\\///\\\___\//\\\__/\\\__\//\\\__/\\\__\/\\\__\/\\\__
       _\/\\\_\///\\\__\///\\\\\/____\///\\\\\/___\//\\\\\\\/\\_
        _\///____\///_____\/////________\/////______\///////\//__

Will output:

[51 _]/\\\__
 __/\\\[44 _]\/\\\__
  _\/\\\[44 _]\/\\\__
   _\/[8 \][2 ________/\\\\\][11 _]\/\\\__
[4  ]_\/\\\[4 /][2 \\\____/\\\///]\\\[4 _]/[9 \]__
[5  ]_\/[8 \]/[4 _][2 /\\\__\//\\\__]/\\\[4 /]\\\__
[6  ]_\/\\\///\\\_[2 __\//\\\__/\\\][2 __\/\\\]__
[7  ]_\/\\\_\///\\\[2 __\///\\\\\/__]_\//[7 \]/\\_
[8  ][2 _\///___][2 __\/////______]\[7 /]\//__
````

### Testing

These tests should be ran in interface directory, 
you have to open another terminal window separate to where the server is running.

## GET / endpoint (Should return HTTP 200):

````
curl -I http://localhost:8080/
````

## POST /decoder with valid input(Should return HTTP 202):

````
curl -X POST -d "text=your_encoded_text&mode=decode" http://localhost:8080/decoder -i
````

## POST /decoder with invalid input (Should return HTTP 400):

````
curl -X POST -d "text=&mode=decode" http://localhost:8080/decoder -i
````

## GET / return main page:

````
curl http://localhost:8080/ | grep -i "diamond decoder"
````




## Author

- [Emil Sundkvist](https://gitea.koodsisu.fi/emilsundkvist)