# The Psychic Poker Player

Rules: read them here https://github.com/Funfun/psychic-poker-player/blob/master/game_rules.md

usage:

1. compile
```
go build -o psychic-poker-player
```

2. run:
```
psychic-poker-player /path/to/input /path/to/output
```

Example:
```
psychic-poker-player sample_input output.out
```

For sample run skip arguments and just run a binary.

testing:
```
go get github.com/stretchr/testify
go test
```

(c) Funfun 2018
