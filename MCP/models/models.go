package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.requested_action JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Oid string `json:"oid"`
	Path string `json:"path"`
	Ref_name string `json:"ref_name"`
	Size int `json:"size"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Thread map[string]interface{} `json:"thread"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Tags []string `json:"tags,omitempty"` // A set of tags applicable for the rule.
	Description string `json:"description,omitempty"` // A short description of the rule used to detect the alert.
	Id string `json:"id,omitempty"` // A unique identifier for the rule used to detect the alert.
	Name string `json:"name,omitempty"` // The name of the rule used to detect the alert.
	Severity string `json:"severity,omitempty"` // The severity of the alert.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation Installation `json:"installation"` // Installation
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue_body_url string `json:"issue_body_url"` // The API URL to get the issue where the secret was detected.
}

// License represents the License schema from the OpenAPI specification
type License struct {
	Description string `json:"description"`
	Html_url string `json:"html_url"`
	Implementation string `json:"implementation"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Body string `json:"body"`
	Conditions []string `json:"conditions"`
	Featured bool `json:"featured"`
	Key string `json:"key"`
	Limitations []string `json:"limitations"`
	Permissions []string `json:"permissions"`
	Spdx_id string `json:"spdx_id"`
}

// Codespace represents the Codespace schema from the OpenAPI specification
type Codespace struct {
	Stop_url string `json:"stop_url"` // API URL to stop this codespace.
	Recent_folders []string `json:"recent_folders"`
	Last_known_stop_notice string `json:"last_known_stop_notice,omitempty"` // The text to display to a user when a codespace has been stopped for a potentially actionable reason.
	Updated_at string `json:"updated_at"`
	Git_status map[string]interface{} `json:"git_status"` // Details about the codespace's git repository.
	Web_url string `json:"web_url"` // URL to access this codespace on the web.
	State string `json:"state"` // State of this codespace.
	Display_name string `json:"display_name,omitempty"` // Display name for this codespace.
	Name string `json:"name"` // Automatically generated name of this codespace.
	Retention_period_minutes int `json:"retention_period_minutes,omitempty"` // Duration in minutes after codespace has gone idle in which it will be deleted. Must be integer minutes between 0 and 43200 (30 days).
	Devcontainer_path string `json:"devcontainer_path,omitempty"` // Path to devcontainer.json from repo root used to create Codespace.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Machine GeneratedType `json:"machine"` // A description of the machine powering a codespace.
	Publish_url string `json:"publish_url,omitempty"` // API URL to publish this codespace to a new repository.
	Url string `json:"url"` // API URL for this codespace.
	Last_used_at string `json:"last_used_at"` // Last known time this codespace was started.
	Billable_owner GeneratedType `json:"billable_owner"` // A GitHub user.
	Idle_timeout_notice string `json:"idle_timeout_notice,omitempty"` // Text to show user when codespace idle timeout minutes has been overriden by an organization policy
	Start_url string `json:"start_url"` // API URL to start this codespace.
	Runtime_constraints map[string]interface{} `json:"runtime_constraints,omitempty"`
	Environment_id string `json:"environment_id"` // UUID identifying this codespace's environment.
	Pending_operation_disabled_reason string `json:"pending_operation_disabled_reason,omitempty"` // Text to show user when codespace is disabled by a pending operation
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Created_at string `json:"created_at"`
	Location string `json:"location"` // The Azure region where this codespace is located.
	Pending_operation bool `json:"pending_operation,omitempty"` // Whether or not a codespace has a pending async operation. This would mean that the codespace is temporarily unavailable. The only thing that you can do with a codespace in this state is delete it.
	Id int `json:"id"`
	Idle_timeout_minutes int `json:"idle_timeout_minutes"` // The number of minutes of inactivity after which this codespace will be automatically stopped.
	Prebuild bool `json:"prebuild"` // Whether the codespace was created from a prebuild.
	Retention_expires_at string `json:"retention_expires_at,omitempty"` // When a codespace will be auto-deleted based on the "retention_period_minutes" and "last_used_at"
	Pulls_url string `json:"pulls_url"` // API URL for the Pull Request associated with this codespace, if any.
	Machines_url string `json:"machines_url"` // API URL to access available alternate machine types for this codespace.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action that was performed.
	Assignee map[string]interface{} `json:"assignee,omitempty"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Pusher_type string `json:"pusher_type"` // The pusher type for the event. Can be either `user` or a deploy key.
	Description string `json:"description"` // The repository's current description.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Ref string `json:"ref"` // The [`git ref`](https://docs.github.com/rest/reference/git#get-a-reference) resource.
	Ref_type string `json:"ref_type"` // The type of Git ref object created in the repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Master_branch string `json:"master_branch"` // The name of the repository's default branch (usually `main`).
}

// Thread represents the Thread schema from the OpenAPI specification
type Thread struct {
	Unread bool `json:"unread"`
	Id string `json:"id"`
	Subject map[string]interface{} `json:"subject"`
	Reason string `json:"reason"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Updated_at string `json:"updated_at"`
	Subscription_url string `json:"subscription_url"`
	Url string `json:"url"`
	Last_read_at string `json:"last_read_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Can_approve_pull_request_reviews bool `json:"can_approve_pull_request_reviews,omitempty"` // Whether GitHub Actions can approve pull requests. Enabling this can be a security risk.
	Default_workflow_permissions string `json:"default_workflow_permissions,omitempty"` // The default workflow permissions granted to the GITHUB_TOKEN when running workflows.
}

