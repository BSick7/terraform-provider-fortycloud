## 0.2.5

IMPROVEMENTS

  * Gracefully handling missing gateways, subnets, and connections.

## 0.2.4

IMPROVEMENTS

  * Aligned subnet resource with terraform plugin operations.

BUG FIXES

  * Subnet create/update omits `gatewayRef` and `resourceGroupRef`.

## 0.2.3

BUG FIXES

  * Gateway update no longer crashes. Omitting `identityServerName` and `release`.

## 0.2.2

BUG FIXES

  * Gateway fields are now editable.

## 0.2.1

FEATURES

  * Added configurable timeout to find gateway.

IMPROVEMENTS

  * Aligned gateway resource with terraform plugin operations. 

## 0.2.0

INCOMPATIBILITIES

  * fortycloud_subnet
    * nat_disabled      -> disable_auto_nat
    * subnet            -> cidr
    * gateway_id        -> assigned in separate REST call; parsed from gatewayRef
    * resource_group_id -> assigned in separate REST call; parsed from ResourceGroupRef
    
  * fortycloud_node (replaced with fortycloud_gateway)

IMPROVEMENTS

  * Running against REST API
