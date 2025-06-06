// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

func (s *Server) cutPrefix(path string) (string, bool) {
	prefix := s.cfg.Prefix
	if prefix == "" {
		return path, true
	}
	if !strings.HasPrefix(path, prefix) {
		// Prefix doesn't match.
		return "", false
	}
	// Cut prefix from the path.
	return strings.TrimPrefix(path, prefix), true
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}

	elem, ok := s.cutPrefix(elem)
	if !ok || len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [2]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"

			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'l': // Prefix: "login"

				if l := len("login"); len(elem) >= l && elem[0:l] == "login" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "POST":
						s.handleLoginPostRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "POST")
					}

					return
				}

			case 't': // Prefix: "task"

				if l := len("task"); len(elem) >= l && elem[0:l] == "task" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case '-': // Prefix: "-groups"

					if l := len("-groups"); len(elem) >= l && elem[0:l] == "-groups" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleTaskGroupsGetRequest([0]string{}, elemIsEscaped, w, r)
						case "POST":
							s.handleTaskGroupsPostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET,POST")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/"

						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "groupId"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							switch r.Method {
							case "DELETE":
								s.handleTaskGroupsGroupIdDeleteRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							case "GET":
								s.handleTaskGroupsGroupIdGetRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "DELETE,GET")
							}

							return
						}
						switch elem[0] {
						case '/': // Prefix: "/"

							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 't': // Prefix: "tasks"

								if l := len("tasks"); len(elem) >= l && elem[0:l] == "tasks" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleTaskGroupsGroupIdTasksGetRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									case "POST":
										s.handleTaskGroupsGroupIdTasksPostRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, "GET,POST")
									}

									return
								}

							case 'u': // Prefix: "users"

								if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch r.Method {
									case "POST":
										s.handleTaskGroupsGroupIdUsersPostRequest([1]string{
											args[0],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, "POST")
									}

									return
								}
								switch elem[0] {
								case '/': // Prefix: "/"

									if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
										elem = elem[l:]
									} else {
										break
									}

									// Param: "userId"
									// Leaf parameter, slashes are prohibited
									idx := strings.IndexByte(elem, '/')
									if idx >= 0 {
										break
									}
									args[1] = elem
									elem = ""

									if len(elem) == 0 {
										// Leaf node.
										switch r.Method {
										case "DELETE":
											s.handleTaskGroupsGroupIdUsersUserIdDeleteRequest([2]string{
												args[0],
												args[1],
											}, elemIsEscaped, w, r)
										default:
											s.notAllowed(w, r, "DELETE")
										}

										return
									}

								}

							}

						}

					}

				case 's': // Prefix: "s/"

					if l := len("s/"); len(elem) >= l && elem[0:l] == "s/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "taskId"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch r.Method {
						case "DELETE":
							s.handleTasksTaskIdDeleteRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "GET":
							s.handleTasksTaskIdGetRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "PUT":
							s.handleTasksTaskIdPutRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "DELETE,GET,PUT")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/status"

						if l := len("/status"); len(elem) >= l && elem[0:l] == "/status" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "PATCH":
								s.handleTasksTaskIdStatusPatchRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "PATCH")
							}

							return
						}

					}

				}

			case 'u': // Prefix: "users"

				if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "POST":
						s.handleUsersPostRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"

					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "userId"
					// Leaf parameter, slashes are prohibited
					idx := strings.IndexByte(elem, '/')
					if idx >= 0 {
						break
					}
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "DELETE":
							s.handleUsersUserIdDeleteRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "DELETE")
						}

						return
					}

				}

			}

		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	summary     string
	operationID string
	pathPattern string
	count       int
	args        [2]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// Summary returns OpenAPI summary.
