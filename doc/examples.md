# CSVSS Examples

## Basic Data Transformation

### Example 1: Adding Calculated Columns

**Input CSV (sales.csv):**
```csv
Product,Price,Quantity
Apple,1.20,10
Banana,0.80,15
Cherry,2.50,8
```

**Script (calculate.csvs):**
```
# Add header for total column
let D1 = "Total";

# Calculate total for each row
let D2 = B2 * C2;
let D3 = B3 * C3; 
let D4 = B4 * C4;

# Add a grand total row
let A5 = "TOTAL";
let D5 = SUM(D2:D4);
```

**Command:**
```bash
csvss -s calculate.csvs sales.csv
```

**Output:**
```csv
Product,Price,Quantity,Total
Apple,1.20,10,12.00
Banana,0.80,15,12.00
Cherry,2.50,8,20.00
TOTAL,,,44.00
```

## String Manipulation

### Example 2: Name Formatting

**Input CSV (names.csv):**
```csv
first_name,last_name,email
john,doe,john.doe@email.com
jane,smith,jane.smith@email.com
```

**Script (format_names.csvs):**
```
# Create formatted full name column
let C1 = "Full Name";
let C2 = UPPER(LEFT(A2, 1)) + LOWER(RIGHT(A2, LEN(A2)-1)) + " " + 
         UPPER(LEFT(B2, 1)) + LOWER(RIGHT(B2, LEN(B2)-1));
let C3 = UPPER(LEFT(A3, 1)) + LOWER(RIGHT(A3, LEN(A3)-1)) + " " + 
         UPPER(LEFT(B3, 1)) + LOWER(RIGHT(B3, LEN(B3)-1));

# Alternative simpler approach
let D1 = "Display Name";
let D2 = CONCATENATE(UPPER(A2), " ", UPPER(B2));
let D3 = CONCATENATE(UPPER(A3), " ", UPPER(B3));
```

**Output:**
```csv
first_name,last_name,email,Full Name,Display Name
john,doe,john.doe@email.com,John Doe,JOHN DOE
jane,smith,jane.smith@email.com,Jane Smith,JANE SMITH
```

## Conditional Logic

### Example 3: Grade Assignment

**Input CSV (scores.csv):**
```csv
Student,Math,Science,English
Alice,85,92,78
Bob,76,84,91
Carol,95,88,89
```

**Script (grades.csvs):**
```
# Add grade columns
let E1 = "Math Grade";
let F1 = "Science Grade"; 
let G1 = "English Grade";
let H1 = "Overall";

# Assign letter grades based on scores
let E2 = IF(B2 >= 90, "A", IF(B2 >= 80, "B", IF(B2 >= 70, "C", "F")));
let F2 = IF(C2 >= 90, "A", IF(C2 >= 80, "B", IF(C2 >= 70, "C", "F")));
let G2 = IF(D2 >= 90, "A", IF(D2 >= 80, "B", IF(D2 >= 70, "C", "F")));

let E3 = IF(B3 >= 90, "A", IF(B3 >= 80, "B", IF(B3 >= 70, "C", "F")));
let F3 = IF(C3 >= 90, "A", IF(C3 >= 80, "B", IF(C3 >= 70, "C", "F")));
let G3 = IF(D3 >= 90, "A", IF(D3 >= 80, "B", IF(D3 >= 70, "C", "F")));

let E4 = IF(B4 >= 90, "A", IF(B4 >= 80, "B", IF(B4 >= 70, "C", "F")));
let F4 = IF(C4 >= 90, "A", IF(C4 >= 80, "B", IF(C4 >= 70, "C", "F")));
let G4 = IF(D4 >= 90, "A", IF(D4 >= 80, "B", IF(D4 >= 70, "C", "F")));

# Calculate overall average and grade
let avg2 = AVERAGE(B2:D2);
let H2 = IF(avg2 >= 90, "A", IF(avg2 >= 80, "B", IF(avg2 >= 70, "C", "F")));

let avg3 = AVERAGE(B3:D3);
let H3 = IF(avg3 >= 90, "A", IF(avg3 >= 80, "B", IF(avg3 >= 70, "C", "F")));

let avg4 = AVERAGE(B4:D4);
let H4 = IF(avg4 >= 90, "A", IF(avg4 >= 80, "B", IF(avg4 >= 70, "C", "F")));
```

## Data Cleaning

### Example 4: Cleaning and Validating Data

**Input CSV (messy_data.csv):**
```csv
name,  age,  salary 
  Alice  ,25,"  50000  "
" Bob ",thirty,"75,000"
Charlie,35,
```

