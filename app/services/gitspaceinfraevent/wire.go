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

package gitspaceinfraevent

import (
	"context"

	gitspaceevents "github.com/harness/gitness/app/events/gitspace"
	gitspaceinfraevents "github.com/harness/gitness/app/events/gitspaceinfra"
	"github.com/harness/gitness/app/gitspace/orchestrator"
	"github.com/harness/gitness/app/services/gitspace"
	"github.com/harness/gitness/app/services/gitspaceevent"
	"github.com/harness/gitness/events"

	"github.com/google/wire"
)

// WireSet provides a wire set for this package.
var WireSet = wire.NewSet(
	ProvideService,
)

func ProvideService(
	ctx context.Context,
	config *gitspaceevent.Config,
	gitspaceInfraEventReaderFactory *events.ReaderFactory[*gitspaceinfraevents.Reader],
	orchestrator orchestrator.Orchestrator,
	gitspaceSvc *gitspace.Service,
	eventReporter *gitspaceevents.Reporter,
) (*Service, error) {
	return NewService(
		ctx,
		config,
		gitspaceInfraEventReaderFactory,
		orchestrator,
		gitspaceSvc,
		eventReporter,
	)
}
