package _go

import (
    "fmt"
)

func test() {
    c := NewCamera[[][]uint8](D300S)
    img := [][]int8{{0, 0, 1}, {2, 0, 0}}
    c.memory = append(c.memory, img)
    fmt.Println(c, len(c.memory), Z9.brand() != A7R_IV.brand())
}

type Camera[I Image] struct {
    model  CameraModel
    memory []I // like a separate, mutable SD card
}

func NewCamera[I Image](m CameraModel) Camera[I] {
    return Camera[I]{m, make([]I, 0)}
}

type CameraModel uint8

const (
    A7R_IV CameraModel = iota
    D300S
    Z9
)

const (
    x CameraModel = 1
)

func (c CameraModel) brand() Brand {
    switch c {
    case A7R_IV:
        s := SonySeries{alpha, 7}
        return Brand{Sony, &s, nil}
    case D300S:
        s := NikonSeries{D | S, 300}
        return Brand{Nikon, nil, &s}
    case Z9:
        s := NikonSeries{Z, 9}
        return Brand{Nikon, nil, &s}
    default:
        return Brand{Unbranded, nil, nil}
    }
}

type Branded interface {
    brand() Brand
}

type Brand struct { // catch all strategy isn't the only way
    BrandKind
    *SonySeries
    *NikonSeries
}

type BrandKind uint8

const (
    Sony BrandKind = iota
    Nikon
    Unbranded
)

type SonySeries struct {
    SonySeriesTags
    level uint8
}

type SonySeriesTags uint8

const ( // note, I don't really know how Sony works
    alpha SonySeriesTags = iota
    NEX
    Cyber
)

func (s SonySeriesTags) String() string {
    switch s {
    case alpha:
        return "alpha"
    case NEX:
        return "NEX"
    case Cyber:
        return "Cyber"
    default:
        return "Unbranded"
    }
}

type NikonSeries struct {
    NikonSeriesTags
    level uint32
}

type NikonSeriesTags uint16

const ( // e.g. Nikon D70, D300S, D2XS, E2NS
    Z NikonSeriesTags = 1 << (iota)
    D
    H
    X
    A
    E
    N
    S
    F
)

const minNS = Z
const maxNS = F

const NikonSeriesStringer = "ZDHXAENSF"

func (s NikonSeriesTags) String() string {
    f := ""
    for b, i := minNS, 0; b <= maxNS; b, i = b<<1, i+1 {
        if (s & b) != 0 {
            f += NikonSeriesStringer[i : i+1]
        }
    }

    return f
}

type Image interface {
    [][]uint32 | [][]uint8
}

func String(brand Brand) string {
    switch brand.BrandKind {
    case Sony:
        return fmt.Sprint("Sony ", brand.SonySeries.SonySeriesTags, brand.SonySeries.level)
    case Nikon:
        return fmt.Sprint("Nikon", brand.NikonSeries.NikonSeriesTags, brand.NikonSeries.level)
    default:
        return "Unbranded"
    }
}
