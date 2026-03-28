# Chargeback Match

Reconciles chargeback records against sales records using `invoice_date`, `invoice_number`, and `program_id` as the composite match key.

## Running

```bash
go run ./cmd/main.go
```

## Testing

```bash
go test ./...
```

## Project Structure

```
match/
├── cmd/
│   └── main.go                     # Entry point
└── internal/
    ├── model/
    │   └── model.go                # ChargeBack and Sale structs
    ├── repository/
    │   ├── factory.go              # Creates repos by type ("csv", ...)
    │   ├── chargebacks.go          # ChargebackRepo interface + CSV impl
    │   └── sales.go                # SalesRepo interface + CSV impl
    └── service/
        ├── dataloader.go           # Loader interface — loads chargebacks and sales via repo
        ├── match.go                # Match struct, runMatch logic, PerformMatch orchestrator
        └── match_test.go           # Table-driven tests for matching logic
```

### Layer responsibilities

| Layer | Package | Responsibility |
|-------|---------|----------------|
| Entry point | `cmd` | Wires context and calls `PerformMatch` |
| Service | `internal/service` | Matching logic, orchestration |
| Repository | `internal/repository` | Data loading, one impl per source type |
| Model | `internal/model` | Shared data structures |

## How Matching Works

1. All chargebacks are indexed by a composite key: `DATE_INVOICENUMBER_PROGRAMID` (uppercased).
2. Each sale is looked up against the index.
3. A `Match` is produced for every chargeback that has a corresponding sale.

Unmatched chargebacks and unmatched sales are logged but not returned.

## Extensibility

### Adding a new data source (e.g. database, API)

The repository layer uses interfaces, so a new source only requires:

1. Implement `ChargebackRepo` in a new file (e.g. `repository/chargebacks_postgres.go`):

```go
type pgChargeBackRepo struct{ db *sql.DB }

func (r pgChargeBackRepo) ChargeBacks(ctx context.Context) ([]model.ChargeBack, error) {
    // query db
}
```

2. Register it in `factory.go`:

```go
case "postgres":
    return newPgChargeBackRepo(db), nil
```

3. Pass `"postgres"` when constructing the loader:

```go
loader := service.NewLoader("postgres")
```

No changes needed in the service or match logic.

### Changing the match key

The composite key is built in a single function in `match.go`:

```go
func matchKey(date, invoiceNumber, programID string) string {
    return strings.ToUpper(date) + "_" + strings.ToUpper(invoiceNumber) + "_" + strings.ToUpper(programID)
}
```

To add or remove fields from the key (e.g. include `dea`), update this function and the two call sites in `runMatch`.

### Adding match output (e.g. write results to file or DB)

`runMatch` returns `[]Match`, where each `Match` carries the full `ChargeBack` and `Sale` structs. `PerformMatch` in `match.go` iterates the results — add any output logic there (file write, DB insert, HTTP call) without touching the matching logic itself.
