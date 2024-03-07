/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	common "github.com/upbound/provider-gcp/config/common"
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *FolderBucketConfig) ResolveReferences( // ResolveReferences of this FolderBucketConfig.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "Folder", "FolderList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Folder),
			Extract:      common.ExtractFolderID(),
			Reference:    mg.Spec.ForProvider.FolderRef,
			Selector:     mg.Spec.ForProvider.FolderSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Folder")
	}
	mg.Spec.ForProvider.Folder = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FolderExclusion.
func (mg *FolderExclusion) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "Folder", "FolderList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Folder),
			Extract:      resource.ExtractParamPath("name", true),
			Reference:    mg.Spec.ForProvider.FolderRef,
			Selector:     mg.Spec.ForProvider.FolderSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Folder")
	}
	mg.Spec.ForProvider.Folder = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FolderSink.
func (mg *FolderSink) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "Folder", "FolderList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Folder),
			Extract:      resource.ExtractParamPath("name", true),
			Reference:    mg.Spec.ForProvider.FolderRef,
			Selector:     mg.Spec.ForProvider.FolderSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Folder")
	}
	mg.Spec.ForProvider.Folder = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LogView.
func (mg *LogView) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("logging.gcp.upbound.io", "v1beta1", "ProjectBucketConfig", "ProjectBucketConfigList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.BucketRef,
			Selector:     mg.Spec.ForProvider.BucketSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Metric.
func (mg *Metric) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("logging.gcp.upbound.io", "v1beta1", "ProjectBucketConfig", "ProjectBucketConfigList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BucketName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.BucketNameRef,
			Selector:     mg.Spec.ForProvider.BucketNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BucketName")
	}
	mg.Spec.ForProvider.BucketName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("logging.gcp.upbound.io", "v1beta1", "ProjectBucketConfig", "ProjectBucketConfigList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BucketName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.BucketNameRef,
			Selector:     mg.Spec.InitProvider.BucketNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.BucketName")
	}
	mg.Spec.InitProvider.BucketName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BucketNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ProjectBucketConfig.
func (mg *ProjectBucketConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.CmekSettings); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("kms.gcp.upbound.io", "v1beta1", "CryptoKey", "CryptoKeyList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CmekSettings[i3].KMSKeyName),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.CmekSettings[i3].KMSKeyNameRef,
				Selector:     mg.Spec.ForProvider.CmekSettings[i3].KMSKeyNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CmekSettings[i3].KMSKeyName")
		}
		mg.Spec.ForProvider.CmekSettings[i3].KMSKeyName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.CmekSettings[i3].KMSKeyNameRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "Project", "ProjectList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Project),
			Extract:      resource.ExtractParamPath("project_id", false),
			Reference:    mg.Spec.ForProvider.ProjectRef,
			Selector:     mg.Spec.ForProvider.ProjectSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Project")
	}
	mg.Spec.ForProvider.Project = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProjectRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.CmekSettings); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("kms.gcp.upbound.io", "v1beta1", "CryptoKey", "CryptoKeyList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CmekSettings[i3].KMSKeyName),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.CmekSettings[i3].KMSKeyNameRef,
				Selector:     mg.Spec.InitProvider.CmekSettings[i3].KMSKeyNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.CmekSettings[i3].KMSKeyName")
		}
		mg.Spec.InitProvider.CmekSettings[i3].KMSKeyName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.CmekSettings[i3].KMSKeyNameRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this ProjectSink.
func (mg *ProjectSink) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomWriterIdentity),
			Extract:      resource.ExtractParamPath("email", true),
			Reference:    mg.Spec.ForProvider.CustomWriterIdentityRef,
			Selector:     mg.Spec.ForProvider.CustomWriterIdentitySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CustomWriterIdentity")
	}
	mg.Spec.ForProvider.CustomWriterIdentity = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CustomWriterIdentityRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CustomWriterIdentity),
			Extract:      resource.ExtractParamPath("email", true),
			Reference:    mg.Spec.InitProvider.CustomWriterIdentityRef,
			Selector:     mg.Spec.InitProvider.CustomWriterIdentitySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CustomWriterIdentity")
	}
	mg.Spec.InitProvider.CustomWriterIdentity = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CustomWriterIdentityRef = rsp.ResolvedReference

	return nil
}
