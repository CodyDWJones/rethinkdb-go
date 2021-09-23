// Code generated by gen_tests.py and process_polyglot.py.
// Do not edit this file directly.
// The template for this file is located at:
// ../template.go.tpl
package reql_tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"gopkg.in/rethinkdb/rethinkdb-go.v6/internal/compare"
)

// These tests are aimed at &&, ||, and !
func TestMathLogicLogicSuite(t *testing.T) {
	suite.Run(t, new(MathLogicLogicSuite))
}

type MathLogicLogicSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *MathLogicLogicSuite) SetupTest() {
	suite.T().Log("Setting up MathLogicLogicSuite")
	// Use imports to prevent errors
	_ = time.Time{}
	_ = compare.AnythingIsFine

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("test").Exec(suite.session)
	err = r.DBCreate("test").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Wait().Exec(suite.session)
	suite.Require().NoError(err)

}

func (suite *MathLogicLogicSuite) TearDownSuite() {
	suite.T().Log("Tearing down MathLogicLogicSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *MathLogicLogicSuite) TestCases() {
	suite.T().Log("Running MathLogicLogicSuite: These tests are aimed at &&, ||, and !")

	{
		// math_logic/logic.yaml line #8
		/* true */
		var expected_ bool = true
		/* r.expr(true) & true */

		suite.T().Log("About to run line #8: r.Expr(true).And(true)")

		runAndAssert(suite.Suite, expected_, r.Expr(true).And(true), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #8")
	}

	{
		// math_logic/logic.yaml line #9
		/* true */
		var expected_ bool = true
		/* true & r.expr(true) */

		suite.T().Log("About to run line #9: r.And(true, r.Expr(true))")

		runAndAssert(suite.Suite, expected_, r.And(true, r.Expr(true)), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #9")
	}

	{
		// math_logic/logic.yaml line #10
		/* true */
		var expected_ bool = true
		/* r.and_(true,true) */

		suite.T().Log("About to run line #10: r.And(true, true)")

		runAndAssert(suite.Suite, expected_, r.And(true, true), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #10")
	}

	{
		// math_logic/logic.yaml line #11
		/* true */
		var expected_ bool = true
		/* r.expr(true).and_(true) */

		suite.T().Log("About to run line #11: r.Expr(true).And(true)")

		runAndAssert(suite.Suite, expected_, r.Expr(true).And(true), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #11")
	}

	{
		// math_logic/logic.yaml line #22
		/* false */
		var expected_ bool = false
		/* r.expr(true) & false */

		suite.T().Log("About to run line #22: r.Expr(true).And(false)")

		runAndAssert(suite.Suite, expected_, r.Expr(true).And(false), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #22")
	}

	{
		// math_logic/logic.yaml line #23
		/* false */
		var expected_ bool = false
		/* r.expr(false) & false */

		suite.T().Log("About to run line #23: r.Expr(false).And(false)")

		runAndAssert(suite.Suite, expected_, r.Expr(false).And(false), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #23")
	}

	{
		// math_logic/logic.yaml line #24
		/* false */
		var expected_ bool = false
		/* true & r.expr(false) */

		suite.T().Log("About to run line #24: r.And(true, r.Expr(false))")

		runAndAssert(suite.Suite, expected_, r.And(true, r.Expr(false)), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #24")
	}

	{
		// math_logic/logic.yaml line #25
		/* false */
		var expected_ bool = false
		/* false & r.expr(false) */

		suite.T().Log("About to run line #25: r.And(false, r.Expr(false))")

		runAndAssert(suite.Suite, expected_, r.And(false, r.Expr(false)), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #25")
	}

	{
		// math_logic/logic.yaml line #26
		/* false */
		var expected_ bool = false
		/* r.and_(true,false) */

		suite.T().Log("About to run line #26: r.And(true, false)")

		runAndAssert(suite.Suite, expected_, r.And(true, false), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #26")
	}

	{
		// math_logic/logic.yaml line #27
		/* false */
		var expected_ bool = false
		/* r.and_(false,false) */

		suite.T().Log("About to run line #27: r.And(false, false)")

		runAndAssert(suite.Suite, expected_, r.And(false, false), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #27")
	}

	{
		// math_logic/logic.yaml line #28
		/* false */
		var expected_ bool = false
		/* r.expr(true).and_(false) */

		suite.T().Log("About to run line #28: r.Expr(true).And(false)")

		runAndAssert(suite.Suite, expected_, r.Expr(true).And(false), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #28")
	}

	{
		// math_logic/logic.yaml line #29
		/* false */
		var expected_ bool = false
		/* r.expr(false).and_(false) */

		suite.T().Log("About to run line #29: r.Expr(false).And(false)")

		runAndAssert(suite.Suite, expected_, r.Expr(false).And(false), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #29")
	}

	{
		// math_logic/logic.yaml line #48
		/* true */
		var expected_ bool = true
		/* r.expr(true) | true */

		suite.T().Log("About to run line #48: r.Expr(true).Or(true)")

		runAndAssert(suite.Suite, expected_, r.Expr(true).Or(true), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #48")
	}

	{
		// math_logic/logic.yaml line #49
		/* true */
		var expected_ bool = true
		/* r.expr(true) | false */

		suite.T().Log("About to run line #49: r.Expr(true).Or(false)")

		runAndAssert(suite.Suite, expected_, r.Expr(true).Or(false), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #49")
	}

	{
		// math_logic/logic.yaml line #50
		/* true */
		var expected_ bool = true
		/* true | r.expr(true) */

		suite.T().Log("About to run line #50: r.Or(true, r.Expr(true))")

		runAndAssert(suite.Suite, expected_, r.Or(true, r.Expr(true)), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #50")
	}

	{
		// math_logic/logic.yaml line #51
		/* true */
		var expected_ bool = true
		/* true | r.expr(false) */

		suite.T().Log("About to run line #51: r.Or(true, r.Expr(false))")

		runAndAssert(suite.Suite, expected_, r.Or(true, r.Expr(false)), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #51")
	}

	{
		// math_logic/logic.yaml line #52
		/* true */
		var expected_ bool = true
		/* r.or_(true,true) */

		suite.T().Log("About to run line #52: r.Or(true, true)")

		runAndAssert(suite.Suite, expected_, r.Or(true, true), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #52")
	}

	{
		// math_logic/logic.yaml line #53
		/* true */
		var expected_ bool = true
		/* r.or_(true,false) */

		suite.T().Log("About to run line #53: r.Or(true, false)")

		runAndAssert(suite.Suite, expected_, r.Or(true, false), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #53")
	}

	{
		// math_logic/logic.yaml line #54
		/* true */
		var expected_ bool = true
		/* r.expr(true).or_(true) */

		suite.T().Log("About to run line #54: r.Expr(true).Or(true)")

		runAndAssert(suite.Suite, expected_, r.Expr(true).Or(true), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #54")
	}

	{
		// math_logic/logic.yaml line #55
		/* true */
		var expected_ bool = true
		/* r.expr(true).or_(false) */

		suite.T().Log("About to run line #55: r.Expr(true).Or(false)")

		runAndAssert(suite.Suite, expected_, r.Expr(true).Or(false), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #55")
	}

	{
		// math_logic/logic.yaml line #72
		/* false */
		var expected_ bool = false
		/* r.expr(false) | false */

		suite.T().Log("About to run line #72: r.Expr(false).Or(false)")

		runAndAssert(suite.Suite, expected_, r.Expr(false).Or(false), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #72")
	}

	{
		// math_logic/logic.yaml line #73
		/* false */
		var expected_ bool = false
		/* false | r.expr(false) */

		suite.T().Log("About to run line #73: r.Or(false, r.Expr(false))")

		runAndAssert(suite.Suite, expected_, r.Or(false, r.Expr(false)), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #73")
	}

	{
		// math_logic/logic.yaml line #74
		/* false */
		var expected_ bool = false
		/* r.and_(false,false) */

		suite.T().Log("About to run line #74: r.And(false, false)")

		runAndAssert(suite.Suite, expected_, r.And(false, false), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #74")
	}

	{
		// math_logic/logic.yaml line #75
		/* false */
		var expected_ bool = false
		/* r.expr(false).and_(false) */

		suite.T().Log("About to run line #75: r.Expr(false).And(false)")

		runAndAssert(suite.Suite, expected_, r.Expr(false).And(false), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #75")
	}

	{
		// math_logic/logic.yaml line #88
		/* false */
		var expected_ bool = false
		/* ~r.expr(True) */

		suite.T().Log("About to run line #88: r.Expr(true).Not()")

		runAndAssert(suite.Suite, expected_, r.Expr(true).Not(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #88")
	}

	{
		// math_logic/logic.yaml line #89
		/* false */
		var expected_ bool = false
		/* r.not_(True) */

		suite.T().Log("About to run line #89: r.Not(true)")

		runAndAssert(suite.Suite, expected_, r.Not(true), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #89")
	}

	{
		// math_logic/logic.yaml line #93
		/* true */
		var expected_ bool = true
		/* ~r.expr(False) */

		suite.T().Log("About to run line #93: r.Expr(false).Not()")

		runAndAssert(suite.Suite, expected_, r.Expr(false).Not(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #93")
	}

	{
		// math_logic/logic.yaml line #94
		/* true */
		var expected_ bool = true
		/* r.not_(False) */

		suite.T().Log("About to run line #94: r.Not(false)")

		runAndAssert(suite.Suite, expected_, r.Not(false), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #94")
	}

	{
		// math_logic/logic.yaml line #97
		/* false */
		var expected_ bool = false
		/* r.expr(True).not_() */

		suite.T().Log("About to run line #97: r.Expr(true).Not()")

		runAndAssert(suite.Suite, expected_, r.Expr(true).Not(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #97")
	}

	{
		// math_logic/logic.yaml line #100
		/* true */
		var expected_ bool = true
		/* r.expr(False).not_() */

		suite.T().Log("About to run line #100: r.Expr(false).Not()")

		runAndAssert(suite.Suite, expected_, r.Expr(false).Not(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #100")
	}

	{
		// math_logic/logic.yaml line #107
		/* true */
		var expected_ bool = true
		/* ~r.and_(True, True) == r.or_(~r.expr(True), ~r.expr(True)) */

		suite.T().Log("About to run line #107: r.And(true, true).Not().Eq(r.Or(r.Expr(true).Not(), r.Expr(true).Not()))")

		runAndAssert(suite.Suite, expected_, r.And(true, true).Not().Eq(r.Or(r.Expr(true).Not(), r.Expr(true).Not())), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #107")
	}

	{
		// math_logic/logic.yaml line #108
		/* true */
		var expected_ bool = true
		/* ~r.and_(True, False) == r.or_(~r.expr(True), ~r.expr(False)) */

		suite.T().Log("About to run line #108: r.And(true, false).Not().Eq(r.Or(r.Expr(true).Not(), r.Expr(false).Not()))")

		runAndAssert(suite.Suite, expected_, r.And(true, false).Not().Eq(r.Or(r.Expr(true).Not(), r.Expr(false).Not())), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #108")
	}

	{
		// math_logic/logic.yaml line #109
		/* true */
		var expected_ bool = true
		/* ~r.and_(False, False) == r.or_(~r.expr(False), ~r.expr(False)) */

		suite.T().Log("About to run line #109: r.And(false, false).Not().Eq(r.Or(r.Expr(false).Not(), r.Expr(false).Not()))")

		runAndAssert(suite.Suite, expected_, r.And(false, false).Not().Eq(r.Or(r.Expr(false).Not(), r.Expr(false).Not())), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #109")
	}

	{
		// math_logic/logic.yaml line #110
		/* true */
		var expected_ bool = true
		/* ~r.and_(False, True) == r.or_(~r.expr(False), ~r.expr(True)) */

		suite.T().Log("About to run line #110: r.And(false, true).Not().Eq(r.Or(r.Expr(false).Not(), r.Expr(true).Not()))")

		runAndAssert(suite.Suite, expected_, r.And(false, true).Not().Eq(r.Or(r.Expr(false).Not(), r.Expr(true).Not())), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #110")
	}

	{
		// math_logic/logic.yaml line #120
		/* true */
		var expected_ bool = true
		/* r.and_(True, True, True, True, True) */

		suite.T().Log("About to run line #120: r.And(true, true, true, true, true)")

		runAndAssert(suite.Suite, expected_, r.And(true, true, true, true, true), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #120")
	}

	{
		// math_logic/logic.yaml line #123
		/* false */
		var expected_ bool = false
		/* r.and_(True, True, True, False, True) */

		suite.T().Log("About to run line #123: r.And(true, true, true, false, true)")

		runAndAssert(suite.Suite, expected_, r.And(true, true, true, false, true), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #123")
	}

	{
		// math_logic/logic.yaml line #126
		/* false */
		var expected_ bool = false
		/* r.and_(True, False, True, False, True) */

		suite.T().Log("About to run line #126: r.And(true, false, true, false, true)")

		runAndAssert(suite.Suite, expected_, r.And(true, false, true, false, true), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #126")
	}

	{
		// math_logic/logic.yaml line #129
		/* false */
		var expected_ bool = false
		/* r.or_(False, False, False, False, False) */

		suite.T().Log("About to run line #129: r.Or(false, false, false, false, false)")

		runAndAssert(suite.Suite, expected_, r.Or(false, false, false, false, false), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #129")
	}

	{
		// math_logic/logic.yaml line #132
		/* true */
		var expected_ bool = true
		/* r.or_(False, False, False, True, False) */

		suite.T().Log("About to run line #132: r.Or(false, false, false, true, false)")

		runAndAssert(suite.Suite, expected_, r.Or(false, false, false, true, false), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #132")
	}

	{
		// math_logic/logic.yaml line #135
		/* true */
		var expected_ bool = true
		/* r.or_(False, True, False, True, False) */

		suite.T().Log("About to run line #135: r.Or(false, true, false, true, false)")

		runAndAssert(suite.Suite, expected_, r.Or(false, true, false, true, false), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #135")
	}

	{
		// math_logic/logic.yaml line #140
		/* err("ReqlQueryLogicError", "Cannot perform bracket on a non-object non-sequence `\"a\"`.", []) */
		var expected_ Err = err("ReqlQueryLogicError", "Cannot perform bracket on a non-object non-sequence `\"a\"`.")
		/* r.expr(r.expr('a')['b']).default(2) */

		suite.T().Log("About to run line #140: r.Expr(r.Expr('a').AtIndex('b')).Default(2)")

		runAndAssert(suite.Suite, expected_, r.Expr(r.Expr("a").AtIndex("b")).Default(2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #140")
	}

	{
		// math_logic/logic.yaml line #145
		/* False */
		var expected_ bool = false
		/* r.expr(r.and_(True, False) == r.or_(False, True)) */

		suite.T().Log("About to run line #145: r.Expr(r.And(true, false).Eq(r.Or(false, true)))")

		runAndAssert(suite.Suite, expected_, r.Expr(r.And(true, false).Eq(r.Or(false, true))), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #145")
	}

	{
		// math_logic/logic.yaml line #151
		/* False */
		var expected_ bool = false
		/* r.expr(r.and_(True, False) >= r.or_(False, True)) */

		suite.T().Log("About to run line #151: r.Expr(r.And(true, false).Ge(r.Or(false, true)))")

		runAndAssert(suite.Suite, expected_, r.Expr(r.And(true, false).Ge(r.Or(false, true))), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #151")
	}

	{
		// math_logic/logic.yaml line #156
		/* true */
		var expected_ bool = true
		/* r.expr(1) & True */

		suite.T().Log("About to run line #156: r.Expr(1).And(true)")

		runAndAssert(suite.Suite, expected_, r.Expr(1).And(true), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #156")
	}

	{
		// math_logic/logic.yaml line #160
		/* ("str") */
		var expected_ string = "str"
		/* r.expr(False) | 'str' */

		suite.T().Log("About to run line #160: r.Expr(false).Or('str')")

		runAndAssert(suite.Suite, expected_, r.Expr(false).Or("str"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #160")
	}

	{
		// math_logic/logic.yaml line #164
		/* false */
		var expected_ bool = false
		/* ~r.expr(1) */

		suite.T().Log("About to run line #164: r.Expr(1).Not()")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Not(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #164")
	}

	{
		// math_logic/logic.yaml line #168
		/* true */
		var expected_ bool = true
		/* ~r.expr(null) */

		suite.T().Log("About to run line #168: r.Expr(nil).Not()")

		runAndAssert(suite.Suite, expected_, r.Expr(nil).Not(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #168")
	}
}
