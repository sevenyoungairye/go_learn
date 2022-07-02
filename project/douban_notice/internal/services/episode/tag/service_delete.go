package tag

import (
	"top.lel.dn/main/internal/repository/mongodb"
)

func (s *service) DeleteById(id string) error {
	info := mongodb.TagInfo{Id: id}
	return info.DeleteById(s.mongoCtx)
}