func (r Route) Summary() string {
	return r.summary
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	elem, ok := s.cutPrefix(elem)
	if !ok {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"

			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'l': // Prefix: "login"

				if l := len("login"); len(elem) >= l && elem[0:l] == "login" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch method {
					case "POST":
						r.name = LoginPostOperation
						r.summary = "ログイン"
						r.operationID = ""
						r.pathPattern = "/login"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}

			case 't': // Prefix: "task"

				if l := len("task"); len(elem) >= l && elem[0:l] == "task" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case '-': // Prefix: "-groups"

					if l := len("-groups"); len(elem) >= l && elem[0:l] == "-groups" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = TaskGroupsGetOperation
							r.summary = "所属グループ一覧取得"
							r.operationID = ""
							r.pathPattern = "/task-groups"
							r.args = args
							r.count = 0
							return r, true
						case "POST":
							r.name = TaskGroupsPostOperation
							r.summary = "タスクグループ作成"
							r.operationID = ""
							r.pathPattern = "/task-groups"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/"

						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "groupId"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							switch method {
							case "DELETE":
								r.name = TaskGroupsGroupIdDeleteOperation
								r.summary = "グループ削除"
								r.operationID = ""
								r.pathPattern = "/task-groups/{groupId}"
								r.args = args
								r.count = 1
								return r, true
							case "GET":
								r.name = TaskGroupsGroupIdGetOperation
								r.summary = "グループ詳細取得"
								r.operationID = ""
								r.pathPattern = "/task-groups/{groupId}"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
						switch elem[0] {
						case '/': // Prefix: "/"

							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case 't': // Prefix: "tasks"

								if l := len("tasks"); len(elem) >= l && elem[0:l] == "tasks" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = TaskGroupsGroupIdTasksGetOperation
										r.summary = "タスク一覧取得"
										r.operationID = ""
										r.pathPattern = "/task-groups/{groupId}/tasks"
										r.args = args
										r.count = 1
										return r, true
									case "POST":
										r.name = TaskGroupsGroupIdTasksPostOperation
										r.summary = "タスク作成"
										r.operationID = ""
										r.pathPattern = "/task-groups/{groupId}/tasks"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}

							case 'u': // Prefix: "users"

								if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									switch method {
									case "POST":
										r.name = TaskGroupsGroupIdUsersPostOperation
										r.summary = "ユーザ招待"
										r.operationID = ""
										r.pathPattern = "/task-groups/{groupId}/users"
										r.args = args
										r.count = 1
										return r, true
									default:
										return
									}
								}
								switch elem[0] {
								case '/': // Prefix: "/"

									if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
										elem = elem[l:]
									} else {
										break
									}

									// Param: "userId"
									// Leaf parameter, slashes are prohibited
									idx := strings.IndexByte(elem, '/')
									if idx >= 0 {
										break
									}
									args[1] = elem
									elem = ""

									if len(elem) == 0 {
										// Leaf node.
										switch method {
										case "DELETE":
											r.name = TaskGroupsGroupIdUsersUserIdDeleteOperation
											r.summary = "グループからユーザを削除"
											r.operationID = ""
											r.pathPattern = "/task-groups/{groupId}/users/{userId}"
											r.args = args
											r.count = 2
											return r, true
										default:
											return
										}
									}

								}

							}

						}

					}

				case 's': // Prefix: "s/"

					if l := len("s/"); len(elem) >= l && elem[0:l] == "s/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "taskId"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch method {
						case "DELETE":
							r.name = TasksTaskIdDeleteOperation
							r.summary = "タスク削除"
							r.operationID = ""
							r.pathPattern = "/tasks/{taskId}"
							r.args = args
							r.count = 1
							return r, true
						case "GET":
							r.name = TasksTaskIdGetOperation
							r.summary = "タスク詳細取得"
							r.operationID = ""
							r.pathPattern = "/tasks/{taskId}"
							r.args = args
							r.count = 1
							return r, true
						case "PUT":
							r.name = TasksTaskIdPutOperation
							r.summary = "タスク更新"
							r.operationID = ""
							r.pathPattern = "/tasks/{taskId}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/status"

						if l := len("/status"); len(elem) >= l && elem[0:l] == "/status" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "PATCH":
								r.name = TasksTaskIdStatusPatchOperation
								r.summary = "タスクステータス変更"
								r.operationID = ""
								r.pathPattern = "/tasks/{taskId}/status"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}

					}

				}

			case 'u': // Prefix: "users"

				if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "POST":
						r.name = UsersPostOperation
						r.summary = "ユーザ登録"
						r.operationID = ""
						r.pathPattern = "/users"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"

					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "userId"
					// Leaf parameter, slashes are prohibited
					idx := strings.IndexByte(elem, '/')
					if idx >= 0 {
						break
					}
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "DELETE":
							r.name = UsersUserIdDeleteOperation
							r.summary = "ユーザ削除"
							r.operationID = ""
							r.pathPattern = "/users/{userId}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}

				}

			}

		}
	}
	return r, false
}
