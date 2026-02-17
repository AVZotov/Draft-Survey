# Draft Survey Tool — Development Roadmap

**Project:** https://github.com/AVZotov/draft-survey  
**Status:** Phase 0 Complete ✅  
**Last Updated:** 2026-02-17

---

## PHASE 0: Foundation ✅ COMPLETE

### Infrastructure
- [x] Repository created (lowercase naming)
- [x] Project structure with Go standard layout
- [x] Allowlist-style `.gitignore` for security
- [x] GitHub Actions CI (ubuntu-22.04, Go 1.25)
- [x] SSH authentication configured

### Decisions Made
- **Workflow:** GitHub Flow (main + feature branches)
- **Go Version:** 1.25 (minimum compatibility)
- **CI Runner:** ubuntu-22.04 (LTS, reproducible builds)
- **PDF Reports:** Generate on-demand only (no auto-save)
- **Backup Strategy:** Each survey = separate JSON file (max 10-20/month)
- **Temp Files:** Auto-deleted after final report generated

### Open Questions for Phase 1
- [ ] Storage backend for open source version (JSON files vs SQLite)
- [ ] Installer/distribution strategy
- [ ] Custom errors location (`internal/errors/`)

---

## PHASE 1: MVP — Open Source (Offline-First)

**Goal:** Working draft survey calculator with local storage

### 1.1 Core Mathematics Module ⭐ HIGHEST PRIORITY
**Location:** `internal/calculation/`

**Tasks:**
- [ ] Implement UNECE 1992 formulas:
  - Quarter Mean Draft
  - First Trim Correction (FTC)
  - Second Trim Correction (STC)
  - Density Correction
  - Displacement calculation
- [ ] Unit tests with hardcoded "golden" test data
- [ ] Validation logic for input ranges

**Deliverable:** Calculation engine that passes all test cases

---

### 1.2 Vessel Data Module
**Location:** `internal/vessel/`

**Tasks:**
- [ ] Vessel information structure (name, IMO, flag, port)
- [ ] Hydrostatic tables storage/parsing
- [ ] Draft readings input (fore/aft/mid — port & starboard)
- [ ] JSON serialization

**Deliverable:** Vessel data can be saved and loaded

---

### 1.3 User Profile (Local)
**Location:** `internal/storage/` (or new `internal/user/`)

**Tasks:**
- [ ] Surveyor profile (name, position, company, signature)
- [ ] Local storage (no authentication yet)
- [ ] Profile persistence in JSON

**Deliverable:** Surveyor can set up profile once

---

### 1.4 Storage Layer
**Location:** `internal/storage/`

**Decision Point:** JSON files
- JSON: Simple, no dependencies, human-readable
- Interim desision with naming convention: YYYY-MM-DD_IMOxxxxxxx_type.json


**Tasks:**
- [ ] Survey CRUD operations
- [ ] Auto-save drafts (temp files in `data/temp/`)
- [ ] List/search surveys
- [ ] Backup mechanism (copy to `data/backups/`)

**Deliverable:** Surveys persist between app restarts

---

### 1.5 Report Generation
**Location:** `internal/report/`

**Tasks:**
- [ ] PDF generation library selection (gofpdf / unidoc)
- [ ] UNECE-compliant report template
- [ ] Include vessel data, calculations, surveyor signature
- [ ] Save to custom location (user chooses path)

**Deliverable:** Generate professional PDF report

---

### 1.6 Web Interface (HTMX)
**Location:** `web/templates/`, `web/static/`

**Tasks:**
- [ ] Fiber server setup (`cmd/server/main.go`)
- [ ] HTMX templates for:
  - Survey list / search
  - New survey form
  - Draft readings input
  - Calculation results display
  - Report generation trigger
- [ ] CSS framework decision (Tailwind / PicoCSS / none)
- [ ] Form validation (client + server side)

**Deliverable:** Working UI at `localhost:8080`

---

### 1.7 Logging
**Location:** `internal/logger/`

