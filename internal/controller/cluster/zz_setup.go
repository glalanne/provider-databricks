// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	aisearchendpoint "github.com/glalanne/provider-databricks/internal/controller/cluster/ai/aisearchendpoint"
	aisearchindex "github.com/glalanne/provider-databricks/internal/controller/cluster/ai/aisearchindex"
	knowledgeassistant "github.com/glalanne/provider-databricks/internal/controller/cluster/ai/knowledgeassistant"
	knowledgeassistantknowledgesource "github.com/glalanne/provider-databricks/internal/controller/cluster/ai/knowledgeassistantknowledgesource"
	supervisoragent "github.com/glalanne/provider-databricks/internal/controller/cluster/ai/supervisoragent"
	supervisoragenttool "github.com/glalanne/provider-databricks/internal/controller/cluster/ai/supervisoragenttool"
	app "github.com/glalanne/provider-databricks/internal/controller/cluster/apps/app"
	customappintegration "github.com/glalanne/provider-databricks/internal/controller/cluster/apps/customappintegration"
	budgetpolicy "github.com/glalanne/provider-databricks/internal/controller/cluster/billing/budgetpolicy"
	cluster "github.com/glalanne/provider-databricks/internal/controller/cluster/compute/cluster"
	clusterpolicy "github.com/glalanne/provider-databricks/internal/controller/cluster/compute/clusterpolicy"
	instancepool "github.com/glalanne/provider-databricks/internal/controller/cluster/compute/instancepool"
	job "github.com/glalanne/provider-databricks/internal/controller/cluster/compute/job"
	library "github.com/glalanne/provider-databricks/internal/controller/cluster/compute/library"
	pipeline "github.com/glalanne/provider-databricks/internal/controller/cluster/compute/pipeline"
	databaseinstance "github.com/glalanne/provider-databricks/internal/controller/cluster/databases/databaseinstance"
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
	disasterrecoveryfailovergroup "github.com/glalanne/provider-databricks/internal/controller/cluster/dr/disasterrecoveryfailovergroup"
	disasterrecoverystableurl "github.com/glalanne/provider-databricks/internal/controller/cluster/dr/disasterrecoverystableurl"
	environmentsdefaultworkspacebaseenvironment "github.com/glalanne/provider-databricks/internal/controller/cluster/envs/environmentsdefaultworkspacebaseenvironment"
	environmentsworkspacebaseenvironment "github.com/glalanne/provider-databricks/internal/controller/cluster/envs/environmentsworkspacebaseenvironment"
	budget "github.com/glalanne/provider-databricks/internal/controller/cluster/finops/budget"
	dataclassificationcatalogconfig "github.com/glalanne/provider-databricks/internal/controller/cluster/governance/dataclassificationcatalogconfig"
	mwslogdelivery "github.com/glalanne/provider-databricks/internal/controller/cluster/log/mwslogdelivery"
	mlflowexperiment "github.com/glalanne/provider-databricks/internal/controller/cluster/mlflow/mlflowexperiment"
	mlflowmodel "github.com/glalanne/provider-databricks/internal/controller/cluster/mlflow/mlflowmodel"
	mlflowwebhook "github.com/glalanne/provider-databricks/internal/controller/cluster/mlflow/mlflowwebhook"
	vectorsearchendpoint "github.com/glalanne/provider-databricks/internal/controller/cluster/mosaic/vectorsearchendpoint"
	vectorsearchindex "github.com/glalanne/provider-databricks/internal/controller/cluster/mosaic/vectorsearchindex"
	accountfederationpolicy "github.com/glalanne/provider-databricks/internal/controller/cluster/oauth/accountfederationpolicy"
	serviceprincipalfederationpolicy "github.com/glalanne/provider-databricks/internal/controller/cluster/oauth/serviceprincipalfederationpolicy"
	postgresbranch "github.com/glalanne/provider-databricks/internal/controller/cluster/postgres/postgresbranch"
	postgrescatalog "github.com/glalanne/provider-databricks/internal/controller/cluster/postgres/postgrescatalog"
	postgrescdfconfig "github.com/glalanne/provider-databricks/internal/controller/cluster/postgres/postgrescdfconfig"
	postgresdatabase "github.com/glalanne/provider-databricks/internal/controller/cluster/postgres/postgresdatabase"
	postgresendpoint "github.com/glalanne/provider-databricks/internal/controller/cluster/postgres/postgresendpoint"
	postgresproject "github.com/glalanne/provider-databricks/internal/controller/cluster/postgres/postgresproject"
	postgresrole "github.com/glalanne/provider-databricks/internal/controller/cluster/postgres/postgresrole"
	postgressyncedtable "github.com/glalanne/provider-databricks/internal/controller/cluster/postgres/postgressyncedtable"
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
	accountnetworkpolicy "github.com/glalanne/provider-databricks/internal/controller/cluster/settings/accountnetworkpolicy"
	accountsettingv2 "github.com/glalanne/provider-databricks/internal/controller/cluster/settings/accountsettingv2"
	aibidashboardembeddingaccesspolicysetting "github.com/glalanne/provider-databricks/internal/controller/cluster/settings/aibidashboardembeddingaccesspolicysetting"
	aibidashboardembeddingapproveddomainssetting "github.com/glalanne/provider-databricks/internal/controller/cluster/settings/aibidashboardembeddingapproveddomainssetting"
	compliancesecurityprofileworkspacesetting "github.com/glalanne/provider-databricks/internal/controller/cluster/settings/compliancesecurityprofileworkspacesetting"
	defaultnamespacesetting "github.com/glalanne/provider-databricks/internal/controller/cluster/settings/defaultnamespacesetting"
	disablelegacyaccesssetting "github.com/glalanne/provider-databricks/internal/controller/cluster/settings/disablelegacyaccesssetting"
	disablelegacydbfssetting "github.com/glalanne/provider-databricks/internal/controller/cluster/settings/disablelegacydbfssetting"
	disablelegacyfeaturessetting "github.com/glalanne/provider-databricks/internal/controller/cluster/settings/disablelegacyfeaturessetting"
	enhancedsecuritymonitoringworkspacesetting "github.com/glalanne/provider-databricks/internal/controller/cluster/settings/enhancedsecuritymonitoringworkspacesetting"
	restrictworkspaceadminssetting "github.com/glalanne/provider-databricks/internal/controller/cluster/settings/restrictworkspaceadminssetting"
	workspacenetworkoption "github.com/glalanne/provider-databricks/internal/controller/cluster/settings/workspacenetworkoption"
	workspacesettingv2 "github.com/glalanne/provider-databricks/internal/controller/cluster/settings/workspacesettingv2"
	provider "github.com/glalanne/provider-databricks/internal/controller/cluster/sharing/provider"
	recipient "github.com/glalanne/provider-databricks/internal/controller/cluster/sharing/recipient"
	share "github.com/glalanne/provider-databricks/internal/controller/cluster/sharing/share"
	alert "github.com/glalanne/provider-databricks/internal/controller/cluster/sql/alert"
	alertv2 "github.com/glalanne/provider-databricks/internal/controller/cluster/sql/alertv2"
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
	tagpolicy "github.com/glalanne/provider-databricks/internal/controller/cluster/tags/tagpolicy"
	workspaceentitytagassignment "github.com/glalanne/provider-databricks/internal/controller/cluster/tags/workspaceentitytagassignment"
	artifactallowlist "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/artifactallowlist"
	catalog "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/catalog"
	catalogworkspacebinding "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/catalogworkspacebinding"
	connection "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/connection"
	credential "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/credential"
	dataqualityrefresh "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/dataqualityrefresh"
	entitytagassignment "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/entitytagassignment"
	externallocation "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/externallocation"
	externalmetadata "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/externalmetadata"
	grant "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/grant"
	grantmap "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/grantmap"
	lakehousemonitor "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/lakehousemonitor"
	metastore "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/metastore"
	metastoreassignment "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/metastoreassignment"
	metastoredataaccess "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/metastoredataaccess"
	onlinetable "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/onlinetable"
	policyinfo "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/policyinfo"
	qualitymonitor "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/qualitymonitor"
	registeredmodel "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/registeredmodel"
	rfaaccessrequestdestinations "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/rfaaccessrequestdestinations"
	schema "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/schema"
	secretuc "github.com/glalanne/provider-databricks/internal/controller/cluster/unity/secretuc"
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
		aisearchendpoint.Setup,
		aisearchindex.Setup,
		knowledgeassistant.Setup,
		knowledgeassistantknowledgesource.Setup,
		supervisoragent.Setup,
		supervisoragenttool.Setup,
		app.Setup,
		customappintegration.Setup,
		budgetpolicy.Setup,
		cluster.Setup,
		clusterpolicy.Setup,
		instancepool.Setup,
		job.Setup,
		library.Setup,
		pipeline.Setup,
		databaseinstance.Setup,
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
		disasterrecoveryfailovergroup.Setup,
		disasterrecoverystableurl.Setup,
		environmentsdefaultworkspacebaseenvironment.Setup,
		environmentsworkspacebaseenvironment.Setup,
		budget.Setup,
		dataclassificationcatalogconfig.Setup,
		mwslogdelivery.Setup,
		mlflowexperiment.Setup,
		mlflowmodel.Setup,
		mlflowwebhook.Setup,
		vectorsearchendpoint.Setup,
		vectorsearchindex.Setup,
		accountfederationpolicy.Setup,
		serviceprincipalfederationpolicy.Setup,
		postgresbranch.Setup,
		postgrescatalog.Setup,
		postgrescdfconfig.Setup,
		postgresdatabase.Setup,
		postgresendpoint.Setup,
		postgresproject.Setup,
		postgresrole.Setup,
		postgressyncedtable.Setup,
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
		accountnetworkpolicy.Setup,
		accountsettingv2.Setup,
		aibidashboardembeddingaccesspolicysetting.Setup,
		aibidashboardembeddingapproveddomainssetting.Setup,
		compliancesecurityprofileworkspacesetting.Setup,
		defaultnamespacesetting.Setup,
		disablelegacyaccesssetting.Setup,
		disablelegacydbfssetting.Setup,
		disablelegacyfeaturessetting.Setup,
		enhancedsecuritymonitoringworkspacesetting.Setup,
		restrictworkspaceadminssetting.Setup,
		workspacenetworkoption.Setup,
		workspacesettingv2.Setup,
		provider.Setup,
		recipient.Setup,
		share.Setup,
		alert.Setup,
		alertv2.Setup,
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
		tagpolicy.Setup,
		workspaceentitytagassignment.Setup,
		artifactallowlist.Setup,
		catalog.Setup,
		catalogworkspacebinding.Setup,
		connection.Setup,
		credential.Setup,
		dataqualityrefresh.Setup,
		entitytagassignment.Setup,
		externallocation.Setup,
		externalmetadata.Setup,
		grant.Setup,
		grantmap.Setup,
		lakehousemonitor.Setup,
		metastore.Setup,
		metastoreassignment.Setup,
		metastoredataaccess.Setup,
		onlinetable.Setup,
		policyinfo.Setup,
		qualitymonitor.Setup,
		registeredmodel.Setup,
		rfaaccessrequestdestinations.Setup,
		schema.Setup,
		secretuc.Setup,
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
		aisearchendpoint.SetupGated,
		aisearchindex.SetupGated,
		knowledgeassistant.SetupGated,
		knowledgeassistantknowledgesource.SetupGated,
		supervisoragent.SetupGated,
		supervisoragenttool.SetupGated,
		app.SetupGated,
		customappintegration.SetupGated,
		budgetpolicy.SetupGated,
		cluster.SetupGated,
		clusterpolicy.SetupGated,
		instancepool.SetupGated,
		job.SetupGated,
		library.SetupGated,
		pipeline.SetupGated,
		databaseinstance.SetupGated,
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
		disasterrecoveryfailovergroup.SetupGated,
		disasterrecoverystableurl.SetupGated,
		environmentsdefaultworkspacebaseenvironment.SetupGated,
		environmentsworkspacebaseenvironment.SetupGated,
		budget.SetupGated,
		dataclassificationcatalogconfig.SetupGated,
		mwslogdelivery.SetupGated,
		mlflowexperiment.SetupGated,
		mlflowmodel.SetupGated,
		mlflowwebhook.SetupGated,
		vectorsearchendpoint.SetupGated,
		vectorsearchindex.SetupGated,
		accountfederationpolicy.SetupGated,
		serviceprincipalfederationpolicy.SetupGated,
		postgresbranch.SetupGated,
		postgrescatalog.SetupGated,
		postgrescdfconfig.SetupGated,
		postgresdatabase.SetupGated,
		postgresendpoint.SetupGated,
		postgresproject.SetupGated,
		postgresrole.SetupGated,
		postgressyncedtable.SetupGated,
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
		accountnetworkpolicy.SetupGated,
		accountsettingv2.SetupGated,
		aibidashboardembeddingaccesspolicysetting.SetupGated,
		aibidashboardembeddingapproveddomainssetting.SetupGated,
		compliancesecurityprofileworkspacesetting.SetupGated,
		defaultnamespacesetting.SetupGated,
		disablelegacyaccesssetting.SetupGated,
		disablelegacydbfssetting.SetupGated,
		disablelegacyfeaturessetting.SetupGated,
		enhancedsecuritymonitoringworkspacesetting.SetupGated,
		restrictworkspaceadminssetting.SetupGated,
		workspacenetworkoption.SetupGated,
		workspacesettingv2.SetupGated,
		provider.SetupGated,
		recipient.SetupGated,
		share.SetupGated,
		alert.SetupGated,
		alertv2.SetupGated,
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
		tagpolicy.SetupGated,
		workspaceentitytagassignment.SetupGated,
		artifactallowlist.SetupGated,
		catalog.SetupGated,
		catalogworkspacebinding.SetupGated,
		connection.SetupGated,
		credential.SetupGated,
		dataqualityrefresh.SetupGated,
		entitytagassignment.SetupGated,
		externallocation.SetupGated,
		externalmetadata.SetupGated,
		grant.SetupGated,
		grantmap.SetupGated,
		lakehousemonitor.SetupGated,
		metastore.SetupGated,
		metastoreassignment.SetupGated,
		metastoredataaccess.SetupGated,
		onlinetable.SetupGated,
		policyinfo.SetupGated,
		qualitymonitor.SetupGated,
		registeredmodel.SetupGated,
		rfaaccessrequestdestinations.SetupGated,
		schema.SetupGated,
		secretuc.SetupGated,
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

