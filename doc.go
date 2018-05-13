// Package gorogue is a simple roguelike engine build in golang.
//
// Goals For This Project:
//
// 1. Keep it simple, stupid
//
// I'm building an engine, not a game. The idea
// is to create a simple tool that any designer can extend to create their game.
// I'm not going to bloat the repository with unnecessary things like lots of items,
// or damage equations or things like that. Users will have to add those features
// per their needs.
//
// 2. Focus on versatility
//
// I want this project to be flexible. That's why I'll be supporting a wide variety
// of  "play styles" including: Turn Based, "Real time", Multi-player, and
// Squad/Party based.
//
// 3. An engine is only as good as its documentation.
//
// 4. Implement as little as possible in the base package.
//
// Use the base package as a skeleton for any game to work from. Implement things
// in subpackages so they can just as easily be extended as removed.
package gorogue
