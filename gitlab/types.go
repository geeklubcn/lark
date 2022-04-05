package gitlab

import (
	"time"
)

// IssueResult https://docs.gitlab.com/ee/api/issues.html
type IssueResult struct {
	Id          int         `json:"id"`
	Iid         int         `json:"iid"`
	ProjectId   int         `json:"project_id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	State       string      `json:"state"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	ClosedAt    interface{} `json:"closed_at"`
	ClosedBy    interface{} `json:"closed_by"`
	Labels      []string    `json:"labels"`
	Milestone   struct {
		Id          int       `json:"id"`
		Iid         int       `json:"iid"`
		ProjectId   int       `json:"project_id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		State       string    `json:"state"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		DueDate     string    `json:"due_date"`
		StartDate   string    `json:"start_date"`
		Expired     bool      `json:"expired"`
		WebUrl      string    `json:"web_url"`
	} `json:"milestone"`
	Assignees []struct {
		Id        int    `json:"id"`
		Username  string `json:"username"`
		Name      string `json:"name"`
		State     string `json:"state"`
		AvatarUrl string `json:"avatar_url"`
		WebUrl    string `json:"web_url"`
	} `json:"assignees"`
	Author struct {
		Id        int    `json:"id"`
		Username  string `json:"username"`
		Name      string `json:"name"`
		State     string `json:"state"`
		AvatarUrl string `json:"avatar_url"`
		WebUrl    string `json:"web_url"`
	} `json:"author"`
	Type     string `json:"type"`
	Assignee struct {
		Id        int    `json:"id"`
		Username  string `json:"username"`
		Name      string `json:"name"`
		State     string `json:"state"`
		AvatarUrl string `json:"avatar_url"`
		WebUrl    string `json:"web_url"`
	} `json:"assignee"`
	UserNotesCount     int         `json:"user_notes_count"`
	MergeRequestsCount int         `json:"merge_requests_count"`
	Upvotes            int         `json:"upvotes"`
	Downvotes          int         `json:"downvotes"`
	DueDate            interface{} `json:"due_date"`
	Confidential       bool        `json:"confidential"`
	DiscussionLocked   interface{} `json:"discussion_locked"`
	IssueType          string      `json:"issue_type"`
	WebUrl             string      `json:"web_url"`
	TimeStats          struct {
		TimeEstimate        int         `json:"time_estimate"`
		TotalTimeSpent      int         `json:"total_time_spent"`
		HumanTimeEstimate   interface{} `json:"human_time_estimate"`
		HumanTotalTimeSpent interface{} `json:"human_total_time_spent"`
	} `json:"time_stats"`
	TaskCompletionStatus struct {
		Count          int `json:"count"`
		CompletedCount int `json:"completed_count"`
	} `json:"task_completion_status"`
	Weight              interface{} `json:"weight"`
	BlockingIssuesCount int         `json:"blocking_issues_count"`
	HasTasks            bool        `json:"has_tasks"`
	Links               struct {
		Self       string `json:"self"`
		Notes      string `json:"notes"`
		AwardEmoji string `json:"award_emoji"`
		Project    string `json:"project"`
	} `json:"_links"`
	References struct {
		Short    string `json:"short"`
		Relative string `json:"relative"`
		Full     string `json:"full"`
	} `json:"references"`
	Severity           string      `json:"severity"`
	MovedToId          interface{} `json:"moved_to_id"`
	ServiceDeskReplyTo interface{} `json:"service_desk_reply_to"`
	HealthStatus       string      `json:"health_status"`
}

