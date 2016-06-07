## 0.2.1

FEATURES

  * Added configurable timeout to find gateway.

IMPROVEMENTS

  * Aligned closer with terraform plugin operations. 

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
