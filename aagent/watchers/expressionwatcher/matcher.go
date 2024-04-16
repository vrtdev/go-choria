// Copyright (c) 2024, R.I. Pienaar and the Choria Project contributors
//
// SPDX-License-Identifier: Apache-2.0

package expression

import (
	"fmt"

	"github.com/expr-lang/expr"
)

func (w *Watcher) evaluateExpression(e string) (bool, error) {
	if e == "" {
		return false, fmt.Errorf("invalid expression")
	}

	env := map[string]any{
		"data":     w.machine.Data(),
		"facts":    w.machine.Facts(),
		"identity": w.machine.Identity(),
	}

	execEnv := expr.Env(env)
	prog, err := expr.Compile(e, execEnv, expr.AsBool())
	if err != nil {
		return false, err
	}

	res, err := expr.Run(prog, env)
	if err != nil {
		return false, err
	}

	b, ok := res.(bool)
	if !ok {
		return false, fmt.Errorf("match was non boolean")
	}

	w.Debugf("Evaluated expression %q returned: %v", e, b)

	return b, nil
}