**Tasks:**
- [ ] Structured logging (slog / zap / zerolog)
- [ ] Log to file (`data/app.log`)
- [ ] Log rotation strategy
- [ ] Log levels (DEBUG/INFO/WARN/ERROR)

**Deliverable:** All operations logged for debugging

---

### 1.8 Error Handling
**Location:** `internal/errors/`

**Tasks:**
- [ ] Custom error types for validation
- [ ] Error messages dictionary (for i18n later)
- [ ] User-friendly error display in UI

**Deliverable:** Consistent error handling throughout app

---

### 1.9 Testing
**Tasks:**
- [ ] Unit tests for calculation module (golden data)
- [ ] Integration tests for storage
- [ ] E2E test: create survey → calculate → generate PDF

**Target Coverage:** >80% for `internal/calculation/`

---

### 1.10 Dictionaries (Static Data)
**Location:** `data/dictionaries/`

**Tasks:**
- [ ] Ports list (JSON)
- [ ] Country flags (JSON)
- [ ] Units of measurement
- [ ] Load from files on startup

**Deliverable:** Dropdown lists populated from JSON

---

### 1.11 Installer/Distribution
**Decision Point:** How users install the app?

**Options:**
- Single binary (Go strength)
- Installer (creates directories, shortcuts)
- Docker image (for server deployment)

**Tasks:**
- [ ] Build script / Makefile
- [ ] Cross-compilation (Windows / macOS / Linux)
- [ ] First-run setup (create directories)

---

## PHASE 2: Commercial Version (Cloud-Sync)

**Goal:** Multi-user system with central repository

### 2.1 Authentication
- [ ] User login/password
- [ ] Long-lived tokens (2 months offline support)
- [ ] Session management

### 2.2 Backend Server
**Tech Stack:** Go + Fiber + PostgreSQL

- [ ] REST API design
- [ ] Database schema (users, surveys, sync state)
- [ ] UUID-based IDs (not auto-increment)

### 2.3 Offline-First Sync
- [ ] Conflict resolution strategy
- [ ] Delta sync (only changed records)
- [ ] Sync status indicator in UI

### 2.4 User Roles
- [ ] Surveyor (create surveys)
- [ ] Coordinator (manage ports/dictionaries)
- [ ] Admin (user management)

### 2.5 Dictionary Management
- [ ] CRUD for ports (coordinators only)
- [ ] IANA timezone selection
- [ ] Sync dictionaries to clients

### 2.6 Deployment
- [ ] Docker Compose (app + PostgreSQL)
- [ ] Terraform (infrastructure as code)
- [ ] Ansible (configuration management)
- [ ] Self-hosted instructions
- [ ] Security hardening

---

## PHASE 3: Polish & Scale

### 3.1 Internationalization (i18n)
- [ ] Russian
- [ ] English
- [ ] Other languages TBD

### 3.2 Performance
- [ ] Optimize for low-end laptops
- [ ] Caching strategy
- [ ] Large dataset handling

### 3.3 Advanced Features
- [ ] Multiple vessel types
- [ ] Historical data analysis
- [ ] Export to Excel
- [ ] Digital signatures (e-signing PDFs)

---

## Development Strategy

### Branching
- `main` — stable releases only
- `feature/*` — work on specific features
- PR to `main` when ready

### Versioning
- Semantic versioning (v0.1.0, v0.2.0, v1.0.0)
- Git tags for releases

### CI/CD
- GitHub Actions on every push
- Automated tests
- Build artifacts for releases

---

## Next Steps (Immediate)

1. **Decision:** JSON files vs SQLite for Phase 1 storage
2. **Decision:** PDF library (gofpdf vs unidoc)
3. **Decision:** CSS framework (Tailwind vs minimal)
4. **Start Phase 1.1:** Implement calculation module
   - Read UNECE 1992 code (already in project files)
   - Define Go structs for draft readings
   - Write first test case

---

## Notes

- This is a living document — update as decisions are made
- Check off items as completed
- Add new tasks as they emerge
- Keep focus on MVP (Phase 1) before expanding

**Ready to start Phase 1.1 — Core Mathematics Module!**
