//! The portsmith line-based trace format (v1), Rust side.
//!
//! Mirrors the Go `internal/trace` package so both tools interoperate. See
//! `docs/FORMAT.md` for the authoritative specification.

use std::fmt;
use std::io::{self, BufRead};

