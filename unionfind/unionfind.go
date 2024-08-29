package unionfind

type Member struct {
	parent    *Member
	UnionSize int
	Val       interface{}
}

func NewMember(val interface{}) *Member {
	newMember := Member{
		UnionSize: 1,
		Val:       val,
	}
	newMember.parent = &newMember
	return &newMember
}

func Union(x, y *Member) *Member {
	xParent := Find(x)
	yParent := Find(y)
	if xParent == yParent {
		return xParent
	}
	yParent.parent = xParent
	xParent.UnionSize = xParent.UnionSize + yParent.UnionSize
	return xParent
}

func Find(x *Member) *Member {
	cur := x
	for cur != cur.parent {
		cur = cur.parent
	}
	return cur
}
