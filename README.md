# provider

Simple package for creating provider functions. Provider functions are packaged functions that provide a single
value. They are useful for encapsulating variables, performing dependency injections (without reflection),
chaining dependencies together, and providing mocked interfaces for testing

## Get Started

```
go get -u github.com/jaredhughes1012/provider
```

## Working with providers

```
type T1 interface {
  DoThing() string
}

type T2 interface {
  DoThingTwice() string
}

// Depends on a string pulled from config somewhere
type t1 struct {
  configStr string
}

func (t t1) DoThing() string {
  return t.configStr
}

// Depends on a T1
type t2 struct {
  t1 T1
}

func (t t2) DoThingTwice() string {
  return fmt.Sprintf("%s %s", t.t1.DoThing(), t.t1.DoThing())
}

func dependent(t2Provider provider.Provider[T2]) {
  t2 := t2Provider(context.Background())
  fmt.Println(t2.DoThingTwice())
}

func main() {
  configStr := os.Args[1]

  // Provider that encapsulates scoped variable
  t1Provider := provider.New[T1](func (ctx context.Context) (T1, error) {
    return &t1 {
      configStr: configStr,
    }, nil
  })

  // Provider that chains providers
  t2Provider := provider.New[T2](func (ctx context.Context) (T2, error) {
    t1, err := p1Provider(ctx)
    if err != nil {
      return nil, er
    }

    return t2 {
      t1: t1,
    }, nil
  })

  dependent(t2Provider)
}

```