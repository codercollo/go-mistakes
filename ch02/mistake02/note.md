# UNNECESSARY NESTED CODE

Nesting is the enemy of readabilty. Every extra level forces your brain to track one more context simultaneously

# RULES:

- If an if block returns, omit the else.
  // GOOD — else is unnecessary after a return
  if foo() {
  return true
  }

- Flip the condition for non-happy paths - handle errors/edge cases first
  // GOOD — non-happy path exits early, happy path stays left
  if s == "" {
  return errors.New("empty string")
  }
  // happy path...
  left column = what should happen(happy path)
  second column = what goes wrong(edge cases, errors)

- Handle edge cases early, return immediately, never use else after a return.
- Keep the happy path flowing down the left edge
