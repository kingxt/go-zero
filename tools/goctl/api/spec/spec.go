package spec

type (
	Doc []string

	Annotation struct {
		Properties map[string]string
	}

	ApiSyntax struct {
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
		List []string
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
		Email      string
		Proterties map[string]string
	}

	Member struct {
		Annotations []Annotation
		Name        string
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
		Annotation   Annotation
		Method       string
		Path         string
		RequestType  Type
		ResponseType Type
		Docs         Doc
		Handler      string
	}

	Service struct {
		Name   string
		Groups []Group
	}

	Type struct {
		Name        string
		Annotations []Annotation
		Members     []Member
	}

	// 系统预设基本数据类型
	BasicType struct {
		StringExpr string
		Name       string
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
	}
	InterfaceType struct {
		StringExpr string
		// do nothing,just for assert
	}
	TimeType struct {
		StringExpr string
	}
	StructType struct {
		StringExpr string
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
