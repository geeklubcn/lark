package gitlab

import (
	"encoding/json"
	"strconv"
)

var FakerGitlab = &fakerGitlab{}

type fakerGitlab struct{}

func (f fakerGitlab) Issues() (map[string]*IssueResult, error) {
	data := `
[
    {
        "id": 30114215,
        "iid": 1,
        "project_id": 16607076,
        "title": "提供issue变更钉钉消息通知",
        "description": "# 背景\n\n通过钉钉消息传递issue变更\n\n## 为什么不通过钉钉自带的gitlab机器人？\n\n1. 发送的消息过多，会忽略有意义的信息\n2. 不能 @ 到相关的人，达到提醒的作用",
        "state": "opened",
        "created_at": "2020-01-30T12:55:54.836Z",
        "updated_at": "2022-04-04T13:26:32.531Z",
        "closed_at": null,
        "closed_by": null,
        "labels": [
            "P0",
            "feature"
        ],
        "milestone": {
            "id": 1139988,
            "iid": 1,
            "project_id": 16607076,
            "title": "2020#02",
            "description": "",
            "state": "active",
            "created_at": "2020-01-30T12:27:29.752Z",
            "updated_at": "2020-01-30T12:27:29.752Z",
            "due_date": "2020-02-29",
            "start_date": "2020-02-01",
            "expired": true,
            "web_url": "https://gitlab.com/wangyuheng77/integration/-/milestones/1"
        },
        "assignees": [
            {
                "id": 5288824,
                "username": "wangyuheng",
                "name": "wangyuheng",
                "state": "active",
                "avatar_url": "https://secure.gravatar.com/avatar/fd8467612d6b157cb036c156d7cb7121?s=80\u0026d=identicon",
                "web_url": "https://gitlab.com/wangyuheng"
            }
        ],
        "author": {
            "id": 5288824,
            "username": "wangyuheng",
            "name": "wangyuheng",
            "state": "active",
            "avatar_url": "https://secure.gravatar.com/avatar/fd8467612d6b157cb036c156d7cb7121?s=80\u0026d=identicon",
            "web_url": "https://gitlab.com/wangyuheng"
        },
        "type": "ISSUE",
        "assignee": {
            "id": 5288824,
            "username": "wangyuheng",
            "name": "wangyuheng",
            "state": "active",
            "avatar_url": "https://secure.gravatar.com/avatar/fd8467612d6b157cb036c156d7cb7121?s=80\u0026d=identicon",
            "web_url": "https://gitlab.com/wangyuheng"
        },
        "user_notes_count": 0,
        "merge_requests_count": 0,
        "upvotes": 0,
        "downvotes": 0,
        "due_date": "2022-04-06",
        "confidential": false,
        "discussion_locked": null,
        "issue_type": "issue",
        "web_url": "https://gitlab.com/wangyuheng77/integration/-/issues/1",
        "time_stats": {
            "time_estimate": 0,
            "total_time_spent": 0,
            "human_time_estimate": null,
            "human_total_time_spent": null
        },
        "task_completion_status": {
            "count": 0,
            "completed_count": 0
        },
        "weight": 1,
        "blocking_issues_count": 0,
        "has_tasks": false,
        "_links": {
            "self": "https://gitlab.com/api/v4/projects/16607076/issues/1",
            "notes": "https://gitlab.com/api/v4/projects/16607076/issues/1/notes",
            "award_emoji": "https://gitlab.com/api/v4/projects/16607076/issues/1/award_emoji",
            "project": "https://gitlab.com/api/v4/projects/16607076"
        },
        "references": {
            "short": "#1",
            "relative": "#1",
            "full": "wangyuheng77/integration#1"
        },
        "severity": "UNKNOWN",
        "moved_to_id": null,
        "service_desk_reply_to": null,
        "health_status": "needs_attention"
    }
]
`
	var issues []*IssueResult
	_ = json.Unmarshal([]byte(data), &issues)

	res := make(map[string]*IssueResult, 0)
	for _, it := range issues {
		res[strconv.Itoa(it.Id)] = it
	}
	return res, nil
}

