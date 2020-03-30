// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

func init() {
	RegisterDatasource("oci_database_autonomous_db_versions", DatabaseAutonomousDbVersionsDataSource())
}

func DatabaseAutonomousDbVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousDbVersions,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_workload": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"autonomous_db_versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"db_workload": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_dedicated": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatabaseAutonomousDbVersions(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDbVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

type DatabaseAutonomousDbVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousDbVersionsResponse
}

func (s *DatabaseAutonomousDbVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousDbVersionsDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousDbVersionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbWorkload, ok := s.D.GetOkExists("db_workload"); ok {
		request.DbWorkload = oci_database.AutonomousDatabaseSummaryDbWorkloadEnum(dbWorkload.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousDbVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutonomousDbVersions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousDbVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autonomousDbVersion := map[string]interface{}{}

		autonomousDbVersion["db_workload"] = r.DbWorkload

		if r.Details != nil {
			autonomousDbVersion["details"] = *r.Details
		}

		if r.IsDedicated != nil {
			autonomousDbVersion["is_dedicated"] = *r.IsDedicated
		}

		if r.Version != nil {
			autonomousDbVersion["version"] = *r.Version
		}

		resources = append(resources, autonomousDbVersion)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DatabaseAutonomousDbVersionsDataSource().Schema["autonomous_db_versions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_db_versions", resources); err != nil {
		return err
	}

	return nil
}