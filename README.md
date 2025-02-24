# Diamond Decoder

Command line tool for decoding and encoding. Encoded text can be decoded into ASCII art and vice versa.

## Features

- Decoding/Encoding of text
- Support for multiline input.
- Can double the size of given input in real-time
- Handles errors effectively and in-depth

## Installation

### Prerequisites

Before running the tool, make sure you have the following installed and or updated:
    - Go (1.18 or later)
    - Git (for cloning the repo)

### Steps to use and run

1. Clone the repo:
````
git clone https://gitea.koodsisu.fi/emilsundkvist/art.git
cd art
````
2. Run the project: Navigate to the decoder directory
````
cd decoder
go run . <flags and input>
````

If you want to you can also build the project:

````
go build -o diamond-decoder
````

Then run it by doing the following:
````
./diamond-decoder -de -2x '[5 A]'
````

## Usage

You can also use a flag to display the usage in the tool
````
go run . -h
````

### Encoding text
- Encoded patterns have to be inside brackets, if brackets are mismatched it will return an error
- There are two arguments inside those brackets, separated by a SINGLE space. Everything proceding will be considered part of the pattern to be expanded
- First argument has to be a number, otherwise it is invalid
- There has to be a second argument, otherwise it is also invalid
- The whole string has to be in quotations, in some cases double quotations will not work.
In that case use single quotations or multiline input.
- Square brackets are not printable for simplicity

## Example test cases
````
Test input: go run . -de '[8 ⣿]⠟⠋⠁[8 ⠀]⠉⠻⣿
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
⠉⠉⠁[8 ⠀]⠈⠙⢿⠗⠂⠄⠀⣴⡟⠀⠀⡃ '
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
Test input #2: go run . -m -en
then press Enter and paste your desired string to encode and press Enter twice.
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

## Author

- [Emil Sundkvist](https://gitea.koodsisu.fi/emilsundkvist)