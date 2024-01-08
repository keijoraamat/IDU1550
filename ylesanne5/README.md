# Ülesanne 5 - 3 kihiline arhitektuur

Rakendus kasutab kolme kihti - esitluskiht (PL), äriloogikakiht (BLL) ja andmekiht (DAL).

## Kasutatud allikad

- https://kool.gunnarpeipman.com/programmeerimine-2/kolmekihiline-arhitektuur/

- https://learning.oreilly.com/library/view/software-architecture-patterns/9781098134280/ch03.html#idm46088527971120

- https://www.ibm.com/topics/three-tier-architecture


## Sõltuvused

- sqlite, kasutamiseks
- Go kompilaator, ehitamiseks

## Kuidas

### Ehitada paigaldamiseks käivitatav fail

```shell
go build .
```

See käsk ehitab ehitamismasina arhitektuurile sobiva binaari. Teistele platvormidele binaari ehitamiseks - https://go.dev/doc/tutorial/compile-install

### Käivitada ilma ehitamata

```shell
go run .
```

### Käivitada kõiki teste

```shell
go test -v ./... -cover
```
