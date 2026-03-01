# MISTAKE-04 Overusing Getters and Setters

# Getters and Setters

- Methods that wrap access to unexported struct fields
- Common in Java/C# but not in idiomatic Go
- Go's standard library itself exposes fields directly eg: timer.C

# The Mistake - Getters/Setters With Zero Value

type Customer struct {
id int
balance float64
}

func (c *Customer) GetId() int { return c.id } // adds nothing
func (c *Customer) SetId(id int) { c.id = id } // adds nothing
func (c \*Customer) GetBalance() float64 { return c.
balance } // adds nothing

-Two problems:

1. Wrong Naming - Go convention is Balance() not GetBalance()
2. No behavior added - pure boilerplate, just wrapping field access

# FIX: Export Fields Directly

type Customer struct {
ID int
Name string
Balance float64
}

// Direct access — clean and idiomatic
c.Balance = 100.0
fmt.Println(c.Balance)

-when getters/setters add no behavior,just export the field

# Go Naming Convention (When you Do use them)

// Getter — named after the field, NO "Get" prefix
func (c \*Customer) Balance() float64 { return c.balance }

// Setter — "Set" + field name
func (c \*Customer) SetBalance(amount float64) { c.balance = amount }

- Usage:
  currentBalance := customer.Balance() // NOT customer.GetBalance()
  customer.SetBalance(0)

# When Getters/Setters are worth using

1. Validation logic in setter:
   func (c \*Customer) SetBalance(amount float64) {
   if amount < 0 {
   amount = 0 // setter guards against invalid state
   }
   c.balance = amount
   }

2. Mutex wrapping for concurrenct access:
   func (s \*SafeCounter) Count() int {
   s.mu.Lock()
   defer s.mu.Unlock()
   return s.count
   }

3. Computed/derived value:
   func (c _Customer) BalanceWithTax() float64 {
   return c.balance _ 1.16 // derived — not stored, computed on access
   }

4. Future-proofing :
   If you know the field will need validation or logic later, a getter/ setter now avoids a breaking API change later

# Core Rule

Don't add getters and setters just because other languages demand it Export fields directly for simple structs.
Only add them when they encapsulate real behavior
