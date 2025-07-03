package slice

func Transform[Src any, Dst any](src []Src, m func(int, Src) Dst, ) []Dst {
    res := make([]Dst, len(src))
    for index, value := range src {
        dest := m(index, value)
        res[index] = dest
    }
    return res
}
