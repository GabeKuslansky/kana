# Timeline

1. Want to study Japanese
2. Thought of CLI tool, but know Anki exists
3. Anki has outdated interface with non-customizable keybinds, and I prefer terminal workflow
4. Find spaced repetition algorithm Anki uses -> https://github.com/open-spaced-repetition/fsrs4anki/wiki/The-Algorithm
5. Asked cousin Tommy to learn what the symbols in the equations mean. Learned about partial derivatives. ∂ = partial derative
6. Used `sesh` from Josh Medeski to identify a good CLI lib in Go, found urfave cli
7. Got some basic CLI commands working
8. learned about huh and charm from a youtube short, decided to go with this
9. Learned there is a whole UI framework called bubbletea made by charmbracelet
10. Started trying to implement FSRS in Go. Tons of reading and pre-requisite knowledge required, feeling exhausted... taking a break
11. Maybe there's an easier and quicker solution. Google if there's a way to query Anki, found anki-connect: exposes an http server to query Anki decks
12. Try it out. Holy shit, it works like magic. This also allows me to use Anki for cross-device syncing, another concern I can outsource, as Anki is available on every platform
13. Get interactive prompts working, need to figure out a storage solution for data like default deck to study from