**Script (clean.csvs):**
```
# Clean up headers
let A1 = "Name";
let B1 = "Age"; 
let C1 = "Salary";

# Clean name column - trim spaces and proper case
let A2 = TRIM(A2);
let A3 = TRIM(A3);
let A4 = TRIM(A4);

# Clean age - convert text numbers, validate
let B2 = B2;  # Already numeric
let B3 = IF(B3 == "thirty", 30, 0);  # Convert text to number
let B4 = B4;  # Already numeric

# Clean salary - remove commas, spaces, convert to number
let C2 = 50000;  # Clean numeric value
let C3 = 75000;  # Remove comma and convert
let C4 = IF(C4 == "", 0, C4);  # Handle empty value

# Add validation column
let D1 = "Valid";
let D2 = IF(AND(LEN(A2) > 0, B2 > 0, C2 > 0), "Yes", "No");
let D3 = IF(AND(LEN(A3) > 0, B3 > 0, C3 > 0), "Yes", "No");
let D4 = IF(AND(LEN(A4) > 0, B4 > 0, C4 > 0), "Yes", "No");
```

## Range Operations

### Example 5: Working with Ranges

**Input CSV (quarterly_sales.csv):**
```csv
Month,Q1,Q2,Q3,Q4
Product A,1000,1200,1100,1300
Product B,800,900,950,1000
Product C,1500,1400,1600,1550
```

**Script (quarterly_analysis.csvs):**
```
# Add total column
let F1 = "Total";
let F2 = SUM(B2:E2);
let F3 = SUM(B3:E3);
let F4 = SUM(B4:E4);

# Add quarterly totals row
let A5 = "Quarter Total";
let B5 = SUM(B2:B4);
let C5 = SUM(C2:C4);
let D5 = SUM(D2:D4);
let E5 = SUM(E2:E4);
let F5 = SUM(F2:F4);

# Add average row
let A6 = "Quarter Average";
let B6 = AVERAGE(B2:B4);
let C6 = AVERAGE(C2:C4);
let D6 = AVERAGE(D2:D4);
let E6 = AVERAGE(E2:E4);
let F6 = AVERAGE(F2:F4);

# Calculate best performing quarter for each product
let G1 = "Best Quarter";
let best_q2 = IF(B2 >= C2 && B2 >= D2 && B2 >= E2, "Q1",
              IF(C2 >= D2 && C2 >= E2, "Q2",
              IF(D2 >= E2, "Q3", "Q4")));
let G2 = best_q2;
```

## Formatting Examples

### Example 6: Number Formatting

**Input CSV (financial.csv):**
```csv
Item,Amount
Revenue,123456.789
Expenses,98765.432
Profit,24691.357
```

**Script (format_financial.csvs):**
```
# Format amounts as currency with 2 decimal places
fmt B2 = "%.2f";
fmt B3 = "%.2f";
fmt B4 = "%.2f";

# Add percentage calculations
let C1 = "Percentage";
let total_revenue = B2;
let C2 = 100;  # Revenue is 100%
let C3 = (B3 / total_revenue) * 100;
let C4 = (B4 / total_revenue) * 100;

# Format percentages
fmt C2 = "%.1f%%";
fmt C3 = "%.1f%%";
fmt C4 = "%.1f%%";
```

## Multiple File Processing

### Example 7: Processing with Variables

**Script (advanced_calculations.csvs):**
```
# Define constants
let tax_rate = 0.08;
let discount_threshold = 1000;
let discount_rate = 0.05;

# Calculate subtotal
let subtotal = SUM(C2:C10);  # Assuming quantity * price in column C

# Apply volume discount
let discount = IF(subtotal > discount_threshold, 
                  subtotal * discount_rate, 
                  0);

# Calculate final amounts
let discounted_total = subtotal - discount;
let tax_amount = discounted_total * tax_rate;
let final_total = discounted_total + tax_amount;

# Output summary
let A12 = "Subtotal";
let B12 = subtotal;

let A13 = "Discount";
let B13 = discount;

let A14 = "Tax";
let B14 = tax_amount;

let A15 = "Total";
let B15 = final_total;

# Format currency
fmt B12:B15 = "%.2f";
```

## Command Line Usage Patterns

### Process and Save to New File
```bash
csvss -s transform.csvs input.csv -o output.csv
```

### Update File In-Place
```bash
csvss -s transform.csvs -u data.csv
```

### Pipe Processing
```bash
cat data.csv | csvss -s transform.csvs > processed.csv
```

### Batch Processing Multiple Files
```bash
for file in *.csv; do
    csvss -s common_transform.csvs -u "$file"
done
```