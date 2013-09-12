package rethinkgo

import (
	"fmt"
	test "launchpad.net/gocheck"
)

func (s *RethinkSuite) TestControlExecList(c *test.C) {
	tree := Expr(List{1, 2, 3, 4, 5, 6, List{7.1, 7.2, 7.3}}).Add(1, 2)
	// fmt.Println(Expr(List{1, 2, 3, 4, 5, 6, List{7.1, 7.2, 7.3}}))
	// fmt.Printf("%T, %+v", tree, tree)
	fmt.Println(tree.compose())
}
