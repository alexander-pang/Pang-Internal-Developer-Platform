
# ADR-002: go control api

## Status
Accepted

## Context
We need the go api

## Decision
use go api

## Reasoning
-  the Go API adds a place for business logic and aggregation that plugins can't easily do                            
- when backstage plugins aren't consistent
## Tradeoffs
- extra layer to build, deploy and mainain
- mostly just proxying data, which doesn't justify the layer
## Consequences
- if backstage plugins covered everything needed, wouldn't need this