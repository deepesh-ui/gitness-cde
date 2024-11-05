// Copyright 2023 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package enum

import (
	gitenum "github.com/harness/gitness/git/enum"
)

// PullReqState defines pull request state.
type PullReqState string

func (PullReqState) Enum() []interface{}                  { return toInterfaceSlice(pullReqStates) }
func (s PullReqState) Sanitize() (PullReqState, bool)     { return Sanitize(s, GetAllPullReqStates) }
func GetAllPullReqStates() ([]PullReqState, PullReqState) { return pullReqStates, "" }

// PullReqState enumeration.
const (
	PullReqStateOpen   PullReqState = "open"
	PullReqStateMerged PullReqState = "merged"
	PullReqStateClosed PullReqState = "closed"
)

var pullReqStates = sortEnum([]PullReqState{
	PullReqStateOpen,
	PullReqStateMerged,
	PullReqStateClosed,
})

// PullReqSort defines pull request attribute that can be used for sorting.
type PullReqSort string

func (PullReqSort) Enum() []interface{}                { return toInterfaceSlice(pullReqSorts) }
func (s PullReqSort) Sanitize() (PullReqSort, bool)    { return Sanitize(s, GetAllPullReqSorts) }
func GetAllPullReqSorts() ([]PullReqSort, PullReqSort) { return pullReqSorts, PullReqSortNumber }

// PullReqSort enumeration.
const (
	PullReqSortNumber  = "number"
	PullReqSortCreated = "created"
	PullReqSortEdited  = "edited"
	PullReqSortMerged  = "merged"
	PullReqSortUpdated = "updated"
)

var pullReqSorts = sortEnum([]PullReqSort{
	PullReqSortNumber,
	PullReqSortCreated,
	PullReqSortEdited,
	PullReqSortMerged,
	PullReqSortUpdated,
})

// PullReqActivityType defines pull request activity message type.
// Essentially, the Type determines the structure of the pull request activity's Payload structure.
type PullReqActivityType string

func (PullReqActivityType) Enum() []interface{} { return toInterfaceSlice(pullReqActivityTypes) }

func (t PullReqActivityType) Sanitize() (PullReqActivityType, bool) {
	return Sanitize(t, GetAllPullReqActivityTypes)
}

func GetAllPullReqActivityTypes() ([]PullReqActivityType, PullReqActivityType) {
	return pullReqActivityTypes, "" // No default value
}

// PullReqActivityType enumeration.
const (
	PullReqActivityTypeComment        PullReqActivityType = "comment"
	PullReqActivityTypeCodeComment    PullReqActivityType = "code-comment"
	PullReqActivityTypeTitleChange    PullReqActivityType = "title-change"
	PullReqActivityTypeStateChange    PullReqActivityType = "state-change"
	PullReqActivityTypeReviewSubmit   PullReqActivityType = "review-submit"
	PullReqActivityTypeReviewerAdd    PullReqActivityType = "reviewer-add"
	PullReqActivityTypeReviewerDelete PullReqActivityType = "reviewer-delete"
	PullReqActivityTypeBranchUpdate   PullReqActivityType = "branch-update"
	PullReqActivityTypeBranchDelete   PullReqActivityType = "branch-delete"
	PullReqActivityTypeBranchRestore  PullReqActivityType = "branch-restore"
	PullReqActivityTypeMerge          PullReqActivityType = "merge"
	PullReqActivityTypeLabelModify    PullReqActivityType = "label-modify"
)

var pullReqActivityTypes = sortEnum([]PullReqActivityType{
	PullReqActivityTypeComment,
	PullReqActivityTypeCodeComment,
	PullReqActivityTypeTitleChange,
	PullReqActivityTypeStateChange,
	PullReqActivityTypeReviewSubmit,
	PullReqActivityTypeReviewerAdd,
	PullReqActivityTypeReviewerDelete,
	PullReqActivityTypeBranchUpdate,
	PullReqActivityTypeBranchDelete,
	PullReqActivityTypeBranchRestore,
	PullReqActivityTypeMerge,
	PullReqActivityTypeLabelModify,
})

// PullReqActivityKind defines kind of pull request activity system message.
// Kind defines the source of the pull request activity entry:
// Whether it's generated by the system, it's a user comment or a part of code review.
type PullReqActivityKind string

func (PullReqActivityKind) Enum() []interface{} { return toInterfaceSlice(pullReqActivityKinds) }

func (k PullReqActivityKind) Sanitize() (PullReqActivityKind, bool) {
	return Sanitize(k, GetAllPullReqActivityKinds)
}

func GetAllPullReqActivityKinds() ([]PullReqActivityKind, PullReqActivityKind) {
	return pullReqActivityKinds, "" // No default value
}

// PullReqActivityKind enumeration.
const (
	PullReqActivityKindSystem        PullReqActivityKind = "system"
	PullReqActivityKindComment       PullReqActivityKind = "comment"
	PullReqActivityKindChangeComment PullReqActivityKind = "change-comment"
)

var pullReqActivityKinds = sortEnum([]PullReqActivityKind{
	PullReqActivityKindSystem,
	PullReqActivityKindComment,
	PullReqActivityKindChangeComment,
})

// PullReqCommentStatus defines status of a pull request comment.
type PullReqCommentStatus string

func (PullReqCommentStatus) Enum() []interface{} { return toInterfaceSlice(pullReqCommentStatuses) }

