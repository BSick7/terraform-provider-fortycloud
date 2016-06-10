## 0.2.7

FEATURES

  * Added `fortycloud_registration_token` resource.
  * Added `fortycloud_gateway.open_vpn_client_cidrs` to be able to set "VPN Client Routes".

IMPROVEMENTS

  * Setting `fortycloud_subnet.gateway_id` to "" will clear the assignment from the subnet.

## 0.2.6

IMPROVEMENTS

  * Not attempting to assign gateway/resource group to subnet if none specified.
  * Matching legacy format for subnet gatewayRef.

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
