package cloudfile

import (
	"context"

	"github.com/suyuan32/simple-admin-file/internal/svc"
	"github.com/suyuan32/simple-admin-file/internal/types"
	"github.com/suyuan32/simple-admin-file/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCloudFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCloudFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCloudFileLogic {
	return &UpdateCloudFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCloudFileLogic) UpdateCloudFile(req *types.CloudFileInfo) (*types.BaseMsgResp, error) {
	err := l.svcCtx.DB.CloudFile.UpdateOneID(uuidx.ParseUUIDString(*req.Id)).
		SetNotNilState(req.State).
		SetNotNilName(req.Name).
		SetNotNilURL(req.Url).
		SetNotNilSize(req.Size).
		SetNotNilFileType(req.FileType).
		SetNotNilUserID(req.UserId).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}
