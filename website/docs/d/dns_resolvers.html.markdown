---
subcategory: "Dns"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_resolvers"
sidebar_current: "docs-oci-datasource-dns-resolvers"
description: |-
  Provides the list of Resolvers in Oracle Cloud Infrastructure Dns service
---

# Data Source: oci_dns_resolvers
This data source provides the list of Resolvers in Oracle Cloud Infrastructure Dns service.

Gets a list of all resolvers within a compartment. The collection can
be filtered by display name, id, or lifecycle state. It can be sorted
on creation time or displayName both in ASC or DESC order. Note that
when no lifecycleState query parameter is provided that the collection
does not include resolvers in the DELETED lifecycleState to be consistent
with other operations of the API.


## Example Usage

```hcl
data "oci_dns_resolvers" "test_resolvers" {
	#Required
	compartment_id = var.compartment_id
	scope = "PRIVATE"

	#Optional
	display_name = var.resolver_display_name
	id = var.resolver_id
	state = var.resolver_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment the resource belongs to.
* `display_name` - (Optional) The displayName of a resource.
* `id` - (Optional) The OCID of a resource.
* `scope` - (Required) Value must be `PRIVATE` when listing private name resolvers.
* `state` - (Optional) The state of a resource.


## Attributes Reference

The following attributes are exported:

* `resolvers` - The list of resolvers.

### Resolver Reference

The following attributes are exported:

* `attached_vcn_id` - The OCID of the attached VCN. 
* `attached_views` - The attached views. Views are evaluated in order.
	* `view_id` - The OCID of the view.
* `compartment_id` - The OCID of the owning compartment.
* `default_view_id` - The OCID of the default view. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Operations": {"CostCenter": "42"}}` 
* `display_name` - The display name of the resolver. 
* `endpoints` - Read-only array of endpoints for the resolver. 
	* `compartment_id` - The OCID of the owning compartment. This will match the resolver that the resolver endpoint is under and will be updated if the resolver's compartment is changed. 
	* `endpoint_type` - The type of resolver endpoint. VNIC is currently the only supported type. 
	* `forwarding_address` - An IP address from which forwarded queries may be sent. For VNIC endpoints, this IP address must be part of the subnet and will be assigned by the system if unspecified when isForwarding is true. 
	* `is_forwarding` - A Boolean flag indicating whether or not the resolver endpoint is for forwarding. 
	* `is_listening` - A Boolean flag indicating whether or not the resolver endpoint is for listening. 
	* `listening_address` - An IP address to listen to queries on. For VNIC endpoints this IP address must be part of the subnet and will be assigned by the system if unspecified. 
	* `name` - The name of the resolver endpoint. Must be unique within the resolver. 
	* `self` - The canonical absolute URL of the resource.
	* `state` - The current state of the resource.
	* `subnet_id` - The OCID of a subnet. Must be part of the VCN that the resolver is attached to.
	* `time_created` - The date and time the resource was created in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.

		**Example:** `2016-07-22T17:23:59:60Z` 
	* `time_updated` - The date and time the resource was last updated in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.

		**Example:** `2016-07-22T17:23:59:60Z` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Department": "Finance"}` 
* `id` - The OCID of the resolver.
* `is_protected` - A Boolean flag indicating whether or not parts of the resource are unable to be explicitly managed. 
* `rules` - Rules for the resolver. Rules are evaluated in order. 
	* `action` - The action determines the behavior of the rule. If a query matches a supplied condition then the action will apply. If there are no conditions on the rule then all queries are subject to the specified action.
		* `FORWARD` - Matching requests will be forwarded from the source interface to the destination address. 
	* `client_address_conditions` - A list of CIDR blocks. The query must come from a client within one of the blocks in order for the rule action to apply. 
	* `destination_addresses` - IP addresses to which queries should be forwarded. Currently limited to a single address. 
	* `qname_cover_conditions` - A list of domain names. The query must be covered by one of the domains in order for the rule action to apply. 
	* `source_endpoint_name` - Name of an endpoint, that is a sub-resource of the resolver, to use as the forwarding interface. The endpoint must have isForwarding set to true. 
* `self` - The canonical absolute URL of the resource.
* `state` - The current state of the resource.
* `time_created` - The date and time the resource was created in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.

	**Example:** `2016-07-22T17:23:59:60Z` 
* `time_updated` - The date and time the resource was last updated in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.

	**Example:** `2016-07-22T17:23:59:60Z` 

