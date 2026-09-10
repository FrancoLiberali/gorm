package tests_test

import (
	"fmt"
	"testing"

	. "gorm.io/gorm/utils/tests"
)

// Extended benchmarks that complement benchmark_test.go's seven canonical
// shapes. They mirror the scenarios in cql's benchmarks/extended_bench_test.go
// so the two projects can be compared line for line on the same machine:
//
//   - Find_5Where     -> chained .Where("name <> ?") x5 (multi-condition dispatch)
//   - First           -> DB.First(&u, id)
//   - JoinedPreload   -> DB.Joins("Account")   (has-one, single JOIN query)
//   - HasMany         -> DB.Preload("Pets")    (separate child SELECT + mount)
//
// Kept here in the fork so the gorm side of that comparison isn't lost.

func seedExtPlain(b *testing.B, n int) {
	DB.Exec("delete from pets")
	DB.Exec("delete from accounts")
	DB.Exec("delete from users")
	users := make([]*User, n)
	for i := range users {
		users[i] = GetUser(fmt.Sprintf("scan-%d", i), Config{})
	}
	if err := DB.CreateInBatches(&users, 100).Error; err != nil {
		b.Fatal(err)
	}
}

func seedExtWithAccount(b *testing.B, n int) {
	DB.Exec("delete from pets")
	DB.Exec("delete from accounts")
	DB.Exec("delete from users")
	users := make([]*User, n)
	for i := range users {
		users[i] = GetUser(fmt.Sprintf("scan-%d", i), Config{Account: true})
	}
	if err := DB.CreateInBatches(&users, 100).Error; err != nil {
		b.Fatal(err)
	}
}

func seedExtWithPets(b *testing.B, nParents, nChildren int) {
	DB.Exec("delete from pets")
	DB.Exec("delete from accounts")
	DB.Exec("delete from users")
	users := make([]*User, nParents)
	for i := range users {
		users[i] = GetUser(fmt.Sprintf("pu-%d", i), Config{Pets: nChildren})
	}
	if err := DB.CreateInBatches(&users, 50).Error; err != nil {
		b.Fatal(err)
	}
}

func benchExtFind5Where(b *testing.B, n int) {
	seedExtPlain(b, n)
	b.ResetTimer()
	b.ReportAllocs()
	for x := 0; x < b.N; x++ {
		var out []*User
		err := DB.
			Where("name <> ?", "nope1").
			Where("name <> ?", "nope2").
			Where("name <> ?", "nope3").
			Where("name <> ?", "nope4").
			Where("name <> ?", "nope5").
			Find(&out).Error
		if err != nil {
			b.Fatal(err)
		}
		if len(out) != n {
			b.Fatalf("expected %d rows, got %d", n, len(out))
		}
	}
}

func BenchmarkExtFind_5Where_1Row(b *testing.B)    { benchExtFind5Where(b, 1) }
func BenchmarkExtFind_5Where_100Rows(b *testing.B) { benchExtFind5Where(b, 100) }

func BenchmarkExtFirst(b *testing.B) {
	seedExtPlain(b, 1)
	var seed User
	DB.First(&seed)
	b.ResetTimer()
	b.ReportAllocs()
	for x := 0; x < b.N; x++ {
		var u User
		if err := DB.First(&u, seed.ID).Error; err != nil {
			b.Fatal(err)
		}
	}
}

func benchExtJoinedPreload(b *testing.B, n int) {
	seedExtWithAccount(b, n)
	b.ResetTimer()
	b.ReportAllocs()
	for x := 0; x < b.N; x++ {
		var out []*User
		if err := DB.Joins("Account").Find(&out).Error; err != nil {
			b.Fatal(err)
		}
		if len(out) != n {
			b.Fatalf("expected %d rows, got %d", n, len(out))
		}
	}
}

func BenchmarkExtJoinedPreload_100(b *testing.B) { benchExtJoinedPreload(b, 100) }
func BenchmarkExtJoinedPreload_1K(b *testing.B)  { benchExtJoinedPreload(b, 1000) }

func BenchmarkExtHasMany_50x4(b *testing.B) {
	seedExtWithPets(b, 50, 4)
	b.ResetTimer()
	b.ReportAllocs()
	for x := 0; x < b.N; x++ {
		var out []*User
		if err := DB.Preload("Pets").Find(&out).Error; err != nil {
			b.Fatal(err)
		}
		if len(out) != 50 {
			b.Fatalf("expected 50 users, got %d", len(out))
		}
	}
}
