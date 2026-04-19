
# ADR-003: argocd-gitops

## Status
Accepted

## Context
- need a way to deploy new features consistently

## Decision
use argocd 

## Reasoning
- there is push vs pull continuous deployment
- push can lead to drift becaue monitoring is stopped after validating and pushing
- pull always reconciles back to source of truth, git
## Tradeoffs
- one more layer of tech
- might not want auto sync in prod like you might want human approval
## Consequences
- lose the confidence whats in cluster matches git