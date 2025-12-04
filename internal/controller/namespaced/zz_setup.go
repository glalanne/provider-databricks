// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	app "github.com/glalanne/provider-databricks/internal/controller/namespaced/apps/app"
	customappintegration "github.com/glalanne/provider-databricks/internal/controller/namespaced/apps/customappintegration"
	cluster "github.com/glalanne/provider-databricks/internal/controller/namespaced/compute/cluster"
	clusterpolicy "github.com/glalanne/provider-databricks/internal/controller/namespaced/compute/clusterpolicy"
	instancepool "github.com/glalanne/provider-databricks/internal/controller/namespaced/compute/instancepool"
	job "github.com/glalanne/provider-databricks/internal/controller/namespaced/compute/job"
	library "github.com/glalanne/provider-databricks/internal/controller/namespaced/compute/library"
	pipeline "github.com/glalanne/provider-databricks/internal/controller/namespaced/compute/pipeline"
	instanceprofile "github.com/glalanne/provider-databricks/internal/controller/namespaced/deployment/instanceprofile"
	mwscredentials "github.com/glalanne/provider-databricks/internal/controller/namespaced/deployment/mwscredentials"
	mwscustomermanagedkeys "github.com/glalanne/provider-databricks/internal/controller/namespaced/deployment/mwscustomermanagedkeys"
	mwsnccbinding "github.com/glalanne/provider-databricks/internal/controller/namespaced/deployment/mwsnccbinding"
	mwsnccprivateendpointrule "github.com/glalanne/provider-databricks/internal/controller/namespaced/deployment/mwsnccprivateendpointrule"
	mwsnetworkconnectivityconfig "github.com/glalanne/provider-databricks/internal/controller/namespaced/deployment/mwsnetworkconnectivityconfig"
	mwsnetworks "github.com/glalanne/provider-databricks/internal/controller/namespaced/deployment/mwsnetworks"
	mwsprivateaccesssettings "github.com/glalanne/provider-databricks/internal/controller/namespaced/deployment/mwsprivateaccesssettings"
	mwsstorageconfigurations "github.com/glalanne/provider-databricks/internal/controller/namespaced/deployment/mwsstorageconfigurations"
	mwsvpcendpoint "github.com/glalanne/provider-databricks/internal/controller/namespaced/deployment/mwsvpcendpoint"
	mwsworkspaces "github.com/glalanne/provider-databricks/internal/controller/namespaced/deployment/mwsworkspaces"
	budget "github.com/glalanne/provider-databricks/internal/controller/namespaced/finops/budget"
	mwslogdelivery "github.com/glalanne/provider-databricks/internal/controller/namespaced/log/mwslogdelivery"
	mlflowexperiment "github.com/glalanne/provider-databricks/internal/controller/namespaced/mlflow/mlflowexperiment"
	mlflowmodel "github.com/glalanne/provider-databricks/internal/controller/namespaced/mlflow/mlflowmodel"
	vectorsearchendpoint "github.com/glalanne/provider-databricks/internal/controller/namespaced/mosaic/vectorsearchendpoint"
	vectorsearchindex "github.com/glalanne/provider-databricks/internal/controller/namespaced/mosaic/vectorsearchindex"
	accountfederationpolicy "github.com/glalanne/provider-databricks/internal/controller/namespaced/oauth/accountfederationpolicy"
	serviceprincipalfederationpolicy "github.com/glalanne/provider-databricks/internal/controller/namespaced/oauth/serviceprincipalfederationpolicy"
	providerconfig "github.com/glalanne/provider-databricks/internal/controller/namespaced/providerconfig"
	accesscontrolruleset "github.com/glalanne/provider-databricks/internal/controller/namespaced/security/accesscontrolruleset"
	entitlements "github.com/glalanne/provider-databricks/internal/controller/namespaced/security/entitlements"
	group "github.com/glalanne/provider-databricks/internal/controller/namespaced/security/group"
	groupinstanceprofile "github.com/glalanne/provider-databricks/internal/controller/namespaced/security/groupinstanceprofile"
	groupmember "github.com/glalanne/provider-databricks/internal/controller/namespaced/security/groupmember"
	grouprole "github.com/glalanne/provider-databricks/internal/controller/namespaced/security/grouprole"
	ipaccesslist "github.com/glalanne/provider-databricks/internal/controller/namespaced/security/ipaccesslist"
	mwspermissionassignment "github.com/glalanne/provider-databricks/internal/controller/namespaced/security/mwspermissionassignment"
	obotoken "github.com/glalanne/provider-databricks/internal/controller/namespaced/security/obotoken"
	permissionassignment "github.com/glalanne/provider-databricks/internal/controller/namespaced/security/permissionassignment"
	permissions "github.com/glalanne/provider-databricks/internal/controller/namespaced/security/permissions"
	secret "github.com/glalanne/provider-databricks/internal/controller/namespaced/security/secret"
	secretacl "github.com/glalanne/provider-databricks/internal/controller/namespaced/security/secretacl"
	secretscope "github.com/glalanne/provider-databricks/internal/controller/namespaced/security/secretscope"
	serviceprincipal "github.com/glalanne/provider-databricks/internal/controller/namespaced/security/serviceprincipal"
	serviceprincipalrole "github.com/glalanne/provider-databricks/internal/controller/namespaced/security/serviceprincipalrole"
	serviceprincipalsecret "github.com/glalanne/provider-databricks/internal/controller/namespaced/security/serviceprincipalsecret"
	sqlpermissions "github.com/glalanne/provider-databricks/internal/controller/namespaced/security/sqlpermissions"
	token "github.com/glalanne/provider-databricks/internal/controller/namespaced/security/token"
	user "github.com/glalanne/provider-databricks/internal/controller/namespaced/security/user"
	userinstanceprofile "github.com/glalanne/provider-databricks/internal/controller/namespaced/security/userinstanceprofile"
	userrole "github.com/glalanne/provider-databricks/internal/controller/namespaced/security/userrole"
	modelserving "github.com/glalanne/provider-databricks/internal/controller/namespaced/serving/modelserving"
	compliancesecurityprofileworkspacesetting "github.com/glalanne/provider-databricks/internal/controller/namespaced/settings/compliancesecurityprofileworkspacesetting"
	defaultnamespacesetting "github.com/glalanne/provider-databricks/internal/controller/namespaced/settings/defaultnamespacesetting"
	enhancedsecuritymonitoringworkspacesetting "github.com/glalanne/provider-databricks/internal/controller/namespaced/settings/enhancedsecuritymonitoringworkspacesetting"
	restrictworkspaceadminssetting "github.com/glalanne/provider-databricks/internal/controller/namespaced/settings/restrictworkspaceadminssetting"
	provider "github.com/glalanne/provider-databricks/internal/controller/namespaced/sharing/provider"
	recipient "github.com/glalanne/provider-databricks/internal/controller/namespaced/sharing/recipient"
	share "github.com/glalanne/provider-databricks/internal/controller/namespaced/sharing/share"
	alert "github.com/glalanne/provider-databricks/internal/controller/namespaced/sql/alert"
	dashboard "github.com/glalanne/provider-databricks/internal/controller/namespaced/sql/dashboard"
	query "github.com/glalanne/provider-databricks/internal/controller/namespaced/sql/query"
	sqlalert "github.com/glalanne/provider-databricks/internal/controller/namespaced/sql/sqlalert"
	sqldashboard "github.com/glalanne/provider-databricks/internal/controller/namespaced/sql/sqldashboard"
	sqlendpoint "github.com/glalanne/provider-databricks/internal/controller/namespaced/sql/sqlendpoint"
	sqlglobalconfig "github.com/glalanne/provider-databricks/internal/controller/namespaced/sql/sqlglobalconfig"
	sqlquery "github.com/glalanne/provider-databricks/internal/controller/namespaced/sql/sqlquery"
	sqlvisualization "github.com/glalanne/provider-databricks/internal/controller/namespaced/sql/sqlvisualization"
	sqlwidget "github.com/glalanne/provider-databricks/internal/controller/namespaced/sql/sqlwidget"
	dbfsfile "github.com/glalanne/provider-databricks/internal/controller/namespaced/storage/dbfsfile"
	file "github.com/glalanne/provider-databricks/internal/controller/namespaced/storage/file"
	mount "github.com/glalanne/provider-databricks/internal/controller/namespaced/storage/mount"
	artifactallowlist "github.com/glalanne/provider-databricks/internal/controller/namespaced/unity/artifactallowlist"
	catalog "github.com/glalanne/provider-databricks/internal/controller/namespaced/unity/catalog"
	catalogworkspacebinding "github.com/glalanne/provider-databricks/internal/controller/namespaced/unity/catalogworkspacebinding"
	connection "github.com/glalanne/provider-databricks/internal/controller/namespaced/unity/connection"
	credential "github.com/glalanne/provider-databricks/internal/controller/namespaced/unity/credential"
	externallocation "github.com/glalanne/provider-databricks/internal/controller/namespaced/unity/externallocation"
	grant "github.com/glalanne/provider-databricks/internal/controller/namespaced/unity/grant"
	grantmap "github.com/glalanne/provider-databricks/internal/controller/namespaced/unity/grantmap"
	lakehousemonitor "github.com/glalanne/provider-databricks/internal/controller/namespaced/unity/lakehousemonitor"
	metastore "github.com/glalanne/provider-databricks/internal/controller/namespaced/unity/metastore"
	metastoreassignment "github.com/glalanne/provider-databricks/internal/controller/namespaced/unity/metastoreassignment"
	metastoredataaccess "github.com/glalanne/provider-databricks/internal/controller/namespaced/unity/metastoredataaccess"
	onlinetable "github.com/glalanne/provider-databricks/internal/controller/namespaced/unity/onlinetable"
	qualitymonitor "github.com/glalanne/provider-databricks/internal/controller/namespaced/unity/qualitymonitor"
	registeredmodel "github.com/glalanne/provider-databricks/internal/controller/namespaced/unity/registeredmodel"
	schema "github.com/glalanne/provider-databricks/internal/controller/namespaced/unity/schema"
	sqltable "github.com/glalanne/provider-databricks/internal/controller/namespaced/unity/sqltable"
	storagecredential "github.com/glalanne/provider-databricks/internal/controller/namespaced/unity/storagecredential"
	systemschema "github.com/glalanne/provider-databricks/internal/controller/namespaced/unity/systemschema"
	volume "github.com/glalanne/provider-databricks/internal/controller/namespaced/unity/volume"
	workspacebinding "github.com/glalanne/provider-databricks/internal/controller/namespaced/unity/workspacebinding"
	directory "github.com/glalanne/provider-databricks/internal/controller/namespaced/workspace/directory"
	gitcredential "github.com/glalanne/provider-databricks/internal/controller/namespaced/workspace/gitcredential"
	globalinitscript "github.com/glalanne/provider-databricks/internal/controller/namespaced/workspace/globalinitscript"
	notebook "github.com/glalanne/provider-databricks/internal/controller/namespaced/workspace/notebook"
	notificationdestination "github.com/glalanne/provider-databricks/internal/controller/namespaced/workspace/notificationdestination"
	repo "github.com/glalanne/provider-databricks/internal/controller/namespaced/workspace/repo"
	workspaceconf "github.com/glalanne/provider-databricks/internal/controller/namespaced/workspace/workspaceconf"
	workspacefile "github.com/glalanne/provider-databricks/internal/controller/namespaced/workspace/workspacefile"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		app.Setup,
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
		accountfederationpolicy.Setup,
		serviceprincipalfederationpolicy.Setup,
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
		app.SetupGated,
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
		accountfederationpolicy.SetupGated,
		serviceprincipalfederationpolicy.SetupGated,
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
