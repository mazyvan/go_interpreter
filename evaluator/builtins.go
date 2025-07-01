package evaluator

import (
	"bytes"
	"fmt"
	"net/http"
	"persistio/lib"
	"persistio/object"
)

func getBuiltins(env *object.Environment) map[string]*object.Builtin {
	return map[string]*object.Builtin{
		"str": {
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got=%d, want=1",
						len(args))
				}
				switch arg := args[0].(type) {
				case *object.Integer:
					return &object.String{Value: fmt.Sprint(arg.Value)}
				case *object.String:
					return arg
				case *object.Boolean:
					if arg.Value {
						return &object.String{Value: "true"}
					}
					return &object.String{Value: "false"}
				default:
					return newError("argument to `str` not supported, got %s",
						args[0].Type())
				}
			},
		},
		"len": {
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got=%d, want=1",
						len(args))
				}
				switch arg := args[0].(type) {
				case *object.Array:
					return &object.Integer{Value: int64(len(arg.Elements))}
				case *object.String:
					return &object.Integer{Value: int64(len(arg.Value))}
				default:
					return newError("argument to `len` not supported, got %s",
						args[0].Type())
				}
			},
		},
		"first": {
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got=%d, want=1",
						len(args))
				}
				if args[0].Type() != object.ARRAY_OBJ {
					return newError("argument to `first` must be ARRAY, got %s",
						args[0].Type())
				}
				arr := args[0].(*object.Array)
				if len(arr.Elements) > 0 {
					return arr.Elements[0]
				}
				return NULL
			},
		},
		"last": {
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got=%d, want=1",
						len(args))
				}
				if args[0].Type() != object.ARRAY_OBJ {
					return newError("argument to `last` must be ARRAY, got %s",
						args[0].Type())
				}
				arr := args[0].(*object.Array)
				length := len(arr.Elements)
				if length > 0 {
					return arr.Elements[length-1]
				}
				return NULL
			},
		},
		"rest": {
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got=%d, want=1",
						len(args))
				}
				if args[0].Type() != object.ARRAY_OBJ {
					return newError("argument to `rest` must be ARRAY, got %s",
						args[0].Type())
				}
				arr := args[0].(*object.Array)
				length := len(arr.Elements)
				if length > 0 {
					newElements := make([]object.Object, length-1)
					copy(newElements, arr.Elements[1:length])
					return &object.Array{Elements: newElements}
				}
				return &object.Array{Elements: []object.Object{}}
			},
		},
		"push": {
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 2 {
					return newError("wrong number of arguments. got=%d, want=2",
						len(args))
				}
				if args[0].Type() != object.ARRAY_OBJ {
					return newError("argument to `push` must be ARRAY, got %s",
						args[0].Type())
				}
				arr := args[0].(*object.Array)
				length := len(arr.Elements)
				newElements := make([]object.Object, length+1)
				copy(newElements, arr.Elements)
				newElements[length] = args[1]
				return &object.Array{Elements: newElements}
			},
		},
		"replace": {
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 3 {
					return newError("wrong number of arguments. got=%d, want=3",
						len(args))
				}
				switch args[0].Type() {
				case object.ARRAY_OBJ:
					arr := args[0].(*object.Array)
					if idx, ok := args[1].(*object.Integer); ok {
						index := idx.Value
						if index < 0 || index >= int64(len(arr.Elements)) {
							return newError("index out of bounds: %d", index)
						}
						arr.Elements[index] = args[2]
						return NULL
					}
				case object.HASH_OBJ:
					hash := args[0].(*object.Hash)
					if key, ok := args[1].(object.Hashable); ok {
						hash.Pairs[key.HashKey()] = object.HashPair{Key: key.(object.Object), Value: args[2]}
						return NULL
					}
				default:
					return newError("argument to `replace` must be ARRAY or HASH, got %s",
						args[0].Type())
				}
				return NULL
			},
		},
		"puts": {
			Fn: func(args ...object.Object) object.Object {
				var out bytes.Buffer
				for _, arg := range args {
					out.WriteString(arg.Inspect() + " ")
				}
				fmt.Println(out.String())
				return NULL
			},
		},
		"http": {
			Fn: func(args ...object.Object) object.Object {
				pairs := make(map[object.HashKey]object.HashPair)

				registerKey := &object.String{Value: "register"}
				registerHashed := registerKey.HashKey()

				listen := &object.String{Value: "listen"}
				listenHashed := listen.HashKey()
				pairs[listenHashed] = object.HashPair{
					Key: listen,
					Value: &object.Builtin{
						Fn: func(args ...object.Object) object.Object {
							if len(args) != 1 {
								return newError("wrong number of arguments. got=%d, want=1",
									len(args))
							}
							if args[0].Type() != object.STRING_OBJ {
								return newError("argument to `listen` must be STRING, got %s",
									args[0].Type())
							}
							port := args[0].(*object.String).Value
							fmt.Printf("Listening on port: %s\n", port)
							lib.StartServer(port)
							return NULL
						},
					},
				}
				pairs[registerHashed] = object.HashPair{
					Key: registerKey,
					Value: &object.Builtin{
						Fn: func(args ...object.Object) object.Object {
							if len(args) != 2 {
								return newError("wrong number of arguments. got=%d, want=2",
									len(args))
							}
							if args[0].Type() != object.STRING_OBJ {
								return newError("argument to `register` must be STRING, got %s",
									args[0].Type())
							}
							if args[1].Type() != object.FUNCTION_OBJ {
								return newError("argument to `register` must be FUNCTION, got %s",
									args[1].Type())
							}
							function := args[1].(*object.Function)
							function.Env = env
							lib.RegisterHandler(
								args[0].(*object.String).Value,
								handlerWrapper(function),
							)
							return NULL
						},
					},
				}

				return &object.Hash{Pairs: pairs}
			},
		},
	}
}

