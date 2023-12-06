// Copyright 2017 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spanner

import (
	"context"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

const statsPrefix = "cloud.google.com/go/spanner/"

var (
	tagClientID   = tag.MustNewKey("client_id")
	tagDatabase   = tag.MustNewKey("database")
	tagInstance   = tag.MustNewKey("instance_id")
	tagLibVersion = tag.MustNewKey("library_version")
	tagAllKeys    = []tag.Key{tagClientID, tagDatabase, tagInstance, tagLibVersion}
)

func recordStat(ctx context.Context, m *stats.Int64Measure, n int64) {
	stats.Record(ctx, m.M(n))
}

var (
	// OpenSessionCount is a measure of the number of sessions currently opened.
	// It is EXPERIMENTAL and subject to change or removal without notice.
	OpenSessionCount = stats.Int64(
		statsPrefix+"open_session_count",
		"Number of sessions currently opened",
		stats.UnitDimensionless,
	)

	// OpenSessionCountView is a view of the last value of OpenSessionCount.
	// It is EXPERIMENTAL and subject to change or removal without notice.
	OpenSessionCountView = &view.View{
		Measure:     OpenSessionCount,
		Aggregation: view.LastValue(),
		TagKeys:     tagAllKeys,
	}

	// MaxAllowedSessionsCount is a measure of the maximum number of sessions
	// allowed. Configurable by the user.
	MaxAllowedSessionsCount = stats.Int64(
		statsPrefix+"max_allowed_sessions",
		"The maximum number of sessions allowed. Configurable by the user.",
		stats.UnitDimensionless,
	)

	// MaxAllowedSessionsCountView is a view of the last value of
	// MaxAllowedSessionsCount.
	MaxAllowedSessionsCountView = &view.View{
		Measure:     MaxAllowedSessionsCount,
		Aggregation: view.LastValue(),
		TagKeys:     tagAllKeys,
	}

	// InUseSessionsCount is a measure of the number of sessions currently
	// opened.
	InUseSessionsCount = stats.Int64(
		statsPrefix+"in_use_sessions",
		"The number of sessions currently in use.",
		stats.UnitDimensionless,
	)

	// InUseSessionsCountView is a view of the last value of
	// InUseSessionsCount.
	InUseSessionsCountView = &view.View{
		Measure:     InUseSessionsCount,
		Aggregation: view.LastValue(),
		TagKeys:     tagAllKeys,
	}

	// MaxInUseSessionsCount is a measure of the maximum number of sessions
	// in use during the last 10 minute interval.
	MaxInUseSessionsCount = stats.Int64(
		statsPrefix+"max_in_use_sessions",
		"The maximum number of sessions in use during the last 10 minute interval.",
		stats.UnitDimensionless,
	)

	// MaxInUseSessionsCountView is a view of the last value of
	// MaxInUseSessionsCount.
	MaxInUseSessionsCountView = &view.View{
		Measure:     MaxInUseSessionsCount,
		Aggregation: view.LastValue(),
		TagKeys:     tagAllKeys,
	}

	// GetSessionTimeoutsCount is a measure of the number of get sessions
	// timeouts due to pool exhaustion.
	GetSessionTimeoutsCount = stats.Int64(
		statsPrefix+"get_session_timeouts",
		"The number of get sessions timeouts due to pool exhaustion.",
		stats.UnitDimensionless,
	)

	// GetSessionTimeoutsCountView is a view of the last value of
	// GetSessionTimeoutsCount.
	GetSessionTimeoutsCountView = &view.View{
		Measure:     GetSessionTimeoutsCount,
		Aggregation: view.Count(),
		TagKeys:     tagAllKeys,
	}

	// AcquiredSessionsCount is the number of sessions acquired from
	// the session pool.
	AcquiredSessionsCount = stats.Int64(
		statsPrefix+"num_acquired_sessions",
		"The number of sessions acquired from the session pool.",
		stats.UnitDimensionless,
	)

	// AcquiredSessionsCountView is a view of the last value of
	// AcquiredSessionsCount.
	AcquiredSessionsCountView = &view.View{
		Measure:     AcquiredSessionsCount,
		Aggregation: view.Count(),
		TagKeys:     tagAllKeys,
	}

	// ReleasedSessionsCount is the number of sessions released by the user
	// and pool maintainer.
	ReleasedSessionsCount = stats.Int64(
		statsPrefix+"num_released_sessions",
		"The number of sessions released by the user and pool maintainer.",
		stats.UnitDimensionless,
	)

	// ReleasedSessionsCountView is a view of the last value of
	// ReleasedSessionsCount.
	ReleasedSessionsCountView = &view.View{
		Measure:     ReleasedSessionsCount,
		Aggregation: view.Count(),
		TagKeys:     tagAllKeys,
	}
)

// EnableStatViews enables all views of metrics relate to session management.
func EnableStatViews() error {
	return view.Register(
		OpenSessionCountView,
		MaxAllowedSessionsCountView,
		InUseSessionsCountView,
		MaxInUseSessionsCountView,
		GetSessionTimeoutsCountView,
		AcquiredSessionsCountView,
		ReleasedSessionsCountView,
	)
}
