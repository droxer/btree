package btree_test

import (
    "github.com/droxer/btree"
    "math/rand"
    "testing"
    "time"
)

func init() {
    seed := time.Now().Unix()
    rand.Seed(seed)
}

func perm(n int) (out []item) {
    for _, v := range rand.Perm(n) {
        out = append(out, item(v))
    }
    return
}

func BenchmarkInsert1000(b *testing.B) {
    benchmarkInsert(1000, b)
}

func BenchmarkInsert10000(b *testing.B) {
    benchmarkInsert(10000, b)
}

func BenchmarkInsert100000(b *testing.B) {
    benchmarkInsert(100000, b)
}

func BenchmarkGet1000(b *testing.B) {
    benchmarkGet(1000, b)
}

func BenchmarkGet10000(b *testing.B) {
    benchmarkGet(10000, b)
}

func BenchmarkGet100000(b *testing.B) {
    benchmarkGet(100000, b)
}

func benchmarkInsert(size int, b *testing.B) {
    b.StopTimer()
    values := perm(size)
    b.StartTimer()

    i := 0
    for i < b.N {
        tr := btree.New(3)
        for _, item := range values {
            tr.Insert(item)
            i++

            if i >= b.N {
                return
            }
        }
    }
}

func benchmarkGet(size int, b *testing.B) {
    b.StopTimer()
    insert := perm(size)
    remove := perm(size)
    b.StartTimer()
    i := 0
    for i < b.N {
        b.StopTimer()
        tr := btree.New(3)
        for _, v := range insert {
            tr.Insert(v)
        }
        b.StartTimer()
        for _, item := range remove {
            tr.Get(item)
            i++
            if i >= b.N {
                return
            }
        }
    }
}
