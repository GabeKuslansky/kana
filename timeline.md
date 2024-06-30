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
14. Study how zoxide's Rust implementation stores data locally in their `db.zo`. They create `db.zo` lazily when updating the file, and write to it by serializing a Rust object into binary with the `bincode` crate, which also effeciently handles deserialization. This tells me I can most likely find a similar binary encoding and decoding lib in go to also efficiently handle serialization.
15. I found three main options: Protobufs (`github.com/golang/protobuf/proto`), `encoding/gob` and `encoding/binary`. Ultimately I went with `encoding/gob` because the binary will only be written to and read from Go, so we can use the Go-specific encoding format from the standard lib. Protobufs are useful for language-agnostic serialization and gRPC, both of which are unneccessary as `kana` is a monolithic Go executable that reads and manages a local file. `encoding/binary` lets you have fine grained control over how the binary is encoded, such as endianness. This is useful for serializing binary that needs to conform to another systems requirements, but not needed here.
16. As for where to store the file, [zoxide](https://github.com/ajeetdsouza/zoxide) by default uses these values, so I'll use the same. I'll just implement Mac for now, since I'm developing on one.
| OS          | Path                                     | Example                                    
    | ----------- | ---------------------------------------- | ------------------------------------------ |
    | Linux / BSD | `$XDG_DATA_HOME` or `$HOME/.local/share` | `/home/alice/.local/share`                 |
    | macOS       | `$HOME/Library/Application Support`      | `/Users/Alice/Library/Application Support` |
    | Windows     | `%LOCALAPPDATA%`                         | `C:\Users\Alice\AppData\Local`             |

They lazily create this directory if it doesn't exist when reading from the db  
17. Added a ton of features, but got when using the API to cards, it opened the Anki app, and merely only populated the fields. It's up to the user to confirm. This is not ideal for a TUI. I found they have a Discord and accept contributions, I'll look into my options here.
18. I found some promising information! I found [this GitHub thread](https://github.com/FooSoft/anki-connect/issues/62), and tried running the docker container for headless Anki, but it seems it's built only for Linux - I had trouble pulling this docker image on my mac. I kept digging, and found [a repo that adds cards](https://github.com/NWuensche/AnkiTerminalImporter) to Anki! Just like magic, the answer was right around the corner, I just had to not give up and keep digging. It uses the `anki` python package which, miraculously, is officially maintained by Anki team `Ankitects Pty Ltd`. Had I known there was an anki package to do ever. I found the database file is (on mac) `/Library/Application Support/Anki2/User 1/collection.anki2` with corresponding locations as mentioned in #16.
