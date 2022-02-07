module llamalauncher

go 1.16

// It will may get public soon. but it's private atm.
replace github.com/er-azh/fortnitemanifest v0.1.0 => ../fortnitemanifest

require (
	entgo.io/ent v0.10.0
	github.com/LlamaNite/llamalog v0.0.0-20210810202603-25f8312a8afd
	github.com/er-azh/fortnitemanifest v0.1.0
	github.com/mattn/go-sqlite3 v1.14.11
	golang.org/x/sys v0.0.0-20220204135822-1c1b9b1eba6a
)
