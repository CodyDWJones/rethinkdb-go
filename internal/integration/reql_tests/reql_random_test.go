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

// Tests randomization functions
func TestRandomSuite(t *testing.T) {
	suite.Run(t, new(RandomSuite))
}

type RandomSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *RandomSuite) SetupTest() {
	suite.T().Log("Setting up RandomSuite")
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

func (suite *RandomSuite) TearDownSuite() {
	suite.T().Log("Tearing down RandomSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *RandomSuite) TestCases() {
	suite.T().Log("Running RandomSuite: Tests randomization functions")

	{
		// random.yaml line #5
		/* 3 */
		var expected_ int = 3
		/* r.expr([1,2,3]).sample(3).distinct().count() */

		suite.T().Log("About to run line #5: r.Expr([]interface{}{1, 2, 3}).Sample(3).Distinct().Count()")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).Sample(3).Distinct().Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #5")
	}

	{
		// random.yaml line #7
		/* 3 */
		var expected_ int = 3
		/* r.expr([1,2,3]).sample(3).count() */

		suite.T().Log("About to run line #7: r.Expr([]interface{}{1, 2, 3}).Sample(3).Count()")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).Sample(3).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #7")
	}

	{
		// random.yaml line #9
		/* 3 */
		var expected_ int = 3
		/* r.expr([1,2,3,4,5,6]).sample(3).distinct().count() */

		suite.T().Log("About to run line #9: r.Expr([]interface{}{1, 2, 3, 4, 5, 6}).Sample(3).Distinct().Count()")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3, 4, 5, 6}).Sample(3).Distinct().Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #9")
	}

	{
		// random.yaml line #11
		/* 3 */
		var expected_ int = 3
		/* r.expr([1,2,3]).sample(4).distinct().count() */

		suite.T().Log("About to run line #11: r.Expr([]interface{}{1, 2, 3}).Sample(4).Distinct().Count()")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).Sample(4).Distinct().Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #11")
	}

	{
		// random.yaml line #15
		/* err('ReqlQueryLogicError', 'Number of items to sample must be non-negative, got `-1`.', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Number of items to sample must be non-negative, got `-1`.")
		/* r.expr([1,2,3]).sample(-1) */

		suite.T().Log("About to run line #15: r.Expr([]interface{}{1, 2, 3}).Sample(-1)")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).Sample(-1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #15")
	}

	{
		// random.yaml line #17
		/* err('ReqlQueryLogicError', 'Cannot convert NUMBER to SEQUENCE', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Cannot convert NUMBER to SEQUENCE")
		/* r.expr(1).sample(1) */

		suite.T().Log("About to run line #17: r.Expr(1).Sample(1)")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Sample(1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #17")
	}

	{
		// random.yaml line #19
		/* err('ReqlQueryLogicError', 'Cannot convert OBJECT to SEQUENCE', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Cannot convert OBJECT to SEQUENCE")
		/* r.expr({}).sample(1) */

		suite.T().Log("About to run line #19: r.Expr(map[interface{}]interface{}{}).Sample(1)")

		runAndAssert(suite.Suite, expected_, r.Expr(map[interface{}]interface{}{}).Sample(1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #19")
	}

	{
		// random.yaml line #25
		/* True */
		var expected_ bool = true
		/* r.random().do(lambda x:r.and_(x.ge(0), x.lt(1))) */

		suite.T().Log("About to run line #25: r.Random().Do(func(x r.Term) interface{} { return r.And(x.Ge(0), x.Lt(1))})")

		runAndAssert(suite.Suite, expected_, r.Random().Do(func(x r.Term) interface{} { return r.And(x.Ge(0), x.Lt(1)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #25")
	}

	{
		// random.yaml line #26
		/* True */
		var expected_ bool = true
		/* r.random(1, float=True).do(lambda x:r.and_(x.ge(0), x.lt(1))) */

		suite.T().Log("About to run line #26: r.Random(1).OptArgs(r.RandomOpts{Float: true, }).Do(func(x r.Term) interface{} { return r.And(x.Ge(0), x.Lt(1))})")

		runAndAssert(suite.Suite, expected_, r.Random(1).OptArgs(r.RandomOpts{Float: true}).Do(func(x r.Term) interface{} { return r.And(x.Ge(0), x.Lt(1)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #26")
	}

	{
		// random.yaml line #27
		/* True */
		var expected_ bool = true
		/* r.random(0, 1, float=True).do(lambda x:r.and_(x.ge(0), x.lt(1))) */

		suite.T().Log("About to run line #27: r.Random(0, 1).OptArgs(r.RandomOpts{Float: true, }).Do(func(x r.Term) interface{} { return r.And(x.Ge(0), x.Lt(1))})")

		runAndAssert(suite.Suite, expected_, r.Random(0, 1).OptArgs(r.RandomOpts{Float: true}).Do(func(x r.Term) interface{} { return r.And(x.Ge(0), x.Lt(1)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #27")
	}

	{
		// random.yaml line #28
		/* True */
		var expected_ bool = true
		/* r.random(1, 0, float=True).do(lambda x:r.and_(x.le(1), x.gt(0))) */

		suite.T().Log("About to run line #28: r.Random(1, 0).OptArgs(r.RandomOpts{Float: true, }).Do(func(x r.Term) interface{} { return r.And(x.Le(1), x.Gt(0))})")

		runAndAssert(suite.Suite, expected_, r.Random(1, 0).OptArgs(r.RandomOpts{Float: true}).Do(func(x r.Term) interface{} { return r.And(x.Le(1), x.Gt(0)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #28")
	}

	{
		// random.yaml line #29
		/* True */
		var expected_ bool = true
		/* r.random(r.expr(0), 1, float=True).do(lambda x:r.and_(x.ge(0), x.lt(1))) */

		suite.T().Log("About to run line #29: r.Random(r.Expr(0), 1).OptArgs(r.RandomOpts{Float: true, }).Do(func(x r.Term) interface{} { return r.And(x.Ge(0), x.Lt(1))})")

		runAndAssert(suite.Suite, expected_, r.Random(r.Expr(0), 1).OptArgs(r.RandomOpts{Float: true}).Do(func(x r.Term) interface{} { return r.And(x.Ge(0), x.Lt(1)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #29")
	}

	{
		// random.yaml line #30
		/* True */
		var expected_ bool = true
		/* r.random(1, r.expr(0), float=True).do(lambda x:r.and_(x.le(1), x.gt(0))) */

		suite.T().Log("About to run line #30: r.Random(1, r.Expr(0)).OptArgs(r.RandomOpts{Float: true, }).Do(func(x r.Term) interface{} { return r.And(x.Le(1), x.Gt(0))})")

		runAndAssert(suite.Suite, expected_, r.Random(1, r.Expr(0)).OptArgs(r.RandomOpts{Float: true}).Do(func(x r.Term) interface{} { return r.And(x.Le(1), x.Gt(0)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #30")
	}

	{
		// random.yaml line #31
		/* True */
		var expected_ bool = true
		/* r.random(r.expr(1), r.expr(0), float=True).do(lambda x:r.and_(x.le(1), x.gt(0))) */

		suite.T().Log("About to run line #31: r.Random(r.Expr(1), r.Expr(0)).OptArgs(r.RandomOpts{Float: true, }).Do(func(x r.Term) interface{} { return r.And(x.Le(1), x.Gt(0))})")

		runAndAssert(suite.Suite, expected_, r.Random(r.Expr(1), r.Expr(0)).OptArgs(r.RandomOpts{Float: true}).Do(func(x r.Term) interface{} { return r.And(x.Le(1), x.Gt(0)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #31")
	}

	{
		// random.yaml line #36
		/* True */
		var expected_ bool = true
		/* r.random(0.495, float=True).do(lambda x:r.and_(x.ge(0), x.lt(0.495))) */

		suite.T().Log("About to run line #36: r.Random(0.495).OptArgs(r.RandomOpts{Float: true, }).Do(func(x r.Term) interface{} { return r.And(x.Ge(0), x.Lt(0.495))})")

		runAndAssert(suite.Suite, expected_, r.Random(0.495).OptArgs(r.RandomOpts{Float: true}).Do(func(x r.Term) interface{} { return r.And(x.Ge(0), x.Lt(0.495)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #36")
	}

	{
		// random.yaml line #37
		/* True */
		var expected_ bool = true
		/* r.random(-0.495, float=True).do(lambda x:r.and_(x.le(0), x.gt(-0.495))) */

		suite.T().Log("About to run line #37: r.Random(-0.495).OptArgs(r.RandomOpts{Float: true, }).Do(func(x r.Term) interface{} { return r.And(x.Le(0), x.Gt(-0.495))})")

		runAndAssert(suite.Suite, expected_, r.Random(-0.495).OptArgs(r.RandomOpts{Float: true}).Do(func(x r.Term) interface{} { return r.And(x.Le(0), x.Gt(-0.495)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #37")
	}

	{
		// random.yaml line #38
		/* True */
		var expected_ bool = true
		/* r.random(1823756.24, float=True).do(lambda x:r.and_(x.ge(0), x.lt(1823756.24))) */

		suite.T().Log("About to run line #38: r.Random(1823756.24).OptArgs(r.RandomOpts{Float: true, }).Do(func(x r.Term) interface{} { return r.And(x.Ge(0), x.Lt(1823756.24))})")

		runAndAssert(suite.Suite, expected_, r.Random(1823756.24).OptArgs(r.RandomOpts{Float: true}).Do(func(x r.Term) interface{} { return r.And(x.Ge(0), x.Lt(1823756.24)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #38")
	}

	{
		// random.yaml line #39
		/* True */
		var expected_ bool = true
		/* r.random(-1823756.24, float=True).do(lambda x:r.and_(x.le(0), x.gt(-1823756.24))) */

		suite.T().Log("About to run line #39: r.Random(-1823756.24).OptArgs(r.RandomOpts{Float: true, }).Do(func(x r.Term) interface{} { return r.And(x.Le(0), x.Gt(-1823756.24))})")

		runAndAssert(suite.Suite, expected_, r.Random(-1823756.24).OptArgs(r.RandomOpts{Float: true}).Do(func(x r.Term) interface{} { return r.And(x.Le(0), x.Gt(-1823756.24)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #39")
	}

	{
		// random.yaml line #44
		/* True */
		var expected_ bool = true
		/* r.random(10.5, 20.153, float=True).do(lambda x:r.and_(x.ge(10.5), x.lt(20.153))) */

		suite.T().Log("About to run line #44: r.Random(10.5, 20.153).OptArgs(r.RandomOpts{Float: true, }).Do(func(x r.Term) interface{} { return r.And(x.Ge(10.5), x.Lt(20.153))})")

		runAndAssert(suite.Suite, expected_, r.Random(10.5, 20.153).OptArgs(r.RandomOpts{Float: true}).Do(func(x r.Term) interface{} { return r.And(x.Ge(10.5), x.Lt(20.153)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #44")
	}

	{
		// random.yaml line #45
		/* True */
		var expected_ bool = true
		/* r.random(20.153, 10.5, float=True).do(lambda x:r.and_(x.le(20.153), x.gt(10.5))) */

		suite.T().Log("About to run line #45: r.Random(20.153, 10.5).OptArgs(r.RandomOpts{Float: true, }).Do(func(x r.Term) interface{} { return r.And(x.Le(20.153), x.Gt(10.5))})")

		runAndAssert(suite.Suite, expected_, r.Random(20.153, 10.5).OptArgs(r.RandomOpts{Float: true}).Do(func(x r.Term) interface{} { return r.And(x.Le(20.153), x.Gt(10.5)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #45")
	}

	{
		// random.yaml line #46
		/* True */
		var expected_ bool = true
		/* r.random(31415926.1, 31415926, float=True).do(lambda x:r.and_(x.le(31415926.1), x.gt(31415926))) */

		suite.T().Log("About to run line #46: r.Random(31415926.1, 31415926).OptArgs(r.RandomOpts{Float: true, }).Do(func(x r.Term) interface{} { return r.And(x.Le(31415926.1), x.Gt(31415926))})")

		runAndAssert(suite.Suite, expected_, r.Random(31415926.1, 31415926).OptArgs(r.RandomOpts{Float: true}).Do(func(x r.Term) interface{} { return r.And(x.Le(31415926.1), x.Gt(31415926)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #46")
	}

	{
		// random.yaml line #51
		/* True */
		var expected_ bool = true
		/* r.random(-10.5, 20.153, float=True).do(lambda x:r.and_(x.ge(-10.5), x.lt(20.153))) */

		suite.T().Log("About to run line #51: r.Random(-10.5, 20.153).OptArgs(r.RandomOpts{Float: true, }).Do(func(x r.Term) interface{} { return r.And(x.Ge(-10.5), x.Lt(20.153))})")

		runAndAssert(suite.Suite, expected_, r.Random(-10.5, 20.153).OptArgs(r.RandomOpts{Float: true}).Do(func(x r.Term) interface{} { return r.And(x.Ge(-10.5), x.Lt(20.153)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #51")
	}

	{
		// random.yaml line #52
		/* True */
		var expected_ bool = true
		/* r.random(-20.153, -10.5, float=True).do(lambda x:r.and_(x.ge(-20.153), x.lt(-10.5))) */

		suite.T().Log("About to run line #52: r.Random(-20.153, -10.5).OptArgs(r.RandomOpts{Float: true, }).Do(func(x r.Term) interface{} { return r.And(x.Ge(-20.153), x.Lt(-10.5))})")

		runAndAssert(suite.Suite, expected_, r.Random(-20.153, -10.5).OptArgs(r.RandomOpts{Float: true}).Do(func(x r.Term) interface{} { return r.And(x.Ge(-20.153), x.Lt(-10.5)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #52")
	}

	{
		// random.yaml line #53
		/* True */
		var expected_ bool = true
		/* r.random(-31415926, -31415926.1, float=True).do(lambda x:r.and_(x.le(-31415926), x.gt(-31415926.1))) */

		suite.T().Log("About to run line #53: r.Random(-31415926, -31415926.1).OptArgs(r.RandomOpts{Float: true, }).Do(func(x r.Term) interface{} { return r.And(x.Le(-31415926), x.Gt(-31415926.1))})")

		runAndAssert(suite.Suite, expected_, r.Random(-31415926, -31415926.1).OptArgs(r.RandomOpts{Float: true}).Do(func(x r.Term) interface{} { return r.And(x.Le(-31415926), x.Gt(-31415926.1)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #53")
	}

	{
		// random.yaml line #58
		/* 2 */
		var expected_ int = 2
		/* r.expr([r.random(), r.random()]).distinct().count() */

		suite.T().Log("About to run line #58: r.Expr([]interface{}{r.Random(), r.Random()}).Distinct().Count()")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{r.Random(), r.Random()}).Distinct().Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #58")
	}

	{
		// random.yaml line #59
		/* 2 */
		var expected_ int = 2
		/* r.expr([r.random(1, float=True), r.random(1, float=True)]).distinct().count() */

		suite.T().Log("About to run line #59: r.Expr([]interface{}{r.Random(1).OptArgs(r.RandomOpts{Float: true, }), r.Random(1).OptArgs(r.RandomOpts{Float: true, })}).Distinct().Count()")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{r.Random(1).OptArgs(r.RandomOpts{Float: true}), r.Random(1).OptArgs(r.RandomOpts{Float: true})}).Distinct().Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #59")
	}

	{
		// random.yaml line #60
		/* 2 */
		var expected_ int = 2
		/* r.expr([r.random(0, 1, float=True), r.random(0, 1, float=True)]).distinct().count() */

		suite.T().Log("About to run line #60: r.Expr([]interface{}{r.Random(0, 1).OptArgs(r.RandomOpts{Float: true, }), r.Random(0, 1).OptArgs(r.RandomOpts{Float: true, })}).Distinct().Count()")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{r.Random(0, 1).OptArgs(r.RandomOpts{Float: true}), r.Random(0, 1).OptArgs(r.RandomOpts{Float: true})}).Distinct().Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #60")
	}

	{
		// random.yaml line #65
		/* True */
		var expected_ bool = true
		/* r.random(0, float=True).eq(0) */

		suite.T().Log("About to run line #65: r.Random(0).OptArgs(r.RandomOpts{Float: true, }).Eq(0)")

		runAndAssert(suite.Suite, expected_, r.Random(0).OptArgs(r.RandomOpts{Float: true}).Eq(0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #65")
	}

	{
		// random.yaml line #66
		/* True */
		var expected_ bool = true
		/* r.random(5, 5, float=True).eq(5) */

		suite.T().Log("About to run line #66: r.Random(5, 5).OptArgs(r.RandomOpts{Float: true, }).Eq(5)")

		runAndAssert(suite.Suite, expected_, r.Random(5, 5).OptArgs(r.RandomOpts{Float: true}).Eq(5), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #66")
	}

	{
		// random.yaml line #67
		/* True */
		var expected_ bool = true
		/* r.random(-499384756758, -499384756758, float=True).eq(-499384756758) */

		suite.T().Log("About to run line #67: r.Random(-499384756758, -499384756758).OptArgs(r.RandomOpts{Float: true, }).Eq(-499384756758)")

		runAndAssert(suite.Suite, expected_, r.Random(-499384756758, -499384756758).OptArgs(r.RandomOpts{Float: true}).Eq(-499384756758), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #67")
	}

	{
		// random.yaml line #68
		/* True */
		var expected_ bool = true
		/* r.random(-93.94757, -93.94757, float=True).eq(-93.94757) */

		suite.T().Log("About to run line #68: r.Random(-93.94757, -93.94757).OptArgs(r.RandomOpts{Float: true, }).Eq(-93.94757)")

		runAndAssert(suite.Suite, expected_, r.Random(-93.94757, -93.94757).OptArgs(r.RandomOpts{Float: true}).Eq(-93.94757), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #68")
	}

	{
		// random.yaml line #69
		/* True */
		var expected_ bool = true
		/* r.random(294.69148, 294.69148, float=True).eq(294.69148) */

		suite.T().Log("About to run line #69: r.Random(294.69148, 294.69148).OptArgs(r.RandomOpts{Float: true, }).Eq(294.69148)")

		runAndAssert(suite.Suite, expected_, r.Random(294.69148, 294.69148).OptArgs(r.RandomOpts{Float: true}).Eq(294.69148), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #69")
	}

	// random.yaml line #74
	// float_max = sys.float_info.max
	suite.T().Log("Possibly executing: var float_max float64 = sys.FloatInfo.Max")

	float_max := sys.FloatInfo.Max
	_ = float_max // Prevent any noused variable errors

	// random.yaml line #78
	// float_min = sys.float_info.min
	suite.T().Log("Possibly executing: var float_min float64 = sys.FloatInfo.Min")

	float_min := sys.FloatInfo.Min
	_ = float_min // Prevent any noused variable errors

	{
		// random.yaml line #82
		/* True */
		var expected_ bool = true
		/* r.random(-float_max, float_max, float=True).do(lambda x:r.and_(x.ge(-float_max), x.lt(float_max))) */

		suite.T().Log("About to run line #82: r.Random(-float_max, float_max).OptArgs(r.RandomOpts{Float: true, }).Do(func(x r.Term) interface{} { return r.And(x.Ge(-float_max), x.Lt(float_max))})")

		runAndAssert(suite.Suite, expected_, r.Random(-float_max, float_max).OptArgs(r.RandomOpts{Float: true}).Do(func(x r.Term) interface{} { return r.And(x.Ge(-float_max), x.Lt(float_max)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #82")
	}

	{
		// random.yaml line #83
		/* True */
		var expected_ bool = true
		/* r.random(float_max, -float_max, float=True).do(lambda x:r.and_(x.le(float_max), x.gt(-float_max))) */

		suite.T().Log("About to run line #83: r.Random(float_max, -float_max).OptArgs(r.RandomOpts{Float: true, }).Do(func(x r.Term) interface{} { return r.And(x.Le(float_max), x.Gt(-float_max))})")

		runAndAssert(suite.Suite, expected_, r.Random(float_max, -float_max).OptArgs(r.RandomOpts{Float: true}).Do(func(x r.Term) interface{} { return r.And(x.Le(float_max), x.Gt(-float_max)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #83")
	}

	{
		// random.yaml line #84
		/* True */
		var expected_ bool = true
		/* r.random(float_min, float_max, float=True).do(lambda x:r.and_(x.ge(float_min), x.lt(float_max))) */

		suite.T().Log("About to run line #84: r.Random(float_min, float_max).OptArgs(r.RandomOpts{Float: true, }).Do(func(x r.Term) interface{} { return r.And(x.Ge(float_min), x.Lt(float_max))})")

		runAndAssert(suite.Suite, expected_, r.Random(float_min, float_max).OptArgs(r.RandomOpts{Float: true}).Do(func(x r.Term) interface{} { return r.And(x.Ge(float_min), x.Lt(float_max)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #84")
	}

	{
		// random.yaml line #85
		/* True */
		var expected_ bool = true
		/* r.random(float_min, -float_max, float=True).do(lambda x:r.and_(x.le(float_min), x.gt(-float_max))) */

		suite.T().Log("About to run line #85: r.Random(float_min, -float_max).OptArgs(r.RandomOpts{Float: true, }).Do(func(x r.Term) interface{} { return r.And(x.Le(float_min), x.Gt(-float_max))})")

		runAndAssert(suite.Suite, expected_, r.Random(float_min, -float_max).OptArgs(r.RandomOpts{Float: true}).Do(func(x r.Term) interface{} { return r.And(x.Le(float_min), x.Gt(-float_max)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #85")
	}

	{
		// random.yaml line #86
		/* True */
		var expected_ bool = true
		/* r.random(-float_min, float_max, float=True).do(lambda x:r.and_(x.ge(-float_min), x.lt(float_max))) */

		suite.T().Log("About to run line #86: r.Random(-float_min, float_max).OptArgs(r.RandomOpts{Float: true, }).Do(func(x r.Term) interface{} { return r.And(x.Ge(-float_min), x.Lt(float_max))})")

		runAndAssert(suite.Suite, expected_, r.Random(-float_min, float_max).OptArgs(r.RandomOpts{Float: true}).Do(func(x r.Term) interface{} { return r.And(x.Ge(-float_min), x.Lt(float_max)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #86")
	}

	{
		// random.yaml line #87
		/* True */
		var expected_ bool = true
		/* r.random(-float_min, -float_max, float=True).do(lambda x:r.and_(x.le(-float_min), x.gt(-float_max))) */

		suite.T().Log("About to run line #87: r.Random(-float_min, -float_max).OptArgs(r.RandomOpts{Float: true, }).Do(func(x r.Term) interface{} { return r.And(x.Le(-float_min), x.Gt(-float_max))})")

		runAndAssert(suite.Suite, expected_, r.Random(-float_min, -float_max).OptArgs(r.RandomOpts{Float: true}).Do(func(x r.Term) interface{} { return r.And(x.Le(-float_min), x.Gt(-float_max)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #87")
	}

	// random.yaml line #92
	// upper_limit = 2**53 - 1
	suite.T().Log("Possibly executing: var upper_limit int = 2<<52 - 1")

	upper_limit := 2<<52 - 1
	_ = upper_limit // Prevent any noused variable errors

	// random.yaml line #96
	// lower_limit = 1 - (2**53)
	suite.T().Log("Possibly executing: var lower_limit int = 1 - 2<<52")

	lower_limit := 1 - 2<<52
	_ = lower_limit // Prevent any noused variable errors

	{
		// random.yaml line #101
		/* True */
		var expected_ bool = true
		/* r.random(256).do(lambda x:r.and_(x.ge(0), x.lt(256))) */

		suite.T().Log("About to run line #101: r.Random(256).Do(func(x r.Term) interface{} { return r.And(x.Ge(0), x.Lt(256))})")

		runAndAssert(suite.Suite, expected_, r.Random(256).Do(func(x r.Term) interface{} { return r.And(x.Ge(0), x.Lt(256)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #101")
	}

	{
		// random.yaml line #102
		/* True */
		var expected_ bool = true
		/* r.random(0, 256).do(lambda x:r.and_(x.ge(0), x.lt(256))) */

		suite.T().Log("About to run line #102: r.Random(0, 256).Do(func(x r.Term) interface{} { return r.And(x.Ge(0), x.Lt(256))})")

		runAndAssert(suite.Suite, expected_, r.Random(0, 256).Do(func(x r.Term) interface{} { return r.And(x.Ge(0), x.Lt(256)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #102")
	}

	{
		// random.yaml line #103
		/* True */
		var expected_ bool = true
		/* r.random(r.expr(256)).do(lambda x:r.and_(x.ge(0), x.lt(256))) */

		suite.T().Log("About to run line #103: r.Random(r.Expr(256)).Do(func(x r.Term) interface{} { return r.And(x.Ge(0), x.Lt(256))})")

		runAndAssert(suite.Suite, expected_, r.Random(r.Expr(256)).Do(func(x r.Term) interface{} { return r.And(x.Ge(0), x.Lt(256)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #103")
	}

	{
		// random.yaml line #104
		/* True */
		var expected_ bool = true
		/* r.random(r.expr(0), 256).do(lambda x:r.and_(x.ge(0), x.lt(256))) */

		suite.T().Log("About to run line #104: r.Random(r.Expr(0), 256).Do(func(x r.Term) interface{} { return r.And(x.Ge(0), x.Lt(256))})")

		runAndAssert(suite.Suite, expected_, r.Random(r.Expr(0), 256).Do(func(x r.Term) interface{} { return r.And(x.Ge(0), x.Lt(256)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #104")
	}

	{
		// random.yaml line #105
		/* True */
		var expected_ bool = true
		/* r.random(0, r.expr(256)).do(lambda x:r.and_(x.ge(0), x.lt(256))) */

		suite.T().Log("About to run line #105: r.Random(0, r.Expr(256)).Do(func(x r.Term) interface{} { return r.And(x.Ge(0), x.Lt(256))})")

		runAndAssert(suite.Suite, expected_, r.Random(0, r.Expr(256)).Do(func(x r.Term) interface{} { return r.And(x.Ge(0), x.Lt(256)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #105")
	}

	{
		// random.yaml line #106
		/* True */
		var expected_ bool = true
		/* r.random(r.expr(0), r.expr(256)).do(lambda x:r.and_(x.ge(0), x.lt(256))) */

		suite.T().Log("About to run line #106: r.Random(r.Expr(0), r.Expr(256)).Do(func(x r.Term) interface{} { return r.And(x.Ge(0), x.Lt(256))})")

		runAndAssert(suite.Suite, expected_, r.Random(r.Expr(0), r.Expr(256)).Do(func(x r.Term) interface{} { return r.And(x.Ge(0), x.Lt(256)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #106")
	}

	{
		// random.yaml line #111
		/* True */
		var expected_ bool = true
		/* r.random(10, 20).do(lambda x:r.and_(x.ge(10), x.lt(20))) */

		suite.T().Log("About to run line #111: r.Random(10, 20).Do(func(x r.Term) interface{} { return r.And(x.Ge(10), x.Lt(20))})")

		runAndAssert(suite.Suite, expected_, r.Random(10, 20).Do(func(x r.Term) interface{} { return r.And(x.Ge(10), x.Lt(20)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #111")
	}

	{
		// random.yaml line #112
		/* True */
		var expected_ bool = true
		/* r.random(9347849, 120937493).do(lambda x:r.and_(x.ge(9347849), x.lt(120937493))) */

		suite.T().Log("About to run line #112: r.Random(9347849, 120937493).Do(func(x r.Term) interface{} { return r.And(x.Ge(9347849), x.Lt(120937493))})")

		runAndAssert(suite.Suite, expected_, r.Random(9347849, 120937493).Do(func(x r.Term) interface{} { return r.And(x.Ge(9347849), x.Lt(120937493)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #112")
	}

	{
		// random.yaml line #123
		/* True */
		var expected_ bool = true
		/* r.random(-10, 20).do(lambda x:r.and_(x.ge(-10), x.lt(20))) */

		suite.T().Log("About to run line #123: r.Random(-10, 20).Do(func(x r.Term) interface{} { return r.And(x.Ge(-10), x.Lt(20))})")

		runAndAssert(suite.Suite, expected_, r.Random(-10, 20).Do(func(x r.Term) interface{} { return r.And(x.Ge(-10), x.Lt(20)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #123")
	}

	{
		// random.yaml line #124
		/* True */
		var expected_ bool = true
		/* r.random(-20, -10).do(lambda x:r.and_(x.ge(-20), x.lt(-10))) */

		suite.T().Log("About to run line #124: r.Random(-20, -10).Do(func(x r.Term) interface{} { return r.And(x.Ge(-20), x.Lt(-10))})")

		runAndAssert(suite.Suite, expected_, r.Random(-20, -10).Do(func(x r.Term) interface{} { return r.And(x.Ge(-20), x.Lt(-10)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #124")
	}

	{
		// random.yaml line #125
		/* True */
		var expected_ bool = true
		/* r.random(-120937493, -9347849).do(lambda x:r.and_(x.ge(-120937493), x.lt(-9347849))) */

		suite.T().Log("About to run line #125: r.Random(-120937493, -9347849).Do(func(x r.Term) interface{} { return r.And(x.Ge(-120937493), x.Lt(-9347849))})")

		runAndAssert(suite.Suite, expected_, r.Random(-120937493, -9347849).Do(func(x r.Term) interface{} { return r.And(x.Ge(-120937493), x.Lt(-9347849)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #125")
	}

	{
		// random.yaml line #137
		/* 2 */
		var expected_ int = 2
		/* r.expr([r.random(upper_limit), r.random(upper_limit)]).distinct().count() */

		suite.T().Log("About to run line #137: r.Expr([]interface{}{r.Random(upper_limit), r.Random(upper_limit)}).Distinct().Count()")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{r.Random(upper_limit), r.Random(upper_limit)}).Distinct().Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #137")
	}

	{
		// random.yaml line #139
		/* 2 */
		var expected_ int = 2
		/* r.expr([upper_limit,upper_limit]).map(lambda x:r.random(x)).distinct().count() */

		suite.T().Log("About to run line #139: r.Expr([]interface{}{upper_limit, upper_limit}).Map(func(x r.Term) interface{} { return r.Random(x)}).Distinct().Count()")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{upper_limit, upper_limit}).Map(func(x r.Term) interface{} { return r.Random(x) }).Distinct().Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #139")
	}

	{
		// random.yaml line #147
		/* err("ReqlQueryLogicError", "Upper bound (-0.5) could not be safely converted to an integer.", []) */
		var expected_ Err = err("ReqlQueryLogicError", "Upper bound (-0.5) could not be safely converted to an integer.")
		/* r.random(-0.5) */

		suite.T().Log("About to run line #147: r.Random(-0.5)")

		runAndAssert(suite.Suite, expected_, r.Random(-0.5), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #147")
	}

	{
		// random.yaml line #149
		/* err("ReqlQueryLogicError", "Upper bound (0.25) could not be safely converted to an integer.", []) */
		var expected_ Err = err("ReqlQueryLogicError", "Upper bound (0.25) could not be safely converted to an integer.")
		/* r.random(0.25) */

		suite.T().Log("About to run line #149: r.Random(0.25)")

		runAndAssert(suite.Suite, expected_, r.Random(0.25), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #149")
	}

	{
		// random.yaml line #151
		/* err("ReqlQueryLogicError", "Upper bound (0.75) could not be safely converted to an integer.", []) */
		var expected_ Err = err("ReqlQueryLogicError", "Upper bound (0.75) could not be safely converted to an integer.")
		/* r.random(-10, 0.75) */

		suite.T().Log("About to run line #151: r.Random(-10, 0.75)")

		runAndAssert(suite.Suite, expected_, r.Random(-10, 0.75), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #151")
	}

	{
		// random.yaml line #153
		/* err("ReqlQueryLogicError", "Lower bound (-120549.25) could not be safely converted to an integer.", []) */
		var expected_ Err = err("ReqlQueryLogicError", "Lower bound (-120549.25) could not be safely converted to an integer.")
		/* r.random(-120549.25, 39458) */

		suite.T().Log("About to run line #153: r.Random(-120549.25, 39458)")

		runAndAssert(suite.Suite, expected_, r.Random(-120549.25, 39458), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #153")
	}

	{
		// random.yaml line #155
		/* err("ReqlQueryLogicError", "Lower bound (-6.5) could not be safely converted to an integer.", []) */
		var expected_ Err = err("ReqlQueryLogicError", "Lower bound (-6.5) could not be safely converted to an integer.")
		/* r.random(-6.5, 8.125) */

		suite.T().Log("About to run line #155: r.Random(-6.5, 8.125)")

		runAndAssert(suite.Suite, expected_, r.Random(-6.5, 8.125), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #155")
	}

	{
		// random.yaml line #159
		/* err("ReqlQueryLogicError", "Generating a random integer requires one or two bounds.", []) */
		var expected_ Err = err("ReqlQueryLogicError", "Generating a random integer requires one or two bounds.")
		/* r.random(float=False) */

		suite.T().Log("About to run line #159: r.Random().OptArgs(r.RandomOpts{Float: false, })")

		runAndAssert(suite.Suite, expected_, r.Random().OptArgs(r.RandomOpts{Float: false}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #159")
	}

	{
		// random.yaml line #165
		/* err("ReqlQueryLogicError", "Lower bound (0) is not less than upper bound (0).", []) */
		var expected_ Err = err("ReqlQueryLogicError", "Lower bound (0) is not less than upper bound (0).")
		/* r.random(0) */

		suite.T().Log("About to run line #165: r.Random(0)")

		runAndAssert(suite.Suite, expected_, r.Random(0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #165")
	}

	{
		// random.yaml line #167
		/* err("ReqlQueryLogicError", "Lower bound (0) is not less than upper bound (0).", []) */
		var expected_ Err = err("ReqlQueryLogicError", "Lower bound (0) is not less than upper bound (0).")
		/* r.random(0, 0) */

		suite.T().Log("About to run line #167: r.Random(0, 0)")

		runAndAssert(suite.Suite, expected_, r.Random(0, 0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #167")
	}

	{
		// random.yaml line #169
		/* err("ReqlQueryLogicError", "Lower bound (515) is not less than upper bound (515).", []) */
		var expected_ Err = err("ReqlQueryLogicError", "Lower bound (515) is not less than upper bound (515).")
		/* r.random(515, 515) */

		suite.T().Log("About to run line #169: r.Random(515, 515)")

		runAndAssert(suite.Suite, expected_, r.Random(515, 515), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #169")
	}

	{
		// random.yaml line #171
		/* err("ReqlQueryLogicError", "Lower bound (-956) is not less than upper bound (-956).", []) */
		var expected_ Err = err("ReqlQueryLogicError", "Lower bound (-956) is not less than upper bound (-956).")
		/* r.random(-956, -956) */

		suite.T().Log("About to run line #171: r.Random(-956, -956)")

		runAndAssert(suite.Suite, expected_, r.Random(-956, -956), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #171")
	}

	{
		// random.yaml line #173
		/* err("ReqlQueryLogicError", "Lower bound (0) is not less than upper bound (-10).", []) */
		var expected_ Err = err("ReqlQueryLogicError", "Lower bound (0) is not less than upper bound (-10).")
		/* r.random(-10) */

		suite.T().Log("About to run line #173: r.Random(-10)")

		runAndAssert(suite.Suite, expected_, r.Random(-10), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #173")
	}

	{
		// random.yaml line #175
		/* err("ReqlQueryLogicError", "Lower bound (20) is not less than upper bound (2).", []) */
		var expected_ Err = err("ReqlQueryLogicError", "Lower bound (20) is not less than upper bound (2).")
		/* r.random(20, 2) */

		suite.T().Log("About to run line #175: r.Random(20, 2)")

		runAndAssert(suite.Suite, expected_, r.Random(20, 2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #175")
	}

	{
		// random.yaml line #177
		/* err("ReqlQueryLogicError", "Lower bound (2) is not less than upper bound (-20).", []) */
		var expected_ Err = err("ReqlQueryLogicError", "Lower bound (2) is not less than upper bound (-20).")
		/* r.random(2, -20) */

		suite.T().Log("About to run line #177: r.Random(2, -20)")

		runAndAssert(suite.Suite, expected_, r.Random(2, -20), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #177")
	}

	{
		// random.yaml line #179
		/* err("ReqlQueryLogicError", "Lower bound (1456) is not less than upper bound (0).", []) */
		var expected_ Err = err("ReqlQueryLogicError", "Lower bound (1456) is not less than upper bound (0).")
		/* r.random(1456, 0) */

		suite.T().Log("About to run line #179: r.Random(1456, 0)")

		runAndAssert(suite.Suite, expected_, r.Random(1456, 0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #179")
	}
}
