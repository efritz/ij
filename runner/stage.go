package runner

import (
	"fmt"

	"github.com/efritz/ij/config"
	"github.com/efritz/ij/environment"
	"github.com/efritz/ij/logging"
	"github.com/efritz/ij/state"
	"github.com/efritz/ij/task/run"
	"github.com/efritz/ij/util"
)

type (
	StageRunner struct {
		state  *state.State
		plan   *config.Plan
		stage  *config.Stage
		prefix *logging.Prefix
	}

	RunnerFunc func() bool
)

func NewStageRunner(
	state *state.State,
	plan *config.Plan,
	stage *config.Stage,
	prefix *logging.Prefix,
) *StageRunner {
	return &StageRunner{
		state:  state,
		plan:   plan,
		stage:  stage,
		prefix: prefix,
	}
}

func (r *StageRunner) Run() bool {
	r.state.Logger.Info(
		r.prefix,
		"Beginning stage",
	)

	runners := []RunnerFunc{}
	for i, stageTask := range r.stage.Tasks {
		runners = append(runners, r.buildRunner(
			stageTask,
			i,
			r.state.Config.Tasks[stageTask.Name],
		))
	}

	if !r.stage.Parallel || r.state.ForceSequential {
		return runSequential(runners)
	}

	return runParallel(runners)
}

func (r *StageRunner) buildRunner(
	stageTask *config.StageTask,
	index int,
	task *config.Task,
) RunnerFunc {
	taskPrefix := r.prefix.Append(fmt.Sprintf(
		"%s.%d",
		task.Name,
		index,
	))

	return func() bool {
		env := environment.Merge(
			environment.New(r.state.Config.Environment),
			environment.New(task.Environment),
			environment.New(r.plan.Environment),
			environment.New(r.stage.Environment),
			environment.New(stageTask.Environment),
			environment.New(r.state.GetExportedEnv()),
			environment.New(r.state.Env),
		)

		runner := run.NewRunner(
			r.state,
			task,
			taskPrefix,
			env,
		)

		if !runner.Run() {
			r.state.ReportError(
				taskPrefix,
				"Task has failed",
			)

			return false
		}

		r.state.Logger.Info(
			taskPrefix,
			"Task has completed successfully",
		)

		return true
	}
}

//
// Helpers

func runSequential(runners []RunnerFunc) bool {
	for _, runner := range runners {
		if !runner() {
			return false
		}
	}

	return true
}

func runParallel(runners []RunnerFunc) bool {
	failures := make(chan bool, len(runners))
	defer close(failures)

	funcs := []func(){}
	for _, runner := range runners {
		funcs = append(funcs, func() {
			if ok := runner(); !ok {
				failures <- false
			}
		})
	}

	util.RunParallel(funcs...)

	select {
	case <-failures:
		return false
	default:
	}

	return true
}