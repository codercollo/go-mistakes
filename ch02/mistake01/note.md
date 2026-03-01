# UNINTENDED VARIABLE SHADOWING

Shadowing - A variable is shadowed when a new variable with the same name is declared inside an inner block, hiding the outer variable. The outer variable still exists but cannot be reached
inside that block

# FIX - cleaner : Pre-declare err, use = not :=

var (
client \*http.Client
err error
)

if tracing {
client, err := createClientWithTracing()
} else {
client, err := createDefaultClient()
}

If a variable needs to survive outside a block, declare it outside
and assign with = inside. Never user := on it inside the block
