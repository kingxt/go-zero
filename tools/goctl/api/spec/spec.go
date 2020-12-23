package spec

type (
	Doc []string

	Annotation struct {
		Properties map[string]string
	}

	LineColumn struct {
		Line   int `json:"-"`
		Column int `json:"-"`
	}

	ApiSyntax struct {
		LineColumn
		Version string
	}

	ApiSpec struct {
		Filename string
		Syntax   ApiSyntax
		Import   ApiImport
		Info     Info
		Types    []Type
		Service  Service
	}

	ApiImport struct {
		List []Import
	}

	Import struct {
		LineColumn
		Value string
	}

	Group struct {
		Annotation Annotation
		Routes     []Route
	}

	Info struct {
		// Deprecated: use Proterties instead
		Title string
		// Deprecated: use Proterties instead
		Desc string
		// Deprecated: use Proterties instead
		Version string
		// Deprecated: use Proterties instead
		Author string
		// Deprecated: use Proterties instead
		Email string
		LineColumn
		Proterties map[string]string
	}

	Member struct {
		Annotations []Annotation
		LineColumn
		Name string
		// 数据类型字面值，如：string、map[int]string、[]int64、[]*User
		Type string
		// it can be asserted as BasicType: int、bool、
		// PointerType: *string、*User、
		// MapType: map[${BasicType}]interface、
		// ArrayType:[]int、[]User、[]*User
		// InterfaceType: interface{}
		// Type
		Expr interface{}
		Tag  string
		// Deprecated
		Comment string // 换成标准struct中将废弃
		// 成员尾部注释说明
		Comments []string
		// 成员头顶注释说明
		Docs     Doc
		IsInline bool
	}

	Route struct {
		LineColumn
		Annotation        Annotation
		Method            string
		Path              string
		RequestType       Type
		ResponseType      Type
		Docs              Doc
		HandlerLineColumn LineColumn `json:"-"`
		Handler           string
	}

	Service struct {
		Name   string
		Groups []Group
	}

	Type struct {
		Name string
		LineColumn
		Annotations []Annotation
		Members     []Member
	}

	// 系统预设基本数据类型
	BasicType struct {
		StringExpr string
		Name       string
		LineColumn
	}

	PointerType struct {
		StringExpr string
		// it can be asserted as BasicType: int、bool、
		// PointerType: *string、*User、
		// MapType: map[${BasicType}]interface、
		// ArrayType:[]int、[]User、[]*User
		// InterfaceType: interface{}
		// Type
		Star interface{}
		LineColumn
	}

	MapType struct {
		StringExpr string
		// only support the BasicType
		Key string
		// it can be asserted as BasicType: int、bool、
		// PointerType: *string、*User、
		// MapType: map[${BasicType}]interface、
		// ArrayType:[]int、[]User、[]*User
		// InterfaceType: interface{}
		// Type
		Value interface{}
		LineColumn
	}
	ArrayType struct {
		StringExpr string
		// it can be asserted as BasicType: int、bool、
		// PointerType: *string、*User、
		// MapType: map[${BasicType}]interface、
		// ArrayType:[]int、[]User、[]*User
		// InterfaceType: interface{}
		// Type
		ArrayType interface{}
		LineColumn
	}
	InterfaceType struct {
		StringExpr string
		// do nothing,just for assert
		LineColumn
	}
	TimeType struct {
		StringExpr string
		LineColumn
	}
	StructType struct {
		StringExpr string
		LineColumn
	}
)

func (spec *ApiSpec) ContainsTime() bool {
	for _, item := range spec.Types {
		members := item.Members
		for _, member := range members {
			if _, ok := member.Expr.(*TimeType); ok {
				return true
			}
		}
	}
	return false
}