func handlerWrapper(fn *object.Function) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		methodKey := &object.String{Value: "method"}
		pathKey := &object.String{Value: "path"}
		headersKey := &object.String{Value: "headers"}
		queryParamKey := &object.String{Value: "query"}
		request := &object.Hash{
			Pairs: map[object.HashKey]object.HashPair{
				methodKey.HashKey(): {
					Key:   methodKey,
					Value: &object.String{Value: r.Method},
				},
				pathKey.HashKey(): {
					Key:   pathKey,
					Value: &object.String{Value: r.URL.Path},
				},
				queryParamKey.HashKey(): {
					Key:   queryParamKey,
					Value: &object.String{Value: r.URL.RawQuery},
				},
				headersKey.HashKey(): {
					Key: headersKey,
					Value: &object.Hash{
						Pairs: make(map[object.HashKey]object.HashPair),
					},
				},
			},
		}
		for name, values := range r.Header {
			for _, value := range values {
				key := &object.String{Value: name}
				request.Pairs[key.HashKey()] = object.HashPair{
					Key:   key,
					Value: &object.String{Value: value},
				}
			}
		}
		statusKey := &object.String{Value: "status"}
		bodyKey := &object.String{Value: "body"}
		response := &object.Hash{
			Pairs: map[object.HashKey]object.HashPair{
				statusKey.HashKey(): {
					Key:   statusKey,
					Value: &object.Integer{Value: 200}, // Default status code
				},
				bodyKey.HashKey(): {
					Key:   bodyKey,
					Value: &object.String{Value: ""}, // Default body
				},
			},
		}
		result := applyFunction(fn, []object.Object{request, response})
		if isError(result) {
			fmt.Printf("Error in handler: %s\n", result.Inspect())
			http.Error(w, result.Inspect(), http.StatusInternalServerError)
			return
		}
		if status, ok := response.Pairs[statusKey.HashKey()].Value.(*object.Integer); ok {
			w.WriteHeader(int(status.Value))
		} else {
			http.Error(w, "Handler did not return a valid status code", http.StatusInternalServerError)
		}

		w.Write([]byte(response.Pairs[bodyKey.HashKey()].Value.(*object.String).Value))
	}
}
