# wordle-helper

This is a small helper program for Wordle players to generate guess suggestions. It was written largely to refamiliarize myself with Go.

## Usage

From the project root execute `go run .` to launch the helper. It will prompt you up to six times for your guess and your results from [Wordle](https://www.nytimes.com/games/wordle/index.html).

```
Guess: crane
Results: xxyyg
Suggested guesses: [ranee manse dance mange rance range lance hanse hance zonae]
```
Guesses must be five-characters long and consist only of letters. 

Results must be five-characters long and consist only of **x** (gray tile) **y** (yellow tile) or **g** (green tile) corresponding to the output from your guess in [Wordle](https://www.nytimes.com/games/wordle/index.html).

After each input set of guess and result it will output the 10 best guesses, based on the strategy described below, or less if there are not 10 words remaining to suggest.

## Strategy

The word-suggestion strategy used in the helper is very naive. The following rules are checked for all words in the list, and if any rules are broken the word is discarded from possible guesses:

* If the word contains a letter marked **x** in the result, discard the word
* If the word has a letter marked **y** in the result in the same position, discard the word
* If the word does not contain a letter marked **y** in the result in some position, discard the word
* If the word does **not** have a letter marked **g** in the result in the same position, discard the word

On each iteration this removes invalid words from the list, allowing a new top-10 suggestions to reach the highest scores.