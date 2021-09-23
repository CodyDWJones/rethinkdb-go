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

// Tests the RQL `map` function
func TestTransformMapSuite(t *testing.T) {
	suite.Run(t, new(TransformMapSuite))
}

type TransformMapSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *TransformMapSuite) SetupTest() {
	suite.T().Log("Setting up TransformMapSuite")
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

func (suite *TransformMapSuite) TearDownSuite() {
	suite.T().Log("Tearing down TransformMapSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *TransformMapSuite) TestCases() {
	suite.T().Log("Running TransformMapSuite: Tests the RQL `map` function")

	{
		// transform/map.yaml line #5
		/* 'STREAM' */
		var expected_ string = "STREAM"
		/* r.range().map(r.range(), lambda x, y:(x, y)).type_of() */

		suite.T().Log("About to run line #5: r.Range().Map(r.Range(), func(x r.Term, y r.Term) interface{} { return []interface{}{x, y}}).TypeOf()")

		runAndAssert(suite.Suite, expected_, r.Range().Map(r.Range(), func(x r.Term, y r.Term) interface{} { return []interface{}{x, y} }).TypeOf(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #5")
	}

	{
		// transform/map.yaml line #10
		/* 'STREAM' */
		var expected_ string = "STREAM"
		/* r.range().map(r.expr([]), lambda x, y:(x, y)).type_of() */

		suite.T().Log("About to run line #10: r.Range().Map(r.Expr([]interface{}{}), func(x r.Term, y r.Term) interface{} { return []interface{}{x, y}}).TypeOf()")

		runAndAssert(suite.Suite, expected_, r.Range().Map(r.Expr([]interface{}{}), func(x r.Term, y r.Term) interface{} { return []interface{}{x, y} }).TypeOf(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #10")
	}

	{
		// transform/map.yaml line #15
		/* 'ARRAY' */
		var expected_ string = "ARRAY"
		/* r.expr([]).map(r.expr([]), lambda x, y:(x, y)).type_of() */

		suite.T().Log("About to run line #15: r.Expr([]interface{}{}).Map(r.Expr([]interface{}{}), func(x r.Term, y r.Term) interface{} { return []interface{}{x, y}}).TypeOf()")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{}).Map(r.Expr([]interface{}{}), func(x r.Term, y r.Term) interface{} { return []interface{}{x, y} }).TypeOf(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #15")
	}

	{
		// transform/map.yaml line #21
		/* [0, 0, 0] */
		var expected_ []interface{} = []interface{}{0, 0, 0}
		/* r.range(3).map(lambda:0) */

		suite.T().Log("About to run line #21: r.Range(3).Map(func() interface{} { return 0})")

		runAndAssert(suite.Suite, expected_, r.Range(3).Map(func() interface{} { return 0 }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #21")
	}

	{
		// transform/map.yaml line #26
		/* [0, 0, 0] */
		var expected_ []interface{} = []interface{}{0, 0, 0}
		/* r.range(3).map(r.range(4), lambda x,y:0) */

		suite.T().Log("About to run line #26: r.Range(3).Map(r.Range(4), func(x r.Term, y r.Term) interface{} { return 0})")

		runAndAssert(suite.Suite, expected_, r.Range(3).Map(r.Range(4), func(x r.Term, y r.Term) interface{} { return 0 }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #26")
	}

	{
		// transform/map.yaml line #31
		/* [[1]] */
		var expected_ []interface{} = []interface{}{[]interface{}{1}}
		/* r.expr([1]).map(lambda x:(x,)) */

		suite.T().Log("About to run line #31: r.Expr([]interface{}{1}).Map(func(x r.Term) interface{} { return []interface{}{x}})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1}).Map(func(x r.Term) interface{} { return []interface{}{x} }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #31")
	}

	{
		// transform/map.yaml line #36
		/* [[1, 1]] */
		var expected_ []interface{} = []interface{}{[]interface{}{1, 1}}
		/* r.expr([1]).map(r.expr([1]), lambda x, y:(x, y)) */

		suite.T().Log("About to run line #36: r.Expr([]interface{}{1}).Map(r.Expr([]interface{}{1}), func(x r.Term, y r.Term) interface{} { return []interface{}{x, y}})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1}).Map(r.Expr([]interface{}{1}), func(x r.Term, y r.Term) interface{} { return []interface{}{x, y} }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #36")
	}

	{
		// transform/map.yaml line #41
		/* [[1, 1, 1]] */
		var expected_ []interface{} = []interface{}{[]interface{}{1, 1, 1}}
		/* r.expr([1]).map(r.expr([1]), r.expr([1]), lambda x, y, z:(x, y, z)) */

		suite.T().Log("About to run line #41: r.Expr([]interface{}{1}).Map(r.Expr([]interface{}{1}), r.Expr([]interface{}{1}), func(x r.Term, y r.Term, z r.Term) interface{} { return []interface{}{x, y, z}})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1}).Map(r.Expr([]interface{}{1}), r.Expr([]interface{}{1}), func(x r.Term, y r.Term, z r.Term) interface{} { return []interface{}{x, y, z} }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #41")
	}

	{
		// transform/map.yaml line #47
		/* err("ReqlQueryLogicError", "The function passed to `map` expects 2 arguments, but 1 sequence was found.", []) */
		var expected_ Err = err("ReqlQueryLogicError", "The function passed to `map` expects 2 arguments, but 1 sequence was found.")
		/* r.expr([1]).map(lambda x, y:(x, y)) */

		suite.T().Log("About to run line #47: r.Expr([]interface{}{1}).Map(func(x r.Term, y r.Term) interface{} { return []interface{}{x, y}})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1}).Map(func(x r.Term, y r.Term) interface{} { return []interface{}{x, y} }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #47")
	}

	{
		// transform/map.yaml line #52
		/* err("ReqlQueryLogicError", "The function passed to `map` expects 1 argument, but 2 sequences were found.", []) */
		var expected_ Err = err("ReqlQueryLogicError", "The function passed to `map` expects 1 argument, but 2 sequences were found.")
		/* r.expr([1]).map(r.expr([1]), lambda x:(x,)) */

		suite.T().Log("About to run line #52: r.Expr([]interface{}{1}).Map(r.Expr([]interface{}{1}), func(x r.Term) interface{} { return []interface{}{x}})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1}).Map(r.Expr([]interface{}{1}), func(x r.Term) interface{} { return []interface{}{x} }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #52")
	}

	{
		// transform/map.yaml line #58
		/* [] */
		var expected_ []interface{} = []interface{}{}
		/* r.range().map(r.expr([]), lambda x, y:(x, y)) */

		suite.T().Log("About to run line #58: r.Range().Map(r.Expr([]interface{}{}), func(x r.Term, y r.Term) interface{} { return []interface{}{x, y}})")

		runAndAssert(suite.Suite, expected_, r.Range().Map(r.Expr([]interface{}{}), func(x r.Term, y r.Term) interface{} { return []interface{}{x, y} }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #58")
	}

	{
		// transform/map.yaml line #63
		/* [[1, 1], [2, 2]] */
		var expected_ []interface{} = []interface{}{[]interface{}{1, 1}, []interface{}{2, 2}}
		/* r.expr([1, 2]).map(r.expr([1, 2, 3, 4]), lambda x, y:(x, y)) */

		suite.T().Log("About to run line #63: r.Expr([]interface{}{1, 2}).Map(r.Expr([]interface{}{1, 2, 3, 4}), func(x r.Term, y r.Term) interface{} { return []interface{}{x, y}})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2}).Map(r.Expr([]interface{}{1, 2, 3, 4}), func(x r.Term, y r.Term) interface{} { return []interface{}{x, y} }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #63")
	}

	{
		// transform/map.yaml line #68
		/* [[0, 0], [1, 1]] */
		var expected_ []interface{} = []interface{}{[]interface{}{0, 0}, []interface{}{1, 1}}
		/* r.range(2).map(r.range(4), lambda x, y:(x, y)) */

		suite.T().Log("About to run line #68: r.Range(2).Map(r.Range(4), func(x r.Term, y r.Term) interface{} { return []interface{}{x, y}})")

		runAndAssert(suite.Suite, expected_, r.Range(2).Map(r.Range(4), func(x r.Term, y r.Term) interface{} { return []interface{}{x, y} }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #68")
	}

	{
		// transform/map.yaml line #73
		/* [[0, 1], [1, 2], [2, 3], [3, 4]] */
		var expected_ []interface{} = []interface{}{[]interface{}{0, 1}, []interface{}{1, 2}, []interface{}{2, 3}, []interface{}{3, 4}}
		/* r.range().map(r.expr([1, 2, 3, 4]), lambda x, y:(x, y)) */

		suite.T().Log("About to run line #73: r.Range().Map(r.Expr([]interface{}{1, 2, 3, 4}), func(x r.Term, y r.Term) interface{} { return []interface{}{x, y}})")

		runAndAssert(suite.Suite, expected_, r.Range().Map(r.Expr([]interface{}{1, 2, 3, 4}), func(x r.Term, y r.Term) interface{} { return []interface{}{x, y} }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #73")
	}

	{
		// transform/map.yaml line #78
		/* [[0, 0], [1, 1], [2, 2]] */
		var expected_ []interface{} = []interface{}{[]interface{}{0, 0}, []interface{}{1, 1}, []interface{}{2, 2}}
		/* r.range(3).map(r.range(5), r.js("(function(x, y){return [x, y];})")) */

		suite.T().Log("About to run line #78: r.Range(3).Map(r.Range(5), r.JS('(function(x, y){return [x, y];})'))")

		runAndAssert(suite.Suite, expected_, r.Range(3).Map(r.Range(5), r.JS("(function(x, y){return [x, y];})")), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #78")
	}

	{
		// transform/map.yaml line #83
		/* err("ReqlQueryLogicError", "Cannot convert NUMBER to SEQUENCE", []) */
		var expected_ Err = err("ReqlQueryLogicError", "Cannot convert NUMBER to SEQUENCE")
		/* r.range().map(r.expr(1), lambda x, y:(x, y)) */

		suite.T().Log("About to run line #83: r.Range().Map(r.Expr(1), func(x r.Term, y r.Term) interface{} { return []interface{}{x, y}})")

		runAndAssert(suite.Suite, expected_, r.Range().Map(r.Expr(1), func(x r.Term, y r.Term) interface{} { return []interface{}{x, y} }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #83")
	}

	{
		// transform/map.yaml line #89
		/* err("ReqlQueryLogicError", "Cannot use an infinite stream with an aggregation function (`reduce`, `count`, etc.) or coerce it to an array.", []) */
		var expected_ Err = err("ReqlQueryLogicError", "Cannot use an infinite stream with an aggregation function (`reduce`, `count`, etc.) or coerce it to an array.")
		/* r.range().map(r.range(), lambda x, y:(x, y)).count() */

		suite.T().Log("About to run line #89: r.Range().Map(r.Range(), func(x r.Term, y r.Term) interface{} { return []interface{}{x, y}}).Count()")

		runAndAssert(suite.Suite, expected_, r.Range().Map(r.Range(), func(x r.Term, y r.Term) interface{} { return []interface{}{x, y} }).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #89")
	}

	{
		// transform/map.yaml line #95
		/* [[0], [1], [2]] */
		var expected_ []interface{} = []interface{}{[]interface{}{0}, []interface{}{1}, []interface{}{2}}
		/* r.map(r.range(3), lambda x:(x,)) */

		suite.T().Log("About to run line #95: r.Map(r.Range(3), func(x r.Term) interface{} { return []interface{}{x}})")

		runAndAssert(suite.Suite, expected_, r.Map(r.Range(3), func(x r.Term) interface{} { return []interface{}{x} }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #95")
	}

	{
		// transform/map.yaml line #100
		/* [1, 2, 3] */
		var expected_ []interface{} = []interface{}{1, 2, 3}
		/* r.map(r.range(3), r.row + 1) */

		suite.T().Log("About to run line #100: r.Map(r.Range(3), r.Row.Add(1))")

		runAndAssert(suite.Suite, expected_, r.Map(r.Range(3), r.Row.Add(1)), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #100")
	}

	{
		// transform/map.yaml line #104
		/* [[0, 0], [1, 1], [2, 2]] */
		var expected_ []interface{} = []interface{}{[]interface{}{0, 0}, []interface{}{1, 1}, []interface{}{2, 2}}
		/* r.map(r.range(3), r.range(5), lambda x, y:(x, y)) */

		suite.T().Log("About to run line #104: r.Map(r.Range(3), r.Range(5), func(x r.Term, y r.Term) interface{} { return []interface{}{x, y}})")

		runAndAssert(suite.Suite, expected_, r.Map(r.Range(3), r.Range(5), func(x r.Term, y r.Term) interface{} { return []interface{}{x, y} }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #104")
	}
}