// ProjectResult https://docs.gitlab.com/ee/api/projects.html#get-single-project
type ProjectResult struct {
	Id            int      `json:"id"`
	Description   string   `json:"description"`
	DefaultBranch string   `json:"default_branch"`
	Visibility    string   `json:"visibility"`
	SshUrlToRepo  string   `json:"ssh_url_to_repo"`
	HttpUrlToRepo string   `json:"http_url_to_repo"`
	WebUrl        string   `json:"web_url"`
	ReadmeUrl     string   `json:"readme_url"`
	TagList       []string `json:"tag_list"`
	Topics        []string `json:"topics"`
	Owner         struct {
		Id        int       `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"owner"`
	Name                             string `json:"name"`
	NameWithNamespace                string `json:"name_with_namespace"`
	Path                             string `json:"path"`
	PathWithNamespace                string `json:"path_with_namespace"`
	IssuesEnabled                    bool   `json:"issues_enabled"`
	OpenIssuesCount                  int    `json:"open_issues_count"`
	MergeRequestsEnabled             bool   `json:"merge_requests_enabled"`
	JobsEnabled                      bool   `json:"jobs_enabled"`
	WikiEnabled                      bool   `json:"wiki_enabled"`
	SnippetsEnabled                  bool   `json:"snippets_enabled"`
	CanCreateMergeRequestIn          bool   `json:"can_create_merge_request_in"`
	ResolveOutdatedDiffDiscussions   bool   `json:"resolve_outdated_diff_discussions"`
	ContainerRegistryEnabled         bool   `json:"container_registry_enabled"`
	ContainerRegistryAccessLevel     string `json:"container_registry_access_level"`
	SecurityAndComplianceAccessLevel string `json:"security_and_compliance_access_level"`
	ContainerExpirationPolicy        struct {
		Cadence         string      `json:"cadence"`
		Enabled         bool        `json:"enabled"`
		KeepN           interface{} `json:"keep_n"`
		OlderThan       interface{} `json:"older_than"`
		NameRegex       interface{} `json:"name_regex"`
		NameRegexDelete interface{} `json:"name_regex_delete"`
		NameRegexKeep   interface{} `json:"name_regex_keep"`
		NextRunAt       time.Time   `json:"next_run_at"`
	} `json:"container_expiration_policy"`
	CreatedAt      time.Time `json:"created_at"`
	LastActivityAt time.Time `json:"last_activity_at"`
	CreatorId      int       `json:"creator_id"`
	Namespace      struct {
		Id        int    `json:"id"`
		Name      string `json:"name"`
		Path      string `json:"path"`
		Kind      string `json:"kind"`
		FullPath  string `json:"full_path"`
		AvatarUrl string `json:"avatar_url"`
		WebUrl    string `json:"web_url"`
	} `json:"namespace"`
	ImportStatus string      `json:"import_status"`
	ImportError  interface{} `json:"import_error"`
	Permissions  struct {
		ProjectAccess struct {
			AccessLevel       int `json:"access_level"`
			NotificationLevel int `json:"notification_level"`
		} `json:"project_access"`
		GroupAccess struct {
			AccessLevel       int `json:"access_level"`
			NotificationLevel int `json:"notification_level"`
		} `json:"group_access"`
	} `json:"permissions"`
	Archived   bool   `json:"archived"`
	AvatarUrl  string `json:"avatar_url"`
	LicenseUrl string `json:"license_url"`
	License    struct {
		Key       string `json:"key"`
		Name      string `json:"name"`
		Nickname  string `json:"nickname"`
		HtmlUrl   string `json:"html_url"`
		SourceUrl string `json:"source_url"`
	} `json:"license"`
	SharedRunnersEnabled       bool   `json:"shared_runners_enabled"`
	ForksCount                 int    `json:"forks_count"`
	StarCount                  int    `json:"star_count"`
	RunnersToken               string `json:"runners_token"`
	CiDefaultGitDepth          int    `json:"ci_default_git_depth"`
	CiForwardDeploymentEnabled bool   `json:"ci_forward_deployment_enabled"`
	PublicJobs                 bool   `json:"public_jobs"`
	SharedWithGroups           []struct {
		GroupId          int    `json:"group_id"`
		GroupName        string `json:"group_name"`
		GroupFullPath    string `json:"group_full_path"`
		GroupAccessLevel int    `json:"group_access_level"`
	} `json:"shared_with_groups"`
	RepositoryStorage                         string      `json:"repository_storage"`
	OnlyAllowMergeIfPipelineSucceeds          bool        `json:"only_allow_merge_if_pipeline_succeeds"`
	AllowMergeOnSkippedPipeline               bool        `json:"allow_merge_on_skipped_pipeline"`
	RestrictUserDefinedVariables              bool        `json:"restrict_user_defined_variables"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool        `json:"only_allow_merge_if_all_discussions_are_resolved"`
	RemoveSourceBranchAfterMerge              bool        `json:"remove_source_branch_after_merge"`
	PrintingMergeRequestsLinkEnabled          bool        `json:"printing_merge_requests_link_enabled"`
	RequestAccessEnabled                      bool        `json:"request_access_enabled"`
	MergeMethod                               string      `json:"merge_method"`
	SquashOption                              string      `json:"squash_option"`
	AutoDevopsEnabled                         bool        `json:"auto_devops_enabled"`
	AutoDevopsDeployStrategy                  string      `json:"auto_devops_deploy_strategy"`
	ApprovalsBeforeMerge                      int         `json:"approvals_before_merge"`
	Mirror                                    bool        `json:"mirror"`
	MirrorUserId                              int         `json:"mirror_user_id"`
	MirrorTriggerBuilds                       bool        `json:"mirror_trigger_builds"`
	OnlyMirrorProtectedBranches               bool        `json:"only_mirror_protected_branches"`
	MirrorOverwritesDivergedBranches          bool        `json:"mirror_overwrites_diverged_branches"`
	ExternalAuthorizationClassificationLabel  interface{} `json:"external_authorization_classification_label"`
	PackagesEnabled                           bool        `json:"packages_enabled"`
	ServiceDeskEnabled                        bool        `json:"service_desk_enabled"`
	ServiceDeskAddress                        interface{} `json:"service_desk_address"`
	AutocloseReferencedIssues                 bool        `json:"autoclose_referenced_issues"`
	SuggestionCommitMessage                   interface{} `json:"suggestion_commit_message"`
	MergeCommitTemplate                       interface{} `json:"merge_commit_template"`
	SquashCommitTemplate                      interface{} `json:"squash_commit_template"`
	MarkedForDeletionAt                       string      `json:"marked_for_deletion_at"`
	MarkedForDeletionOn                       string      `json:"marked_for_deletion_on"`
	ComplianceFrameworks                      []string    `json:"compliance_frameworks"`
	Statistics                                struct {
		CommitCount           int `json:"commit_count"`
		StorageSize           int `json:"storage_size"`
		RepositorySize        int `json:"repository_size"`
		WikiSize              int `json:"wiki_size"`
		LfsObjectsSize        int `json:"lfs_objects_size"`
		JobArtifactsSize      int `json:"job_artifacts_size"`
		PipelineArtifactsSize int `json:"pipeline_artifacts_size"`
		PackagesSize          int `json:"packages_size"`
		SnippetsSize          int `json:"snippets_size"`
		UploadsSize           int `json:"uploads_size"`
	} `json:"statistics"`
	ContainerRegistryImagePrefix string `json:"container_registry_image_prefix"`
	Links                        struct {
		Self          string `json:"self"`
		Issues        string `json:"issues"`
		MergeRequests string `json:"merge_requests"`
		RepoBranches  string `json:"repo_branches"`
		Labels        string `json:"labels"`
		Events        string `json:"events"`
		Members       string `json:"members"`
	} `json:"_links"`
}
