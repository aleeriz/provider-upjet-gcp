// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *JobTrigger) ResolveReferences( // ResolveReferences of this JobTrigger.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.InspectJob != nil {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.InspectJob.Actions); i4++ {
			if mg.Spec.ForProvider.InspectJob.Actions[i4].Deidentify != nil {
				if mg.Spec.ForProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig != nil {
					if mg.Spec.ForProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table != nil {
						{
							m, l, err = apisresolver.GetManagedResource("bigquery.gcp.upbound.io", "v1beta2", "Dataset", "DatasetList")
							if err != nil {
								return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
							}
							rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
								CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table.DatasetID),
								Extract:      reference.ExternalName(),
								Reference:    mg.Spec.ForProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table.DatasetIDRef,
								Selector:     mg.Spec.ForProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table.DatasetIDSelector,
								To:           reference.To{List: l, Managed: m},
							})
						}
						if err != nil {
							return errors.Wrap(err, "mg.Spec.ForProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table.DatasetID")
						}
						mg.Spec.ForProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table.DatasetID = reference.ToPtrValue(rsp.ResolvedValue)
						mg.Spec.ForProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table.DatasetIDRef = rsp.ResolvedReference

					}
				}
			}
		}
	}
	if mg.Spec.ForProvider.InspectJob != nil {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.InspectJob.Actions); i4++ {
			if mg.Spec.ForProvider.InspectJob.Actions[i4].Deidentify != nil {
				if mg.Spec.ForProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig != nil {
					if mg.Spec.ForProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table != nil {
						{
							m, l, err = apisresolver.GetManagedResource("bigquery.gcp.upbound.io", "v1beta2", "Table", "TableList")
							if err != nil {
								return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
							}
							rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
								CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table.TableID),
								Extract:      reference.ExternalName(),
								Reference:    mg.Spec.ForProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table.TableIDRef,
								Selector:     mg.Spec.ForProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table.TableIDSelector,
								To:           reference.To{List: l, Managed: m},
							})
						}
						if err != nil {
							return errors.Wrap(err, "mg.Spec.ForProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table.TableID")
						}
						mg.Spec.ForProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table.TableID = reference.ToPtrValue(rsp.ResolvedValue)
						mg.Spec.ForProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table.TableIDRef = rsp.ResolvedReference

					}
				}
			}
		}
	}
	if mg.Spec.InitProvider.InspectJob != nil {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.InspectJob.Actions); i4++ {
			if mg.Spec.InitProvider.InspectJob.Actions[i4].Deidentify != nil {
				if mg.Spec.InitProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig != nil {
					if mg.Spec.InitProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table != nil {
						{
							m, l, err = apisresolver.GetManagedResource("bigquery.gcp.upbound.io", "v1beta2", "Dataset", "DatasetList")
							if err != nil {
								return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
							}
							rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
								CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table.DatasetID),
								Extract:      reference.ExternalName(),
								Reference:    mg.Spec.InitProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table.DatasetIDRef,
								Selector:     mg.Spec.InitProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table.DatasetIDSelector,
								To:           reference.To{List: l, Managed: m},
							})
						}
						if err != nil {
							return errors.Wrap(err, "mg.Spec.InitProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table.DatasetID")
						}
						mg.Spec.InitProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table.DatasetID = reference.ToPtrValue(rsp.ResolvedValue)
						mg.Spec.InitProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table.DatasetIDRef = rsp.ResolvedReference

					}
				}
			}
		}
	}
	if mg.Spec.InitProvider.InspectJob != nil {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.InspectJob.Actions); i4++ {
			if mg.Spec.InitProvider.InspectJob.Actions[i4].Deidentify != nil {
				if mg.Spec.InitProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig != nil {
					if mg.Spec.InitProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table != nil {
						{
							m, l, err = apisresolver.GetManagedResource("bigquery.gcp.upbound.io", "v1beta2", "Table", "TableList")
							if err != nil {
								return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
							}
							rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
								CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table.TableID),
								Extract:      reference.ExternalName(),
								Reference:    mg.Spec.InitProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table.TableIDRef,
								Selector:     mg.Spec.InitProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table.TableIDSelector,
								To:           reference.To{List: l, Managed: m},
							})
						}
						if err != nil {
							return errors.Wrap(err, "mg.Spec.InitProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table.TableID")
						}
						mg.Spec.InitProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table.TableID = reference.ToPtrValue(rsp.ResolvedValue)
						mg.Spec.InitProvider.InspectJob.Actions[i4].Deidentify.TransformationDetailsStorageConfig.Table.TableIDRef = rsp.ResolvedReference

					}
				}
			}
		}
	}

	return nil
}
