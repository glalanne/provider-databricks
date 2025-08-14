// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	customappintegration "github.com/glalanne/provider-databricks/internal/controller/cluster/apps/customappintegration"
	cluster "github.com/glalanne/provider-databricks/internal/controller/cluster/compute/cluster"
	clusterpolicy "github.com/glalanne/provider-databricks/internal/controller/cluster/compute/clusterpolicy"
	instancepool "github.com/glalanne/provider-databricks/internal/controller/cluster/compute/instancepool"
	job "github.com/glalanne/provider-databricks/internal/controller/cluster/compute/job"
	library "github.com/glalanne/provider-databricks/internal/controller/cluster/compute/library"
	pipeline "github.com/glalanne/provider-databricks/internal/controller/cluster/compute/pipeline"
	instanceprofile "github.com/glalanne/provider-databricks/internal/controller/cluster/deployment/instanceprofile"
	mwscredentials "github.com/glalanne/provider-databricks/internal/controller/cluster/deployment/mwscredentials"
	mwscustomermanagedkeys "github.com/glalanne/provider-databricks/internal/controller/cluster/deployment/mwscustomermanagedkeys"
	mwsnccbinding "github.com/glalanne/provider-databricks/internal/controller/cluster/deployment/mwsnccbinding"
	mwsnccprivateendpointrule "github.com/glalanne/provider-databricks/internal/controller/cluster/deployment/mwsnccprivateendpointrule"
	mwsnetworkconnectivityconfig "github.com/glalanne/provider-databricks/internal/controller/cluster/deployment/mwsnetworkconnectivityconfig"
	mwsnetworks "github.com/glalanne/provider-databricks/internal/controller/cluster/deployment/mwsnetworks"
	mwsprivateaccesssettings "github.com/glalanne/provider-databricks/internal/controller/cluster/deployment/mwsprivateaccesssettings"
	mwsstorageconfigurations "github.com/glalanne/provider-databricks/internal/controller/cluster/deployment/mwsstorageconfigurations"
	mwsvpcendpoint "github.com/glalanne/provider-databricks/internal/controller/cluster/deployment/mwsvpcendpoint"
	mwsworkspaces "github.com/glalanne/provider-databricks/internal/controller/cluster/deployment/mwsworkspaces"
	budget "github.com/glalanne/provider-databricks/internal/controller/cluster/finops/budget"
	mwslogdelivery "github.com/glalanne/provider-databricks/internal/controller/cluster/log/mwslogdelivery"
	mlflowexperiment "github.com/glalanne/provider-databricks/internal/controller/cluster/mlflow/mlflowexperiment"
	mlflowmodel "github.com/glalanne/provider-databricks/internal/controller/cluster/mlflow/mlflowmodel"
	vectorsearchendpoint "github.com/glalanne/provider-databricks/internal/controller/cluster/mosaic/vectorsearchendpoint"
	vectorsearchindex "github.com/glalanne/provider-databricks/internal/controller/cluster/mosaic/vectorsearchindex"
	providerconfig "github.com/glalanne/provider-databricks/internal/controller/cluster/providerconfig"
	accesscontrolruleset "github.com/glalanne/provider-databricks/internal/controller/cluster/security/accesscontrolruleset"
	entitlements "github.com/glalanne/provider-databricks/internal/controller/cluster/security/entitlements"
	group "github.com/glalanne/provider-databricks/internal/controller/cluster/security/group"
	groupinstanceprofile "github.com/glalanne/provider-databricks/internal/controller/cluster/security/groupinstanceprofile"
	groupmember "github.com/glalanne/provider-databricks/internal/controller/cluster/security/groupmember"
	grouprole "github.com/glalanne/provider-databricks/internal/controller/cluster/security/grouprole"
	ipaccesslist "github.com/glalanne/provider-databricks/internal/controller/cluster/security/ipaccesslist"
	mwspermissionassignment "github.com/glalanne/provider-databricks/internal/controller/cluster/security/mwspermissionassignment"
	obotoken "github.com/glalanne/provider-databricks/internal/controller/cluster/security/obotoken"
	permissionassignment "github.com/glalanne/provider-databricks/internal/controller/cluster/security/permissionassignment"
	permissions "github.com/glalanne/provider-databricks/internal/controller/cluster/security/permissions"
	secret "github.com/glalanne/provider-databricks/internal/controller/cluster/security/secret"
	secretacl "github.com/glalanne/provider-databricks/internal/controller/cluster/security/secretacl"
	secretscope "github.com/glalanne/provider-databricks/internal/controller/cluster/security/secretscope"
	serviceprincipal "github.com/glalanne/provider-databricks/internal/controller/cluster/security/serviceprincipal"
	serviceprincipalrole "github.com/glalanne/provider-databricks/internal/controller/cluster/security/serviceprincipalrole"
	serviceprincipalsecret "github.com/glalanne/provider-databricks/internal/controller/cluster/security/serviceprincipalsecret"
	sqlpermissions "github.com/glalanne/provider-databricks/internal/controller/cluster/security/sqlpermissions"
	token "github.com/glalanne/provider-databricks/internal/controller/cluster/security/token"
	user "github.com/glalanne/provider-databricks/internal/controller/cluster/security/user"
	userinstanceprofile "github.com/glalanne/provider-databricks/internal/controller/cluster/security/userinstanceprofile"
	userrole "github.com/glalanne/provider-databricks/internal/controller/cluster/security/userrole"
	modelserving "github.com/glalanne/provider-databricks/internal/controller/cluster/serving/modelserving"
	compliancesecurityprofileworkspacesetting "github.com/glalanne/provider-databricks/internal/controller/cluster/settings/compliancesecurityprofileworkspacesetting"
	defaultnamespacesetting "github.com/glalanne/provider-databricks/internal/controller/cluster/settings/defaultnamespacesetting"
	enhancedsecuritymonitoringworkspacesetting "github.com/glalanne/provider-databricks/internal/controller/cluster/settings/enhancedsecuritymonitoringworkspacesetting"
	restrictworkspaceadminssetting "github.com/glalanne/provider-databricks/internal/controller/cluster/settings/restrictworkspaceadminssetting"
	provider "github.com/glalanne/provider-databricks/internal/controller/cluster/sharing/provider"
	recipient "github.com/glalanne/provider-databricks/internal/controller/cluster/sharing/recipient"
	share "github.com/glalanne/provider-databricks/internal/controller/cluster/sharing/share"
	alert "github.com/glalanne/provider-databricks/internal/controller/cluster/sql/alert"
	dashboard "github.com/glalanne/provider-databricks/internal/controller/cluster/sql/dashboard"
	query "github.com/glalanne/provider-databricks/internal/controller/cluster/sql/query"
	sqlalert "github.com/glalanne/provider-databricks/internal/controller/cluster/sql/sqlalert"
	sqldashboard "github.com/glalanne/provider-databricks/internal/controller/cluster/sql/sqldashboard"
	sqlendpoint "github.com/glalanne/provider-databricks/internal/controller/cluster/sql/sqlendpoint"
	sqlglobalconfig "github.com/glalanne/provider-databricks/internal/controller/cluster/sql/sqlglobalconfig"
	sqlquery "github.com/glalanne/provider-databricks/internal/controller/cluster/sql/sqlquery"
	sqlvisualization "github.com/glalanne/provider-databricks/internal/controller/cluster/sql/sqlvisualization"
	sqlwidget "github.com/glalanne/provider-databricks/internal/controller/cluster/sql/sqlwidget"
	dbfsfile "github.com/glalanne/provider-databricks/internal/controller/cluster/storage/dbfsfile"
	file "github.com/glalanne/provider-databricks/internal/controller/cluster/storage/file"
	mount "github.com/glalanne/provider-databricks/internal/controller/cluster/storage/mount"
	artifactallowlist "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/artifactallowlist"
	catalog "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/catalog"
	catalogworkspacebinding "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/catalogworkspacebinding"
	connection "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/connection"
	credential "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/credential"
	externallocation "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/externallocation"
	grant "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/grant"
	grantmap "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/grantmap"
	lakehousemonitor "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/lakehousemonitor"
	metastore "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/metastore"
	metastoreassignment "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/metastoreassignment"
	metastoredataaccess "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/metastoredataaccess"
	onlinetable "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/onlinetable"
	qualitymonitor "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/qualitymonitor"
	registeredmodel "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/registeredmodel"
	schema "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/schema"
	sqltable "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/sqltable"
	storagecredential "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/storagecredential"
	systemschema "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/systemschema"
	volume "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/volume"
	workspacebinding "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/workspacebinding"
	directory "github.com/glalanne/provider-databricks/internal/controller/cluster/workspace/directory"
	gitcredential "github.com/glalanne/provider-databricks/internal/controller/cluster/workspace/gitcredential"
	globalinitscript "github.com/glalanne/provider-databricks/internal/controller/cluster/workspace/globalinitscript"
	notebook "github.com/glalanne/provider-databricks/internal/controller/cluster/workspace/notebook"
	notificationdestination "github.com/glalanne/provider-databricks/internal/controller/cluster/workspace/notificationdestination"
	repo "github.com/glalanne/provider-databricks/internal/controller/cluster/workspace/repo"
	workspaceconf "github.com/glalanne/provider-databricks/internal/controller/cluster/workspace/workspaceconf"
	workspacefile "github.com/glalanne/provider-databricks/internal/controller/cluster/workspace/workspacefile"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		customappintegration.Setup,
		cluster.Setup,
		clusterpolicy.Setup,
		instancepool.Setup,
		job.Setup,
		library.Setup,
		pipeline.Setup,
		instanceprofile.Setup,
		mwscredentials.Setup,
		mwscustomermanagedkeys.Setup,
		mwsnccbinding.Setup,
		mwsnccprivateendpointrule.Setup,
		mwsnetworkconnectivityconfig.Setup,
		mwsnetworks.Setup,
		mwsprivateaccesssettings.Setup,
		mwsstorageconfigurations.Setup,
		mwsvpcendpoint.Setup,
		mwsworkspaces.Setup,
		budget.Setup,
		mwslogdelivery.Setup,
		mlflowexperiment.Setup,
		mlflowmodel.Setup,
		vectorsearchendpoint.Setup,
		vectorsearchindex.Setup,
		providerconfig.Setup,
		accesscontrolruleset.Setup,
		entitlements.Setup,
		group.Setup,
		groupinstanceprofile.Setup,
		groupmember.Setup,
		grouprole.Setup,
		ipaccesslist.Setup,
		mwspermissionassignment.Setup,
		obotoken.Setup,
		permissionassignment.Setup,
		permissions.Setup,
		secret.Setup,
		secretacl.Setup,
		secretscope.Setup,
		serviceprincipal.Setup,
		serviceprincipalrole.Setup,
		serviceprincipalsecret.Setup,
		sqlpermissions.Setup,
		token.Setup,
		user.Setup,
		userinstanceprofile.Setup,
		userrole.Setup,
		modelserving.Setup,
		compliancesecurityprofileworkspacesetting.Setup,
		defaultnamespacesetting.Setup,
		enhancedsecuritymonitoringworkspacesetting.Setup,
		restrictworkspaceadminssetting.Setup,
		provider.Setup,
		recipient.Setup,
		share.Setup,
		alert.Setup,
		dashboard.Setup,
		query.Setup,
		sqlalert.Setup,
		sqldashboard.Setup,
		sqlendpoint.Setup,
		sqlglobalconfig.Setup,
		sqlquery.Setup,
		sqlvisualization.Setup,
		sqlwidget.Setup,
		dbfsfile.Setup,
		file.Setup,
		mount.Setup,
		artifactallowlist.Setup,
		catalog.Setup,
		catalogworkspacebinding.Setup,
		connection.Setup,
		credential.Setup,
		externallocation.Setup,
		grant.Setup,
		grantmap.Setup,
		lakehousemonitor.Setup,
		metastore.Setup,
		metastoreassignment.Setup,
		metastoredataaccess.Setup,
		onlinetable.Setup,
		qualitymonitor.Setup,
		registeredmodel.Setup,
		schema.Setup,
		sqltable.Setup,
		storagecredential.Setup,
		systemschema.Setup,
		volume.Setup,
		workspacebinding.Setup,
		directory.Setup,
		gitcredential.Setup,
		globalinitscript.Setup,
		notebook.Setup,
		notificationdestination.Setup,
		repo.Setup,
		workspaceconf.Setup,
		workspacefile.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		customappintegration.SetupGated,
		cluster.SetupGated,
		clusterpolicy.SetupGated,
		instancepool.SetupGated,
		job.SetupGated,
		library.SetupGated,
		pipeline.SetupGated,
		instanceprofile.SetupGated,
		mwscredentials.SetupGated,
		mwscustomermanagedkeys.SetupGated,
		mwsnccbinding.SetupGated,
		mwsnccprivateendpointrule.SetupGated,
		mwsnetworkconnectivityconfig.SetupGated,
		mwsnetworks.SetupGated,
		mwsprivateaccesssettings.SetupGated,
		mwsstorageconfigurations.SetupGated,
		mwsvpcendpoint.SetupGated,
		mwsworkspaces.SetupGated,
		budget.SetupGated,
		mwslogdelivery.SetupGated,
		mlflowexperiment.SetupGated,
		mlflowmodel.SetupGated,
		vectorsearchendpoint.SetupGated,
		vectorsearchindex.SetupGated,
		providerconfig.SetupGated,
		accesscontrolruleset.SetupGated,
		entitlements.SetupGated,
		group.SetupGated,
		groupinstanceprofile.SetupGated,
		groupmember.SetupGated,
		grouprole.SetupGated,
		ipaccesslist.SetupGated,
		mwspermissionassignment.SetupGated,
		obotoken.SetupGated,
		permissionassignment.SetupGated,
		permissions.SetupGated,
		secret.SetupGated,
		secretacl.SetupGated,
		secretscope.SetupGated,
		serviceprincipal.SetupGated,
		serviceprincipalrole.SetupGated,
		serviceprincipalsecret.SetupGated,
		sqlpermissions.SetupGated,
		token.SetupGated,
		user.SetupGated,
		userinstanceprofile.SetupGated,
		userrole.SetupGated,
		modelserving.SetupGated,
		compliancesecurityprofileworkspacesetting.SetupGated,
		defaultnamespacesetting.SetupGated,
		enhancedsecuritymonitoringworkspacesetting.SetupGated,
		restrictworkspaceadminssetting.SetupGated,
		provider.SetupGated,
		recipient.SetupGated,
		share.SetupGated,
		alert.SetupGated,
		dashboard.SetupGated,
		query.SetupGated,
		sqlalert.SetupGated,
		sqldashboard.SetupGated,
		sqlendpoint.SetupGated,
		sqlglobalconfig.SetupGated,
		sqlquery.SetupGated,
		sqlvisualization.SetupGated,
		sqlwidget.SetupGated,
		dbfsfile.SetupGated,
		file.SetupGated,
		mount.SetupGated,
		artifactallowlist.SetupGated,
		catalog.SetupGated,
		catalogworkspacebinding.SetupGated,
		connection.SetupGated,
		credential.SetupGated,
		externallocation.SetupGated,
		grant.SetupGated,
		grantmap.SetupGated,
		lakehousemonitor.SetupGated,
		metastore.SetupGated,
		metastoreassignment.SetupGated,
		metastoredataaccess.SetupGated,
		onlinetable.SetupGated,
		qualitymonitor.SetupGated,
		registeredmodel.SetupGated,
		schema.SetupGated,
		sqltable.SetupGated,
		storagecredential.SetupGated,
		systemschema.SetupGated,
		volume.SetupGated,
		workspacebinding.SetupGated,
		directory.SetupGated,
		gitcredential.SetupGated,
		globalinitscript.SetupGated,
		notebook.SetupGated,
		notificationdestination.SetupGated,
		repo.SetupGated,
		workspaceconf.SetupGated,
		workspacefile.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