// SetupWebhookWithManager registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		aisearchendpoint.SetupWebhookWithManager,
		aisearchindex.SetupWebhookWithManager,
		knowledgeassistant.SetupWebhookWithManager,
		knowledgeassistantknowledgesource.SetupWebhookWithManager,
		supervisoragent.SetupWebhookWithManager,
		supervisoragenttool.SetupWebhookWithManager,
		app.SetupWebhookWithManager,
		customappintegration.SetupWebhookWithManager,
		budgetpolicy.SetupWebhookWithManager,
		cluster.SetupWebhookWithManager,
		clusterpolicy.SetupWebhookWithManager,
		instancepool.SetupWebhookWithManager,
		job.SetupWebhookWithManager,
		library.SetupWebhookWithManager,
		pipeline.SetupWebhookWithManager,
		databaseinstance.SetupWebhookWithManager,
		instanceprofile.SetupWebhookWithManager,
		mwscredentials.SetupWebhookWithManager,
		mwscustomermanagedkeys.SetupWebhookWithManager,
		mwsnccbinding.SetupWebhookWithManager,
		mwsnccprivateendpointrule.SetupWebhookWithManager,
		mwsnetworkconnectivityconfig.SetupWebhookWithManager,
		mwsnetworks.SetupWebhookWithManager,
		mwsprivateaccesssettings.SetupWebhookWithManager,
		mwsstorageconfigurations.SetupWebhookWithManager,
		mwsvpcendpoint.SetupWebhookWithManager,
		mwsworkspaces.SetupWebhookWithManager,
		disasterrecoveryfailovergroup.SetupWebhookWithManager,
		disasterrecoverystableurl.SetupWebhookWithManager,
		environmentsdefaultworkspacebaseenvironment.SetupWebhookWithManager,
		environmentsworkspacebaseenvironment.SetupWebhookWithManager,
		budget.SetupWebhookWithManager,
		dataclassificationcatalogconfig.SetupWebhookWithManager,
		mwslogdelivery.SetupWebhookWithManager,
		mlflowexperiment.SetupWebhookWithManager,
		mlflowmodel.SetupWebhookWithManager,
		mlflowwebhook.SetupWebhookWithManager,
		vectorsearchendpoint.SetupWebhookWithManager,
		vectorsearchindex.SetupWebhookWithManager,
		accountfederationpolicy.SetupWebhookWithManager,
		serviceprincipalfederationpolicy.SetupWebhookWithManager,
		postgresbranch.SetupWebhookWithManager,
		postgrescatalog.SetupWebhookWithManager,
		postgrescdfconfig.SetupWebhookWithManager,
		postgresdatabase.SetupWebhookWithManager,
		postgresendpoint.SetupWebhookWithManager,
		postgresproject.SetupWebhookWithManager,
		postgresrole.SetupWebhookWithManager,
		postgressyncedtable.SetupWebhookWithManager,
		providerconfig.SetupWebhookWithManager,
		accesscontrolruleset.SetupWebhookWithManager,
		entitlements.SetupWebhookWithManager,
		group.SetupWebhookWithManager,
		groupinstanceprofile.SetupWebhookWithManager,
		groupmember.SetupWebhookWithManager,
		grouprole.SetupWebhookWithManager,
		ipaccesslist.SetupWebhookWithManager,
		mwspermissionassignment.SetupWebhookWithManager,
		obotoken.SetupWebhookWithManager,
		permissionassignment.SetupWebhookWithManager,
		permissions.SetupWebhookWithManager,
		secret.SetupWebhookWithManager,
		secretacl.SetupWebhookWithManager,
		secretscope.SetupWebhookWithManager,
		serviceprincipal.SetupWebhookWithManager,
		serviceprincipalrole.SetupWebhookWithManager,
		serviceprincipalsecret.SetupWebhookWithManager,
		sqlpermissions.SetupWebhookWithManager,
		token.SetupWebhookWithManager,
		user.SetupWebhookWithManager,
		userinstanceprofile.SetupWebhookWithManager,
		userrole.SetupWebhookWithManager,
		modelserving.SetupWebhookWithManager,
		accountnetworkpolicy.SetupWebhookWithManager,
		accountsettingv2.SetupWebhookWithManager,
		aibidashboardembeddingaccesspolicysetting.SetupWebhookWithManager,
		aibidashboardembeddingapproveddomainssetting.SetupWebhookWithManager,
		compliancesecurityprofileworkspacesetting.SetupWebhookWithManager,
		defaultnamespacesetting.SetupWebhookWithManager,
		disablelegacyaccesssetting.SetupWebhookWithManager,
		disablelegacydbfssetting.SetupWebhookWithManager,
		disablelegacyfeaturessetting.SetupWebhookWithManager,
		enhancedsecuritymonitoringworkspacesetting.SetupWebhookWithManager,
		restrictworkspaceadminssetting.SetupWebhookWithManager,
		workspacenetworkoption.SetupWebhookWithManager,
		workspacesettingv2.SetupWebhookWithManager,
		provider.SetupWebhookWithManager,
		recipient.SetupWebhookWithManager,
		share.SetupWebhookWithManager,
		alert.SetupWebhookWithManager,
		alertv2.SetupWebhookWithManager,
		dashboard.SetupWebhookWithManager,
		query.SetupWebhookWithManager,
		sqlalert.SetupWebhookWithManager,
		sqldashboard.SetupWebhookWithManager,
		sqlendpoint.SetupWebhookWithManager,
		sqlglobalconfig.SetupWebhookWithManager,
		sqlquery.SetupWebhookWithManager,
		sqlvisualization.SetupWebhookWithManager,
		sqlwidget.SetupWebhookWithManager,
		dbfsfile.SetupWebhookWithManager,
		file.SetupWebhookWithManager,
		mount.SetupWebhookWithManager,
		tagpolicy.SetupWebhookWithManager,
		workspaceentitytagassignment.SetupWebhookWithManager,
		artifactallowlist.SetupWebhookWithManager,
		catalog.SetupWebhookWithManager,
		catalogworkspacebinding.SetupWebhookWithManager,
		connection.SetupWebhookWithManager,
		credential.SetupWebhookWithManager,
		dataqualityrefresh.SetupWebhookWithManager,
		entitytagassignment.SetupWebhookWithManager,
		externallocation.SetupWebhookWithManager,
		externalmetadata.SetupWebhookWithManager,
		grant.SetupWebhookWithManager,
		grantmap.SetupWebhookWithManager,
		lakehousemonitor.SetupWebhookWithManager,
		metastore.SetupWebhookWithManager,
		metastoreassignment.SetupWebhookWithManager,
		metastoredataaccess.SetupWebhookWithManager,
		onlinetable.SetupWebhookWithManager,
		policyinfo.SetupWebhookWithManager,
		qualitymonitor.SetupWebhookWithManager,
		registeredmodel.SetupWebhookWithManager,
		rfaaccessrequestdestinations.SetupWebhookWithManager,
		schema.SetupWebhookWithManager,
		secretuc.SetupWebhookWithManager,
		sqltable.SetupWebhookWithManager,
		storagecredential.SetupWebhookWithManager,
		systemschema.SetupWebhookWithManager,
		volume.SetupWebhookWithManager,
		workspacebinding.SetupWebhookWithManager,
		directory.SetupWebhookWithManager,
		gitcredential.SetupWebhookWithManager,
		globalinitscript.SetupWebhookWithManager,
		notebook.SetupWebhookWithManager,
		notificationdestination.SetupWebhookWithManager,
		repo.SetupWebhookWithManager,
		workspaceconf.SetupWebhookWithManager,
		workspacefile.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
