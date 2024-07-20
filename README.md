# abx kubectl

A `kubectl` plugin to run commands based on custom attributes over fleets of clusters.

## Quick Start

```
kubectl krew install abx
kubectl abx -sa staging,us-east-1 -na ecommerce get pods -n monitoring 
```

`-sa` or `select attributes` is the first initial pass over your context attribute mapping, 
the execution will filter for all contexts that have the attributes `staging` and `us-east-1`. 

`-na` or `negate attributes` will apply a filter over your select and remove
any contexts that match this filter. Note a negation will always take precedence 
over a selection or in other words, negation always wins. 


Take the following `kube config`

```yaml
contexts:
- context:
  name: stag01
- context:
  name: prod02
- context:
  name: prod01
```

Attribute mappings are then defined in a yaml file for example

```yaml
---
prod01:
  - production
  - us-east-1
  - ecommerce
prod02:
  - production
  - us-east-1
  - data-analytics
stag01:
  - staging
  - us-east-1
  - ecommerce
stag02:
  - staging
  - us-east-1 
  - data-analytics
```

Based on our original query our `get pods -n monitoring` would be ran against `stag02`. 
