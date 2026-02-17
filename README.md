# Draft Survey Tool

A tool for calculating cargo weight on maritime vessels using the draft survey method (UNECE Code 1992).

## Status
🚧 Under active development

## Tech Stack
- **Backend:** Go + Fiber
- **Frontend:** HTMX + html/template
- **Storage:** JSON files (open source version)

## Project Structure
```
cmd/server/       — entry point
internal/         — core business logic
  calculation/    — draft survey math
  vessel/         — vessel data
  report/         — PDF generation
  storage/        — data persistence
  errors/         — custom errors
  logger/         — logging
web/              — templates and static assets
data/             — runtime data (not committed)
  dictionaries/   — ports, flags (committed)
  surveys/        — survey records (local only)
  temp/           — auto-save drafts (local only)
docs/             — documentation
```

## License
MIT
