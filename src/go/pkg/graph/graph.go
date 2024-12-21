package graph

import "iter"

type INode interface {
	Edges() iter.Seq[INode]
}