// Environment represents the Environment schema from the OpenAPI specification
type Environment struct {
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Id int `json:"id"` // The id of the environment.
	Name string `json:"name"` // The name of the environment.
	Protection_rules []interface{} `json:"protection_rules,omitempty"`
	Url string `json:"url"`
	Created_at string `json:"created_at"` // The time that the environment was created, in ISO 8601 format.
	Updated_at string `json:"updated_at"` // The time that the environment was last updated, in ISO 8601 format.
	Deployment_branch_policy GeneratedType `json:"deployment_branch_policy,omitempty"` // The type of deployment branch policy for this environment. To allow all branches to deploy, set to `null`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository_url string `json:"repository_url,omitempty"`
	Subscribed bool `json:"subscribed"`
	Thread_url string `json:"thread_url,omitempty"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Ignored bool `json:"ignored"`
	Reason string `json:"reason"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Cve_id string `json:"cve_id"` // The unique CVE ID assigned to the advisory.
	Summary string `json:"summary"` // A short, plain text summary of the advisory.
	Vulnerabilities []GeneratedType `json:"vulnerabilities"` // Vulnerable version range information for the advisory.
	Description string `json:"description"` // A long-form Markdown-supported description of the advisory.
	Ghsa_id string `json:"ghsa_id"` // The unique GitHub Security Advisory ID assigned to the advisory.
	References []map[string]interface{} `json:"references"` // Links to additional advisory information.
	Withdrawn_at string `json:"withdrawn_at"` // The time that the advisory was withdrawn in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Identifiers []map[string]interface{} `json:"identifiers"` // Values that identify this advisory among security information sources.
	Published_at string `json:"published_at"` // The time that the advisory was published in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Cvss map[string]interface{} `json:"cvss"` // Details for the advisory pertaining to the Common Vulnerability Scoring System.
	Cwes []map[string]interface{} `json:"cwes"` // Details for the advisory pertaining to Common Weakness Enumeration.
	Severity string `json:"severity"` // The severity of the advisory.
	Updated_at string `json:"updated_at"` // The time that the advisory was last modified in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Duration int `json:"duration"`
	ErrorField map[string]interface{} `json:"error"`
	Pusher GeneratedType `json:"pusher"` // A GitHub user.
	Status string `json:"status"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Commit string `json:"commit"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Comment map[string]interface{} `json:"comment"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Gravatar_id string `json:"gravatar_id"`
	Starred_url string `json:"starred_url"`
	Events_url string `json:"events_url"`
	Following_url string `json:"following_url"`
	Gists_url string `json:"gists_url"`
	Subscriptions_url string `json:"subscriptions_url"`
	Html_url string `json:"html_url"`
	Repos_url string `json:"repos_url"`
	Node_id string `json:"node_id"`
	Starred_at string `json:"starred_at,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Site_admin bool `json:"site_admin"`
	Followers_url string `json:"followers_url"`
	Login string `json:"login"`
	Email string `json:"email,omitempty"`
	Id int `json:"id"`
	Name string `json:"name,omitempty"`
	Organizations_url string `json:"organizations_url"`
	Url string `json:"url"`
	Received_events_url string `json:"received_events_url"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// Integration represents the Integration schema from the OpenAPI specification
type Integration struct {
	Installations_count int `json:"installations_count,omitempty"` // The number of installations associated with the GitHub app
	Events []string `json:"events"` // The list of events for the GitHub app
	Pem string `json:"pem,omitempty"`
	Updated_at string `json:"updated_at"`
	Webhook_secret string `json:"webhook_secret,omitempty"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Slug string `json:"slug,omitempty"` // The slug name of the GitHub app
	Html_url string `json:"html_url"`
	Permissions map[string]interface{} `json:"permissions"` // The set of permissions for the GitHub app
	Created_at string `json:"created_at"`
	Description string `json:"description"`
	Client_secret string `json:"client_secret,omitempty"`
	Client_id string `json:"client_id,omitempty"`
	Name string `json:"name"` // The name of the GitHub app
	Node_id string `json:"node_id"`
	External_url string `json:"external_url"`
	Id int `json:"id"` // Unique identifier of the GitHub app
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name,omitempty"` // The name pattern that branches must match in order to deploy to the environment.
	Node_id string `json:"node_id,omitempty"`
	Id int `json:"id,omitempty"` // The unique identifier of the branch policy.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship map[string]interface{} `json:"sponsorship"`
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue interface{} `json:"issue"`
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name pattern that branches must match in order to deploy to the environment. Wildcard characters will not match `/`. For example, to match branches that begin with `release/` and contain an additional single slash, use `release/*/*`. For more information about pattern matching syntax, see the [Ruby File.fnmatch documentation](https://ruby-doc.org/core-2.5.1/File.html#method-c-fnmatch).
}

// Hovercard represents the Hovercard schema from the OpenAPI specification
type Hovercard struct {
	Contexts []map[string]interface{} `json:"contexts"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repo Repository `json:"repo"` // A repository on GitHub.
	Starred_at string `json:"starred_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
	Commit_id string `json:"commit_id"`
	Url string `json:"url"`
	Assigner GeneratedType `json:"assigner"` // A GitHub user.
	Node_id string `json:"node_id"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Slug string `json:"slug"`
	Updated_at string `json:"updated_at"`
	Organization GeneratedType `json:"organization"` // Team Organization
	Parent GeneratedType `json:"parent,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Repositories_url string `json:"repositories_url"`
	Description string `json:"description"`
	Ldap_dn string `json:"ldap_dn,omitempty"` // Distinguished Name (DN) that team maps to within LDAP environment
	Members_url string `json:"members_url"`
	Node_id string `json:"node_id"`
	Repos_count int `json:"repos_count"`
	Members_count int `json:"members_count"`
	Privacy string `json:"privacy,omitempty"` // The level of privacy this team should have
	Id int `json:"id"` // Unique identifier of the team
	Url string `json:"url"` // URL for the team
	Name string `json:"name"` // Name of the team
	Created_at string `json:"created_at"`
	Permission string `json:"permission"` // Permission that the team will have for its repositories
	Html_url string `json:"html_url"`
}

// Deployment represents the Deployment schema from the OpenAPI specification
type Deployment struct {
	Created_at string `json:"created_at"`
	Description string `json:"description"`
	Ref string `json:"ref"` // The ref to deploy. This can be a branch, tag, or sha.
	Sha string `json:"sha"`
	Production_environment bool `json:"production_environment,omitempty"` // Specifies if the given environment is one that end-users directly interact with. Default: false.
	Repository_url string `json:"repository_url"`
	Node_id string `json:"node_id"`
	Statuses_url string `json:"statuses_url"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Payload interface{} `json:"payload"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Transient_environment bool `json:"transient_environment,omitempty"` // Specifies if the given environment is will no longer exist at some point in the future. Default: false.
	Environment string `json:"environment"` // Name for the target deployment environment.
	Url string `json:"url"`
	Updated_at string `json:"updated_at"`
	Task string `json:"task"` // Parameter to specify a task to execute
	Id int `json:"id"` // Unique identifier of the deployment
	Original_environment string `json:"original_environment,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Action string `json:"action"`
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Browser_download_url string `json:"browser_download_url"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Label string `json:"label"`
	State string `json:"state"` // State of the release asset.
	Download_count int `json:"download_count"`
	Size int `json:"size"`
	Uploader GeneratedType `json:"uploader"` // A GitHub user.
	Content_type string `json:"content_type"`
	Name string `json:"name"` // The file name of the asset.
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Author map[string]interface{} `json:"author"` // Identifying information for the git-user
	Committer map[string]interface{} `json:"committer"` // Identifying information for the git-user
	Node_id string `json:"node_id"`
	Sha string `json:"sha"` // SHA for the commit
	Message string `json:"message"` // Message describing the purpose of the commit
	Tree map[string]interface{} `json:"tree"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Verification map[string]interface{} `json:"verification"`
	Parents []map[string]interface{} `json:"parents"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Milestone map[string]interface{} `json:"milestone"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Event string `json:"event"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Membership map[string]interface{} `json:"membership,omitempty"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Message string `json:"message"`
	Documentation_url string `json:"documentation_url"`
	Errors []map[string]interface{} `json:"errors,omitempty"`
}

// Hook represents the Hook schema from the OpenAPI specification
type Hook struct {
	Updated_at string `json:"updated_at"`
	Active bool `json:"active"` // Determines whether the hook is actually triggered on pushes.
	Created_at string `json:"created_at"`
	Id int `json:"id"` // Unique identifier of the webhook.
	Last_response GeneratedType `json:"last_response"`
	Ping_url string `json:"ping_url"`
	Deliveries_url string `json:"deliveries_url,omitempty"`
	TypeField string `json:"type"`
	Url string `json:"url"`
	Events []string `json:"events"` // Determines what events the hook is triggered for. Default: ['push'].
	Name string `json:"name"` // The name of a valid service, use 'web' for a webhook.
	Test_url string `json:"test_url"`
	Config map[string]interface{} `json:"config"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total int `json:"total"`
	Weeks []map[string]interface{} `json:"weeks"`
	Author GeneratedType `json:"author"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Price_model string `json:"price_model"`
	Unit_name string `json:"unit_name"`
	Url string `json:"url"`
	Yearly_price_in_cents int `json:"yearly_price_in_cents"`
	Accounts_url string `json:"accounts_url"`
	Has_free_trial bool `json:"has_free_trial"`
	Id int `json:"id"`
	Monthly_price_in_cents int `json:"monthly_price_in_cents"`
	Name string `json:"name"`
	Number int `json:"number"`
	State string `json:"state"`
	Bullets []string `json:"bullets"`
	Description string `json:"description"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Clone_url string `json:"clone_url"`
	Downloads_url string `json:"downloads_url"`
	Subscription_url string `json:"subscription_url"`
	Git_url string `json:"git_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Hooks_url string `json:"hooks_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	License GeneratedType `json:"license"` // License Simple
	Topics []string `json:"topics,omitempty"`
	Statuses_url string `json:"statuses_url"`
	Git_commits_url string `json:"git_commits_url"`
	Forks int `json:"forks"`
	Branches_url string `json:"branches_url"`
	Trees_url string `json:"trees_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Watchers int `json:"watchers"`
	Archive_url string `json:"archive_url"`
	Blobs_url string `json:"blobs_url"`
	Score float64 `json:"score"`
	Notifications_url string `json:"notifications_url"`
	Issues_url string `json:"issues_url"`
	Forks_url string `json:"forks_url"`
	Subscribers_url string `json:"subscribers_url"`
	Has_projects bool `json:"has_projects"`
	Commits_url string `json:"commits_url"`
	Html_url string `json:"html_url"`
	Stargazers_count int `json:"stargazers_count"`
	Name string `json:"name"`
	Fork bool `json:"fork"`
	Created_at string `json:"created_at"`
	Deployments_url string `json:"deployments_url"`
	Contents_url string `json:"contents_url"`
	Updated_at string `json:"updated_at"`
	Language string `json:"language"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"`
	Languages_url string `json:"languages_url"`
	Assignees_url string `json:"assignees_url"`
	Git_tags_url string `json:"git_tags_url"`
	Archived bool `json:"archived"`
	Has_discussions bool `json:"has_discussions,omitempty"`
	Has_pages bool `json:"has_pages"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Private bool `json:"private"`
	Homepage string `json:"homepage"`
	Ssh_url string `json:"ssh_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Tags_url string `json:"tags_url"`
	Git_refs_url string `json:"git_refs_url"`
	Open_issues int `json:"open_issues"`
	Pulls_url string `json:"pulls_url"`
	Comments_url string `json:"comments_url"`
	Keys_url string `json:"keys_url"`
	Default_branch string `json:"default_branch"`
	Node_id string `json:"node_id"`
	Has_wiki bool `json:"has_wiki"`
	Compare_url string `json:"compare_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"`
	Size int `json:"size"`
	Open_issues_count int `json:"open_issues_count"`
	Is_template bool `json:"is_template,omitempty"`
	Has_downloads bool `json:"has_downloads"`
	Master_branch string `json:"master_branch,omitempty"`
	Teams_url string `json:"teams_url"`
	Mirror_url string `json:"mirror_url"`
	Svn_url string `json:"svn_url"`
	Releases_url string `json:"releases_url"`
	Contributors_url string `json:"contributors_url"`
	Stargazers_url string `json:"stargazers_url"`
	Id int `json:"id"`
	Has_issues bool `json:"has_issues"`
	Labels_url string `json:"labels_url"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"`
	Events_url string `json:"events_url"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Description string `json:"description"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Collaborators_url string `json:"collaborators_url"`
	Issue_events_url string `json:"issue_events_url"`
	Milestones_url string `json:"milestones_url"`
	Url string `json:"url"`
	Watchers_count int `json:"watchers_count"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Full_name string `json:"full_name"`
	Forks_count int `json:"forks_count"`
	Pushed_at string `json:"pushed_at"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Merges_url string `json:"merges_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
	Marketplace_purchase interface{} `json:"marketplace_purchase"`
	Effective_date string `json:"effective_date"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	State string `json:"state"` // The state of the member in the organization. The `pending` state indicates the user has not yet accepted an invitation.
	Url string `json:"url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Organization_url string `json:"organization_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Role string `json:"role"` // The user's membership type in the organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body"` // The comment text.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Number int `json:"number"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Reason string `json:"reason"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Branches []map[string]interface{} `json:"branches"` // An array of branch objects containing the status' SHA. Each branch contains the given SHA, but the SHA may or may not be the head of the branch. The array includes a maximum of 10 branches.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	State string `json:"state"` // The new state. Can be `pending`, `success`, `failure`, or `error`.
	Description string `json:"description"` // The optional human-readable description added to the status.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Id int `json:"id"` // The unique identifier of the status.
	Target_url string `json:"target_url"` // The optional link added to the status.
	Commit map[string]interface{} `json:"commit"`
	Context string `json:"context"`
	Avatar_url string `json:"avatar_url,omitempty"`
	Name string `json:"name"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Updated_at string `json:"updated_at"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Sha string `json:"sha"` // The Commit SHA.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Git_url string `json:"git_url"`
	Sha string `json:"sha"`
	Target string `json:"target"`
	Size int `json:"size"`
	Download_url string `json:"download_url"`
	TypeField string `json:"type"`
	Links map[string]interface{} `json:"_links"`
	Html_url string `json:"html_url"`
	Name string `json:"name"`
	Path string `json:"path"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Archived_at string `json:"archived_at"`
	Content_node_id string `json:"content_node_id"`
	Created_at string `json:"created_at"`
	Project_node_id string `json:"project_node_id,omitempty"`
	Updated_at string `json:"updated_at"`
	Content_type string `json:"content_type"` // The type of content tracked in a project item
	Id float64 `json:"id"`
	Node_id string `json:"node_id,omitempty"`
	Creator GeneratedType `json:"creator,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Actions_caches []map[string]interface{} `json:"actions_caches"` // Array of caches
	Total_count int `json:"total_count"` // Total number of caches
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Check_suite map[string]interface{} `json:"check_suite"` // The [check_suite](https://docs.github.com/rest/reference/checks#suites).
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Actions_meta map[string]interface{} `json:"actions_meta,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Deleted_at string `json:"deleted_at"`
	Created_at string `json:"created_at"`
	Deleted_by GeneratedType `json:"deleted_by"` // A GitHub user.
	Public bool `json:"public"`
	Closed_at string `json:"closed_at"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Id float64 `json:"id"`
	Node_id string `json:"node_id"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Short_description string `json:"short_description"`
	Description string `json:"description"`
	Number int `json:"number"`
	Title string `json:"title"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action that was performed.
	Assignee map[string]interface{} `json:"assignee,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue_title_url string `json:"issue_title_url"` // The API URL to get the issue where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/reference/pulls#comments) itself.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	State_reason string `json:"state_reason,omitempty"`
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
	Event string `json:"event"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Commit_url string `json:"commit_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Effective_date string `json:"effective_date,omitempty"` // The `pending_cancellation` and `pending_tier_change` event types will include the date the cancellation or tier change will take effect.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Changes map[string]interface{} `json:"changes"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Sponsorship map[string]interface{} `json:"sponsorship"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Count int `json:"count"`
	Referrer string `json:"referrer"`
	Uniques int `json:"uniques"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_id string `json:"commit_id"` // A commit SHA for the review.
	State string `json:"state"`
	Links map[string]interface{} `json:"_links"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body"` // The text of the review.
	Event string `json:"event"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Pull_request_url string `json:"pull_request_url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Id int `json:"id"` // Unique identifier of the review
	Submitted_at string `json:"submitted_at,omitempty"`
	Body_html string `json:"body_html,omitempty"`
	Body_text string `json:"body_text,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Marketplace_purchase interface{} `json:"marketplace_purchase"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Effective_date string `json:"effective_date"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key_id string `json:"key_id"` // The identifier for the key.
	Key string `json:"key"` // The Base64 encoded public key.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Body_text string `json:"body_text,omitempty"`
	Body_html string `json:"body_html,omitempty"`
	Event string `json:"event"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Body string `json:"body,omitempty"` // Contents of the issue comment
	Id int `json:"id"` // Unique identifier of the issue comment
	Url string `json:"url"` // URL for the issue comment
	Reactions GeneratedType `json:"reactions,omitempty"`
	Node_id string `json:"node_id"`
	Html_url string `json:"html_url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Issue_url string `json:"issue_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/reference/issues#comments) itself.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) the comment belongs to.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alert interface{} `json:"alert"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// Package represents the Package schema from the OpenAPI specification
type Package struct {
	Name string `json:"name"` // The name of the package.
	Owner GeneratedType `json:"owner,omitempty"` // A GitHub user.
	Repository GeneratedType `json:"repository,omitempty"` // Minimal Repository
	Url string `json:"url"`
	Visibility string `json:"visibility"`
	Id int `json:"id"` // Unique identifier of the package.
	Package_type string `json:"package_type"`
	Updated_at string `json:"updated_at"`
	Version_count int `json:"version_count"` // The number of versions of the package.
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Node_id string `json:"node_id"`
	Deployment_url string `json:"deployment_url"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Log_url string `json:"log_url,omitempty"` // The URL to associate with this status.
	Id int `json:"id"`
	Repository_url string `json:"repository_url"`
	Created_at string `json:"created_at"`
	Description string `json:"description"` // A short description of the status.
	Updated_at string `json:"updated_at"`
	Target_url string `json:"target_url"` // Deprecated: the URL to associate with this status.
	Environment string `json:"environment,omitempty"` // The environment of the deployment that the status is for.
	State string `json:"state"` // The state of the status.
	Environment_url string `json:"environment_url,omitempty"` // The URL for accessing your environment.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	State string `json:"state"`
	Base map[string]interface{} `json:"base"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Html_url string `json:"html_url"`
	Review_comments_url string `json:"review_comments_url"`
	Body string `json:"body"`
	Issue_url string `json:"issue_url"`
	Merge_commit_sha string `json:"merge_commit_sha"`
	Draft bool `json:"draft,omitempty"` // Indicates whether or not the pull request is a draft.
	Merged_at string `json:"merged_at"`
	Number int `json:"number"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Locked bool `json:"locked"`
	Statuses_url string `json:"statuses_url"`
	Commits_url string `json:"commits_url"`
	Links map[string]interface{} `json:"_links"`
	Patch_url string `json:"patch_url"`
	Title string `json:"title"`
	Id int `json:"id"`
	Requested_teams []Team `json:"requested_teams,omitempty"`
	Requested_reviewers []GeneratedType `json:"requested_reviewers,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Closed_at string `json:"closed_at"`
	Node_id string `json:"node_id"`
	Head map[string]interface{} `json:"head"`
	Diff_url string `json:"diff_url"`
	Auto_merge GeneratedType `json:"auto_merge"` // The status of auto merging a pull request.
	Updated_at string `json:"updated_at"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	User GeneratedType `json:"user"` // A GitHub user.
	Comments_url string `json:"comments_url"`
	Labels []map[string]interface{} `json:"labels"`
	Review_comment_url string `json:"review_comment_url"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
}

// Import represents the Import schema from the OpenAPI specification
type Import struct {
	Tfvc_project string `json:"tfvc_project,omitempty"`
	Url string `json:"url"`
	Error_message string `json:"error_message,omitempty"`
	Large_files_size int `json:"large_files_size,omitempty"`
	Use_lfs bool `json:"use_lfs,omitempty"`
	Project_choices []map[string]interface{} `json:"project_choices,omitempty"`
	Push_percent int `json:"push_percent,omitempty"`
	Import_percent int `json:"import_percent,omitempty"`
	Status_text string `json:"status_text,omitempty"`
	Svn_root string `json:"svn_root,omitempty"`
	Failed_step string `json:"failed_step,omitempty"`
	Status string `json:"status"`
	Has_large_files bool `json:"has_large_files,omitempty"`
	Vcs string `json:"vcs"`
	Authors_count int `json:"authors_count,omitempty"`
	Commit_count int `json:"commit_count,omitempty"`
	Authors_url string `json:"authors_url"`
	Large_files_count int `json:"large_files_count,omitempty"`
	Message string `json:"message,omitempty"`
	Repository_url string `json:"repository_url"`
	Html_url string `json:"html_url"`
	Vcs_url string `json:"vcs_url"` // The URL of the originating repository.
	Svc_root string `json:"svc_root,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Expires_at string `json:"expires_at"` // The time this token expires
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Repositories []Repository `json:"repositories,omitempty"` // The repositories this token has access to
	Repository_selection string `json:"repository_selection,omitempty"` // Describe whether all repositories have been selected or there's a selection involved
	Single_file string `json:"single_file,omitempty"`
	Token string `json:"token"` // The token used for authentication
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.rerequested JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Id int `json:"id"`
	Label map[string]interface{} `json:"label"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Merge_group map[string]interface{} `json:"merge_group"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Access_level string `json:"access_level"` // Defines the level of access that workflows outside of the repository have to actions and reusable workflows within the repository. `none` means the access is only possible from workflows in this repository. `user` level access allows sharing across user owned private repos only. `organization` level access allows sharing across the organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Project map[string]interface{} `json:"project"`
	Repository GeneratedType `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"` // The changes to the project if the action was `edited`.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project map[string]interface{} `json:"project"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Body string `json:"body"` // The main text of the discussion.
	Comments_count int `json:"comments_count"`
	Author GeneratedType `json:"author"` // A GitHub user.
	Body_version string `json:"body_version"` // The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server.
	Last_edited_at string `json:"last_edited_at"`
	Body_html string `json:"body_html"`
	Updated_at string `json:"updated_at"`
	Comments_url string `json:"comments_url"`
	Node_id string `json:"node_id"`
	Pinned bool `json:"pinned"` // Whether or not this discussion should be pinned for easy retrieval.
	Title string `json:"title"` // The title of the discussion.
	Reactions GeneratedType `json:"reactions,omitempty"`
	Number int `json:"number"` // The unique sequence number of a team discussion.
	Private bool `json:"private"` // Whether or not this discussion should be restricted to team members and organization administrators.
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Team_url string `json:"team_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Requester map[string]interface{} `json:"requester"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Repositories_added []map[string]interface{} `json:"repositories_added"` // An array of repository objects, which were added to the installation.
	Repositories_removed []map[string]interface{} `json:"repositories_removed"` // An array of repository objects, which were removed from the installation.
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Action string `json:"action"`
	Installation Installation `json:"installation"` // Installation
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"` // The location type. Because secrets may be found in different types of resources (ie. code, comments, issues), this field identifies the type of resource where the secret was found.
	Details interface{} `json:"details"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Description string `json:"description"`
	Domains []string `json:"domains"` // Array of the domain set and its alternate name (if it is configured)
	Expires_at string `json:"expires_at,omitempty"`
	State string `json:"state"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Hook_id int `json:"hook_id,omitempty"` // The ID of the webhook that triggered the ping.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Zen string `json:"zen,omitempty"` // Random string of GitHub zen.
	Hook map[string]interface{} `json:"hook,omitempty"` // The webhook that is being pinged
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Label map[string]interface{} `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"` // The changes to the label if the action was `edited`.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Head_sha string `json:"head_sha"` // The SHA of the head commit that is being checked.
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	After string `json:"after"`
	Before string `json:"before"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Runs_rerequestable bool `json:"runs_rerequestable,omitempty"`
	App GeneratedType `json:"app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Conclusion string `json:"conclusion"`
	Id int `json:"id"`
	Latest_check_runs_count int `json:"latest_check_runs_count"`
	Rerequestable bool `json:"rerequestable,omitempty"`
	Check_runs_url string `json:"check_runs_url"`
	Head_branch string `json:"head_branch"`
	Pull_requests []GeneratedType `json:"pull_requests"`
	Status string `json:"status"`
	Created_at string `json:"created_at"`
	Url string `json:"url"`
	Head_commit GeneratedType `json:"head_commit"` // A commit.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit map[string]interface{} `json:"commit"`
	Html_url string `json:"html_url"`
	Sha string `json:"sha"`
	Comments_url string `json:"comments_url"`
	Committer GeneratedType `json:"committer"` // Metaproperties for Git author/committer information.
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Url string `json:"url"`
	Node_id string `json:"node_id"`
	Parents []map[string]interface{} `json:"parents"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Score float64 `json:"score"`
	Author GeneratedType `json:"author"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Include_claim_keys []string `json:"include_claim_keys"` // Array of unique strings. Each claim key can only contain alphanumeric characters and underscores.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Marketplace_purchase interface{} `json:"marketplace_purchase"`
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Effective_date string `json:"effective_date"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the collaborator permissions
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Member map[string]interface{} `json:"member"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// Snapshot represents the Snapshot schema from the OpenAPI specification
type Snapshot struct {
	Manifests map[string]interface{} `json:"manifests,omitempty"` // A collection of package manifests, which are a collection of related dependencies declared in a file or representing a logical group of dependencies.
	Metadata Metadata `json:"metadata,omitempty"` // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
	Ref string `json:"ref"` // The repository branch that triggered this snapshot.
	Scanned string `json:"scanned"` // The time at which the snapshot was scanned.
	Sha string `json:"sha"` // The commit SHA associated with this dependency snapshot. Maximum length: 40 characters.
	Version int `json:"version"` // The version of the repository snapshot submission.
	Detector map[string]interface{} `json:"detector"` // A description of the detector used.
	Job map[string]interface{} `json:"job"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_id string `json:"commit_id"`
	Node_id string `json:"node_id"`
	Path string `json:"path"`
	Position int `json:"position"`
	Created_at string `json:"created_at"`
	Line int `json:"line"`
	User GeneratedType `json:"user"` // A GitHub user.
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Url string `json:"url"`
	Body string `json:"body"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Updated_at string `json:"updated_at"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Blocked_user map[string]interface{} `json:"blocked_user"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Milestone map[string]interface{} `json:"milestone"`
	Url string `json:"url"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	License GeneratedType `json:"license"` // License Simple
	Links map[string]interface{} `json:"_links"`
	Download_url string `json:"download_url"`
	Git_url string `json:"git_url"`
	Html_url string `json:"html_url"`
	Name string `json:"name"`
	Size int `json:"size"`
	TypeField string `json:"type"`
	Url string `json:"url"`
	Encoding string `json:"encoding"`
	Path string `json:"path"`
	Sha string `json:"sha"`
	Content string `json:"content"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Rule GeneratedType `json:"rule"`
	Dismissed_comment string `json:"dismissed_comment,omitempty"` // The dismissal comment associated with the dismissal of the alert.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Tool GeneratedType `json:"tool"`
	Number int `json:"number"` // The security alert number.
	State string `json:"state"` // State of a code scanning alert.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Fixed_at string `json:"fixed_at,omitempty"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Most_recent_instance GeneratedType `json:"most_recent_instance"`
	Url string `json:"url"` // The REST API URL of the alert resource.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Instances_url string `json:"instances_url"` // The REST API URL for fetching the list of instances for an alert.
	Dismissed_reason string `json:"dismissed_reason"` // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Branch string `json:"branch"`
	Path string `json:"path"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Dependency map[string]interface{} `json:"dependency"` // Details for the vulnerable dependency.
	Dismissed_reason string `json:"dismissed_reason"` // The reason that the alert was dismissed.
	State string `json:"state"` // The state of the Dependabot alert.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Number int `json:"number"` // The security alert number.
	Fixed_at string `json:"fixed_at"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Repository GeneratedType `json:"repository"` // A GitHub repository.
	Updated_at string `json:"updated_at"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Security_advisory GeneratedType `json:"security_advisory"` // Details for the GitHub Security Advisory.
	Dismissed_comment string `json:"dismissed_comment"` // An optional comment associated with the alert's dismissal.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Security_vulnerability GeneratedType `json:"security_vulnerability"` // Details pertaining to one vulnerable version range for the advisory.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Selected_actions_url string `json:"selected_actions_url,omitempty"` // The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
	Allowed_actions string `json:"allowed_actions,omitempty"` // The permissions policy that controls the actions and reusable workflows that are allowed to run.
	Enabled bool `json:"enabled"` // Whether GitHub Actions is enabled on the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Label map[string]interface{} `json:"label,omitempty"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Pull_request map[string]interface{} `json:"pull_request"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Label map[string]interface{} `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Release interface{} `json:"release"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	User GeneratedType `json:"user"` // A GitHub user.
	Body_text string `json:"body_text,omitempty"`
	Commit_id string `json:"commit_id"` // A commit SHA for the review. If the commit object was garbage collected or forcibly deleted, then it no longer exists in Git and this value will be `null`.
	Node_id string `json:"node_id"`
	Links map[string]interface{} `json:"_links"`
	Body string `json:"body"` // The text of the review.
	Body_html string `json:"body_html,omitempty"`
	Pull_request_url string `json:"pull_request_url"`
	State string `json:"state"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Html_url string `json:"html_url"`
	Id int `json:"id"` // Unique identifier of the review
	Submitted_at string `json:"submitted_at,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repositories []Repository `json:"repositories,omitempty"`
	Repository_selection string `json:"repository_selection,omitempty"`
	Single_file string `json:"single_file,omitempty"`
	Single_file_paths []string `json:"single_file_paths,omitempty"`
	Token string `json:"token"`
	Expires_at string `json:"expires_at"`
	Has_multiple_single_files bool `json:"has_multiple_single_files,omitempty"`
	Permissions GeneratedType `json:"permissions,omitempty"` // The permissions granted to the user-to-server access token.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"` // The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the secret
	Selected_repositories_url string `json:"selected_repositories_url"` // The API URL at which the list of repositories this secret is visible to can be retrieved
	Updated_at string `json:"updated_at"` // The date and time at which the secret was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Visibility string `json:"visibility"` // The type of repositories in the organization that the secret is visible to
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Hook_id int `json:"hook_id"` // The id of the modified webhook.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository GeneratedType `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Hook map[string]interface{} `json:"hook"` // The modified webhook. This will contain different keys based on the type of webhook it is: repository, organization, business, app, or GitHub Marketplace.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Forkee interface{} `json:"forkee"` // The created [`repository`](https://docs.github.com/rest/reference/repos#get-a-repository) resource.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Active_caches_count int `json:"active_caches_count"` // The number of active caches in the repository.
	Active_caches_size_in_bytes int `json:"active_caches_size_in_bytes"` // The sum of the size in bytes of all the active cache items in the repository.
	Full_name string `json:"full_name"` // The repository owner and name for the cache usage being shown.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Membership map[string]interface{} `json:"membership"` // The membership between the user and the organization. Not present when the action is `member_invited`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Release interface{} `json:"release"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project_column map[string]interface{} `json:"project_column"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Links map[string]interface{} `json:"_links"`
	Content string `json:"content"`
	Download_url string `json:"download_url"`
	Html_url string `json:"html_url"`
	Name string `json:"name"`
	Path string `json:"path"`
	Size int `json:"size"`
	TypeField string `json:"type"`
	Encoding string `json:"encoding"`
	Git_url string `json:"git_url"`
	Sha string `json:"sha"`
	Submodule_git_url string `json:"submodule_git_url,omitempty"`
	Target string `json:"target,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Machine GeneratedType `json:"machine"` // A description of the machine powering a codespace.
	Id int `json:"id"`
	State string `json:"state"` // State of this codespace.
	Last_used_at string `json:"last_used_at"` // Last known time this codespace was started.
	Stop_url string `json:"stop_url"` // API URL to stop this codespace.
	Git_status map[string]interface{} `json:"git_status"` // Details about the codespace's git repository.
	Prebuild bool `json:"prebuild"` // Whether the codespace was created from a prebuild.
	Repository GeneratedType `json:"repository"` // Full Repository
	Environment_id string `json:"environment_id"` // UUID identifying this codespace's environment.
	Location string `json:"location"` // The Azure region where this codespace is located.
	Start_url string `json:"start_url"` // API URL to start this codespace.
	Idle_timeout_minutes int `json:"idle_timeout_minutes"` // The number of minutes of inactivity after which this codespace will be automatically stopped.
	Display_name string `json:"display_name,omitempty"` // Display name for this codespace.
	Billable_owner GeneratedType `json:"billable_owner"` // A GitHub user.
	Retention_period_minutes int `json:"retention_period_minutes,omitempty"` // Duration in minutes after codespace has gone idle in which it will be deleted. Must be integer minutes between 0 and 43200 (30 days).
	Machines_url string `json:"machines_url"` // API URL to access available alternate machine types for this codespace.
	Runtime_constraints map[string]interface{} `json:"runtime_constraints,omitempty"`
	Web_url string `json:"web_url"` // URL to access this codespace on the web.
	Created_at string `json:"created_at"`
	Pending_operation_disabled_reason string `json:"pending_operation_disabled_reason,omitempty"` // Text to show user when codespace is disabled by a pending operation
	Retention_expires_at string `json:"retention_expires_at,omitempty"` // When a codespace will be auto-deleted based on the "retention_period_minutes" and "last_used_at"
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Idle_timeout_notice string `json:"idle_timeout_notice,omitempty"` // Text to show user when codespace idle timeout minutes has been overriden by an organization policy
	Pulls_url string `json:"pulls_url"` // API URL for the Pull Request associated with this codespace, if any.
	Pending_operation bool `json:"pending_operation,omitempty"` // Whether or not a codespace has a pending async operation. This would mean that the codespace is temporarily unavailable. The only thing that you can do with a codespace in this state is delete it.
	Url string `json:"url"` // API URL for this codespace.
	Name string `json:"name"` // Automatically generated name of this codespace.
	Devcontainer_path string `json:"devcontainer_path,omitempty"` // Path to devcontainer.json from repo root used to create Codespace.
	Recent_folders []string `json:"recent_folders"`
	Publish_url string `json:"publish_url,omitempty"` // API URL to publish this codespace to a new repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Body string `json:"body,omitempty"`
	Html_url string `json:"html_url"`
	Key string `json:"key"`
	Name string `json:"name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Security_advisory map[string]interface{} `json:"security_advisory"` // The details of the security advisory, including summary, description, and severity.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project map[string]interface{} `json:"project"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sha string `json:"sha"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Run_duration_ms int `json:"run_duration_ms,omitempty"`
	Billable map[string]interface{} `json:"billable"`
}

// Collaborator represents the Collaborator schema from the OpenAPI specification
type Collaborator struct {
	Html_url string `json:"html_url"`
	Starred_url string `json:"starred_url"`
	Gravatar_id string `json:"gravatar_id"`
	Repos_url string `json:"repos_url"`
	Organizations_url string `json:"organizations_url"`
	Gists_url string `json:"gists_url"`
	Id int `json:"id"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Login string `json:"login"`
	Url string `json:"url"`
	Events_url string `json:"events_url"`
	Received_events_url string `json:"received_events_url"`
	Email string `json:"email,omitempty"`
	Followers_url string `json:"followers_url"`
	Subscriptions_url string `json:"subscriptions_url"`
	Role_name string `json:"role_name"`
	Following_url string `json:"following_url"`
	Name string `json:"name,omitempty"`
	Node_id string `json:"node_id"`
	Site_admin bool `json:"site_admin"`
	TypeField string `json:"type"`
	Avatar_url string `json:"avatar_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Marketplace_purchase interface{} `json:"marketplace_purchase"`
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
	Effective_date string `json:"effective_date"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Membership map[string]interface{} `json:"membership,omitempty"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Base map[string]interface{} `json:"base"`
	Head map[string]interface{} `json:"head"`
	Id int `json:"id"`
	Number int `json:"number"`
}

// Feed represents the Feed schema from the OpenAPI specification
type Feed struct {
	Timeline_url string `json:"timeline_url"`
	Current_user_url string `json:"current_user_url,omitempty"`
	Repository_discussions_category_url string `json:"repository_discussions_category_url,omitempty"` // A feed of discussions for a given repository and category.
	Links map[string]interface{} `json:"_links"`
	Current_user_organization_url string `json:"current_user_organization_url,omitempty"`
	Current_user_public_url string `json:"current_user_public_url,omitempty"`
	Repository_discussions_url string `json:"repository_discussions_url,omitempty"` // A feed of discussions for a given repository.
	Current_user_actor_url string `json:"current_user_actor_url,omitempty"`
	Current_user_organization_urls []string `json:"current_user_organization_urls,omitempty"`
	User_url string `json:"user_url"`
	Security_advisories_url string `json:"security_advisories_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Merged bool `json:"merged"`
	Message string `json:"message"`
	Sha string `json:"sha"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alt_domain map[string]interface{} `json:"alt_domain,omitempty"`
	Domain map[string]interface{} `json:"domain,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Wait_timer int `json:"wait_timer"` // The set duration of the wait timer
	Wait_timer_started_at string `json:"wait_timer_started_at"` // The time that the wait timer began.
	Current_user_can_approve bool `json:"current_user_can_approve"` // Whether the currently authenticated user can approve the deployment
	Environment map[string]interface{} `json:"environment"`
	Reviewers []map[string]interface{} `json:"reviewers"` // The people or teams that may approve jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Effective_date string `json:"effective_date"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Marketplace_purchase interface{} `json:"marketplace_purchase"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Reaction represents the Reaction schema from the OpenAPI specification
type Reaction struct {
	Content string `json:"content"` // The reaction to use
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	User GeneratedType `json:"user"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"` // The node ID of the pull request review comment.
	Position int `json:"position"` // The line index in the diff to which the comment applies. This field is deprecated; use `line` instead.
	Reactions GeneratedType `json:"reactions,omitempty"`
	Original_line int `json:"original_line,omitempty"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Pull_request_review_id int `json:"pull_request_review_id"` // The ID of the pull request review to which the comment belongs.
	Side string `json:"side,omitempty"` // The side of the diff to which the comment applies. The side of the last line of the range for a multi-line comment
	Url string `json:"url"` // URL for the pull request review comment
	Original_position int `json:"original_position"` // The index of the original line in the diff to which the comment applies. This field is deprecated; use `original_line` instead.
	Original_start_line int `json:"original_start_line,omitempty"` // The first line of the range for a multi-line comment.
	Original_commit_id string `json:"original_commit_id"` // The SHA of the original commit to which the comment applies.
	Html_url string `json:"html_url"` // HTML URL for the pull request review comment.
	Pull_request_url string `json:"pull_request_url"` // URL for the pull request that the review comment belongs to.
	User GeneratedType `json:"user"` // A GitHub user.
	Body_html string `json:"body_html,omitempty"`
	Path string `json:"path"` // The relative path of the file to which the comment applies.
	Body_text string `json:"body_text,omitempty"`
	Id int `json:"id"` // The ID of the pull request review comment.
	Commit_id string `json:"commit_id"` // The SHA of the commit to which the comment applies.
	In_reply_to_id int `json:"in_reply_to_id,omitempty"` // The comment ID to reply to.
	Body string `json:"body"` // The text of the comment.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Diff_hunk string `json:"diff_hunk"` // The diff of the line that the comment refers to.
	Start_side string `json:"start_side,omitempty"` // The side of the first line of the range for a multi-line comment.
	Links map[string]interface{} `json:"_links"`
	Line int `json:"line,omitempty"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Start_line int `json:"start_line,omitempty"` // The first line of the range for a multi-line comment.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Status string `json:"status"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Verification represents the Verification schema from the OpenAPI specification
type Verification struct {
	Signature string `json:"signature"`
	Verified bool `json:"verified"`
	Payload string `json:"payload"`
	Reason string `json:"reason"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commits []map[string]interface{} `json:"commits"` // An array of commit objects describing the pushed commits. (Pushed commits are all commits that are included in the `compare` between the `before` commit and the `after` commit.) The array includes a maximum of 20 commits. If necessary, you can use the [Commits API](https://docs.github.com/rest/reference/repos#commits) to fetch additional commits. This limit is applied to timeline events only and isn't applied to webhook deliveries.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	After string `json:"after"` // The SHA of the most recent commit on `ref` after the push.
	Compare string `json:"compare"` // URL that shows the changes in this `ref` update, from the `before` commit to the `after` commit. For a newly created `ref` that is directly based on the default branch, this is the comparison between the head of the default branch and the `after` commit. Otherwise, this shows all commits until the `after` commit.
	Pusher map[string]interface{} `json:"pusher"` // Metaproperties for Git author/committer information.
	Repository map[string]interface{} `json:"repository"` // A git repository
	Before string `json:"before"` // The SHA of the most recent commit on `ref` before the push.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Created bool `json:"created"` // Whether this push created the `ref`.
	Head_commit map[string]interface{} `json:"head_commit"`
	Deleted bool `json:"deleted"` // Whether this push deleted the `ref`.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Base_ref string `json:"base_ref"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Ref string `json:"ref"` // The full git ref that was pushed. Example: `refs/heads/main` or `refs/tags/v3.14.1`.
	Forced bool `json:"forced"` // Whether this push was a force push of the `ref`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Github_owned_allowed bool `json:"github_owned_allowed,omitempty"` // Whether GitHub-owned actions are allowed. For example, this includes the actions in the `actions` organization.
	Patterns_allowed []string `json:"patterns_allowed,omitempty"` // Specifies a list of string-matching patterns to allow specific action(s) and reusable workflow(s). Wildcards, tags, and SHAs are allowed. For example, `monalisa/octocat@*`, `monalisa/octocat@v2`, `monalisa/*`. **Note**: The `patterns_allowed` setting only applies to public repositories.
	Verified_allowed bool `json:"verified_allowed,omitempty"` // Whether actions from GitHub Marketplace verified creators are allowed. Set to `true` to allow all actions by GitHub Marketplace verified creators.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Permission string `json:"permission"` // Permission that the team will have for its repositories
	Url string `json:"url"` // URL for the team
	Ldap_dn string `json:"ldap_dn,omitempty"` // Distinguished Name (DN) that team maps to within LDAP environment
	Node_id string `json:"node_id"`
	Description string `json:"description"` // Description of the team
	Privacy string `json:"privacy,omitempty"` // The level of privacy this team should have
	Repositories_url string `json:"repositories_url"`
	Html_url string `json:"html_url"`
	Members_url string `json:"members_url"`
	Name string `json:"name"` // Name of the team
	Slug string `json:"slug"`
	Id int `json:"id"` // Unique identifier of the team
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Column_name string `json:"column_name,omitempty"`
	Content_url string `json:"content_url,omitempty"`
	Id int `json:"id"` // The project card's ID
	Project_id string `json:"project_id,omitempty"`
	Project_url string `json:"project_url"`
	Node_id string `json:"node_id"`
	Note string `json:"note"`
	Updated_at string `json:"updated_at"`
	Column_url string `json:"column_url"`
	Created_at string `json:"created_at"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Archived bool `json:"archived,omitempty"` // Whether or not the card is archived
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Rule map[string]interface{} `json:"rule"` // The branch protection rule. Includes a `name` and all the [branch protection settings](https://docs.github.com/github/administering-a-repository/defining-the-mergeability-of-pull-requests/about-protected-branches#about-branch-protection-settings) applied to branches that match the name. Binary settings are boolean. Multi-level configurations are one of `off`, `non_admins`, or `everyone`. Actor and build lists are arrays of strings.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Start_line int `json:"start_line,omitempty"`
	End_column int `json:"end_column,omitempty"`
	End_line int `json:"end_line,omitempty"`
	Path string `json:"path,omitempty"`
	Start_column int `json:"start_column,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ecosystem string `json:"ecosystem"` // The package's language or package management ecosystem.
	Name string `json:"name"` // The unique package name within its ecosystem.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Projects_v2 GeneratedType `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Organization GeneratedType `json:"organization"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name of the GitHub app
	Client_secret string `json:"client_secret,omitempty"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Pem string `json:"pem,omitempty"`
	External_url string `json:"external_url"`
	Installations_count int `json:"installations_count,omitempty"` // The number of installations associated with the GitHub app
	Permissions map[string]interface{} `json:"permissions"` // The set of permissions for the GitHub app
	Slug string `json:"slug,omitempty"` // The slug name of the GitHub app
	Client_id string `json:"client_id,omitempty"`
	Created_at string `json:"created_at"`
	Description string `json:"description"`
	Updated_at string `json:"updated_at"`
	Events []string `json:"events"` // The list of events for the GitHub app
	Node_id string `json:"node_id"`
	Webhook_secret string `json:"webhook_secret,omitempty"`
	Html_url string `json:"html_url"`
	Id int `json:"id"` // Unique identifier of the GitHub app
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Download_url string `json:"download_url"`
	TypeField string `json:"type"`
	Html_url string `json:"html_url"`
	Path string `json:"path"`
	Size int `json:"size"`
	Submodule_git_url string `json:"submodule_git_url"`
	Links map[string]interface{} `json:"_links"`
	Git_url string `json:"git_url"`
	Name string `json:"name"`
	Sha string `json:"sha"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Expiry string `json:"expiry,omitempty"` // The duration of the interaction restriction. Default: `one_day`.
	Limit string `json:"limit"` // The type of GitHub user that can comment, open issues, or create pull requests while the interaction limit is in effect.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name of the variable.
	Updated_at string `json:"updated_at"` // The date and time at which the variable was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Value string `json:"value"` // The value of the variable.
	Created_at string `json:"created_at"` // The date and time at which the variable was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_job map[string]interface{} `json:"workflow_job"`
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue interface{} `json:"issue"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Protected bool `json:"protected"`
	Protection GeneratedType `json:"protection"` // Branch Protection
	Protection_url string `json:"protection_url"`
	Required_approving_review_count int `json:"required_approving_review_count,omitempty"`
	Links map[string]interface{} `json:"_links"`
	Commit Commit `json:"commit"` // Commit
	Name string `json:"name"`
	Pattern string `json:"pattern,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Documentation_url string `json:"documentation_url,omitempty"`
	Message string `json:"message,omitempty"`
	Schemas []string `json:"schemas,omitempty"`
	Scimtype string `json:"scimType,omitempty"`
	Status int `json:"status,omitempty"`
	Detail string `json:"detail,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Required_conversation_resolution map[string]interface{} `json:"required_conversation_resolution,omitempty"`
	Required_linear_history map[string]interface{} `json:"required_linear_history,omitempty"`
	Url string `json:"url"`
	Allow_deletions map[string]interface{} `json:"allow_deletions,omitempty"`
	Required_signatures map[string]interface{} `json:"required_signatures,omitempty"`
	Lock_branch map[string]interface{} `json:"lock_branch,omitempty"` // Whether to set the branch as read-only. If this is true, users will not be able to push to the branch.
	Allow_fork_syncing map[string]interface{} `json:"allow_fork_syncing,omitempty"` // Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing.
	Required_pull_request_reviews map[string]interface{} `json:"required_pull_request_reviews,omitempty"`
	Required_status_checks GeneratedType `json:"required_status_checks,omitempty"` // Status Check Policy
	Restrictions GeneratedType `json:"restrictions,omitempty"` // Branch Restriction Policy
	Allow_force_pushes map[string]interface{} `json:"allow_force_pushes,omitempty"`
	Block_creations map[string]interface{} `json:"block_creations,omitempty"`
	Enforce_admins map[string]interface{} `json:"enforce_admins,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Subscriptions_url string `json:"subscriptions_url"`
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"`
	Id int `json:"id"`
	Events_url string `json:"events_url"`
	Repos_url string `json:"repos_url"`
	Avatar_url string `json:"avatar_url"`
	Following_url string `json:"following_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Followers_url string `json:"followers_url"`
	Html_url string `json:"html_url"`
	Starred_url string `json:"starred_url"`
	Url string `json:"url"`
	Login string `json:"login"`
	Node_id string `json:"node_id"`
	Site_admin bool `json:"site_admin"`
	Email string `json:"email,omitempty"`
	Organizations_url string `json:"organizations_url"`
	Role_name string `json:"role_name"`
	Gists_url string `json:"gists_url"`
	Gravatar_id string `json:"gravatar_id"`
	Received_events_url string `json:"received_events_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name of the check.
	Status string `json:"status"` // The phase of the lifecycle that the check is currently in.
	Pull_requests []GeneratedType `json:"pull_requests"`
	External_id string `json:"external_id"`
	Head_sha string `json:"head_sha"` // The SHA of the commit that is being checked.
	Started_at string `json:"started_at"`
	Completed_at string `json:"completed_at"`
	Details_url string `json:"details_url"`
	Output map[string]interface{} `json:"output"`
	Conclusion string `json:"conclusion"`
	Url string `json:"url"`
	App GeneratedType `json:"app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Node_id string `json:"node_id"`
	Check_suite map[string]interface{} `json:"check_suite"`
	Deployment GeneratedType `json:"deployment,omitempty"` // A deployment created as the result of an Actions check run from a workflow that references an environment
	Html_url string `json:"html_url"`
	Id int `json:"id"` // The id of the check.
}

// Autolink represents the Autolink schema from the OpenAPI specification
type Autolink struct {
	Id int `json:"id"`
	Is_alphanumeric bool `json:"is_alphanumeric"` // Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters.
	Key_prefix string `json:"key_prefix"` // The prefix of a key that is linkified.
	Url_template string `json:"url_template"` // A template for the target URL that is generated if a key was found.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Starred_at string `json:"starred_at"` // The time the star was created. This is a timestamp in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`. Will be `null` for the `deleted` action.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	User GeneratedType `json:"user"` // A GitHub user.
	Version string `json:"version"`
	Change_status map[string]interface{} `json:"change_status"`
	Committed_at string `json:"committed_at"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Reason string `json:"reason,omitempty"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"`
	Pull_request map[string]interface{} `json:"pull_request"`
	Action string `json:"action"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Milestone Milestone `json:"milestone,omitempty"` // A collection of related issues and pull requests.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// Authorization represents the Authorization schema from the OpenAPI specification
type Authorization struct {
	Url string `json:"url"`
	Scopes []string `json:"scopes"` // A list of scopes that this authorization is in.
	Installation GeneratedType `json:"installation,omitempty"`
	Expires_at string `json:"expires_at"`
	User GeneratedType `json:"user,omitempty"` // A GitHub user.
	Hashed_token string `json:"hashed_token"`
	Token_last_eight string `json:"token_last_eight"`
	Token string `json:"token"`
	Note string `json:"note"`
	Note_url string `json:"note_url"`
	App map[string]interface{} `json:"app"`
	Fingerprint string `json:"fingerprint"`
	Updated_at string `json:"updated_at"`
	Id int `json:"id"`
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) the comment belongs to.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/reference/issues#comments) itself.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Action string `json:"action"`
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
}

// Release represents the Release schema from the OpenAPI specification
type Release struct {
	Tarball_url string `json:"tarball_url"`
	Upload_url string `json:"upload_url"`
	Body_text string `json:"body_text,omitempty"`
	Created_at string `json:"created_at"`
	Name string `json:"name"`
	Tag_name string `json:"tag_name"` // The name of the tag.
	Discussion_url string `json:"discussion_url,omitempty"` // The URL of the release discussion.
	Published_at string `json:"published_at"`
	Author GeneratedType `json:"author"` // A GitHub user.
	Body_html string `json:"body_html,omitempty"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Body string `json:"body,omitempty"`
	Url string `json:"url"`
	Id int `json:"id"`
	Target_commitish string `json:"target_commitish"` // Specifies the commitish value that determines where the Git tag is created from.
	Zipball_url string `json:"zipball_url"`
	Assets_url string `json:"assets_url"`
	Mentions_count int `json:"mentions_count,omitempty"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Assets []GeneratedType `json:"assets"`
	Draft bool `json:"draft"` // true to create a draft (unpublished) release, false to create a published one.
	Prerelease bool `json:"prerelease"` // Whether to identify the release as a prerelease or a full release.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action,omitempty"`
	Check_run GeneratedType `json:"check_run"` // A check performed on the code of a given code change
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type,omitempty"` // The type of label. Read-only labels are applied automatically when the runner is configured.
	Id int `json:"id,omitempty"` // Unique identifier of the label.
	Name string `json:"name"` // Name of the label.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Performed_via_github_app Integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Assigner GeneratedType `json:"assigner"` // A GitHub user.
	Event string `json:"event"`
	Commit_url string `json:"commit_url"`
	Id int `json:"id"`
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Commit map[string]interface{} `json:"commit"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Tarball_url string `json:"tarball_url"`
	Zipball_url string `json:"zipball_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total_active_caches_count int `json:"total_active_caches_count"` // The count of active caches across all repositories of an enterprise or an organization.
	Total_active_caches_size_in_bytes int `json:"total_active_caches_size_in_bytes"` // The total size in bytes of all active cache items across all repositories of an enterprise or an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
}

// Project represents the Project schema from the OpenAPI specification
type Project struct {
	Name string `json:"name"` // Name of the project
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Body string `json:"body"` // Body of the project
	Updated_at string `json:"updated_at"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Columns_url string `json:"columns_url"`
	Html_url string `json:"html_url"`
	Organization_permission string `json:"organization_permission,omitempty"` // The baseline permission that all organization members have on this project. Only present if owner is an organization.
	Number int `json:"number"`
	Owner_url string `json:"owner_url"`
	Id int `json:"id"`
	Private bool `json:"private,omitempty"` // Whether or not this project can be seen by everyone. Only present if owner is an organization.
	Created_at string `json:"created_at"`
	State string `json:"state"` // State of the project; either 'open' or 'closed'
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_id string `json:"commit_id"`
	Event string `json:"event"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Id int `json:"id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
}

// Installation represents the Installation schema from the OpenAPI specification
type Installation struct {
	Single_file_paths []string `json:"single_file_paths,omitempty"`
	Created_at string `json:"created_at"`
	Permissions GeneratedType `json:"permissions"` // The permissions granted to the user-to-server access token.
	Id int `json:"id"` // The ID of the installation.
	Repositories_url string `json:"repositories_url"`
	Account interface{} `json:"account"`
	Access_tokens_url string `json:"access_tokens_url"`
	Suspended_at string `json:"suspended_at"`
	Events []string `json:"events"`
	App_id int `json:"app_id"`
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Target_id int `json:"target_id"` // The ID of the user or organization this token is being scoped to.
	Has_multiple_single_files bool `json:"has_multiple_single_files,omitempty"`
	Target_type string `json:"target_type"`
	Suspended_by GeneratedType `json:"suspended_by"` // A GitHub user.
	App_slug string `json:"app_slug"`
	Single_file_name string `json:"single_file_name"`
	Updated_at string `json:"updated_at"`
	Contact_email string `json:"contact_email,omitempty"`
	Html_url string `json:"html_url"`
}

// Traffic represents the Traffic schema from the OpenAPI specification
type Traffic struct {
	Count int `json:"count"`
	Timestamp string `json:"timestamp"`
	Uniques int `json:"uniques"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Release map[string]interface{} `json:"release"` // The [release](https://docs.github.com/rest/reference/repos/#get-a-release) object.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alert map[string]interface{} `json:"alert"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// Issue represents the Issue schema from the OpenAPI specification
type Issue struct {
	Title string `json:"title"` // Title of the issue
	Body_html string `json:"body_html,omitempty"`
	Repository_url string `json:"repository_url"`
	Body string `json:"body,omitempty"` // Contents of the issue
	Comments int `json:"comments"`
	Number int `json:"number"` // Number uniquely identifying the issue within its repository
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Id int `json:"id"`
	Closed_by GeneratedType `json:"closed_by,omitempty"` // A GitHub user.
	Html_url string `json:"html_url"`
	Timeline_url string `json:"timeline_url,omitempty"`
	Draft bool `json:"draft,omitempty"`
	State string `json:"state"` // State of the issue; either 'open' or 'closed'
	Url string `json:"url"` // URL for the issue
	Locked bool `json:"locked"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Comments_url string `json:"comments_url"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Updated_at string `json:"updated_at"`
	Closed_at string `json:"closed_at"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Labels_url string `json:"labels_url"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	User GeneratedType `json:"user"` // A GitHub user.
	Labels []interface{} `json:"labels"` // Labels to associate with this issue; pass one or more label names to replace the set of labels on this issue; send an empty array to clear all labels from the issue; note that the labels are silently dropped for users without push access to the repository
	Body_text string `json:"body_text,omitempty"`
	State_reason string `json:"state_reason,omitempty"` // The reason for the current state
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Events_url string `json:"events_url"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Security_severity_level string `json:"security_severity_level,omitempty"` // The security severity of the alert.
	Full_description string `json:"full_description,omitempty"` // description of the rule used to detect the alert.
	Help string `json:"help,omitempty"` // Detailed documentation for the rule as GitHub Flavored Markdown.
	Severity string `json:"severity,omitempty"` // The severity of the alert.
	Description string `json:"description,omitempty"` // A short description of the rule used to detect the alert.
	Help_uri string `json:"help_uri,omitempty"` // A link to the documentation for the rule used to detect the alert.
	Name string `json:"name,omitempty"` // The name of the rule used to detect the alert.
	Id string `json:"id,omitempty"` // A unique identifier for the rule used to detect the alert.
	Tags []string `json:"tags,omitempty"` // A set of tags applicable for the rule.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Errors []map[string]interface{} `json:"errors"`
}

// Repository represents the Repository schema from the OpenAPI specification
type Repository struct {
	Language string `json:"language"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Mirror_url string `json:"mirror_url"`
	Subscribers_url string `json:"subscribers_url"`
	Html_url string `json:"html_url"`
	Keys_url string `json:"keys_url"`
	License GeneratedType `json:"license"` // License Simple
	Node_id string `json:"node_id"`
	Stargazers_url string `json:"stargazers_url"`
	Hooks_url string `json:"hooks_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Network_count int `json:"network_count,omitempty"`
	Languages_url string `json:"languages_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Assignees_url string `json:"assignees_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Open_issues int `json:"open_issues"`
	Fork bool `json:"fork"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property has been deprecated. Please use `squash_merge_commit_title` instead.
	Description string `json:"description"`
	Pushed_at string `json:"pushed_at"`
	Clone_url string `json:"clone_url"`
	Milestones_url string `json:"milestones_url"`
	Name string `json:"name"` // The name of the repository.
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Issues_url string `json:"issues_url"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Commits_url string `json:"commits_url"`
	Contents_url string `json:"contents_url"`
	Forks_url string `json:"forks_url"`
	Comments_url string `json:"comments_url"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub user.
	Deployments_url string `json:"deployments_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Watchers int `json:"watchers"`
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Svn_url string `json:"svn_url"`
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Blobs_url string `json:"blobs_url"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Labels_url string `json:"labels_url"`
	Statuses_url string `json:"statuses_url"`
	Homepage string `json:"homepage"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Trees_url string `json:"trees_url"`
	Notifications_url string `json:"notifications_url"`
	Archive_url string `json:"archive_url"`
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Starred_at string `json:"starred_at,omitempty"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Url string `json:"url"`
	Compare_url string `json:"compare_url"`
	Topics []string `json:"topics,omitempty"`
	Tags_url string `json:"tags_url"`
	Merges_url string `json:"merges_url"`
	Git_url string `json:"git_url"`
	Releases_url string `json:"releases_url"`
	Size int `json:"size"` // The size of the repository. Size is calculated hourly. When a repository is initially created, the size is 0.
	Pulls_url string `json:"pulls_url"`
	Id int `json:"id"` // Unique identifier of the repository
	Subscription_url string `json:"subscription_url"`
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Downloads_url string `json:"downloads_url"`
	Git_commits_url string `json:"git_commits_url"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Events_url string `json:"events_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Git_refs_url string `json:"git_refs_url"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Forks int `json:"forks"`
	Stargazers_count int `json:"stargazers_count"`
	Teams_url string `json:"teams_url"`
	Contributors_url string `json:"contributors_url"`
	Watchers_count int `json:"watchers_count"`
	Master_branch string `json:"master_branch,omitempty"`
	Template_repository map[string]interface{} `json:"template_repository,omitempty"`
	Private bool `json:"private"` // Whether the repository is private or public.
	Issue_events_url string `json:"issue_events_url"`
	Open_issues_count int `json:"open_issues_count"`
	Updated_at string `json:"updated_at"`
	Git_tags_url string `json:"git_tags_url"`
	Created_at string `json:"created_at"`
	Forks_count int `json:"forks_count"`
	Ssh_url string `json:"ssh_url"`
	Has_pages bool `json:"has_pages"`
	Collaborators_url string `json:"collaborators_url"`
	Branches_url string `json:"branches_url"`
	Full_name string `json:"full_name"`
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship map[string]interface{} `json:"sponsorship"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id,omitempty"`
	Comments []GeneratedType `json:"comments,omitempty"`
	Commit_id string `json:"commit_id,omitempty"`
	Event string `json:"event,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"`
	Protected bool `json:"protected"`
	Commit map[string]interface{} `json:"commit"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Columns_url string `json:"columns_url"`
	Number int `json:"number"`
	Permissions map[string]interface{} `json:"permissions"`
	Url string `json:"url"`
	Updated_at string `json:"updated_at"`
	Owner_url string `json:"owner_url"`
	State string `json:"state"`
	Id int `json:"id"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Private bool `json:"private,omitempty"` // Whether the project is private or not. Only present when owner is an organization.
	Body string `json:"body"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Organization_permission string `json:"organization_permission,omitempty"` // The organization permission for this project. Only present when owner is an organization.
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Status string `json:"status"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Action string `json:"action"`
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Days []int `json:"days"`
	Total int `json:"total"`
	Week int `json:"week"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Comment map[string]interface{} `json:"comment"` // The [commit comment](https://docs.github.com/rest/reference/repos#get-a-commit-comment) resource.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action performed. Can be `created`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Merged bool `json:"merged"`
	Requested_reviewers []GeneratedType `json:"requested_reviewers,omitempty"`
	Created_at string `json:"created_at"`
	Requested_teams []GeneratedType `json:"requested_teams,omitempty"`
	Deletions int `json:"deletions"`
	Merged_by GeneratedType `json:"merged_by"` // A GitHub user.
	Additions int `json:"additions"`
	Review_comments_url string `json:"review_comments_url"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Node_id string `json:"node_id"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Url string `json:"url"`
	Labels []map[string]interface{} `json:"labels"`
	Commits_url string `json:"commits_url"`
	Patch_url string `json:"patch_url"`
	Auto_merge GeneratedType `json:"auto_merge"` // The status of auto merging a pull request.
	Mergeable bool `json:"mergeable"`
	Html_url string `json:"html_url"`
	Title string `json:"title"` // The title of the pull request.
	Review_comment_url string `json:"review_comment_url"`
	Updated_at string `json:"updated_at"`
	Head map[string]interface{} `json:"head"`
	Statuses_url string `json:"statuses_url"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Comments int `json:"comments"`
	Issue_url string `json:"issue_url"`
	Links map[string]interface{} `json:"_links"`
	Commits int `json:"commits"`
	Mergeable_state string `json:"mergeable_state"`
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Base map[string]interface{} `json:"base"`
	Rebaseable bool `json:"rebaseable,omitempty"`
	Merge_commit_sha string `json:"merge_commit_sha"`
	Maintainer_can_modify bool `json:"maintainer_can_modify"` // Indicates whether maintainers can modify the pull request.
	Body string `json:"body"`
	Locked bool `json:"locked"`
	Changed_files int `json:"changed_files"`
	Diff_url string `json:"diff_url"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Comments_url string `json:"comments_url"`
	Draft bool `json:"draft,omitempty"` // Indicates whether or not the pull request is a draft.
	State string `json:"state"` // State of this Pull Request. Either `open` or `closed`.
	User GeneratedType `json:"user"` // A GitHub user.
	Number int `json:"number"` // Number uniquely identifying the pull request within its repository.
	Merged_at string `json:"merged_at"`
	Closed_at string `json:"closed_at"`
	Id int `json:"id"`
	Review_comments int `json:"review_comments"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Effective_date string `json:"effective_date,omitempty"` // The `pending_cancellation` and `pending_tier_change` event types will include the date the cancellation or tier change will take effect.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship map[string]interface{} `json:"sponsorship"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // Full Repository
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action,omitempty"`
	Alert GeneratedType `json:"alert"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Location GeneratedType `json:"location"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Number int `json:"number,omitempty"` // The security alert number.
	Resolution_comment string `json:"resolution_comment,omitempty"` // The comment that was optionally added when this alert was closed
	Secret string `json:"secret,omitempty"` // The secret that was detected.
	Created_at string `json:"created_at,omitempty"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Resolution string `json:"resolution,omitempty"` // **Required when the `state` is `resolved`.** The reason for resolving the alert.
	Repository GeneratedType `json:"repository,omitempty"` // A GitHub repository.
	Resolved_by GeneratedType `json:"resolved_by,omitempty"` // A GitHub user.
	Push_protection_bypassed bool `json:"push_protection_bypassed,omitempty"` // Whether push protection was bypassed for the detected secret.
	Push_protection_bypassed_at string `json:"push_protection_bypassed_at,omitempty"` // The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Secret_type string `json:"secret_type,omitempty"` // The type of secret that secret scanning detected.
	State string `json:"state,omitempty"` // Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
	Url string `json:"url,omitempty"` // The REST API URL of the alert resource.
	Locations_url string `json:"locations_url,omitempty"` // The REST API URL of the code locations for this alert.
	Html_url string `json:"html_url,omitempty"` // The GitHub URL of the alert resource.
	Secret_type_display_name string `json:"secret_type_display_name,omitempty"` // User-friendly name for the detected secret, matching the `secret_type`. For a list of built-in patterns, see "[Secret scanning patterns](https://docs.github.com/code-security/secret-scanning/secret-scanning-patterns#supported-secrets-for-advanced-security)."
	Push_protection_bypassed_by GeneratedType `json:"push_protection_bypassed_by,omitempty"` // A GitHub user.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Resolved_at string `json:"resolved_at,omitempty"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url,omitempty"`
	Key string `json:"key"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Spdx_id string `json:"spdx_id"`
	Url string `json:"url"`
}

// Artifact represents the Artifact schema from the OpenAPI specification
type Artifact struct {
	Id int `json:"id"`
	Size_in_bytes int `json:"size_in_bytes"` // The size in bytes of the artifact.
	Updated_at string `json:"updated_at"`
	Workflow_run map[string]interface{} `json:"workflow_run,omitempty"`
	Created_at string `json:"created_at"`
	Expired bool `json:"expired"` // Whether or not the artifact has expired.
	Url string `json:"url"`
	Archive_download_url string `json:"archive_download_url"`
	Expires_at string `json:"expires_at"`
	Name string `json:"name"` // The name of the artifact.
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Email string `json:"email,omitempty"`
	Id int `json:"id"`
	Login string `json:"login"`
	Node_id string `json:"node_id,omitempty"`
	Organization_billing_email string `json:"organization_billing_email,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow map[string]interface{} `json:"workflow"`
	Workflow_run interface{} `json:"workflow_run"`
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Export_url string `json:"export_url,omitempty"` // Url for fetching export details
	Html_url string `json:"html_url,omitempty"` // Web url for the exported branch
	Id string `json:"id,omitempty"` // Id for the export details
	Sha string `json:"sha,omitempty"` // Git commit SHA of the exported branch
	State string `json:"state,omitempty"` // State of the latest export
	Branch string `json:"branch,omitempty"` // Name of the exported branch
	Completed_at string `json:"completed_at,omitempty"` // Completion time of the last export operation
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Login string `json:"login"`
	Url string `json:"url"`
	Avatar_url string `json:"avatar_url"`
	Organizations_url string `json:"organizations_url"`
	Starred_at string `json:"starred_at,omitempty"`
	Email string `json:"email,omitempty"`
	Gists_url string `json:"gists_url"`
	Gravatar_id string `json:"gravatar_id"`
	Repos_url string `json:"repos_url"`
	Starred_url string `json:"starred_url"`
	Node_id string `json:"node_id"`
	Id int `json:"id"`
	Html_url string `json:"html_url"`
	TypeField string `json:"type"`
	Subscriptions_url string `json:"subscriptions_url"`
	Following_url string `json:"following_url"`
	Name string `json:"name,omitempty"`
	Received_events_url string `json:"received_events_url"`
	Events_url string `json:"events_url"`
	Site_admin bool `json:"site_admin"`
	Followers_url string `json:"followers_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Secret_scanning_alerts string `json:"secret_scanning_alerts,omitempty"` // The level of permission to grant the access token to view and manage secret scanning alerts.
	Packages string `json:"packages,omitempty"` // The level of permission to grant the access token for packages published to GitHub Packages.
	Members string `json:"members,omitempty"` // The level of permission to grant the access token for organization teams and members.
	Organization_plan string `json:"organization_plan,omitempty"` // The level of permission to grant the access token for viewing an organization's plan.
	Statuses string `json:"statuses,omitempty"` // The level of permission to grant the access token for commit statuses.
	Vulnerability_alerts string `json:"vulnerability_alerts,omitempty"` // The level of permission to grant the access token to manage Dependabot alerts.
	Single_file string `json:"single_file,omitempty"` // The level of permission to grant the access token to manage just a single file.
	Organization_projects string `json:"organization_projects,omitempty"` // The level of permission to grant the access token to manage organization projects and projects beta (where available).
	Organization_packages string `json:"organization_packages,omitempty"` // The level of permission to grant the access token for organization packages published to GitHub Packages.
	Pull_requests string `json:"pull_requests,omitempty"` // The level of permission to grant the access token for pull requests and related comments, assignees, labels, milestones, and merges.
	Repository_hooks string `json:"repository_hooks,omitempty"` // The level of permission to grant the access token to manage the post-receive hooks for a repository.
	Issues string `json:"issues,omitempty"` // The level of permission to grant the access token for issues and related comments, assignees, labels, and milestones.
	Repository_projects string `json:"repository_projects,omitempty"` // The level of permission to grant the access token to manage repository projects, columns, and cards.
	Team_discussions string `json:"team_discussions,omitempty"` // The level of permission to grant the access token to manage team discussions and related comments.
	Repository_announcement_banners string `json:"repository_announcement_banners,omitempty"` // The level of permission to grant the access token to view and manage announcement banners for a repository.
	Security_events string `json:"security_events,omitempty"` // The level of permission to grant the access token to view and manage security events like code scanning alerts.
	Organization_secrets string `json:"organization_secrets,omitempty"` // The level of permission to grant the access token to manage organization secrets.
	Deployments string `json:"deployments,omitempty"` // The level of permission to grant the access token for deployments and deployment statuses.
	Pages string `json:"pages,omitempty"` // The level of permission to grant the access token to retrieve Pages statuses, configuration, and builds, as well as create new builds.
	Organization_user_blocking string `json:"organization_user_blocking,omitempty"` // The level of permission to grant the access token to view and manage users blocked by the organization.
	Environments string `json:"environments,omitempty"` // The level of permission to grant the access token for managing repository environments.
	Organization_administration string `json:"organization_administration,omitempty"` // The level of permission to grant the access token to manage access to an organization.
	Actions string `json:"actions,omitempty"` // The level of permission to grant the access token for GitHub Actions workflows, workflow runs, and artifacts.
	Workflows string `json:"workflows,omitempty"` // The level of permission to grant the access token to update GitHub Actions workflow files.
	Organization_self_hosted_runners string `json:"organization_self_hosted_runners,omitempty"` // The level of permission to grant the access token to view and manage GitHub Actions self-hosted runners available to an organization.
	Contents string `json:"contents,omitempty"` // The level of permission to grant the access token for repository contents, commits, branches, downloads, releases, and merges.
	Organization_custom_roles string `json:"organization_custom_roles,omitempty"` // The level of permission to grant the access token for custom repository roles management. This property is in beta and is subject to change.
	Checks string `json:"checks,omitempty"` // The level of permission to grant the access token for checks on code.
	Organization_announcement_banners string `json:"organization_announcement_banners,omitempty"` // The level of permission to grant the access token to view and manage announcement banners for an organization.
	Organization_hooks string `json:"organization_hooks,omitempty"` // The level of permission to grant the access token to manage the post-receive hooks for an organization.
	Metadata string `json:"metadata,omitempty"` // The level of permission to grant the access token to search repositories, list collaborators, and access repository metadata.
	Secrets string `json:"secrets,omitempty"` // The level of permission to grant the access token to manage repository secrets.
	Administration string `json:"administration,omitempty"` // The level of permission to grant the access token for repository creation, deletion, settings, teams, and collaborators creation.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Users_url string `json:"users_url"`
	Apps []map[string]interface{} `json:"apps"`
	Apps_url string `json:"apps_url"`
	Teams []map[string]interface{} `json:"teams"`
	Teams_url string `json:"teams_url"`
	Url string `json:"url"`
	Users []map[string]interface{} `json:"users"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Team map[string]interface{} `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
}

// Topic represents the Topic schema from the OpenAPI specification
type Topic struct {
	Names []string `json:"names"`
}

// Dependency represents the Dependency schema from the OpenAPI specification
type Dependency struct {
	Metadata Metadata `json:"metadata,omitempty"` // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
	Package_url string `json:"package_url,omitempty"` // Package-url (PURL) of dependency. See https://github.com/package-url/purl-spec for more details.
	Relationship string `json:"relationship,omitempty"` // A notation of whether a dependency is requested directly by this manifest or is a dependency of another dependency.
	Scope string `json:"scope,omitempty"` // A notation of whether the dependency is required for the primary build artifact (runtime) or is only used for development. Future versions of this specification may allow for more granular scopes.
	Dependencies []string `json:"dependencies,omitempty"` // Array of package-url (PURLs) of direct child dependencies.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project_card map[string]interface{} `json:"project_card"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action,omitempty"`
	Check_run GeneratedType `json:"check_run"` // A check performed on the code of a given code change
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project map[string]interface{} `json:"project"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Registry_package map[string]interface{} `json:"registry_package"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship map[string]interface{} `json:"sponsorship"`
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Requested_team Team `json:"requested_team,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Created_at string `json:"created_at"`
	Requested_reviewer GeneratedType `json:"requested_reviewer,omitempty"` // A GitHub user.
	Review_requester GeneratedType `json:"review_requester"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Event string `json:"event"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Commit_url string `json:"commit_url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Review map[string]interface{} `json:"review"` // The review that was affected.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Pull_request map[string]interface{} `json:"pull_request"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Number int `json:"number"`
	Reason string `json:"reason"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pages []map[string]interface{} `json:"pages"` // The pages that were updated.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Milestone represents the Milestone schema from the OpenAPI specification
type Milestone struct {
	Created_at string `json:"created_at"`
	Due_on string `json:"due_on"`
	Closed_at string `json:"closed_at"`
	State string `json:"state"` // The state of the milestone.
	Closed_issues int `json:"closed_issues"`
	Id int `json:"id"`
	Labels_url string `json:"labels_url"`
	Url string `json:"url"`
	Open_issues int `json:"open_issues"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Html_url string `json:"html_url"`
	Number int `json:"number"` // The number of the milestone.
	Updated_at string `json:"updated_at"`
	Description string `json:"description"`
	Node_id string `json:"node_id"`
	Title string `json:"title"` // The title of the milestone.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2 GeneratedType `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Expires_at string `json:"expires_at"`
	Limit string `json:"limit"` // The type of GitHub user that can comment, open issues, or create pull requests while the interaction limit is in effect.
	Origin string `json:"origin"`
}

// Link represents the Link schema from the OpenAPI specification
type Link struct {
	Href string `json:"href"`
}

// Status represents the Status schema from the OpenAPI specification
type Status struct {
	Creator GeneratedType `json:"creator"` // A GitHub user.
	State string `json:"state"`
	Url string `json:"url"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Avatar_url string `json:"avatar_url"`
	Context string `json:"context"`
	Description string `json:"description"`
	Target_url string `json:"target_url"`
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Requester interface{} `json:"requester,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key_id string `json:"key_id"` // The identifier for the key.
	Title string `json:"title,omitempty"`
	Url string `json:"url,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Id int `json:"id,omitempty"`
	Key string `json:"key"` // The Base64 encoded public key.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Limit int `json:"limit"`
	Remaining int `json:"remaining"`
	Reset int `json:"reset"`
	Used int `json:"used"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Label map[string]interface{} `json:"label,omitempty"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the issue.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Label map[string]interface{} `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Visibility string `json:"visibility"` // Visibility of a secret
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes,omitempty"` // If the action was `edited`, the changes to the rule.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Rule map[string]interface{} `json:"rule"` // The branch protection rule. Includes a `name` and all the [branch protection settings](https://docs.github.com/github/administering-a-repository/defining-the-mergeability-of-pull-requests/about-protected-branches#about-branch-protection-settings) applied to branches that match the name. Binary settings are boolean. Multi-level configurations are one of `off`, `non_admins`, or `everyone`. Actor and build lists are arrays of strings.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// Team represents the Team schema from the OpenAPI specification
type Team struct {
	Members_url string `json:"members_url"`
	Node_id string `json:"node_id"`
	Permission string `json:"permission"`
	Repositories_url string `json:"repositories_url"`
	Html_url string `json:"html_url"`
	Parent GeneratedType `json:"parent"` // Groups of organization members that gives permissions on specified repositories.
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Privacy string `json:"privacy,omitempty"`
	Url string `json:"url"`
	Description string `json:"description"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Milestone Milestone `json:"milestone,omitempty"` // A collection of related issues and pull requests.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Event string `json:"event,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Comments []GeneratedType `json:"comments,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commits_url string `json:"commits_url,omitempty"`
	Description string `json:"description,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Owner GeneratedType `json:"owner,omitempty"` // A GitHub user.
	Forks []map[string]interface{} `json:"forks,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Fork_of map[string]interface{} `json:"fork_of,omitempty"` // Gist
	History []GeneratedType `json:"history,omitempty"`
	User string `json:"user,omitempty"`
	Id string `json:"id,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Public bool `json:"public,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	Comments_url string `json:"comments_url,omitempty"`
	Git_push_url string `json:"git_push_url,omitempty"`
	Truncated bool `json:"truncated,omitempty"`
	Comments int `json:"comments,omitempty"`
	Forks_url string `json:"forks_url,omitempty"`
	Git_pull_url string `json:"git_pull_url,omitempty"`
	Url string `json:"url,omitempty"`
	Files map[string]interface{} `json:"files,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Path string `json:"path"`
	Ref string `json:"ref,omitempty"`
	Sha string `json:"sha"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_job interface{} `json:"workflow_job"`
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// Enterprise represents the Enterprise schema from the OpenAPI specification
type Enterprise struct {
	Updated_at string `json:"updated_at"`
	Description string `json:"description,omitempty"` // A short description of the enterprise.
	Id int `json:"id"` // Unique identifier of the enterprise
	Slug string `json:"slug"` // The slug url identifier for the enterprise.
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the enterprise.
	Node_id string `json:"node_id"`
	Website_url string `json:"website_url,omitempty"` // The enterprise's website URL.
	Avatar_url string `json:"avatar_url"`
	Html_url string `json:"html_url"`
}

// Key represents the Key schema from the OpenAPI specification
type Key struct {
	Url string `json:"url"`
	Verified bool `json:"verified"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Key string `json:"key"`
	Read_only bool `json:"read_only"`
	Title string `json:"title"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name of the check.
	Url string `json:"url"`
	Id int `json:"id"` // The id of the check.
	Details_url string `json:"details_url"`
	Pull_requests []GeneratedType `json:"pull_requests"`
	Started_at string `json:"started_at"`
	Deployment GeneratedType `json:"deployment,omitempty"` // A deployment created as the result of an Actions check run from a workflow that references an environment
	External_id string `json:"external_id"`
	Html_url string `json:"html_url"`
	Status string `json:"status"` // The phase of the lifecycle that the check is currently in.
	Conclusion string `json:"conclusion"`
	Check_suite GeneratedType `json:"check_suite"` // A suite of checks performed on the code of a given code change
	Completed_at string `json:"completed_at"`
	App GeneratedType `json:"app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Head_sha string `json:"head_sha"` // The SHA of the commit that is being checked.
	Node_id string `json:"node_id"`
	Output map[string]interface{} `json:"output"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url,omitempty"` // The URL to which the payloads will be delivered.
	Content_type string `json:"content_type,omitempty"` // The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
	Insecure_ssl GeneratedType `json:"insecure_ssl,omitempty"`
	Secret string `json:"secret,omitempty"` // If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers).
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Deployment map[string]interface{} `json:"deployment"` // The [deployment](https://docs.github.com/rest/reference/deployments#list-deployments).
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Workflow map[string]interface{} `json:"workflow"`
	Action string `json:"action"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Committed_at string `json:"committed_at,omitempty"`
	Url string `json:"url,omitempty"`
	User GeneratedType `json:"user,omitempty"` // A GitHub user.
	Version string `json:"version,omitempty"`
	Change_status map[string]interface{} `json:"change_status,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Alert interface{} `json:"alert"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Node_id string `json:"node_id"`
	Event string `json:"event"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Rename map[string]interface{} `json:"rename"`
	Commit_url string `json:"commit_url"`
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Forks_url string `json:"forks_url"`
	Pushed_at string `json:"pushed_at,omitempty"`
	Archived bool `json:"archived,omitempty"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Languages_url string `json:"languages_url"`
	Git_commits_url string `json:"git_commits_url"`
	Blobs_url string `json:"blobs_url"`
	Keys_url string `json:"keys_url"`
	Notifications_url string `json:"notifications_url"`
	Homepage string `json:"homepage,omitempty"`
	Collaborators_url string `json:"collaborators_url"`
	Svn_url string `json:"svn_url,omitempty"`
	Contents_url string `json:"contents_url"`
	Watchers_count int `json:"watchers_count,omitempty"`
	Issues_url string `json:"issues_url"`
	Private bool `json:"private"`
	Topics []string `json:"topics,omitempty"`
	Clone_url string `json:"clone_url,omitempty"`
	Forks_count int `json:"forks_count,omitempty"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Trees_url string `json:"trees_url"`
	Archive_url string `json:"archive_url"`
	Branches_url string `json:"branches_url"`
	Open_issues_count int `json:"open_issues_count,omitempty"`
	Role_name string `json:"role_name,omitempty"`
	Git_url string `json:"git_url,omitempty"`
	Subscribers_url string `json:"subscribers_url"`
	Html_url string `json:"html_url"`
	Merges_url string `json:"merges_url"`
	Stargazers_url string `json:"stargazers_url"`
	Has_issues bool `json:"has_issues,omitempty"`
	Url string `json:"url"`
	Language string `json:"language,omitempty"`
	Forks int `json:"forks,omitempty"`
	Has_projects bool `json:"has_projects,omitempty"`
	Open_issues int `json:"open_issues,omitempty"`
	Ssh_url string `json:"ssh_url,omitempty"`
	Issue_events_url string `json:"issue_events_url"`
	Name string `json:"name"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Fork bool `json:"fork"`
	Full_name string `json:"full_name"`
	Git_tags_url string `json:"git_tags_url"`
	Labels_url string `json:"labels_url"`
	Assignees_url string `json:"assignees_url"`
	Security_and_analysis GeneratedType `json:"security_and_analysis,omitempty"`
	Disabled bool `json:"disabled,omitempty"`
	Releases_url string `json:"releases_url"`
	Network_count int `json:"network_count,omitempty"`
	Description string `json:"description"`
	Comments_url string `json:"comments_url"`
	Events_url string `json:"events_url"`
	Pulls_url string `json:"pulls_url"`
	Size int `json:"size,omitempty"` // The size of the repository. Size is calculated hourly. When a repository is initially created, the size is 0.
	License map[string]interface{} `json:"license,omitempty"`
	Tags_url string `json:"tags_url"`
	Stargazers_count int `json:"stargazers_count,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	Deployments_url string `json:"deployments_url"`
	Has_wiki bool `json:"has_wiki,omitempty"`
	Milestones_url string `json:"milestones_url"`
	Node_id string `json:"node_id"`
	Visibility string `json:"visibility,omitempty"`
	Code_of_conduct GeneratedType `json:"code_of_conduct,omitempty"` // Code Of Conduct
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Mirror_url string `json:"mirror_url,omitempty"`
	Has_discussions bool `json:"has_discussions,omitempty"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Teams_url string `json:"teams_url"`
	Hooks_url string `json:"hooks_url"`
	Downloads_url string `json:"downloads_url"`
	Has_downloads bool `json:"has_downloads,omitempty"`
	Git_refs_url string `json:"git_refs_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Default_branch string `json:"default_branch,omitempty"`
	Issue_comment_url string `json:"issue_comment_url"`
	Id int `json:"id"`
	Created_at string `json:"created_at,omitempty"`
	Is_template bool `json:"is_template,omitempty"`
	Subscription_url string `json:"subscription_url"`
	Commits_url string `json:"commits_url"`
	Statuses_url string `json:"statuses_url"`
	Compare_url string `json:"compare_url"`
	Contributors_url string `json:"contributors_url"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Watchers int `json:"watchers,omitempty"`
	Has_pages bool `json:"has_pages,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Rocket int `json:"rocket"`
	Url string `json:"url"`
	Hooray int `json:"hooray"`
	Confused int `json:"confused"`
	Eyes int `json:"eyes"`
	Heart int `json:"heart"`
	Field1 int `json:"-1"`
	Total_count int `json:"total_count"`
	Field1 int `json:"+1"`
	Laugh int `json:"laugh"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Custom_branch_policies bool `json:"custom_branch_policies"` // Whether only branches that match the specified name patterns can deploy to this environment. If `custom_branch_policies` is `true`, `protected_branches` must be `false`; if `custom_branch_policies` is `false`, `protected_branches` must be `true`.
	Protected_branches bool `json:"protected_branches"` // Whether only branches with branch protection rules can deploy to this environment. If `protected_branches` is `true`, `custom_branch_policies` must be `false`; if `protected_branches` is `false`, `custom_branch_policies` must be `true`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Thread map[string]interface{} `json:"thread"`
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Resolution_comment string `json:"resolution_comment,omitempty"` // An optional comment to resolve an alert.
	Secret_type string `json:"secret_type,omitempty"` // The type of secret that secret scanning detected.
	Url string `json:"url,omitempty"` // The REST API URL of the alert resource.
	Resolution string `json:"resolution,omitempty"` // **Required when the `state` is `resolved`.** The reason for resolving the alert.
	Resolved_at string `json:"resolved_at,omitempty"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Push_protection_bypassed_by GeneratedType `json:"push_protection_bypassed_by,omitempty"` // A GitHub user.
	Created_at string `json:"created_at,omitempty"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Secret string `json:"secret,omitempty"` // The secret that was detected.
	Push_protection_bypassed bool `json:"push_protection_bypassed,omitempty"` // Whether push protection was bypassed for the detected secret.
	Resolved_by GeneratedType `json:"resolved_by,omitempty"` // A GitHub user.
	State string `json:"state,omitempty"` // Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
	Number int `json:"number,omitempty"` // The security alert number.
	Locations_url string `json:"locations_url,omitempty"` // The REST API URL of the code locations for this alert.
	Secret_type_display_name string `json:"secret_type_display_name,omitempty"` // User-friendly name for the detected secret, matching the `secret_type`. For a list of built-in patterns, see "[Secret scanning patterns](https://docs.github.com/code-security/secret-scanning/secret-scanning-patterns#supported-secrets-for-advanced-security)."
	Push_protection_bypassed_at string `json:"push_protection_bypassed_at,omitempty"` // The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Html_url string `json:"html_url,omitempty"` // The GitHub URL of the alert resource.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Base_branch string `json:"base_branch,omitempty"`
	Merge_type string `json:"merge_type,omitempty"`
	Message string `json:"message,omitempty"`
}

// Manifest represents the Manifest schema from the OpenAPI specification
type Manifest struct {
	Resolved map[string]interface{} `json:"resolved,omitempty"` // A collection of resolved package dependencies.
	File map[string]interface{} `json:"file,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"` // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
	Name string `json:"name"` // The name of the manifest.
}

// Label represents the Label schema from the OpenAPI specification
type Label struct {
	Url string `json:"url"` // URL for the label
	Color string `json:"color"` // 6-character hex code, without the leading #, identifying the color
	DefaultField bool `json:"default"`
	Description string `json:"description"`
	Id int64 `json:"id"`
	Name string `json:"name"` // The name of the label.
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Repositories_added []map[string]interface{} `json:"repositories_added"` // An array of repository objects, which were added to the installation.
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Requester map[string]interface{} `json:"requester"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repositories_removed []map[string]interface{} `json:"repositories_removed"` // An array of repository objects, which were removed from the installation.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Action string `json:"action"`
	Installation Installation `json:"installation"` // Installation
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	PackageField map[string]interface{} `json:"package"` // Information about the package.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// Event represents the Event schema from the OpenAPI specification
type Event struct {
	TypeField string `json:"type"`
	Actor Actor `json:"actor"` // Actor
	Created_at string `json:"created_at"`
	Id string `json:"id"`
	Org Actor `json:"org,omitempty"` // Actor
	Payload map[string]interface{} `json:"payload"`
	Public bool `json:"public"`
	Repo map[string]interface{} `json:"repo"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the comment if the action was `edited`.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name of the secret.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
	Updated_at string `json:"updated_at"`
	Visibility string `json:"visibility"` // Visibility of a secret
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship map[string]interface{} `json:"sponsorship"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Inviter GeneratedType `json:"inviter"` // A GitHub user.
	Node_id string `json:"node_id"`
	Url string `json:"url"` // URL for the repository invitation
	Id int `json:"id"` // Unique identifier of the repository invitation.
	Html_url string `json:"html_url"`
	Created_at string `json:"created_at"`
	Permissions string `json:"permissions"` // The permission associated with the invitation.
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Expired bool `json:"expired,omitempty"` // Whether or not the invitation has expired
	Invitee GeneratedType `json:"invitee"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/reference/pulls#comments) itself.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Changes map[string]interface{} `json:"changes"` // The changes to the comment.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Temp_download_token string `json:"temp_download_token,omitempty"` // A short lived bearer token used to download the runner, if needed.
	Architecture string `json:"architecture"`
	Download_url string `json:"download_url"`
	Filename string `json:"filename"`
	Os string `json:"os"`
	Sha256_checksum string `json:"sha256_checksum,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes interface{} `json:"changes,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Starred_at interface{} `json:"starred_at"` // The time the star was created. This is a timestamp in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`. Will be `null` for the `deleted` action.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue interface{} `json:"issue"`
	Milestone map[string]interface{} `json:"milestone,omitempty"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request interface{} `json:"pull_request"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Included_gigabytes_bandwidth int `json:"included_gigabytes_bandwidth"` // Free storage space (GB) for GitHub Packages.
	Total_gigabytes_bandwidth_used int `json:"total_gigabytes_bandwidth_used"` // Sum of the free and paid storage space (GB) for GitHuub Packages.
	Total_paid_gigabytes_bandwidth_used int `json:"total_paid_gigabytes_bandwidth_used"` // Total paid storage space (GB) for GitHuub Packages.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Billable map[string]interface{} `json:"billable"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id string `json:"id,omitempty"` // An identifier for the upload.
	Url string `json:"url,omitempty"` // The REST API URL for checking the status of the upload.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project map[string]interface{} `json:"project"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Membership map[string]interface{} `json:"membership"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enabled bool `json:"enabled"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Preferences map[string]interface{} `json:"preferences"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project_column map[string]interface{} `json:"project_column"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Releases_url string `json:"releases_url"`
	Ssh_url string `json:"ssh_url"`
	Fork bool `json:"fork"`
	Watchers_count int `json:"watchers_count"`
	Starred_at string `json:"starred_at,omitempty"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Merges_url string `json:"merges_url"`
	Forks_url string `json:"forks_url"`
	Stargazers_url string `json:"stargazers_url"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Private bool `json:"private"` // Whether the repository is private or public.
	Master_branch string `json:"master_branch,omitempty"`
	Trees_url string `json:"trees_url"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Svn_url string `json:"svn_url"`
	Git_tags_url string `json:"git_tags_url"`
	Open_issues int `json:"open_issues"`
	Language string `json:"language"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Forks_count int `json:"forks_count"`
	Size int `json:"size"` // The size of the repository. Size is calculated hourly. When a repository is initially created, the size is 0.
	Deployments_url string `json:"deployments_url"`
	Url string `json:"url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Downloads_url string `json:"downloads_url"`
	Pushed_at string `json:"pushed_at"`
	Name string `json:"name"` // The name of the repository.
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Stargazers_count int `json:"stargazers_count"`
	Comments_url string `json:"comments_url"`
	Node_id string `json:"node_id"`
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Compare_url string `json:"compare_url"`
	Description string `json:"description"`
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Updated_at string `json:"updated_at"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Id int `json:"id"` // Unique identifier of the repository
	Contents_url string `json:"contents_url"`
	Commits_url string `json:"commits_url"`
	Labels_url string `json:"labels_url"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Open_issues_count int `json:"open_issues_count"`
	Teams_url string `json:"teams_url"`
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Hooks_url string `json:"hooks_url"`
	Issue_events_url string `json:"issue_events_url"`
	Contributors_url string `json:"contributors_url"`
	Notifications_url string `json:"notifications_url"`
	Full_name string `json:"full_name"`
	Watchers int `json:"watchers"`
	Template_repository map[string]interface{} `json:"template_repository,omitempty"`
	Pulls_url string `json:"pulls_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Events_url string `json:"events_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Archive_url string `json:"archive_url"`
	Forks int `json:"forks"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Topics []string `json:"topics,omitempty"`
	Issues_url string `json:"issues_url"`
	Keys_url string `json:"keys_url"`
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Collaborators_url string `json:"collaborators_url"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Network_count int `json:"network_count,omitempty"`
	Git_url string `json:"git_url"`
	Subscription_url string `json:"subscription_url"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub user.
	Html_url string `json:"html_url"`
	Subscribers_url string `json:"subscribers_url"`
	Assignees_url string `json:"assignees_url"`
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Statuses_url string `json:"statuses_url"`
	Languages_url string `json:"languages_url"`
	Clone_url string `json:"clone_url"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Blobs_url string `json:"blobs_url"`
	Has_pages bool `json:"has_pages"`
	Created_at string `json:"created_at"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property has been deprecated. Please use `squash_merge_commit_title` instead.
	License GeneratedType `json:"license"` // License Simple
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Homepage string `json:"homepage"`
	Milestones_url string `json:"milestones_url"`
	Git_refs_url string `json:"git_refs_url"`
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Mirror_url string `json:"mirror_url"`
	Git_commits_url string `json:"git_commits_url"`
	Branches_url string `json:"branches_url"`
	Tags_url string `json:"tags_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Id int `json:"id,omitempty"`
	Pattern string `json:"pattern"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Confirm_delete_url string `json:"confirm_delete_url"` // Next deletable analysis in chain, with last analysis deletion confirmation
	Next_analysis_url string `json:"next_analysis_url"` // Next deletable analysis in chain, without last analysis deletion confirmation
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit map[string]interface{} `json:"commit"`
	Content map[string]interface{} `json:"content"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Email string `json:"email"`
	Id int `json:"id"`
	Invitation_source string `json:"invitation_source,omitempty"`
	Invitation_teams_url string `json:"invitation_teams_url"`
	Created_at string `json:"created_at"`
	Role string `json:"role"`
	Failed_reason string `json:"failed_reason,omitempty"`
	Inviter GeneratedType `json:"inviter"` // A GitHub user.
	Team_count int `json:"team_count"`
	Failed_at string `json:"failed_at,omitempty"`
	Login string `json:"login"`
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	From string `json:"from"`
	To string `json:"to"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Title string `json:"title"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Project_card interface{} `json:"project_card"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2 GeneratedType `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion interface{} `json:"discussion"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Url string `json:"url"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Event string `json:"event"`
	Id int `json:"id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_id string `json:"commit_id"`
	Node_id string `json:"node_id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Allowed_actions string `json:"allowed_actions,omitempty"` // The permissions policy that controls the actions and reusable workflows that are allowed to run.
	Enabled_repositories string `json:"enabled_repositories"` // The policy that controls the repositories in the organization that are allowed to run GitHub Actions.
	Selected_actions_url string `json:"selected_actions_url,omitempty"` // The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"` // The API URL to use to get or set the selected repositories that are allowed to run GitHub Actions, when `enabled_repositories` is set to `selected`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.completed JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Login string `json:"login"`
	Marketplace_pending_change map[string]interface{} `json:"marketplace_pending_change,omitempty"`
	Marketplace_purchase map[string]interface{} `json:"marketplace_purchase"`
	Organization_billing_email string `json:"organization_billing_email,omitempty"`
	TypeField string `json:"type"`
	Url string `json:"url"`
	Email string `json:"email,omitempty"`
	Id int `json:"id"`
}

// Discussion represents the Discussion schema from the OpenAPI specification
type Discussion struct {
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Answer_chosen_by map[string]interface{} `json:"answer_chosen_by"`
	Title string `json:"title"`
	Created_at string `json:"created_at"`
	Timeline_url string `json:"timeline_url,omitempty"`
	Comments int `json:"comments"`
	Repository_url string `json:"repository_url"`
	State string `json:"state"` // The current state of the discussion. `converting` means that the discussion is being converted from an issue. `transferring` means that the discussion is being transferred from another repository.
	Answer_html_url string `json:"answer_html_url"`
	Category map[string]interface{} `json:"category"`
	Id int `json:"id"`
	Number int `json:"number"`
	Active_lock_reason string `json:"active_lock_reason"`
	Answer_chosen_at string `json:"answer_chosen_at"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	User map[string]interface{} `json:"user"`
	Body string `json:"body"`
	Locked bool `json:"locked"`
	Reactions map[string]interface{} `json:"reactions,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Assignee map[string]interface{} `json:"assignee"`
	Number int `json:"number"` // The pull request number.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Pull_request map[string]interface{} `json:"pull_request"`
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// Blob represents the Blob schema from the OpenAPI specification
type Blob struct {
	Content string `json:"content"`
	Encoding string `json:"encoding"`
	Highlighted_content string `json:"highlighted_content,omitempty"`
	Node_id string `json:"node_id"`
	Sha string `json:"sha"`
	Size int `json:"size"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Title string `json:"title"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Key string `json:"key"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Count int `json:"count"`
	Path string `json:"path"`
	Title string `json:"title"`
	Uniques int `json:"uniques"`
}

// Runner represents the Runner schema from the OpenAPI specification
type Runner struct {
	Busy bool `json:"busy"`
	Id int `json:"id"` // The id of the runner.
	Labels []GeneratedType `json:"labels"`
	Name string `json:"name"` // The name of the runner.
	Os string `json:"os"` // The Operating System of the runner.
	Status string `json:"status"` // The status of the runner.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Comment map[string]interface{} `json:"comment"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.created JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) the comment belongs to.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Changes map[string]interface{} `json:"changes"` // The changes to the comment.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/reference/issues#comments) itself.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pages []string `json:"pages,omitempty"`
	Actions []string `json:"actions,omitempty"`
	Api []string `json:"api,omitempty"`
	Dependabot []string `json:"dependabot,omitempty"`
	Importer []string `json:"importer,omitempty"`
	Git []string `json:"git,omitempty"`
	Packages []string `json:"packages,omitempty"`
	Ssh_key_fingerprints map[string]interface{} `json:"ssh_key_fingerprints,omitempty"`
	Ssh_keys []string `json:"ssh_keys,omitempty"`
	Verifiable_password_authentication bool `json:"verifiable_password_authentication"`
	Web []string `json:"web,omitempty"`
	Hooks []string `json:"hooks,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Followers int `json:"followers,omitempty"`
	Site_admin bool `json:"site_admin"`
	Score float64 `json:"score"`
	Created_at string `json:"created_at,omitempty"`
	TypeField string `json:"type"`
	Gravatar_id string `json:"gravatar_id"`
	Login string `json:"login"`
	Public_gists int `json:"public_gists,omitempty"`
	Starred_url string `json:"starred_url"`
	Following_url string `json:"following_url"`
	Gists_url string `json:"gists_url"`
	Hireable bool `json:"hireable,omitempty"`
	Bio string `json:"bio,omitempty"`
	Suspended_at string `json:"suspended_at,omitempty"`
	Id int `json:"id"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	Node_id string `json:"node_id"`
	Repos_url string `json:"repos_url"`
	Company string `json:"company,omitempty"`
	Received_events_url string `json:"received_events_url"`
	Blog string `json:"blog,omitempty"`
	Following int `json:"following,omitempty"`
	Location string `json:"location,omitempty"`
	Name string `json:"name,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Followers_url string `json:"followers_url"`
	Events_url string `json:"events_url"`
	Email string `json:"email,omitempty"`
	Public_repos int `json:"public_repos,omitempty"`
	Organizations_url string `json:"organizations_url"`
	Subscriptions_url string `json:"subscriptions_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Clones []Traffic `json:"clones"`
	Count int `json:"count"`
	Uniques int `json:"uniques"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	State string `json:"state"`
	Dismissal_commit_id string `json:"dismissal_commit_id,omitempty"`
	Dismissal_message string `json:"dismissal_message"`
	Review_id int `json:"review_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url"`
	Key string `json:"key"`
	Name string `json:"name"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Registry_package map[string]interface{} `json:"registry_package"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// Actor represents the Actor schema from the OpenAPI specification
type Actor struct {
	Url string `json:"url"`
	Avatar_url string `json:"avatar_url"`
	Display_login string `json:"display_login,omitempty"`
	Gravatar_id string `json:"gravatar_id"`
	Id int `json:"id"`
	Login string `json:"login"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Status string `json:"status"`
	Behind_by int `json:"behind_by"`
	Merge_base_commit Commit `json:"merge_base_commit"` // Commit
	Permalink_url string `json:"permalink_url"`
	Ahead_by int `json:"ahead_by"`
	Files []GeneratedType `json:"files,omitempty"`
	Patch_url string `json:"patch_url"`
	Base_commit Commit `json:"base_commit"` // Commit
	Html_url string `json:"html_url"`
	Total_commits int `json:"total_commits"`
	Url string `json:"url"`
	Commits []Commit `json:"commits"`
	Diff_url string `json:"diff_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Deleted_at string `json:"deleted_at,omitempty"`
	Description string `json:"description,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Name string `json:"name"` // The name of the package version.
	Package_html_url string `json:"package_html_url"`
	Created_at string `json:"created_at"`
	Id int `json:"id"` // Unique identifier of the package version.
	Url string `json:"url"`
	Html_url string `json:"html_url,omitempty"`
	License string `json:"license,omitempty"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Author map[string]interface{} `json:"author"`
	Committer map[string]interface{} `json:"committer"`
	Id string `json:"id"`
	Message string `json:"message"`
	Timestamp string `json:"timestamp"`
	Tree_id string `json:"tree_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Status string `json:"status,omitempty"`
	Url string `json:"url,omitempty"`
	Documentation_url string `json:"documentation_url,omitempty"`
	Message string `json:"message,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	PackageField map[string]interface{} `json:"package"` // Information about the package.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"` // URL for the issue comment
	Body_html string `json:"body_html,omitempty"`
	Body_text string `json:"body_text,omitempty"`
	User GeneratedType `json:"user"` // A GitHub user.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body,omitempty"` // Contents of the issue comment
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Id int `json:"id"` // Unique identifier of the issue comment
	Issue_url string `json:"issue_url"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Can_approve_pull_request_reviews bool `json:"can_approve_pull_request_reviews"` // Whether GitHub Actions can approve pull requests. Enabling this can be a security risk.
	Default_workflow_permissions string `json:"default_workflow_permissions"` // The default workflow permissions granted to the GITHUB_TOKEN when running workflows.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Committer map[string]interface{} `json:"committer"` // Identifying information for the git-user
	Url string `json:"url"`
	Verification map[string]interface{} `json:"verification"`
	Html_url string `json:"html_url"`
	Author map[string]interface{} `json:"author"` // Identifying information for the git-user
	Event string `json:"event,omitempty"`
	Message string `json:"message"` // Message describing the purpose of the commit
	Parents []map[string]interface{} `json:"parents"`
	Tree map[string]interface{} `json:"tree"`
	Node_id string `json:"node_id"`
	Sha string `json:"sha"` // SHA for the commit
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key string `json:"key"` // The Base64 encoded public key.
	Key_id string `json:"key_id"` // The identifier for the key.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Date string `json:"date,omitempty"`
	Email string `json:"email,omitempty"`
	Name string `json:"name,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project_card map[string]interface{} `json:"project_card"`
	Repository GeneratedType `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total_count int `json:"total_count"`
	Url string `json:"url"`
	Commit_url string `json:"commit_url"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Sha string `json:"sha"`
	State string `json:"state"`
	Statuses []GeneratedType `json:"statuses"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow map[string]interface{} `json:"workflow"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Value string `json:"value"` // The value of the variable.
	Visibility string `json:"visibility"` // Visibility of a variable
	Created_at string `json:"created_at"` // The date and time at which the variable was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the variable.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
	Updated_at string `json:"updated_at"` // The date and time at which the variable was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ref string `json:"ref"` // The [`git ref`](https://docs.github.com/rest/reference/git#get-a-reference) resource.
	Ref_type string `json:"ref_type"` // The type of Git ref object deleted in the repository.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pusher_type string `json:"pusher_type"` // The pusher type for the event. Can be either `user` or a deploy key.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_job map[string]interface{} `json:"workflow_job"`
	Action string `json:"action"`
}

// Job represents the Job schema from the OpenAPI specification
type Job struct {
	Url string `json:"url"`
	Workflow_name string `json:"workflow_name"` // The name of the workflow.
	Head_sha string `json:"head_sha"` // The SHA of the commit that is being run.
	Run_url string `json:"run_url"`
	Name string `json:"name"` // The name of the job.
	Steps []map[string]interface{} `json:"steps,omitempty"` // Steps in this job.
	Runner_name string `json:"runner_name"` // The name of the runner to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Started_at string `json:"started_at"` // The time that the job started, in ISO 8601 format.
	Check_run_url string `json:"check_run_url"`
	Runner_group_name string `json:"runner_group_name"` // The name of the runner group to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Run_id int `json:"run_id"` // The id of the associated workflow run.
	Labels []string `json:"labels"` // Labels for the workflow job. Specified by the "runs_on" attribute in the action's workflow file.
	Node_id string `json:"node_id"`
	Runner_group_id int `json:"runner_group_id"` // The ID of the runner group to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Runner_id int `json:"runner_id"` // The ID of the runner to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Status string `json:"status"` // The phase of the lifecycle that the job is currently in.
	Head_branch string `json:"head_branch"` // The name of the current branch.
	Html_url string `json:"html_url"`
	Conclusion string `json:"conclusion"` // The outcome of the job.
	Id int `json:"id"` // The id of the job.
	Run_attempt int `json:"run_attempt,omitempty"` // Attempt number of the associated workflow run, 1 for first attempt and higher if the workflow was re-run.
	Completed_at string `json:"completed_at"` // The time that the job finished, in ISO 8601 format.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alert GeneratedType `json:"alert"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	State string `json:"state"` // State of a code scanning alert.
	Repository GeneratedType `json:"repository"` // A GitHub repository.
	Dismissed_reason string `json:"dismissed_reason"` // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
	Instances_url string `json:"instances_url"` // The REST API URL for fetching the list of instances for an alert.
	Tool GeneratedType `json:"tool"`
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Most_recent_instance GeneratedType `json:"most_recent_instance"`
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Fixed_at string `json:"fixed_at,omitempty"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Number int `json:"number"` // The security alert number.
	Dismissed_comment string `json:"dismissed_comment,omitempty"` // The dismissal comment associated with the dismissal of the alert.
	Rule GeneratedType `json:"rule"`
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Score float64 `json:"score"`
	Curated bool `json:"curated"`
	Related []map[string]interface{} `json:"related,omitempty"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Logo_url string `json:"logo_url,omitempty"`
	Released string `json:"released"`
	Description string `json:"description"`
	Featured bool `json:"featured"`
	Display_name string `json:"display_name"`
	Repository_count int `json:"repository_count,omitempty"`
	Created_by string `json:"created_by"`
	Aliases []map[string]interface{} `json:"aliases,omitempty"`
	Name string `json:"name"`
	Short_description string `json:"short_description"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Requester map[string]interface{} `json:"requester,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation Installation `json:"installation"` // Installation
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Project_column map[string]interface{} `json:"project_column"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Use_default bool `json:"use_default"` // Whether to use the default template or not. If `true`, the `include_claim_keys` field is ignored.
	Include_claim_keys []string `json:"include_claim_keys,omitempty"` // Array of unique strings. Each claim key can only contain alphanumeric characters and underscores.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Blocked_user map[string]interface{} `json:"blocked_user"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Previous_attempt_url string `json:"previous_attempt_url,omitempty"` // The URL to the previous attempted run of this workflow, if one exists.
	Jobs_url string `json:"jobs_url"` // The URL to the jobs for the workflow run.
	Head_commit GeneratedType `json:"head_commit"` // A commit.
	Status string `json:"status"`
	Head_repository_id int `json:"head_repository_id,omitempty"`
	Created_at string `json:"created_at"`
	Head_repository GeneratedType `json:"head_repository"` // Minimal Repository
	Name string `json:"name,omitempty"` // The name of the workflow run.
	Conclusion string `json:"conclusion"`
	Node_id string `json:"node_id"`
	Cancel_url string `json:"cancel_url"` // The URL to cancel the workflow run.
	Check_suite_node_id string `json:"check_suite_node_id,omitempty"` // The node ID of the associated check suite.
	Display_title string `json:"display_title"` // The event-specific title associated with the run or the run-name if set, or the value of `run-name` if it is set in the workflow.
	Workflow_id int `json:"workflow_id"` // The ID of the parent workflow.
	Pull_requests []GeneratedType `json:"pull_requests"`
	Check_suite_id int `json:"check_suite_id,omitempty"` // The ID of the associated check suite.
	Logs_url string `json:"logs_url"` // The URL to download the logs for the workflow run.
	Run_started_at string `json:"run_started_at,omitempty"` // The start time of the latest run. Resets on re-run.
	Url string `json:"url"` // The URL to the workflow run.
	Rerun_url string `json:"rerun_url"` // The URL to rerun the workflow run.
	Referenced_workflows []GeneratedType `json:"referenced_workflows,omitempty"`
	Updated_at string `json:"updated_at"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Artifacts_url string `json:"artifacts_url"` // The URL to the artifacts for the workflow run.
	Head_sha string `json:"head_sha"` // The SHA of the head commit that points to the version of the workflow being run.
	Check_suite_url string `json:"check_suite_url"` // The URL to the associated check suite.
	Workflow_url string `json:"workflow_url"` // The URL to the workflow.
	Head_branch string `json:"head_branch"`
	Event string `json:"event"`
	Html_url string `json:"html_url"`
	Actor GeneratedType `json:"actor,omitempty"` // A GitHub user.
	Id int `json:"id"` // The ID of the workflow run.
	Path string `json:"path"` // The full path of the workflow
	Run_number int `json:"run_number"` // The auto incrementing run number for the workflow run.
	Run_attempt int `json:"run_attempt,omitempty"` // Attempt number of the run, 1 for first attempt and higher if the workflow was re-run.
	Triggering_actor GeneratedType `json:"triggering_actor,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Key string `json:"key"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_job interface{} `json:"workflow_job"`
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/reference/pulls#comments) itself.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Key string `json:"key"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Minutes_used_breakdown map[string]interface{} `json:"minutes_used_breakdown"`
	Total_minutes_used int `json:"total_minutes_used"` // The sum of the free and paid GitHub Actions minutes used.
	Total_paid_minutes_used int `json:"total_paid_minutes_used"` // The total paid GitHub Actions minutes used.
	Included_minutes int `json:"included_minutes"` // The amount of free GitHub Actions minutes available.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Strict bool `json:"strict,omitempty"`
	Url string `json:"url,omitempty"`
	Checks []map[string]interface{} `json:"checks"`
	Contexts []string `json:"contexts"`
	Contexts_url string `json:"contexts_url,omitempty"`
	Enforcement_level string `json:"enforcement_level,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue_comment_url string `json:"issue_comment_url"` // The API URL to get the issue comment where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Build map[string]interface{} `json:"build"` // The [List GitHub Pages builds](https://docs.github.com/rest/reference/repos#list-github-pages-builds) itself.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Id int `json:"id"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Secret_scanning map[string]interface{} `json:"secret_scanning,omitempty"`
	Secret_scanning_push_protection map[string]interface{} `json:"secret_scanning_push_protection,omitempty"`
	Advanced_security map[string]interface{} `json:"advanced_security,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Pull_request map[string]interface{} `json:"pull_request"`
	Review map[string]interface{} `json:"review"` // The review that was affected.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Project_url string `json:"project_url"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Cards_url string `json:"cards_url"`
	Created_at string `json:"created_at"`
	Id int `json:"id"` // The unique identifier of the project column
	Name string `json:"name"` // Name of the project column
}

// Language represents the Language schema from the OpenAPI specification
type Language struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Requested_action map[string]interface{} `json:"requested_action,omitempty"` // The action requested by the user.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Check_run GeneratedType `json:"check_run"` // A check performed on the code of a given code change
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Team map[string]interface{} `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"` // The global node ID of the installation.
	Id int `json:"id"` // The ID of the installation.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	First_patched_version map[string]interface{} `json:"first_patched_version"` // Details pertaining to the package version that patches this vulnerability.
	PackageField GeneratedType `json:"package"` // Details for the vulnerable package.
	Severity string `json:"severity"` // The severity of the vulnerability.
	Vulnerable_version_range string `json:"vulnerable_version_range"` // Conditions that identify vulnerable versions of this vulnerability's package.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Warning string `json:"warning"` // Warning generated when processing the analysis
	Results_count int `json:"results_count"` // The total number of results in the analysis.
	Id int `json:"id"` // Unique identifier for this analysis.
	Created_at string `json:"created_at"` // The time that the analysis was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Commit_sha string `json:"commit_sha"` // The SHA of the commit to which the analysis you are uploading relates.
	ErrorField string `json:"error"`
	Deletable bool `json:"deletable"`
	Tool GeneratedType `json:"tool"`
	Environment string `json:"environment"` // Identifies the variable values associated with the environment in which this analysis was performed.
	Analysis_key string `json:"analysis_key"` // Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name.
	Url string `json:"url"` // The REST API URL of the analysis resource.
	Ref string `json:"ref"` // The full Git reference, formatted as `refs/heads/<branch name>`, `refs/pull/<number>/merge`, or `refs/pull/<number>/head`.
	Rules_count int `json:"rules_count"` // The total number of rules used in the analysis.
	Sarif_id string `json:"sarif_id"` // An identifier for the upload.
	Category string `json:"category,omitempty"` // Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Pull_request map[string]interface{} `json:"pull_request"`
	Action string `json:"action"`
	Before string `json:"before"`
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	After string `json:"after"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit map[string]interface{} `json:"commit"`
	Name string `json:"name"`
	Protected bool `json:"protected"`
	Protection GeneratedType `json:"protection,omitempty"` // Branch Protection
	Protection_url string `json:"protection_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Source map[string]interface{} `json:"source"`
	Updated_at string `json:"updated_at"`
	Actor GeneratedType `json:"actor,omitempty"` // A GitHub user.
	Created_at string `json:"created_at"`
	Event string `json:"event"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Original_commit_id string `json:"original_commit_id"`
	Body string `json:"body"`
	Original_position int `json:"original_position"`
	Original_start_line int `json:"original_start_line,omitempty"` // The original first line of the range for a multi-line comment.
	Pull_request_review_id int `json:"pull_request_review_id"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Updated_at string `json:"updated_at"`
	Body_html string `json:"body_html,omitempty"`
	Diff_hunk string `json:"diff_hunk"`
	In_reply_to_id int `json:"in_reply_to_id,omitempty"`
	Body_text string `json:"body_text,omitempty"`
	Pull_request_url string `json:"pull_request_url"`
	Created_at string `json:"created_at"`
	Path string `json:"path"`
	Position int `json:"position"`
	Start_side string `json:"start_side,omitempty"` // The side of the first line of the range for a multi-line comment.
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
	Original_line int `json:"original_line,omitempty"` // The original line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Side string `json:"side,omitempty"` // The side of the first line of the range for a multi-line comment.
	Links map[string]interface{} `json:"_links"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Line int `json:"line,omitempty"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Reactions GeneratedType `json:"reactions,omitempty"`
	Start_line int `json:"start_line,omitempty"` // The first line of the range for a multi-line comment.
	User GeneratedType `json:"user"` // A GitHub user.
	Node_id string `json:"node_id"`
}

// Workflow represents the Workflow schema from the OpenAPI specification
type Workflow struct {
	Updated_at string `json:"updated_at"`
	Deleted_at string `json:"deleted_at,omitempty"`
	Id int `json:"id"`
	Name string `json:"name"`
	Url string `json:"url"`
	Badge_url string `json:"badge_url"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Created_at string `json:"created_at"`
	Path string `json:"path"`
	State string `json:"state"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Required_approving_review_count int `json:"required_approving_review_count,omitempty"`
	Url string `json:"url,omitempty"`
	Bypass_pull_request_allowances map[string]interface{} `json:"bypass_pull_request_allowances,omitempty"` // Allow specific users, teams, or apps to bypass pull request requirements.
	Dismiss_stale_reviews bool `json:"dismiss_stale_reviews"`
	Dismissal_restrictions map[string]interface{} `json:"dismissal_restrictions,omitempty"`
	Require_code_owner_reviews bool `json:"require_code_owner_reviews"`
	Require_last_push_approval bool `json:"require_last_push_approval,omitempty"` // Whether the most recent push must be approved by someone other than the person who pushed it.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Content_reports_enabled bool `json:"content_reports_enabled,omitempty"`
	Description string `json:"description"`
	Documentation string `json:"documentation"`
	Files map[string]interface{} `json:"files"`
	Health_percentage int `json:"health_percentage"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Release interface{} `json:"release"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Team map[string]interface{} `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the team if the action was `edited`.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Estimated_paid_storage_for_month int `json:"estimated_paid_storage_for_month"` // Estimated storage space (GB) used in billing cycle.
	Estimated_storage_for_month int `json:"estimated_storage_for_month"` // Estimated sum of free and paid storage space (GB) used in billing cycle.
	Days_left_in_billing_cycle int `json:"days_left_in_billing_cycle"` // Numbers of days left in billing cycle.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Requested_reviewer GeneratedType `json:"requested_reviewer,omitempty"` // A GitHub user.
	Requested_team Team `json:"requested_team,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Issue GeneratedType `json:"issue,omitempty"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Commit_url string `json:"commit_url"`
	Label GeneratedType `json:"label,omitempty"` // Issue Event Label
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Dismissed_review GeneratedType `json:"dismissed_review,omitempty"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Assignee GeneratedType `json:"assignee,omitempty"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Id int `json:"id"`
	Lock_reason string `json:"lock_reason,omitempty"`
	Created_at string `json:"created_at"`
	Author_association string `json:"author_association,omitempty"` // How the author is associated with the repository.
	Project_card GeneratedType `json:"project_card,omitempty"` // Issue Event Project Card
	Rename GeneratedType `json:"rename,omitempty"` // Issue Event Rename
	Review_requester GeneratedType `json:"review_requester,omitempty"` // A GitHub user.
	Assigner GeneratedType `json:"assigner,omitempty"` // A GitHub user.
	Event string `json:"event"`
	Milestone GeneratedType `json:"milestone,omitempty"` // Issue Event Milestone
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Branch string `json:"branch"`
	Client_payload map[string]interface{} `json:"client_payload"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Fixed_at string `json:"fixed_at,omitempty"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Most_recent_instance GeneratedType `json:"most_recent_instance"`
	Dismissed_comment string `json:"dismissed_comment,omitempty"` // The dismissal comment associated with the dismissal of the alert.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Tool GeneratedType `json:"tool"`
	Dismissed_reason string `json:"dismissed_reason"` // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
	Instances_url string `json:"instances_url"` // The REST API URL for fetching the list of instances for an alert.
	Number int `json:"number"` // The security alert number.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Rule GeneratedType `json:"rule"`
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	State string `json:"state"` // State of a code scanning alert.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Assignee map[string]interface{} `json:"assignee,omitempty"`
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Action string `json:"action"`
	Number int `json:"number"` // The pull request number.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url,omitempty"`
	Key string `json:"key"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Spdx_id string `json:"spdx_id"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Rule map[string]interface{} `json:"rule"` // The branch protection rule. Includes a `name` and all the [branch protection settings](https://docs.github.com/github/administering-a-repository/defining-the-mergeability-of-pull-requests/about-protected-branches#about-branch-protection-settings) applied to branches that match the name. Binary settings are boolean. Multi-level configurations are one of `off`, `non_admins`, or `everyone`. Actor and build lists are arrays of strings.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// Contributor represents the Contributor schema from the OpenAPI specification
type Contributor struct {
	Avatar_url string `json:"avatar_url,omitempty"`
	Id int `json:"id,omitempty"`
	Organizations_url string `json:"organizations_url,omitempty"`
	Site_admin bool `json:"site_admin,omitempty"`
	Gravatar_id string `json:"gravatar_id,omitempty"`
	Name string `json:"name,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Events_url string `json:"events_url,omitempty"`
	Following_url string `json:"following_url,omitempty"`
	Followers_url string `json:"followers_url,omitempty"`
	Received_events_url string `json:"received_events_url,omitempty"`
	Contributions int `json:"contributions"`
	Subscriptions_url string `json:"subscriptions_url,omitempty"`
	TypeField string `json:"type"`
	Gists_url string `json:"gists_url,omitempty"`
	Repos_url string `json:"repos_url,omitempty"`
	Login string `json:"login,omitempty"`
	Starred_url string `json:"starred_url,omitempty"`
	Url string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"` // The type of activity for the event that triggered the delivery.
	Duration float64 `json:"duration"` // Time spent delivering.
	Repository_id int `json:"repository_id"` // The id of the repository associated with this event.
	Id int `json:"id"` // Unique identifier of the webhook delivery.
	Status string `json:"status"` // Describes the response returned after attempting the delivery.
	Delivered_at string `json:"delivered_at"` // Time when the webhook delivery occurred.
	Installation_id int `json:"installation_id"` // The id of the GitHub App installation associated with this event.
	Event string `json:"event"` // The event that triggered the delivery.
	Guid string `json:"guid"` // Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event).
	Redelivery bool `json:"redelivery"` // Whether the webhook delivery is a redelivery.
	Status_code int `json:"status_code"` // Status code received when delivery was made.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Permission string `json:"permission"`
	Role_name string `json:"role_name"`
	User GeneratedType `json:"user"` // Collaborator
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Release map[string]interface{} `json:"release"` // The [release](https://docs.github.com/rest/reference/repos/#get-a-release) object.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	All []int `json:"all"`
	Owner []int `json:"owner"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Member map[string]interface{} `json:"member"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// Metadata represents the Metadata schema from the OpenAPI specification
type Metadata struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Release map[string]interface{} `json:"release"` // The [release](https://docs.github.com/rest/reference/repos/#get-a-release) object.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Workflow string `json:"workflow"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Inputs map[string]interface{} `json:"inputs"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Ref string `json:"ref"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes int `json:"changes"`
	Patch string `json:"patch,omitempty"`
	Sha string `json:"sha"`
	Additions int `json:"additions"`
	Previous_filename string `json:"previous_filename,omitempty"`
	Raw_url string `json:"raw_url"`
	Status string `json:"status"`
	Blob_url string `json:"blob_url"`
	Contents_url string `json:"contents_url"`
	Deletions int `json:"deletions"`
	Filename string `json:"filename"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Analyses_url string `json:"analyses_url,omitempty"` // The REST API URL for getting the analyses associated with the upload.
	Errors []string `json:"errors,omitempty"` // Any errors that ocurred during processing of the delivery.
	Processing_status string `json:"processing_status,omitempty"` // `pending` files have not yet been processed, while `complete` means results from the SARIF have been stored. `failed` files have either not been processed at all, or could only be partially processed.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"`
	Source string `json:"source"`
}

// Email represents the Email schema from the OpenAPI specification
type Email struct {
	Email string `json:"email"`
	Primary bool `json:"primary"`
	Verified bool `json:"verified"`
	Visibility string `json:"visibility"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Locked bool `json:"locked"`
	Score float64 `json:"score"`
	Body_html string `json:"body_html,omitempty"`
	State_reason string `json:"state_reason,omitempty"`
	Created_at string `json:"created_at"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Updated_at string `json:"updated_at"`
	Body string `json:"body,omitempty"`
	Closed_at string `json:"closed_at"`
	Html_url string `json:"html_url"`
	Comments int `json:"comments"`
	Repository_url string `json:"repository_url"`
	Body_text string `json:"body_text,omitempty"`
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Id int `json:"id"`
	Timeline_url string `json:"timeline_url,omitempty"`
	Events_url string `json:"events_url"`
	Number int `json:"number"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Comments_url string `json:"comments_url"`
	Url string `json:"url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Labels []map[string]interface{} `json:"labels"`
	State string `json:"state"`
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Node_id string `json:"node_id"`
	Labels_url string `json:"labels_url"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Title string `json:"title"`
	Draft bool `json:"draft,omitempty"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Checks []map[string]interface{} `json:"checks"`
	Contexts []string `json:"contexts"`
	Contexts_url string `json:"contexts_url"`
	Strict bool `json:"strict"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Hooks_url string `json:"hooks_url"` // The API URL to list the hooks on the repository.
	Issue_events_url string `json:"issue_events_url"` // A template for the API URL to get information about issue events on the repository.
	Issues_url string `json:"issues_url"` // A template for the API URL to get information about issues on the repository.
	Git_commits_url string `json:"git_commits_url"` // A template for the API URL to get information about Git commits of the repository.
	Milestones_url string `json:"milestones_url"` // A template for the API URL to get information about milestones of the repository.
	Git_refs_url string `json:"git_refs_url"` // A template for the API URL to get information about Git refs of the repository.
	Commits_url string `json:"commits_url"` // A template for the API URL to get information about commits on the repository.
	Comments_url string `json:"comments_url"` // A template for the API URL to get information about comments on the repository.
	Contents_url string `json:"contents_url"` // A template for the API URL to get the contents of the repository.
	Teams_url string `json:"teams_url"` // The API URL to list the teams on the repository.
	Pulls_url string `json:"pulls_url"` // A template for the API URL to get information about pull requests on the repository.
	Languages_url string `json:"languages_url"` // The API URL to get information about the languages of the repository.
	Assignees_url string `json:"assignees_url"` // A template for the API URL to list the available assignees for issues in the repository.
	Merges_url string `json:"merges_url"` // The API URL to merge branches in the repository.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Name string `json:"name"` // The name of the repository.
	Tags_url string `json:"tags_url"` // The API URL to get information about tags on the repository.
	Compare_url string `json:"compare_url"` // A template for the API URL to compare two commits or refs.
	Labels_url string `json:"labels_url"` // A template for the API URL to get information about labels of the repository.
	Git_tags_url string `json:"git_tags_url"` // A template for the API URL to get information about Git tags of the repository.
	Statuses_url string `json:"statuses_url"` // A template for the API URL to get information about statuses of a commit.
	Notifications_url string `json:"notifications_url"` // A template for the API URL to get information about notifications on the repository.
	Subscription_url string `json:"subscription_url"` // The API URL to subscribe to notifications for this repository.
	Subscribers_url string `json:"subscribers_url"` // The API URL to list the subscribers on the repository.
	Full_name string `json:"full_name"` // The full, globally unique, name of the repository.
	Url string `json:"url"` // The URL to get more information about the repository from the GitHub API.
	Blobs_url string `json:"blobs_url"` // A template for the API URL to create or retrieve a raw Git blob in the repository.
	Fork bool `json:"fork"` // Whether the repository is a fork.
	Deployments_url string `json:"deployments_url"` // The API URL to list the deployments of the repository.
	Issue_comment_url string `json:"issue_comment_url"` // A template for the API URL to get information about issue comments on the repository.
	Stargazers_url string `json:"stargazers_url"` // The API URL to list the stargazers on the repository.
	Id int `json:"id"` // A unique identifier of the repository.
	Private bool `json:"private"` // Whether the repository is private.
	Forks_url string `json:"forks_url"` // The API URL to list the forks of the repository.
	Collaborators_url string `json:"collaborators_url"` // A template for the API URL to get information about collaborators of the repository.
	Downloads_url string `json:"downloads_url"` // The API URL to list the downloads on the repository.
	Description string `json:"description"` // The repository description.
	Keys_url string `json:"keys_url"` // A template for the API URL to get information about deploy keys on the repository.
	Releases_url string `json:"releases_url"` // A template for the API URL to get information about releases on the repository.
	Html_url string `json:"html_url"` // The URL to view the repository on GitHub.com.
	Trees_url string `json:"trees_url"` // A template for the API URL to create or retrieve a raw Git tree of the repository.
	Archive_url string `json:"archive_url"` // A template for the API URL to download the repository as an archive.
	Contributors_url string `json:"contributors_url"` // A template for the API URL to list the contributors to the repository.
	Branches_url string `json:"branches_url"` // A template for the API URL to get information about branches in the repository.
	Node_id string `json:"node_id"` // The GraphQL identifier of the repository.
	Events_url string `json:"events_url"` // The API URL to list the events of the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project_card map[string]interface{} `json:"project_card"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Page represents the Page schema from the OpenAPI specification
type Page struct {
	Protected_domain_state string `json:"protected_domain_state,omitempty"` // The state if the domain is verified
	Custom_404 bool `json:"custom_404"` // Whether the Page has a custom 404 page.
	Https_certificate GeneratedType `json:"https_certificate,omitempty"`
	Https_enforced bool `json:"https_enforced,omitempty"` // Whether https is enabled on the domain
	Public bool `json:"public"` // Whether the GitHub Pages site is publicly visible. If set to `true`, the site is accessible to anyone on the internet. If set to `false`, the site will only be accessible to users who have at least `read` access to the repository that published the site.
	Build_type string `json:"build_type,omitempty"` // The process in which the Page will be built.
	Cname string `json:"cname"` // The Pages site's custom domain
	Source GeneratedType `json:"source,omitempty"`
	Url string `json:"url"` // The API address for accessing this Page resource.
	Status string `json:"status"` // The status of the most recent build of the Page.
	Html_url string `json:"html_url,omitempty"` // The web address the Page can be accessed from.
	Pending_domain_unverified_at string `json:"pending_domain_unverified_at,omitempty"` // The timestamp when a pending domain becomes unverified.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Members_can_create_public_pages bool `json:"members_can_create_public_pages,omitempty"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Repos_url string `json:"repos_url"`
	Collaborators int `json:"collaborators,omitempty"`
	Name string `json:"name,omitempty"`
	Html_url string `json:"html_url"`
	Company string `json:"company,omitempty"`
	Public_members_url string `json:"public_members_url"`
	Events_url string `json:"events_url"`
	Login string `json:"login"`
	Members_url string `json:"members_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Blog string `json:"blog,omitempty"`
	Email string `json:"email,omitempty"`
	Public_repos int `json:"public_repos"`
	Location string `json:"location,omitempty"`
	Updated_at string `json:"updated_at"`
	Following int `json:"following"`
	Avatar_url string `json:"avatar_url"`
	Members_can_create_private_repositories bool `json:"members_can_create_private_repositories,omitempty"`
	Members_can_create_public_repositories bool `json:"members_can_create_public_repositories,omitempty"`
	Node_id string `json:"node_id"`
	TypeField string `json:"type"`
	Public_gists int `json:"public_gists"`
	Id int `json:"id"`
	Description string `json:"description"`
	Disk_usage int `json:"disk_usage,omitempty"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Members_can_fork_private_repositories bool `json:"members_can_fork_private_repositories,omitempty"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Owned_private_repos int `json:"owned_private_repos,omitempty"`
	Is_verified bool `json:"is_verified,omitempty"`
	Followers int `json:"followers"`
	Hooks_url string `json:"hooks_url"`
	Has_organization_projects bool `json:"has_organization_projects"`
	Members_can_create_private_pages bool `json:"members_can_create_private_pages,omitempty"`
	Issues_url string `json:"issues_url"`
	Total_private_repos int `json:"total_private_repos,omitempty"`
	Two_factor_requirement_enabled bool `json:"two_factor_requirement_enabled,omitempty"`
	Members_can_create_repositories bool `json:"members_can_create_repositories,omitempty"`
	Default_repository_permission string `json:"default_repository_permission,omitempty"`
	Has_repository_projects bool `json:"has_repository_projects"`
	Billing_email string `json:"billing_email,omitempty"`
	Members_allowed_repository_creation_type string `json:"members_allowed_repository_creation_type,omitempty"`
	Members_can_create_pages bool `json:"members_can_create_pages,omitempty"`
	Private_gists int `json:"private_gists,omitempty"`
	Members_can_create_internal_repositories bool `json:"members_can_create_internal_repositories,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Message string `json:"message"`
	Documentation_url string `json:"documentation_url"`
	Errors []string `json:"errors,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Workflow map[string]interface{} `json:"workflow,omitempty"`
	Workflow_run map[string]interface{} `json:"workflow_run,omitempty"`
	Check_run map[string]interface{} `json:"check_run,omitempty"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Deployment map[string]interface{} `json:"deployment"` // The [deployment](https://docs.github.com/rest/reference/deployments#list-deployments).
	Deployment_status map[string]interface{} `json:"deployment_status"` // The [deployment status](https://docs.github.com/rest/reference/deployments#list-deployment-statuses).
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Role string `json:"role"` // The role of the user in the team.
	State string `json:"state"` // The state of the user's membership in the team.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sha string `json:"sha"`
	Tag string `json:"tag"` // Name of the tag
	Tagger map[string]interface{} `json:"tagger"`
	Url string `json:"url"` // URL for the tag
	Verification Verification `json:"verification,omitempty"`
	Message string `json:"message"` // Message describing the purpose of the tag
	Node_id string `json:"node_id"`
	Object map[string]interface{} `json:"object"`
}

// Commit represents the Commit schema from the OpenAPI specification
type Commit struct {
	Comments_url string `json:"comments_url"`
	Sha string `json:"sha"`
	Commit map[string]interface{} `json:"commit"`
	Html_url string `json:"html_url"`
	Parents []map[string]interface{} `json:"parents"`
	Committer GeneratedType `json:"committer"` // A GitHub user.
	Stats map[string]interface{} `json:"stats,omitempty"`
	Url string `json:"url"`
	Files []GeneratedType `json:"files,omitempty"`
	Node_id string `json:"node_id"`
	Author GeneratedType `json:"author"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Environments []map[string]interface{} `json:"environments"` // The list of environments that were approved or rejected
	State string `json:"state"` // Whether deployment to the environment(s) was approved or rejected or pending (with comments)
	User GeneratedType `json:"user"` // A GitHub user.
	Comment string `json:"comment"` // The comment submitted with the deployment review
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Team map[string]interface{} `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Committer map[string]interface{} `json:"committer"`
	Id string `json:"id"`
	Message string `json:"message"`
	Timestamp string `json:"timestamp"`
	Tree_id string `json:"tree_id"`
	Author map[string]interface{} `json:"author"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Column_name string `json:"column_name"`
	Id int `json:"id"`
	Previous_column_name string `json:"previous_column_name,omitempty"`
	Project_id int `json:"project_id"`
	Project_url string `json:"project_url"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Label map[string]interface{} `json:"label,omitempty"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Member map[string]interface{} `json:"member"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert interface{} `json:"alert"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Body string `json:"body"` // The generated body describing the contents of the release supporting markdown formatting
	Name string `json:"name"` // The generated name of the release
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	State string `json:"state"` // State of the required workflow
	Ref string `json:"ref"` // Ref at which the workflow file will be selected
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Scope string `json:"scope"` // Scope of the required workflow
	Created_at string `json:"created_at"`
	Id float64 `json:"id"` // Unique identifier for a required workflow
	Path string `json:"path"` // Path of the workflow file
	Name string `json:"name"` // Name present in the workflow file
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the milestone if the action was `edited`.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Locked bool `json:"locked"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Created_at string `json:"created_at"`
	Repository_url string `json:"repository_url"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	User GeneratedType `json:"user"` // A GitHub user.
	Updated_at string `json:"updated_at"`
	State string `json:"state"` // State of the issue; either 'open' or 'closed'
	State_reason string `json:"state_reason,omitempty"` // The reason for the current state
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Closed_at string `json:"closed_at"`
	Number int `json:"number"` // Number uniquely identifying the issue within its repository
	Id int `json:"id"`
	Body string `json:"body,omitempty"` // Contents of the issue
	Closed_by GeneratedType `json:"closed_by,omitempty"` // A GitHub user.
	Timeline_url string `json:"timeline_url,omitempty"`
	Body_html string `json:"body_html,omitempty"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Comments int `json:"comments"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body_text string `json:"body_text,omitempty"`
	Comments_url string `json:"comments_url"`
	Events_url string `json:"events_url"`
	Title string `json:"title"` // Title of the issue
	Html_url string `json:"html_url"`
	Labels_url string `json:"labels_url"`
	Draft bool `json:"draft,omitempty"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Url string `json:"url"` // URL for the issue
	Node_id string `json:"node_id"`
	Labels []interface{} `json:"labels"` // Labels to associate with this issue; pass one or more label names to replace the set of labels on this issue; send an empty array to clear all labels from the issue; note that the labels are silently dropped for users without push access to the repository
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Category string `json:"category,omitempty"` // Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code.
	Analysis_key string `json:"analysis_key,omitempty"` // Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name.
	Html_url string `json:"html_url,omitempty"`
	Message map[string]interface{} `json:"message,omitempty"`
	Commit_sha string `json:"commit_sha,omitempty"`
	State string `json:"state,omitempty"` // State of a code scanning alert.
	Classifications []string `json:"classifications,omitempty"` // Classifications that have been applied to the file that triggered the alert. For example identifying it as documentation, or a generated file.
	Environment string `json:"environment,omitempty"` // Identifies the variable values associated with the environment in which the analysis that generated this alert instance was performed, such as the language that was analyzed.
	Location GeneratedType `json:"location,omitempty"` // Describe a region within a file for the alert.
	Ref string `json:"ref,omitempty"` // The full Git reference, formatted as `refs/heads/<branch name>`, `refs/pull/<number>/merge`, or `refs/pull/<number>/head`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Views []Traffic `json:"views"`
	Count int `json:"count"`
	Uniques int `json:"uniques"`
}

// Stargazer represents the Stargazer schema from the OpenAPI specification
type Stargazer struct {
	User GeneratedType `json:"user"` // A GitHub user.
	Starred_at string `json:"starred_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue interface{} `json:"issue"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Team map[string]interface{} `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the secret_scanning_alert_location.created JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert interface{} `json:"alert"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Answer map[string]interface{} `json:"answer"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Status string `json:"status"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issues_url string `json:"issues_url"`
	Updated_at string `json:"updated_at"`
	Secret_scanning_push_protection_enabled_for_new_repositories bool `json:"secret_scanning_push_protection_enabled_for_new_repositories,omitempty"` // Whether secret scanning push protection is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Has_organization_projects bool `json:"has_organization_projects"`
	Billing_email string `json:"billing_email,omitempty"`
	Has_repository_projects bool `json:"has_repository_projects"`
	Created_at string `json:"created_at"`
	Email string `json:"email,omitempty"`
	Id int `json:"id"`
	Members_can_create_public_repositories bool `json:"members_can_create_public_repositories,omitempty"`
	Dependency_graph_enabled_for_new_repositories bool `json:"dependency_graph_enabled_for_new_repositories,omitempty"` // Whether dependency graph is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Dependabot_alerts_enabled_for_new_repositories bool `json:"dependabot_alerts_enabled_for_new_repositories,omitempty"` // Whether GitHub Advanced Security is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Node_id string `json:"node_id"`
	TypeField string `json:"type"`
	Default_repository_permission string `json:"default_repository_permission,omitempty"`
	Members_can_create_public_pages bool `json:"members_can_create_public_pages,omitempty"`
	Members_can_create_pages bool `json:"members_can_create_pages,omitempty"`
	Html_url string `json:"html_url"`
	Members_allowed_repository_creation_type string `json:"members_allowed_repository_creation_type,omitempty"`
	Location string `json:"location,omitempty"`
	Members_can_create_repositories bool `json:"members_can_create_repositories,omitempty"`
	Following int `json:"following"`
	Hooks_url string `json:"hooks_url"`
	Name string `json:"name,omitempty"`
	Login string `json:"login"`
	Members_url string `json:"members_url"`
	Events_url string `json:"events_url"`
	Blog string `json:"blog,omitempty"`
	Members_can_fork_private_repositories bool `json:"members_can_fork_private_repositories,omitempty"`
	Public_repos int `json:"public_repos"`
	Owned_private_repos int `json:"owned_private_repos,omitempty"`
	Url string `json:"url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Description string `json:"description"`
	Two_factor_requirement_enabled bool `json:"two_factor_requirement_enabled,omitempty"`
	Is_verified bool `json:"is_verified,omitempty"`
	Advanced_security_enabled_for_new_repositories bool `json:"advanced_security_enabled_for_new_repositories,omitempty"` // Whether GitHub Advanced Security is enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Secret_scanning_enabled_for_new_repositories bool `json:"secret_scanning_enabled_for_new_repositories,omitempty"` // Whether secret scanning is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Members_can_create_private_pages bool `json:"members_can_create_private_pages,omitempty"`
	Repos_url string `json:"repos_url"`
	Company string `json:"company,omitempty"`
	Disk_usage int `json:"disk_usage,omitempty"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Secret_scanning_push_protection_custom_link_enabled bool `json:"secret_scanning_push_protection_custom_link_enabled,omitempty"` // Whether a custom link is shown to contributors who are blocked from pushing a secret by push protection.
	Twitter_username string `json:"twitter_username,omitempty"`
	Members_can_create_internal_repositories bool `json:"members_can_create_internal_repositories,omitempty"`
	Followers int `json:"followers"`
	Avatar_url string `json:"avatar_url"`
	Members_can_create_private_repositories bool `json:"members_can_create_private_repositories,omitempty"`
	Dependabot_security_updates_enabled_for_new_repositories bool `json:"dependabot_security_updates_enabled_for_new_repositories,omitempty"` // Whether dependabot security updates are automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Public_gists int `json:"public_gists"`
	Secret_scanning_push_protection_custom_link string `json:"secret_scanning_push_protection_custom_link,omitempty"` // An optional URL string to display to contributors who are blocked from pushing a secret.
	Total_private_repos int `json:"total_private_repos,omitempty"`
	Collaborators int `json:"collaborators,omitempty"`
	Private_gists int `json:"private_gists,omitempty"`
	Public_members_url string `json:"public_members_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Workflow_run interface{} `json:"workflow_run"`
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow map[string]interface{} `json:"workflow"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Url string `json:"url"`
	Commit_url string `json:"commit_url"`
	Id int `json:"id"`
	Label map[string]interface{} `json:"label"`
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repositories_url string `json:"repositories_url"`
	Slug string `json:"slug"`
	Ldap_dn string `json:"ldap_dn,omitempty"` // Distinguished Name (DN) that team maps to within LDAP environment
	Html_url string `json:"html_url"`
	Id int `json:"id"` // Unique identifier of the team
	Name string `json:"name"` // Name of the team
	Url string `json:"url"` // URL for the team
	Description string `json:"description"` // Description of the team
	Members_url string `json:"members_url"`
	Node_id string `json:"node_id"`
	Permission string `json:"permission"` // Permission that the team will have for its repositories
	Privacy string `json:"privacy,omitempty"` // The level of privacy this team should have
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Invitation map[string]interface{} `json:"invitation"` // The invitation for the user or email if the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	User map[string]interface{} `json:"user,omitempty"`
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Last_edited_at string `json:"last_edited_at"`
	Node_id string `json:"node_id"`
	Number int `json:"number"` // The unique sequence number of a team discussion comment.
	Body string `json:"body"` // The main text of the comment.
	Body_html string `json:"body_html"`
	Body_version string `json:"body_version"` // The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server.
	Created_at string `json:"created_at"`
	Discussion_url string `json:"discussion_url"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Author GeneratedType `json:"author"` // A GitHub user.
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Account GeneratedType `json:"account"`
	Billing_cycle string `json:"billing_cycle"`
	Free_trial_ends_on string `json:"free_trial_ends_on"`
	Next_billing_date string `json:"next_billing_date"`
	On_free_trial bool `json:"on_free_trial"`
	Plan GeneratedType `json:"plan"` // Marketplace Listing Plan
	Unit_count int `json:"unit_count"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Required_linear_history map[string]interface{} `json:"required_linear_history,omitempty"`
	Allow_deletions map[string]interface{} `json:"allow_deletions,omitempty"`
	Allow_fork_syncing map[string]interface{} `json:"allow_fork_syncing,omitempty"` // Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing.
	Lock_branch map[string]interface{} `json:"lock_branch,omitempty"` // Whether to set the branch as read-only. If this is true, users will not be able to push to the branch.
	Name string `json:"name,omitempty"`
	Required_conversation_resolution map[string]interface{} `json:"required_conversation_resolution,omitempty"`
	Required_signatures map[string]interface{} `json:"required_signatures,omitempty"`
	Required_status_checks GeneratedType `json:"required_status_checks,omitempty"` // Protected Branch Required Status Check
	Restrictions GeneratedType `json:"restrictions,omitempty"` // Branch Restriction Policy
	Allow_force_pushes map[string]interface{} `json:"allow_force_pushes,omitempty"`
	Enforce_admins GeneratedType `json:"enforce_admins,omitempty"` // Protected Branch Admin Enforced
	Block_creations map[string]interface{} `json:"block_creations,omitempty"`
	Protection_url string `json:"protection_url,omitempty"`
	Required_pull_request_reviews GeneratedType `json:"required_pull_request_reviews,omitempty"` // Protected Branch Pull Request Review
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Old_answer map[string]interface{} `json:"old_answer"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"` // The REST API URL of the alert resource.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Dismissed_comment string `json:"dismissed_comment"` // An optional comment associated with the alert's dismissal.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Security_advisory GeneratedType `json:"security_advisory"` // Details for the GitHub Security Advisory.
	Security_vulnerability GeneratedType `json:"security_vulnerability"` // Details pertaining to one vulnerable version range for the advisory.
	Updated_at string `json:"updated_at"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_reason string `json:"dismissed_reason"` // The reason that the alert was dismissed.
	Number int `json:"number"` // The security alert number.
	Dependency map[string]interface{} `json:"dependency"` // Details for the vulnerable dependency.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Fixed_at string `json:"fixed_at"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	State string `json:"state"` // The state of the Dependabot alert.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Scope string `json:"scope"` // The scope of the membership. Currently, can only be `team`.
	Team map[string]interface{} `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Member map[string]interface{} `json:"member"`
	Sender map[string]interface{} `json:"sender"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Added_by string `json:"added_by,omitempty"`
	Created_at string `json:"created_at"`
	Key string `json:"key"`
	Title string `json:"title"`
	Last_used string `json:"last_used,omitempty"`
	Url string `json:"url"`
	Id int `json:"id"`
	Read_only bool `json:"read_only"`
	Verified bool `json:"verified"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Links map[string]interface{} `json:"_links"`
	Download_url string `json:"download_url"`
	Html_url string `json:"html_url"`
	Name string `json:"name"`
	TypeField string `json:"type"`
	Url string `json:"url"`
	Git_url string `json:"git_url"`
	Entries []map[string]interface{} `json:"entries,omitempty"`
	Path string `json:"path"`
	Sha string `json:"sha"`
	Size int `json:"size"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Single_file_name string `json:"single_file_name"`
	Single_file_paths []string `json:"single_file_paths,omitempty"`
	Account GeneratedType `json:"account"` // A GitHub user.
	Has_multiple_single_files bool `json:"has_multiple_single_files,omitempty"`
	Permissions GeneratedType `json:"permissions"` // The permissions granted to the user-to-server access token.
	Repositories_url string `json:"repositories_url"`
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Preview_url string `json:"preview_url,omitempty"` // The URI to the deployed GitHub Pages preview.
	Status_url string `json:"status_url"` // The URI to monitor GitHub Pages deployment status.
	Page_url string `json:"page_url"` // The URI to the deployed GitHub Pages.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Label map[string]interface{} `json:"label,omitempty"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Number int `json:"number"` // The pull request number.
	Action string `json:"action"`
	Pull_request map[string]interface{} `json:"pull_request"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Label map[string]interface{} `json:"label,omitempty"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Teams []Team `json:"teams"`
	Users []GeneratedType `json:"users"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"`
	Href string `json:"href"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Event string `json:"event"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_url string `json:"commit_url"`
	Id int `json:"id"`
	Lock_reason string `json:"lock_reason"`
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Rate GeneratedType `json:"rate"`
	Resources map[string]interface{} `json:"resources"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_sha string `json:"commit_sha"` // SHA-1 hash ID of the associated commit
	Path string `json:"path"` // The file path in the repository
	Start_column float64 `json:"start_column"` // The column at which the secret starts within the start line when the file is interpreted as 8BIT ASCII
	Blob_sha string `json:"blob_sha"` // SHA-1 hash ID of the associated blob
	Blob_url string `json:"blob_url"` // The API URL to get the associated blob resource
	End_line float64 `json:"end_line"` // Line number at which the secret ends in the file
	Commit_url string `json:"commit_url"` // The API URL to get the associated commit resource
	End_column float64 `json:"end_column"` // The column at which the secret ends within the end line when the file is interpreted as 8BIT ASCII
	Start_line float64 `json:"start_line"` // Line number at which the secret starts in the file
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enabled_by GeneratedType `json:"enabled_by"` // A GitHub user.
	Merge_method string `json:"merge_method"` // The merge method to use.
	Commit_message string `json:"commit_message"` // Commit message for the merge commit.
	Commit_title string `json:"commit_title"` // Title for the merge commit message.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Title string `json:"title"` // The title of the milestone.
	Updated_at string `json:"updated_at"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Number int `json:"number"` // The number of the milestone.
	Description string `json:"description"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Created_at string `json:"created_at"`
	Due_on string `json:"due_on"`
	Open_issues int `json:"open_issues"`
	State string `json:"state"` // The state of the milestone.
	Closed_at string `json:"closed_at"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Closed_issues int `json:"closed_issues"`
	Labels_url string `json:"labels_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Review map[string]interface{} `json:"review"` // The review that was affected.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Label map[string]interface{} `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Has_downloads bool `json:"has_downloads,omitempty"`
	Id int `json:"id"`
	Clone_url string `json:"clone_url,omitempty"`
	Statuses_url string `json:"statuses_url"`
	Commits_url string `json:"commits_url"`
	Stargazers_count int `json:"stargazers_count,omitempty"`
	Events_url string `json:"events_url"`
	Collaborators_url string `json:"collaborators_url"`
	Archived bool `json:"archived,omitempty"`
	Languages_url string `json:"languages_url"`
	Private bool `json:"private"`
	Notifications_url string `json:"notifications_url"`
	Git_refs_url string `json:"git_refs_url"`
	Keys_url string `json:"keys_url"`
	Open_issues_count int `json:"open_issues_count,omitempty"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Allow_forking bool `json:"allow_forking,omitempty"`
	Security_and_analysis GeneratedType `json:"security_and_analysis,omitempty"`
	Branches_url string `json:"branches_url"`
	License map[string]interface{} `json:"license,omitempty"`
	Has_wiki bool `json:"has_wiki,omitempty"`
	Name string `json:"name"`
	Full_name string `json:"full_name"`
	Size int `json:"size,omitempty"` // The size of the repository. Size is calculated hourly. When a repository is initially created, the size is 0.
	Ssh_url string `json:"ssh_url,omitempty"`
	Blobs_url string `json:"blobs_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Tags_url string `json:"tags_url"`
	Node_id string `json:"node_id"`
	Subscription_url string `json:"subscription_url"`
	Forks_count int `json:"forks_count,omitempty"`
	Html_url string `json:"html_url"`
	Labels_url string `json:"labels_url"`
	Default_branch string `json:"default_branch,omitempty"`
	Forks int `json:"forks,omitempty"`
	Description string `json:"description"`
	Visibility string `json:"visibility,omitempty"`
	Teams_url string `json:"teams_url"`
	Fork bool `json:"fork"`
	Comments_url string `json:"comments_url"`
	Deployments_url string `json:"deployments_url"`
	Watchers_count int `json:"watchers_count,omitempty"`
	Compare_url string `json:"compare_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Forks_url string `json:"forks_url"`
	Merges_url string `json:"merges_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Archive_url string `json:"archive_url"`
	Watchers int `json:"watchers,omitempty"`
	Has_projects bool `json:"has_projects,omitempty"`
	Subscribers_url string `json:"subscribers_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Has_pages bool `json:"has_pages,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Git_commits_url string `json:"git_commits_url"`
	Url string `json:"url"`
	Stargazers_url string `json:"stargazers_url"`
	Is_template bool `json:"is_template,omitempty"`
	Has_discussions bool `json:"has_discussions,omitempty"`
	Language string `json:"language,omitempty"`
	Git_url string `json:"git_url,omitempty"`
	Role_name string `json:"role_name,omitempty"`
	Pulls_url string `json:"pulls_url"`
	Git_tags_url string `json:"git_tags_url"`
	Releases_url string `json:"releases_url"`
	Code_of_conduct GeneratedType `json:"code_of_conduct,omitempty"` // Code Of Conduct
	Pushed_at string `json:"pushed_at,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	Hooks_url string `json:"hooks_url"`
	Disabled bool `json:"disabled,omitempty"`
	Has_issues bool `json:"has_issues,omitempty"`
	Trees_url string `json:"trees_url"`
	Milestones_url string `json:"milestones_url"`
	Mirror_url string `json:"mirror_url,omitempty"`
	Network_count int `json:"network_count,omitempty"`
	Topics []string `json:"topics,omitempty"`
	Svn_url string `json:"svn_url,omitempty"`
	Contributors_url string `json:"contributors_url"`
	Issue_events_url string `json:"issue_events_url"`
	Assignees_url string `json:"assignees_url"`
	Downloads_url string `json:"downloads_url"`
	Open_issues int `json:"open_issues,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Issues_url string `json:"issues_url"`
	Contents_url string `json:"contents_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Project_card map[string]interface{} `json:"project_card"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Object map[string]interface{} `json:"object"`
	Ref string `json:"ref"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total_private_repos int `json:"total_private_repos"`
	Created_at string `json:"created_at"`
	Following_url string `json:"following_url"`
	Hireable bool `json:"hireable"`
	Url string `json:"url"`
	Received_events_url string `json:"received_events_url"`
	Private_gists int `json:"private_gists"`
	Updated_at string `json:"updated_at"`
	Company string `json:"company"`
	Followers_url string `json:"followers_url"`
	Bio string `json:"bio"`
	Following int `json:"following"`
	Email string `json:"email"`
	Followers int `json:"followers"`
	Location string `json:"location"`
	Suspended_at string `json:"suspended_at,omitempty"`
	Organizations_url string `json:"organizations_url"`
	Public_gists int `json:"public_gists"`
	Blog string `json:"blog"`
	Business_plus bool `json:"business_plus,omitempty"`
	Ldap_dn string `json:"ldap_dn,omitempty"`
	Owned_private_repos int `json:"owned_private_repos"`
	Two_factor_authentication bool `json:"two_factor_authentication"`
	Gravatar_id string `json:"gravatar_id"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Disk_usage int `json:"disk_usage"`
	Starred_url string `json:"starred_url"`
	Subscriptions_url string `json:"subscriptions_url"`
	Node_id string `json:"node_id"`
	Name string `json:"name"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Repos_url string `json:"repos_url"`
	Gists_url string `json:"gists_url"`
	Public_repos int `json:"public_repos"`
	TypeField string `json:"type"`
	Login string `json:"login"`
	Collaborators int `json:"collaborators"`
	Id int `json:"id"`
	Site_admin bool `json:"site_admin"`
	Avatar_url string `json:"avatar_url"`
	Events_url string `json:"events_url"`
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Avatar_url string `json:"avatar_url"`
	Events_url string `json:"events_url"`
	Hooks_url string `json:"hooks_url"`
	Id int `json:"id"`
	Issues_url string `json:"issues_url"`
	Description string `json:"description"`
	Login string `json:"login"`
	Node_id string `json:"node_id"`
	Members_url string `json:"members_url"`
	Public_members_url string `json:"public_members_url"`
	Repos_url string `json:"repos_url"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Event string `json:"event"`
	Id int `json:"id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_url string `json:"commit_url"`
	Dismissed_review map[string]interface{} `json:"dismissed_review"`
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ssh_url string `json:"ssh_url"`
	Git_commits_url string `json:"git_commits_url"`
	Blobs_url string `json:"blobs_url"`
	Pushed_at string `json:"pushed_at"`
	Merges_url string `json:"merges_url"`
	Subscription_url string `json:"subscription_url"`
	Archive_url string `json:"archive_url"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Watchers_count int `json:"watchers_count"`
	Languages_url string `json:"languages_url"`
	Clone_url string `json:"clone_url"`
	Forks int `json:"forks"`
	Network_count int `json:"network_count,omitempty"`
	Hooks_url string `json:"hooks_url"`
	Has_pages bool `json:"has_pages"`
	Tags_url string `json:"tags_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Statuses_url string `json:"statuses_url"`
	Id int `json:"id"` // Unique identifier of the repository
	Comments_url string `json:"comments_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Open_issues int `json:"open_issues"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Branches_url string `json:"branches_url"`
	Stargazers_count int `json:"stargazers_count"`
	Homepage string `json:"homepage"`
	Trees_url string `json:"trees_url"`
	Events_url string `json:"events_url"`
	Issue_events_url string `json:"issue_events_url"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Full_name string `json:"full_name"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Git_refs_url string `json:"git_refs_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Updated_at string `json:"updated_at"`
	Forks_url string `json:"forks_url"`
	Role_name string `json:"role_name,omitempty"`
	Fork bool `json:"fork"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Topics []string `json:"topics,omitempty"`
	Description string `json:"description"`
	Issue_comment_url string `json:"issue_comment_url"`
	Pulls_url string `json:"pulls_url"`
	Open_issues_count int `json:"open_issues_count"`
	Stargazers_url string `json:"stargazers_url"`
	Html_url string `json:"html_url"`
	Contents_url string `json:"contents_url"`
	Master_branch string `json:"master_branch,omitempty"`
	Git_tags_url string `json:"git_tags_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Subscribers_url string `json:"subscribers_url"`
	Commits_url string `json:"commits_url"`
	Downloads_url string `json:"downloads_url"`
	Language string `json:"language"`
	Contributors_url string `json:"contributors_url"`
	Releases_url string `json:"releases_url"`
	Assignees_url string `json:"assignees_url"`
	Labels_url string `json:"labels_url"`
	License GeneratedType `json:"license"` // License Simple
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Node_id string `json:"node_id"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Notifications_url string `json:"notifications_url"`
	Deployments_url string `json:"deployments_url"`
	Name string `json:"name"` // The name of the repository.
	Collaborators_url string `json:"collaborators_url"`
	Size int `json:"size"`
	Url string `json:"url"`
	Teams_url string `json:"teams_url"`
	Svn_url string `json:"svn_url"`
	Compare_url string `json:"compare_url"`
	Template_repository GeneratedType `json:"template_repository,omitempty"` // A repository on GitHub.
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Issues_url string `json:"issues_url"`
	Watchers int `json:"watchers"`
	Git_url string `json:"git_url"`
	Mirror_url string `json:"mirror_url"`
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Milestones_url string `json:"milestones_url"`
	Private bool `json:"private"` // Whether the repository is private or public.
	Forks_count int `json:"forks_count"`
	Keys_url string `json:"keys_url"`
	Created_at string `json:"created_at"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Storage_in_bytes int `json:"storage_in_bytes"` // How much storage is available to the codespace.
	Cpus int `json:"cpus"` // How many cores are available to the codespace.
	Display_name string `json:"display_name"` // The display name of the machine includes cores, memory, and storage.
	Memory_in_bytes int `json:"memory_in_bytes"` // How much memory is available to the codespace.
	Name string `json:"name"` // The name of the machine.
	Operating_system string `json:"operating_system"` // The operating system of the machine.
	Prebuild_availability string `json:"prebuild_availability"` // Whether a prebuild is currently available when creating a codespace for this machine and repository. If a branch was not specified as a ref, the default branch will be assumed. Value will be "null" if prebuilds are not supported or prebuild availability could not be determined. Value will be "none" if no prebuild is available. Latest values "ready" and "in_progress" indicate the prebuild availability status.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository_url string `json:"repository_url"`
	Subscribed bool `json:"subscribed"` // Determines if notifications should be received from this repository.
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Ignored bool `json:"ignored"` // Determines if all notifications should be blocked from this repository.
	Reason string `json:"reason"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Url string `json:"url"`
	Commit_url string `json:"commit_url"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Requested_reviewer GeneratedType `json:"requested_reviewer,omitempty"` // A GitHub user.
	Requested_team Team `json:"requested_team,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Review_requester GeneratedType `json:"review_requester"` // A GitHub user.
	Event string `json:"event"`
	Id int `json:"id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_url string `json:"commit_url"`
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Commit_id string `json:"commit_id"`
	Event string `json:"event"`
	Id int `json:"id"`
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Security_advisory map[string]interface{} `json:"security_advisory"` // The details of the security advisory, including summary, description, and severity.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Expires_at string `json:"expires_at"`
	Can_certify bool `json:"can_certify"`
	Primary_key_id int `json:"primary_key_id"`
	Public_key string `json:"public_key"`
	Can_encrypt_comms bool `json:"can_encrypt_comms"`
	Created_at string `json:"created_at"`
	Revoked bool `json:"revoked"`
	Can_encrypt_storage bool `json:"can_encrypt_storage"`
	Emails []map[string]interface{} `json:"emails"`
	Raw_key string `json:"raw_key"`
	Id int `json:"id"`
	Can_sign bool `json:"can_sign"`
	Name string `json:"name,omitempty"`
	Key_id string `json:"key_id"`
	Subkeys []map[string]interface{} `json:"subkeys"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Key map[string]interface{} `json:"key"` // The [`deploy key`](https://docs.github.com/rest/reference/deployments#get-a-deploy-key) resource.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Permission string `json:"permission"`
	User GeneratedType `json:"user"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Actions_meta map[string]interface{} `json:"actions_meta,omitempty"`
	Check_suite map[string]interface{} `json:"check_suite"` // The [check_suite](https://docs.github.com/rest/reference/checks#suites).
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue interface{} `json:"issue"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Team map[string]interface{} `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_requests []GeneratedType `json:"pull_requests,omitempty"`
	Before string `json:"before,omitempty"`
	Id int `json:"id,omitempty"`
	App Integration `json:"app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Created_at string `json:"created_at,omitempty"`
	Repository GeneratedType `json:"repository,omitempty"` // Minimal Repository
	Updated_at string `json:"updated_at,omitempty"`
	After string `json:"after,omitempty"`
	Conclusion string `json:"conclusion,omitempty"`
	Head_branch string `json:"head_branch,omitempty"`
	Head_sha string `json:"head_sha,omitempty"` // The SHA of the head commit that is being checked.
	Status string `json:"status,omitempty"`
	Url string `json:"url,omitempty"`
	Node_id string `json:"node_id,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Deliveries_url string `json:"deliveries_url,omitempty"`
	Name string `json:"name"`
	Updated_at string `json:"updated_at"`
	Config map[string]interface{} `json:"config"`
	Events []string `json:"events"`
	TypeField string `json:"type"`
	Url string `json:"url"`
	Active bool `json:"active"`
	Id int `json:"id"`
	Ping_url string `json:"ping_url"`
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"` // Unique identifier of the deployment
	Transient_environment bool `json:"transient_environment,omitempty"` // Specifies if the given environment is will no longer exist at some point in the future. Default: false.
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Environment string `json:"environment"` // Name for the target deployment environment.
	Original_environment string `json:"original_environment,omitempty"`
	Production_environment bool `json:"production_environment,omitempty"` // Specifies if the given environment is one that end-users directly interact with. Default: false.
	Repository_url string `json:"repository_url"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Task string `json:"task"` // Parameter to specify a task to execute
	Description string `json:"description"`
	Node_id string `json:"node_id"`
	Statuses_url string `json:"statuses_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Node_id string `json:"node_id"`
	Id int `json:"id"`
	Performed_via_github_app Integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_url string `json:"commit_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"` // Unique identifier of the delivery.
	Redelivery bool `json:"redelivery"` // Whether the delivery is a redelivery.
	Status string `json:"status"` // Description of the status of the attempted delivery
	Event string `json:"event"` // The event that triggered the delivery.
	Repository_id int `json:"repository_id"` // The id of the repository associated with this event.
	Duration float64 `json:"duration"` // Time spent delivering.
	Request map[string]interface{} `json:"request"`
	Response map[string]interface{} `json:"response"`
	Url string `json:"url,omitempty"` // The URL target of the delivery.
	Action string `json:"action"` // The type of activity for the event that triggered the delivery.
	Delivered_at string `json:"delivered_at"` // Time when the delivery was delivered.
	Guid string `json:"guid"` // Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event).
	Installation_id int `json:"installation_id"` // The id of the GitHub App installation associated with this event.
	Status_code int `json:"status_code"` // Status code received when delivery was made.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Title string `json:"title"`
	End_line int `json:"end_line"`
	Message string `json:"message"`
	Path string `json:"path"`
	Raw_details string `json:"raw_details"`
	Start_column int `json:"start_column"`
	Blob_href string `json:"blob_href"`
	End_column int `json:"end_column"`
	Annotation_level string `json:"annotation_level"`
	Start_line int `json:"start_line"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Badge_url string `json:"badge_url"`
	Html_url string `json:"html_url"`
	Path string `json:"path"`
	Source_repository GeneratedType `json:"source_repository"` // Minimal Repository
	State string `json:"state"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Color string `json:"color"`
	DefaultField bool `json:"default"`
	Id int `json:"id"`
	Name string `json:"name"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Score float64 `json:"score"`
	Description string `json:"description"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Label map[string]interface{} `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion map[string]interface{} `json:"discussion"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/reference/issues) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action that was performed.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Context string `json:"context"`
	Target_url string `json:"target_url"`
	Updated_at string `json:"updated_at"`
	Avatar_url string `json:"avatar_url"`
	Description string `json:"description"`
	Required bool `json:"required,omitempty"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	State string `json:"state"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Member map[string]interface{} `json:"member"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Scope string `json:"scope"` // The scope of the membership. Currently, can only be `team`.
	Sender map[string]interface{} `json:"sender"`
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Team map[string]interface{} `json:"team"` // Groups of organization members that gives permissions on specified repositories.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Release map[string]interface{} `json:"release"` // The [release](https://docs.github.com/rest/reference/repos/#get-a-release) object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Remote_name string `json:"remote_name"`
	Url string `json:"url"`
	Email string `json:"email"`
	Id int `json:"id"`
	Import_url string `json:"import_url"`
	Name string `json:"name"`
	Remote_id string `json:"remote_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Key map[string]interface{} `json:"key"` // The [`deploy key`](https://docs.github.com/rest/reference/deployments#get-a-deploy-key) resource.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"`
	Color string `json:"color"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Display_name string `json:"display_name"` // The display name of the machine includes cores, memory, and storage.
	Memory_in_bytes int `json:"memory_in_bytes"` // How much memory is available to the codespace.
	Name string `json:"name"` // The name of the machine.
	Operating_system string `json:"operating_system"` // The operating system of the machine.
	Prebuild_availability string `json:"prebuild_availability"` // Whether a prebuild is currently available when creating a codespace for this machine and repository. If a branch was not specified as a ref, the default branch will be assumed. Value will be "null" if prebuilds are not supported or prebuild availability could not be determined. Value will be "none" if no prebuild is available. Latest values "ready" and "in_progress" indicate the prebuild availability status.
	Storage_in_bytes int `json:"storage_in_bytes"` // How much storage is available to the codespace.
	Cpus int `json:"cpus"` // How many cores are available to the codespace.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Projects_v2 GeneratedType `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sha string `json:"sha"`
	Tree []map[string]interface{} `json:"tree"` // Objects specifying a tree structure
	Truncated bool `json:"truncated"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key string `json:"key"` // The Base64 encoded public key.
	Key_id string `json:"key_id"` // The identifier for the key.
	Title string `json:"title,omitempty"`
	Url string `json:"url,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Id int `json:"id,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Deployments_url string `json:"deployments_url"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"`
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Code_of_conduct GeneratedType `json:"code_of_conduct,omitempty"` // Code of Conduct Simple
	Collaborators_url string `json:"collaborators_url"`
	Private bool `json:"private"`
	Labels_url string `json:"labels_url"`
	Watchers int `json:"watchers"`
	Master_branch string `json:"master_branch,omitempty"`
	Comments_url string `json:"comments_url"`
	Contributors_url string `json:"contributors_url"`
	Git_url string `json:"git_url"`
	Forks_count int `json:"forks_count"`
	Stargazers_count int `json:"stargazers_count"`
	Subscribers_count int `json:"subscribers_count"`
	Languages_url string `json:"languages_url"`
	Issues_url string `json:"issues_url"`
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Updated_at string `json:"updated_at"`
	Id int `json:"id"`
	Topics []string `json:"topics,omitempty"`
	Size int `json:"size"` // The size of the repository. Size is calculated hourly. When a repository is initially created, the size is 0.
	Homepage string `json:"homepage"`
	Watchers_count int `json:"watchers_count"`
	Issue_events_url string `json:"issue_events_url"`
	Default_branch string `json:"default_branch"`
	Html_url string `json:"html_url"`
	Has_wiki bool `json:"has_wiki"`
	Parent Repository `json:"parent,omitempty"` // A repository on GitHub.
	Source Repository `json:"source,omitempty"` // A repository on GitHub.
	Keys_url string `json:"keys_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"`
	Compare_url string `json:"compare_url"`
	Open_issues_count int `json:"open_issues_count"`
	Mirror_url string `json:"mirror_url"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Git_refs_url string `json:"git_refs_url"`
	Created_at string `json:"created_at"`
	Forks_url string `json:"forks_url"`
	Releases_url string `json:"releases_url"`
	Subscribers_url string `json:"subscribers_url"`
	Is_template bool `json:"is_template,omitempty"`
	Merges_url string `json:"merges_url"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Tags_url string `json:"tags_url"`
	Branches_url string `json:"branches_url"`
	Full_name string `json:"full_name"`
	Has_downloads bool `json:"has_downloads"`
	Assignees_url string `json:"assignees_url"`
	Statuses_url string `json:"statuses_url"`
	Has_projects bool `json:"has_projects"`
	Has_pages bool `json:"has_pages"`
	Downloads_url string `json:"downloads_url"`
	Git_tags_url string `json:"git_tags_url"`
	Pushed_at string `json:"pushed_at"`
	Commits_url string `json:"commits_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Network_count int `json:"network_count"`
	License GeneratedType `json:"license"` // License Simple
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub user.
	Node_id string `json:"node_id"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"`
	Hooks_url string `json:"hooks_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Security_and_analysis GeneratedType `json:"security_and_analysis,omitempty"`
	Notifications_url string `json:"notifications_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"`
	Has_issues bool `json:"has_issues"`
	Events_url string `json:"events_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Blobs_url string `json:"blobs_url"`
	Template_repository GeneratedType `json:"template_repository,omitempty"` // A repository on GitHub.
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Svn_url string `json:"svn_url"`
	Milestones_url string `json:"milestones_url"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"`
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is allowed.
	Description string `json:"description"`
	Url string `json:"url"`
	Name string `json:"name"`
	Teams_url string `json:"teams_url"`
	Fork bool `json:"fork"`
	Subscription_url string `json:"subscription_url"`
	Archive_url string `json:"archive_url"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Allow_forking bool `json:"allow_forking,omitempty"`
	Stargazers_url string `json:"stargazers_url"`
	Trees_url string `json:"trees_url"`
	Contents_url string `json:"contents_url"`
	Ssh_url string `json:"ssh_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Clone_url string `json:"clone_url"`
	Language string `json:"language"`
	Pulls_url string `json:"pulls_url"`
	Has_discussions bool `json:"has_discussions"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"`
	Forks int `json:"forks"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Open_issues int `json:"open_issues"`
	Archived bool `json:"archived"`
	Git_commits_url string `json:"git_commits_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Git_push_url string `json:"git_push_url"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	History []interface{} `json:"history,omitempty"`
	Public bool `json:"public"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Git_pull_url string `json:"git_pull_url"`
	Commits_url string `json:"commits_url"`
	Files map[string]interface{} `json:"files"`
	Forks_url string `json:"forks_url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Comments int `json:"comments"`
	Comments_url string `json:"comments_url"`
	Forks []interface{} `json:"forks,omitempty"`
	Owner GeneratedType `json:"owner,omitempty"` // A GitHub user.
	Id string `json:"id"`
	Html_url string `json:"html_url"`
	Truncated bool `json:"truncated,omitempty"`
	Description string `json:"description"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action,omitempty"`
	Check_run GeneratedType `json:"check_run"` // A check performed on the code of a given code change
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// Migration represents the Migration schema from the OpenAPI specification
type Migration struct {
	Exclude_releases bool `json:"exclude_releases"`
	Exclude []interface{} `json:"exclude,omitempty"`
	State string `json:"state"`
	Lock_repositories bool `json:"lock_repositories"`
	Updated_at string `json:"updated_at"`
	Archive_url string `json:"archive_url,omitempty"`
	Exclude_attachments bool `json:"exclude_attachments"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Org_metadata_only bool `json:"org_metadata_only"`
	Url string `json:"url"`
	Exclude_git_data bool `json:"exclude_git_data"`
	Id int `json:"id"`
	Repositories []Repository `json:"repositories"` // The repositories included in the migration. Only returned for export migrations.
	Exclude_metadata bool `json:"exclude_metadata"`
	Guid string `json:"guid"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Exclude_owner_projects bool `json:"exclude_owner_projects"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment map[string]interface{} `json:"comment"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sha string `json:"sha"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Git_url string `json:"git_url"`
	Html_url string `json:"html_url"`
	Language string `json:"language,omitempty"`
	Path string `json:"path"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Score float64 `json:"score"`
	Name string `json:"name"`
	Url string `json:"url"`
	File_size int `json:"file_size,omitempty"`
	Last_modified_at string `json:"last_modified_at,omitempty"`
	Line_numbers []string `json:"line_numbers,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Version string `json:"version,omitempty"` // The version of the tool used to generate the code scanning analysis.
	Guid string `json:"guid,omitempty"` // The GUID of the tool used to generate the code scanning analysis, if provided in the uploaded SARIF data.
	Name string `json:"name,omitempty"` // The name of the tool used to generate the code scanning analysis.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Account map[string]interface{} `json:"account"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Changes map[string]interface{} `json:"changes"`
	Target_type string `json:"target_type"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Security_advisory map[string]interface{} `json:"security_advisory"` // The details of the security advisory, including summary, description, and severity.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the ping JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Company string `json:"company"`
	Private_gists int `json:"private_gists,omitempty"`
	Public_repos int `json:"public_repos"`
	Following int `json:"following"`
	Email string `json:"email"`
	Location string `json:"location"`
	Repos_url string `json:"repos_url"`
	Followers_url string `json:"followers_url"`
	Updated_at string `json:"updated_at"`
	Site_admin bool `json:"site_admin"`
	Disk_usage int `json:"disk_usage,omitempty"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Blog string `json:"blog"`
	Avatar_url string `json:"avatar_url"`
	Created_at string `json:"created_at"`
	Url string `json:"url"`
	Starred_url string `json:"starred_url"`
	Suspended_at string `json:"suspended_at,omitempty"`
	Hireable bool `json:"hireable"`
	Owned_private_repos int `json:"owned_private_repos,omitempty"`
	Total_private_repos int `json:"total_private_repos,omitempty"`
	Login string `json:"login"`
	Following_url string `json:"following_url"`
	TypeField string `json:"type"`
	Public_gists int `json:"public_gists"`
	Name string `json:"name"`
	Subscriptions_url string `json:"subscriptions_url"`
	Received_events_url string `json:"received_events_url"`
	Collaborators int `json:"collaborators,omitempty"`
	Events_url string `json:"events_url"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Organizations_url string `json:"organizations_url"`
	Gravatar_id string `json:"gravatar_id"`
	Followers int `json:"followers"`
	Html_url string `json:"html_url"`
	Gists_url string `json:"gists_url"`
	Bio string `json:"bio"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Content_type string `json:"content_type"` // The MIME type of the CodeQL database file.
	Created_at string `json:"created_at"` // The date and time at which the CodeQL database was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Size int `json:"size"` // The size of the CodeQL database file in bytes.
	Uploader GeneratedType `json:"uploader"` // A GitHub user.
	Id int `json:"id"` // The ID of the CodeQL database.
	Updated_at string `json:"updated_at"` // The date and time at which the CodeQL database was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Language string `json:"language"` // The language of the CodeQL database.
	Name string `json:"name"` // The name of the CodeQL database.
	Url string `json:"url"` // The URL at which to download the CodeQL database. The `Accept` header must be set to the value of the `content_type` property.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Visibility string `json:"visibility"` // The type of repositories in the organization that the secret is visible to
	Created_at string `json:"created_at"` // The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the secret
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"` // The API URL at which the list of repositories this secret is visible to can be retrieved
	Updated_at string `json:"updated_at"` // The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Actions_meta map[string]interface{} `json:"actions_meta,omitempty"`
	Check_suite map[string]interface{} `json:"check_suite"` // The [check_suite](https://docs.github.com/rest/reference/checks#suites).
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Repository Repository `json:"repository"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Project_column map[string]interface{} `json:"project_column"`
	Repository GeneratedType `json:"repository,omitempty"` // A repository on GitHub.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise Enterprise `json:"enterprise,omitempty"` // An enterprise on GitHub.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. This property is included when the event is configured for and sent to a GitHub App.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
}

// Root represents the Root schema from the OpenAPI specification
type Root struct {
	Starred_gists_url string `json:"starred_gists_url"`
	Topic_search_url string `json:"topic_search_url,omitempty"`
	Gists_url string `json:"gists_url"`
	Issue_search_url string `json:"issue_search_url"`
	User_organizations_url string `json:"user_organizations_url"`
	Issues_url string `json:"issues_url"`
	Emojis_url string `json:"emojis_url"`
	Events_url string `json:"events_url"`
	User_url string `json:"user_url"`
	User_search_url string `json:"user_search_url"`
	User_repositories_url string `json:"user_repositories_url"`
	Starred_url string `json:"starred_url"`
	Repository_url string `json:"repository_url"`
	Emails_url string `json:"emails_url"`
	Repository_search_url string `json:"repository_search_url"`
	Authorizations_url string `json:"authorizations_url"`
	Rate_limit_url string `json:"rate_limit_url"`
	Notifications_url string `json:"notifications_url"`
	Current_user_repositories_url string `json:"current_user_repositories_url"`
	Current_user_authorizations_html_url string `json:"current_user_authorizations_html_url"`
	Following_url string `json:"following_url"`
	Followers_url string `json:"followers_url"`
	Organization_repositories_url string `json:"organization_repositories_url"`
	Keys_url string `json:"keys_url"`
	Current_user_url string `json:"current_user_url"`
	Commit_search_url string `json:"commit_search_url"`
	Organization_teams_url string `json:"organization_teams_url"`
	Hub_url string `json:"hub_url"`
	Organization_url string `json:"organization_url"`
	Public_gists_url string `json:"public_gists_url"`
	Feeds_url string `json:"feeds_url"`
	Code_search_url string `json:"code_search_url"`
	Label_search_url string `json:"label_search_url"`
}
