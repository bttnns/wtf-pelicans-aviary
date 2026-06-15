# Contributing to the aviary

This repo holds the Pelican Ground School lesson content for
[pelicans.wtf](https://pelicans.wtf). It is mounted into the site as a pinned
git submodule and rendered through a hardened Markdown pipeline.

## How to add or edit a lesson

1. Copy `lessons/_template.md` to `lessons/<your-slug>.md`. (Files starting with
   `_` are excluded from the site and never render, so the template is safe.)
2. Fill in the frontmatter (`slug`, `nav`, `blurb`, `teaches`, `order`) and the
   body. Use plain Markdown plus the directive vocabulary in
   [`SYNTAX.md`](./SYNTAX.md).
3. Run the validator locally:

   ```sh
   node scripts/validate-lessons.mjs
   ```

4. Open a PR. The `validate lessons` GitHub Action runs the same checks on your
   branch. Green check, and a maintainer reviews the voice + accuracy.

To publish: a maintainer bumps the submodule pointer in the main repo and pushes
`release`. The push is the deploy.

## The rules (enforced by CI and the site build)

- **Standard Markdown only.** No raw HTML, no JSX, no imports, no expressions.
  Raw HTML is escaped, never rendered.
- **Directives are the only rich content:** `:::beakman`, `::art{id}`,
  `::promptdemo`, `::ralphloop`, `:::sources`. Unknown directives fail the build.
- **Art** must reference an `id` that matches `^[a-z0-9][a-z0-9-]*$` and already
  exists in the site's `illustrations.json`. New illustrations are generated in
  the main repo, not here.
- **Links:** body prose links are internal only (relative, `#anchors`,
  `pelicans.wtf`, `mailto:`). External `https` links go in `:::sources` only.
  No `http:`, no `javascript:`, no `data:`.
- **No em-dashes or en-dashes (`—` / `–`), ever.** Commas, colons, parentheses,
  or "to" for ranges.
- **Voice:** comedy first, real AI underneath. Dr. Beakman is the host: one or
  two punchy, accurate, funny sentences, never a paragraph.

## Why so strict

Lesson HTML runs in visitors' browsers. The hardening is layered defense in
depth (Markdown-only, escaped HTML, a directive whitelist, and `rehype-sanitize`
as the final transform). The strictness is what lets us accept content PRs
safely. See `SYNTAX.md` for examples.
