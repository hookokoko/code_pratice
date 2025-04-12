package rank_list

// RankMgr 排行榜管理器
type RankMgr interface {
	// 结构体需要有榜单汇总，比如一个叫做rankObjHash的struct，作为保存用户id和用户信息的map

	// GetRankObj 获取某个榜单对象
	GetRankObj()
	// UpdateUserInfo 更新用户信息
	UpdateUserInfo()
	// UpdateRank 刷新排行榜
	UpdateRank()
}

type UserInfoMgr interface {
	// 用户信息管理
	// 这里会调用RankMgr用户排行榜管理器
}

type RankInstance struct {
	Head       Node
	Tail       Node
	RankLength int64 // 排行榜长度
	Data       any   // 用户信息汇总
}

type Node struct {
}

// 更新用户排名
func (ri *RankInstance) UpdateRankValue() {}

// 获取整个榜单的数据
func (ri *RankInstance) GetData() {}

// 根据排名范围获取榜单数据
func (ri *RankInstance) GetRangeData() {}
