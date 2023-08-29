// Copyright (c) 2023 Symbol Not Found L.L.C.
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
//
// github:SymbolNotFound/ggp-site/gamewiki/gameinfo.go

package gamewiki

type GamesList struct {
	Games []GameInfo
}

// Represents metadata and derived data about a single game.  Includes solo
// and multiplayer games, deterministic and nondeterministic, turn-based and
// real-time, table
type GameInfo struct {
	// Identifies the game; limited URL-safe characters.
	GameId string `json:"id"`

	// The displayed name of the game.
	Name string `json:"name"`

	// A string representation of the number of players allowed by the game rules.
	RoleCount string `json:"players"`

	// A brief description of the game (longer descriptions found in STORY.md).
	Summary string `json:"summary"`

	// Reflects the playability of the game and its associated rulesheet & UI.
	Playable GameStatus `json:"status"`

	// A list of tags/labels assigned to this game.
	Tags []string `json:"tags"`

	// The description part of an external reference link.
	RefDesc string `json:"ref_desc"`

	// The URL pointed to by the external reference link.
	RefLink string `json:"ref_link"`

	// Local path where the game's rule definition can be found (if it exists).
	Rulesheet string `json:"rulesheet,omitempty"`

	// Additional notes I was making when first compiling this list.
	ExtraNotes []string `json:"notes,omitempty"`

	// The measurable features of this game.
	Complexity GameComplexity `json:"complexity,omitempty"`

	// The measurable features of this game.
	LegacyComplexity GameComplexity `json:"measures,omitempty"`

	// Comparative values for this game relative to other games.
	LegacyMeasures GameMeasures `json:"scores,omitempty"`

	// Comparative values for this game relative to other games.
	Measures GameMeasures `json:"measures,omitempty"`
}

type GameStatus string

// These definitions are self-documenting.
const (
	UNKNOWN      GameStatus = ""
	DOCUMENTED   GameStatus = "documented, but not implemented"
	IN_DEVELOP   GameStatus = "currently in development"
	ALPHA_DEV    GameStatus = "ALPHA (development + testing)"
	BETA_TEST    GameStatus = "BETA (feature complete, in beta testing)"
	PRIVATE_ONLY GameStatus = "HIDDEN"
	LAUNCHED     GameStatus = "LAUNCHED!"
	ARCHIVED     GameStatus = "[archived]"
)

// Game Complexity is a game design measurement, derived from game properties
// (e.g. how deep & dense the game tree is, how much of the board is used in a
// game, percentage of pieces in play, etc.).  These measures can sometimes be
// measured directly from the rules but many of these measures require a
// simulation of many playthroughs to sufficiently approximate an estimate.
type GameComplexity struct {
	BoardUsage           float32 `json:"Board Usage"`
	MinMatchLength       uint    `json:"min(Match Length)"`
	MaxMatchLength       uint    `json:"max(Match Length)"`
	RuleComplexity       float32 `json:"Rule Complexity"`
	EquipmentComplexity  float32 `json:"Equipment Complexity"`
	BranchingFactor      float32 `json:"Branching Factor"`
	GameTreeComplexity   float32 `json:"Game Tree Complexity"`
	GameLength           float32 `json:"Game Length"`
	StateSpaceComplexity float32 `json:"Statespace Complexity"`
	MoveSeparability     float32 `json:"Move Separability"`
	MutationalComplexity float32 `json:"Mutational Complexity"`
}

// Measurements based on the complexity scores above, attempting to model the
// qualities of a game relative to other (similar and different) games.
type GameMeasures struct {
	Simplicity float32 `json:"Simplicity"`
	Clarity    float32 `json:"Clarity"`
	Efficiency float32 `json:"Efficiency"`
	RelDepth   float32 `json:"Relative Depth"`
	Elegance   float32 `json:"Elegance"`
	Shibui     float32 `json:"Shibui"`
}