func (f fakerGitlab) Project(id string) (*ProjectResult, error) {
	data := `
{
    "id": 16607076,
    "description": "manage feature & bug by issue",
    "name": "integration",
    "name_with_namespace": "wangyuheng / integration",
    "path": "integration",
    "path_with_namespace": "wangyuheng77/integration",
    "created_at": "2020-01-30T12:25:00.751Z",
    "default_branch": "master",
    "tag_list": [],
    "topics": [],
    "ssh_url_to_repo": "git@gitlab.com:wangyuheng77/integration.git",
    "http_url_to_repo": "https://gitlab.com/wangyuheng77/integration.git",
    "web_url": "https://gitlab.com/wangyuheng77/integration",
    "readme_url": "https://gitlab.com/wangyuheng77/integration/-/blob/master/README.md",
    "avatar_url": null,
    "forks_count": 0,
    "star_count": 0,
    "last_activity_at": "2020-01-30T12:25:00.751Z",
    "namespace": {
        "id": 7044374,
        "name": "wangyuheng",
        "path": "wangyuheng77",
        "kind": "group",
        "full_path": "wangyuheng77",
        "parent_id": null,
        "avatar_url": null,
        "web_url": "https://gitlab.com/groups/wangyuheng77"
    },
    "container_registry_image_prefix": "registry.gitlab.com/wangyuheng77/integration",
    "_links": {
        "self": "https://gitlab.com/api/v4/projects/16607076",
        "issues": "https://gitlab.com/api/v4/projects/16607076/issues",
        "merge_requests": "https://gitlab.com/api/v4/projects/16607076/merge_requests",
        "repo_branches": "https://gitlab.com/api/v4/projects/16607076/repository/branches",
        "labels": "https://gitlab.com/api/v4/projects/16607076/labels",
        "events": "https://gitlab.com/api/v4/projects/16607076/events",
        "members": "https://gitlab.com/api/v4/projects/16607076/members"
    },
    "packages_enabled": true,
    "empty_repo": false,
    "archived": false,
    "visibility": "public",
    "resolve_outdated_diff_discussions": false,
    "container_expiration_policy": {
        "cadence": "7d",
        "enabled": false,
        "keep_n": null,
        "older_than": null,
        "name_regex": null,
        "name_regex_keep": null,
        "next_run_at": "2020-02-06T12:25:00.780Z"
    },
    "issues_enabled": true,
    "merge_requests_enabled": true,
    "wiki_enabled": true,
    "jobs_enabled": true,
    "snippets_enabled": true,
    "container_registry_enabled": true,
    "service_desk_enabled": true,
    "service_desk_address": "contact-project+wangyuheng77-integration-16607076-issue-@incoming.gitlab.com",
    "can_create_merge_request_in": true,
    "issues_access_level": "enabled",
    "repository_access_level": "enabled",
    "merge_requests_access_level": "enabled",
    "forking_access_level": "enabled",
    "wiki_access_level": "enabled",
    "builds_access_level": "enabled",
    "snippets_access_level": "enabled",
    "pages_access_level": "enabled",
    "operations_access_level": "enabled",
    "analytics_access_level": "enabled",
    "container_registry_access_level": "enabled",
    "security_and_compliance_access_level": "private",
    "emails_disabled": null,
    "shared_runners_enabled": true,
    "lfs_enabled": true,
    "creator_id": 5288824,
    "import_url": null,
    "import_type": null,
    "import_status": "none",
    "import_error": null,
    "open_issues_count": 1,
    "runners_token": "GR1348941LyDDzpJNxRcJcEKYFa-H",
    "ci_default_git_depth": 50,
    "ci_forward_deployment_enabled": null,
    "ci_job_token_scope_enabled": false,
    "public_jobs": true,
    "build_git_strategy": "fetch",
    "build_timeout": 3600,
    "auto_cancel_pending_pipelines": "enabled",
    "build_coverage_regex": null,
    "ci_config_path": null,
    "shared_with_groups": [],
    "only_allow_merge_if_pipeline_succeeds": false,
    "allow_merge_on_skipped_pipeline": null,
    "restrict_user_defined_variables": false,
    "request_access_enabled": true,
    "only_allow_merge_if_all_discussions_are_resolved": false,
    "remove_source_branch_after_merge": true,
    "printing_merge_request_link_enabled": true,
    "merge_method": "merge",
    "squash_option": "default_off",
    "suggestion_commit_message": null,
    "merge_commit_template": null,
    "squash_commit_template": null,
    "auto_devops_enabled": false,
    "auto_devops_deploy_strategy": "continuous",
    "autoclose_referenced_issues": true,
    "keep_latest_artifact": true,
    "runner_token_expiration_interval": null,
    "approvals_before_merge": 0,
    "mirror": false,
    "external_authorization_classification_label": "",
    "marked_for_deletion_at": null,
    "marked_for_deletion_on": null,
    "requirements_enabled": true,
    "requirements_access_level": "enabled",
    "security_and_compliance_enabled": true,
    "compliance_frameworks": [],
    "issues_template": null,
    "merge_requests_template": null,
    "merge_pipelines_enabled": false,
    "merge_trains_enabled": false,
    "permissions": {
        "project_access": null,
        "group_access": {
            "access_level": 50,
            "notification_level": 3
        }
    }
}
`
	var res ProjectResult
	_ = json.Unmarshal([]byte(data), &res)
	return &res, nil
}
