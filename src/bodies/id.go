package bodies

type ID struct {
  Num int
  Type string
}

func NewID(n int, t string) *ID {
  return &ID{
    Num: n,
    Type: t,
  }
}
