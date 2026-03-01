# MISUSING INIT FUNCTIONS

init function:

- Takes no arguments, returns no result :func init()
- Runs automatically when the package is initialized
- Runs after all variable and constant declarations in the package
- Cannot be called directly - init() in code is a compile error
- You can have multiple init functions per file or per package
- Multiple files: alphabetical order of filenames
- Multiple in same file: source order

Execution Order:
dependency packages vars -> dependency package init() -> main package vars ->
main package init() -> main()

# 1 Mistake: USING init for DB connection

var db \*sql.DB // global variable

func init() {
d, err := sql.Open("mysql", os.Getenv("DSN"))
if err != nil {
log.Panic(err) // only option is panic
}
db = d
}

- Limited error handling init can't return an error. Only option is panic or
  log.Fatal. The caller(another package, a test) has no chance to retry, fallback or
  handle the error gracefully.

- Break unit tests init runs before ALL tests, even unit tests that have nothing to do with the database. Every test file in the package now requires a live DB connection just to run.

- Forces global variables any function in the package can read or mutate db
  Functions depending on db can't be isolated in tests. Global state - hidden dependency = fragile code

# FIX :Use Plain Function

func createClient(dsn string) (\*sql.DB, error) {
db, err := sql.Open("mysql", dsn)
if err != nil {
return nil, err // caller decides what to do
}
if err = db.Ping(); err != nil {
return nil, err
}
return db, nil // no global, encapsulated
}

- Caller handles thee error - retry, fallback, fatal
- Unit tests don't rtigger DB connection unless they call ths function
- db is passed where needed, no hidden global dependency
- Easy to integrate test this function in isolation

# When init is used/acceptable

func init() {
redirect := func(w http.ResponseWriter, r \*http.Request) {
http.Redirect(w, r, "/", http.StatusFound)
}
http.HandleFunc("/blog", redirect)
http.HandleFunc("/blog/", redirect)
}

- Cannot fail (http.HandlerFunc) only panics for nil handler
- No global variables being set
- No external dependencies (no DB no network)
- No impact on unit tests
- Pure static config

# Core Rule

Use init only for static configuration that cannot fail and has no
external dependencies. For everything else, i.e DB, Redis, or I/O use
a plain function and return the error
