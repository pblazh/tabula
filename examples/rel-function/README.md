# REL Function Example

This example demonstrates the REL (relative reference) function for creating dynamic cell references.

## What REL does

REL creates relative cell references based on the target cell being assigned to:

- `REL(1, 0)` - 1 column right, same row
- `REL(-1, 0)` - 1 column left, same row
- `REL(0, -1)` - same column, 1 row up
- `REL(2, 1)` - 2 columns right, 1 row down
