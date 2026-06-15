# wtf-pelicans-aviary

Private content for [pelicans.wtf](https://pelicans.wtf), mounted into the site
as a pinned git submodule at `site/src/content/aviary/` and read through an
Astro content collection.

## Layout

- `lessons/*.md`: Pelican Ground School lessons, as **hardened standard
  Markdown** (NOT MDX). Each file carries frontmatter (`slug`, `nav`, `blurb`,
  `teaches`, `order`) and a body in the comedy-pelican voice. The only rich
  content is a fixed safe directive vocabulary (`:::beakman`, `::art{id}`,
  `::promptdemo`, `::ralphloop`, `:::sources`); see `SYNTAX.md`. Files starting
  with `_` (e.g. `lessons/_template.md`) are excluded from the collection and
  never render.
- `SYNTAX.md`: the one-page directive + link-policy cheat-sheet.
- `CONTRIBUTING.md`: how to add or edit a lesson.
- `scripts/validate-lessons.mjs`: zero-dependency validator (run it locally;
  CI runs it on every PR).
- `.github/`: the PR template and the `validate lessons` workflow.

## Hardening (why the strict rules)

Lesson HTML runs in visitors' browsers, so the content layer is locked down as
layered defense in depth, enforced in the site build:

- `.md` only (no JSX, imports, or expressions can exist).
- Raw HTML is escaped, never rendered (no `rehype-raw`).
- Only the whitelisted directives map to vetted components; unknown directives
  fail the build.
- `rehype-sanitize` runs as the final transform with a strict allowlist (no
  `script`/`style`/`iframe`/`object`/`embed`/`form`/`base`, no `on*` handlers,
  no inline `style`).
- Link policy: body prose is internal-only; external `https` lives in
  `:::sources` only; no `http:`/`javascript:`/`data:`.
- No em-dashes or en-dashes anywhere.

## Contributing

Open a PR here (the `validate lessons` Action checks it without the private site
code). To publish, a maintainer bumps the submodule pointer in the main repo and
pushes `release` (the push is the deploy). The build is offline and reproducible:
the submodule is cloned, never fetched at build time.

HARD RULE: no em-dashes or en-dashes anywhere.
