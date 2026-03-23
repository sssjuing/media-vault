# Media Vault - Agent Guidelines

This document provides guidelines for agentic coding assistants working in the Media Vault repository.

## Project Overview

**Media Vault** is a full-stack media management application with:

- **Backend**: Go + Echo v5 + GORM + SQLite/PostgreSQL
- **Frontend**: React 19 + TypeScript + Vite + Tailwind CSS + TanStack Query
- **Package Manager**: pnpm (monorepo workspace)

## Build & Development Commands

### Root Commands (pnpm workspace)

```bash
pnpm dev          # Start all dev servers
pnpm lint         # Lint all packages
pnpm test         # Test all packages
pnpm build        # Build all packages
pnpm format:check # Check Prettier formatting
pnpm commit       # Conventional commits with emoji
```

### Go Backend

```bash
cd /home/user/source/media-vault

# Development
air                    # Live-reload dev server (uses .air.toml)

# Build
go build -o ./tmp/main ./cmd/server

# Testing
go test ./...                               # Run all tests
go test ./internal/app/server/repository    # Test specific package
go test ./internal/app/server/repository -run TestVideoRepository_Count  # Run single test
go test -v ./...                            # Verbose test output

# Linting
# (Use standard Go tools - no specific linter configured beyond go fmt)
```

### Frontend (web/login)

```bash
cd /home/user/source/media-vault/web/login

pnpm dev           # Vite dev server
pnpm build         # TypeScript + Vite build
pnpm lint          # ESLint
pnpm format        # Prettier write
pnpm format:check  # Prettier check
pnpm preview       # Preview production build
```

## Code Style Guidelines

### Go Code Style

**Imports**:

- Standard library first
- Third-party libraries next
- Local packages last
- Groups separated by blank lines

Example:

```go
import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v5"
	"github.com/samber/lo"

	"github.com/sssjuing/media-vault/internal/app/server/model"
	"github.com/sssjuing/media-vault/internal/app/server/repository"
)
```

**Naming & Structure**:

- Use camelCase for variables/functions, PascalCase for exported identifiers
- Organize code into handler/model/repository layers
- Prefer `samber/lo` for functional utilities
- Use `jinzhu/copier` for struct copying
- Use `gorm` for database operations
- Use `echo/v5` for HTTP routing

**Error Handling**:

- Return errors early
- Wrap errors with context
- Use `utils.NewError()` for API responses
- Validate requests with `validateRequest()` middleware

### TypeScript/React Code Style

**Imports** (sorted by Prettier):

- React imports first
- Third-party libraries
- Local absolute imports (@/)
- Local relative imports
- Asset files (svg/png) last

**Formatting**:

- 2-space indentation
- 120-character line length
- Single quotes for strings
- Trailing commas enabled
- Prettier configured to sort imports automatically

**Linting**:

- ESLint with typescript-eslint
- React hooks plugin enabled
- React refresh plugin enabled

**Frontend Stack**:

- React 19 with TypeScript
- Tailwind CSS v4 for styling
- `clsx` + `tailwind-merge` (cn utility)
- TanStack Query for data fetching
- Radix UI components
- Lucide React icons
- Sonner for toasts
- Next.js themes for dark mode

## Testing Guidelines

### Go Testing

- Test files: `*_test.go` alongside source files
- Use `github.com/stretchr/testify` for assertions
- Test naming: `Test<Struct>_<Method>` or `Test<Function>`
- Integration tests use SQLite in-memory databases

Example single test run:

```bash
go test ./internal/app/server/repository -run TestVideoRepository_Count -v
```

### Frontend Testing

- (No specific test framework configured yet - follow standard React testing patterns when adding tests)

## Git & Commit Guidelines

**Conventional Commits**:

- Use `pnpm commit` for interactive commit messages
- Emoji-enabled commits via cz-git
- Commitlint enforced via Husky hooks

**Branch Naming**:

- Follow conventional patterns (feature/, fix/, docs/, etc.)

## Project Structure

```
media-vault/
├── cmd/server/              # Go entry point
├── internal/
│   ├── app/server/          # API handlers, models, repositories
│   └── pkg/                 # Shared utilities (config, downloader, OSS)
├── pkg/taskqueue/          # Task queue implementation
├── web/login/              # React frontend (pnpm workspace)
├── configs/                # Configuration files
├── .air.toml               # Go live-reload config
├── .husky/                 # Git hooks
└── package.json            # pnpm workspace root
```

## Key Dependencies

### Go

- `github.com/labstack/echo/v5` - HTTP framework
- `gorm.io/gorm` - ORM
- `github.com/samber/lo` - Functional utilities
- `github.com/jinzhu/copier` - Struct copying
- `github.com/spf13/viper` - Configuration
- `github.com/stretchr/testify` - Testing

### Frontend

- `react@19` - UI library
- `@tanstack/react-query` - Data fetching
- `tailwindcss@4` - Styling
- `radix-ui` - UI primitives
- `lucide-react` - Icons

## Notes for Agents

1. **Always run lint/format** before committing: `pnpm lint`, `pnpm format:check`
2. **Test patterns**: When modifying Go code, run relevant tests with `go test ./path/to/package -run TestName`
3. **Code organization**: Follow existing handler/model/repository pattern for Go code
4. **Type safety**: Use TypeScript strictly - avoid `any` types
5. **Imports**: Let Prettier sort imports automatically
6. **Error handling**: Always handle errors explicitly in Go; don't ignore them with `_` unless documented
