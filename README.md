# sqlmorph

[![Join the chat at https://gitter.im/s2gatev/sqlmorph](https://badges.gitter.im/s2gatev/sqlmorph.svg)](https://gitter.im/s2gatev/sqlmorph?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Build Status](https://travis-ci.org/s2gatev/sqlmorph.svg?branch=master)](https://travis-ci.org/s2gatev/sqlmorph)
[![Coverage Status](https://coveralls.io/repos/s2gatev/sqlmorph/badge.svg?branch=master&service=github)](https://coveralls.io/github/s2gatev/sqlmorph?branch=master)
[![Go Report Card](http://goreportcard.com/badge/s2gatev/sqlmorph)](http://goreportcard.com/report/s2gatev/sqlmorph)
[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

**sqlmorph** is an SQL AST parser and generator.

```go
package example_test

import (
  "strings"

  "github.com/s2gatev/sqlmorph/parsing"
)

func ExampleSelectParsing() {
  query := `SELECT u.Name FROM User u WHERE u.Age=21`
  ast := parsing.NewParser(strings.NewReader(query)).Parse()

  fmt.Println("%+v", ast)

  // Output:
  // &Select{
  //   Fields: []*Field{
  //     &Field{Target: "u", Name: "Name"},
  //   },
  //   Conditions: []*EqualsCondition{
  //     &EqualsCondition{
  //       Field: &Field{Target: "u", Name: "Age"},
  //       Value: "21",
  //     },
  //   },
  //   Table: &Table{Name: "User", Alias: "u"},
  // }
```