func (s PullReqCommentStatus) Sanitize() (PullReqCommentStatus, bool) {
	return Sanitize(s, GetAllPullReqCommentStatuses)
}

func GetAllPullReqCommentStatuses() ([]PullReqCommentStatus, PullReqCommentStatus) {
	return pullReqCommentStatuses, "" // No default value
}

// PullReqCommentStatus enumeration.
const (
	PullReqCommentStatusActive   PullReqCommentStatus = "active"
	PullReqCommentStatusResolved PullReqCommentStatus = "resolved"
)

var pullReqCommentStatuses = sortEnum([]PullReqCommentStatus{
	PullReqCommentStatusActive,
	PullReqCommentStatusResolved,
})

// PullReqReviewDecision defines state of a pull request review.
type PullReqReviewDecision string

func (PullReqReviewDecision) Enum() []interface{} {
	return toInterfaceSlice(pullReqReviewDecisions)
}

func (decision PullReqReviewDecision) Sanitize() (PullReqReviewDecision, bool) {
	return Sanitize(decision, GetAllPullReqReviewDecisions)
}

func GetAllPullReqReviewDecisions() ([]PullReqReviewDecision, PullReqReviewDecision) {
	return pullReqReviewDecisions, "" // No default value
}

// PullReqReviewDecision enumeration.
const (
	PullReqReviewDecisionPending   PullReqReviewDecision = "pending"
	PullReqReviewDecisionReviewed  PullReqReviewDecision = "reviewed"
	PullReqReviewDecisionApproved  PullReqReviewDecision = "approved"
	PullReqReviewDecisionChangeReq PullReqReviewDecision = "changereq"
)

var pullReqReviewDecisions = sortEnum([]PullReqReviewDecision{
	PullReqReviewDecisionPending,
	PullReqReviewDecisionReviewed,
	PullReqReviewDecisionApproved,
	PullReqReviewDecisionChangeReq,
})

// PullReqReviewerType defines type of a pull request reviewer.
type PullReqReviewerType string

func (PullReqReviewerType) Enum() []interface{} { return toInterfaceSlice(pullReqReviewerTypes) }

func (reviewerType PullReqReviewerType) Sanitize() (PullReqReviewerType, bool) {
	return Sanitize(reviewerType, GetAllPullReqReviewerTypes)
}

func GetAllPullReqReviewerTypes() ([]PullReqReviewerType, PullReqReviewerType) {
	return pullReqReviewerTypes, "" // No default value
}

// PullReqReviewerType enumeration.
const (
	PullReqReviewerTypeRequested    PullReqReviewerType = "requested"
	PullReqReviewerTypeAssigned     PullReqReviewerType = "assigned"
	PullReqReviewerTypeSelfAssigned PullReqReviewerType = "self_assigned"
)

var pullReqReviewerTypes = sortEnum([]PullReqReviewerType{
	PullReqReviewerTypeRequested,
	PullReqReviewerTypeAssigned,
	PullReqReviewerTypeSelfAssigned,
})

type MergeMethod gitenum.MergeMethod

// MergeMethod enumeration.
const (
	MergeMethodMerge       = MergeMethod(gitenum.MergeMethodMerge)
	MergeMethodSquash      = MergeMethod(gitenum.MergeMethodSquash)
	MergeMethodRebase      = MergeMethod(gitenum.MergeMethodRebase)
	MergeMethodFastForward = MergeMethod(gitenum.MergeMethodFastForward)
)

var MergeMethods = sortEnum([]MergeMethod{
	MergeMethodMerge,
	MergeMethodSquash,
	MergeMethodRebase,
	MergeMethodFastForward,
})

func (MergeMethod) Enum() []interface{} { return toInterfaceSlice(MergeMethods) }
func (m MergeMethod) Sanitize() (MergeMethod, bool) {
	s, ok := gitenum.MergeMethod(m).Sanitize()
	return MergeMethod(s), ok
}

type MergeCheckStatus string

const (
	// MergeCheckStatusUnchecked merge status has not been checked.
	MergeCheckStatusUnchecked MergeCheckStatus = "unchecked"
	// MergeCheckStatusConflict can’t merge into the target branch due to a potential conflict.
	MergeCheckStatusConflict MergeCheckStatus = "conflict"
	// MergeCheckStatusMergeable branch can merged cleanly into the target branch.
	MergeCheckStatusMergeable MergeCheckStatus = "mergeable"
)

type PullReqLabelActivityType string

func (PullReqLabelActivityType) Enum() []interface{} { return toInterfaceSlice(LabelActivityTypes) }
func (t PullReqLabelActivityType) Sanitize() (PullReqLabelActivityType, bool) {
	return Sanitize(t, GetAllLabelActivityTypes)
}
func GetAllLabelActivityTypes() ([]PullReqLabelActivityType, PullReqLabelActivityType) {
	return LabelActivityTypes, LabelActivityNoop
}

const (
	LabelActivityAssign   PullReqLabelActivityType = "assign"
	LabelActivityUnassign PullReqLabelActivityType = "unassign"
	LabelActivityReassign PullReqLabelActivityType = "reassign"
	LabelActivityNoop     PullReqLabelActivityType = "noop"
)

var LabelActivityTypes = sortEnum([]PullReqLabelActivityType{
	LabelActivityAssign,
	LabelActivityUnassign,
	LabelActivityReassign,
	LabelActivityNoop,
})
