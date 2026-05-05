# ✈️ Cahier d'exos Go — Vol 5h offline

**Setup avant le vol :**
1. Vérifie que Go est installé : `go version`
2. Crée un dossier `avion/`
3. À l'intérieur : `go mod init avion`
4. Pour chaque exo : crée un sous-dossier (ex: `exo01/`), mets `xxx.go` ET `xxx_test.go` dedans
5. Lance avec : `cd exo01 && go test -v`

**Règles du jeu :**
- Tu fais l'exo SEUL, sans regarder la correction
- Si tu bloques 10 min, tu vas lire la correction, tu fermes, tu retapes
- 12 exos progressifs, ~25 min chacun = 5h pile
- Les 4 derniers sont costauds, prends ton temps

**Note importante** : pas de `import "github.com/..."` parce que t'auras pas internet pour télécharger les modules. On reste 100% stdlib.

---

# 📋 LES EXOS

---

## Exo 1 — Reverse de string (15 min)

Écris une fonction `Reverse(s string) string` qui inverse une chaîne.

**Exemples :**
- `Reverse("hello")` → `"olleh"`
- `Reverse("")` → `""`
- `Reverse("a")` → `"a"`
- `Reverse("héllo")` → `"olléh"` (gère l'unicode)

**Indices :**
- Convertis en `[]rune` pour gérer l'unicode (sinon tu casses les caractères accentués)
- Boucle classique : `for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1`

**Test à écrire** : table-driven avec au moins 4 cas (vide, un caractère, classique, avec accent).

---

## Exo 2 — Min de slice avec erreur (15 min)

Écris `Min(nums []int) (int, error)`.

- Retourne le plus petit élément
- Si slice vide → erreur `ErrEmptySlice` (sentinel exportée)

**Test** : table-driven avec champ `wantErr bool`. Au minimum 5 cas dont vide et nil.

**Tu dois utiliser `errors.Is` dans le test pour vérifier l'erreur.**

---

## Exo 3 — Compter les voyelles (15 min)

Écris `CountVowels(s string) int` qui compte les voyelles (a, e, i, o, u, y, MAJ et min).

**Exemples :**
- `CountVowels("hello")` → 2
- `CountVowels("HELLO")` → 2
- `CountVowels("")` → 0
- `CountVowels("rythm")` → 1 (le y compte)

**Indices :**
- Tu peux mettre les voyelles dans une `map[rune]bool` ou utiliser `strings.ContainsRune`
- Itère avec `for _, r := range s` (range sur string te donne des runes)
- N'oublie pas de gérer les majuscules : `strings.ToLower` AVANT, ou check les deux cas

**Test** : table-driven, 6+ cas.

---

## Exo 4 — Filtrer un slice (20 min)

Écris `Filter(nums []int, pred func(int) bool) []int` qui retourne un nouveau slice contenant uniquement les éléments satisfaisant le prédicat.

**Exemples :**
```go
isEven := func(n int) bool { return n%2 == 0 }
Filter([]int{1, 2, 3, 4, 5}, isEven) // [2, 4]

isPositive := func(n int) bool { return n > 0 }
Filter([]int{-2, -1, 0, 1, 2}, isPositive) // [1, 2]
```

**Indices :**
- Une fonction peut être passée en paramètre (first-class functions en Go)
- Initialise `result := []int{}` ou `make([]int, 0)`
- `append(result, v)` quand `pred(v)` est true

**Test à écrire** : 3+ cas avec différents prédicats. Pour comparer 2 slices dans le test, utilise `reflect.DeepEqual` :

```go
import "reflect"

if !reflect.DeepEqual(got, tc.want) {
    t.Errorf("Filter(%v) = %v, want %v", tc.in, got, tc.want)
}
```

---

## Exo 5 — Compter les mots dans une phrase (20 min)

Écris `WordCount(s string) map[string]int` qui retourne un map du nombre d'occurrences de chaque mot.

**Exemples :**
- `WordCount("le chat le chien")` → `map[string]int{"le": 2, "chat": 1, "chien": 1}`
- `WordCount("")` → `map[string]int{}`
- Doit ignorer la casse : `WordCount("Le le LE")` → `{"le": 3}`

**Indices :**
- `strings.Fields(s)` splitte sur les whitespaces
- `strings.ToLower(s)` pour la casse
- Init la map : `result := make(map[string]int)`
- Pour incrémenter : `result[word]++` (Go gère le cas "n'existe pas encore" : ça met 0 + 1 = 1)

**Test** : utilise `reflect.DeepEqual` pour comparer des maps. 4+ cas.

---

## Exo 6 — Validateur d'email (basique) (20 min)

Écris `IsValidEmail(s string) bool`.

Règles simplifiées :
- Doit contenir exactement UN `@`
- Le `@` ne peut être ni au début ni à la fin
- Doit y avoir au moins un `.` APRÈS le `@`
- Le `.` ne peut être juste après le `@`
- Pas de spaces

**Exemples :**
- `IsValidEmail("yoann@gmail.com")` → true
- `IsValidEmail("yoann@gmail")` → false (pas de point)
- `IsValidEmail("@gmail.com")` → false (commence par @)
- `IsValidEmail("yoann@.com")` → false (point juste après @)
- `IsValidEmail("yo ann@gmail.com")` → false (espace)
- `IsValidEmail("")` → false

**Indices :**
- `strings.Count(s, "@")` pour compter les @
- `strings.Index(s, "@")` pour trouver la position
- `strings.Contains(s, " ")` pour les espaces

**Test** : 8+ cas couvrant tous les cas valides et invalides.

---

## Exo 7 — Stack (LIFO) avec generics (25 min)

Implémente une pile générique :

```go
type Stack[T any] struct {
    items []T
}

func NewStack[T any]() *Stack[T]
func (s *Stack[T]) Push(item T)
func (s *Stack[T]) Pop() (T, error)  // erreur si stack vide
func (s *Stack[T]) Peek() (T, error) // top sans pop
func (s *Stack[T]) Size() int
func (s *Stack[T]) IsEmpty() bool
```

**Sentinel error** : `ErrEmptyStack`.

**Indices :**
- Push = `append(s.items, item)`
- Pop = récupère le dernier, slice s.items pour enlever, retourne
- Peek = juste lis le dernier sans modifier
- Pour la valeur zéro de T dans le `return ?, err` : utilise `var zero T` puis `return zero, err`

**Test** : teste avec `Stack[int]` ET `Stack[string]`. Au moins 6 cas couvrant push, pop, peek vide, size.

---

## Exo 8 — Roman numerals → int (30 min)

Écris `RomanToInt(s string) (int, error)` qui convertit un nombre romain en entier.

**Règles :**
- Symboles : I=1, V=5, X=10, L=50, C=100, D=500, M=1000
- Si un symbole plus petit précède un plus grand, on soustrait : IV = 4, IX = 9, XL = 40, etc.
- Sinon on additionne

**Exemples :**
- `RomanToInt("III")` → 3
- `RomanToInt("IV")` → 4
- `RomanToInt("IX")` → 9
- `RomanToInt("LVIII")` → 58 (L=50, V=5, III=3)
- `RomanToInt("MCMXCIV")` → 1994 (M=1000, CM=900, XC=90, IV=4)
- `RomanToInt("")` → erreur
- `RomanToInt("ABC")` → erreur (chars invalides)

**Indices :**
- `map[rune]int` pour les valeurs des symboles
- Boucle de gauche à droite : compare la valeur courante avec la SUIVANTE
- Si `current < next` → soustraire, sinon additionner

**Test** : 8+ cas dont 2 cas d'erreur.

---

## Exo 9 — Binary search (30 min)

Écris `BinarySearch(sorted []int, target int) (int, bool)` qui retourne :
- L'index du target si trouvé + true
- -1 + false sinon

Le slice est garanti trié en entrée. **Tu DOIS utiliser l'algo binary search** (O(log n)), pas un range simple.

**Exemples :**
- `BinarySearch([]int{1,3,5,7,9}, 5)` → (2, true)
- `BinarySearch([]int{1,3,5,7,9}, 4)` → (-1, false)
- `BinarySearch([]int{}, 1)` → (-1, false)
- `BinarySearch([]int{1}, 1)` → (0, true)

**Algo (en pseudo) :**
```
low = 0
high = len(sorted) - 1
while low <= high:
    mid = (low + high) / 2
    if sorted[mid] == target: return mid, true
    if sorted[mid] < target: low = mid + 1
    else: high = mid - 1
return -1, false
```

**Test** : 6+ cas. Pense aux edge cases : début, fin, milieu, absent, slice vide, single element.

---

## Exo 10 — Parser une URL (avec stdlib) (25 min)

Écris `ParseURL(s string) (URLParts, error)` qui parse une URL et retourne ses composants.

```go
type URLParts struct {
    Scheme string  // "http", "https", "ftp"
    Host   string  // "example.com"
    Port   string  // "8080" ou "" si absent
    Path   string  // "/api/users"
}
```

**Exemples :**
- `ParseURL("https://example.com/path")` → `{Scheme: "https", Host: "example.com", Path: "/path"}`
- `ParseURL("http://localhost:8080/api")` → `{Scheme: "http", Host: "localhost", Port: "8080", Path: "/api"}`
- `ParseURL("ftp://files.org")` → `{Scheme: "ftp", Host: "files.org"}`
- `ParseURL("not-a-url")` → erreur

**Tu DOIS utiliser `net/url` de la stdlib.** C'est exactement le genre d'exo où tu dois lire la doc.

**Doc stdlib offline** : tape `go doc net/url` dans ton terminal, ou `go doc net/url.Parse`. Go embarque toute la doc en local !

**Test** : 5+ cas dont au moins 1 avec port et 1 cas d'erreur.

---

## Exo 11 — Tri par insertion (30 min)

Écris `InsertionSort(nums []int)` qui trie un slice in-place avec l'algo insertion sort.

**Algo :**
```
pour i de 1 à len(nums)-1:
    courant = nums[i]
    j = i - 1
    tant que j >= 0 et nums[j] > courant:
        nums[j+1] = nums[j]
        j--
    nums[j+1] = courant
```

C'est lent (O(n²)) mais simple à implémenter et c'est un classique.

**Test** : 5+ cas. Cas tordus à inclure :
- Slice vide
- Slice déjà trié
- Slice trié inverse
- Avec doublons
- Single element

**Comparaison de slices** : `reflect.DeepEqual`.

---

## Exo 12 — BOSS : Mini parser CSV (45 min)

Écris une fonction `ParseCSV(s string) [][]string` qui parse une string CSV simple.

**Règles :**
- Champs séparés par `,`
- Lignes séparées par `\n`
- Trim les espaces autour des champs
- Ignore les lignes vides
- Pas besoin de gérer les guillemets ou échappements (juste basique)

**Exemple :**
```
Entrée:
"name, age, city
Alice, 30, Paris
Bob,25,Lyon

Charlie, 40, Marseille"

Sortie:
[][]string{
    {"name", "age", "city"},
    {"Alice", "30", "Paris"},
    {"Bob", "25", "Lyon"},
    {"Charlie", "40", "Marseille"},
}
```

**Bonus 1** : ajoute une fonction `ParseCSVAsMap(s string) []map[string]string` qui utilise la première ligne comme headers et retourne un slice de maps.

**Bonus 2** : ajoute la gestion des erreurs si une ligne n'a pas le même nombre de champs que les headers.

**Indices :**
- `strings.Split(s, "\n")` pour les lignes
- `strings.Split(line, ",")` pour les champs
- `strings.TrimSpace(field)` pour trim

**Test** : au moins 5 cas. Inclure : vide, une ligne, plusieurs lignes, lignes vides au milieu, espaces partout.

---

# ✅ LES CORRECTIONS

⚠️ **NE LIS PAS LES CORRECTIONS AVANT D'AVOIR FAIT L'EXO** ⚠️

Tu te connais : si tu lis avant, ton cerveau ne fait pas le boulot d'écriture, donc les automatismes ne reviennent pas. Force-toi à galérer 10-15 min seul. C'est la galère qui construit les réflexes.

---

## ✅ Correction Exo 1 — Reverse

```go
// reverse.go
package exo01

func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
```

```go
// reverse_test.go
package exo01

import "testing"

func TestReverse(t *testing.T) {
    tests := []struct {
        name string
        in   string
        want string
    }{
        {"vide", "", ""},
        {"un caractère", "a", "a"},
        {"classique", "hello", "olleh"},
        {"avec accent", "héllo", "olléh"},
        {"palindrome", "kayak", "kayak"},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got := Reverse(tc.in)
            if got != tc.want {
                t.Errorf("Reverse(%q) = %q, want %q", tc.in, got, tc.want)
            }
        })
    }
}
```

**Points clés :**
- `[]rune(s)` est CRUCIAL. Si tu fais `[]byte(s)`, le test "avec accent" foire (le caractère é fait 2 bytes en UTF-8).
- Le swap multi-affectation `r[i], r[j] = r[j], r[i]` est idiomatique Go.
- `i < j` (pas `<=`) pour ne pas échanger l'élément du milieu avec lui-même.

---

## ✅ Correction Exo 2 — Min

```go
// min.go
package exo02

import "errors"

var ErrEmptySlice = errors.New("slice is empty")

func Min(nums []int) (int, error) {
    if len(nums) == 0 {
        return 0, ErrEmptySlice
    }
    m := nums[0]
    for _, v := range nums[1:] {
        if v < m {
            m = v
        }
    }
    return m, nil
}
```

```go
// min_test.go
package exo02

import (
    "errors"
    "testing"
)

func TestMin(t *testing.T) {
    tests := []struct {
        name    string
        in      []int
        want    int
        wantErr bool
    }{
        {"un élément", []int{5}, 5, false},
        {"plusieurs", []int{3, 1, 4, 1, 5, 9, 2, 6}, 1, false},
        {"négatifs", []int{-5, -2, -10}, -10, false},
        {"slice vide", []int{}, 0, true},
        {"slice nil", nil, 0, true},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got, err := Min(tc.in)
            if tc.wantErr {
                if !errors.Is(err, ErrEmptySlice) {
                    t.Errorf("attendait ErrEmptySlice, got %v", err)
                }
                return
            }
            if err != nil {
                t.Fatalf("erreur inattendue: %v", err)
            }
            if got != tc.want {
                t.Errorf("Min(%v) = %d, want %d", tc.in, got, tc.want)
            }
        })
    }
}
```

**Points clés :**
- Sentinel error exportée pour permettre `errors.Is` côté appelant.
- `nums[1:]` permet d'éviter de comparer `nums[0]` à lui-même au premier tour.
- `return zero, err` quand erreur (pas d'accès à `nums[0]` !).

---

## ✅ Correction Exo 3 — Voyelles

```go
// vowels.go
package exo03

import "strings"

func CountVowels(s string) int {
    vowels := "aeiouy"
    s = strings.ToLower(s)
    count := 0
    for _, r := range s {
        if strings.ContainsRune(vowels, r) {
            count++
        }
    }
    return count
}
```

```go
// vowels_test.go
package exo03

import "testing"

func TestCountVowels(t *testing.T) {
    tests := []struct {
        name string
        in   string
        want int
    }{
        {"vide", "", 0},
        {"hello", "hello", 2},
        {"HELLO", "HELLO", 2},
        {"rythm avec y", "rythm", 1},
        {"que des voyelles", "aeiouy", 6},
        {"que des consonnes", "bcdfgh", 0},
        {"avec espaces", "hello world", 3},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got := CountVowels(tc.in)
            if got != tc.want {
                t.Errorf("CountVowels(%q) = %d, want %d", tc.in, got, tc.want)
            }
        })
    }
}
```

**Points clés :**
- `range s` sur une string te donne des `rune` (pas des `byte`), parfait pour l'unicode.
- `strings.ContainsRune` est plus simple qu'une map pour 6 valeurs.
- `ToLower` AVANT la boucle évite de checker les 12 cas (a,A,e,E...).

---

## ✅ Correction Exo 4 — Filter

```go
// filter.go
package exo04

func Filter(nums []int, pred func(int) bool) []int {
    result := []int{}
    for _, v := range nums {
        if pred(v) {
            result = append(result, v)
        }
    }
    return result
}
```

```go
// filter_test.go
package exo04

import (
    "reflect"
    "testing"
)

func TestFilter(t *testing.T) {
    isEven := func(n int) bool { return n%2 == 0 }
    isPositive := func(n int) bool { return n > 0 }
    isMagic := func(n int) bool { return n == 42 }

    tests := []struct {
        name string
        in   []int
        pred func(int) bool
        want []int
    }{
        {"pairs", []int{1, 2, 3, 4, 5}, isEven, []int{2, 4}},
        {"positifs", []int{-2, -1, 0, 1, 2}, isPositive, []int{1, 2}},
        {"slice vide", []int{}, isEven, []int{}},
        {"aucun match", []int{1, 3, 5}, isEven, []int{}},
        {"que des matches", []int{42, 42, 42}, isMagic, []int{42, 42, 42}},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got := Filter(tc.in, tc.pred)
            if !reflect.DeepEqual(got, tc.want) {
                t.Errorf("Filter(%v) = %v, want %v", tc.in, got, tc.want)
            }
        })
    }
}
```

**Points clés :**
- Initialise `result := []int{}` (slice vide) plutôt que `var result []int` (nil). Différence visible pour `reflect.DeepEqual` !
- Le predicat est juste une `func(int) bool` — Go traite les fonctions comme valeurs.

---

## ✅ Correction Exo 5 — WordCount

```go
// wordcount.go
package exo05

import "strings"

func WordCount(s string) map[string]int {
    result := make(map[string]int)
    s = strings.ToLower(s)
    for _, word := range strings.Fields(s) {
        result[word]++
    }
    return result
}
```

```go
// wordcount_test.go
package exo05

import (
    "reflect"
    "testing"
)

func TestWordCount(t *testing.T) {
    tests := []struct {
        name string
        in   string
        want map[string]int
    }{
        {"vide", "", map[string]int{}},
        {"un mot", "bonjour", map[string]int{"bonjour": 1}},
        {"répétition", "le chat le chien le chat", map[string]int{"le": 3, "chat": 2, "chien": 1}},
        {"casse mixte", "Le le LE", map[string]int{"le": 3}},
        {"espaces multiples", "  bonjour   monde  ", map[string]int{"bonjour": 1, "monde": 1}},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got := WordCount(tc.in)
            if !reflect.DeepEqual(got, tc.want) {
                t.Errorf("WordCount(%q) = %v, want %v", tc.in, got, tc.want)
            }
        })
    }
}
```

**Points clés :**
- `strings.Fields` est PLUS BIEN que `strings.Split(s, " ")` car il gère les espaces multiples et les whitespace exotiques (tabs, etc.).
- `result[word]++` : si `word` n'existe pas, Go retourne 0 (valeur zéro de `int`), puis incrémente à 1. Pas besoin de check préalable.
- `make(map[string]int)` plutôt que `var m map[string]int` (qui est nil et ne supporte pas l'écriture).

---

## ✅ Correction Exo 6 — Email

```go
// email.go
package exo06

import "strings"

func IsValidEmail(s string) bool {
    if s == "" {
        return false
    }
    if strings.Contains(s, " ") {
        return false
    }
    if strings.Count(s, "@") != 1 {
        return false
    }
    atIdx := strings.Index(s, "@")
    if atIdx == 0 || atIdx == len(s)-1 {
        return false
    }
    domain := s[atIdx+1:]
    if !strings.Contains(domain, ".") {
        return false
    }
    if domain[0] == '.' {
        return false
    }
    return true
}
```

```go
// email_test.go
package exo06

import "testing"

func TestIsValidEmail(t *testing.T) {
    tests := []struct {
        name string
        in   string
        want bool
    }{
        {"valide simple", "yoann@gmail.com", true},
        {"valide complexe", "first.last+tag@domain.co.uk", true},
        {"vide", "", false},
        {"sans @", "yoanngmail.com", false},
        {"deux @", "yoann@@gmail.com", false},
        {"@ au début", "@gmail.com", false},
        {"@ à la fin", "yoann@", false},
        {"pas de point dans domain", "yoann@gmail", false},
        {"point juste après @", "yoann@.com", false},
        {"avec espace", "yo ann@gmail.com", false},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got := IsValidEmail(tc.in)
            if got != tc.want {
                t.Errorf("IsValidEmail(%q) = %v, want %v", tc.in, got, tc.want)
            }
        })
    }
}
```

**Points clés :**
- Pattern "early return" : check chaque condition d'invalidité, retourne `false` direct. Plus lisible qu'un gros `if a && b && c && d`.
- `s[atIdx+1:]` pour avoir la partie domaine.
- En vrai en prod tu utilises `net/mail.ParseAddress`, mais l'exo ici c'est de coder à la main pour pratiquer.

---

## ✅ Correction Exo 7 — Stack générique

```go
// stack.go
package exo07

import "errors"

var ErrEmptyStack = errors.New("stack is empty")

type Stack[T any] struct {
    items []T
}

func NewStack[T any]() *Stack[T] {
    return &Stack[T]{items: []T{}}
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
    var zero T
    if len(s.items) == 0 {
        return zero, ErrEmptyStack
    }
    n := len(s.items) - 1
    item := s.items[n]
    s.items = s.items[:n]
    return item, nil
}

func (s *Stack[T]) Peek() (T, error) {
    var zero T
    if len(s.items) == 0 {
        return zero, ErrEmptyStack
    }
    return s.items[len(s.items)-1], nil
}

func (s *Stack[T]) Size() int {
    return len(s.items)
}

func (s *Stack[T]) IsEmpty() bool {
    return len(s.items) == 0
}
```

```go
// stack_test.go
package exo07

import (
    "errors"
    "testing"
)

func TestStackInt(t *testing.T) {
    s := NewStack[int]()

    if !s.IsEmpty() {
        t.Error("nouveau stack devrait être vide")
    }

    _, err := s.Pop()
    if !errors.Is(err, ErrEmptyStack) {
        t.Errorf("Pop sur stack vide: attendait ErrEmptyStack, got %v", err)
    }

    s.Push(1)
    s.Push(2)
    s.Push(3)

    if s.Size() != 3 {
        t.Errorf("Size = %d, want 3", s.Size())
    }

    top, err := s.Peek()
    if err != nil || top != 3 {
        t.Errorf("Peek = %d, %v, want 3, nil", top, err)
    }

    v, _ := s.Pop()
    if v != 3 {
        t.Errorf("Pop = %d, want 3", v)
    }
    v, _ = s.Pop()
    if v != 2 {
        t.Errorf("Pop = %d, want 2", v)
    }

    if s.Size() != 1 {
        t.Errorf("après 2 pops, Size = %d, want 1", s.Size())
    }
}

func TestStackString(t *testing.T) {
    s := NewStack[string]()
    s.Push("a")
    s.Push("b")
    v, _ := s.Pop()
    if v != "b" {
        t.Errorf("got %q, want %q", v, "b")
    }
}
```

**Points clés :**
- `[T any]` après le nom du type/fonction = type parameter.
- `var zero T` te donne la valeur zéro pour n'importe quel type T (utile pour les returns d'erreur).
- `s.items[:n]` pour retirer le dernier — ne réalloue pas, juste réduit la longueur.
- Tester avec différents types (`Stack[int]` ET `Stack[string]`) prouve que les generics marchent.

---

## ✅ Correction Exo 8 — Roman to Int

```go
// roman.go
package exo08

import (
    "errors"
    "fmt"
)

var ErrInvalidRoman = errors.New("invalid roman numeral")

func RomanToInt(s string) (int, error) {
    if s == "" {
        return 0, ErrInvalidRoman
    }
    values := map[rune]int{
        'I': 1, 'V': 5, 'X': 10, 'L': 50,
        'C': 100, 'D': 500, 'M': 1000,
    }

    runes := []rune(s)
    total := 0
    for i, r := range runes {
        v, ok := values[r]
        if !ok {
            return 0, fmt.Errorf("%w: %q", ErrInvalidRoman, r)
        }
        // Si le caractère suivant existe et est plus grand, soustraire
        if i+1 < len(runes) {
            nextV, ok := values[runes[i+1]]
            if !ok {
                return 0, fmt.Errorf("%w: %q", ErrInvalidRoman, runes[i+1])
            }
            if v < nextV {
                total -= v
                continue
            }
        }
        total += v
    }
    return total, nil
}
```

```go
// roman_test.go
package exo08

import (
    "errors"
    "testing"
)

func TestRomanToInt(t *testing.T) {
    tests := []struct {
        name    string
        in      string
        want    int
        wantErr bool
    }{
        {"III", "III", 3, false},
        {"IV", "IV", 4, false},
        {"IX", "IX", 9, false},
        {"LVIII", "LVIII", 58, false},
        {"MCMXCIV", "MCMXCIV", 1994, false},
        {"MMXXIV", "MMXXIV", 2024, false},
        {"vide", "", 0, true},
        {"chars invalides", "ABC", 0, true},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got, err := RomanToInt(tc.in)
            if tc.wantErr {
                if !errors.Is(err, ErrInvalidRoman) {
                    t.Errorf("attendait ErrInvalidRoman, got %v", err)
                }
                return
            }
            if err != nil {
                t.Fatalf("erreur inattendue: %v", err)
            }
            if got != tc.want {
                t.Errorf("RomanToInt(%q) = %d, want %d", tc.in, got, tc.want)
            }
        })
    }
}
```

**Points clés :**
- Convertir en `[]rune` pour pouvoir indexer (`runes[i+1]`) — sur une string, l'indexation directe te donne des bytes.
- L'astuce du soustractif : si la valeur courante est plus petite que la suivante, on soustrait au lieu d'ajouter.
- Wrapping d'erreur avec `%w` permet à `errors.Is` de remonter la chaîne.

---

## ✅ Correction Exo 9 — Binary Search

```go
// search.go
package exo09

func BinarySearch(sorted []int, target int) (int, bool) {
    low, high := 0, len(sorted)-1
    for low <= high {
        mid := (low + high) / 2
        if sorted[mid] == target {
            return mid, true
        }
        if sorted[mid] < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return -1, false
}
```

```go
// search_test.go
package exo09

import "testing"

func TestBinarySearch(t *testing.T) {
    tests := []struct {
        name      string
        sorted    []int
        target    int
        wantIdx   int
        wantFound bool
    }{
        {"trouvé milieu", []int{1, 3, 5, 7, 9}, 5, 2, true},
        {"trouvé début", []int{1, 3, 5, 7, 9}, 1, 0, true},
        {"trouvé fin", []int{1, 3, 5, 7, 9}, 9, 4, true},
        {"non trouvé", []int{1, 3, 5, 7, 9}, 4, -1, false},
        {"plus grand que tous", []int{1, 3, 5}, 100, -1, false},
        {"plus petit que tous", []int{1, 3, 5}, -10, -1, false},
        {"slice vide", []int{}, 1, -1, false},
        {"single match", []int{42}, 42, 0, true},
        {"single no match", []int{42}, 0, -1, false},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            gotIdx, gotFound := BinarySearch(tc.sorted, tc.target)
            if gotIdx != tc.wantIdx || gotFound != tc.wantFound {
                t.Errorf("BinarySearch(%v, %d) = (%d, %v), want (%d, %v)",
                    tc.sorted, tc.target, gotIdx, gotFound, tc.wantIdx, tc.wantFound)
            }
        })
    }
}
```

**Points clés :**
- Le piège classique : `mid := (low + high) / 2` peut overflower si low et high sont énormes (low+high dépasse maxInt). En vrai prod, fais `low + (high-low)/2`. Pour des slices Go normaux ça arrivera jamais.
- `low <= high` (pas `<`) pour gérer le cas où il reste un seul élément.

---

## ✅ Correction Exo 10 — Parse URL

```go
// url.go
package exo10

import (
    "errors"
    "net/url"
)

var ErrInvalidURL = errors.New("invalid URL")

type URLParts struct {
    Scheme string
    Host   string
    Port   string
    Path   string
}

func ParseURL(s string) (URLParts, error) {
    u, err := url.Parse(s)
    if err != nil {
        return URLParts{}, ErrInvalidURL
    }
    if u.Scheme == "" || u.Host == "" {
        return URLParts{}, ErrInvalidURL
    }
    return URLParts{
        Scheme: u.Scheme,
        Host:   u.Hostname(),
        Port:   u.Port(),
        Path:   u.Path,
    }, nil
}
```

```go
// url_test.go
package exo10

import (
    "errors"
    "testing"
)

func TestParseURL(t *testing.T) {
    tests := []struct {
        name    string
        in      string
        want    URLParts
        wantErr bool
    }{
        {
            "https simple",
            "https://example.com/path",
            URLParts{Scheme: "https", Host: "example.com", Path: "/path"},
            false,
        },
        {
            "avec port",
            "http://localhost:8080/api",
            URLParts{Scheme: "http", Host: "localhost", Port: "8080", Path: "/api"},
            false,
        },
        {
            "ftp sans path",
            "ftp://files.org",
            URLParts{Scheme: "ftp", Host: "files.org"},
            false,
        },
        {"non URL", "not-a-url", URLParts{}, true},
        {"vide", "", URLParts{}, true},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got, err := ParseURL(tc.in)
            if tc.wantErr {
                if !errors.Is(err, ErrInvalidURL) {
                    t.Errorf("attendait ErrInvalidURL, got %v", err)
                }
                return
            }
            if err != nil {
                t.Fatalf("erreur inattendue: %v", err)
            }
            if got != tc.want {
                t.Errorf("ParseURL(%q) = %+v, want %+v", tc.in, got, tc.want)
            }
        })
    }
}
```

**Points clés :**
- `net/url.Parse` est très permissif (accepte plein de strings comme valides). On doit checker manuellement que Scheme et Host sont remplis.
- `u.Hostname()` retourne le host sans le port. `u.Port()` retourne le port (ou "" si absent).
- Pour `go doc net/url`, tape ça dans ton terminal sans connexion — tu auras la doc complète.

---

## ✅ Correction Exo 11 — Insertion Sort

```go
// sort.go
package exo11

func InsertionSort(nums []int) {
    for i := 1; i < len(nums); i++ {
        current := nums[i]
        j := i - 1
        for j >= 0 && nums[j] > current {
            nums[j+1] = nums[j]
            j--
        }
        nums[j+1] = current
    }
}
```

```go
// sort_test.go
package exo11

import (
    "reflect"
    "testing"
)

func TestInsertionSort(t *testing.T) {
    tests := []struct {
        name string
        in   []int
        want []int
    }{
        {"vide", []int{}, []int{}},
        {"un élément", []int{42}, []int{42}},
        {"déjà trié", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
        {"trié inverse", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
        {"avec doublons", []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}, []int{1, 1, 2, 3, 3, 4, 5, 5, 6, 9}},
        {"négatifs", []int{-3, 5, -1, 0, 2}, []int{-3, -1, 0, 2, 5}},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            // Copie l'input pour ne pas modifier la table entre tests
            input := make([]int, len(tc.in))
            copy(input, tc.in)
            InsertionSort(input)
            if !reflect.DeepEqual(input, tc.want) {
                t.Errorf("InsertionSort(%v) = %v, want %v", tc.in, input, tc.want)
            }
        })
    }
}
```

**Points clés :**
- L'algo modifie le slice EN PLACE (pas de retour). Convention courante en Go pour les opérations sur slices.
- Le `copy` dans le test est important pour que le `tc.in` reste intact entre les tests (sinon tu modifies la table partagée).
- En vrai prod, tu utilises `slices.Sort` de la stdlib (Go 1.21+) qui utilise pdqsort, beaucoup plus rapide.

---

## ✅ Correction Exo 12 — Mini parser CSV

```go
// csv.go
package exo12

import (
    "errors"
    "fmt"
    "strings"
)

var ErrColumnMismatch = errors.New("column count mismatch")

func ParseCSV(s string) [][]string {
    if s == "" {
        return [][]string{}
    }
    lines := strings.Split(s, "\n")
    result := [][]string{}
    for _, line := range lines {
        line = strings.TrimSpace(line)
        if line == "" {
            continue
        }
        fields := strings.Split(line, ",")
        for i := range fields {
            fields[i] = strings.TrimSpace(fields[i])
        }
        result = append(result, fields)
    }
    return result
}

func ParseCSVAsMap(s string) ([]map[string]string, error) {
    rows := ParseCSV(s)
    if len(rows) == 0 {
        return []map[string]string{}, nil
    }
    headers := rows[0]
    result := []map[string]string{}
    for i, row := range rows[1:] {
        if len(row) != len(headers) {
            return nil, fmt.Errorf("%w: line %d has %d fields, expected %d",
                ErrColumnMismatch, i+2, len(row), len(headers))
        }
        m := make(map[string]string)
        for j, val := range row {
            m[headers[j]] = val
        }
        result = append(result, m)
    }
    return result, nil
}
```

```go
// csv_test.go
package exo12

import (
    "errors"
    "reflect"
    "testing"
)

func TestParseCSV(t *testing.T) {
    tests := []struct {
        name string
        in   string
        want [][]string
    }{
        {"vide", "", [][]string{}},
        {
            "une ligne",
            "a,b,c",
            [][]string{{"a", "b", "c"}},
        },
        {
            "plusieurs lignes",
            "name,age\nAlice,30\nBob,25",
            [][]string{
                {"name", "age"},
                {"Alice", "30"},
                {"Bob", "25"},
            },
        },
        {
            "avec espaces",
            "  name , age \n Alice , 30 ",
            [][]string{
                {"name", "age"},
                {"Alice", "30"},
            },
        },
        {
            "avec lignes vides",
            "a,b\n\nc,d\n\n",
            [][]string{{"a", "b"}, {"c", "d"}},
        },
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got := ParseCSV(tc.in)
            if !reflect.DeepEqual(got, tc.want) {
                t.Errorf("ParseCSV(%q) = %v, want %v", tc.in, got, tc.want)
            }
        })
    }
}

func TestParseCSVAsMap(t *testing.T) {
    in := "name,age\nAlice,30\nBob,25"
    got, err := ParseCSVAsMap(in)
    if err != nil {
        t.Fatalf("erreur inattendue: %v", err)
    }
    want := []map[string]string{
        {"name": "Alice", "age": "30"},
        {"name": "Bob", "age": "25"},
    }
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v, want %v", got, want)
    }
}

func TestParseCSVAsMapMismatch(t *testing.T) {
    in := "a,b,c\n1,2"
    _, err := ParseCSVAsMap(in)
    if !errors.Is(err, ErrColumnMismatch) {
        t.Errorf("attendait ErrColumnMismatch, got %v", err)
    }
}
```

**Points clés :**
- En vrai prod, tu utilises `encoding/csv` de la stdlib qui gère les guillemets, les escaping, les multi-lignes. L'exo ici c'est de pratiquer le pattern de parsing.
- Pattern courant : une fonction de bas niveau (`ParseCSV`) renvoie la structure brute, une de plus haut niveau (`ParseCSVAsMap`) ajoute la sémantique.
- `i+2` dans le message d'erreur car `rows[1:]` enlève les headers (ligne 1) et `i` part de 0 → ligne réelle = i + 2.

---

# 🎯 RÉCAP À LA FIN DU VOL

Si tu as fait les 12 exos en 5h, t'as **vraiment** réactivé tes automatismes Go. Quels patterns sont maintenant dans tes doigts :

- ✅ Table-driven tests (12 fois)
- ✅ Sentinel errors + `errors.Is` (6 fois)
- ✅ `range` avec `_` pour ignorer index
- ✅ `[]rune` vs `string` pour Unicode
- ✅ `make` pour maps et slices
- ✅ `reflect.DeepEqual` pour comparer collections
- ✅ Generics (sur Stack)
- ✅ Algos classiques (binary search, insertion sort)
- ✅ Parsing simple

Si tu n'as fait QUE 6 exos, c'est déjà bien. Il vaut mieux faire 6 exos en comprenant tout que survoler les 12 sans rien retenir.

**Garde ce fichier**, tu peux refaire les exos dans 1 semaine. Tu seras choqué de voir comme c'est plus rapide la 2e fois.

Bon vol ✈️
