package core

import (
	"context"
)

type Api interface {
	PackageList(ctx context.Context) (*PackageListResponse, error)
	PackageGet(ctx context.Context, id string) (*PackageGetResponse, error)
	PackageFind(ctx context.Context, name string) (*Package, error)
	SystemInfo(ctx context.Context) (*SystemInfoResponse, error)
	PackageServerList(ctx context.Context, req PackageServerListRequest) (*PackageServerListResponse, error)
	PackageInstall(ctx context.Context, req PackageInstallRequest) (*PackageInstallResponse, error)
	PackageInstallCompound(ctx context.Context, req PackageInstallCompoundRequest) error
	PackageInstallStatus(ctx context.Context, req PackageInstallStatusRequest) (*PackageInstallStatusResponse, error)
	PackageInstallDelete(ctx context.Context, req PackageInstallDeleteRequest) error
	PackageUninstall(ctx context.Context, req PackageUninstallRequest) (*PackageUninstallResponse, error)
	PackageUninstallCompound(ctx context.Context, name string) error
	ContentLength(ctx context.Context, url string) (int64, error)
	PackageFeedList(ctx context.Context) (*PackageFeedListResponse, error)
	PackageFeedAdd(ctx context.Context, req PackageFeedAddRequest) error
	PackageFeedDelete(ctx context.Context, req PackageFeedDeleteRequest) error
	PackageSettingGet(ctx context.Context, req PackageSettingGetRequest) (*PackageSettingGetResponse, error)
	ShareList(ctx context.Context) (*ShareListResponse, error)
	ShareGet(ctx context.Context, name string) (*ShareGetResponse, error)
	ShareGetByID(ctx context.Context, id string) (*Share, error)
	ShareCreate(ctx context.Context, share ShareInfo) error
	ShareDelete(ctx context.Context, name string) error
	VolumeList(ctx context.Context) (*VolumeListResponse, error)

	TaskList(ctx context.Context, req ListTaskRequest) (*ListTaskResponse, error)
	TaskCreate(ctx context.Context, req TaskRequest) (*TaskResult, error)
	RootTaskCreate(ctx context.Context, req TaskRequest) (*TaskResult, error)
	TaskUpdate(ctx context.Context, req TaskRequest) (*TaskResult, error)
	RootTaskUpdate(ctx context.Context, req TaskRequest) (*TaskResult, error)
	TaskDelete(ctx context.Context, ids ...int64) error
	TaskRun(ctx context.Context, ids ...int64) error
	TaskGet(ctx context.Context, id int64) (*TaskResult, error)
	TaskFind(ctx context.Context, name string) (*TaskResult, error)

	PasswordConfirm(ctx context.Context, password string) (*PasswordConfirmResponse, error)
}
