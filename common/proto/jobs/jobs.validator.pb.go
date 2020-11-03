// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: jobs.proto

package jobs

import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"
import _ "github.com/pydio/cells/common/proto/activity"
import _ "github.com/pydio/cells/common/proto/idm"
import _ "github.com/pydio/cells/common/proto/tree"
import _ "github.com/pydio/cells/common/service/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *NodesSelector) Validate() error {
	if this.Query != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Query); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Query", err)
		}
	}
	return nil
}
func (this *IdmSelector) Validate() error {
	if this.Query != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Query); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Query", err)
		}
	}
	return nil
}
func (this *UsersSelector) Validate() error {
	for _, item := range this.Users {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Users", err)
			}
		}
	}
	if this.Query != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Query); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Query", err)
		}
	}
	return nil
}
func (this *ActionOutputFilter) Validate() error {
	if this.Query != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Query); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Query", err)
		}
	}
	return nil
}
func (this *ContextMetaFilter) Validate() error {
	if this.Query != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Query); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Query", err)
		}
	}
	return nil
}
func (this *ContextMetaSingleQuery) Validate() error {
	if this.Condition != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Condition); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Condition", err)
		}
	}
	return nil
}
func (this *Schedule) Validate() error {
	return nil
}
func (this *Action) Validate() error {
	if this.NodesSelector != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.NodesSelector); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("NodesSelector", err)
		}
	}
	if this.UsersSelector != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UsersSelector); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UsersSelector", err)
		}
	}
	if this.NodesFilter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.NodesFilter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("NodesFilter", err)
		}
	}
	if this.UsersFilter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UsersFilter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UsersFilter", err)
		}
	}
	if this.IdmSelector != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.IdmSelector); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("IdmSelector", err)
		}
	}
	if this.IdmFilter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.IdmFilter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("IdmFilter", err)
		}
	}
	if this.ActionOutputFilter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ActionOutputFilter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ActionOutputFilter", err)
		}
	}
	if this.ContextMetaFilter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ContextMetaFilter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ContextMetaFilter", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	for _, item := range this.ChainedActions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ChainedActions", err)
			}
		}
	}
	for _, item := range this.FailedFilterActions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("FailedFilterActions", err)
			}
		}
	}
	return nil
}
func (this *Job) Validate() error {
	if this.Schedule != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Schedule); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Schedule", err)
		}
	}
	for _, item := range this.Actions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Actions", err)
			}
		}
	}
	for _, item := range this.Tasks {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Tasks", err)
			}
		}
	}
	if this.NodeEventFilter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.NodeEventFilter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("NodeEventFilter", err)
		}
	}
	if this.UserEventFilter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UserEventFilter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UserEventFilter", err)
		}
	}
	if this.IdmFilter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.IdmFilter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("IdmFilter", err)
		}
	}
	if this.ContextMetaFilter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ContextMetaFilter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ContextMetaFilter", err)
		}
	}
	for _, item := range this.Parameters {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Parameters", err)
			}
		}
	}
	return nil
}
func (this *JobParameter) Validate() error {
	return nil
}
func (this *JobChangeEvent) Validate() error {
	if this.JobUpdated != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.JobUpdated); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("JobUpdated", err)
		}
	}
	return nil
}
func (this *TaskChangeEvent) Validate() error {
	if this.TaskUpdated != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TaskUpdated); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TaskUpdated", err)
		}
	}
	if this.Job != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Job); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Job", err)
		}
	}
	return nil
}
func (this *PutJobRequest) Validate() error {
	if this.Job != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Job); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Job", err)
		}
	}
	return nil
}
func (this *PutJobResponse) Validate() error {
	if this.Job != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Job); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Job", err)
		}
	}
	return nil
}
func (this *GetJobRequest) Validate() error {
	return nil
}
func (this *GetJobResponse) Validate() error {
	if this.Job != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Job); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Job", err)
		}
	}
	return nil
}
func (this *DeleteJobRequest) Validate() error {
	return nil
}
func (this *DeleteJobResponse) Validate() error {
	return nil
}
func (this *ListJobsRequest) Validate() error {
	return nil
}
func (this *ListJobsResponse) Validate() error {
	if this.Job != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Job); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Job", err)
		}
	}
	return nil
}
func (this *ListTasksRequest) Validate() error {
	return nil
}
func (this *ListTasksResponse) Validate() error {
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	return nil
}
func (this *PutTaskRequest) Validate() error {
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	return nil
}
func (this *PutTaskResponse) Validate() error {
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	return nil
}
func (this *DeleteTasksRequest) Validate() error {
	return nil
}
func (this *DeleteTasksResponse) Validate() error {
	return nil
}
func (this *DetectStuckTasksRequest) Validate() error {
	return nil
}
func (this *DetectStuckTasksResponse) Validate() error {
	return nil
}
func (this *Task) Validate() error {
	for _, item := range this.ActionsLogs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ActionsLogs", err)
			}
		}
	}
	return nil
}
func (this *CtrlCommand) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *CtrlCommandResponse) Validate() error {
	return nil
}
func (this *ActionLog) Validate() error {
	if this.Action != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Action); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Action", err)
		}
	}
	if this.InputMessage != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.InputMessage); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("InputMessage", err)
		}
	}
	if this.OutputMessage != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.OutputMessage); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("OutputMessage", err)
		}
	}
	return nil
}
func (this *JobTriggerEvent) Validate() error {
	if this.Schedule != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Schedule); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Schedule", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *ActionOutput) Validate() error {
	return nil
}
func (this *ActionOutputSingleQuery) Validate() error {
	return nil
}
func (this *ActionMessage) Validate() error {
	if this.Event != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Event); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Event", err)
		}
	}
	for _, item := range this.Nodes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Nodes", err)
			}
		}
	}
	for _, item := range this.Users {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Users", err)
			}
		}
	}
	for _, item := range this.Roles {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Roles", err)
			}
		}
	}
	for _, item := range this.Workspaces {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Workspaces", err)
			}
		}
	}
	for _, item := range this.Acls {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Acls", err)
			}
		}
	}
	for _, item := range this.Activities {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Activities", err)
			}
		}
	}
	for _, item := range this.OutputChain {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("OutputChain", err)
			}
		}
	}
	return nil
}
